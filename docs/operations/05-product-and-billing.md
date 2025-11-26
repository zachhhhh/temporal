# Product, Billing, and Payments

## User journey (0→payment)
- Awareness → signup → workspace/namespace creation → first workflow → scale-up → billing activation → ongoing usage and support.
- Default free tier with limits; guided setup with SDK quickstarts; in-product observability and quota indicators.

## Billing architecture
- Metering service records workflow starts, executions, history size, storage, and egress; events signed and buffered with idempotent sinks.
- Rating engine applies plan rules (free/standard/enterprise), discounts, and credits; entitlements enforced at request time and during scheduling.
- Invoicing via provider (e.g., Stripe) with dual-write to ledger; dunning and payment retries automated; webhooks verified with signatures.
- PCI boundaries: payment processor handles card data; service stores tokens only; segregated subnet and service accounts for billing components.

## Backend contracts
- APIs for metering ingestion, entitlement checks, invoice generation, and customer portal links; all with strict schemas and idempotency keys.
- Audit every billing mutation; reconciliation job compares processor data with internal ledger daily; alert on deltas.

## UI flows
- Signup with SSO/passwordless; workspace wizard; usage dashboard with projected spend and alerts; plan management with upgrade/downgrade paths.
- Support embedded (chat/email links), incident banners, and status page integration.

## Operational policies
- Currency/timezone normalization; tax handling via provider; invoicing schedule aligned to contract; credit memo workflow.
- SLA for billing issues: P1 (blocking) 4h, P2 1d, P3 next sprint. Require reproduction notes and billing logs.

## Data lifecycle
- Keep billing events for 7 years cold storage; PII minimized and tokenized; access via least-privilege roles; backups with separate encryption keys.
