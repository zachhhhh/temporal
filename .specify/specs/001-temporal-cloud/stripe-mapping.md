# Stripe Product Mapping

## Overview

Maps internal plan tiers and usage metrics to Stripe Products and Prices for billing.

## Product Structure

```
Temporal Cloud (Product)
├── Essential Plan (Price - recurring)
├── Business Plan (Price - recurring)
├── Enterprise Plan (Price - recurring)
├── Actions (Price - metered)
├── Active Storage (Price - metered)
└── Retained Storage (Price - metered)
```

## Price ID Mapping

### Production Environment

| Item             | Stripe Price ID                | Type      | Amount       |
| ---------------- | ------------------------------ | --------- | ------------ |
| Essential Plan   | `price_essential_monthly_prod` | Recurring | $100/mo      |
| Business Plan    | `price_business_monthly_prod`  | Recurring | $500/mo      |
| Actions (per M)  | `price_actions_metered_prod`   | Metered   | $50/M        |
| Active Storage   | `price_active_storage_prod`    | Metered   | $0.042/GBh   |
| Retained Storage | `price_retained_storage_prod`  | Metered   | $0.00105/GBh |
| SCIM Add-on      | `price_scim_addon_prod`        | Recurring | $500/mo      |

### Test Environment

| Item             | Stripe Price ID                | Type      | Amount       |
| ---------------- | ------------------------------ | --------- | ------------ |
| Essential Plan   | `price_essential_monthly_test` | Recurring | $100/mo      |
| Business Plan    | `price_business_monthly_test`  | Recurring | $500/mo      |
| Actions (per M)  | `price_actions_metered_test`   | Metered   | $50/M        |
| Active Storage   | `price_active_storage_test`    | Metered   | $0.042/GBh   |
| Retained Storage | `price_retained_storage_test`  | Metered   | $0.00105/GBh |
| SCIM Add-on      | `price_scim_addon_test`        | Recurring | $500/mo      |

## Configuration

```go
// config/stripe.go
type StripeConfig struct {
    APIKey     string
    WebhookKey string
    Prices     StripePrices
}

type StripePrices struct {
    EssentialPlan   string `env:"STRIPE_PRICE_ESSENTIAL"`
    BusinessPlan    string `env:"STRIPE_PRICE_BUSINESS"`
    Actions         string `env:"STRIPE_PRICE_ACTIONS"`
    ActiveStorage   string `env:"STRIPE_PRICE_ACTIVE_STORAGE"`
    RetainedStorage string `env:"STRIPE_PRICE_RETAINED_STORAGE"`
    SCIMAddon       string `env:"STRIPE_PRICE_SCIM"`
}

func (c *StripeConfig) GetPlanPrice(plan PlanTier) string {
    switch plan {
    case PlanEssential:
        return c.Prices.EssentialPlan
    case PlanBusiness:
        return c.Prices.BusinessPlan
    case PlanEnterprise:
        return "" // Custom pricing
    default:
        return ""
    }
}
```

## Subscription Creation

```go
func CreateSubscription(ctx context.Context, org *Organization, plan PlanTier) (*stripe.Subscription, error) {
    // Create or get customer
    customer, err := getOrCreateCustomer(ctx, org)
    if err != nil {
        return nil, err
    }

    // Build subscription items
    items := []*stripe.SubscriptionItemsParams{
        // Base plan
        {Price: stripe.String(config.GetPlanPrice(plan))},
        // Metered usage items
        {Price: stripe.String(config.Prices.Actions)},
        {Price: stripe.String(config.Prices.ActiveStorage)},
        {Price: stripe.String(config.Prices.RetainedStorage)},
    }

    // Create subscription
    params := &stripe.SubscriptionParams{
        Customer: stripe.String(customer.ID),
        Items:    items,
        Metadata: map[string]string{
            "organization_id": org.ID,
            "plan":            string(plan),
        },
    }

    return subscription.New(params)
}
```

## Usage Reporting

