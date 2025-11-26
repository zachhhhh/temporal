package workflows

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// BillingCycleInput is the input for the billing cycle workflow.
type BillingCycleInput struct {
	OrganizationID string
	PeriodStart    time.Time
	PeriodEnd      time.Time
}

// BillingCycleWorkflow runs the monthly billing cycle for an organization.
func BillingCycleWorkflow(ctx workflow.Context, input BillingCycleInput) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting billing cycle", "org_id", input.OrganizationID)

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Minute,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    time.Minute,
			MaximumAttempts:    5,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// Step 1: Aggregate all usage for the billing period
	var usageSummary UsageSummaryOutput
	var a *Activities
	err := workflow.ExecuteActivity(ctx, a.AggregateUsageActivity, AggregateUsageInput{
		OrganizationID: input.OrganizationID,
		PeriodStart:    input.PeriodStart,
		PeriodEnd:      input.PeriodEnd,
	}).Get(ctx, &usageSummary)
	if err != nil {
		return err
	}

	// Step 2: Generate invoice
	var invoiceID string
	err = workflow.ExecuteActivity(ctx, a.GenerateInvoiceActivity, GenerateInvoiceInput{
		OrganizationID: input.OrganizationID,
		PeriodStart:    input.PeriodStart,
		PeriodEnd:      input.PeriodEnd,
		UsageSummary:   usageSummary,
	}).Get(ctx, &invoiceID)
	if err != nil {
		return err
	}

	// Step 3: Report usage to Stripe
	err = workflow.ExecuteActivity(ctx, a.ReportStripeUsageActivity, ReportStripeUsageInput{
		OrganizationID: input.OrganizationID,
		InvoiceID:      invoiceID,
		UsageSummary:   usageSummary,
	}).Get(ctx, nil)
	if err != nil {
		logger.Warn("Failed to report usage to Stripe", "error", err)
	}

	// Step 4: Wait for Stripe to finalize invoice (async)
	_ = workflow.Sleep(ctx, 24*time.Hour)

	// Step 5: Send invoice notification
	err = workflow.ExecuteActivity(ctx, a.SendInvoiceEmailActivity, SendInvoiceEmailInput{
		OrganizationID: input.OrganizationID,
		InvoiceID:      invoiceID,
	}).Get(ctx, nil)
	if err != nil {
		logger.Warn("Failed to send invoice email", "error", err)
	}

	logger.Info("Billing cycle completed", "org_id", input.OrganizationID, "invoice_id", invoiceID)
	return nil
}

// DunningInput is the input for the dunning workflow.
type DunningInput struct {
	OrganizationID string
	InvoiceID      string
}

// DunningWorkflow handles failed payment retry.
func DunningWorkflow(ctx workflow.Context, input DunningInput) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting dunning workflow", "invoice_id", input.InvoiceID)

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 5 * time.Minute,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    time.Minute,
			MaximumAttempts:    3,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var a *Activities
	retrySchedule := []time.Duration{
		3 * 24 * time.Hour,  // Day 3
		7 * 24 * time.Hour,  // Day 7
		14 * 24 * time.Hour, // Day 14
	}

	for attempt, delay := range retrySchedule {
		_ = workflow.Sleep(ctx, delay)

		// Send reminder email
		_ = workflow.ExecuteActivity(ctx, a.SendPaymentReminderActivity, SendPaymentReminderInput{
			OrganizationID: input.OrganizationID,
			InvoiceID:      input.InvoiceID,
			Attempt:        attempt + 1,
		}).Get(ctx, nil)

		// Check if paid
		var paid bool
		err := workflow.ExecuteActivity(ctx, a.CheckInvoicePaidActivity, input.InvoiceID).Get(ctx, &paid)
		if err != nil {
			logger.Warn("Failed to check invoice status", "error", err)
			continue
		}

		if paid {
			logger.Info("Invoice paid", "invoice_id", input.InvoiceID)
			return nil
		}
	}

	// Final: Suspend account
	err := workflow.ExecuteActivity(ctx, a.SuspendAccountActivity, input.OrganizationID).Get(ctx, nil)
	if err != nil {
		return err
	}

	logger.Info("Account suspended due to non-payment", "org_id", input.OrganizationID)
	return nil
}

// UsageAggregationInput is the input for usage aggregation workflow.
type UsageAggregationInput struct {
	PeriodType string // "daily" or "monthly"
	PeriodDate time.Time
}

// UsageAggregationWorkflow aggregates usage data.
func UsageAggregationWorkflow(ctx workflow.Context, input UsageAggregationInput) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting usage aggregation", "period_type", input.PeriodType)

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Minute,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    5 * time.Minute,
			MaximumAttempts:    3,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// Get list of organizations to process
	var orgIDs []string
	var a *Activities
	err := workflow.ExecuteActivity(ctx, a.ListActiveOrganizationsActivity).Get(ctx, &orgIDs)
	if err != nil {
		return err
	}

	// Process each organization
	for _, orgID := range orgIDs {
		err := workflow.ExecuteActivity(ctx, a.AggregateOrgUsageActivity, AggregateOrgUsageInput{
			OrganizationID: orgID,
			PeriodType:     input.PeriodType,
			PeriodDate:     input.PeriodDate,
		}).Get(ctx, nil)
		if err != nil {
			logger.Warn("Failed to aggregate usage for org", "org_id", orgID, "error", err)
		}
	}

	logger.Info("Usage aggregation completed", "period_type", input.PeriodType)
	return nil
}

// Activity input/output types for billing

type AggregateUsageInput struct {
	OrganizationID string
	PeriodStart    time.Time
	PeriodEnd      time.Time
}

type UsageSummaryOutput struct {
	TotalActions       int64
	ActiveStorageGBH   float64
	RetainedStorageGBH float64
}

type GenerateInvoiceInput struct {
	OrganizationID string
	PeriodStart    time.Time
	PeriodEnd      time.Time
	UsageSummary   UsageSummaryOutput
}

type ReportStripeUsageInput struct {
	OrganizationID string
	InvoiceID      string
	UsageSummary   UsageSummaryOutput
}

type SendInvoiceEmailInput struct {
	OrganizationID string
	InvoiceID      string
}

type SendPaymentReminderInput struct {
	OrganizationID string
	InvoiceID      string
	Attempt        int
}

type AggregateOrgUsageInput struct {
	OrganizationID string
	PeriodType     string
	PeriodDate     time.Time
}
