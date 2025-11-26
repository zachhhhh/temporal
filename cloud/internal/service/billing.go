package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"go.temporal.io/cloud/internal/config"
	"go.temporal.io/cloud/internal/repository"
	"go.temporal.io/server/common/log"
)

// BillingService handles billing business logic.
type BillingService struct {
	repos        *repository.Repositories
	stripeConfig config.StripeConfig
	logger       log.Logger
}

// NewBillingService creates a new billing service.
func NewBillingService(repos *repository.Repositories, stripeCfg config.StripeConfig, logger log.Logger) *BillingService {
	return &BillingService{repos: repos, stripeConfig: stripeCfg, logger: logger}
}

// GetSubscription retrieves the subscription for an organization.
func (s *BillingService) GetSubscription(ctx context.Context, orgID uuid.UUID) (*repository.Subscription, error) {
	return s.repos.Subscriptions.GetByOrganizationID(ctx, orgID)
}

// UpdateSubscriptionInput is the input for updating a subscription.
type UpdateSubscriptionInput struct {
	OrganizationID uuid.UUID
	Plan           string
}

// UpdateSubscription updates the subscription plan.
func (s *BillingService) UpdateSubscription(ctx context.Context, input *UpdateSubscriptionInput) (*repository.Subscription, error) {
	sub, err := s.repos.Subscriptions.GetByOrganizationID(ctx, input.OrganizationID)
	if err != nil {
		return nil, err
	}
	if sub == nil {
		return nil, fmt.Errorf("subscription not found")
	}

	// Validate plan transition
	if !isValidPlanTransition(sub.Plan, input.Plan) {
		return nil, fmt.Errorf("invalid plan transition from %s to %s", sub.Plan, input.Plan)
	}

	// Update plan limits
	actions, activeGB, retainedGB := repository.GetPlanLimits(input.Plan)
	sub.Plan = input.Plan
	sub.ActionsIncluded = actions
	sub.ActiveStorageGB = activeGB
	sub.RetainedStorageGB = retainedGB

	// TODO: Update Stripe subscription if exists

	if err := s.repos.Subscriptions.Update(ctx, sub); err != nil {
		return nil, fmt.Errorf("failed to update subscription: %w", err)
	}

	return sub, nil
}

// UsageSummary represents usage data for a period.
type UsageSummary struct {
	OrganizationID     uuid.UUID
	PeriodStart        time.Time
	PeriodEnd          time.Time
	TotalActions       int64
	ActiveStorageGBH   decimal.Decimal
	RetainedStorageGBH decimal.Decimal
	NamespaceUsage     []NamespaceUsageSummary
	ActionBreakdown    ActionBreakdown
}

// NamespaceUsageSummary represents usage for a single namespace.
type NamespaceUsageSummary struct {
	NamespaceID        string
	NamespaceName      string
	TotalActions       int64
	ActiveStorageGBH   decimal.Decimal
	RetainedStorageGBH decimal.Decimal
}

// ActionBreakdown represents usage breakdown by action type.
type ActionBreakdown struct {
	WorkflowStarted       int64
	WorkflowReset         int64
	TimerStarted          int64
	SignalSent            int64
	QueryReceived         int64
	UpdateReceived        int64
	ActivityStarted       int64
	ActivityHeartbeat     int64
	LocalActivityBatch    int64
	ChildWorkflowStarted  int64
	ScheduleExecution     int64
	NexusOperation        int64
	SearchAttributeUpsert int64
	SideEffectRecorded    int64
	WorkflowExported      int64
}

