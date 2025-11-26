# Pricing & Plans

## Plan Tiers

| Plan             | Base Price | Actions Included | Active Storage | Retained Storage |
| ---------------- | ---------- | ---------------- | -------------- | ---------------- |
| Free             | $0/mo      | 100K             | 0.1 GB         | 4 GB             |
| Essential        | $100/mo    | 1M               | 1 GB           | 40 GB            |
| Business         | $500/mo    | 2.5M             | 2.5 GB         | 100 GB           |
| Enterprise       | Custom     | 10M              | 10 GB          | 400 GB           |
| Mission Critical | Custom     | Custom           | Custom         | Custom           |

## Overage Pricing

### Actions (per million)

| Tier    | Price |
| ------- | ----- |
| 0-5M    | $50   |
| 5-10M   | $45   |
| 10-20M  | $40   |
| 20-50M  | $35   |
| 50-100M | $30   |
| 100M+   | $25   |

### Storage

| Type             | Price        |
| ---------------- | ------------ |
| Active Storage   | $0.042/GBh   |
| Retained Storage | $0.00105/GBh |

### Storage Conversion

- 1 GB = 744 GBh (per month)
- Active: $0.042 × 744 = ~$31.25/GB/month
- Retained: $0.00105 × 744 = ~$0.78/GB/month

## Add-ons

| Feature           | Price    | Included In                      |
| ----------------- | -------- | -------------------------------- |
| SAML SSO          | Included | Business+                        |
| SCIM              | $500/mo  | Enterprise+ (or Business add-on) |
| Multi-region HA   | 2x usage | All plans                        |
| Dedicated Support | Custom   | Enterprise+                      |

## Billing Cycle

- Monthly billing on the 1st
- Prorated for partial months
- Usage calculated in UTC
- Invoices due NET 30

## Payment Methods

- Credit card (Stripe)
- ACH bank transfer (Enterprise)
- Wire transfer (Enterprise)
- Temporal Credits (prepaid)

## Credits

### Commitment Pricing

| Commitment | Discount |
| ---------- | -------- |
| 1 year     | 10%      |
| 2 years    | 15%      |
| 3 years    | 20%      |

### Credit Application

- Credits apply to: Actions, Storage, Plan fees
- Credits do NOT apply to: Add-ons
- Credits expire after commitment period

## Example Calculations

### Small Workload (Essential Plan)

```
Base plan:                    $100
Actions: 2M (1M overage)      $50
Active Storage: 0.5 GBh       $0 (included)
Retained Storage: 20 GBh      $0 (included)
─────────────────────────────────
Total:                        $150/month
```

### Medium Workload (Business Plan)

```
Base plan:                    $500
Actions: 10M (7.5M overage)
  - 5M × $50/M = $250
  - 2.5M × $45/M = $112.50
Active Storage: 5 GB          $156 (2.5 GB overage)
Retained Storage: 200 GB      $78 (100 GB overage)
─────────────────────────────────
Total:                        $1,096.50/month
```
