# Billing Reconciliation Logic

## Problem Statement

Usage is aggregated hourly. Invoices are monthly. Stripe subscriptions handle the fixed fee, but usage must be reported accurately and idempotently.

## Architecture

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│  Metering   │────▶│   Usage     │────▶│   Stripe    │
│ Interceptor │     │   Store     │     │   Billing   │
└─────────────┘     └─────────────┘     └─────────────┘
      │                   │                   │
      ▼                   ▼                   ▼
  Real-time          Hourly              Monthly
  counting         aggregation          invoicing
```

## Algorithm: ReportUsageWorkflow

### Trigger

- Runs every 1 hour per organization
- Scheduled via Temporal Schedule

### Workflow Definition

```go
func ReportUsageWorkflow(ctx workflow.Context, orgID string) error {
    // Step 1: Get usage for last hour
    var usage UsageSummary
    err := workflow.ExecuteActivity(ctx, GetHourlyUsage, orgID).Get(ctx, &usage)
    if err != nil {
        return err
    }

    // Step 2: Check if already reported (idempotency)
    var reported bool
    err = workflow.ExecuteActivity(ctx, CheckUsageReported, usage.ID).Get(ctx, &reported)
    if err != nil {
        return err
    }
    if reported {
        return nil // Already processed
    }

    // Step 3: Report to Stripe
    var stripeRecordID string
    err = workflow.ExecuteActivity(ctx, ReportToStripe, ReportUsageInput{
        OrgID:     orgID,
        Actions:   usage.ActionCount,
        ActiveGBh: usage.ActiveStorageGBh,
        RetainedGBh: usage.RetainedStorageGBh,
        Timestamp: usage.PeriodEnd.Unix(),
    }).Get(ctx, &stripeRecordID)
    if err != nil {
        return err
    }

    // Step 4: Mark as reported
    return workflow.ExecuteActivity(ctx, MarkUsageReported, MarkReportedInput{
        UsageID:        usage.ID,
        StripeRecordID: stripeRecordID,
    }).Get(ctx, nil)
}
```

### Stripe Reporting

```go
func ReportToStripe(ctx context.Context, input ReportUsageInput) (string, error) {
    // Use 'increment' action for idempotency
    // Stripe deduplicates based on subscription_item + timestamp

    // Report actions
    if input.Actions > 0 {
        _, err := stripe.UsageRecord.New(&stripe.UsageRecordParams{
            SubscriptionItem: stripe.String(input.ActionsItemID),
            Quantity:         stripe.Int64(input.Actions),
            Timestamp:        stripe.Int64(input.Timestamp),
            Action:           stripe.String("increment"),
        })
        if err != nil {
            return "", err
        }
    }

    // Report active storage
    if input.ActiveGBh > 0 {
        _, err := stripe.UsageRecord.New(&stripe.UsageRecordParams{
            SubscriptionItem: stripe.String(input.ActiveStorageItemID),
            Quantity:         stripe.Int64(int64(input.ActiveGBh * 1000)), // milliunits
            Timestamp:        stripe.Int64(input.Timestamp),
            Action:           stripe.String("increment"),
        })
        if err != nil {
            return "", err
        }
    }

    return recordID, nil
}
```

## Proration Logic

### Plan Upgrades

- **Timing**: Immediate
- **Calculation**: `(DaysRemaining / DaysInMonth) * (NewPrice - OldPrice)`
- **Included quotas**: Full month's quota from upgrade date

```go
func CalculateUpgradeProration(upgrade UpgradeRequest) int64 {
    daysRemaining := daysUntilEndOfMonth(time.Now())
    daysInMonth := daysInCurrentMonth()

    priceDiff := PlanPrices[upgrade.NewPlan] - PlanPrices[upgrade.OldPlan]
    prorated := (float64(daysRemaining) / float64(daysInMonth)) * float64(priceDiff)

    return int64(prorated)
}
```

### Plan Downgrades

- **Timing**: Next billing cycle
- **Calculation**: No immediate charge
- **Included quotas**: Current plan until end of cycle

## Invoice Generation

### Monthly Invoice Workflow

```go
func GenerateInvoiceWorkflow(ctx workflow.Context, orgID string, month time.Time) error {
    // Step 1: Get subscription
    var sub Subscription
    workflow.ExecuteActivity(ctx, GetSubscription, orgID).Get(ctx, &sub)

    // Step 2: Get usage summary
    var usage MonthlyUsage
    workflow.ExecuteActivity(ctx, GetMonthlyUsage, orgID, month).Get(ctx, &usage)

    // Step 3: Calculate line items
    lineItems := []LineItem{}

    // Base plan fee
    lineItems = append(lineItems, LineItem{
        Description: fmt.Sprintf("%s Plan", sub.Plan),
        Amount:      PlanPrices[sub.Plan],
    })

    // Action overage
    if usage.Actions > sub.ActionsIncluded {
        overage := usage.Actions - sub.ActionsIncluded
        cost := calculateTieredActionCost(overage)
        lineItems = append(lineItems, LineItem{
            Description: fmt.Sprintf("Actions Overage (%d)", overage),
            Amount:      cost,
        })
    }

    // Storage overage (similar logic)

    // Step 4: Create invoice
    return workflow.ExecuteActivity(ctx, CreateInvoice, CreateInvoiceInput{
        OrgID:     orgID,
        Month:     month,
        LineItems: lineItems,
    }).Get(ctx, nil)
}
```

### Tiered Action Pricing

```go
func calculateTieredActionCost(actions int64) int64 {
    tiers := []struct {
        UpTo  int64
        Price int64 // per million
    }{
        {5_000_000, 5000},   // $50/M
        {10_000_000, 4500},  // $45/M
        {20_000_000, 4000},  // $40/M
        {50_000_000, 3500},  // $35/M
        {100_000_000, 3000}, // $30/M
        {0, 2500},           // $25/M (unlimited)
    }

    remaining := actions
    totalCost := int64(0)
    prevTier := int64(0)

    for _, tier := range tiers {
        if remaining <= 0 {
            break
        }

        var tierSize int64
        if tier.UpTo == 0 {
            tierSize = remaining
        } else {
            tierSize = tier.UpTo - prevTier
        }

        actionsInTier := min(remaining, tierSize)
        totalCost += (actionsInTier * tier.Price) / 1_000_000
        remaining -= actionsInTier
        prevTier = tier.UpTo
    }

    return totalCost
}
```

## Error Handling

### Stripe API Failures

- Retry with exponential backoff
- Max 5 retries over 1 hour
- Alert on persistent failure
- Manual intervention if still failing

### Usage Data Missing

- Log warning
- Use last known good data
- Flag invoice for review

### Duplicate Prevention

- Idempotency key: `{org_id}:{period_start}:{period_end}`
- Check before processing
- Mark as processed atomically