// GetUsage retrieves usage data for an organization.
func (s *BillingService) GetUsage(ctx context.Context, orgID uuid.UUID, start, end time.Time) (*UsageSummary, error) {
	summary, err := s.repos.Usage.GetSummaryByOrganization(ctx, orgID, start, end)
	if err != nil {
		return nil, fmt.Errorf("failed to get usage summary: %w", err)
	}

	return &UsageSummary{
		OrganizationID:     orgID,
		PeriodStart:        start,
		PeriodEnd:          end,
		TotalActions:       summary.ActionCount,
		ActiveStorageGBH:   summary.ActiveStorageGBH,
		RetainedStorageGBH: summary.RetainedStorageGBH,
		ActionBreakdown: ActionBreakdown{
			WorkflowStarted:       summary.WorkflowStarted,
			WorkflowReset:         summary.WorkflowReset,
			TimerStarted:          summary.TimerStarted,
			SignalSent:            summary.SignalSent,
			QueryReceived:         summary.QueryReceived,
			UpdateReceived:        summary.UpdateReceived,
			ActivityStarted:       summary.ActivityStarted,
			ActivityHeartbeat:     summary.ActivityHeartbeat,
			LocalActivityBatch:    summary.LocalActivityBatch,
			ChildWorkflowStarted:  summary.ChildWorkflowStarted,
			ScheduleExecution:     summary.ScheduleExecution,
			NexusOperation:        summary.NexusOperation,
			SearchAttributeUpsert: summary.SearchAttributeUpsert,
			SideEffectRecorded:    summary.SideEffectRecorded,
			WorkflowExported:      summary.WorkflowExported,
		},
	}, nil
}

// GetUsageByNamespace retrieves usage for a specific namespace.
func (s *BillingService) GetUsageByNamespace(ctx context.Context, namespaceID string, start, end time.Time) (*NamespaceUsageSummary, error) {
	summary, err := s.repos.Usage.GetSummaryByNamespace(ctx, namespaceID, start, end)
	if err != nil {
		return nil, fmt.Errorf("failed to get namespace usage: %w", err)
	}

	return &NamespaceUsageSummary{
		NamespaceID:        namespaceID,
		TotalActions:       summary.ActionCount,
		ActiveStorageGBH:   summary.ActiveStorageGBH,
		RetainedStorageGBH: summary.RetainedStorageGBH,
	}, nil
}

// ListInvoices lists invoices for an organization.
func (s *BillingService) ListInvoices(ctx context.Context, orgID uuid.UUID, limit, offset int) ([]*repository.Invoice, error) {
	return s.repos.Invoices.ListByOrganization(ctx, orgID, limit, offset)
}

// GetInvoice retrieves a specific invoice.
func (s *BillingService) GetInvoice(ctx context.Context, invoiceID uuid.UUID) (*repository.Invoice, error) {
	return s.repos.Invoices.GetByID(ctx, invoiceID)
}

// CreditBalance represents an organization's credit balance.
type CreditBalance struct {
	OrganizationID uuid.UUID
	BalanceCents   int64
}

// GetCreditBalance retrieves the credit balance for an organization.
func (s *BillingService) GetCreditBalance(ctx context.Context, orgID uuid.UUID) (*CreditBalance, error) {
	// TODO: Implement credit balance retrieval
	return &CreditBalance{
		OrganizationID: orgID,
		BalanceCents:   0,
	}, nil
}

// RecordUsage records usage for billing.
func (s *BillingService) RecordUsage(ctx context.Context, record *repository.UsageRecord) error {
	return s.repos.Usage.Create(ctx, record)
}

// GenerateInvoice generates an invoice for a billing period.
func (s *BillingService) GenerateInvoice(ctx context.Context, orgID uuid.UUID, periodStart, periodEnd time.Time) (*repository.Invoice, error) {
	// Get subscription
	sub, err := s.repos.Subscriptions.GetByOrganizationID(ctx, orgID)
	if err != nil {
		return nil, fmt.Errorf("failed to get subscription: %w", err)
	}

	// Get usage
	usage, err := s.repos.Usage.GetSummaryByOrganization(ctx, orgID, periodStart, periodEnd)
	if err != nil {
		return nil, fmt.Errorf("failed to get usage: %w", err)
	}

	// Calculate charges
	lineItems := s.calculateLineItems(sub, usage)
	lineItemsJSON, _ := json.Marshal(lineItems)

	var subtotal int64
	for _, item := range lineItems {
		subtotal += item.AmountCents
	}

	// Generate invoice number
	invoiceNumber, err := s.repos.Invoices.GenerateInvoiceNumber(ctx, orgID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate invoice number: %w", err)
	}

	invoice := &repository.Invoice{
		OrganizationID: orgID,
		InvoiceNumber:  invoiceNumber,
		PeriodStart:    periodStart,
		PeriodEnd:      periodEnd,
		LineItems:      lineItemsJSON,
		SubtotalCents:  subtotal,
		TotalCents:     subtotal,
		Status:         "draft",
		DueAt:          sql.NullTime{Time: periodEnd.AddDate(0, 0, 30), Valid: true},
	}

	if err := s.repos.Invoices.Create(ctx, invoice); err != nil {
		return nil, fmt.Errorf("failed to create invoice: %w", err)
	}

	return invoice, nil
}

