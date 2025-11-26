# Service Level Agreements (SLA)

## Availability

### Definitions

- **Unavailable**: When valid API requests return 5xx errors or time out for > 1 minute continuously.
- **Maintenance**: Planned downtime with 24h notice (excluded from SLA).

### Service Tiers

| Plan       | Availability Target        | Financial Credit                                 |
| ---------- | -------------------------- | ------------------------------------------------ |
| Essential  | 99.9%                      | 10% if < 99.9%, 25% if < 99.0%                   |
| Business   | 99.9%                      | 10% if < 99.9%, 25% if < 99.0%                   |
| Enterprise | 99.99% (with Multi-Region) | 10% if < 99.99%, 30% if < 99.9%, 100% if < 99.0% |

### Calculation

`Availability = (Total Minutes - Downtime Minutes) / Total Minutes`

## Latency SLO

Target response times (P99) measured at the load balancer:

- **StartWorkflow**: < 200ms
- **SignalWorkflow**: < 200ms
- **QueryWorkflow**: < 500ms (excluding worker processing time)
- **History Events**: < 100ms

## Support Response Times

| Priority      | Essential       | Business        | Enterprise     |
| ------------- | --------------- | --------------- | -------------- |
| P0 (Critical) | N/A             | 1 hour          | 15 mins        |
| P1 (High)     | N/A             | 4 hours         | 1 hour         |
| P2 (Normal)   | 2 business days | 1 business day  | 4 hours        |
| P3 (Low)      | Best effort     | 2 business days | 1 business day |

## Exclusions

The SLA does not apply to:

1. Beta/Experimental features
2. Client-side errors (bad input, rate limiting)
3. Force Majeure events
4. Planned maintenance
5. Suspended accounts (due to non-payment)

## Claim Process

1. Customer must submit claim via Support Portal within 30 days.
2. Claim must include timestamps and error logs.
3. Temporal validates against server-side logs.
4. Credits applied to next invoice.
