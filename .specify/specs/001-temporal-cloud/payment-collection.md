# Payment Collection & Dunning

## Payment Lifecycle

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Invoice   │────▶│   Payment   │────▶│   Paid      │
│   Created   │     │   Attempted │     │             │
└─────────────┘     └─────────────┘     └─────────────┘
                          │
                          │ Failed
                          ▼
                    ┌─────────────┐
                    │   Retry     │──┐
                    │   (Dunning) │  │ (up to 4 times)
                    └─────────────┘◀─┘
                          │
                          │ All retries failed
                          ▼
                    ┌─────────────┐
                    │   Suspended │
                    └─────────────┘
                          │
                          │ 30 days
                          ▼
                    ┌─────────────┐
                    │   Canceled  │
                    └─────────────┘
```

## Invoice Generation

### Monthly Invoice Workflow

```go
func GenerateMonthlyInvoices(ctx workflow.Context, month time.Time) error {
    // Get all active subscriptions
    var subs []Subscription
    workflow.ExecuteActivity(ctx, GetActiveSubscriptions).Get(ctx, &subs)

    for _, sub := range subs {
        // Generate invoice for each
        workflow.ExecuteChildWorkflow(ctx, GenerateInvoiceWorkflow, sub, month)
    }

    return nil
}

func GenerateInvoiceWorkflow(ctx workflow.Context, sub Subscription, month time.Time) error {
    // Calculate usage
    var usage MonthlyUsage
    workflow.ExecuteActivity(ctx, CalculateUsage, sub.OrgID, month).Get(ctx, &usage)

    // Create invoice
    var invoice Invoice
    workflow.ExecuteActivity(ctx, CreateInvoice, CreateInvoiceInput{
        Subscription: sub,
        Usage:        usage,
        Month:        month,
    }).Get(ctx, &invoice)

    // Finalize in Stripe
    workflow.ExecuteActivity(ctx, FinalizeStripeInvoice, invoice.StripeInvoiceID)

    // Send invoice email
    workflow.ExecuteActivity(ctx, SendInvoiceEmail, invoice)

    return nil
}
```

## Payment Processing

### Automatic Payment

Stripe automatically charges the default payment method when invoice is finalized.

### Manual Payment

For enterprise customers with NET 30 terms:

1. Invoice marked as "send_invoice"
2. Customer pays via bank transfer
3. Finance team marks as paid manually

## Dunning Process

### Retry Schedule

| Attempt | Day | Action                           |
| ------- | --- | -------------------------------- |
| 1       | 0   | Initial charge                   |
| 2       | 3   | Retry + email reminder           |
| 3       | 7   | Retry + email warning            |
| 4       | 14  | Final retry + email final notice |
| -       | 21  | Account suspended                |
| -       | 51  | Account canceled + data deletion |

### Dunning Workflow

```go
func DunningWorkflow(ctx workflow.Context, invoiceID string) error {
    retrySchedule := []time.Duration{
        3 * 24 * time.Hour,
        4 * 24 * time.Hour,
        7 * 24 * time.Hour,
    }

    for i, delay := range retrySchedule {
        // Wait for next retry
        workflow.Sleep(ctx, delay)

        // Check if paid in the meantime
        var invoice Invoice
        workflow.ExecuteActivity(ctx, GetInvoice, invoiceID).Get(ctx, &invoice)
        if invoice.Status == "paid" {
            return nil
        }

        // Retry payment
        var result PaymentResult
        workflow.ExecuteActivity(ctx, RetryPayment, invoiceID).Get(ctx, &result)

        if result.Success {
            workflow.ExecuteActivity(ctx, SendPaymentSuccessEmail, invoice)
            return nil
        }

        // Send appropriate reminder
        workflow.ExecuteActivity(ctx, SendDunningEmail, DunningEmailInput{
            Invoice: invoice,
            Attempt: i + 2,
        })
    }

    // All retries exhausted - suspend account
    workflow.ExecuteActivity(ctx, SuspendAccount, invoice.OrgID)

    // Wait 30 days then cancel
    workflow.Sleep(ctx, 30*24*time.Hour)

    // Check one more time
    var invoice Invoice
    workflow.ExecuteActivity(ctx, GetInvoice, invoiceID).Get(ctx, &invoice)
    if invoice.Status == "paid" {
        workflow.ExecuteActivity(ctx, ReactivateAccount, invoice.OrgID)
        return nil
    }

    // Cancel and schedule data deletion
    workflow.ExecuteActivity(ctx, CancelSubscription, invoice.OrgID)
    workflow.ExecuteActivity(ctx, ScheduleDataDeletion, invoice.OrgID)

    return nil
}
```

## Email Templates

### Invoice Created

```
Subject: Your Temporal Cloud Invoice for {{month}}

Hi {{name}},

Your invoice for {{month}} is ready.

Amount Due: ${{amount}}
Due Date: {{due_date}}

[View Invoice]  [Pay Now]
```

### Payment Failed (Attempt 1)

```
Subject: Payment failed for your Temporal Cloud invoice

Hi {{name}},

We were unable to charge your payment method for your
{{month}} invoice of ${{amount}}.

Please update your payment method to avoid service interruption.

[Update Payment Method]
```

### Account Suspended

```
Subject: Your Temporal Cloud account has been suspended

Hi {{name}},

Due to an unpaid invoice of ${{amount}}, your Temporal Cloud
account has been suspended.

Your namespaces are no longer accessible. Please pay the
outstanding balance to restore access.

If payment is not received within 30 days, your account and
all data will be permanently deleted.

[Pay Now]
```

## Account Suspension

### What Gets Suspended

- API access disabled
- Workers cannot connect
- Console shows "Account Suspended" banner
- New namespace creation blocked

### What Continues

- Data is retained (for 30 days)
- Audit logs accessible
- Billing portal accessible

### Reactivation

Upon payment:

1. Invoice marked as paid
2. Subscription status → active
3. Access restored within 5 minutes
4. Send "Account Reactivated" email

## Failed Payment Handling

### Common Failure Reasons

| Code                      | Reason             | Action          |
| ------------------------- | ------------------ | --------------- |
| `card_declined`           | Generic decline    | Update card     |
| `insufficient_funds`      | Not enough balance | Retry later     |
| `expired_card`            | Card expired       | Update card     |
| `authentication_required` | 3DS needed         | Customer action |

### Smart Retry

```go
func shouldRetry(failureCode string) bool {
    retryable := map[string]bool{
        "insufficient_funds":   true,
        "processing_error":     true,
        "rate_limit":           true,
    }
    return retryable[failureCode]
}
```

## Collections (Enterprise)

For enterprise customers with large outstanding balances:

1. Account Manager notified at 14 days overdue
2. Finance team reviews at 21 days
3. Collections process at 30 days
4. Legal action consideration at 60 days

## Metrics

| Metric                          | Target |
| ------------------------------- | ------ |
| On-time payment rate            | >95%   |
| Dunning recovery rate           | >70%   |
| Suspension rate                 | <2%    |
| Cancellation rate (non-payment) | <0.5%  |

## Fraud Prevention

### Red Flags

- Multiple failed payment methods
- Rapid usage spike after signup
- Disposable email domains
- High-risk countries

### Actions

- Manual review before trial conversion
- Require upfront payment for high-risk
- Rate limit API key creation