// InvoiceLineItem represents a line item on an invoice.
type InvoiceLineItem struct {
	Description    string  `json:"description"`
	Quantity       float64 `json:"quantity"`
	Unit           string  `json:"unit"`
	UnitPriceCents int64   `json:"unit_price_cents"`
	AmountCents    int64   `json:"amount_cents"`
}

func (s *BillingService) calculateLineItems(sub *repository.Subscription, usage *repository.UsageRecord) []InvoiceLineItem {
	var items []InvoiceLineItem

	// Base plan fee
	planFee := getPlanFee(sub.Plan)
	if planFee > 0 {
		items = append(items, InvoiceLineItem{
			Description:    fmt.Sprintf("%s Plan", sub.Plan),
			Quantity:       1,
			Unit:           "month",
			UnitPriceCents: planFee,
			AmountCents:    planFee,
		})
	}

	// Action overage
	actionsOverage := usage.ActionCount - sub.ActionsIncluded
	if actionsOverage > 0 {
		pricePerMillion := getActionPrice(actionsOverage)
		millions := float64(actionsOverage) / 1000000.0
		amount := int64(millions * float64(pricePerMillion))
		items = append(items, InvoiceLineItem{
			Description:    "Actions (overage)",
			Quantity:       millions,
			Unit:           "million",
			UnitPriceCents: pricePerMillion,
			AmountCents:    amount,
		})
	}

	// Active storage overage
	activeOverage := usage.ActiveStorageGBH.Sub(sub.ActiveStorageGB.Mul(decimal.NewFromInt(744)))
	if activeOverage.IsPositive() {
		pricePerGBH := int64(4200) // $0.042 per GBh in cents * 100
		amount := activeOverage.Mul(decimal.NewFromInt(pricePerGBH)).Div(decimal.NewFromInt(100)).IntPart()
		items = append(items, InvoiceLineItem{
			Description:    "Active Storage (overage)",
			Quantity:       activeOverage.InexactFloat64(),
			Unit:           "GB-hour",
			UnitPriceCents: pricePerGBH / 100,
			AmountCents:    amount,
		})
	}

	// Retained storage overage
	retainedOverage := usage.RetainedStorageGBH.Sub(sub.RetainedStorageGB.Mul(decimal.NewFromInt(744)))
	if retainedOverage.IsPositive() {
		pricePerGBH := int64(105) // $0.00105 per GBh in cents * 100
		amount := retainedOverage.Mul(decimal.NewFromInt(pricePerGBH)).Div(decimal.NewFromInt(100)).IntPart()
		items = append(items, InvoiceLineItem{
			Description:    "Retained Storage (overage)",
			Quantity:       retainedOverage.InexactFloat64(),
			Unit:           "GB-hour",
			UnitPriceCents: pricePerGBH / 100,
			AmountCents:    amount,
		})
	}

	return items
}

func isValidPlanTransition(from, to string) bool {
	planOrder := map[string]int{
		"free": 0, "essentials": 1, "business": 2, "enterprise": 3, "mission_critical": 4,
	}
	return planOrder[to] >= planOrder[from]
}

func getPlanFee(plan string) int64 {
	fees := map[string]int64{
		"free": 0, "essentials": 10000, "business": 50000, "enterprise": 0, "mission_critical": 0,
	}
	return fees[plan]
}

func getActionPrice(overage int64) int64 {
	millions := overage / 1000000
	switch {
	case millions < 5:
		return 5000 // $50/M
	case millions < 10:
		return 4500 // $45/M
	case millions < 20:
		return 4000 // $40/M
	case millions < 50:
		return 3500 // $35/M
	case millions < 100:
		return 3000 // $30/M
	default:
		return 2500 // $25/M
	}
}