```go
func ReportUsage(ctx context.Context, sub *Subscription, usage *UsageRecord) error {
    // Report actions
    if usage.ActionCount > 0 {
        _, err := usagerecord.New(&stripe.UsageRecordParams{
            SubscriptionItem: stripe.String(sub.ActionsItemID),
            Quantity:         stripe.Int64(usage.ActionCount),
            Timestamp:        stripe.Int64(usage.PeriodEnd.Unix()),
            Action:           stripe.String("increment"),
        })
        if err != nil {
            return fmt.Errorf("failed to report actions: %w", err)
        }
    }

    // Report active storage (in milli-GBh for precision)
    if usage.ActiveStorageGBh > 0 {
        milliGBh := int64(usage.ActiveStorageGBh * 1000)
        _, err := usagerecord.New(&stripe.UsageRecordParams{
            SubscriptionItem: stripe.String(sub.ActiveStorageItemID),
            Quantity:         stripe.Int64(milliGBh),
            Timestamp:        stripe.Int64(usage.PeriodEnd.Unix()),
            Action:           stripe.String("increment"),
        })
        if err != nil {
            return fmt.Errorf("failed to report active storage: %w", err)
        }
    }

    // Report retained storage
    if usage.RetainedStorageGBh > 0 {
        milliGBh := int64(usage.RetainedStorageGBh * 1000)
        _, err := usagerecord.New(&stripe.UsageRecordParams{
            SubscriptionItem: stripe.String(sub.RetainedStorageItemID),
            Quantity:         stripe.Int64(milliGBh),
            Timestamp:        stripe.Int64(usage.PeriodEnd.Unix()),
            Action:           stripe.String("increment"),
        })
        if err != nil {
            return fmt.Errorf("failed to report retained storage: %w", err)
        }
    }

    return nil
}
```

## Webhook Handling

### Supported Events

| Event                           | Handler                  |
| ------------------------------- | ------------------------ |
| `customer.subscription.created` | Sync subscription to DB  |
| `customer.subscription.updated` | Update plan/status       |
| `customer.subscription.deleted` | Mark as canceled         |
| `invoice.paid`                  | Update invoice status    |
| `invoice.payment_failed`        | Trigger dunning workflow |
| `invoice.finalized`             | Store invoice details    |

### Webhook Handler

```go
func HandleStripeWebhook(w http.ResponseWriter, r *http.Request) {
    payload, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "read error", http.StatusBadRequest)
        return
    }

    event, err := webhook.ConstructEvent(
        payload,
        r.Header.Get("Stripe-Signature"),
        config.WebhookKey,
    )
    if err != nil {
        http.Error(w, "signature error", http.StatusBadRequest)
        return
    }

    switch event.Type {
    case "customer.subscription.updated":
        var sub stripe.Subscription
        json.Unmarshal(event.Data.Raw, &sub)
        handleSubscriptionUpdated(r.Context(), &sub)

    case "invoice.paid":
        var inv stripe.Invoice
        json.Unmarshal(event.Data.Raw, &inv)
        handleInvoicePaid(r.Context(), &inv)

    case "invoice.payment_failed":
        var inv stripe.Invoice
        json.Unmarshal(event.Data.Raw, &inv)
        handlePaymentFailed(r.Context(), &inv)
    }

    w.WriteHeader(http.StatusOK)
}
```

### Plan Change from Stripe Portal

When customer changes plan via Stripe Billing Portal:

```go
func handleSubscriptionUpdated(ctx context.Context, sub *stripe.Subscription) error {
    orgID := sub.Metadata["organization_id"]

    // Determine new plan from price
    var newPlan PlanTier
    for _, item := range sub.Items.Data {
        switch item.Price.ID {
        case config.Prices.EssentialPlan:
            newPlan = PlanEssential
        case config.Prices.BusinessPlan:
            newPlan = PlanBusiness
        }
    }

    // Update internal subscription
    return store.UpdateSubscription(ctx, orgID, UpdateSubscriptionInput{
        Plan:                newPlan,
        Status:              mapStripeStatus(sub.Status),
        StripeSubscriptionID: sub.ID,
    })
}
```

## Testing

### Test Mode Setup

```bash
# Set test API key
export STRIPE_API_KEY=sk_test_xxx

# Create test products/prices
stripe products create --name="Temporal Cloud"
stripe prices create \
  --product=prod_xxx \
  --unit-amount=10000 \
  --currency=usd \
  --recurring[interval]=month \
  --nickname="Essential Plan Test"
```

### Webhook Testing

```bash
# Forward webhooks to local
stripe listen --forward-to localhost:8081/webhooks/stripe

# Trigger test events
stripe trigger invoice.paid
stripe trigger customer.subscription.updated
```
