# Temporal Cloud Platform Specification

## 1. Overview

### 1.1 Vision

Transform the open-source Temporal server into a managed cloud service with pay-as-you-go billing, enterprise security, and self-service management.

### 1.2 Success Metrics

| Metric                           | Target       |
| -------------------------------- | ------------ |
| Upstream sync success rate       | 100%         |
| Platform availability            | 99.9%        |
| Invoice accuracy                 | 99.99%       |
| Mean time to provision namespace | < 30 seconds |
| P99 API latency                  | < 200ms      |

### 1.3 Users

- **Platform Operators**: Deploy and manage Temporal Cloud infrastructure
- **Organization Admins**: Manage billing, users, security settings
- **Namespace Admins**: Create namespaces, manage access
- **Developers**: Use Temporal via SDKs

## 2. Functional Requirements

### 2.1 Organization Management

- Create, update, delete organizations
- Invite members with role-based access (Owner, Admin, Member)
- Organization-level settings (SSO, security policies)

### 2.2 Billing & Metering

- Track billable actions (workflow starts, signals, activities, etc.)
- Track storage usage (active GBh, retained GBh)
- Generate monthly invoices
- Process payments via Stripe
- Usage dashboards and alerts

### 2.3 Subscription Management

- Plan tiers: Free, Essentials ($100/mo), Business ($500/mo), Enterprise (custom)
- Quota enforcement based on plan
- Self-service upgrade/downgrade
- Overage billing

### 2.4 Enterprise Security

- SAML 2.0 SSO integration
- SCIM 2.0 user provisioning
- Audit logging (90-day hot, 7-year archive)
- Service accounts and API keys

### 2.5 Cloud Console

- Web UI for all management operations
- Usage dashboards with charts
- Invoice viewing and download
- Settings management

### 2.6 Infrastructure

- Multi-region deployment
- Automated failover
- Disaster recovery
- Auto-scaling (future)

## 3. Non-Functional Requirements

### 3.1 Performance

- API response time P99 < 200ms
- Console page load < 2 seconds
- Metering lag < 5 minutes

### 3.2 Availability

- Platform SLA: 99.9% (Essentials/Business), 99.99% (Enterprise)
- Planned maintenance windows: < 4 hours/month
- Zero-downtime deployments

### 3.3 Security

- SOC 2 Type II compliance
- Data encryption at rest (AES-256)
- Data encryption in transit (TLS 1.3)
- Regular penetration testing

### 3.4 Scalability

- Support 10,000+ organizations
- Support 100,000+ namespaces
- Handle 1M+ actions/second globally

---

## Consolidated Specifications

<!-- Source: README.md -->

## Temporal Cloud Platform Specification

> Complete specification for building a managed Temporal Cloud service.

### Quick Links

| Document                       | Purpose                               |
| ------------------------------ | ------------------------------------- |
| [spec.md](spec.md)             | Vision, requirements, success metrics |
| [plan.md](plan.md)             | Technical architecture and phases     |
| [tasks.md](tasks.md)           | Implementation task breakdown         |
| [quickstart.md](quickstart.md) | Developer setup guide                 |

### Document Index

#### 📋 Core Specifications

| Document                       | Description                                |
| ------------------------------ | ------------------------------------------ |
| [spec.md](spec.md)             | Functional and non-functional requirements |
| [plan.md](plan.md)             | Architecture, technology stack, phases     |
| [tasks.md](tasks.md)           | Detailed implementation tasks              |
| [data-model.md](data-model.md) | Database schema (14 tables)                |
| [quickstart.md](quickstart.md) | Local development setup                    |

#### 🏗️ Infrastructure

| Document                                       | Description                           |
| ---------------------------------------------- | ------------------------------------- |
| [infra.md](infra.md)                           | Cloud architecture, Terraform modules |
| [ha.md](ha.md)                                 | High availability design              |
| [dr.md](dr.md)                                 | Disaster recovery procedures          |
| [regions.md](regions.md)                       | Supported AWS/GCP regions             |
| [bootstrap.md](bootstrap.md)                   | Day 0 cluster bootstrap               |
| [capacity-planning.md](capacity-planning.md)   | Scaling and resource planning         |
| [multi-cloud.md](multi-cloud.md)               | AWS + GCP multi-cloud strategy        |
| [state-management.md](state-management.md)     | Terraform, app, session state         |
| [environment-config.md](environment-config.md) | Environment configuration             |

#### 🔄 Operations

| Document                                         | Description                   |
| ------------------------------------------------ | ----------------------------- |
| [cicd.md](cicd.md)                               | CI/CD pipeline configuration  |
| [monitoring.md](monitoring.md)                   | Observability and alerting    |
| [runbooks.md](runbooks.md)                       | Operational procedures        |
| [incident-management.md](incident-management.md) | Incident response process     |
| [database-operations.md](database-operations.md) | DB maintenance and migrations |

#### 🔐 Security & Access

| Document                                             | Description                 |
| ---------------------------------------------------- | --------------------------- |
| [security.md](security.md)                           | Security architecture       |
| [security-hardening.md](security-hardening.md)       | Defense in depth, hardening |
| [certificates.md](certificates.md)                   | mTLS certificate management |
| [api-keys.md](api-keys.md)                           | API key lifecycle           |
| [secrets-management.md](secrets-management.md)       | Secrets storage & rotation  |
| [saml.md](saml.md)                                   | SAML SSO integration        |
| [scim.md](scim.md)                                   | SCIM user provisioning      |
| [users.md](users.md)                                 | User roles and permissions  |
| [compliance-automation.md](compliance-automation.md) | SOC 2 automation            |
| [audit-logs.md](audit-logs.md)                       | Audit logging               |
| [bot-protection.md](bot-protection.md)               | Bot & abuse prevention      |
| [zero-day-response.md](zero-day-response.md)         | Vulnerability response      |

#### 💰 Billing & Payments

| Document                                       | Description                  |
| ---------------------------------------------- | ---------------------------- |
| [pricing.md](pricing.md)                       | Plan tiers and pricing       |
| [actions-detailed.md](actions-detailed.md)     | Billable actions reference   |
| [limits.md](limits.md)                         | System limits and quotas     |
| [stripe-mapping.md](stripe-mapping.md)         | Stripe product configuration |
| [payment-collection.md](payment-collection.md) | Dunning and collections      |

#### 🔌 APIs & Tools

| Document                                   | Description               |
| ------------------------------------------ | ------------------------- |
| [cloud-ops-api.md](cloud-ops-api.md)       | gRPC API reference        |
| [tcld.md](tcld.md)                         | CLI command reference     |
| [terraform.md](terraform.md)               | Terraform provider        |
| [metrics-endpoint.md](metrics-endpoint.md) | Prometheus metrics export |
| [sdk-extensions.md](sdk-extensions.md)     | SDK helper libraries      |
| [api-versioning.md](api-versioning.md)     | API versioning policy     |

#### 🎨 Frontend

| Document                               | Description          |
| -------------------------------------- | -------------------- |
| [console-design.md](console-design.md) | UI/UX specifications |
| [namespaces.md](namespaces.md)         | Namespace management |

#### 📦 Development Process

| Document                                             | Description                      |
| ---------------------------------------------------- | -------------------------------- |
| [git.md](git.md)                                     | Git workflow and branch strategy |
| [upstream-repos.md](upstream-repos.md)               | Managing 190+ temporalio repos   |
| [repo-automation.md](repo-automation.md)             | Automated repo sync & discovery  |
| [release-management.md](release-management.md)       | Release process                  |
| [upgrade-policies.md](upgrade-policies.md)           | Update & upgrade policies        |
| [testing.md](testing.md)                             | Testing strategy                 |
| [qa-process.md](qa-process.md)                       | QA workflow                      |
| [bug-triage.md](bug-triage.md)                       | Issue management                 |
| [feature-flags.md](feature-flags.md)                 | Feature flag system              |
| [dependency-management.md](dependency-management.md) | Dependency updates               |
| [plugin-management.md](plugin-management.md)         | Plugin safety & architecture     |

#### 🚀 Optimization & Reliability

| Document                                         | Description                     |
| ------------------------------------------------ | ------------------------------- |
| [cost-optimization.md](cost-optimization.md)     | Compute, storage, network costs |
| [system-optimization.md](system-optimization.md) | Kernel, runtime, DB tuning      |
| [crash-proofing.md](crash-proofing.md)           | Circuit breakers, chaos eng     |
| [log-management.md](log-management.md)           | Centralized logging & retention |
| [cdn.md](cdn.md)                                 | Edge caching & global routing   |

#### 🤝 Customer & Support

| Document                                         | Description                  |
| ------------------------------------------------ | ---------------------------- |
| [customer-onboarding.md](customer-onboarding.md) | Signup to first workflow     |
| [sla.md](sla.md)                                 | Service level agreements     |
| [support-escalation.md](support-escalation.md)   | Support tiers and escalation |

#### 🧠 Business Logic

| Document                                                           | Description                  |
| ------------------------------------------------------------------ | ---------------------------- |
| [logic/billing-reconciliation.md](logic/billing-reconciliation.md) | Billing workflow algorithms  |
| [logic/quota-enforcement.md](logic/quota-enforcement.md)           | Rate limiting implementation |
| [logic/namespace-failover.md](logic/namespace-failover.md)         | Failover state machine       |

#### 📝 Architecture Decisions

| Document                                                                       | Description            |
| ------------------------------------------------------------------------------ | ---------------------- |
| [adr/README.md](adr/README.md)                                                 | ADR index and template |
| [adr/ADR-001-postgresql-persistence.md](adr/ADR-001-postgresql-persistence.md) | Database choice        |

#### 📜 Governance

| Document                                               | Description        |
| ------------------------------------------------------ | ------------------ |
| [../memory/constitution.md](../memory/constitution.md) | Project principles |

---

### End-to-End Coverage

This specification covers the complete lifecycle:

```
┌─────────────────────────────────────────────────────────────────────┐
│                         CUSTOMER JOURNEY                             │
├─────────────────────────────────────────────────────────────────────┤
│  Signup → Verify → Trial → Subscribe → Use → Pay → Support → Renew │
│    ↓        ↓       ↓         ↓        ↓     ↓       ↓         ↓   │
│  [onboarding] [saml]  [pricing] [stripe] [limits] [dunning] [sla] │
└─────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────┐
│                         ENGINEERING JOURNEY                          │
├─────────────────────────────────────────────────────────────────────┤
│  Plan → Code → Test → Review → Deploy → Monitor → Debug → Improve   │
│   ↓      ↓      ↓       ↓        ↓        ↓        ↓        ↓      │
│ [spec] [git] [testing] [qa]   [cicd]  [monitor] [bugs]  [release] │
└─────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────┐
│                         INFRASTRUCTURE JOURNEY                       │
├─────────────────────────────────────────────────────────────────────┤
│  Bootstrap → Deploy → Scale → Monitor → Incident → Recover          │
│     ↓          ↓       ↓        ↓          ↓          ↓            │
│ [bootstrap] [infra] [capacity] [monitor] [incident]  [dr]          │
└─────────────────────────────────────────────────────────────────────┘
```

### Getting Started

1. **Read the spec**: Start with [spec.md](spec.md) for requirements
2. **Understand the plan**: Review [plan.md](plan.md) for architecture
3. **Set up locally**: Follow [quickstart.md](quickstart.md)
4. **Pick a task**: Check [tasks.md](tasks.md) for work items

### Document Count

**Total: 70 specification documents**

- Core: 5
- Infrastructure: 9
- Operations: 5
- Security: 12
- Billing: 5
- APIs: 6
- Frontend: 2
- Process: 11
- Optimization: 5
- Customer: 3
- Logic: 3
- ADR: 2
- Governance: 1

### Complete Coverage

This spec covers **everything** from start to receiving payment:

| Phase          | Coverage                                               |
| -------------- | ------------------------------------------------------ |
| Repo Sync      | `repo-automation.md`, `upstream-repos.md` (190+ repos) |
| Git/Process    | `git.md`, `plugin-management.md`                       |
| Development    | `quickstart.md`, `dependency-management.md`            |
| Optimization   | `cost-optimization.md`, `system-optimization.md`       |
| Reliability    | `crash-proofing.md`, `log-management.md`, `cdn.md`     |
| Testing/QA     | `testing.md`, `qa-process.md`                          |
| Deployment     | `cicd.md`, `release-management.md`                     |
| Infrastructure | `infra.md`, `multi-cloud.md`, `bootstrap.md`           |
| State          | `state-management.md` (Terraform, app, session)        |
| Environments   | `environment-config.md`                                |
| HA/DR          | `ha.md`, `dr.md`, `capacity-planning.md`               |
| Security       | `security-hardening.md`, `secrets-management.md`       |
| Bot/Abuse      | `bot-protection.md`                                    |
| Zero-Day       | `zero-day-response.md`                                 |
| Upgrades       | `upgrade-policies.md`                                  |
| Billing        | `pricing.md`, `stripe-mapping.md`                      |
| Payments       | `payment-collection.md` (dunning)                      |
| Onboarding     | `customer-onboarding.md`                               |
| Support        | `support-escalation.md`, `sla.md`                      |

<!-- Source: actions-detailed.md -->

## Billable Actions - Complete List

### Workflow Actions

| Action                   | Count | Notes                                |
| ------------------------ | ----- | ------------------------------------ |
| Workflow started         | 1     | Via client, Continue-As-New, Child   |
| Workflow reset           | 1     | Actions before reset still count     |
| Timer started            | 1     | Includes implicit SDK timers         |
| Search Attribute upsert  | 1     | Per UpsertSearchAttributes call      |
| Signal sent              | 1     | From client or workflow              |
| Query received           | 1     | Including UI stack trace             |
| Update received          | 1     | Successful or rejected               |
| Side Effect recorded     | 1     | Mutable: only on change              |
| Workflow options updated | 1     | Callback attach, versioning override |

#### Non-Billable Workflow Operations

- De-duplicated workflow starts (same Workflow ID)
- De-duplicated updates (same Update ID)
- Search attributes at workflow start
- TemporalChangeVersion search attribute

### Child Workflow Actions

| Action               | Count | Notes                      |
| -------------------- | ----- | -------------------------- |
| Start Child Workflow | 2     | Intent (1) + Execution (1) |

### Activity Actions

| Action               | Count | Notes                    |
| -------------------- | ----- | ------------------------ |
| Activity started     | 1     | Each attempt             |
| Activity retried     | 1     | Each retry attempt       |
| Local Activity batch | 1     | All in one Workflow Task |
| Activity Heartbeat   | 1     | Only if reaches server   |

#### Local Activity Details

- All Local Activities in one Workflow Task = 1 Action
- Each Workflow Task heartbeat = 1 additional Action
- Retries after heartbeat = 1 Action (capped at 100)

#### Heartbeat Throttling

- SDKs throttle heartbeats (default: 80% of timeout)
- Only heartbeats reaching server are billed
- Local Activities don't have heartbeats

### Schedule Actions

| Action             | Count | Notes                       |
| ------------------ | ----- | --------------------------- |
| Schedule execution | 3     | 2 (schedule) + 1 (workflow) |

### Export Actions

| Action            | Count | Notes                |
| ----------------- | ----- | -------------------- |
| Workflow exported | 1     | Per workflow history |

### Nexus Actions

| Action              | Namespace | Count          |
| ------------------- | --------- | -------------- |
| Operation scheduled | Caller    | 1              |
| Operation canceled  | Caller    | 1              |
| Handler primitives  | Handler   | Normal billing |

#### Nexus Notes

- Retries of Nexus Operation itself: Not billed
- Underlying Activities/Workflows: Normal billing on handler namespace

### Action Estimation

#### In UI

- Workflow history shows "Billable Actions" column
- Summary at top of workflow view
- Experimental feature - may not include all actions

#### Not Included in UI Estimate

- Query
- Activity Heartbeats
- Rejected Updates
- Export
- Schedule overhead

### Cost Optimization Tips

1. **Batch operations**: Use Local Activities for small, fast operations
2. **Reduce signals**: Combine multiple signals into one
3. **Optimize timers**: Use longer durations when possible
4. **Heartbeat wisely**: Increase heartbeat timeout to reduce frequency
5. **Use Continue-As-New**: Prevent unbounded history growth

<!-- Source: adr/ADR-001-postgresql-persistence.md -->

## ADR-001: Use PostgreSQL for Cloud Platform Persistence

### Status

Accepted

### Context

The Temporal Cloud Platform needs a reliable, scalable database for storing:

- Organizations and users
- Subscriptions and billing data
- Usage records and invoices
- Audit logs
- Namespace metadata

We need to choose a database that:

- Is highly available
- Scales with our expected growth
- Has strong consistency guarantees
- Is operationally mature in cloud environments
- Supports our team's expertise

### Decision

We will use **PostgreSQL** as the primary database for the Cloud Platform, deployed on AWS RDS with Multi-AZ configuration.

Specifically:

- AWS RDS PostgreSQL 15
- Multi-AZ for high availability
- Read replicas for read-heavy workloads
- Point-in-time recovery enabled
- Encryption at rest (AES-256)

### Consequences

#### Positive

- **Proven reliability**: PostgreSQL is battle-tested at scale
- **AWS managed**: RDS handles backups, patching, failover
- **Team expertise**: Team has deep PostgreSQL experience
- **Rich ecosystem**: Excellent tooling (pgAdmin, pg_dump, etc.)
- **Strong consistency**: ACID compliance for financial data
- **JSON support**: JSONB for flexible schema where needed

#### Negative

- **Vertical scaling limits**: Eventually may need sharding
- **Vendor lock-in**: Tied to AWS RDS features
- **Cost**: More expensive than self-managed
- **Single-region**: Cross-region replication adds complexity

#### Neutral

- Requires careful schema design upfront
- Migration tooling (golang-migrate) adds operational overhead

### Alternatives Considered

#### Option A: CockroachDB

**Pros**: Distributed, automatic scaling, multi-region native
**Cons**: Less team experience, more complex operations, higher cost
**Decision**: Overkill for initial scale; revisit at 100K+ orgs

#### Option B: MongoDB

**Pros**: Flexible schema, good for rapid iteration
**Cons**: Weaker consistency guarantees, team prefers SQL
**Decision**: Rejected due to consistency requirements for billing

#### Option C: Self-managed PostgreSQL on EC2

**Pros**: More control, potentially cheaper
**Cons**: Operational burden, HA complexity, no managed backups
**Decision**: RDS overhead justified by reduced ops burden

### References

- [AWS RDS PostgreSQL Documentation](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_PostgreSQL.html)
- [PostgreSQL 15 Release Notes](https://www.postgresql.org/docs/15/release-15.html)
- Team discussion: 2025-01-05

<!-- Source: adr/README.md -->

## Architecture Decision Records

### What is an ADR?

An Architecture Decision Record (ADR) captures an important architectural decision made along with its context and consequences.

### When to Write an ADR

- New technology choice
- Major design pattern
- Significant refactoring
- Integration approach
- Security decisions
- Breaking changes

### ADR Template

```markdown
# ADR-NNN: [Title]

## Status

[Proposed | Accepted | Deprecated | Superseded by ADR-XXX]

## Context

What is the issue that we're seeing that is motivating this decision?

## Decision

What is the change that we're proposing and/or doing?

## Consequences

What becomes easier or harder because of this change?

### Positive

-

### Negative

-

### Neutral

-

## Alternatives Considered

### Option A

Description and trade-offs

### Option B

Description and trade-offs

## References

- [Link to relevant docs]
- [Link to discussion]
```

### Process

1. **Propose**: Create ADR in `proposed/` folder
2. **Review**: Team reviews in PR
3. **Accept**: Move to `accepted/` folder
4. **Implement**: Reference ADR in code
5. **Supersede**: When replaced, update status

### Naming Convention

`ADR-{number}-{short-title}.md`

Example: `ADR-001-use-postgresql-for-persistence.md`

### Index

| ADR | Title                              | Status   | Date       |
| --- | ---------------------------------- | -------- | ---------- |
| 001 | Use PostgreSQL for persistence     | Accepted | 2025-01-01 |
| 002 | Adopt gRPC for Cloud Ops API       | Accepted | 2025-01-15 |
| 003 | Multi-region namespace replication | Accepted | 2025-02-01 |

<!-- Source: api-keys.md -->

## API Key Management

### Overview

API keys provide programmatic access to Temporal Cloud without mTLS certificates. They are tied to a user or service account identity.

### Key Types

| Type                    | Owner           | Use Case                 |
| ----------------------- | --------------- | ------------------------ |
| User API Key            | User            | Personal automation, CLI |
| Service Account API Key | Service Account | CI/CD, Workers           |

### Key Lifecycle

#### Create

```bash
# Via tcld
tcld apikey create \
  --name "CI Pipeline" \
  --duration 90d

# Output: API key (shown only once)
# temporal_ak_xxxxxxxxxxxxxxxxxxxx
```

#### List

```bash
tcld apikey list
```

#### Disable/Enable

```bash
tcld apikey disable --id key-123
tcld apikey enable --id key-123
```

#### Delete

```bash
tcld apikey delete --id key-123
```

#### Rotate

```bash
tcld apikey rotate --id key-123
# Returns new key, old key disabled after 24h
```

### Permissions

API keys inherit permissions from their owner:

- User API Key → User's account role + namespace permissions
- Service Account API Key → Service account's role + permissions

### Security Best Practices

1. **Short expiration**: Set 90-day expiration
2. **Least privilege**: Use service accounts with minimal permissions
3. **Rotate regularly**: Rotate keys every 90 days
4. **Monitor usage**: Review last_used_at regularly
5. **Disable unused**: Disable keys not used in 30+ days

### Using API Keys

#### Temporal CLI

```bash
temporal workflow list \
  --api-key temporal_ak_xxxx \
  --address my-namespace.tmprl.cloud:443
```

#### SDKs

```go
// Go SDK
client, err := client.Dial(client.Options{
    HostPort:  "my-namespace.tmprl.cloud:443",
    Namespace: "my-namespace",
    Credentials: client.NewAPIKeyStaticCredentials("temporal_ak_xxxx"),
})
```

#### tcld

```bash
tcld --api-key temporal_ak_xxxx namespace list
```

#### Cloud Ops API

```bash
curl -H "Authorization: Bearer temporal_ak_xxxx" \
  https://api.temporal.io/api/v1/namespaces
```

### Schema

```sql
CREATE TABLE api_keys (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    owner_type VARCHAR(50) NOT NULL,  -- 'user' or 'service_account'
    owner_id UUID NOT NULL,
    key_hash VARCHAR(255) NOT NULL,   -- SHA-256 hash
    key_prefix VARCHAR(10) NOT NULL,  -- First 10 chars for identification
    name VARCHAR(255),
    expires_at TIMESTAMPTZ,
    disabled BOOLEAN DEFAULT FALSE,
    last_used_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_api_keys_owner ON api_keys(owner_type, owner_id);
CREATE INDEX idx_api_keys_prefix ON api_keys(key_prefix);
```

### Rate Limits

| Scope       | Limit            |
| ----------- | ---------------- |
| Per API key | 20 requests/sec  |
| Per account | 200 requests/sec |

### Audit Events

All API key operations are logged:

- `CreateAPIKey`
- `DeleteAPIKey`
- `UpdateAPIKey` (enable/disable)
- `RotateAPIKey`

<!-- Source: api-versioning.md -->

## API Versioning

### Versioning Strategy

We use **URL path versioning** for the Cloud Ops API.

```
https://api.temporal.io/api/v1/namespaces
https://api.temporal.io/api/v2/namespaces
```

### Version Lifecycle

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Alpha     │────▶│   Beta      │────▶│   Stable    │
│   (v1alpha1)│     │   (v1beta1) │     │   (v1)      │
└─────────────┘     └─────────────┘     └─────────────┘
                                               │
                                               │ (12+ months)
                                               ▼
                                        ┌─────────────┐
                                        │  Deprecated │
                                        └─────────────┘
                                               │
                                               │ (6 months)
                                               ▼
                                        ┌─────────────┐
                                        │   Sunset    │
                                        └─────────────┘
```

#### Stability Levels

| Level      | Breaking Changes           | Support Period    |
| ---------- | -------------------------- | ----------------- |
| Alpha      | Any time                   | None              |
| Beta       | Major releases only        | 6 months          |
| Stable     | Never (new version needed) | 24 months minimum |
| Deprecated | None                       | 6 months          |

### Breaking vs Non-Breaking Changes

#### Non-Breaking (OK)

- Adding new optional fields
- Adding new endpoints
- Adding new enum values
- Relaxing validation (e.g., increasing limit)
- Bug fixes to match documented behavior

#### Breaking (Requires New Version)

- Removing or renaming fields
- Changing field types
- Removing endpoints
- Adding required fields
- Changing validation to be more strict
- Changing default values

### Proto Compatibility

#### Reserved Fields

```protobuf
message Namespace {
  string id = 1;
  string name = 2;
  // Removed fields - NEVER reuse these numbers
  reserved 3, 4;
  reserved "old_field_name";

  // New fields start from next available
  string region = 5;
}
```

#### Backwards Compatible Changes

```protobuf
// Original v1
message CreateNamespaceRequest {
  string name = 1;
  string region = 2;
}

// Updated v1 (compatible)
message CreateNamespaceRequest {
  string name = 1;
  string region = 2;
  optional int32 retention_days = 3; // NEW - optional
}
```

### Version Migration

#### Deprecation Notice

When deprecating an API version:

1. **Announcement** (T-6 months)

   - Blog post
   - Email to all customers
   - Console banner

2. **Warning Headers** (T-6 months)

   ```
   Deprecation: true
   Sunset: Sat, 01 Jul 2025 00:00:00 GMT
   Link: <https://docs.temporal.io/migration/v1-to-v2>; rel="deprecation"
   ```

3. **Migration Guide**

   - Document all changes
   - Provide code examples
   - Offer migration tooling if complex

4. **Sunset** (T-0)
   - Return 410 Gone
   - Log attempts for customer outreach

#### Migration Support

```go
// Dual-write during migration period
func CreateNamespace(ctx context.Context, req *v2.CreateNamespaceRequest) {
    // Write to v2 storage
    writeV2(req)

    // Also maintain v1 compatibility layer
    if featureFlags.IsEnabled("v1_compat") {
        v1Req := convertV2ToV1(req)
        writeV1(v1Req)
    }
}
```

### Client SDK Versioning

| SDK        | Cloud API v1 | Cloud API v2 |
| ---------- | ------------ | ------------ |
| sdk-go 1.x | ✅           | ❌           |
| sdk-go 2.x | ✅ (compat)  | ✅           |

### Deprecation Policy

#### Minimum Support Periods

| API Type           | Stable Support | Deprecation Warning |
| ------------------ | -------------- | ------------------- |
| Public REST/gRPC   | 24 months      | 6 months            |
| Terraform Provider | 18 months      | 6 months            |
| CLI                | 12 months      | 3 months            |
| Internal           | 6 months       | 1 month             |

#### Customer Communication

| Channel       | Timing                                            |
| ------------- | ------------------------------------------------- |
| Documentation | Immediately on deprecation                        |
| Email         | 6 months, 3 months, 1 month, 1 week before sunset |
| Console       | Warning banner on affected pages                  |
| API Response  | Deprecation header                                |

### Example: v1 to v2 Migration

#### What Changed

| v1                          | v2                     | Change             |
| --------------------------- | ---------------------- | ------------------ |
| `GET /v1/namespaces`        | `GET /v2/namespaces`   | Pagination changed |
| `namespace.settings`        | `namespace.config`     | Field renamed      |
| `retention_period` (string) | `retention_days` (int) | Type changed       |

#### Migration Code

```go
// Helper for clients
func MigrateV1ToV2Response(v1 *v1.Namespace) *v2.Namespace {
    return &v2.Namespace{
        Id:            v1.Id,
        Name:          v1.Name,
        Config:        convertSettings(v1.Settings),
        RetentionDays: parseRetention(v1.RetentionPeriod),
    }
}
```

### Testing Compatibility

```go
func TestBackwardsCompatibility(t *testing.T) {
    // Load v1 fixtures
    v1Fixtures := loadFixtures("testdata/v1/*.json")

    for _, fixture := range v1Fixtures {
        // Should still parse with current code
        var req v1.CreateNamespaceRequest
        err := json.Unmarshal(fixture, &req)
        require.NoError(t, err)

        // Should produce valid response
        resp, err := handler.CreateNamespace(ctx, &req)
        require.NoError(t, err)
        require.NotNil(t, resp)
    }
}
```

### Documentation

Each version has its own API reference:

- `docs.temporal.io/api/v1`
- `docs.temporal.io/api/v2`

With clear migration guide linking the two.

<!-- Source: audit-logs.md -->

## Audit Logging

### Overview

Audit logs provide a record of all administrative actions in Temporal Cloud for compliance and security monitoring.

### Supported Events

#### Account

| Event                 | Description                         |
| --------------------- | ----------------------------------- |
| ChangeAccountPlanType | Plan upgrade/downgrade              |
| UpdateAccountAPI      | Configure audit logs, observability |

#### API Keys

| Event        | Description              |
| ------------ | ------------------------ |
| CreateAPIKey | API key created          |
| DeleteAPIKey | API key deleted          |
| UpdateAPIKey | API key enabled/disabled |

#### Namespace

| Event                          | Description                |
| ------------------------------ | -------------------------- |
| CreateNamespaceAPI             | Namespace created          |
| DeleteNamespaceAPI             | Namespace deleted          |
| UpdateNamespaceAPI             | Namespace settings changed |
| FailoverNamespacesAPI          | HA namespace failover      |
| RenameCustomSearchAttributeAPI | Search attribute renamed   |

#### Users

| Event                     | Description                  |
| ------------------------- | ---------------------------- |
| CreateUserAPI             | User created                 |
| DeleteUserAPI             | User deleted                 |
| InviteUsersAPI            | User invited                 |
| UpdateUserAPI             | User role changed            |
| SetUserNamespaceAccessAPI | Namespace permission changed |

#### Service Accounts

| Event                      | Description             |
| -------------------------- | ----------------------- |
| CreateServiceAccount       | Service account created |
| DeleteServiceAccount       | Service account deleted |
| UpdateServiceAccount       | Service account updated |
| CreateServiceAccountAPIKey | API key created         |

#### Nexus

| Event               | Description      |
| ------------------- | ---------------- |
| CreateNexusEndpoint | Endpoint created |
| DeleteNexusEndpoint | Endpoint deleted |
| UpdateNexusEndpoint | Endpoint updated |

#### Connectivity

| Event                  | Description  |
| ---------------------- | ------------ |
| CreateConnectivityRule | Rule created |
| DeleteConnectivityRule | Rule deleted |

### Log Format

```json
{
  "operation": "CreateNamespaceAPI",
  "status": "OK",
  "version": 2,
  "log_id": "uuid-here",
  "x_forwarded_for": "10.1.2.3",
  "emit_time": "2025-01-01T00:00:00Z",
  "principal": {
    "type": "user",
    "id": "user-123",
    "name": "user@example.com",
    "api_key_id": ""
  },
  "raw_details": {
    "namespace_name": "my-namespace",
    "region": "aws-us-east-1"
  }
}
```

### Export Configuration

#### Supported Sinks

- AWS S3
- GCP Cloud Storage
- Datadog
- Splunk

#### Configure via Console

1. Go to Settings → Audit Logs
2. Click "Add Sink"
3. Select sink type
4. Configure credentials
5. Test connection
6. Enable

#### Configure via tcld

```bash
tcld account audit-log-sink create \
  --type s3 \
  --bucket my-audit-logs \
  --region us-east-1 \
  --role-arn arn:aws:iam::123456789:role/temporal-audit
```

### Retention

| Storage   | Retention | Purpose      |
| --------- | --------- | ------------ |
| Hot (API) | 90 days   | Quick access |
| Cold (S3) | 7 years   | Compliance   |

### Querying Logs

#### Via Console

1. Go to Settings → Audit Logs
2. Filter by date, operation, user
3. Export to CSV

#### Via API

```bash
curl -H "Authorization: Bearer $API_KEY" \
  "https://api.temporal.io/api/v1/audit-logs?start_time=2025-01-01&end_time=2025-01-31"
```

### Compliance

#### SOC 2

- All state-changing operations logged
- Immutable storage
- 7-year retention

#### GDPR

- User actions logged
- Data access logged
- Export capability

### Best Practices

1. **Export to SIEM**: Send logs to your security monitoring system
2. **Alert on anomalies**: Set up alerts for unusual activity
3. **Regular review**: Review logs weekly for security
4. **Retain exports**: Keep exported logs beyond 90 days

<!-- Source: bootstrap.md -->

## Day 0 Bootstrap Sequence

### Problem Statement

The cloud platform runs on Temporal, but it also manages Temporal. This creates a circular dependency that must be resolved during initial setup.

### Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                    Seed Cluster                              │
│  ┌─────────────────────────────────────────────────────┐    │
│  │  Standalone Temporal (NOT multi-tenant)              │    │
│  │  - Runs control plane workflows only                 │    │
│  │  - Invoice generation, usage aggregation, etc.       │    │
│  └─────────────────────────────────────────────────────┘    │
│                           │                                  │
│                           │ Manages                          │
│                           ▼                                  │
│  ┌─────────────────────────────────────────────────────┐    │
│  │  Customer-Facing Temporal Clusters                   │    │
│  │  - Multi-tenant                                      │    │
│  │  - Per-region deployment                             │    │
│  └─────────────────────────────────────────────────────┘    │
└─────────────────────────────────────────────────────────────┘
```

### Bootstrap Phases

#### Phase 0: Prerequisites

- [ ] AWS accounts created (dev, staging, prod)
- [ ] Domain registered (temporal-cloud.io)
- [ ] SSL certificates provisioned
- [ ] Terraform state backend configured

#### Phase 1: Global Infrastructure

```bash
# 1.1 Create Terraform state backend
cd terraform/bootstrap
terraform init
terraform apply -target=module.state_backend

# 1.2 Create global resources
cd terraform/global
terraform init
terraform apply
```

**Resources created:**

- Route53 hosted zone
- ACM certificates
- IAM roles
- S3 buckets (audit logs, backups)

#### Phase 2: Seed Cluster

```bash
# 2.1 Deploy seed VPC
cd terraform/seed
terraform apply -target=module.vpc

# 2.2 Deploy seed database
terraform apply -target=module.rds

# 2.3 Deploy seed Temporal cluster
terraform apply -target=module.temporal_seed
```

**Seed cluster configuration:**

```yaml
# values-seed.yaml
temporal:
  server:
    replicaCount: 1 # Single node for seed
    config:
      persistence:
        default:
          driver: sql
          sql:
            host: seed-db.internal
            port: 5432
            database: temporal

  # No multi-tenancy
  namespaces:
    - name: cloud-platform
      retention: 30d
```

#### Phase 3: Cloud Platform Database

```bash
# 3.1 Run cloud platform migrations
cd temporal-cloud-platform
DATABASE_URL="postgres://..." go run cmd/migrate/main.go up

# 3.2 Seed initial data
go run cmd/seed/main.go \
  --admin-email admin@temporal.io \
  --org-name "Temporal" \
  --org-slug "temporal"
```

**Seed data:**

```sql
-- System organization
INSERT INTO organizations (id, name, slug)
VALUES ('00000000-0000-0000-0000-000000000000', 'System', 'system');

-- Initial admin user
INSERT INTO users (id, email, name)
VALUES ('00000000-0000-0000-0000-000000000001', 'admin@temporal.io', 'Admin');

INSERT INTO organization_members (organization_id, user_id, role)
VALUES ('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000001', 'owner');
```

#### Phase 4: Cloud Platform Services

```bash
# 4.1 Deploy cloud platform to seed cluster
cd terraform/seed
terraform apply -target=module.cloud_platform

# 4.2 Verify health
curl https://api.temporal-cloud.io/health
```

#### Phase 5: First Customer Region

```bash
# 5.1 Trigger region creation workflow
temporal workflow start \
  --task-queue cloud-platform \
  --type CreateRegionWorkflow \
  --input '{"region": "aws-us-east-1", "size": "production"}'

# 5.2 Monitor workflow
temporal workflow show --workflow-id create-region-aws-us-east-1
```

**CreateRegionWorkflow:**

```go
func CreateRegionWorkflow(ctx workflow.Context, input CreateRegionInput) error {
    // Step 1: Provision infrastructure
    workflow.ExecuteActivity(ctx, ProvisionRegionInfra, input)

    // Step 2: Deploy Temporal cluster
    workflow.ExecuteActivity(ctx, DeployTemporalCluster, input)

    // Step 3: Configure networking
    workflow.ExecuteActivity(ctx, ConfigureNetworking, input)

    // Step 4: Register region in control plane
    workflow.ExecuteActivity(ctx, RegisterRegion, input)

    // Step 5: Health check
    workflow.ExecuteActivity(ctx, HealthCheckRegion, input)

    return nil
}
```

#### Phase 6: Verification

```bash
# 6.1 Create test namespace
tcld namespace create \
  --name test-namespace \
  --region aws-us-east-1

# 6.2 Run test workflow
temporal workflow start \
  --address test-namespace.tmprl.cloud:443 \
  --tls-cert-path client.pem \
  --tls-key-path client.key \
  --task-queue test \
  --type TestWorkflow

# 6.3 Verify billing
tcld account usage
```

### Disaster Recovery for Seed Cluster

The seed cluster is critical infrastructure. Special DR considerations:

#### Backup Strategy

- Database: Continuous WAL + daily snapshots
- Workflow state: Replicated to S3
- Configuration: Git (infrastructure as code)

#### Recovery Procedure

1. Provision new seed cluster from Terraform
2. Restore database from backup
3. Workflows will resume from last checkpoint
4. Update DNS to new seed cluster

#### Seed Cluster Monitoring

- 24/7 alerting
- Separate from customer monitoring
- Direct PagerDuty escalation

### Security Considerations

#### Seed Cluster Access

- No customer data
- Restricted to platform team
- Separate AWS account
- VPN-only access

#### Secrets

- Stored in AWS Secrets Manager
- Rotated every 30 days
- Accessed via IAM roles only

### Runbook Location

Detailed bootstrap runbook:
`temporal-cloud-infra/docs/runbooks/bootstrap.md`

<!-- Source: bot-protection.md -->

## Bot Protection

### Threat Model

| Bot Type            | Impact             | Mitigation                           |
| ------------------- | ------------------ | ------------------------------------ |
| Credential stuffing | Account takeover   | Rate limiting, MFA, breach detection |
| Scraping            | Data exfiltration  | Rate limiting, fingerprinting        |
| DDoS                | Service disruption | WAF, Shield, auto-scaling            |
| Spam/abuse          | Resource waste     | CAPTCHA, reputation                  |
| API abuse           | Quota bypass       | API keys, rate limiting              |

### Multi-Layer Defense

```
┌─────────────────────────────────────────────────────────────────┐
│  Layer 1: Edge (Cloudflare/AWS WAF)                             │
│  - IP reputation                                                 │
│  - Known bot signatures                                          │
│  - JavaScript challenge                                          │
├─────────────────────────────────────────────────────────────────┤
│  Layer 2: Rate Limiting                                          │
│  - Per-IP limits                                                 │
│  - Per-account limits                                            │
│  - Per-endpoint limits                                           │
├─────────────────────────────────────────────────────────────────┤
│  Layer 3: Behavioral Analysis                                    │
│  - Request patterns                                              │
│  - Session anomalies                                             │
│  - Impossible travel                                             │
├─────────────────────────────────────────────────────────────────┤
│  Layer 4: Challenge/Response                                     │
│  - CAPTCHA (high-risk actions)                                  │
│  - Email verification                                            │
│  - MFA                                                          │
└─────────────────────────────────────────────────────────────────┘
```

### Rate Limiting

#### Implementation

```go
type RateLimiter struct {
    redis *redis.Client
}

type RateLimitConfig struct {
    // Per IP
    IPRequestsPerMinute     int
    IPRequestsPerHour       int

    // Per account
    AccountRequestsPerMinute int
    AccountRequestsPerHour   int

    // Per endpoint
    LoginAttemptsPerMinute   int
    SignupAttemptsPerHour    int
}

func (rl *RateLimiter) Check(ctx context.Context, key string, limit int, window time.Duration) (bool, int, error) {
    now := time.Now()
    windowKey := fmt.Sprintf("%s:%d", key, now.Unix()/int64(window.Seconds()))

    pipe := rl.redis.Pipeline()
    incr := pipe.Incr(ctx, windowKey)
    pipe.Expire(ctx, windowKey, window)
    pipe.Exec(ctx)

    count := int(incr.Val())
    remaining := limit - count

    if count > limit {
        return false, 0, nil
    }

    return true, remaining, nil
}
```

#### Rate Limits by Endpoint

| Endpoint              | Per IP/min | Per Account/min | Per IP/hour |
| --------------------- | ---------- | --------------- | ----------- |
| Login                 | 10         | N/A             | 100         |
| Signup                | 5          | N/A             | 20          |
| Password Reset        | 3          | 3               | 10          |
| API (authenticated)   | 1000       | 500             | 10000       |
| API (unauthenticated) | 100        | N/A             | 1000        |

#### Response Headers

```
X-RateLimit-Limit: 1000
X-RateLimit-Remaining: 999
X-RateLimit-Reset: 1704067200
Retry-After: 60
```

### Bot Detection

#### Device Fingerprinting

```typescript
// Client-side fingerprint collection
import FingerprintJS from "@fingerprintjs/fingerprintjs";

async function getFingerprint() {
  const fp = await FingerprintJS.load();
  const result = await fp.get();
  return result.visitorId;
}

// Send with requests
fetch("/api/action", {
  headers: {
    "X-Device-Fingerprint": await getFingerprint(),
  },
});
```

#### Behavioral Signals

```go
type BehaviorAnalyzer struct {
    redis *redis.Client
}

type RequestSignals struct {
    IP            string
    UserAgent     string
    Fingerprint   string
    RequestPath   string
    RequestMethod string
    Timestamp     time.Time
    SessionID     string
}

func (ba *BehaviorAnalyzer) AnalyzeRequest(ctx context.Context, signals RequestSignals) (riskScore float64, reasons []string) {
    // Check velocity
    if ba.isVelocityTooHigh(ctx, signals.IP) {
        riskScore += 0.3
        reasons = append(reasons, "high_velocity")
    }

    // Check user agent
    if ba.isSuspiciousUserAgent(signals.UserAgent) {
        riskScore += 0.2
        reasons = append(reasons, "suspicious_ua")
    }

    // Check fingerprint consistency
    if ba.isFingerprintMismatch(ctx, signals.SessionID, signals.Fingerprint) {
        riskScore += 0.4
        reasons = append(reasons, "fingerprint_mismatch")
    }

    // Check impossible travel
    if ba.isImpossibleTravel(ctx, signals.SessionID, signals.IP) {
        riskScore += 0.5
        reasons = append(reasons, "impossible_travel")
    }

    return riskScore, reasons
}
```

#### IP Reputation

```go
// Check IP against threat intelligence
func CheckIPReputation(ip string) (score float64, threats []string) {
    // Check against known bad IP lists
    if isInBotnet(ip) {
        return 1.0, []string{"botnet"}
    }

    // Check against Tor exit nodes
    if isTorExitNode(ip) {
        return 0.8, []string{"tor"}
    }

    // Check against VPN/proxy lists
    if isVPN(ip) || isProxy(ip) {
        return 0.3, []string{"vpn_proxy"}
    }

    // Check recent abuse reports
    abuseScore := getAbuseScore(ip)
    if abuseScore > 0.5 {
        return abuseScore, []string{"abuse_reports"}
    }

    return 0.0, nil
}
```

### CAPTCHA Integration

#### When to Show CAPTCHA

| Trigger                             | Action       |
| ----------------------------------- | ------------ |
| 5+ failed logins                    | Show CAPTCHA |
| Suspicious fingerprint              | Show CAPTCHA |
| High-risk signup (disposable email) | Show CAPTCHA |
| Rate limit approaching              | Show CAPTCHA |
| VPN/Tor detected                    | Show CAPTCHA |

#### Implementation (hCaptcha)

```typescript
// React component
import HCaptcha from "@hcaptcha/react-hcaptcha";

function LoginForm() {
  const [captchaToken, setCaptchaToken] = useState<string | null>(null);
  const [showCaptcha, setShowCaptcha] = useState(false);

  const onSubmit = async (data: LoginData) => {
    const response = await fetch("/api/login", {
      method: "POST",
      body: JSON.stringify({
        ...data,
        captcha_token: captchaToken,
      }),
    });

    if (response.status === 428) {
      // Server requires CAPTCHA
      setShowCaptcha(true);
    }
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      {/* ... form fields ... */}

      {showCaptcha && (
        <HCaptcha
          sitekey={process.env.HCAPTCHA_SITE_KEY}
          onVerify={setCaptchaToken}
        />
      )}

      <button type="submit">Login</button>
    </form>
  );
}
```

#### Server-Side Verification

```go
func VerifyCaptcha(token string) (bool, error) {
    resp, err := http.PostForm("https://hcaptcha.com/siteverify", url.Values{
        "secret":   {os.Getenv("HCAPTCHA_SECRET")},
        "response": {token},
    })

    var result struct {
        Success bool `json:"success"`
    }
    json.NewDecoder(resp.Body).Decode(&result)

    return result.Success, nil
}
```

### Account Protection

#### Credential Stuffing Prevention

```go
// Check password against breach databases
func IsBreachedPassword(password string) bool {
    // Use k-anonymity to check against HaveIBeenPwned
    hash := sha1.Sum([]byte(password))
    prefix := hex.EncodeToString(hash[:])[:5]
    suffix := hex.EncodeToString(hash[:])[5:]

    resp, _ := http.Get("https://api.pwnedpasswords.com/range/" + prefix)
    body, _ := io.ReadAll(resp.Body)

    return strings.Contains(string(body), strings.ToUpper(suffix))
}

// On registration/password change
if IsBreachedPassword(newPassword) {
    return errors.New("This password has been found in a data breach. Please choose a different password.")
}
```

#### Account Lockout

```go
func HandleLoginAttempt(ctx context.Context, email string, success bool) error {
    key := "login_attempts:" + email

    if success {
        // Clear attempts on success
        redis.Del(ctx, key)
        return nil
    }

    // Increment failed attempts
    attempts := redis.Incr(ctx, key).Val()
    redis.Expire(ctx, key, 15*time.Minute)

    if attempts >= 5 {
        // Lock account
        lockAccount(ctx, email)
        sendAccountLockedEmail(email)
        return ErrAccountLocked
    }

    return nil
}
```

### Monitoring & Alerts

#### Bot Detection Metrics

```yaml
alerts:
  - name: HighBotTraffic
    expr: |
      sum(rate(blocked_requests_total{reason="bot"}[5m])) > 100
    severity: warning

  - name: CredentialStuffingAttack
    expr: |
      sum(rate(failed_logins_total[5m])) by (ip) > 10
    severity: critical

  - name: ScrapingDetected
    expr: |
      sum(rate(requests_total[5m])) by (ip) > 1000
      AND sum(rate(requests_total[5m])) by (user_id) == 0
    severity: warning
```

#### Response Playbook

| Alert               | Response                               |
| ------------------- | -------------------------------------- |
| High bot traffic    | Review WAF rules, add blocks           |
| Credential stuffing | Block IP ranges, notify affected users |
| Scraping            | Add rate limits, consider legal action |
| DDoS                | Activate Shield, scale infrastructure  |

<!-- Source: bug-triage.md -->

## Bug Triage & Issue Management

### Issue Sources

1. **Customer Support** - Zendesk tickets escalated to engineering
2. **Internal** - Found by team during development/testing
3. **Monitoring** - Automated alerts that indicate bugs
4. **Security** - Vulnerability reports

### Issue Classification

#### Severity

| Severity | Description                             | Response | Resolution |
| -------- | --------------------------------------- | -------- | ---------- |
| P0       | Production down, data loss              | 15 min   | 4 hours    |
| P1       | Major feature broken, workaround exists | 1 hour   | 24 hours   |
| P2       | Feature partially broken                | 4 hours  | 1 week     |
| P3       | Minor issue, cosmetic                   | 1 day    | 1 month    |
| P4       | Nice to have, improvement               | 1 week   | Backlog    |

#### Type Labels

| Label               | Description                    |
| ------------------- | ------------------------------ |
| `bug`               | Something is broken            |
| `security`          | Security vulnerability         |
| `performance`       | Performance degradation        |
| `regression`        | Previously working, now broken |
| `customer-reported` | From customer support          |

#### Component Labels

| Label        | Owner Team     |
| ------------ | -------------- |
| `billing`    | Platform       |
| `auth`       | Platform       |
| `namespaces` | Platform       |
| `console`    | Frontend       |
| `api`        | Platform       |
| `infra`      | Infrastructure |

### Triage Process

#### Daily Triage Meeting (15 min)

**Attendees**: On-call engineer, Tech Lead, PM (optional)

**Agenda**:

1. Review new issues (5 min)
2. Assign severity and owner (5 min)
3. Review P0/P1 progress (5 min)

#### Triage Workflow

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   New       │────▶│   Triage    │────▶│  Assigned   │
│   Issue     │     │   Meeting   │     │             │
└─────────────┘     └─────────────┘     └─────────────┘
                                               │
                    ┌─────────────┐             │
                    │   Closed    │◀────────────┤
                    │  (Invalid)  │             │
                    └─────────────┘             ▼
                                        ┌─────────────┐
┌─────────────┐     ┌─────────────┐     │    In       │
│   Closed    │◀────│   Review    │◀────│  Progress   │
│  (Resolved) │     │             │     │             │
└─────────────┘     └─────────────┘     └─────────────┘
```

### Issue Template

```markdown
## Bug Report

### Description

[Clear description of the issue]

### Steps to Reproduce

1.
2.
3.

### Expected Behavior

[What should happen]

### Actual Behavior

[What actually happens]

### Environment

- Account ID:
- Namespace:
- Region:
- SDK Version:
- Browser (if UI):

### Logs/Screenshots

[Attach relevant logs or screenshots]

### Impact

[Number of customers affected, revenue impact if known]
```

### Bug Fix Workflow

#### 1. Reproduce

```bash
# Create test case that reproduces the bug
func TestBug123_NamespaceCreationFails(t *testing.T) {
    // This should fail until bug is fixed
}
```

#### 2. Root Cause Analysis

Document in the issue:

- What went wrong
- Why it wasn't caught earlier
- Related code/commits

#### 3. Fix

- Branch from `cloud/develop`
- Implement fix
- Add regression test
- Update any affected documentation

#### 4. Review

- Code review required
- QA verification required for P0/P1
- Security review if security-related

#### 5. Deploy

- P0/P1: Hotfix process (immediate)
- P2/P3: Next regular release
- P4: When convenient

### Escalation Path

```
On-Call Engineer
      │
      ▼
Tech Lead (after 30 min for P0)
      │
      ▼
Engineering Manager (after 1 hour for P0)
      │
      ▼
VP Engineering (after 2 hours for P0)
```

### Metrics

#### Bug Metrics (Monthly Review)

- **MTTR** (Mean Time To Resolve): Target < 24h for P1
- **Bug Escape Rate**: Bugs found in production vs testing
- **Regression Rate**: Bugs that reoccur
- **Customer-Reported Ratio**: % of bugs from customers

#### SLA Compliance

Track resolution time against severity SLAs.
Alert if approaching SLA breach.

### Post-Mortem

Required for all P0 and P1 bugs.

#### Template

```markdown
## Post-Mortem: [Title]

### Summary

[1-2 sentence summary]

### Timeline

- HH:MM - Issue detected
- HH:MM - Investigation started
- HH:MM - Root cause identified
- HH:MM - Fix deployed
- HH:MM - Confirmed resolved

### Root Cause

[Technical explanation]

### Impact

- Duration: X hours
- Customers affected: N
- Revenue impact: $X

### What Went Well

-

### What Went Wrong

-

### Action Items

- [ ] [Action 1] - Owner - Due Date
- [ ] [Action 2] - Owner - Due Date
```

<!-- Source: capacity-planning.md -->

## Capacity Planning

### Capacity Model

#### Key Metrics

| Metric           | Unit  | Current | 6mo Forecast | 12mo Forecast |
| ---------------- | ----- | ------- | ------------ | ------------- |
| Organizations    | count | 1,000   | 2,500        | 5,000         |
| Namespaces       | count | 10,000  | 25,000       | 50,000        |
| Actions/sec      | rate  | 100K    | 300K         | 1M            |
| Active Storage   | GB    | 500     | 1,500        | 5,000         |
| Retained Storage | TB    | 10      | 30           | 100           |

#### Resource Mapping

| Workload   | Primary Bottleneck | Scaling Factor                    |
| ---------- | ------------------ | --------------------------------- |
| Workflows  | History service    | 1 pod per 1K concurrent workflows |
| Activities | Matching service   | 1 pod per 10K activities/sec      |
| Storage    | PostgreSQL IOPS    | Scale vertically + read replicas  |
| Search     | Elasticsearch      | 1 node per 10K queries/sec        |

### Infrastructure Sizing

#### Per-Region Requirements

##### Small (< 10K namespaces)

```yaml
eks:
  nodes: 10 x m6i.xlarge
temporal:
  frontend: 3 replicas
  history: 3 replicas
  matching: 3 replicas
  worker: 2 replicas
database:
  instance: db.r6g.xlarge
  storage: 500GB
  iops: 3000
redis:
  instance: cache.r6g.large
  nodes: 3
```

##### Medium (10K-50K namespaces)

```yaml
eks:
  nodes: 25 x m6i.2xlarge
temporal:
  frontend: 6 replicas
  history: 9 replicas
  matching: 6 replicas
  worker: 4 replicas
database:
  instance: db.r6g.2xlarge
  storage: 2TB
  iops: 10000
  read_replicas: 2
redis:
  instance: cache.r6g.xlarge
  nodes: 6
```

##### Large (50K+ namespaces)

```yaml
eks:
  nodes: 50 x m6i.4xlarge
temporal:
  frontend: 12 replicas
  history: 18 replicas
  matching: 12 replicas
  worker: 8 replicas
database:
  instance: db.r6g.4xlarge
  storage: 10TB
  iops: 30000
  read_replicas: 4
redis:
  instance: cache.r6g.2xlarge
  nodes: 9
```

### Scaling Triggers

#### Automatic Scaling

```yaml
# HPA configuration
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: temporal-frontend
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: temporal-frontend
  minReplicas: 3
  maxReplicas: 20
  metrics:
    - type: Resource
      resource:
        name: cpu
        target:
          type: Utilization
          averageUtilization: 70
    - type: Pods
      pods:
        metric:
          name: grpc_requests_per_second
        target:
          type: AverageValue
          averageValue: 1000
```

#### Manual Scaling Thresholds

| Metric               | Warning | Action           |
| -------------------- | ------- | ---------------- |
| CPU > 70% sustained  | Alert   | Scale pods       |
| Memory > 80%         | Alert   | Scale pods       |
| DB CPU > 70%         | Alert   | Scale instance   |
| DB connections > 80% | Alert   | Scale instance   |
| Disk > 80%           | Alert   | Increase storage |
| Queue depth > 1000   | Alert   | Scale workers    |

### Capacity Review

#### Weekly Review

1. Review resource utilization dashboards
2. Check growth trends
3. Identify hotspots
4. Update forecasts

#### Monthly Review

1. Compare actual vs forecast
2. Adjust resource allocations
3. Plan next quarter infrastructure
4. Budget review

#### Quarterly Planning

1. Update 6/12 month forecasts
2. Plan major infrastructure changes
3. Reserved instance purchasing
4. Budget approval

### Cost Optimization

#### Reserved Instances

| Resource    | On-Demand | 1yr Reserved | 3yr Reserved | Recommendation      |
| ----------- | --------- | ------------ | ------------ | ------------------- |
| EKS nodes   | $0.192/hr | $0.120/hr    | $0.080/hr    | 1yr for stable load |
| RDS         | $0.456/hr | $0.285/hr    | $0.190/hr    | 1yr for prod        |
| ElastiCache | $0.156/hr | $0.097/hr    | $0.065/hr    | 1yr for prod        |

#### Right-Sizing

Weekly job to identify:

- Underutilized instances (< 20% CPU)
- Oversized storage
- Unused resources

```bash
# Generate right-sizing report
aws compute-optimizer get-ec2-instance-recommendations \
  --output json > rightsizing-report.json
```

### Disaster Capacity

#### Reserve Capacity

Maintain 50% headroom in secondary regions for failover.

| Region       | Primary Capacity | Reserve Capacity |
| ------------ | ---------------- | ---------------- |
| us-east-1    | 100%             | 0%               |
| us-west-2    | 50%              | 50% (failover)   |
| eu-west-1    | 100%             | 0%               |
| eu-central-1 | 50%              | 50% (failover)   |

#### Burst Capacity

Use spot instances for non-critical workloads:

- Batch processing
- Analytics
- Development environments

### Monitoring

#### Capacity Dashboard

Key panels:

1. Resource utilization by service
2. Growth trends (7d, 30d, 90d)
3. Forecast vs actual
4. Cost per customer
5. Scaling events

#### Alerts

```yaml
alerts:
  - name: CapacityWarning
    condition: |
      avg(cpu_utilization) > 70% for 1h
      OR avg(memory_utilization) > 80% for 30m
      OR disk_utilization > 80%
    severity: warning

  - name: CapacityCritical
    condition: |
      avg(cpu_utilization) > 90% for 15m
      OR avg(memory_utilization) > 95% for 15m
      OR disk_utilization > 95%
    severity: critical
```

### Runbook: Emergency Scaling

When capacity limits are hit:

1. **Immediate** (< 5 min)

   - Scale pods via HPA override
   - Enable autoscaling if disabled

2. **Short-term** (< 1 hour)

   - Add nodes to cluster
   - Scale database read replicas

3. **Medium-term** (< 1 day)
   - Scale database instance
   - Increase storage
   - Add new region if needed

<!-- Source: cdn.md -->

## Content Delivery Network (CDN)

### Strategy

Temporal Cloud is largely API-based (dynamic), but static assets and global routing benefit from CDN.

### Static Assets

#### Cloud Console

- **HTML/JS/CSS**: Cached at edge (Cloudflare/CloudFront).
- **TTL**: 1 year (immutable hashes).
- **Invalidation**: On deployment.

#### Documentation

- Cached globally.
- Stale-while-revalidate strategy.

### Dynamic Routing (Global Accelerator)

Accelerate gRPC/API traffic using AWS Global Accelerator or Cloudflare Spectrum.

1. **Anycast IP**: User connects to nearest edge POP.
2. **Backbone**: Traffic traverses provider's private fiber, not public internet.
3. **Origin**: Handoff to nearest regional ALB.

**Benefit**: ~30% latency reduction for cross-continent API calls.

### Security at Edge

- **WAF**: Block threats at edge before they hit origin.
- **DDoS**: Absorb volumetric attacks.
- **TLS Termination**: Offload crypto handshake to edge.

### Configuration

#### AWS CloudFront

```hcl
resource "aws_cloudfront_distribution" "console" {
  origin {
    domain_name = aws_s3_bucket.frontend.bucket_regional_domain_name
    origin_id   = "s3-console"
  }

  # SPA routing
  custom_error_response {
    error_code         = 404
    response_code      = 200
    response_page_path = "/index.html"
  }
}
```

#### AWS Global Accelerator

```hcl
resource "aws_globalaccelerator_accelerator" "api" {
  name            = "temporal-cloud-api"
  ip_address_type = "IPV4"
  enabled         = true
}

resource "aws_globalaccelerator_listener" "api" {
  accelerator_arn = aws_globalaccelerator_accelerator.api.id
  protocol        = "TCP"

  port_range {
    from_port = 443
    to_port   = 443
  }
}
```

<!-- Source: certificates.md -->

## Certificate Management

### mTLS Authentication

Temporal Cloud uses mutual TLS (mTLS) for namespace authentication. Clients must present a valid certificate signed by a CA registered with the namespace.

### Certificate Requirements

#### CA Certificates

- Format: X.509 v3
- Encoding: PEM
- Key size: RSA 2048+ or ECDSA P-256+
- Max per namespace: 16 certificates or 32KB total

#### End-Entity Certificates

- Must be signed by registered CA
- Validity: Recommended max 1 year
- Must include Extended Key Usage: Client Authentication

### Certificate Operations

| Operation               | Console | tcld | API | Terraform |
| ----------------------- | ------- | ---- | --- | --------- |
| Add CA certificate      | ✅      | ✅   | ✅  | ✅        |
| Remove CA certificate   | ✅      | ✅   | ✅  | ✅        |
| List certificates       | ✅      | ✅   | ✅  | ✅        |
| Set certificate filters | ✅      | ✅   | ✅  | ✅        |

### Certificate Filters

Fine-grained access control based on certificate attributes.

#### Filter Types

| Filter     | Description          | Example              |
| ---------- | -------------------- | -------------------- |
| Subject CN | Common Name match    | `*.example.com`      |
| Subject OU | Organizational Unit  | `engineering`        |
| SAN DNS    | DNS Subject Alt Name | `worker.example.com` |

#### Filter Configuration

```json
{
  "filters": [
    {
      "type": "subject_cn",
      "value": "*.prod.example.com"
    },
    {
      "type": "subject_ou",
      "value": "production"
    }
  ]
}
```

### Generating Certificates

#### Using OpenSSL

```bash
# Generate CA key and certificate
openssl genrsa -out ca.key 4096
openssl req -new -x509 -days 365 -key ca.key -out ca.pem \
  -subj "/CN=Temporal CA/O=MyOrg"

# Generate client key and CSR
openssl genrsa -out client.key 4096
openssl req -new -key client.key -out client.csr \
  -subj "/CN=temporal-worker/O=MyOrg"

# Sign client certificate
openssl x509 -req -days 365 -in client.csr \
  -CA ca.pem -CAkey ca.key -CAcreateserial \
  -out client.pem
```

#### Using tcld

```bash
tcld generate-certificates \
  --namespace my-namespace \
  --ca-cert ca.pem \
  --output-dir ./certs
```

### Certificate Rotation

#### Process

1. Generate new CA certificate
2. Add new CA to namespace (both old and new active)
3. Update all workers with new client certificates
4. Remove old CA certificate
5. Monitor for authentication failures

#### Zero-Downtime Rotation

```bash
# Step 1: Add new CA (keep old)
tcld namespace certificates add \
  --namespace my-namespace \
  --ca-certificate new-ca.pem

# Step 2: Update workers (deploy with new certs)

# Step 3: Remove old CA
tcld namespace certificates remove \
  --namespace my-namespace \
  --ca-certificate-fingerprint <old-fingerprint>
```

### Expiration Monitoring

#### Check Expiration

```bash
# Check certificate expiration
openssl x509 -enddate -noout -in ca.pem

# Output: notAfter=Jan  1 00:00:00 2026 GMT
```

#### Alerts

- 30 days before expiry: Warning notification
- 7 days before expiry: Critical notification
- On expiry: Certificate rejected

### Troubleshooting

#### Certificate Rejected

1. Verify CA is registered with namespace
2. Check certificate not expired
3. Verify certificate chain is complete
4. Check certificate filters match

#### Connection Failed

```bash
# Test connection with certificate
grpcurl -cert client.pem -key client.key -cacert ca.pem \
  my-namespace.tmprl.cloud:443 \
  temporal.api.workflowservice.v1.WorkflowService/GetSystemInfo
```

<!-- Source: cicd.md -->

## CI/CD Pipeline

### Pipeline Overview

```
┌─────────┐   ┌─────────┐   ┌─────────┐   ┌─────────┐   ┌─────────┐
│  Lint   │──▶│  Test   │──▶│  Build  │──▶│ Deploy  │──▶│ Verify  │
│         │   │         │   │         │   │ Staging │   │         │
└─────────┘   └─────────┘   └─────────┘   └─────────┘   └─────────┘
                                                │
                                                ▼
                                          ┌─────────┐
                                          │ Deploy  │
                                          │  Prod   │
                                          └─────────┘
```

### Workflow Files

#### ci.yaml

```yaml
name: CI
on:
  push:
    branches: [cloud/main, cloud/develop]
  pull_request:
    branches: [cloud/main]

jobs:
  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: "1.22"
      - run: make lint

  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: "1.22"
      - run: make test

  build:
    needs: [lint, test]
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - run: make build
      - uses: docker/build-push-action@v5
        with:
          push: ${{ github.ref == 'refs/heads/cloud/main' }}
          tags: ${{ env.REGISTRY }}/${{ env.IMAGE }}:${{ github.sha }}
```

#### deploy-staging.yaml

```yaml
name: Deploy Staging
on:
  push:
    branches: [cloud/develop]

jobs:
  deploy:
    runs-on: ubuntu-latest
    environment: staging
    steps:
      - uses: actions/checkout@v4
      - uses: aws-actions/configure-aws-credentials@v4
      - run: |
          helm upgrade --install cloud-platform ./charts/cloud-platform \
            --namespace cloud-platform \
            --set image.tag=${{ github.sha }}
```

#### deploy-prod.yaml

```yaml
name: Deploy Production
on:
  workflow_dispatch:
    inputs:
      version:
        description: "Version to deploy"
        required: true

jobs:
  deploy:
    runs-on: ubuntu-latest
    environment: production
    steps:
      - uses: actions/checkout@v4
      - uses: aws-actions/configure-aws-credentials@v4
      - run: |
          helm upgrade --install cloud-platform ./charts/cloud-platform \
            --namespace cloud-platform \
            --set image.tag=${{ inputs.version }}
```

#### sync-upstream.yaml

```yaml
name: Sync Upstream
on:
  schedule:
    - cron: "0 0 * * *"
  workflow_dispatch:

jobs:
  sync:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0
      - run: |
          git remote add upstream https://github.com/temporalio/temporal.git
          git fetch upstream
          git checkout cloud/main
          git merge upstream/main -m "Sync upstream"
          git push origin cloud/main
```

### Environments

| Environment | Trigger               | Approval    | URL                       |
| ----------- | --------------------- | ----------- | ------------------------- |
| Development | Push to feature/\*    | None        | dev.temporal-cloud.io     |
| Staging     | Push to cloud/develop | None        | staging.temporal-cloud.io |
| Production  | Manual                | 2 approvers | temporal-cloud.io         |

### Deployment Strategy

#### Rolling Update

```yaml
strategy:
  type: RollingUpdate
  rollingUpdate:
    maxSurge: 1
    maxUnavailable: 0
```

#### Canary (Future)

- Deploy to 10% of traffic
- Monitor error rates
- Gradually increase to 100%
- Automatic rollback on errors

### Rollback Procedure

#### Automatic Rollback

```yaml
# Helm rollback on failed deployment
- run: |
    helm upgrade ... || helm rollback cloud-platform
```

#### Manual Rollback

```bash
# List releases
helm history cloud-platform -n cloud-platform

# Rollback to previous
helm rollback cloud-platform 1 -n cloud-platform
```

### Secrets Management

#### GitHub Secrets

- `AWS_ACCESS_KEY_ID`
- `AWS_SECRET_ACCESS_KEY`
- `DOCKER_REGISTRY_TOKEN`
- `SLACK_WEBHOOK_URL`

#### Runtime Secrets

- Stored in AWS Secrets Manager
- Injected via External Secrets Operator
- Rotated every 90 days

<!-- Source: cloud-ops-api.md -->

## Cloud Ops API

### Overview

The Cloud Ops API is a gRPC service for managing Temporal Cloud resources programmatically. It is the foundation for the Console, CLI (`tcld`), and Terraform provider.

### Service Definition

#### Endpoint

`saas-api.tmprl.cloud:443`

#### Authentication

- Bearer Token (API Key)
- mTLS (Service Identity)

### Proto Definitions

#### Service: `AccountService`

Manage account-level settings and users.

- `GetAccount`
- `UpdateAccount`
- `ListUsers`
- `InviteUser`
- `UpdateUser`
- `DeleteUser`

#### Service: `NamespaceService`

Manage namespaces.

- `CreateNamespace`
- `GetNamespace`
- `UpdateNamespace`
- `DeleteNamespace`
- `ListNamespaces`
- `FailoverNamespace`
- `AddNamespaceRegion`

#### Service: `AccessService`

Manage permissions and API keys.

- `CreateAPIKey`
- `ListAPIKeys`
- `UpdateAPIKey`
- `RotateAPIKey`
- `DeleteAPIKey`

#### Service: `UsageService`

Retrieve usage data.

- `GetUsageSummary` (Daily/Monthly)
- `GetRequestUsage` (Detailed)

### Rate Limiting

| Scope            | Limit    | Burst |
| ---------------- | -------- | ----- |
| Per Account      | 200 RPS  | 300   |
| Per User         | 20 RPS   | 40    |
| Read Operations  | 1000 RPS | 2000  |
| Write Operations | 50 RPS   | 100   |

### Error Handling

Standard gRPC error codes:

- `INVALID_ARGUMENT` (400): Validation failed
- `UNAUTHENTICATED` (401): Invalid/missing token
- `PERMISSION_DENIED` (403): Insufficient role
- `NOT_FOUND` (404): Resource doesn't exist
- `RESOURCE_EXHAUSTED` (429): Rate limit exceeded
- `UNAVAILABLE` (503): Maintenance or outage

### Idempotency

All write operations support `request_id` field.

- Clients MUST generate a UUID for `request_id`.
- Server stores result of operations for 24 hours.
- Retrying with same `request_id` returns the original result without re-executing.

<!-- Source: compliance-automation.md -->

## Compliance Automation

### Overview

Automated evidence collection and compliance checks for SOC 2 Type II, GDPR, and other regulatory requirements.

### Tagging Strategy

All Terraform resources must have compliance tags:

```hcl
locals {
  common_tags = {
    Environment        = var.environment
    Project            = "temporal-cloud"
    Compliance         = "SOC2"
    DataClassification = "Confidential"
    Owner              = "platform-team"
    ManagedBy          = "terraform"
  }
}

resource "aws_db_instance" "main" {
  # ... configuration ...
  tags = merge(local.common_tags, {
    DataClassification = "Restricted"
    BackupRequired     = "true"
  })
}
```

#### Data Classification Levels

| Level        | Description        | Examples                   |
| ------------ | ------------------ | -------------------------- |
| Public       | No restrictions    | Marketing content          |
| Internal     | Internal use only  | Architecture docs          |
| Confidential | Business sensitive | Usage metrics              |
| Restricted   | Highly sensitive   | Customer data, credentials |

### SOC 2 Control Mapping

#### CC6.1 - Logical Access Controls

| Control        | Implementation | Evidence                 |
| -------------- | -------------- | ------------------------ |
| Authentication | SAML SSO, MFA  | IAM policies, SSO config |
| Authorization  | RBAC           | Role assignments         |
| Access review  | Quarterly      | Access review reports    |

#### CC6.6 - Encryption

| Control         | Implementation | Evidence                 |
| --------------- | -------------- | ------------------------ |
| Data at rest    | AES-256        | RDS/S3 encryption config |
| Data in transit | TLS 1.3        | ALB/cert config          |
| Key management  | AWS KMS        | KMS key policies         |

#### CC7.2 - Monitoring

| Control             | Implementation        | Evidence        |
| ------------------- | --------------------- | --------------- |
| Security monitoring | CloudTrail, GuardDuty | Alert configs   |
| Anomaly detection   | CloudWatch Anomaly    | Detection rules |
| Incident response   | PagerDuty             | Runbooks        |

### Automated Compliance Checks

#### Daily CI Job

```yaml
# .github/workflows/compliance-check.yaml
name: Compliance Check
on:
  schedule:
    - cron: "0 6 * * *" # Daily at 6 AM UTC
  workflow_dispatch:

jobs:
  compliance:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Setup AWS credentials
        uses: aws-actions/configure-aws-credentials@v4
        with:
          role-to-assume: ${{ secrets.COMPLIANCE_ROLE_ARN }}
          aws-region: us-east-1

      - name: Run compliance checks
        run: |
          python scripts/compliance-check.py \
            --output compliance-report.json

      - name: Upload report
        uses: actions/upload-artifact@v4
        with:
          name: compliance-report
          path: compliance-report.json

      - name: Fail on critical issues
        run: |
          CRITICAL=$(jq '.critical_count' compliance-report.json)
          if [ "$CRITICAL" -gt 0 ]; then
            echo "Critical compliance issues found!"
            exit 1
          fi
```

#### Compliance Check Script

```python
# scripts/compliance-check.py
import boto3
import json
from datetime import datetime, timedelta

class ComplianceChecker:
    def __init__(self):
        self.results = {
            "timestamp": datetime.utcnow().isoformat(),
            "checks": [],
            "critical_count": 0,
            "warning_count": 0
        }

    def check_rds_encryption(self):
        """CC6.6 - All RDS instances must be encrypted"""
        rds = boto3.client('rds')
        instances = rds.describe_db_instances()

        for db in instances['DBInstances']:
            encrypted = db.get('StorageEncrypted', False)
            self.add_result(
                control="CC6.6",
                resource=db['DBInstanceIdentifier'],
                check="RDS Encryption",
                passed=encrypted,
                severity="critical" if not encrypted else "pass"
            )

    def check_s3_encryption(self):
        """CC6.6 - All S3 buckets must have encryption"""
        s3 = boto3.client('s3')
        buckets = s3.list_buckets()

        for bucket in buckets['Buckets']:
            try:
                encryption = s3.get_bucket_encryption(Bucket=bucket['Name'])
                encrypted = True
            except:
                encrypted = False

            self.add_result(
                control="CC6.6",
                resource=bucket['Name'],
                check="S3 Encryption",
                passed=encrypted,
                severity="critical" if not encrypted else "pass"
            )

    def check_s3_public_access(self):
        """CC6.1 - No S3 buckets should be public"""
        s3 = boto3.client('s3')
        buckets = s3.list_buckets()

        for bucket in buckets['Buckets']:
            try:
                acl = s3.get_bucket_acl(Bucket=bucket['Name'])
                public = any(
                    grant['Grantee'].get('URI', '').endswith('AllUsers')
                    for grant in acl['Grants']
                )
            except:
                public = False

            self.add_result(
                control="CC6.1",
                resource=bucket['Name'],
                check="S3 Public Access",
                passed=not public,
                severity="critical" if public else "pass"
            )

    def check_backup_age(self):
        """CC7.1 - Backups must be recent"""
        rds = boto3.client('rds')
        snapshots = rds.describe_db_snapshots(SnapshotType='automated')

        for snapshot in snapshots['DBSnapshots']:
            age = datetime.utcnow() - snapshot['SnapshotCreateTime'].replace(tzinfo=None)
            recent = age < timedelta(hours=24)

            self.add_result(
                control="CC7.1",
                resource=snapshot['DBSnapshotIdentifier'],
                check="Backup Age",
                passed=recent,
                severity="warning" if not recent else "pass",
                details=f"Age: {age}"
            )

    def check_iam_mfa(self):
        """CC6.1 - All IAM users must have MFA"""
        iam = boto3.client('iam')
        users = iam.list_users()

        for user in users['Users']:
            mfa_devices = iam.list_mfa_devices(UserName=user['UserName'])
            has_mfa = len(mfa_devices['MFADevices']) > 0

            self.add_result(
                control="CC6.1",
                resource=user['UserName'],
                check="IAM MFA",
                passed=has_mfa,
                severity="critical" if not has_mfa else "pass"
            )

    def check_unused_credentials(self):
        """CC6.1 - No unused credentials > 90 days"""
        iam = boto3.client('iam')
        report = iam.generate_credential_report()
        # ... parse and check

    def add_result(self, control, resource, check, passed, severity, details=None):
        result = {
            "control": control,
            "resource": resource,
            "check": check,
            "passed": passed,
            "severity": severity,
            "details": details
        }
        self.results["checks"].append(result)

        if severity == "critical" and not passed:
            self.results["critical_count"] += 1
        elif severity == "warning" and not passed:
            self.results["warning_count"] += 1

    def run_all_checks(self):
        self.check_rds_encryption()
        self.check_s3_encryption()
        self.check_s3_public_access()
        self.check_backup_age()
        self.check_iam_mfa()
        return self.results

if __name__ == "__main__":
    checker = ComplianceChecker()
    results = checker.run_all_checks()
    print(json.dumps(results, indent=2))
```

### Evidence Collection

#### Automated Evidence

| Evidence            | Collection Method | Frequency  |
| ------------------- | ----------------- | ---------- |
| IAM policies        | AWS Config        | Daily      |
| Encryption status   | Compliance script | Daily      |
| Access logs         | CloudTrail → S3   | Continuous |
| Change history      | Terraform state   | On change  |
| Vulnerability scans | Trivy/Snyk        | On build   |

#### Manual Evidence

| Evidence         | Owner           | Frequency |
| ---------------- | --------------- | --------- |
| Access reviews   | Security team   | Quarterly |
| Penetration test | Third party     | Annual    |
| Policy reviews   | Compliance team | Annual    |
| Training records | HR              | Annual    |

### GDPR Compliance

#### Data Subject Rights

| Right         | Implementation        |
| ------------- | --------------------- |
| Access        | Export API endpoint   |
| Erasure       | Delete workflow       |
| Portability   | Export in JSON format |
| Rectification | Update API            |

#### Data Processing Records

```sql
CREATE TABLE data_processing_records (
    id UUID PRIMARY KEY,
    organization_id UUID NOT NULL,
    purpose VARCHAR(255) NOT NULL,
    legal_basis VARCHAR(100) NOT NULL,
    data_categories TEXT[],
    retention_period INTERVAL,
    created_at TIMESTAMPTZ NOT NULL
);
```

### Audit Trail

All compliance-relevant events are logged:

```json
{
  "event_type": "compliance_check",
  "timestamp": "2025-01-01T00:00:00Z",
  "check_type": "rds_encryption",
  "resource": "temporal-cloud-prod",
  "result": "pass",
  "control": "CC6.6",
  "evidence_url": "s3://compliance-evidence/2025/01/01/rds-encryption.json"
}
```

### Reporting

#### Monthly Compliance Report

Generated automatically and sent to:

- Security team
- Compliance officer
- Engineering leadership

#### Audit Preparation

Before SOC 2 audit:

1. Run full compliance check suite
2. Generate evidence package
3. Review and remediate findings
4. Prepare documentation index

<!-- Source: console-design.md -->

## Cloud Console Design

### Architecture

- **Framework**: Next.js 14 (App Router)
- **Styling**: Tailwind CSS + shadcn/ui
- **State**: TanStack Query + Zustand
- **Auth**: JWT in HttpOnly cookie
- **API**: gRPC-Web (Connect)

### Site Map

```
/login
/sso/callback
/dashboard (Redirects to first account)
/accounts/{account_id}/
├── overview               # Usage, health
├── namespaces/
│   ├── [list]
│   ├── create
│   └── {namespace_id}/
│       ├── overview       # Metrics
│       ├── workflows/     # Web UI
│       ├── settings       # Retention, CA certs
│       └── limits         # Quotas
├── members/               # Users & Roles
├── service-accounts/
├── billing/
│   ├── overview           # Current usage & estimated cost
│   ├── invoices/          # Past invoices
│   └── settings           # Payment methods, address
├── settings/
    ├── general
    ├── sso                # SAML config
    ├── audit-logs         # Export config
    └── api-keys
```

### Key Screens

#### 1. Namespace List

**Columns**: Status, Name, Region, Retention, Active Workflows, Actions/sec (24h avg).
**Actions**: Create Namespace, Search.

#### 2. Namespace Detail > Overview

**Charts** (Recharts):

- Success Rate (Line)
- Latency (Line)
- Actions/sec (Area)
  **Info**: Certificates (expiry warning), Region, Grpc Endpoint (copy button).

#### 3. Create Namespace Wizard

**Step 1: Basics**: Name, Region selection (with map/latency hints).
**Step 2: Config**: Retention period slider (1-90 days).
**Step 3: Security**: Upload CA certificate (drag & drop).
**Step 4: Review**: Summary & "Create" button.

#### 4. Billing Overview

**Header**: Current Month-to-Date Cost.
**Breakdown**:

- Plan Base Fee
- Actions Overage (Progress bar vs included amount)
- Storage (Active vs Retained)
  **Usage Chart**: Daily bar chart of actions.

#### 5. SSO Configuration

**Status**: Enabled/Disabled toggle.
**Config**:

- IdP Metadata URL input.
- Manual file upload.
  **Attribute Mapping**:
- Role mapping table (IdP Group -> Cloud Role).

### UX Patterns

- **Loading**: Skeleton loaders (shadcn/ui `Skeleton`).
- **Errors**: Toast notifications (Sonner) + inline error boundaries.
- **Copying**: Click-to-copy for IDs, keys, endpoints.
- **Time**: All times in UTC by default, toggle for Local.
- **Dates**: Relative ("2 hours ago") with tooltip for absolute.

### Theming

- **Mode**: Dark/Light mode toggle (default Dark for dev tools).
- **Brand Colors**:
  - Primary: Temporal Black/White.
  - Accents: Blue for links, Green for success, Red for errors.

<!-- Source: cost-optimization.md -->

## Cost Optimization Strategy

### Philosophy

**Pay for value, not waste.** Optimize unit economics without sacrificing reliability.

### Compute Optimization

#### 1. Spot Instances / Preemptible VMs

Use Spot instances for stateless workloads:

- **Worker Service**: 100% Spot (stateless, auto-recovering)
- **Frontend Service**: 50% Spot (behind load balancer)
- **Matching Service**: On-Demand (stateful in-memory)
- **History Service**: On-Demand (critical state)

**Savings**: ~70% on compute.

#### 2. Graviton (ARM64) Migration

Migrate all services to AWS Graviton3 (m7g, r7g instances).

- Temporal Go code supports ARM64.
- **Savings**: ~20% price-performance improvement.

#### 3. Kubernetes Right-Sizing

- **Vertical Pod Autoscaler (VPA)**: Recommend request/limits based on actual usage.
- **Karpenter**: Bin-pack pods onto optimal node sizes.
- **Over-provisioning**: Keep buffer small (10%).

### Storage Optimization

#### 1. Tiered Storage (S3)

Move workflow history blobs to cheaper storage tiers.

- **Standard**: Hot data (0-30 days)
- **Intelligent Tiering**: Warm data (30-90 days)
- **Glacier Instant Retrieval**: Cold data (90+ days)

**Savings**: ~40-60% on long-term retention.

#### 2. Database pruning

- Aggressively prune `task_executions` and completed workflow records (transfer to S3).
- Use partial indexes to reduce index size.

#### 3. EBS Volume Types

- Use **gp3** instead of gp2 (20% cheaper per GB).
- Monitor IOPS usage and provision accurately.

### Networking Optimization

#### 1. Keep Traffic Local

- Ensure AZ affinity: Frontend -> History in same AZ preferred.
- Cross-AZ traffic costs $0.01/GB.
- **Topology Aware Hints** in K8s.

#### 2. Endpoint Services (PrivateLink)

- Avoid NAT Gateway data processing charges ($0.045/GB).
- Use VPC Endpoints for S3, DynamoDB, etc.

### Observability Cost Control

#### 1. Metric Cardinality

- Drop high-cardinality labels (workflowID, runID) from metrics.
- Whitelist only essential metrics for Datadog/Prometheus.

#### 2. Log Sampling

- Sample INFO logs at 10%.
- Keep ERROR/WARN at 100%.
- Use structural logging to allow backend filtering.

### Governance

#### 1. Budgets & Alerts

- Set AWS Budgets per team/service.
- Alert when forecast exceeds budget by 10%.

#### 2. Tagging Policy

Every resource MUST have:

- `CostCenter`
- `Service`
- `Environment`
- `Owner`

Resources without tags are auto-terminated after warning.

### Cost Dashboard

**Unit Metrics**:

- Cost per million Actions
- Cost per GB-hour Storage
- Cost per Active Workflow

If unit cost increases >10%, trigger investigation.

<!-- Source: crash-proofing.md -->

## Crash Proofing & Reliability

### Circuit Breakers

Protect the system from cascading failures.

```go
// Client-side circuit breaker (e.g., in Cloud API calling Billing)
cb := circuitbreaker.New(circuitbreaker.Options{
    Name:        "billing-service",
    MaxRequests: 100,
    Interval:    5 * time.Second,
    Timeout:     30 * time.Second,
    ReadyToTrip: func(counts circuitbreaker.Counts) bool {
        return counts.ConsecutiveFailures > 5
    },
})

result, err := cb.Execute(func() (interface{}, error) {
    return billingClient.Charge(ctx, req)
})
```

### Bulkheading

Isolate failures to specific components or customers.

#### 1. Per-Namespace Isolation

- Shard heavy namespaces to dedicated Task Queues.
- In extreme cases, move noisy neighbor to dedicated History hosts (Isolation Groups).

#### 2. Thread Pools

- Separate thread pools for:
  - Critical (API, Membership)
  - User (StartWorkflow, Signal)
  - Background (Replication, Retention)

### Graceful Degradation

When system is overloaded:

1. **Shed Load**: Reject low-priority traffic (e.g., `ListWorkflows`, `Query`) with `503 Service Unavailable`. Keep `StartWorkflow` and `Signal` working.
2. **Disable Features**: Turn off Visibility (search) updates if Elasticsearch is down.
3. **Increase Latency**: Slow down background replication to save IOPS for active traffic.

### Chaos Engineering

#### Scheduled Drills

- **Simian Army**: Randomly kill pods (daily).
- **Network Partition**: Block traffic between AZs (monthly).
- **Latency Injection**: Add 100ms lag to DB calls (quarterly).

#### "Game Days"

Quarterly manual drills:

1. Kill primary region.
2. Corrupt database WAL.
3. Expire root CA certificate.

### Recovery Oriented Computing

#### Fast Restart

- Optimize startup time (< 10s).
- Lazy load caches.

#### Stateless Frontends

- Frontend service must be killable at any time with zero impact (drains connections).

#### Idempotency

- ALL write APIs must be idempotent.
- Clients must retry indefinitely on transient errors.

<!-- Source: customer-onboarding.md -->

## Customer Onboarding

### Onboarding Flow

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Signup    │────▶│   Verify    │────▶│   Setup     │
│   Form      │     │   Email     │     │   Account   │
└─────────────┘     └─────────────┘     └─────────────┘
                                               │
                                               ▼
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   First     │◀────│   Create    │◀────│   Select    │
│   Workflow  │     │   Namespace │     │   Plan      │
└─────────────┘     └─────────────┘     └─────────────┘
```

### Step 1: Signup

#### Form Fields

- Email (required)
- Full Name (required)
- Company Name (required)
- Password (8+ chars, 1 uppercase, 1 number)

#### Validations

- Email not already registered
- Email domain not blocked (disposable emails)
- Password strength check

#### Actions

1. Create user record (unverified)
2. Create organization record
3. Send verification email
4. Log signup event to analytics

### Step 2: Email Verification

#### Email Content

```
Subject: Verify your Temporal Cloud account

Hi {{name}},

Click the link below to verify your email:
{{verification_link}}

This link expires in 24 hours.
```

#### Verification Link

`https://cloud.temporal.io/verify?token={{jwt_token}}`

Token contains: user_id, email, expiry

#### Actions

1. Mark user as verified
2. Redirect to account setup

### Step 3: Account Setup

#### Organization Details

- Organization name (pre-filled)
- Organization slug (auto-generated, editable)
- Industry (dropdown)
- Company size (dropdown)

#### Actions

1. Update organization record
2. Create default settings
3. Proceed to plan selection

### Step 4: Plan Selection

#### Plan Display

```
┌─────────────────────────────────────────────────────────┐
│  Essential           Business           Enterprise      │
│  $100/mo             $500/mo            Contact Sales   │
│                                                         │
│  ✓ 1M actions        ✓ 2.5M actions     ✓ Custom        │
│  ✓ 1 GB active       ✓ 2.5 GB active    ✓ Custom        │
│  ✓ Email support     ✓ Chat support     ✓ Dedicated     │
│                      ✓ SSO              ✓ SCIM          │
│                                                         │
│  [Start Free Trial]  [Start Free Trial] [Contact Sales] │
└─────────────────────────────────────────────────────────┘
```

#### Free Trial

- 14 days on selected plan
- No credit card required
- Full feature access

#### Actions

1. Create subscription (trial status)
2. Set trial_ends_at = now + 14 days
3. Schedule trial expiry reminder workflow

### Step 5: Create First Namespace

#### Guided Wizard

```
Step 1 of 3: Name Your Namespace

  Namespace Name: [my-first-namespace]

  This will be part of your endpoint:
  my-first-namespace.abc123.tmprl.cloud

  [Next]
```

```
Step 2 of 3: Select Region

  ○ US East (Virginia)     - Recommended
  ○ US West (Oregon)
  ○ EU West (Ireland)
  ○ Asia Pacific (Singapore)

  [Next]
```

```
Step 3 of 3: Security

  Upload CA Certificate (optional)
  [Drop file here or click to upload]

  You can also generate certificates using our CLI:
  $ tcld generate-certificates --namespace my-first-namespace

  [Skip for now]  [Create Namespace]
```

#### Actions

1. Create namespace via Cloud Ops API
2. Wait for namespace to be ready (~30s)
3. Show success with connection details

### Step 6: First Workflow (Getting Started)

#### Connection Details

```
Your namespace is ready!

Endpoint: my-first-namespace.abc123.tmprl.cloud:443
Namespace: my-first-namespace

Quick Start:
$ temporal workflow start \
    --address my-first-namespace.abc123.tmprl.cloud:443 \
    --namespace my-first-namespace \
    --task-queue my-queue \
    --type MyWorkflow
```

#### Code Examples (Tabs)

- Go
- Java
- TypeScript
- Python

#### Next Steps Checklist

- [ ] Download and install SDK
- [ ] Run your first workflow
- [ ] Invite team members
- [ ] Set up billing (before trial ends)

### Onboarding Workflow (Backend)

```go
func CustomerOnboardingWorkflow(ctx workflow.Context, input OnboardingInput) error {
    // Step 1: Create user and org
    var user User
    workflow.ExecuteActivity(ctx, CreateUser, input).Get(ctx, &user)

    // Step 2: Send verification email
    workflow.ExecuteActivity(ctx, SendVerificationEmail, user.Email)

    // Step 3: Wait for verification (up to 7 days)
    verified := workflow.GetSignalChannel(ctx, "email-verified")

    selector := workflow.NewSelector(ctx)
    selector.AddReceive(verified, func(c workflow.ReceiveChannel, more bool) {
        // Continue onboarding
    })
    selector.AddFuture(workflow.NewTimer(ctx, 7*24*time.Hour), func(f workflow.Future) {
        // Send reminder or expire
    })
    selector.Select(ctx)

    // Step 4: Wait for first namespace
    workflow.ExecuteActivity(ctx, WaitForFirstNamespace, user.OrgID)

    // Step 5: Send welcome email with tips
    workflow.ExecuteActivity(ctx, SendWelcomeEmail, user.Email)

    // Step 6: Schedule trial reminders
    workflow.ExecuteActivity(ctx, ScheduleTrialReminders, input.TrialEndsAt)

    return nil
}
```

### Trial Reminders

| Day | Email                                                 |
| --- | ----------------------------------------------------- |
| 7   | "Your trial is halfway through"                       |
| 12  | "2 days left - add payment method"                    |
| 14  | "Trial ended - add payment to continue"               |
| 17  | "Account suspended - data will be deleted in 30 days" |

### Metrics

- Signup → Verified: Target 80%
- Verified → First Namespace: Target 70%
- First Namespace → First Workflow: Target 60%
- Trial → Paid Conversion: Target 25%

<!-- Source: data-model.md -->

## Data Model

### Entity Relationship Diagram

```
┌──────────────┐       ┌──────────────┐       ┌──────────────┐
│ Organization │───┬───│ Subscription │       │    User      │
└──────────────┘   │   └──────────────┘       └──────────────┘
       │           │          │                      │
       │           │          │                      │
       ▼           │          ▼                      ▼
┌──────────────┐   │   ┌──────────────┐       ┌──────────────┐
│   Namespace  │   │   │ UsageRecord  │       │  OrgMember   │
└──────────────┘   │   └──────────────┘       └──────────────┘
                   │          │
                   │          ▼
                   │   ┌──────────────┐
                   └──▶│   Invoice    │
                       └──────────────┘

┌──────────────┐
│ AuditEvent   │ (standalone)
└──────────────┘
```

### Schema Definitions

#### 001_organizations.sql

```sql
CREATE TABLE organizations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(63) NOT NULL UNIQUE,
    settings JSONB DEFAULT '{}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE organization_members (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id),
    user_id UUID NOT NULL,
    role VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(organization_id, user_id)
);
```

#### 002_subscriptions.sql

```sql
CREATE TYPE plan_tier AS ENUM ('free', 'essentials', 'business', 'enterprise');

CREATE TABLE subscriptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id) UNIQUE,
    plan plan_tier NOT NULL DEFAULT 'free',
    status VARCHAR(50) NOT NULL DEFAULT 'active',
    actions_included BIGINT NOT NULL,
    active_storage_gb DECIMAL(10,2) NOT NULL,
    retained_storage_gb DECIMAL(10,2) NOT NULL,
    stripe_customer_id VARCHAR(255),
    stripe_subscription_id VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

#### 003_usage.sql

```sql
CREATE TABLE usage_records (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id),
    namespace_id VARCHAR(255) NOT NULL,
    period_start TIMESTAMPTZ NOT NULL,
    period_end TIMESTAMPTZ NOT NULL,
    action_count BIGINT NOT NULL DEFAULT 0,
    active_storage_gbh DECIMAL(20,6) NOT NULL DEFAULT 0,
    retained_storage_gbh DECIMAL(20,6) NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(organization_id, namespace_id, period_start)
);

CREATE INDEX idx_usage_org_period ON usage_records(organization_id, period_start);
```

#### 004_invoices.sql

```sql
CREATE TABLE invoices (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id),
    period_start TIMESTAMPTZ NOT NULL,
    period_end TIMESTAMPTZ NOT NULL,
    line_items JSONB NOT NULL,
    subtotal_cents BIGINT NOT NULL,
    tax_cents BIGINT NOT NULL DEFAULT 0,
    total_cents BIGINT NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'draft',
    stripe_invoice_id VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    paid_at TIMESTAMPTZ
);
```

#### 005_audit.sql

```sql
CREATE TABLE audit_events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL,
    actor_type VARCHAR(50) NOT NULL,
    actor_id VARCHAR(255) NOT NULL,
    actor_ip INET,
    action VARCHAR(100) NOT NULL,
    resource_type VARCHAR(100) NOT NULL,
    resource_id VARCHAR(255),
    request_id VARCHAR(255),
    metadata JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_audit_org_time ON audit_events(organization_id, created_at DESC);
CREATE INDEX idx_audit_actor ON audit_events(actor_id, created_at DESC);
```

#### 006_users.sql

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE user_account_roles (
    user_id UUID NOT NULL REFERENCES users(id),
    organization_id UUID NOT NULL REFERENCES organizations(id),
    role VARCHAR(50) NOT NULL,
    PRIMARY KEY (user_id, organization_id)
);

CREATE TABLE user_namespace_permissions (
    user_id UUID NOT NULL REFERENCES users(id),
    namespace_id VARCHAR(255) NOT NULL,
    permission VARCHAR(50) NOT NULL,
    PRIMARY KEY (user_id, namespace_id)
);
```

#### 007_service_accounts.sql

```sql
CREATE TABLE service_accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    account_role VARCHAR(50),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE service_account_namespace_permissions (
    service_account_id UUID NOT NULL REFERENCES service_accounts(id),
    namespace_id VARCHAR(255) NOT NULL,
    permission VARCHAR(50) NOT NULL,
    PRIMARY KEY (service_account_id, namespace_id)
);
```

#### 008_api_keys.sql

```sql
CREATE TABLE api_keys (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    owner_type VARCHAR(50) NOT NULL,
    owner_id UUID NOT NULL,
    key_hash VARCHAR(255) NOT NULL,
    key_prefix VARCHAR(10) NOT NULL,
    name VARCHAR(255),
    expires_at TIMESTAMPTZ,
    disabled BOOLEAN DEFAULT FALSE,
    last_used_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_api_keys_owner ON api_keys(owner_type, owner_id);
CREATE INDEX idx_api_keys_prefix ON api_keys(key_prefix);
```

#### 009_namespaces.sql

```sql
CREATE TABLE cloud_namespaces (
    id VARCHAR(255) PRIMARY KEY,
    organization_id UUID NOT NULL REFERENCES organizations(id),
    name VARCHAR(255) NOT NULL,
    region VARCHAR(50) NOT NULL,
    retention_days INT NOT NULL DEFAULT 7,
    deletion_protected BOOLEAN DEFAULT FALSE,
    tags JSONB DEFAULT '{}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(organization_id, name)
);

CREATE TABLE namespace_certificates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    namespace_id VARCHAR(255) NOT NULL REFERENCES cloud_namespaces(id),
    certificate_pem TEXT NOT NULL,
    fingerprint VARCHAR(255) NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE namespace_search_attributes (
    namespace_id VARCHAR(255) NOT NULL REFERENCES cloud_namespaces(id),
    name VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL,
    PRIMARY KEY (namespace_id, name)
);
```

#### 010_connectivity.sql

```sql
CREATE TABLE connectivity_rules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id),
    name VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL,
    config JSONB NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE namespace_connectivity_bindings (
    namespace_id VARCHAR(255) NOT NULL REFERENCES cloud_namespaces(id),
    connectivity_rule_id UUID NOT NULL REFERENCES connectivity_rules(id),
    PRIMARY KEY (namespace_id, connectivity_rule_id)
);
```

#### 011_nexus.sql

```sql
CREATE TABLE nexus_endpoints (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id),
    name VARCHAR(255) NOT NULL,
    target_namespace_id VARCHAR(255) NOT NULL,
    handler_name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(organization_id, name)
);

CREATE TABLE nexus_endpoint_allowlist (
    endpoint_id UUID NOT NULL REFERENCES nexus_endpoints(id),
    caller_namespace_id VARCHAR(255) NOT NULL,
    PRIMARY KEY (endpoint_id, caller_namespace_id)
);
```

#### 012_export.sql

```sql
CREATE TABLE export_sinks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    namespace_id VARCHAR(255) NOT NULL REFERENCES cloud_namespaces(id),
    sink_type VARCHAR(50) NOT NULL,
    config JSONB NOT NULL,
    enabled BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

#### 013_credits.sql

```sql
CREATE TABLE credit_purchases (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id),
    amount_cents BIGINT NOT NULL,
    purchased_at TIMESTAMPTZ NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE credit_balance (
    organization_id UUID PRIMARY KEY REFERENCES organizations(id),
    balance_cents BIGINT NOT NULL DEFAULT 0,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE credit_transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id),
    amount_cents BIGINT NOT NULL,
    description VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

#### 014_user_groups.sql

```sql
CREATE TABLE user_groups (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(organization_id, name)
);

CREATE TABLE user_group_members (
    group_id UUID NOT NULL REFERENCES user_groups(id),
    user_id UUID NOT NULL REFERENCES users(id),
    PRIMARY KEY (group_id, user_id)
);

CREATE TABLE user_group_namespace_permissions (
    group_id UUID NOT NULL REFERENCES user_groups(id),
    namespace_id VARCHAR(255) NOT NULL,
    permission VARCHAR(50) NOT NULL,
    PRIMARY KEY (group_id, namespace_id)
);
```

<!-- Source: database-operations.md -->

## Database Operations

### Migration Strategy

#### Tool: golang-migrate

```bash
# Install
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Create new migration
migrate create -ext sql -dir schema -seq add_user_groups

# Run migrations
migrate -path schema -database "$DATABASE_URL" up

# Rollback last migration
migrate -path schema -database "$DATABASE_URL" down 1
```

#### Migration Naming

`{version}_{description}.up.sql` and `{version}_{description}.down.sql`

Example:

- `000001_create_organizations.up.sql`
- `000001_create_organizations.down.sql`

#### Migration Rules

1. **Always reversible**: Every `up` must have a `down`
2. **Backwards compatible**: Old code must work with new schema
3. **No data loss**: Never drop columns without migration period
4. **Atomic**: Each migration is a single transaction

#### Deployment Process

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Deploy    │────▶│   Migrate   │────▶│   Deploy    │
│   DB First  │     │   (auto)    │     │   App       │
└─────────────┘     └─────────────┘     └─────────────┘
```

1. Migration runs in CI/CD pipeline
2. App deployment waits for migration success
3. App must handle both old and new schema during rollout

#### Breaking Changes

When schema change is not backwards compatible:

**Phase 1: Add**

```sql
-- Add new column (nullable)
ALTER TABLE users ADD COLUMN new_email VARCHAR(255);
```

**Phase 2: Migrate**

```sql
-- Backfill data
UPDATE users SET new_email = email WHERE new_email IS NULL;
```

**Phase 3: Switch**

- Deploy new app code using new column
- Old code still works with old column

**Phase 4: Remove**

```sql
-- Drop old column (after all apps updated)
ALTER TABLE users DROP COLUMN email;
ALTER TABLE users RENAME COLUMN new_email TO email;
```

### Backup & Recovery

#### Automated Backups (RDS)

| Type         | Frequency  | Retention |
| ------------ | ---------- | --------- |
| Snapshot     | Daily      | 35 days   |
| WAL          | Continuous | 7 days    |
| Cross-region | Daily      | 7 days    |

#### Manual Backup

```bash
# Create snapshot
aws rds create-db-snapshot \
  --db-instance-identifier temporal-cloud-prod \
  --db-snapshot-identifier manual-backup-$(date +%Y%m%d)
```

#### Point-in-Time Recovery

```bash
# Restore to specific time
aws rds restore-db-instance-to-point-in-time \
  --source-db-instance-identifier temporal-cloud-prod \
  --target-db-instance-identifier temporal-cloud-pitr \
  --restore-time 2025-01-15T10:00:00Z
```

#### Disaster Recovery

See `dr.md` for full DR procedures.

### Performance Tuning

#### Connection Pooling

Use PgBouncer for connection pooling:

```yaml
# pgbouncer.ini
[databases]
temporal_cloud = host=rds-endpoint port=5432 dbname=temporal_cloud

[pgbouncer]
pool_mode = transaction
max_client_conn = 1000
default_pool_size = 50
```

#### Index Management

```sql
-- Find unused indexes
SELECT
  schemaname || '.' || relname AS table,
  indexrelname AS index,
  pg_size_pretty(pg_relation_size(i.indexrelid)) AS index_size,
  idx_scan as index_scans
FROM pg_stat_user_indexes ui
JOIN pg_index i ON ui.indexrelid = i.indexrelid
WHERE NOT indisunique AND idx_scan < 50
ORDER BY pg_relation_size(i.indexrelid) DESC;

-- Find missing indexes
SELECT
  schemaname || '.' || relname AS table,
  seq_scan,
  seq_tup_read,
  idx_scan,
  n_live_tup
FROM pg_stat_user_tables
WHERE seq_scan > 1000 AND n_live_tup > 10000
ORDER BY seq_tup_read DESC;
```

#### Query Analysis

```sql
-- Enable pg_stat_statements
CREATE EXTENSION IF NOT EXISTS pg_stat_statements;

-- Find slow queries
SELECT
  query,
  calls,
  total_time / 1000 as total_seconds,
  mean_time as avg_ms,
  rows
FROM pg_stat_statements
ORDER BY total_time DESC
LIMIT 20;
```

### Maintenance Tasks

#### Vacuum

Automatic vacuum is enabled, but monitor:

```sql
-- Check vacuum stats
SELECT
  schemaname,
  relname,
  last_vacuum,
  last_autovacuum,
  n_dead_tup
FROM pg_stat_user_tables
ORDER BY n_dead_tup DESC;
```

#### Analyze

Run after bulk operations:

```sql
ANALYZE usage_records;
```

#### Reindex

Schedule during low traffic:

```sql
REINDEX INDEX CONCURRENTLY idx_usage_org_period;
```

### Monitoring

#### Key Metrics

| Metric          | Warning   | Critical  |
| --------------- | --------- | --------- |
| CPU             | > 70%     | > 90%     |
| Connections     | > 80% max | > 95% max |
| Disk            | > 80%     | > 90%     |
| Replication Lag | > 30s     | > 60s     |
| Query Time P99  | > 500ms   | > 1s      |

#### Alerts

```yaml
alerts:
  - name: DatabaseHighCPU
    expr: aws_rds_cpuutilization_average > 70
    for: 10m
    severity: warning

  - name: DatabaseConnectionsHigh
    expr: aws_rds_database_connections_average / aws_rds_database_connections_maximum > 0.8
    for: 5m
    severity: warning
```

### Access Control

#### Roles

| Role       | Access                   | Used By               |
| ---------- | ------------------------ | --------------------- |
| app_rw     | Read/write to app tables | Application           |
| app_ro     | Read-only                | Reporting, Analytics  |
| migrations | Schema changes           | CI/CD pipeline        |
| admin      | Full access              | DBAs (emergency only) |

```sql
-- Create application role
CREATE ROLE app_rw;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO app_rw;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO app_rw;
```

#### Auditing

Enable RDS audit logging:

```
rds.force_ssl = 1
log_connections = on
log_disconnections = on
log_statement = 'ddl'
```

<!-- Source: dependency-management.md -->

## Dependency Management

### Dependency Types

| Type              | Examples              | Update Strategy                   |
| ----------------- | --------------------- | --------------------------------- |
| Upstream Temporal | temporal server, SDKs | Daily sync, careful review        |
| Infrastructure    | Terraform, Helm       | Monthly, test thoroughly          |
| Go modules        | grpc, stripe-go       | Weekly automated, review breaking |
| NPM packages      | React, Next.js        | Weekly automated, review major    |
| Security          | Any CVE               | Immediate                         |

### Go Dependencies

#### go.mod Management

```bash
# Update all dependencies
go get -u ./...
go mod tidy

# Update specific dependency
go get -u github.com/stripe/stripe-go/v76

# Check for outdated
go list -u -m all
```

#### Dependabot Configuration

```yaml
# .github/dependabot.yml
version: 2
updates:
  - package-ecosystem: "gomod"
    directory: "/"
    schedule:
      interval: "weekly"
    commit-message:
      prefix: "deps(go):"
    labels:
      - "dependencies"
      - "go"
    groups:
      temporal:
        patterns:
          - "go.temporal.io/*"
      grpc:
        patterns:
          - "google.golang.org/grpc*"
          - "google.golang.org/protobuf*"
```

#### Version Pinning

```go
// go.mod
module github.com/YOUR_ORG/temporal-cloud-platform

go 1.22

require (
    go.temporal.io/api v1.29.0
    go.temporal.io/sdk v1.26.0
    github.com/stripe/stripe-go/v76 v76.0.0
)
```

### NPM Dependencies

#### package.json Management

```bash
# Check outdated
pnpm outdated

# Update all (within semver)
pnpm update

# Update specific package
pnpm update react@latest
```

#### Renovate Configuration

```json
{
  "extends": ["config:base"],
  "packageRules": [
    {
      "matchPackagePatterns": ["^@types/"],
      "groupName": "type definitions"
    },
    {
      "matchPackagePatterns": ["react", "next"],
      "groupName": "react ecosystem"
    }
  ],
  "schedule": ["after 10pm on sunday"]
}
```

### Temporal SDK Compatibility

#### Version Matrix

Maintain compatibility between components:

```yaml
# version-matrix.yaml
cloud_platform: v1.2.0
temporal_server: v1.24.2
sdk_go: v1.29.0
sdk_java: v1.25.0
sdk_typescript: v1.10.0
sdk_python: v1.7.0
```

#### Upgrade Process

1. **Test**: Run SDK compatibility tests
2. **Stage**: Deploy to staging
3. **Soak**: Run for 24 hours
4. **Prod**: Gradual rollout
5. **Document**: Update version matrix

### Security Vulnerabilities

#### Scanning

```yaml
# .github/workflows/security-scan.yaml
name: Security Scan
on:
  push:
  schedule:
    - cron: "0 0 * * *"

jobs:
  scan:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Run Snyk
        uses: snyk/actions/golang@master
        with:
          command: test
          args: --severity-threshold=high
```

#### Vulnerability Response

| Severity | Response Time | Action                         |
| -------- | ------------- | ------------------------------ |
| Critical | 24 hours      | Immediate patch, hotfix deploy |
| High     | 7 days        | Next release                   |
| Medium   | 30 days       | Regular release                |
| Low      | 90 days       | When convenient                |

#### Patching Process

1. **Alert**: Snyk/Dependabot creates issue
2. **Assess**: Determine exploitability
3. **Patch**: Update dependency
4. **Test**: Run full test suite
5. **Deploy**: Based on severity
6. **Verify**: Confirm scan passes

### Vendoring

We **do not** vendor Go dependencies.

**Rationale**:

- `go mod` provides reproducibility
- Vendoring bloats repository
- Harder to audit for security

**Exception**: Fork and vendor if:

- Upstream is unmaintained
- Critical patch needed immediately

### License Compliance

#### Allowed Licenses

| License        | Allowed   |
| -------------- | --------- |
| MIT            | ✅        |
| Apache 2.0     | ✅        |
| BSD 2/3-Clause | ✅        |
| ISC            | ✅        |
| MPL 2.0        | ⚠️ Review |
| GPL            | ❌        |
| AGPL           | ❌        |

#### License Scanning

```bash
# Go
go-licenses check ./...

# NPM
npx license-checker --onlyAllow "MIT;Apache-2.0;BSD-2-Clause;BSD-3-Clause;ISC"
```

### Upgrade Checklist

Before upgrading major dependencies:

- [ ] Read changelog for breaking changes
- [ ] Check our usage of deprecated APIs
- [ ] Run full test suite
- [ ] Deploy to staging
- [ ] Monitor for 24 hours
- [ ] Deploy to production
- [ ] Update documentation

### Rollback

If dependency update causes issues:

```bash
# Go: Revert go.mod and go.sum
git checkout HEAD~1 -- go.mod go.sum
go mod download

# NPM: Revert package.json and lockfile
git checkout HEAD~1 -- package.json pnpm-lock.yaml
pnpm install
```

<!-- Source: dr.md -->

## Disaster Recovery

### Recovery Objectives

| Metric                         | Target    |
| ------------------------------ | --------- |
| RPO (Recovery Point Objective) | 5 minutes |
| RTO (Recovery Time Objective)  | 1 hour    |

### Backup Strategy

#### Database Backups

| Data                | Frequency  | Retention | Location        |
| ------------------- | ---------- | --------- | --------------- |
| PostgreSQL WAL      | Continuous | 35 days   | Cross-region S3 |
| PostgreSQL Snapshot | Daily      | 90 days   | Cross-region S3 |
| Redis Snapshot      | Hourly     | 7 days    | Same region S3  |

#### Configuration Backups

| Data                 | Frequency | Retention | Location        |
| -------------------- | --------- | --------- | --------------- |
| Terraform state      | On change | Forever   | S3 + versioning |
| Kubernetes manifests | On change | Forever   | Git             |
| Secrets              | On change | 90 days   | AWS Backup      |

#### Audit Log Archive

| Data         | Frequency | Retention | Location                |
| ------------ | --------- | --------- | ----------------------- |
| Audit events | Real-time | 7 years   | Cross-region S3 Glacier |

### Recovery Procedures

#### Scenario 1: Data Corruption

1. Identify corruption scope and time
2. Stop writes to affected tables
3. Restore from point-in-time backup
4. Validate data integrity
5. Resume operations
6. Post-incident review

**RTO**: 30 minutes  
**RPO**: 5 minutes

#### Scenario 2: Single Region Failure

1. Confirm region is unavailable
2. Trigger DNS failover to secondary
3. Verify secondary region health
4. Notify customers of degraded service
5. Begin data sync when primary recovers
6. Failback when primary is stable

**RTO**: 15 minutes  
**RPO**: 5 minutes

#### Scenario 3: Complete Platform Failure

1. Activate incident response team
2. Provision new infrastructure from Terraform
3. Restore database from latest backup
4. Deploy applications from CI/CD
5. Validate all services
6. Update DNS to new infrastructure
7. Comprehensive post-mortem

**RTO**: 4 hours  
**RPO**: 5 minutes

### DR Testing

#### Test Schedule

| Test Type          | Frequency | Duration |
| ------------------ | --------- | -------- |
| Backup restoration | Monthly   | 2 hours  |
| Failover drill     | Quarterly | 4 hours  |
| Full DR exercise   | Annually  | 1 day    |

#### Test Checklist

- [ ] Verify backup integrity
- [ ] Test restore procedure
- [ ] Validate data consistency
- [ ] Test failover automation
- [ ] Measure actual RTO/RPO
- [ ] Document lessons learned

### Runbook Location

All detailed runbooks are in:
`temporal-cloud-infra/docs/runbooks/`

- `disaster-recovery.md`
- `database-restore.md`
- `region-failover.md`
- `incident-response.md`

<!-- Source: environment-config.md -->

## Environment Configuration

### Environment Hierarchy

```
┌─────────────────────────────────────────────────────────────────┐
│                         Production                               │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐              │
│  │  us-east-1  │  │  eu-west-1  │  │ ap-south-1  │              │
│  └─────────────┘  └─────────────┘  └─────────────┘              │
└─────────────────────────────────────────────────────────────────┘
                              ▲
┌─────────────────────────────────────────────────────────────────┐
│                         Staging                                  │
│  ┌─────────────┐                                                │
│  │  us-east-1  │  (Production mirror, sanitized data)          │
│  └─────────────┘                                                │
└─────────────────────────────────────────────────────────────────┘
                              ▲
┌─────────────────────────────────────────────────────────────────┐
│                           QA                                     │
│  ┌─────────────┐                                                │
│  │  us-east-1  │  (Testing, seed data)                         │
│  └─────────────┘                                                │
└─────────────────────────────────────────────────────────────────┘
                              ▲
┌─────────────────────────────────────────────────────────────────┐
│                      Development                                 │
│  ┌─────────────┐                                                │
│  │   Local     │  (Docker Compose)                              │
│  └─────────────┘                                                │
└─────────────────────────────────────────────────────────────────┘
```

### Environment Definitions

| Environment | Purpose             | Data           | Access      | Deploy              |
| ----------- | ------------------- | -------------- | ----------- | ------------------- |
| Local       | Developer machine   | Seed           | Developer   | Manual              |
| Dev         | Shared development  | Seed           | Engineering | On push             |
| QA          | QA testing          | Sanitized      | QA + Eng    | On merge to develop |
| Staging     | Pre-prod validation | Sanitized prod | All         | On merge to main    |
| Production  | Live service        | Real           | Restricted  | Manual approval     |

### Configuration Management

#### Configuration Sources (Priority Order)

1. **Environment Variables** (highest priority)
2. **Config Files** (per-environment)
3. **Defaults** (in code)

#### Configuration Structure

```yaml
# config/base.yaml (defaults)
server:
  port: 8080
  read_timeout: 30s
  write_timeout: 30s

database:
  max_connections: 100
  idle_connections: 10

logging:
  level: info
  format: json

features:
  enable_scim: false
  enable_multi_region: false
```

```yaml
# config/production.yaml (overrides)
server:
  port: 8080

database:
  max_connections: 500
  idle_connections: 50

logging:
  level: warn

features:
  enable_scim: true
  enable_multi_region: true
```

#### Environment Variables

```bash
# Naming convention: TEMPORAL_CLOUD_{SECTION}_{KEY}
TEMPORAL_CLOUD_SERVER_PORT=8080
TEMPORAL_CLOUD_DATABASE_HOST=prod-db.internal
TEMPORAL_CLOUD_DATABASE_MAX_CONNECTIONS=500
TEMPORAL_CLOUD_LOGGING_LEVEL=warn
```

#### Go Configuration Loading

```go
type Config struct {
    Server   ServerConfig   `yaml:"server"`
    Database DatabaseConfig `yaml:"database"`
    Logging  LoggingConfig  `yaml:"logging"`
    Features FeatureConfig  `yaml:"features"`
}

func LoadConfig(env string) (*Config, error) {
    // Load base config
    cfg := loadYAML("config/base.yaml")

    // Overlay environment-specific config
    envConfig := loadYAML(fmt.Sprintf("config/%s.yaml", env))
    mergeConfig(cfg, envConfig)

    // Override with environment variables
    applyEnvVars(cfg)

    // Validate
    if err := cfg.Validate(); err != nil {
        return nil, err
    }

    return cfg, nil
}
```

### Environment-Specific Settings

#### Database

| Setting         | Local     | Dev    | QA    | Staging    | Prod    |
| --------------- | --------- | ------ | ----- | ---------- | ------- |
| Host            | localhost | dev-db | qa-db | staging-db | prod-db |
| Max Connections | 10        | 50     | 100   | 200        | 500     |
| SSL             | false     | true   | true  | true       | true    |
| Read Replicas   | 0         | 0      | 0     | 1          | 3       |

#### Caching

| Setting      | Local     | Dev       | QA       | Staging       | Prod       |
| ------------ | --------- | --------- | -------- | ------------- | ---------- |
| Redis Host   | localhost | dev-redis | qa-redis | staging-redis | prod-redis |
| TTL          | 60s       | 60s       | 300s     | 300s          | 300s       |
| Cluster Mode | false     | false     | false    | true          | true       |

#### External Services

| Service   | Local     | Dev       | QA        | Staging   | Prod      |
| --------- | --------- | --------- | --------- | --------- | --------- |
| Stripe    | Test mode | Test mode | Test mode | Test mode | Live mode |
| SendGrid  | Sandbox   | Sandbox   | Sandbox   | Live      | Live      |
| PagerDuty | Disabled  | Disabled  | Disabled  | Test      | Live      |

#### Feature Flags

| Feature        | Local | Dev  | QA   | Staging | Prod     |
| -------------- | ----- | ---- | ---- | ------- | -------- |
| SCIM           | true  | true | true | true    | Per plan |
| Multi-region   | true  | true | true | true    | Per plan |
| New Billing UI | true  | true | true | 50%     | 10%      |

### Kubernetes ConfigMaps

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: cloud-platform-config
  namespace: cloud-platform
data:
  TEMPORAL_CLOUD_ENV: production
  TEMPORAL_CLOUD_REGION: us-east-1
  TEMPORAL_CLOUD_LOG_LEVEL: warn
  TEMPORAL_CLOUD_METRICS_ENABLED: "true"
```

#### Per-Environment Kustomization

```yaml
# kustomization.yaml (base)
resources:
  - deployment.yaml
  - service.yaml
  - configmap.yaml

# overlays/production/kustomization.yaml
resources:
  - ../../base
patchesStrategicMerge:
  - deployment-patch.yaml
configMapGenerator:
  - name: cloud-platform-config
    behavior: merge
    literals:
      - TEMPORAL_CLOUD_ENV=production
      - TEMPORAL_CLOUD_REPLICAS=3
```

### Environment Promotion

#### Promotion Flow

```
feature/* → develop → staging → production
    │          │          │          │
    ▼          ▼          ▼          ▼
  Local       Dev        QA      Staging → Prod
```

#### Promotion Checklist

**Dev → QA**

- [ ] Unit tests pass
- [ ] Integration tests pass
- [ ] Code review approved

**QA → Staging**

- [ ] QA sign-off
- [ ] No P0/P1 bugs
- [ ] Performance benchmarks pass

**Staging → Production**

- [ ] 24h soak in staging
- [ ] Security scan clean
- [ ] Rollback plan confirmed
- [ ] On-call notified
- [ ] 2 approvals

### Environment Isolation

#### Network Isolation

```hcl
# Each environment in separate VPC
module "vpc_prod" {
  source = "./modules/vpc"
  cidr   = "10.0.0.0/16"
  env    = "production"
}

module "vpc_staging" {
  source = "./modules/vpc"
  cidr   = "10.1.0.0/16"
  env    = "staging"
}

# No VPC peering between prod and non-prod
```

#### Data Isolation

- Production data NEVER in non-prod environments
- Staging uses sanitized copy (PII removed)
- QA/Dev use synthetic seed data

#### Access Isolation

| Environment | Access                                           |
| ----------- | ------------------------------------------------ |
| Production  | SRE, On-call (emergency), Senior Eng (read-only) |
| Staging     | All Engineering                                  |
| QA          | All Engineering, QA                              |
| Dev         | All Engineering                                  |

### Environment Debugging

#### Local Development

```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f api

# Access database
psql postgres://temporal:temporal@localhost:5432/temporal_cloud
```

#### Remote Environments

```bash
# Port forward to pod
kubectl port-forward svc/cloud-api 8080:8080 -n cloud-platform

# View logs
kubectl logs -f deployment/cloud-api -n cloud-platform

# Exec into pod
kubectl exec -it deployment/cloud-api -n cloud-platform -- /bin/sh
```

### Environment Health Checks

#### Health Endpoints

| Endpoint   | Purpose                  |
| ---------- | ------------------------ |
| `/health`  | Basic liveness           |
| `/ready`   | Readiness (dependencies) |
| `/metrics` | Prometheus metrics       |

#### Environment Dashboard

Each environment has a health dashboard showing:

- Service status (up/down)
- Error rates
- Latency percentiles
- Database connections
- Cache hit rate
- Queue depth

<!-- Source: feature-flags.md -->

## Feature Flags

### Overview

Feature flags enable safe rollout of new features, A/B testing, and quick rollback without deployment.

### Flag Types

| Type        | Purpose             | Example                          |
| ----------- | ------------------- | -------------------------------- |
| Release     | Gate new features   | `enable_multi_region_namespaces` |
| Ops         | Operational toggles | `enable_rate_limiting`           |
| Experiment  | A/B testing         | `new_billing_ui_variant`         |
| Kill Switch | Emergency disable   | `disable_namespace_creation`     |

### Flag Definition

```yaml
# flags/enable_scim.yaml
name: enable_scim
description: Enable SCIM provisioning for enterprise customers
type: release
owner: platform-team
default: false
targeting:
  - match:
      plan: enterprise
    value: true
  - match:
      org_id:
        - org-beta-1
        - org-beta-2
    value: true
```

### Targeting Rules

#### By Plan

```yaml
targeting:
  - match:
      plan: [enterprise, business]
    value: true
```

#### By Organization

```yaml
targeting:
  - match:
      org_id: org-123
    value: true
```

#### By Percentage

```yaml
targeting:
  - match:
      percentage: 10 # 10% of all users
    value: true
```

#### By User Email Domain

```yaml
targeting:
  - match:
      email_domain: temporal.io
    value: true
```

### Backend Usage

#### Go SDK

```go
import "github.com/YOUR_ORG/temporal-cloud-platform/pkg/flags"

func CreateNamespace(ctx context.Context, req *CreateNamespaceRequest) error {
    // Check feature flag
    if flags.IsEnabled(ctx, "enable_multi_region") {
        // New multi-region logic
        return createMultiRegionNamespace(ctx, req)
    }

    // Existing single-region logic
    return createSingleRegionNamespace(ctx, req)
}
```

#### Flag Client

```go
type FlagClient interface {
    IsEnabled(ctx context.Context, flagName string) bool
    GetVariant(ctx context.Context, flagName string) string
    GetValue(ctx context.Context, flagName string, defaultValue interface{}) interface{}
}

// Implementation with caching
type cachedFlagClient struct {
    store  FlagStore
    cache  *lru.Cache
    ttl    time.Duration
}

func (c *cachedFlagClient) IsEnabled(ctx context.Context, flagName string) bool {
    // Extract targeting context
    org := GetOrgFromContext(ctx)
    user := GetUserFromContext(ctx)

    // Check cache
    cacheKey := fmt.Sprintf("%s:%s:%s", flagName, org.ID, user.ID)
    if cached, ok := c.cache.Get(cacheKey); ok {
        return cached.(bool)
    }

    // Evaluate flag
    result := c.evaluate(flagName, EvalContext{
        OrgID:   org.ID,
        Plan:    org.Plan,
        UserID:  user.ID,
        Email:   user.Email,
    })

    c.cache.Add(cacheKey, result)
    return result
}
```

### Frontend Usage

#### React Hook

```typescript
import { useFeatureFlag } from "@/hooks/useFeatureFlags";

function BillingPage() {
  const showNewBillingUI = useFeatureFlag("new_billing_ui");

  if (showNewBillingUI) {
    return <NewBillingDashboard />;
  }

  return <LegacyBillingDashboard />;
}
```

#### Flag Provider

```typescript
// Fetch flags on app load
export function FlagProvider({ children }) {
  const { data: flags } = useQuery({
    queryKey: ["feature-flags"],
    queryFn: () => api.getFlags(),
    staleTime: 5 * 60 * 1000, // 5 minutes
  });

  return <FlagContext.Provider value={flags}>{children}</FlagContext.Provider>;
}
```

### Flag Lifecycle

#### 1. Create Flag

```bash
# Add flag definition
cat > flags/enable_new_feature.yaml << EOF
name: enable_new_feature
description: New feature description
type: release
owner: @engineer
default: false
EOF

# Deploy flag
make deploy-flags
```

#### 2. Implement Behind Flag

```go
if flags.IsEnabled(ctx, "enable_new_feature") {
    newFeatureLogic()
} else {
    existingLogic()
}
```

#### 3. Gradual Rollout

```yaml
# Week 1: Internal
targeting:
  - match: { email_domain: temporal.io }
    value: true

# Week 2: Beta customers
targeting:
  - match: { email_domain: temporal.io }
    value: true
  - match: { org_id: [beta-org-1, beta-org-2] }
    value: true

# Week 3: 10% of all
targeting:
  - match: { percentage: 10 }
    value: true

# Week 4: 100%
default: true
```

#### 4. Remove Flag

After 100% rollout and monitoring:

1. Remove flag checks from code
2. Delete flag definition
3. Mark as deprecated (7 days)
4. Delete flag

### Operations

#### Emergency Kill Switch

```bash
# Disable feature immediately
curl -X PUT https://config.internal/flags/enable_new_billing \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -d '{"default": false, "targeting": []}'
```

#### View Flag Status

```bash
# List all flags
curl https://config.internal/flags

# Get specific flag
curl https://config.internal/flags/enable_scim
```

#### Audit Trail

All flag changes are logged:

```json
{
  "timestamp": "2025-01-01T12:00:00Z",
  "action": "update",
  "flag": "enable_scim",
  "actor": "admin@temporal.io",
  "before": { "default": false },
  "after": { "default": true }
}
```

### Best Practices

1. **Short-lived flags**: Remove within 30 days of 100% rollout
2. **Clear naming**: `enable_`, `disable_`, `show_`, `use_`
3. **Document**: Always include description and owner
4. **Test both paths**: Unit tests for flag on and off
5. **Monitor**: Add metrics for flag evaluation

### Metrics

```go
// Track flag usage
flagEvaluations.WithLabelValues(flagName, result).Inc()
```

Dashboard:

- Flag evaluation count by flag
- Flag distribution (true vs false)
- Flags enabled over time

<!-- Source: git.md -->

## Git Management

### Repository Strategy

#### Forked Repositories

| Repo                             | Upstream                                    | Purpose                          |
| -------------------------------- | ------------------------------------------- | -------------------------------- |
| temporal                         | temporalio/temporal                         | Core server + cloud interceptors |
| cloud-api                        | temporalio/cloud-api                        | Cloud API protos                 |
| tcld                             | temporalio/tcld                             | Cloud CLI                        |
| terraform-provider-temporalcloud | temporalio/terraform-provider-temporalcloud | Terraform                        |

#### New Repositories

| Repo                    | Purpose                |
| ----------------------- | ---------------------- |
| temporal-cloud-platform | Backend services       |
| temporal-cloud-console  | Web UI                 |
| temporal-cloud-infra    | Infrastructure as Code |

### Branch Strategy

```
main                    ← Mirror of upstream (auto-sync)
├── cloud/main          ← Production releases
├── cloud/staging       ← Staging environment
├── cloud/develop       ← Development integration
└── feature/*           ← Feature branches
```

### Sync Policy

#### Golden Rule

**NEVER modify existing upstream files. Only ADD new files/packages.**

#### Allowed Changes (temporal fork)

```
✅ ADD: common/cloud/           (new package)
✅ ADD: common/cloud/metering/
✅ ADD: common/cloud/quota/
✅ ADD: common/cloud/audit/
❌ MODIFY: common/config/config.go
❌ MODIFY: service/frontend/fx.go
```

#### Protected Paths

```yaml
# .github/protected-paths.yaml
temporal:
  protected:
    - "api/"
    - "cmd/server/"
    - "common/authorization/authorizer.go"
    - "common/config/config.go"
    - "service/frontend/workflow_handler.go"
    - "service/history/"
    - "service/matching/"
  allowed_additions:
    - "common/cloud/"
```

### Sync Automation

#### Daily Sync Workflow

```yaml
name: Sync Upstream
on:
  schedule:
    - cron: "0 0 * * *"
  workflow_dispatch:

jobs:
  sync:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
        with:
          ref: cloud/main
          fetch-depth: 0
          token: ${{ secrets.GH_PAT }}

      - name: Setup Git
        run: |
          git config user.name "Sync Bot"
          git config user.email "bot@your-org.com"

      - name: Sync
        run: |
          git remote add upstream https://github.com/temporalio/temporal.git || true
          git fetch upstream
          git merge upstream/main --no-edit || exit 1
          git push origin cloud/main
```

### Pre-Commit Hooks

#### Installation

```bash
# Install pre-commit
pip install pre-commit
pre-commit install
```

#### Configuration

```yaml
# .pre-commit-config.yaml
repos:
  - repo: local
    hooks:
      - id: check-protected-paths
        name: Check Protected Paths
        entry: scripts/check-protected-paths.sh
        language: script
        pass_filenames: true

      - id: go-fmt
        name: Go Format
        entry: gofmt -w
        language: system
        files: \.go$

      - id: go-lint
        name: Go Lint
        entry: golangci-lint run
        language: system
        files: \.go$
```

#### Protected Path Check Script

```bash
#!/bin/bash
# scripts/check-protected-paths.sh

PROTECTED=(
  "common/config/config.go"
  "common/authorization/authorizer.go"
  "service/frontend/fx.go"
  "service/history/"
  "service/matching/"
)

for file in "$@"; do
  for p in "${PROTECTED[@]}"; do
    if [[ "$file" == "$p"* ]]; then
      echo "❌ BLOCKED: Cannot modify protected path: $file"
      echo "   Add cloud code to common/cloud/ instead"
      exit 1
    fi
  done
done
```

### Branch Protection Rules

#### cloud/main

- Require pull request reviews: 2
- Require status checks: lint, test, build
- Require branches to be up to date
- No force pushes
- No deletions

#### cloud/develop

- Require pull request reviews: 1
- Require status checks: lint, test
- Allow force pushes (with lease)

### Commit Conventions

#### Format

```
<type>(<scope>): <subject>

<body>

<footer>
```

#### Types

- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation
- `style`: Formatting
- `refactor`: Code restructuring
- `test`: Tests
- `chore`: Maintenance

#### Example

```
feat(billing): add invoice generation workflow

Implements the monthly invoice generation workflow that:
- Calculates usage from metering data
- Generates line items for actions and storage
- Creates Stripe invoice

Closes #123
```

<!-- Source: ha.md -->

## High Availability

### Redundancy Design

#### Component Redundancy

| Component          | Replicas | Spread         | Failover  |
| ------------------ | -------- | -------------- | --------- |
| Temporal Frontend  | 3        | Multi-AZ       | Automatic |
| Temporal History   | 3        | Multi-AZ       | Automatic |
| Temporal Matching  | 3        | Multi-AZ       | Automatic |
| Cloud Platform API | 3        | Multi-AZ       | Automatic |
| Cloud Console      | 2        | Multi-AZ       | Automatic |
| PostgreSQL         | 2        | Multi-AZ (RDS) | Automatic |
| Redis              | 3        | Cluster mode   | Automatic |

#### Pod Anti-Affinity

```yaml
affinity:
  podAntiAffinity:
    requiredDuringSchedulingIgnoredDuringExecution:
      - labelSelector:
          matchLabels:
            app: temporal-frontend
        topologyKey: topology.kubernetes.io/zone
```

### Failure Scenarios

| Scenario       | Detection        | Recovery            | RTO   |
| -------------- | ---------------- | ------------------- | ----- |
| Pod failure    | K8s health check | Auto-restart        | < 30s |
| Node failure   | K8s node monitor | Reschedule pods     | < 2m  |
| AZ failure     | ALB health check | Route to healthy AZ | < 1m  |
| Region failure | Route53 health   | DNS failover        | < 5m  |
| DB failure     | RDS monitoring   | Auto-failover       | < 2m  |
| Redis failure  | ElastiCache      | Replica promotion   | < 1m  |

### Health Checks

#### Kubernetes Probes

```yaml
livenessProbe:
  httpGet:
    path: /health
    port: 8080
  initialDelaySeconds: 30
  periodSeconds: 10
  failureThreshold: 3

readinessProbe:
  httpGet:
    path: /ready
    port: 8080
  initialDelaySeconds: 5
  periodSeconds: 5
  failureThreshold: 3
```

#### ALB Health Checks

- Path: `/health`
- Interval: 10 seconds
- Timeout: 5 seconds
- Healthy threshold: 2
- Unhealthy threshold: 3

### Load Balancing

#### External Traffic

- Route53 latency-based routing
- ALB with cross-zone load balancing
- Connection draining: 30 seconds

#### Internal Traffic

- Kubernetes Service (ClusterIP)
- gRPC load balancing via headless service

### Namespace High Availability

#### Same-Region Replication

- Data replicated within region
- Automatic failover between AZs
- 99.99% SLA

#### Multi-Region Replication

- Async replication to secondary region
- Manual or automatic failover
- RPO: < 5 minutes
- RTO: < 15 minutes

#### Failover Process

1. Detect primary region failure
2. Promote secondary to primary
3. Update DNS records
4. Notify affected customers
5. Begin data reconciliation

<!-- Source: incident-management.md -->

## Incident Management

### Incident Lifecycle

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Detect    │────▶│   Respond   │────▶│   Mitigate  │
│             │     │             │     │             │
└─────────────┘     └─────────────┘     └─────────────┘
                                               │
                                               ▼
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Improve   │◀────│   Review    │◀────│   Resolve   │
│             │     │             │     │             │
└─────────────┘     └─────────────┘     └─────────────┘
```

### Severity Levels

| Level | Description         | Examples                           | Response Time |
| ----- | ------------------- | ---------------------------------- | ------------- |
| SEV1  | Complete outage     | All namespaces down, data loss     | 15 min        |
| SEV2  | Major degradation   | Region down, 50%+ error rate       | 30 min        |
| SEV3  | Partial degradation | Single feature broken, 10%+ errors | 2 hours       |
| SEV4  | Minor issue         | UI bug, minor latency increase     | 24 hours      |

### Detection

#### Automatic (PagerDuty)

```yaml
# Alert rules that trigger incidents
alerts:
  - name: HighErrorRate
    condition: error_rate > 5%
    duration: 5m
    severity: SEV2

  - name: ServiceDown
    condition: up == 0
    duration: 1m
    severity: SEV1

  - name: HighLatency
    condition: p99_latency > 1s
    duration: 10m
    severity: SEV3
```

#### Manual (Slack)

```
/incident create
Title: Users unable to create namespaces
Severity: SEV2
Description: Multiple customers reporting 500 errors
```

### Response

#### Incident Commander (IC)

The first responder becomes IC until handoff.

**IC Responsibilities**:

- Coordinate response
- Communicate status
- Make decisions
- Delegate tasks

#### Communication Channels

| Channel        | Purpose                        |
| -------------- | ------------------------------ |
| #incident-{id} | Real-time coordination         |
| #incidents     | All incident notifications     |
| Status Page    | External communication         |
| Email          | Customer notification (SEV1/2) |

#### Status Page Updates

```
[Investigating] We are investigating issues with namespace creation.

[Identified] We have identified the root cause and are implementing a fix.

[Monitoring] A fix has been deployed. We are monitoring the situation.

[Resolved] This incident has been resolved.
```

### Roles

#### During Incident

| Role                | Responsibility                 |
| ------------------- | ------------------------------ |
| Incident Commander  | Overall coordination           |
| Technical Lead      | Technical investigation        |
| Communications Lead | Status updates, customer comms |
| Scribe              | Document timeline, actions     |

#### Escalation

```
On-Call Engineer (0 min)
      ↓
Tech Lead (15 min for SEV1)
      ↓
Engineering Manager (30 min for SEV1)
      ↓
VP Engineering (60 min for SEV1)
      ↓
CTO (customer-impacting SEV1)
```

### Response Procedures

#### SEV1 Response

```
1. [0 min] Alert fires → On-call paged
2. [5 min] Acknowledge → Join #incident channel
3. [10 min] Assess impact → Confirm severity
4. [15 min] Assemble team → Page additional help if needed
5. [20 min] Initial status update → Post to status page
6. [30 min] Identify root cause OR escalate
7. [Ongoing] Updates every 15 minutes
```

#### Common Actions

**Rollback**:

```bash
# Helm rollback
helm rollback cloud-platform -n cloud-platform

# Verify
kubectl get pods -n cloud-platform
```

**Feature Flag Disable**:

```bash
# Disable problematic feature
curl -X PUT https://config.internal/flags/new-billing-flow \
  -d '{"enabled": false}'
```

**Scale Up**:

```bash
# Scale replicas
kubectl scale deployment cloud-api --replicas=10 -n cloud-platform
```

**Traffic Shift**:

```bash
# Shift traffic to healthy region
aws route53 change-resource-record-sets ...
```

### Communication Templates

#### Internal (Slack)

```
🚨 *INCIDENT STARTED* 🚨

*Incident ID*: INC-2025-001
*Severity*: SEV1
*Title*: API returning 500 errors
*IC*: @oncall-engineer
*Channel*: #incident-inc-2025-001

*Impact*: All API requests failing
*Customers Affected*: All

*Current Status*: Investigating
```

#### External (Customer Email)

```
Subject: [Temporal Cloud Incident] Service Degradation

Dear Customer,

We are currently experiencing service issues affecting
Temporal Cloud. Our team is actively investigating.

Impact: [Description]
Start Time: [Time UTC]

We will provide updates every 30 minutes.

Current Status: [Status]

For real-time updates, visit: status.temporal.io

Temporal Cloud Team
```

### Post-Incident

#### Timeline (Required for SEV1/2)

Within 24 hours:

- [ ] Detailed timeline documented
- [ ] Root cause identified
- [ ] Impact quantified

#### Post-Mortem (Required for SEV1/2)

Within 5 business days:

- [ ] Post-mortem document completed
- [ ] Review meeting held
- [ ] Action items assigned

#### Post-Mortem Template

```markdown
# Post-Mortem: INC-2025-001

## Summary

One-paragraph summary of what happened.

## Impact

- Duration: 45 minutes
- Customers affected: 1,234
- Namespaces affected: 5,678
- Revenue impact: $X,XXX
- SLA impact: Yes/No

## Timeline (all times UTC)

- 14:00 - Deployment started
- 14:05 - Error rate increased
- 14:10 - Alert fired, on-call paged
- 14:15 - Investigation started
- 14:25 - Root cause identified
- 14:30 - Rollback initiated
- 14:35 - Rollback complete
- 14:45 - Confirmed resolved

## Root Cause

Technical explanation of what went wrong.

## Contributing Factors

- Factor 1
- Factor 2

## What Went Well

- Fast detection (5 min)
- Clear runbook followed

## What Went Wrong

- Missing test case for edge condition
- Deployment during peak hours

## Action Items

| Action              | Owner | Due        |
| ------------------- | ----- | ---------- |
| Add regression test | @eng1 | 2025-02-01 |
| Update runbook      | @eng2 | 2025-02-01 |
| Add circuit breaker | @eng3 | 2025-02-15 |

## Lessons Learned

Key takeaways for the team.
```

### Metrics

| Metric                          | Target          |
| ------------------------------- | --------------- |
| MTTA (Mean Time To Acknowledge) | < 5 min         |
| MTTD (Mean Time To Detect)      | < 5 min         |
| MTTR (Mean Time To Resolve)     | < 1 hour (SEV1) |
| Post-mortem completion rate     | 100% (SEV1/2)   |
| Action item completion rate     | 100%            |

<!-- Source: infra.md -->

## Infrastructure

### Architecture Overview

#### Multi-Region Deployment

- 3 regions: us-east-1 (primary), eu-west-1, ap-south-1
- Multi-AZ within each region
- EKS for container orchestration
- RDS PostgreSQL (Multi-AZ)
- ElastiCache Redis (Cluster mode)

### Terraform Modules

| Module     | Path                 | Purpose                   |
| ---------- | -------------------- | ------------------------- |
| vpc        | `modules/vpc`        | Network, subnets, NAT     |
| eks        | `modules/eks`        | Kubernetes cluster        |
| rds        | `modules/rds`        | PostgreSQL                |
| redis      | `modules/redis`      | ElastiCache               |
| monitoring | `modules/monitoring` | Prometheus, Grafana       |
| alb        | `modules/alb`        | Application Load Balancer |
| waf        | `modules/waf`        | Web Application Firewall  |

### Environments

| Env     | Account         | Region       | Size   |
| ------- | --------------- | ------------ | ------ |
| dev     | dev-account     | us-east-1    | Small  |
| staging | staging-account | us-east-1    | Medium |
| prod    | prod-account    | Multi-region | Large  |

### Resource Sizing

#### Development

```yaml
eks:
  node_groups:
    - name: general
      instance_types: [t3.medium]
      min_size: 2
      max_size: 4
rds:
  instance_class: db.t3.medium
  storage: 100GB
redis:
  node_type: cache.t3.medium
  num_cache_nodes: 1
```

#### Staging

```yaml
eks:
  node_groups:
    - name: general
      instance_types: [t3.large]
      min_size: 3
      max_size: 6
rds:
  instance_class: db.r6g.large
  storage: 500GB
  multi_az: true
redis:
  node_type: cache.r6g.large
  num_cache_nodes: 2
```

#### Production

```yaml
eks:
  node_groups:
    - name: general
      instance_types: [m6i.xlarge]
      min_size: 6
      max_size: 20
    - name: temporal
      instance_types: [m6i.2xlarge]
      min_size: 3
      max_size: 10
rds:
  instance_class: db.r6g.2xlarge
  storage: 2TB
  multi_az: true
  read_replicas: 2
redis:
  node_type: cache.r6g.xlarge
  num_cache_nodes: 3
  cluster_mode: true
```

### Network Design

#### VPC CIDR Allocation

| Environment | VPC CIDR     | Public Subnets | Private Subnets | Database Subnets |
| ----------- | ------------ | -------------- | --------------- | ---------------- |
| dev         | 10.0.0.0/16  | 10.0.0.0/20    | 10.0.16.0/20    | 10.0.32.0/20     |
| staging     | 10.1.0.0/16  | 10.1.0.0/20    | 10.1.16.0/20    | 10.1.32.0/20     |
| prod-us     | 10.10.0.0/16 | 10.10.0.0/20   | 10.10.16.0/20   | 10.10.32.0/20    |
| prod-eu     | 10.20.0.0/16 | 10.20.0.0/20   | 10.20.16.0/20   | 10.20.32.0/20    |
| prod-ap     | 10.30.0.0/16 | 10.30.0.0/20   | 10.30.16.0/20   | 10.30.32.0/20    |

#### Security Groups

| Name     | Inbound            | Outbound   |
| -------- | ------------------ | ---------- |
| alb-sg   | 443 from 0.0.0.0/0 | All to VPC |
| eks-sg   | All from alb-sg    | All        |
| rds-sg   | 5432 from eks-sg   | None       |
| redis-sg | 6379 from eks-sg   | None       |

### Kubernetes Resources

#### Namespaces

- `temporal-system` - Temporal server components
- `cloud-platform` - Cloud platform services
- `monitoring` - Prometheus, Grafana, Loki
- `ingress` - ALB Ingress Controller

#### Deployments

| Service           | Replicas (Prod) | CPU | Memory |
| ----------------- | --------------- | --- | ------ |
| temporal-frontend | 3               | 2   | 4Gi    |
| temporal-history  | 3               | 4   | 8Gi    |
| temporal-matching | 3               | 2   | 4Gi    |
| temporal-worker   | 3               | 2   | 4Gi    |
| cloud-api         | 3               | 1   | 2Gi    |
| cloud-console     | 2               | 0.5 | 1Gi    |

### DNS & Certificates

#### Domain Structure

```
temporal-cloud.io
├── api.temporal-cloud.io        → Cloud API
├── console.temporal-cloud.io    → Web Console
├── *.tmprl.cloud                → Namespace endpoints
└── grpc.tmprl.cloud             → gRPC endpoints
```

#### Certificate Management

- AWS Certificate Manager for ALB
- cert-manager for internal TLS
- Let's Encrypt for public endpoints

<!-- Source: limits.md -->

## System Limits

### Account Level

| Limit             | Default | Max             | How to Increase           |
| ----------------- | ------- | --------------- | ------------------------- |
| Users             | 300     | Contact support | Support ticket            |
| Namespaces        | 10      | 100 (auto)      | Auto-scales, then support |
| Service Accounts  | 100     | Contact support | Support ticket            |
| API Keys per user | 10      | 50              | Support ticket            |

### Namespace Level

| Limit                    | Default               | Notes                            |
| ------------------------ | --------------------- | -------------------------------- |
| Actions/sec (APS)        | 400                   | Auto-scales based on 7-day usage |
| Requests/sec (RPS)       | 1600                  | Auto-scales                      |
| Operations/sec (OPS)     | 3200                  | Auto-scales                      |
| Schedules RPS            | 10                    | Fixed                            |
| Visibility API RPS       | 30                    | Fixed                            |
| Nexus RPS                | Part of namespace RPS | Shared limit                     |
| Certificates             | 16 or 32KB            | Whichever first                  |
| Retention period         | 1-90 days             | Per namespace config             |
| Concurrent pollers       | 2000                  | Per task queue                   |
| Custom Search Attributes | 100                   | Per namespace                    |

### Workflow Level

| Limit                      | Value      | Notes                          |
| -------------------------- | ---------- | ------------------------------ |
| Identifier length          | 1000 bytes | Unicode may use multiple bytes |
| gRPC message size          | 4 MB       | All endpoints                  |
| Event History transaction  | 4 MB       | Non-configurable               |
| Payload size               | 2 MB       | Single request                 |
| Concurrent Activities      | 2000       | Per workflow                   |
| Concurrent Child Workflows | 2000       | Per workflow                   |
| Concurrent Signals         | 2000       | Per workflow                   |
| Signals per execution      | 10,000     | Total lifetime                 |
| In-flight Updates          | 10         | Concurrent                     |
| Total Updates              | 2000       | In history                     |
| Event History events       | 51,200     | Warning at 10,240              |
| Event History size         | 50 MB      | Warning at 10 MB               |
| Callbacks per workflow     | 32         | Total                          |
| In-flight Nexus ops        | 30         | Concurrent                     |

### Nexus Limits

| Limit                          | Value              |
| ------------------------------ | ------------------ |
| Endpoints per account          | 100                |
| Allowlist entries per endpoint | 100                |
| Operation request timeout      | 10 seconds         |
| Async operation timeout        | 24 hours (default) |

### Worker Versioning Limits

| Limit                     | Value |
| ------------------------- | ----- |
| Deployments per namespace | 100   |
| Versions per deployment   | 100   |
| Task queues per version   | 1000  |

### Rate Limit Behavior

#### Throttling Priority

1. External events (Critical) - Never throttled
2. Workflow progress updates
3. Visibility API calls
4. Cloud operations

#### When Throttled

- Requests are delayed, not dropped
- High-priority calls never delayed
- Workers may take longer to complete

### Increasing Limits

#### Automatic Scaling

- APS, RPS, OPS auto-scale based on 7-day usage
- Never falls below default
- Scales up within minutes of increased usage

#### Manual Increase

1. Open support ticket
2. Provide justification and expected usage
3. Temporal reviews and approves
4. Limit increased within 24-48 hours

<!-- Source: log-management.md -->

## Log Management

### Architecture

```
[Pod] -> [Stdout] -> [FluentBit] -> [Kinesis] -> [OpenSearch/Loki] -> [S3 Archive]
```

### Log Levels

| Level     | Meaning                                   | Example                              |
| --------- | ----------------------------------------- | ------------------------------------ |
| **ERROR** | Action failed, human intervention needed. | DB connection lost.                  |
| **WARN**  | Action failed, handled gracefully.        | Workflow retried.                    |
| **INFO**  | Major lifecycle events.                   | Workflow started, Namespace created. |
| **DEBUG** | Granular logic flow (Sampled).            | Request received, state changed.     |

### Structural Logging

JSON format mandatory for parsing.

```json
{
  "level": "info",
  "ts": "2025-01-01T12:00:00Z",
  "caller": "history/handler.go:123",
  "msg": "Processing task",
  "service": "history",
  "shard_id": 42,
  "namespace_id": "ns-123",
  "workflow_id": "wf-abc",
  "trace_id": "1a2b3c"
}
```

### Sensitive Data Redaction

- **PII**: Email, Names, IP (hash or mask).
- **Secrets**: API Keys, Passwords (REDACTED).
- **Payloads**: Workflow inputs/results (NEVER log full payloads).

### Retention Policies

| Log Type       | Hot Storage (Searchable) | Cold Storage (Archive) |
| -------------- | ------------------------ | ---------------------- |
| Access/Audit   | 90 days                  | 7 years                |
| Error/Warn     | 30 days                  | 1 year                 |
| Info (Sampled) | 7 days                   | 90 days                |
| Debug          | 3 days                   | 30 days                |

### Cost Control

#### Indexing Rules

- **Index**: `level`, `service`, `trace_id`, `namespace_id`, `error`.
- **Do Not Index**: `stack_trace`, free-text `msg`.

#### Sampling

- Use dynamic sampling based on volume. If 1000 logs/sec from one workflow, sample at 0.1%.

### Access Control

- **Devs**: Read access to non-PII logs.
- **SRE**: Full access.
- **Customers**: Access ONLY to their namespace's logs via Export API (not direct access).

<!-- Source: logic/billing-reconciliation.md -->

## Billing Reconciliation Logic

### Problem Statement

Usage is aggregated hourly. Invoices are monthly. Stripe subscriptions handle the fixed fee, but usage must be reported accurately and idempotently.

### Architecture

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

### Algorithm: ReportUsageWorkflow

#### Trigger

- Runs every 1 hour per organization
- Scheduled via Temporal Schedule

#### Workflow Definition

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

#### Stripe Reporting

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

### Proration Logic

#### Plan Upgrades

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

#### Plan Downgrades

- **Timing**: Next billing cycle
- **Calculation**: No immediate charge
- **Included quotas**: Current plan until end of cycle

### Invoice Generation

#### Monthly Invoice Workflow

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

#### Tiered Action Pricing

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

### Error Handling

#### Stripe API Failures

- Retry with exponential backoff
- Max 5 retries over 1 hour
- Alert on persistent failure
- Manual intervention if still failing

#### Usage Data Missing

- Log warning
- Use last known good data
- Flag invoice for review

#### Duplicate Prevention

- Idempotency key: `{org_id}:{period_start}:{period_end}`
- Check before processing
- Mark as processed atomically

<!-- Source: logic/namespace-failover.md -->

## Namespace Failover Logic

### Problem Statement

Fail over a namespace from one region to another without data loss, maintaining workflow consistency.

### Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                    Primary Region                            │
│  ┌─────────────┐     ┌─────────────┐     ┌─────────────┐   │
│  │   Workers   │────▶│  Temporal   │────▶│  Database   │   │
│  └─────────────┘     └─────────────┘     └─────────────┘   │
│                             │                               │
│                             │ Replication                   │
│                             ▼                               │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                   Secondary Region                           │
│  ┌─────────────┐     ┌─────────────┐     ┌─────────────┐   │
│  │   Workers   │────▶│  Temporal   │────▶│  Database   │   │
│  │  (standby)  │     │  (standby)  │     │  (replica)  │   │
│  └─────────────┘     └─────────────┘     └─────────────┘   │
└─────────────────────────────────────────────────────────────┘
```

### Failover State Machine

```
┌─────────┐    Initiate    ┌─────────────┐    Fence    ┌─────────┐
│ Active  │───────────────▶│  Fencing    │────────────▶│ Syncing │
└─────────┘                └─────────────┘             └─────────┘
                                                            │
                                                            │ Sync Complete
                                                            ▼
┌─────────┐    Complete    ┌─────────────┐    Promote  ┌─────────┐
│ Active  │◀───────────────│  Switching  │◀────────────│ Syncing │
│ (new)   │                └─────────────┘             └─────────┘
```

### Failover Workflow

```go
func NamespaceFailoverWorkflow(ctx workflow.Context, input FailoverInput) error {
    logger := workflow.GetLogger(ctx)

    // Step 1: Validate failover request
    var validation ValidationResult
    err := workflow.ExecuteActivity(ctx, ValidateFailover, input).Get(ctx, &validation)
    if err != nil || !validation.Valid {
        return fmt.Errorf("failover validation failed: %v", validation.Reason)
    }

    // Step 2: Fence the primary region
    // This prevents new writes to the primary
    logger.Info("Fencing primary region", "region", input.SourceRegion)
    err = workflow.ExecuteActivity(ctx, FencePrimaryRegion, FenceInput{
        NamespaceID: input.NamespaceID,
        Region:      input.SourceRegion,
    }).Get(ctx, nil)
    if err != nil {
        return fmt.Errorf("failed to fence primary: %w", err)
    }

    // Step 3: Wait for replication to catch up
    logger.Info("Waiting for replication sync")
    var syncStatus SyncStatus
    for {
        err = workflow.ExecuteActivity(ctx, CheckReplicationLag, input.NamespaceID).Get(ctx, &syncStatus)
        if err != nil {
            return fmt.Errorf("failed to check replication: %w", err)
        }

        if syncStatus.LagSeconds < 1 {
            break
        }

        logger.Info("Replication lag", "seconds", syncStatus.LagSeconds)
        workflow.Sleep(ctx, 5*time.Second)
    }

    // Step 4: Promote secondary to primary
    logger.Info("Promoting secondary region", "region", input.TargetRegion)
    err = workflow.ExecuteActivity(ctx, PromoteSecondary, PromoteInput{
        NamespaceID: input.NamespaceID,
        Region:      input.TargetRegion,
    }).Get(ctx, nil)
    if err != nil {
        // Attempt rollback
        workflow.ExecuteActivity(ctx, UnfencePrimary, input.SourceRegion)
        return fmt.Errorf("failed to promote secondary: %w", err)
    }

    // Step 5: Update DNS/routing
    logger.Info("Updating DNS routing")
    err = workflow.ExecuteActivity(ctx, UpdateDNSRouting, DNSInput{
        NamespaceID: input.NamespaceID,
        NewRegion:   input.TargetRegion,
    }).Get(ctx, nil)
    if err != nil {
        return fmt.Errorf("failed to update DNS: %w", err)
    }

    // Step 6: Notify customers
    workflow.ExecuteActivity(ctx, NotifyFailoverComplete, input)

    // Step 7: Demote old primary to secondary
    workflow.ExecuteActivity(ctx, DemoteToSecondary, DemoteInput{
        NamespaceID: input.NamespaceID,
        Region:      input.SourceRegion,
    })

    logger.Info("Failover complete")
    return nil
}
```

### Fencing Mechanism

#### Purpose

Prevent split-brain by ensuring only one region accepts writes.

#### Implementation

```go
func FencePrimaryRegion(ctx context.Context, input FenceInput) error {
    // Step 1: Set namespace to read-only mode
    err := setNamespaceReadOnly(ctx, input.NamespaceID, input.Region)
    if err != nil {
        return err
    }

    // Step 2: Wait for in-flight operations to complete
    timeout := time.After(30 * time.Second)
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-timeout:
            return fmt.Errorf("timeout waiting for in-flight operations")
        case <-ticker.C:
            count, err := getInFlightOperations(ctx, input.NamespaceID)
            if err != nil {
                return err
            }
            if count == 0 {
                return nil
            }
        }
    }
}
```

### Replication Sync Check

```go
func CheckReplicationLag(ctx context.Context, namespaceID string) (SyncStatus, error) {
    // Get last written position in primary
    primaryPos, err := getPrimaryPosition(ctx, namespaceID)
    if err != nil {
        return SyncStatus{}, err
    }

    // Get last replicated position in secondary
    secondaryPos, err := getSecondaryPosition(ctx, namespaceID)
    if err != nil {
        return SyncStatus{}, err
    }

    // Calculate lag
    lag := primaryPos.Timestamp.Sub(secondaryPos.Timestamp)

    return SyncStatus{
        PrimaryPosition:   primaryPos,
        SecondaryPosition: secondaryPos,
        LagSeconds:        int(lag.Seconds()),
        InSync:            lag < time.Second,
    }, nil
}
```

### DNS Update

```go
func UpdateDNSRouting(ctx context.Context, input DNSInput) error {
    // Get namespace endpoint
    endpoint := fmt.Sprintf("%s.tmprl.cloud", input.NamespaceID)

    // Update Route53 record to point to new region
    _, err := route53Client.ChangeResourceRecordSets(ctx, &route53.ChangeResourceRecordSetsInput{
        HostedZoneId: aws.String(hostedZoneID),
        ChangeBatch: &types.ChangeBatch{
            Changes: []types.Change{
                {
                    Action: types.ChangeActionUpsert,
                    ResourceRecordSet: &types.ResourceRecordSet{
                        Name: aws.String(endpoint),
                        Type: types.RRTypeA,
                        AliasTarget: &types.AliasTarget{
                            DNSName:              aws.String(regionEndpoints[input.NewRegion]),
                            HostedZoneId:         aws.String(regionHostedZones[input.NewRegion]),
                            EvaluateTargetHealth: aws.Bool(true),
                        },
                    },
                },
            },
        },
    })

    return err
}
```

### Rollback Procedure

If failover fails mid-way:

```go
func RollbackFailover(ctx workflow.Context, input FailoverInput) error {
    // Step 1: Unfence primary (if fenced)
    workflow.ExecuteActivity(ctx, UnfencePrimary, input.SourceRegion)

    // Step 2: Revert DNS (if changed)
    workflow.ExecuteActivity(ctx, UpdateDNSRouting, DNSInput{
        NamespaceID: input.NamespaceID,
        NewRegion:   input.SourceRegion, // Back to original
    })

    // Step 3: Demote secondary (if promoted)
    workflow.ExecuteActivity(ctx, DemoteToSecondary, DemoteInput{
        NamespaceID: input.NamespaceID,
        Region:      input.TargetRegion,
    })

    // Step 4: Notify of rollback
    workflow.ExecuteActivity(ctx, NotifyFailoverRollback, input)

    return nil
}
```

### Timing Expectations

| Phase            | Expected Duration            |
| ---------------- | ---------------------------- |
| Validation       | < 5 seconds                  |
| Fencing          | < 30 seconds                 |
| Replication sync | < 5 minutes (depends on lag) |
| Promotion        | < 10 seconds                 |
| DNS update       | < 60 seconds (propagation)   |
| **Total RTO**    | **< 15 minutes**             |

### Monitoring

#### Metrics

- `failover_duration_seconds`
- `replication_lag_seconds`
- `failover_success_total`
- `failover_failure_total`

#### Alerts

- Replication lag > 5 minutes
- Failover duration > 30 minutes
- Failover failure

<!-- Source: logic/quota-enforcement.md -->

## Quota Enforcement Logic

### Problem Statement

Enforce plan-based quotas (actions/sec, storage) across a distributed system with minimal latency impact.

### Architecture

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Request   │────▶│   Quota     │────▶│   Temporal  │
│             │     │  Enforcer   │     │   Server    │
└─────────────┘     └─────────────┘     └─────────────┘
                          │
                          ▼
                    ┌─────────────┐
                    │   Redis     │
                    │   Cache     │
                    └─────────────┘
```

### Token Bucket Algorithm

#### Concept

- Each namespace has a "bucket" of tokens
- Tokens replenish at a fixed rate (APS limit)
- Each action consumes one token
- Request rejected if bucket empty

#### Implementation

```go
type TokenBucket struct {
    redis     *redis.Client
    namespace string
    limit     int64 // tokens per second
}

func (tb *TokenBucket) Allow(ctx context.Context) (bool, error) {
    now := time.Now().UnixNano()
    key := fmt.Sprintf("quota:%s", tb.namespace)

    // Lua script for atomic check-and-decrement
    script := `
        local key = KEYS[1]
        local limit = tonumber(ARGV[1])
        local now = tonumber(ARGV[2])
        local window = 1000000000 -- 1 second in nanoseconds

        -- Get current bucket state
        local bucket = redis.call('HMGET', key, 'tokens', 'last_update')
        local tokens = tonumber(bucket[1]) or limit
        local last_update = tonumber(bucket[2]) or now

        -- Calculate tokens to add based on time elapsed
        local elapsed = now - last_update
        local tokens_to_add = (elapsed / window) * limit
        tokens = math.min(limit, tokens + tokens_to_add)

        -- Try to consume a token
        if tokens >= 1 then
            tokens = tokens - 1
            redis.call('HMSET', key, 'tokens', tokens, 'last_update', now)
            redis.call('EXPIRE', key, 10)
            return 1
        else
            return 0
        end
    `

    result, err := tb.redis.Eval(ctx, script, []string{key}, tb.limit, now).Int()
    if err != nil {
        // On Redis failure, allow request (fail open)
        return true, err
    }

    return result == 1, nil
}
```

### Distributed Rate Limiting

#### Challenge

Multiple frontend instances need coordinated rate limiting.

#### Solution: Sliding Window with Redis

```go
type SlidingWindowLimiter struct {
    redis     *redis.Client
    namespace string
    limit     int64
    window    time.Duration
}

func (swl *SlidingWindowLimiter) Allow(ctx context.Context) (bool, int64, error) {
    now := time.Now()
    windowStart := now.Add(-swl.window).UnixNano()
    key := fmt.Sprintf("ratelimit:%s", swl.namespace)

    pipe := swl.redis.Pipeline()

    // Remove old entries
    pipe.ZRemRangeByScore(ctx, key, "0", fmt.Sprintf("%d", windowStart))

    // Count current entries
    countCmd := pipe.ZCard(ctx, key)

    // Add current request (optimistically)
    pipe.ZAdd(ctx, key, &redis.Z{
        Score:  float64(now.UnixNano()),
        Member: fmt.Sprintf("%d:%s", now.UnixNano(), uuid.New().String()),
    })

    // Set expiry
    pipe.Expire(ctx, key, swl.window*2)

    _, err := pipe.Exec(ctx)
    if err != nil {
        return true, 0, err // Fail open
    }

    count := countCmd.Val()
    if count >= swl.limit {
        // Remove the optimistically added entry
        swl.redis.ZRemRangeByRank(ctx, key, -1, -1)
        return false, swl.limit - count, nil
    }

    return true, swl.limit - count - 1, nil
}
```

### Quota Cache

#### Local Cache for Performance

```go
type QuotaCache struct {
    cache    *lru.Cache
    store    QuotaStore
    ttl      time.Duration
}

type CachedQuota struct {
    Limit     int64
    ExpiresAt time.Time
}

func (qc *QuotaCache) GetLimit(ctx context.Context, namespace string) (int64, error) {
    // Check local cache
    if cached, ok := qc.cache.Get(namespace); ok {
        cq := cached.(*CachedQuota)
        if time.Now().Before(cq.ExpiresAt) {
            return cq.Limit, nil
        }
    }

    // Fetch from store
    limit, err := qc.store.GetNamespaceLimit(ctx, namespace)
    if err != nil {
        return 0, err
    }

    // Update cache
    qc.cache.Add(namespace, &CachedQuota{
        Limit:     limit,
        ExpiresAt: time.Now().Add(qc.ttl),
    })

    return limit, nil
}
```

### gRPC Interceptor

```go
func QuotaInterceptor(enforcer *QuotaEnforcer) grpc.UnaryServerInterceptor {
    return func(
        ctx context.Context,
        req interface{},
        info *grpc.UnaryServerInfo,
        handler grpc.UnaryHandler,
    ) (interface{}, error) {
        // Extract namespace from request
        namespace := extractNamespace(req)
        if namespace == "" {
            return handler(ctx, req)
        }

        // Check if this is a billable action
        if !isBillableAction(info.FullMethod) {
            return handler(ctx, req)
        }

        // Check quota
        allowed, remaining, err := enforcer.Check(ctx, namespace)
        if err != nil {
            // Log error but allow request (fail open)
            log.Warn("quota check failed", "error", err)
            return handler(ctx, req)
        }

        if !allowed {
            return nil, status.Errorf(
                codes.ResourceExhausted,
                "rate limit exceeded for namespace %s, retry after %v",
                namespace,
                time.Second,
            )
        }

        // Add remaining quota to response headers
        grpc.SetHeader(ctx, metadata.Pairs(
            "x-ratelimit-remaining", fmt.Sprintf("%d", remaining),
        ))

        return handler(ctx, req)
    }
}
```

### Priority Levels

#### Request Prioritization

```go
type Priority int

const (
    PriorityCritical Priority = iota // External events, never throttled
    PriorityHigh                     // Workflow progress
    PriorityMedium                   // Visibility API
    PriorityLow                      // Cloud operations
)

func getPriority(method string) Priority {
    switch {
    case isExternalEvent(method):
        return PriorityCritical
    case isWorkflowProgress(method):
        return PriorityHigh
    case isVisibilityAPI(method):
        return PriorityMedium
    default:
        return PriorityLow
    }
}

func (e *QuotaEnforcer) Check(ctx context.Context, namespace string) (bool, int64, error) {
    priority := getPriority(getMethod(ctx))

    // Critical priority never throttled
    if priority == PriorityCritical {
        return true, -1, nil
    }

    // Check rate limit
    return e.limiter.Allow(ctx, namespace)
}
```

### Auto-Scaling Limits

#### Dynamic Limit Adjustment

```go
func (e *QuotaEnforcer) AdjustLimits(ctx context.Context) error {
    // Run daily
    namespaces, err := e.store.ListNamespaces(ctx)
    if err != nil {
        return err
    }

    for _, ns := range namespaces {
        // Get 7-day usage
        usage, err := e.store.Get7DayUsage(ctx, ns.ID)
        if err != nil {
            continue
        }

        // Calculate new limit (max of default and 1.5x peak usage)
        peakAPS := usage.PeakActionsPerSecond
        newLimit := max(DefaultAPS, int64(float64(peakAPS)*1.5))

        // Update limit
        e.store.UpdateLimit(ctx, ns.ID, newLimit)
    }

    return nil
}
```

<!-- Source: metrics-endpoint.md -->

## Metrics Endpoint

### Overview

The Metrics Endpoint allows customers to scrape Prometheus-compatible metrics for their own namespaces. This enables them to build custom alerts and dashboards in their own monitoring systems (Datadog, Grafana Cloud, etc.).

### Configuration

#### 1. Generate Certificate

Metrics are protected by mTLS. You must generate a dedicated client certificate for scraping.

```bash
tcld namespace metrics-cert create \
  --namespace my-ns \
  --cert metrics-client.pem
```

#### 2. Endpoint URL

`https://<namespace>.<account>.tmprl.cloud/prometheus/metrics`

### Scrape Configuration

#### Prometheus (`prometheus.yml`)

```yaml
scrape_configs:
  - job_name: "temporal-cloud"
    scrape_interval: 1m
    scheme: https
    static_configs:
      - targets: ["my-ns.a1b2c.tmprl.cloud:443"]
    metrics_path: /prometheus/metrics
    tls_config:
      cert_file: /path/to/metrics-client.pem
      key_file: /path/to/metrics-client.key
      insecure_skip_verify: false
```

#### Datadog

Use the OpenMetrics integration.

```yaml
# conf.d/openmetrics.d/conf.yaml
instances:
  - prometheus_url: https://my-ns.a1b2c.tmprl.cloud/prometheus/metrics
    namespace: temporal_cloud
    metrics:
      - temporal_cloud*
    tls_cert: /path/to/cert.pem
    tls_private_key: /path/to/key.pem
```

### Available Metrics

#### Workflow Metrics

- `temporal_cloud_workflow_start_count`: Counter
- `temporal_cloud_workflow_success_count`: Counter
- `temporal_cloud_workflow_failed_count`: Counter
- `temporal_cloud_workflow_latency_seconds`: Histogram

#### Activity Metrics

- `temporal_cloud_activity_start_count`: Counter
- `temporal_cloud_activity_failed_count`: Counter
- `temporal_cloud_activity_latency_seconds`: Histogram

#### Task Queue Metrics

- `temporal_cloud_task_queue_backlog_count`: Gauge
- `temporal_cloud_task_queue_latency_seconds`: Histogram (Schedule-to-Start)

#### Resource Metrics

- `temporal_cloud_action_count`: Counter (Billable actions)
- `temporal_cloud_storage_active_bytes`: Gauge
- `temporal_cloud_storage_retained_bytes`: Gauge

### Labels

All metrics include:

- `namespace`
- `operation` (e.g., StartWorkflowExecution)
- `task_queue`
- `workflow_type`
- `activity_type`
- `status` (for counters)

### Cardinality Limits

To protect the system, high-cardinality labels (like Workflow ID) are **NOT** included in this endpoint.

<!-- Source: monitoring.md -->

## Monitoring & Alerting

### SLIs and SLOs

| Service   | SLI              | SLO     | Alert Threshold |
| --------- | ---------------- | ------- | --------------- |
| Cloud API | Availability     | 99.9%   | < 99.5% (5m)    |
| Cloud API | Latency P99      | < 200ms | > 500ms (5m)    |
| Console   | Page Load        | < 2s    | > 5s (5m)       |
| Billing   | Invoice Accuracy | 99.99%  | Any error       |
| Metering  | Lag              | < 5m    | > 10m           |

### Monitoring Stack

| Component  | Tool         | Purpose             |
| ---------- | ------------ | ------------------- |
| Metrics    | Prometheus   | Time-series metrics |
| Dashboards | Grafana      | Visualization       |
| Logs       | Loki         | Log aggregation     |
| Traces     | Jaeger       | Distributed tracing |
| Alerts     | Alertmanager | Alert routing       |
| Uptime     | Pingdom      | External monitoring |

### Key Metrics

#### Application Metrics

```
# Request rate
sum(rate(http_requests_total[5m])) by (service)

# Error rate
sum(rate(http_requests_total{status=~"5.."}[5m])) / sum(rate(http_requests_total[5m]))

# Latency P99
histogram_quantile(0.99, rate(http_request_duration_seconds_bucket[5m]))
```

#### Infrastructure Metrics

```
# CPU usage
avg(rate(container_cpu_usage_seconds_total[5m])) by (pod)

# Memory usage
container_memory_usage_bytes / container_spec_memory_limit_bytes

# Disk usage
(node_filesystem_size_bytes - node_filesystem_free_bytes) / node_filesystem_size_bytes
```

#### Business Metrics

```
# Active organizations
count(temporal_cloud_organizations_total)

# Actions per second
sum(rate(temporal_cloud_actions_total[5m]))

# Revenue (monthly)
sum(temporal_cloud_invoice_total_cents) / 100
```

### Dashboards

| Dashboard         | Purpose                     | Audience    |
| ----------------- | --------------------------- | ----------- |
| Platform Overview | High-level health           | On-call     |
| API Performance   | Latency, errors, throughput | Engineering |
| Billing Metrics   | Revenue, usage              | Finance     |
| Infrastructure    | K8s, RDS, Redis             | Platform    |
| Customer Health   | Per-org metrics             | Support     |

### Alert Routing

| Severity      | Response Time     | Channel           |
| ------------- | ----------------- | ----------------- |
| P1 (Critical) | 15 min            | PagerDuty + Slack |
| P2 (High)     | 1 hour            | PagerDuty + Slack |
| P3 (Medium)   | 4 hours           | Slack             |
| P4 (Low)      | Next business day | Email             |

### Alert Definitions

#### P1 - Critical

```yaml
- alert: APIDown
  expr: up{job="cloud-api"} == 0
  for: 1m
  labels:
    severity: critical
  annotations:
    summary: "Cloud API is down"

- alert: DatabaseDown
  expr: pg_up == 0
  for: 1m
  labels:
    severity: critical
```

#### P2 - High

```yaml
- alert: HighErrorRate
  expr: rate(http_requests_total{status=~"5.."}[5m]) > 0.05
  for: 5m
  labels:
    severity: high

- alert: HighLatency
  expr: histogram_quantile(0.99, rate(http_request_duration_seconds_bucket[5m])) > 0.5
  for: 5m
  labels:
    severity: high
```

#### P3 - Medium

```yaml
- alert: DiskSpaceLow
  expr: (node_filesystem_free_bytes / node_filesystem_size_bytes) < 0.2
  for: 15m
  labels:
    severity: medium

- alert: MeteringLag
  expr: temporal_cloud_metering_lag_seconds > 600
  for: 10m
  labels:
    severity: medium
```

### On-Call Rotation

#### Schedule

- Primary: 1 week rotation
- Secondary: 1 week rotation (backup)
- Escalation: Engineering Manager → VP Engineering

#### Responsibilities

- Acknowledge alerts within SLA
- Investigate and mitigate issues
- Document incidents
- Hand off to next on-call

<!-- Source: multi-cloud.md -->

## Multi-Cloud Strategy

### Overview

Temporal Cloud supports deployment across AWS and GCP to provide:

- Geographic coverage
- Vendor redundancy
- Customer preference accommodation
- Regulatory compliance (data residency)

### Cloud Coverage

#### Primary Cloud: AWS

| Region      | Code           | Services       |
| ----------- | -------------- | -------------- |
| N. Virginia | us-east-1      | Full (Primary) |
| Oregon      | us-west-2      | Full           |
| Ireland     | eu-west-1      | Full           |
| Frankfurt   | eu-central-1   | Full           |
| Singapore   | ap-southeast-1 | Full           |
| Sydney      | ap-southeast-2 | Full           |
| Tokyo       | ap-northeast-1 | Full           |
| Mumbai      | ap-south-1     | Full           |

#### Secondary Cloud: GCP

| Region      | Code         | Services |
| ----------- | ------------ | -------- |
| Iowa        | us-central1  | Full     |
| Oregon      | us-west1     | Full     |
| N. Virginia | us-east4     | Full     |
| Frankfurt   | europe-west3 | Full     |
| Mumbai      | asia-south1  | Full     |

### Architecture

#### Control Plane

Single control plane on AWS, manages all clouds:

```
┌─────────────────────────────────────────────────────────────────┐
│                    Control Plane (AWS us-east-1)                 │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐              │
│  │  Cloud API  │  │   Billing   │  │  Provision  │              │
│  └─────────────┘  └─────────────┘  └─────────────┘              │
└───────────────────────────┬─────────────────────────────────────┘
                            │
            ┌───────────────┼───────────────┐
            ▼               ▼               ▼
      ┌──────────┐    ┌──────────┐    ┌──────────┐
      │   AWS    │    │   AWS    │    │   GCP    │
      │us-east-1 │    │eu-west-1 │    │us-central│
      └──────────┘    └──────────┘    └──────────┘
```

#### Data Plane

Each cloud region runs independently:

```yaml
# Per-region components
- Temporal Server (Frontend, History, Matching, Worker)
- PostgreSQL (Regional, no cross-cloud replication by default)
- Redis (Regional)
- Load Balancer (Cloud-native: ALB/Cloud Load Balancing)
```

### Cross-Cloud Replication

#### Namespace Replication Pairs

| Primary (AWS) | Secondary (GCP) | Latency |
| ------------- | --------------- | ------- |
| us-east-1     | us-central1     | ~30ms   |
| us-west-2     | us-west1        | ~20ms   |
| eu-central-1  | europe-west3    | ~10ms   |
| ap-south-1    | asia-south1     | ~5ms    |

#### Replication Architecture

```
┌─────────────────────┐         ┌─────────────────────┐
│    AWS us-east-1    │         │   GCP us-central1   │
│  ┌───────────────┐  │         │  ┌───────────────┐  │
│  │   Temporal    │──┼─────────┼─▶│   Temporal    │  │
│  │   Primary     │  │  gRPC   │  │   Standby     │  │
│  └───────────────┘  │         │  └───────────────┘  │
│  ┌───────────────┐  │         │  ┌───────────────┐  │
│  │  PostgreSQL   │──┼─────────┼─▶│  PostgreSQL   │  │
│  │   Primary     │  │  Async  │  │   Replica     │  │
│  └───────────────┘  │         │  └───────────────┘  │
└─────────────────────┘         └─────────────────────┘
```

#### Cross-Cloud Connectivity

**Option 1: Public Internet (Encrypted)**

- TLS 1.3 for all traffic
- IP allowlisting
- Simpler, higher latency

**Option 2: Dedicated Interconnect**

- AWS Direct Connect + GCP Partner Interconnect
- Private connectivity
- Lower latency, higher cost
- Used for high-volume replication

### Terraform Multi-Cloud

#### Provider Configuration

```hcl
# providers.tf
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
    google = {
      source  = "hashicorp/google"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
  alias  = "primary"
}

provider "google" {
  project = var.gcp_project
  region  = var.gcp_region
  alias   = "secondary"
}
```

#### Module Structure

```
terraform/
├── modules/
│   ├── aws/
│   │   ├── vpc/
│   │   ├── eks/
│   │   ├── rds/
│   │   └── temporal/
│   ├── gcp/
│   │   ├── vpc/
│   │   ├── gke/
│   │   ├── cloudsql/
│   │   └── temporal/
│   └── common/
│       ├── monitoring/
│       └── dns/
├── environments/
│   ├── prod-aws-us-east-1/
│   ├── prod-aws-eu-west-1/
│   ├── prod-gcp-us-central1/
│   └── staging/
```

#### Unified Resource Abstraction

```hcl
# modules/common/kubernetes-cluster/main.tf
module "aws_cluster" {
  count  = var.cloud == "aws" ? 1 : 0
  source = "../../aws/eks"
  # ...
}

module "gcp_cluster" {
  count  = var.cloud == "gcp" ? 1 : 0
  source = "../../gcp/gke"
  # ...
}

output "cluster_endpoint" {
  value = var.cloud == "aws" ? module.aws_cluster[0].endpoint : module.gcp_cluster[0].endpoint
}
```

### Cloud-Agnostic Components

#### Kubernetes (EKS/GKE)

Same Helm charts, different values:

```yaml
# values-aws.yaml
ingress:
  class: alb
  annotations:
    alb.ingress.kubernetes.io/scheme: internet-facing

# values-gcp.yaml
ingress:
  class: gce
  annotations:
    kubernetes.io/ingress.class: gce
```

#### Database (RDS/Cloud SQL)

Same schema, different provisioning:

```hcl
# AWS
resource "aws_db_instance" "temporal" {
  engine         = "postgres"
  engine_version = "15"
  instance_class = "db.r6g.xlarge"
}

# GCP
resource "google_sql_database_instance" "temporal" {
  database_version = "POSTGRES_15"
  settings {
    tier = "db-custom-4-16384"
  }
}
```

#### Object Storage (S3/GCS)

```go
// Cloud-agnostic storage interface
type ObjectStore interface {
    Put(ctx context.Context, key string, data []byte) error
    Get(ctx context.Context, key string) ([]byte, error)
    Delete(ctx context.Context, key string) error
}

// Factory
func NewObjectStore(cloud, bucket string) ObjectStore {
    switch cloud {
    case "aws":
        return &S3Store{bucket: bucket}
    case "gcp":
        return &GCSStore{bucket: bucket}
    }
}
```

### DNS & Traffic Management

#### Global Load Balancing

```
                    ┌─────────────────────┐
                    │   Route53 / Cloud   │
                    │       DNS           │
                    └──────────┬──────────┘
                               │ Latency-based routing
           ┌───────────────────┼───────────────────┐
           ▼                   ▼                   ▼
    ┌─────────────┐     ┌─────────────┐     ┌─────────────┐
    │ AWS ALB     │     │ AWS ALB     │     │ GCP LB      │
    │ us-east-1   │     │ eu-west-1   │     │ us-central1 │
    └─────────────┘     └─────────────┘     └─────────────┘
```

#### Failover Configuration

```hcl
resource "aws_route53_health_check" "aws_region" {
  fqdn              = "us-east-1.temporal-cloud.io"
  port              = 443
  type              = "HTTPS"
  resource_path     = "/health"
  failure_threshold = 3
  request_interval  = 10
}

resource "aws_route53_record" "api" {
  zone_id = aws_route53_zone.main.zone_id
  name    = "api.temporal-cloud.io"
  type    = "A"

  set_identifier = "aws-us-east-1"

  alias {
    name    = aws_lb.api.dns_name
    zone_id = aws_lb.api.zone_id
  }

  latency_routing_policy {
    region = "us-east-1"
  }

  health_check_id = aws_route53_health_check.aws_region.id
}
```

### Cost Optimization

#### Cloud Pricing Comparison

| Resource         | AWS (us-east-1) | GCP (us-central1) |
| ---------------- | --------------- | ----------------- |
| Compute (4 vCPU) | $0.17/hr        | $0.15/hr          |
| PostgreSQL       | $0.46/hr        | $0.42/hr          |
| Egress (per GB)  | $0.09           | $0.12             |
| Load Balancer    | $0.025/hr       | $0.025/hr         |

#### Multi-Cloud Cost Strategy

1. **Committed Use**: 1-year reserved instances on primary cloud
2. **Spot/Preemptible**: For non-critical workloads
3. **Right-sizing**: Regular review of instance sizes
4. **Egress Optimization**: Minimize cross-cloud data transfer

### Monitoring Across Clouds

#### Unified Observability

```yaml
# Single Grafana instance with multiple data sources
datasources:
  - name: Prometheus-AWS
    type: prometheus
    url: https://prometheus.aws.temporal-cloud.io

  - name: Prometheus-GCP
    type: prometheus
    url: https://prometheus.gcp.temporal-cloud.io

  - name: CloudWatch
    type: cloudwatch

  - name: Stackdriver
    type: stackdriver
```

#### Cross-Cloud Alerting

```yaml
# Alert on cross-cloud replication lag
alert: CrossCloudReplicationLag
expr: |
  temporal_replication_lag_seconds{
    source_cloud="aws",
    target_cloud="gcp"
  } > 60
severity: warning
```

<!-- Source: namespaces.md -->

## Namespace Management

### Overview

A Namespace is the unit of isolation in Temporal Cloud. Each namespace has its own:

- Workflow executions
- Task queues
- Search attributes
- Certificates
- Retention period

### Namespace Properties

| Property            | Description                   | Configurable     |
| ------------------- | ----------------------------- | ---------------- |
| Name                | Unique identifier             | At creation only |
| Region              | Cloud region                  | At creation only |
| Retention           | History retention (1-90 days) | Yes              |
| Certificates        | mTLS CA certificates          | Yes              |
| Search Attributes   | Custom search attributes      | Yes              |
| Tags                | Key-value metadata            | Yes              |
| Deletion Protection | Prevent accidental delete     | Yes              |

### Operations

| Operation     | Console | tcld | API | Terraform |
| ------------- | ------- | ---- | --- | --------- |
| Create        | ✅      | ✅   | ✅  | ✅        |
| Update        | ✅      | ✅   | ✅  | ✅        |
| Delete        | ✅      | ✅   | ✅  | ✅        |
| Failover (HA) | ✅      | ✅   | ✅  | ❌        |

### Creating a Namespace

#### Via Console

1. Go to Namespaces → Create Namespace
2. Enter name and select region
3. Configure retention period
4. Add certificates
5. Click Create

#### Via tcld

```bash
tcld namespace create \
  --name my-namespace \
  --region aws-us-east-1 \
  --retention-days 7 \
  --ca-certificate ca.pem
```

#### Via Terraform

```hcl
resource "temporalcloud_namespace" "example" {
  name           = "my-namespace"
  region         = "aws-us-east-1"
  retention_days = 7

  certificate {
    certificate = file("ca.pem")
  }
}
```

### Namespace Naming

#### Rules

- 2-63 characters
- Lowercase letters, numbers, hyphens
- Must start with letter
- Must end with letter or number
- Globally unique (account-scoped)

#### Best Practices

- Include environment: `myapp-prod`, `myapp-staging`
- Include team: `payments-prod`, `orders-prod`
- Avoid generic names: `test`, `dev`

### Custom Search Attributes

#### Built-in Attributes

- `WorkflowId`
- `WorkflowType`
- `StartTime`
- `CloseTime`
- `ExecutionStatus`

#### Adding Custom Attributes

```bash
tcld namespace search-attributes add \
  --namespace my-namespace \
  --name CustomerId \
  --type Keyword

tcld namespace search-attributes add \
  --namespace my-namespace \
  --name OrderTotal \
  --type Double
```

#### Attribute Types

| Type        | Description        | Example       |
| ----------- | ------------------ | ------------- |
| Keyword     | Exact match string | `CustomerId`  |
| Text        | Full-text search   | `Description` |
| Int         | Integer            | `RetryCount`  |
| Double      | Decimal            | `OrderTotal`  |
| Bool        | Boolean            | `IsVIP`       |
| Datetime    | Timestamp          | `DueDate`     |
| KeywordList | List of keywords   | `Tags`        |

### Retention Period

- Minimum: 1 day
- Maximum: 90 days
- Default: 7 days

#### Changing Retention

```bash
tcld namespace update \
  --namespace my-namespace \
  --retention-days 30
```

**Note**: Reducing retention does not immediately delete old data.

### Deletion Protection

Prevents accidental namespace deletion.

```bash
# Enable
tcld namespace update \
  --namespace my-namespace \
  --deletion-protection enabled

# Disable (required before delete)
tcld namespace update \
  --namespace my-namespace \
  --deletion-protection disabled
```

### Tags

Key-value metadata for organization and billing.

```bash
tcld namespace update \
  --namespace my-namespace \
  --tag environment=production \
  --tag team=payments \
  --tag cost-center=12345
```

#### Tag Limits

- Max 50 tags per namespace
- Key: 1-128 characters
- Value: 0-256 characters

### Endpoints

#### gRPC Endpoint

```
<namespace>.<account-id>.tmprl.cloud:443
```

#### Regional Endpoint

```
<region>.region.tmprl.cloud:443
```

### High Availability

#### Enable Multi-Region

```bash
tcld namespace update \
  --namespace my-namespace \
  --add-region aws-us-west-2
```

#### Failover

```bash
tcld namespace failover \
  --namespace my-namespace \
  --target-region aws-us-west-2
```

<!-- Source: payment-collection.md -->

## Payment Collection & Dunning

### Payment Lifecycle

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

### Invoice Generation

#### Monthly Invoice Workflow

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

### Payment Processing

#### Automatic Payment

Stripe automatically charges the default payment method when invoice is finalized.

#### Manual Payment

For enterprise customers with NET 30 terms:

1. Invoice marked as "send_invoice"
2. Customer pays via bank transfer
3. Finance team marks as paid manually

### Dunning Process

#### Retry Schedule

| Attempt | Day | Action                           |
| ------- | --- | -------------------------------- |
| 1       | 0   | Initial charge                   |
| 2       | 3   | Retry + email reminder           |
| 3       | 7   | Retry + email warning            |
| 4       | 14  | Final retry + email final notice |
| -       | 21  | Account suspended                |
| -       | 51  | Account canceled + data deletion |

#### Dunning Workflow

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

### Email Templates

#### Invoice Created

```
Subject: Your Temporal Cloud Invoice for {{month}}

Hi {{name}},

Your invoice for {{month}} is ready.

Amount Due: ${{amount}}
Due Date: {{due_date}}

[View Invoice]  [Pay Now]
```

#### Payment Failed (Attempt 1)

```
Subject: Payment failed for your Temporal Cloud invoice

Hi {{name}},

We were unable to charge your payment method for your
{{month}} invoice of ${{amount}}.

Please update your payment method to avoid service interruption.

[Update Payment Method]
```

#### Account Suspended

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

### Account Suspension

#### What Gets Suspended

- API access disabled
- Workers cannot connect
- Console shows "Account Suspended" banner
- New namespace creation blocked

#### What Continues

- Data is retained (for 30 days)
- Audit logs accessible
- Billing portal accessible

#### Reactivation

Upon payment:

1. Invoice marked as paid
2. Subscription status → active
3. Access restored within 5 minutes
4. Send "Account Reactivated" email

### Failed Payment Handling

#### Common Failure Reasons

| Code                      | Reason             | Action          |
| ------------------------- | ------------------ | --------------- |
| `card_declined`           | Generic decline    | Update card     |
| `insufficient_funds`      | Not enough balance | Retry later     |
| `expired_card`            | Card expired       | Update card     |
| `authentication_required` | 3DS needed         | Customer action |

#### Smart Retry

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

### Collections (Enterprise)

For enterprise customers with large outstanding balances:

1. Account Manager notified at 14 days overdue
2. Finance team reviews at 21 days
3. Collections process at 30 days
4. Legal action consideration at 60 days

### Metrics

| Metric                          | Target |
| ------------------------------- | ------ |
| On-time payment rate            | >95%   |
| Dunning recovery rate           | >70%   |
| Suspension rate                 | <2%    |
| Cancellation rate (non-payment) | <0.5%  |

### Fraud Prevention

#### Red Flags

- Multiple failed payment methods
- Rapid usage spike after signup
- Disposable email domains
- High-risk countries

#### Actions

- Manual review before trial conversion
- Require upfront payment for high-risk
- Rate limit API key creation

<!-- Source: plan.md -->

## Technical Implementation Plan

### 1. Repository Structure

#### 1.1 Repositories

| ID  | Name                             | Type | Purpose                          |
| --- | -------------------------------- | ---- | -------------------------------- |
| R1  | temporal                         | FORK | Core server + cloud interceptors |
| R2  | cloud-api                        | FORK | Cloud API proto definitions      |
| R3  | tcld                             | FORK | Cloud CLI                        |
| R4  | terraform-provider-temporalcloud | FORK | Terraform provider               |
| R5  | temporal-cloud-platform          | NEW  | Backend services                 |
| R6  | temporal-cloud-console           | NEW  | Web UI                           |
| R7  | temporal-cloud-infra             | NEW  | Infrastructure as Code           |

#### 1.2 Branch Strategy

```
main              ← Protected, requires PR + review
├── cloud/main    ← Cloud production branch
├── cloud/staging ← Staging environment
├── cloud/develop ← Development integration
└── feature/*     ← Feature branches
```

### 2. Technology Stack

#### 2.1 Backend

| Component     | Technology         | Version   |
| ------------- | ------------------ | --------- |
| Language      | Go                 | 1.22+     |
| API Protocol  | gRPC + Connect     | -         |
| Database      | PostgreSQL         | 15        |
| Cache         | Redis              | 7         |
| Message Queue | Temporal Workflows | -         |
| Billing       | Stripe             | API v2024 |
| SAML          | crewjam/saml       | 0.4.x     |
| SCIM          | elimity-com/scim   | 2.x       |

#### 2.2 Frontend

| Component  | Technology     | Version |
| ---------- | -------------- | ------- |
| Framework  | Next.js        | 14      |
| UI Library | React          | 18      |
| Styling    | Tailwind CSS   | 3.4     |
| Components | shadcn/ui      | latest  |
| State      | TanStack Query | 5       |
| API Client | Connect-Web    | latest  |

#### 2.3 Infrastructure

| Component               | Technology           | Version |
| ----------------------- | -------------------- | ------- |
| IaC                     | Terraform            | 1.6+    |
| Container Orchestration | Kubernetes           | 1.28+   |
| Package Manager         | Helm                 | 3.13+   |
| CI/CD                   | GitHub Actions       | -       |
| Secrets                 | AWS Secrets Manager  | -       |
| Monitoring              | Prometheus + Grafana | -       |
| Logging                 | Loki                 | -       |
| Tracing                 | Jaeger               | -       |

### 3. Infrastructure Architecture

#### 3.1 Multi-Region Topology

```
                    ┌─────────────────────┐
                    │   Global Load       │
                    │   Balancer (DNS)    │
                    └──────────┬──────────┘
           ┌───────────────────┼───────────────────┐
           ▼                   ▼                   ▼
    ┌─────────────┐     ┌─────────────┐     ┌─────────────┐
    │  us-east-1  │     │  eu-west-1  │     │ ap-south-1  │
    │  (Primary)  │     │ (Secondary) │     │ (Secondary) │
    └─────────────┘     └─────────────┘     └─────────────┘
```

#### 3.2 Single Region Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                         VPC                                  │
│  ┌─────────────────────────────────────────────────────┐    │
│  │                 Public Subnets                       │    │
│  │  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐  │    │
│  │  │   ALB/NLB   │  │   ALB/NLB   │  │   ALB/NLB   │  │    │
│  │  │    AZ-a     │  │    AZ-b     │  │    AZ-c     │  │    │
│  │  └─────────────┘  └─────────────┘  └─────────────┘  │    │
│  └─────────────────────────────────────────────────────┘    │
│  ┌─────────────────────────────────────────────────────┐    │
│  │                 Private Subnets                      │    │
│  │  ┌─────────────────────────────────────────────┐    │    │
│  │  │              EKS Cluster                     │    │    │
│  │  │  ┌─────────┐ ┌─────────┐ ┌─────────────┐    │    │    │
│  │  │  │Temporal │ │ Cloud   │ │   Cloud     │    │    │    │
│  │  │  │ Server  │ │Platform │ │  Console    │    │    │    │
│  │  │  └─────────┘ └─────────┘ └─────────────┘    │    │    │
│  │  └─────────────────────────────────────────────┘    │    │
│  └─────────────────────────────────────────────────────┘    │
│  ┌─────────────────────────────────────────────────────┐    │
│  │                 Database Subnets                     │    │
│  │  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐  │    │
│  │  │  RDS (Pri)  │  │ RDS (Read)  │  │   Redis     │  │    │
│  │  │    AZ-a     │  │    AZ-b     │  │  Cluster    │  │    │
│  │  └─────────────┘  └─────────────┘  └─────────────┘  │    │
│  └─────────────────────────────────────────────────────┘    │
└─────────────────────────────────────────────────────────────┘
```

### 4. Implementation Phases

#### Phase 1: Foundation (Weeks 1-4)

- Fork repos, setup branches
- Create new repos (platform, console, infra)
- Setup sync automation
- Create DB schemas
- Create org/billing protos

#### Phase 2: Metering (Weeks 5-8)

- Metering interceptor
- Action collector
- Storage calculator
- Usage aggregation workflow

#### Phase 3: Billing (Weeks 9-14)

- Stripe integration
- Invoice generation
- Quota enforcement
- Payment processing

#### Phase 4: Security (Weeks 15-20)

- SAML SSO
- SCIM provisioning
- Audit logging
- API keys

#### Phase 5: Console (Weeks 16-22)

- Project setup
- Auth pages
- Organization pages
- Billing pages
- Settings pages

#### Phase 6: Infrastructure (Weeks 20-24)

- Staging environment
- Production environment
- DR setup
- Monitoring & alerting

#### Phase 7: Testing & Launch (Weeks 22-26)

- Complete test coverage
- Load testing
- Security audit
- Beta launch
- GA launch

<!-- Source: plugin-management.md -->

## Plugin Management

### Overview

Temporal allows plugins (interceptors, custom code). In Cloud, we must manage these strictly for security and stability.

### Supported Plugin Types

1. **Interceptors (Client-side)**

   - Run in customer's SDK.
   - **Management**: Distribute via SDK extensions (see `sdk-extensions.md`).
   - **Safety**: Safe (customer's risk).

2. **Remote Data Converter (Codec Server)**

   - Decrypts payloads for display in Console.
   - **Architecture**: Browser calls customer's HTTP endpoint directly.
   - **Management**: Configured per-namespace in Console settings.
   - **Safety**: Cloud never sees unencrypted data.

3. **Search Attributes (Server-side)**
   - Custom index fields.
   - **Limit**: 100 per namespace.
   - **Types**: Strict validation (Keyword, Int, etc.).

### Prohibited Plugins

- **Custom Workflow Logic (Server-side)**: No user code runs on Temporal servers.
- **Custom Interceptors (Server-side)**: Only standard Cloud interceptors allowed (Metering, Auth).

### Plugin Update Strategy

#### SDK Plugins

- Versioned via standard package managers (npm, maven, go mod).
- Deprecation warnings in SDK logs.

#### Codec Server

- Customer responsible for maintaining availability.
- Console handles connection errors gracefully ("Unable to decode payload").

### Internal Plugins (Cloud Platform)

We use internal plugins for extensibility:

```go
// Platform plugin interface
type CloudPlugin interface {
    OnNamespaceCreated(ctx context.Context, ns *Namespace) error
    OnNamespaceDeleted(ctx context.Context, ns *Namespace) error
}

// Implementation example: Datadog Metrics
type DatadogPlugin struct {}

func (p *DatadogPlugin) OnNamespaceCreated(ctx context.Context, ns *Namespace) error {
    // Register new tag in Datadog
    return nil
}
```

Managed via dependency injection at startup.

<!-- Source: pricing.md -->

## Pricing & Plans

### Plan Tiers

| Plan             | Base Price | Actions Included | Active Storage | Retained Storage |
| ---------------- | ---------- | ---------------- | -------------- | ---------------- |
| Free             | $0/mo      | 100K             | 0.1 GB         | 4 GB             |
| Essential        | $100/mo    | 1M               | 1 GB           | 40 GB            |
| Business         | $500/mo    | 2.5M             | 2.5 GB         | 100 GB           |
| Enterprise       | Custom     | 10M              | 10 GB          | 400 GB           |
| Mission Critical | Custom     | Custom           | Custom         | Custom           |

### Overage Pricing

#### Actions (per million)

| Tier    | Price |
| ------- | ----- |
| 0-5M    | $50   |
| 5-10M   | $45   |
| 10-20M  | $40   |
| 20-50M  | $35   |
| 50-100M | $30   |
| 100M+   | $25   |

#### Storage

| Type             | Price        |
| ---------------- | ------------ |
| Active Storage   | $0.042/GBh   |
| Retained Storage | $0.00105/GBh |

#### Storage Conversion

- 1 GB = 744 GBh (per month)
- Active: $0.042 × 744 = ~$31.25/GB/month
- Retained: $0.00105 × 744 = ~$0.78/GB/month

### Add-ons

| Feature           | Price    | Included In                      |
| ----------------- | -------- | -------------------------------- |
| SAML SSO          | Included | Business+                        |
| SCIM              | $500/mo  | Enterprise+ (or Business add-on) |
| Multi-region HA   | 2x usage | All plans                        |
| Dedicated Support | Custom   | Enterprise+                      |

### Billing Cycle

- Monthly billing on the 1st
- Prorated for partial months
- Usage calculated in UTC
- Invoices due NET 30

### Payment Methods

- Credit card (Stripe)
- ACH bank transfer (Enterprise)
- Wire transfer (Enterprise)
- Temporal Credits (prepaid)

### Credits

#### Commitment Pricing

| Commitment | Discount |
| ---------- | -------- |
| 1 year     | 10%      |
| 2 years    | 15%      |
| 3 years    | 20%      |

#### Credit Application

- Credits apply to: Actions, Storage, Plan fees
- Credits do NOT apply to: Add-ons
- Credits expire after commitment period

### Example Calculations

#### Small Workload (Essential Plan)

```
Base plan:                    $100
Actions: 2M (1M overage)      $50
Active Storage: 0.5 GBh       $0 (included)
Retained Storage: 20 GBh      $0 (included)
─────────────────────────────────
Total:                        $150/month
```

#### Medium Workload (Business Plan)

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

<!-- Source: qa-process.md -->

## QA Process

### QA Environments

| Environment | Purpose             | Data                | Access         |
| ----------- | ------------------- | ------------------- | -------------- |
| Dev         | Developer testing   | Seed data           | All engineers  |
| QA          | QA team testing     | Sanitized prod copy | QA + Engineers |
| Staging     | Pre-prod validation | Sanitized prod copy | All            |
| Prod        | Production          | Real                | Restricted     |

### QA Workflow

#### Feature Development

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Dev       │────▶│   QA        │────▶│   Staging   │
│   Testing   │     │   Testing   │     │   Testing   │
└─────────────┘     └─────────────┘     └─────────────┘
     │                    │                    │
     ▼                    ▼                    ▼
  Unit Tests         Functional          E2E + Load
  Integration        Exploratory         UAT
```

#### QA Entry Criteria

Before QA testing begins:

- [ ] Unit tests passing (>80% coverage)
- [ ] Integration tests passing
- [ ] Code review approved
- [ ] Feature deployed to QA environment
- [ ] Test data prepared
- [ ] Acceptance criteria documented

#### QA Exit Criteria

Before promotion to staging:

- [ ] All test cases passed
- [ ] No P0/P1 bugs open
- [ ] Performance benchmarks met
- [ ] Security checklist completed
- [ ] Documentation reviewed

### Test Types

#### 1. Functional Testing

Manual and automated testing of features.

**Test Case Template**:

```markdown
## TC-001: Create Namespace

**Preconditions**:

- User logged in with admin role
- No namespace with name "test-ns" exists

**Steps**:

1. Navigate to Namespaces
2. Click "Create Namespace"
3. Enter name: "test-ns"
4. Select region: "aws-us-east-1"
5. Click "Create"

**Expected Result**:

- Namespace created successfully
- Appears in namespace list
- Correct region shown

**Priority**: High
**Automated**: Yes
```

#### 2. Exploratory Testing

Unscripted testing to find edge cases.

**Session Template**:

```markdown
## Exploratory Session: Billing Flow

**Charter**: Explore billing edge cases around plan upgrades mid-cycle

**Time Box**: 2 hours

**Areas to Explore**:

- Upgrade with pending invoice
- Upgrade then immediate downgrade
- Upgrade with low credit balance

**Findings**:

- [Bug] Proration calculation off by one day
- [Observation] Slow response when loading invoices
```

#### 3. Regression Testing

Ensure existing features still work after changes.

**Regression Suite**:

- Core Workflows: 50 cases (automated)
- Billing: 30 cases (automated)
- Authentication: 25 cases (automated)
- Console UI: 100 cases (80% automated)

#### 4. Performance Testing

```bash
# Run load test
k6 run --vus 100 --duration 10m load/api-test.js

# Expected Results
# - P99 latency < 200ms
# - Error rate < 0.1%
# - No memory leaks
```

#### 5. Security Testing

- **SAST**: Static analysis on every PR (Semgrep)
- **DAST**: Weekly dynamic scanning (OWASP ZAP)
- **Dependency Scan**: On every build (Snyk)
- **Penetration Test**: Annual (third party)

### Test Data Management

#### Seed Data

```sql
-- QA seed data
INSERT INTO organizations (id, name, slug) VALUES
  ('qa-org-1', 'QA Test Org 1', 'qa-test-1'),
  ('qa-org-2', 'QA Test Org 2', 'qa-test-2');

INSERT INTO users (id, email, name) VALUES
  ('qa-user-1', 'qa1@test.temporal.io', 'QA User 1'),
  ('qa-user-2', 'qa2@test.temporal.io', 'QA User 2');
```

#### Data Sanitization

For staging/QA from production:

```bash
# Sanitize production data
pg_dump prod_db | \
  sed 's/[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}/redacted@example.com/g' | \
  psql qa_db
```

### Bug Reporting

#### From QA Testing

```markdown
## Bug: [Short Title]

**Found In**: QA Environment
**Build**: v1.2.0-rc1
**Test Case**: TC-042

**Steps to Reproduce**:

1. ...

**Expected**: ...
**Actual**: ...

**Attachments**:

- Screenshot
- Console logs
- Network trace

**Severity**: P2
**Assigned To**: @engineer
```

### QA Metrics

#### Quality Metrics

| Metric               | Target | Measurement                  |
| -------------------- | ------ | ---------------------------- |
| Test Coverage        | >80%   | Codecov                      |
| Bug Escape Rate      | <5%    | Bugs in prod / total bugs    |
| Automation Rate      | >70%   | Automated / total test cases |
| Regression Pass Rate | >99%   | Passing / total regression   |

#### Release Quality Gate

Release blocked if:

- Test coverage < 80%
- Any P0/P1 bugs open
- Security scan has critical findings
- Performance regression > 10%

### Tools

| Purpose           | Tool              |
| ----------------- | ----------------- |
| Test Management   | TestRail / Linear |
| Automated Testing | Playwright, k6    |
| Bug Tracking      | Linear            |
| Security Scanning | Snyk, Semgrep     |
| Performance       | k6, Grafana       |

<!-- Source: quickstart.md -->

## Quickstart Guide

### Prerequisites

```bash
# Install required tools
brew install go node@20 docker kubectl helm terraform
brew install bufbuild/buf/buf golangci-lint gh
npm install -g pnpm

# Verify versions
go version      # 1.22+
node --version  # 20.x
docker --version
kubectl version --client
helm version
terraform version
```

### 1. Clone Repositories

```bash
mkdir -p ~/temporal-cloud && cd ~/temporal-cloud

# Clone all repos
gh repo clone YOUR_ORG/temporal
gh repo clone YOUR_ORG/cloud-api
gh repo clone YOUR_ORG/temporal-cloud-platform
gh repo clone YOUR_ORG/temporal-cloud-console
gh repo clone YOUR_ORG/temporal-cloud-infra

# Setup upstream remotes for forks
cd temporal && git remote add upstream https://github.com/temporalio/temporal.git
cd ../cloud-api && git remote add upstream https://github.com/temporalio/cloud-api.git
```

### 2. Start Local Environment

#### Docker Compose

```yaml
# docker-compose.dev.yaml
version: "3.8"

services:
  postgresql:
    image: postgres:15
    environment:
      POSTGRES_USER: temporal
      POSTGRES_PASSWORD: temporal
      POSTGRES_DB: temporal_cloud
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

  redis:
    image: redis:7
    ports:
      - "6379:6379"

  temporal:
    image: temporalio/auto-setup:latest
    ports:
      - "7233:7233"
    environment:
      - DB=postgresql
      - DB_PORT=5432
      - POSTGRES_USER=temporal
      - POSTGRES_PWD=temporal
      - POSTGRES_SEEDS=postgresql
    depends_on:
      - postgresql

  temporal-ui:
    image: temporalio/ui:latest
    ports:
      - "8080:8080"
    environment:
      - TEMPORAL_ADDRESS=temporal:7233
    depends_on:
      - temporal

volumes:
  postgres_data:
```

```bash
# Start infrastructure
docker-compose -f docker-compose.dev.yaml up -d

# Verify services
docker-compose ps
```

### 3. Run Migrations

```bash
cd temporal-cloud-platform

# Install golang-migrate
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Run migrations
migrate -path schema -database "postgres://temporal:temporal@localhost:5432/temporal_cloud?sslmode=disable" up
```

### 4. Start Backend

```bash
cd temporal-cloud-platform

# Install dependencies
go mod download

# Run the server
go run cmd/server/main.go
```

### 5. Start Frontend

```bash
cd temporal-cloud-console

# Install dependencies
pnpm install

# Start dev server
pnpm dev
```

### 6. Access Services

| Service       | URL                   |
| ------------- | --------------------- |
| Temporal UI   | http://localhost:8080 |
| Cloud Console | http://localhost:3000 |
| Cloud API     | http://localhost:8081 |
| PostgreSQL    | localhost:5432        |
| Redis         | localhost:6379        |

### 7. Run Tests

```bash
# Backend tests
cd temporal-cloud-platform
make test

# Frontend tests
cd temporal-cloud-console
pnpm test

# Integration tests
make integration-test
```

### 8. Generate Code

```bash
# Generate proto code
cd cloud-api
buf generate

# Generate TypeScript types
cd temporal-cloud-console
pnpm generate
```

### Common Issues

#### Port Already in Use

```bash
# Find process using port
lsof -i :5432

# Kill process
kill -9 <PID>
```

#### Database Connection Failed

```bash
# Check PostgreSQL is running
docker-compose ps postgresql

# Check logs
docker-compose logs postgresql
```

#### Proto Generation Failed

```bash
# Ensure buf is installed
buf --version

# Clear cache and retry
buf mod update
buf generate
```

<!-- Source: regions.md -->

## Supported Regions

### AWS Regions

| Region      | Code           | Cloud API Code     | PrivateLink | Same-Region HA | Multi-Region HA |
| ----------- | -------------- | ------------------ | ----------- | -------------- | --------------- |
| N. Virginia | us-east-1      | aws-us-east-1      | ✅          | ✅             | ✅              |
| Ohio        | us-east-2      | aws-us-east-2      | ✅          | ❌             | ✅              |
| Oregon      | us-west-2      | aws-us-west-2      | ✅          | ✅             | ✅              |
| Canada      | ca-central-1   | aws-ca-central-1   | ✅          | ❌             | ✅              |
| Ireland     | eu-west-1      | aws-eu-west-1      | ✅          | ✅             | ✅              |
| London      | eu-west-2      | aws-eu-west-2      | ✅          | ❌             | ✅              |
| Frankfurt   | eu-central-1   | aws-eu-central-1   | ✅          | ✅             | ✅              |
| Singapore   | ap-southeast-1 | aws-ap-southeast-1 | ✅          | ❌             | ✅              |
| Sydney      | ap-southeast-2 | aws-ap-southeast-2 | ✅          | ❌             | ✅              |
| Tokyo       | ap-northeast-1 | aws-ap-northeast-1 | ✅          | ❌             | ✅              |
| Seoul       | ap-northeast-2 | aws-ap-northeast-2 | ✅          | ❌             | ✅              |
| Mumbai      | ap-south-1     | aws-ap-south-1     | ✅          | ❌             | ✅              |
| Hyderabad   | ap-south-2     | aws-ap-south-2     | ✅          | ❌             | ✅              |
| São Paulo   | sa-east-1      | aws-sa-east-1      | ✅          | ❌             | ❌              |

### GCP Regions

| Region      | Code         | Cloud API Code   | PSC | Same-Region HA | Multi-Region HA |
| ----------- | ------------ | ---------------- | --- | -------------- | --------------- |
| Iowa        | us-central1  | gcp-us-central1  | ✅  | ❌             | ✅              |
| Oregon      | us-west1     | gcp-us-west1     | ✅  | ❌             | ✅              |
| N. Virginia | us-east4     | gcp-us-east4     | ✅  | ❌             | ✅              |
| Frankfurt   | europe-west3 | gcp-europe-west3 | ✅  | ❌             | ✅              |
| Mumbai      | asia-south1  | gcp-asia-south1  | ✅  | ❌             | ✅              |

### Multi-Region Replication Pairs

#### US Regions

| Primary         | Secondary Options                              |
| --------------- | ---------------------------------------------- |
| aws-us-east-1   | aws-us-east-2, aws-us-west-2, aws-ca-central-1 |
| aws-us-west-2   | aws-us-east-1, aws-us-east-2, aws-ca-central-1 |
| gcp-us-central1 | gcp-us-west1, gcp-us-east4                     |

#### EU Regions

| Primary          | Secondary Options               |
| ---------------- | ------------------------------- |
| aws-eu-west-1    | aws-eu-west-2, aws-eu-central-1 |
| aws-eu-central-1 | aws-eu-west-1, aws-eu-west-2    |

#### APAC Regions

| Primary            | Secondary Options                                          |
| ------------------ | ---------------------------------------------------------- |
| aws-ap-northeast-1 | aws-ap-northeast-2, aws-ap-southeast-1, aws-ap-southeast-2 |
| aws-ap-south-1     | aws-ap-south-2, aws-ap-southeast-1                         |

### Multi-Cloud Replication

| AWS Region       | GCP Region                    |
| ---------------- | ----------------------------- |
| aws-us-east-1    | gcp-us-central1, gcp-us-east4 |
| aws-us-west-2    | gcp-us-west1                  |
| aws-eu-central-1 | gcp-europe-west3              |
| aws-ap-south-1   | gcp-asia-south1               |

### Regional Endpoints

#### Format

```
<cloud-api-code>.region.tmprl.cloud
```

#### Examples

```
aws-us-east-1.region.tmprl.cloud
gcp-us-central1.region.tmprl.cloud
```

### Latency Considerations

- Choose region closest to your workers
- Multi-region adds ~50-100ms latency for replication
- Cross-cloud replication may have higher latency

<!-- Source: release-management.md -->

## Release Management

### Versioning

#### Semantic Versioning

`MAJOR.MINOR.PATCH`

- **MAJOR**: Breaking API changes
- **MINOR**: New features, backwards compatible
- **PATCH**: Bug fixes, security patches

#### Version Components

| Component          | Version | Release Cadence |
| ------------------ | ------- | --------------- |
| Cloud Platform API | v1.x.x  | Monthly         |
| Cloud Console      | v1.x.x  | Weekly          |
| tcld CLI           | v1.x.x  | Monthly         |
| Terraform Provider | v0.x.x  | Monthly         |
| SDK Extensions     | v1.x.x  | As needed       |

### Release Process

#### 1. Feature Freeze (T-5 days)

```bash
# Create release branch
git checkout cloud/develop
git checkout -b release/v1.2.0
git push origin release/v1.2.0
```

- No new features after this point
- Only bug fixes and polish

#### 2. QA Validation (T-4 to T-2 days)

- Deploy release branch to staging
- Run full E2E test suite
- QA team performs manual testing
- Security scan

#### 3. Release Candidate (T-2 days)

```bash
# Tag RC
git tag v1.2.0-rc1
git push origin v1.2.0-rc1
```

- Deploy RC to production (canary)
- Monitor metrics for 24 hours

#### 4. Release (T-0)

```bash
# Merge to cloud/main
git checkout cloud/main
git merge release/v1.2.0
git tag v1.2.0
git push origin cloud/main --tags

# Update changelog
# Generate release notes
```

#### 5. Post-Release

- Announce in #releases Slack
- Update documentation
- Email customers (major releases only)
- Monitor for issues

### Changelog

#### Format (Keep a Changelog)

```markdown
# Changelog

## [1.2.0] - 2025-02-01

### Added

- Multi-region namespace support (#123)
- SCIM group sync (#456)

### Changed

- Improved billing dashboard performance

### Fixed

- Certificate expiry notification timing (#789)

### Security

- Updated dependencies for CVE-2025-1234
```

#### Generation

```bash
# Auto-generate from commits
git log v1.1.0..v1.2.0 --pretty=format:"- %s (%h)" > CHANGELOG_DRAFT.md
```

### Rollout Strategy

#### Canary Deployment

1. Deploy to 5% of traffic
2. Monitor error rates, latency for 1 hour
3. If healthy, increase to 25%
4. Wait 2 hours
5. Full rollout (100%)

```yaml
# ArgoCD rollout
apiVersion: argoproj.io/v1alpha1
kind: Rollout
spec:
  strategy:
    canary:
      steps:
        - setWeight: 5
        - pause: { duration: 1h }
        - setWeight: 25
        - pause: { duration: 2h }
        - setWeight: 100
      analysis:
        templates:
          - templateName: success-rate
        startingStep: 1
```

#### Rollback Criteria

Automatic rollback if:

- Error rate > 1%
- P99 latency > 500ms
- Any P0/P1 bugs reported

```bash
# Manual rollback
helm rollback cloud-platform --namespace cloud-platform
```

### Hotfix Process

For critical production issues:

```bash
# Branch from cloud/main
git checkout cloud/main
git checkout -b hotfix/v1.2.1

# Fix the issue
# ... commit ...

# Tag and deploy immediately
git tag v1.2.1
git push origin v1.2.1

# Merge back
git checkout cloud/main && git merge hotfix/v1.2.1
git checkout cloud/develop && git merge hotfix/v1.2.1
```

### Release Checklist

#### Pre-Release

- [ ] All tests passing
- [ ] Security scan clean
- [ ] Changelog updated
- [ ] Documentation updated
- [ ] Breaking changes documented
- [ ] Migration guide (if needed)
- [ ] Rollback plan confirmed

#### Release

- [ ] Tag created
- [ ] Canary deployed
- [ ] Metrics monitored
- [ ] Full rollout complete

#### Post-Release

- [ ] Release notes published
- [ ] Customers notified (if applicable)
- [ ] Retrospective scheduled (major releases)

### Communication

#### Internal

- Slack #releases: All releases
- Slack #incidents: Hotfixes

#### External

- Status page: Maintenance notices
- Email: Major releases, breaking changes
- Blog: Feature announcements

<!-- Source: repo-automation.md -->

## Automated Repository Sync & Management

### Overview

We must maintain synchronization with 190+ repositories in `temporalio` organization, including automatically discovering and syncing new ones.

### Architecture

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│  GitHub     │────▶│   Sync      │────▶│   Our       │
│  API        │     │   Bot       │     │   Org       │
└─────────────┘     └─────────────┘     └─────────────┘
```

### Automated Discovery

#### Sync Bot Job

Runs daily (or triggered by webhook):

```go
func DiscoverAndSync(ctx context.Context) error {
    // 1. List all upstream repos
    upstreamRepos, err := github.ListOrgRepos("temporalio")

    // 2. List our repos
    ourRepos, err := github.ListOrgRepos("YOUR_ORG")

    // 3. Find new repos
    for _, upstream := range upstreamRepos {
        if !contains(ourRepos, upstream.Name) {
            // New repo found!
            handleNewRepo(ctx, upstream)
        } else {
            // Existing repo - check for updates
            syncRepo(ctx, upstream)
        }
    }

    return nil
}
```

#### Handling New Repos

When a new repo appears in `temporalio`:

1. **Classify**: Determine repo type (Go SDK, Java SDK, Core, Tool, etc.) based on languages and topics.
2. **Fork/Mirror**:
   - If it's a core component we modify -> **Fork**
   - If it's a dependency/SDK -> **Mirror**
3. **Configure**: Apply standard branch protection, CI workflows, and team access.
4. **Notify**: Alert #engineering-ops about the new repo.

### Sync Strategy

#### 1. Mirror Repos (Read-Only)

For repos we don't modify (e.g. `sdk-java`, `samples-go`):

```yaml
# .github/workflows/mirror-sync.yaml
name: Mirror Sync
on:
  schedule:
    - cron: "0 */4 * * *" # Every 4 hours

jobs:
  sync:
    runs-on: ubuntu-latest
    steps:
      - name: Sync
        run: |
          git remote add upstream https://github.com/temporalio/${REPO_NAME}.git
          git fetch upstream
          git checkout main
          git merge upstream/main --ff-only
          git push origin main
          # Sync tags too
          git fetch upstream --tags
          git push origin --tags
```

#### 2. Forked Repos (Modified)

For repos we patch (e.g. `temporal`, `ui`):

```bash
# Logic:
# 1. Fetch upstream/main
# 2. Rebase cloud/main onto upstream/main? NO.
# 3. Merge upstream/main into cloud/main? YES.

git checkout cloud/main
git merge upstream/main -m "Sync upstream"

if conflict:
    create_pr_for_manual_resolution()
else:
    push origin cloud/main
    trigger_build()
```

### Release Synchronization

#### Automatic Release Detection

When `temporalio/temporal` tags `v1.25.0`:

1. **Detect**: Webhook or polling sees new tag.
2. **Branch**: Create `release/v1.25.0` from `upstream/v1.25.0`.
3. **Apply Patches**: Cherry-pick our cloud-specific commits (if structured as patches) OR merge `cloud/main` features.
4. **Build**: Trigger release build.
5. **Test**: Run integration tests.
6. **Promote**: If tests pass, mark as `cloud-v1.25.0` ready for staging.

### Conflict Management

#### Prevention

- **Strict Isolation**: Cloud code goes in `common/cloud/` directory.
- **Interfaces**: Use Go interfaces to inject cloud logic, avoiding changes to core files.
- **Interceptors**: Use gRPC interceptors instead of modifying handlers.

#### Resolution

If auto-sync fails:

1. **Alert**: "Sync failed for `temporal`: Merge conflict in `service/history/workflow.go`"
2. **PR**: Bot creates a PR with the conflict markers.
3. **Block**: Deployments blocked until resolved.

### Tooling

#### Repo Config as Code

Manage settings for all 190+ repos centrally:

```yaml
# repos.yaml
defaults:
  has_issues: false
  has_projects: false
  has_wiki: false
  allow_squash_merge: true
  allow_merge_commit: false
  allow_rebase_merge: false

repos:
  temporal:
    has_issues: true
    protected_branches:
      main:
        required_reviews: 1
      cloud/main:
        required_reviews: 2
```

Run `github-settings-sync` tool to apply.

### Metrics

- **Sync Latency**: Time from upstream commit to our repo (Target: < 1h)
- **Sync Failures**: Number of manual interventions required (Target: < 1/week)
- **Repo Coverage**: % of upstream repos mirrored (Target: 100%)

<!-- Source: runbooks.md -->

## Runbooks

### Runbook Index

| Runbook                 | Purpose                     | Location             |
| ----------------------- | --------------------------- | -------------------- |
| incident-response.md    | Handle production incidents | infra/docs/runbooks/ |
| disaster-recovery.md    | DR procedures               | infra/docs/runbooks/ |
| database-maintenance.md | DB operations               | infra/docs/runbooks/ |
| scaling.md              | Scale services              | infra/docs/runbooks/ |
| certificate-rotation.md | Rotate TLS certs            | infra/docs/runbooks/ |
| secret-rotation.md      | Rotate secrets              | infra/docs/runbooks/ |
| deployment.md           | Deploy new versions         | infra/docs/runbooks/ |
| rollback.md             | Rollback procedures         | infra/docs/runbooks/ |

---

### Incident Response

#### Severity Classification

| Level | Description      | Response Time     |
| ----- | ---------------- | ----------------- |
| SEV1  | Complete outage  | 15 minutes        |
| SEV2  | Partial outage   | 1 hour            |
| SEV3  | Degraded service | 4 hours           |
| SEV4  | Minor issue      | Next business day |

#### Response Process

1. **Acknowledge** - Claim incident in PagerDuty
2. **Assess** - Determine severity and impact
3. **Communicate** - Update status page, notify stakeholders
4. **Mitigate** - Apply temporary fix if needed
5. **Resolve** - Implement permanent fix
6. **Review** - Post-incident review within 48 hours

---

### Database Maintenance

#### Backup Verification

```bash
# List recent backups
aws rds describe-db-snapshots \
  --db-instance-identifier temporal-cloud-prod \
  --query 'DBSnapshots[*].[DBSnapshotIdentifier,SnapshotCreateTime]'

# Restore to test instance
aws rds restore-db-instance-from-db-snapshot \
  --db-instance-identifier temporal-cloud-restore-test \
  --db-snapshot-identifier <snapshot-id>
```

#### Point-in-Time Recovery

```bash
# Restore to specific time
aws rds restore-db-instance-to-point-in-time \
  --source-db-instance-identifier temporal-cloud-prod \
  --target-db-instance-identifier temporal-cloud-pitr \
  --restore-time 2025-01-01T12:00:00Z
```

---

### Scaling

#### Horizontal Scaling (Pods)

```bash
# Scale deployment
kubectl scale deployment cloud-api \
  --replicas=5 \
  -n cloud-platform

# Verify
kubectl get pods -n cloud-platform
```

#### Vertical Scaling (Resources)

```bash
# Edit deployment
kubectl edit deployment cloud-api -n cloud-platform

# Update resources
resources:
  requests:
    cpu: "2"
    memory: "4Gi"
  limits:
    cpu: "4"
    memory: "8Gi"
```

#### Database Scaling

```bash
# Modify RDS instance
aws rds modify-db-instance \
  --db-instance-identifier temporal-cloud-prod \
  --db-instance-class db.r6g.2xlarge \
  --apply-immediately
```

---

### Certificate Rotation

#### Check Expiration

```bash
# Check certificate expiration
openssl x509 -enddate -noout -in /path/to/cert.pem

# List all certs expiring in 30 days
kubectl get certificates -A -o json | jq '.items[] | select(.status.notAfter | fromdateiso8601 < (now + 2592000))'
```

#### Rotate Certificate

```bash
# Trigger cert-manager renewal
kubectl delete secret <cert-secret-name> -n <namespace>

# Verify new certificate
kubectl get certificate <cert-name> -n <namespace>
```

---

### Deployment

#### Standard Deployment

```bash
# Deploy to staging
helm upgrade --install cloud-platform ./charts/cloud-platform \
  --namespace cloud-platform \
  --set image.tag=v1.2.3 \
  -f values-staging.yaml

# Verify deployment
kubectl rollout status deployment/cloud-api -n cloud-platform
```

#### Canary Deployment

```bash
# Deploy canary (10%)
helm upgrade --install cloud-platform-canary ./charts/cloud-platform \
  --namespace cloud-platform \
  --set image.tag=v1.2.3 \
  --set replicaCount=1

# Monitor metrics
# If healthy, scale up canary, scale down stable
```

---

### Rollback

#### Helm Rollback

```bash
# List releases
helm history cloud-platform -n cloud-platform

# Rollback to previous
helm rollback cloud-platform 1 -n cloud-platform

# Rollback to specific revision
helm rollback cloud-platform 5 -n cloud-platform
```

#### Kubernetes Rollback

```bash
# Rollback deployment
kubectl rollout undo deployment/cloud-api -n cloud-platform

# Rollback to specific revision
kubectl rollout undo deployment/cloud-api --to-revision=2 -n cloud-platform
```

---

### Secret Rotation

#### Database Credentials

```bash
# Generate new password
NEW_PASSWORD=$(openssl rand -base64 32)

# Update in Secrets Manager
aws secretsmanager update-secret \
  --secret-id temporal-cloud/db-password \
  --secret-string "$NEW_PASSWORD"

# Update RDS
aws rds modify-db-instance \
  --db-instance-identifier temporal-cloud-prod \
  --master-user-password "$NEW_PASSWORD"

# Restart pods to pick up new secret
kubectl rollout restart deployment/cloud-api -n cloud-platform
```

#### API Keys

```bash
# Rotate via API
curl -X POST https://api.temporal-cloud.io/v1/api-keys/rotate \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -d '{"key_id": "key-123"}'
```

<!-- Source: saml.md -->

## SAML SSO

### Overview

SAML 2.0 Single Sign-On allows users to authenticate via their organization's Identity Provider (IdP).

### Supported IdPs

- Microsoft Entra ID (Azure AD)
- Okta
- OneLogin
- Any SAML 2.0 compliant IdP

### Configuration

#### Step 1: Create SAML Application in IdP

##### Okta

1. Create new SAML 2.0 application
2. Set ACS URL: `https://login.tmprl.cloud/saml/acs/<account-id>`
3. Set Entity ID: `urn:auth0:prod-tmprl:saml-<account-id>`
4. Configure attribute mappings

##### Azure AD

1. Create Enterprise Application
2. Set up SAML SSO
3. Set Reply URL: `https://login.tmprl.cloud/saml/acs/<account-id>`
4. Set Identifier: `urn:auth0:prod-tmprl:saml-<account-id>`

#### Step 2: Configure Temporal Cloud

```bash
# Via tcld
tcld account saml configure \
  --idp-metadata-url "https://your-idp.com/metadata.xml"

# Or upload metadata file
tcld account saml configure \
  --idp-metadata-file ./idp-metadata.xml
```

#### Step 3: Enable SAML

```bash
tcld account saml enable
```

### Attribute Mapping

| SAML Attribute | Temporal Attribute | Required |
| -------------- | ------------------ | -------- |
| email          | User email         | Yes      |
| firstName      | First name         | No       |
| lastName       | Last name          | No       |
| groups         | Role mapping       | No       |

### Group-to-Role Mapping

Map IdP groups to Temporal Cloud roles:

```json
{
  "group_mappings": [
    {
      "idp_group": "temporal-admins",
      "account_role": "global_admin"
    },
    {
      "idp_group": "temporal-developers",
      "account_role": "developer"
    },
    {
      "idp_group": "temporal-readonly",
      "account_role": "read_only"
    }
  ]
}
```

### Just-in-Time Provisioning

When enabled, users are automatically created on first login:

- User created with mapped role
- Email from SAML assertion
- Name from SAML attributes (if provided)

```bash
tcld account saml configure \
  --jit-provisioning enabled \
  --default-role developer
```

### Pricing

| Plan             | SAML SSO     |
| ---------------- | ------------ |
| Essential        | Not included |
| Business         | Included     |
| Enterprise       | Included     |
| Mission Critical | Included     |

### Troubleshooting

#### Login Failed

1. Verify IdP metadata is correct
2. Check ACS URL matches exactly
3. Verify Entity ID matches
4. Check SAML assertion contains email

#### User Not Created

1. Verify JIT provisioning is enabled
2. Check email attribute is mapped
3. Verify user limit not reached

#### Role Not Assigned

1. Check group mappings configuration
2. Verify groups attribute is sent
3. Check group name matches exactly

### Security Considerations

1. **Require SAML**: Disable password login after SAML setup
2. **Session timeout**: Configure appropriate session duration
3. **Group sync**: Use SCIM for real-time group sync
4. **Audit**: Monitor SAML login events

<!-- Source: scim.md -->

## SCIM Provisioning

### Overview

System for Cross-domain Identity Management (SCIM 2.0) enables automatic user and group provisioning from your Identity Provider.

### Requirements

- SAML SSO must be configured first
- Plan: Enterprise/Mission Critical, or Business + $500/mo add-on

### Supported IdPs

- Okta
- Microsoft Entra ID (Azure AD)
- OneLogin
- Any SCIM 2.0 compliant IdP

### Endpoints

| Endpoint                         | Methods                  | Description        |
| -------------------------------- | ------------------------ | ------------------ |
| `/scim/v2/Users`                 | GET, POST, PATCH, DELETE | User management    |
| `/scim/v2/Groups`                | GET, POST, PATCH, DELETE | Group management   |
| `/scim/v2/ServiceProviderConfig` | GET                      | SCIM configuration |
| `/scim/v2/Schemas`               | GET                      | Schema definitions |

### User Attributes

| SCIM Attribute  | Temporal Attribute | Required |
| --------------- | ------------------ | -------- |
| userName        | email              | Yes      |
| displayName     | name               | No       |
| active          | enabled            | Yes      |
| emails[primary] | email              | Yes      |

### Group Attributes

| SCIM Attribute | Temporal Attribute |
| -------------- | ------------------ |
| displayName    | Group name         |
| members        | Group members      |

### Configuration

#### Step 1: Enable SCIM

```bash
tcld account scim enable
```

#### Step 2: Generate SCIM Token

```bash
tcld account scim token create
# Output: SCIM bearer token (shown only once)
```

#### Step 3: Configure IdP

##### Okta

1. Go to Applications → Your SAML App → Provisioning
2. Enable SCIM provisioning
3. Set Base URL: `https://api.temporal.io/scim/v2`
4. Set API Token: `<scim-token>`
5. Enable: Create Users, Update User Attributes, Deactivate Users

##### Azure AD

1. Go to Enterprise App → Provisioning
2. Set Provisioning Mode: Automatic
3. Set Tenant URL: `https://api.temporal.io/scim/v2`
4. Set Secret Token: `<scim-token>`
5. Test connection and save

### Provisioning Behavior

#### User Created in IdP

1. SCIM POST to `/Users`
2. User created in Temporal Cloud
3. Default role assigned (or from group mapping)

#### User Updated in IdP

1. SCIM PATCH to `/Users/{id}`
2. User attributes updated
3. Role unchanged (unless group changed)

#### User Deactivated in IdP

1. SCIM PATCH with `active: false`
2. User disabled in Temporal Cloud
3. Active sessions terminated

#### User Deleted in IdP

1. SCIM DELETE to `/Users/{id}`
2. User removed from Temporal Cloud
3. API keys revoked

### Group Sync

#### Group Created

1. SCIM POST to `/Groups`
2. User group created in Temporal Cloud

#### Members Added

1. SCIM PATCH to `/Groups/{id}`
2. Users added to group
3. Namespace permissions applied

#### Members Removed

1. SCIM PATCH to `/Groups/{id}`
2. Users removed from group
3. Namespace permissions revoked

### Namespace Permissions via Groups

Map IdP groups to namespace permissions:

```bash
tcld user-group namespace-access set \
  --group-id grp-123 \
  --namespace my-namespace \
  --permission write
```

### Pricing

| Plan             | SCIM           |
| ---------------- | -------------- |
| Essential        | Not available  |
| Business         | $500/mo add-on |
| Enterprise       | Included       |
| Mission Critical | Included       |

### Troubleshooting

#### Provisioning Failed

1. Check SCIM token is valid
2. Verify Base URL is correct
3. Check IdP logs for errors
4. Verify user limit not reached

#### User Not Synced

1. Check user is assigned to app in IdP
2. Verify required attributes are mapped
3. Check provisioning logs in IdP

#### Group Permissions Not Applied

1. Verify group exists in Temporal Cloud
2. Check namespace permissions are configured
3. Verify user is member of group

<!-- Source: sdk-extensions.md -->

## SDK Extensions

### Overview

While standard Temporal SDKs work with Temporal Cloud, we provide wrapper libraries to improve Developer Experience (DX), specifically for certificate rotation and cloud-specific metrics.

### Go SDK Extension (`go.temporal.io/sdk/contrib/cloud`)

#### 1. Automatic Certificate Rotation

Standard SDK requires restarting the client when certificates change. This extension watches the file system and hot-reloads TLS config.

```go
options := client.Options{
    HostPort: "my-ns.tmprl.cloud:443",
    Namespace: "my-ns",
    ConnectionOptions: client.ConnectionOptions{
        TLS: cloud.NewRotatingTLSConfig(cloud.TLSParams{
            CertPath: "/etc/certs/client.pem",
            KeyPath:  "/etc/certs/client.key",
            CheckInterval: 1 * time.Minute,
        }),
    },
}
```

#### 2. Cloud Metrics Tagger

Automatically adds `namespace` and `region` tags to all metrics and converts them to OpenTelemetry format preferred by our observability stack.

```go
options := client.Options{
    MetricsHandler: cloud.NewMetricsHandler(cloud.MetricsParams{
        Prometheus: true,
    }),
}
```

#### 3. Connection Tuner

Sets gRPC keepalive parameters optimal for Temporal Cloud load balancers to prevent connection resets.

```go
// Automatically sets:
// - KeepAliveTime: 30s
// - KeepAliveTimeout: 10s
// - PermitWithoutStream: true
client, err := cloud.Dial(ctx, options)
```

### Java SDK Extension

#### 1. KeyStore Reloader

Watches the JKS/PKCS12 file and reloads the `SslContext`.

```java
Scope scope = new CloudScopeBuilder()
    .setCertPath(Paths.get("client.pem"))
    .setKeyPath(Paths.get("client.key"))
    .enableRotation(Duration.ofMinutes(1))
    .build();

WorkflowServiceStubsOptions options = WorkflowServiceStubsOptions.newBuilder()
    .setSslContext(scope.getSslContext())
    .build();
```

### TypeScript SDK Extension

#### 1. MTLS Reloader

Uses `fs.watch` to update the connection options.

```typescript
const connection = await CloudConnection.connect({
  address: "my-ns.tmprl.cloud",
  tls: {
    certPath: "./client.pem",
    keyPath: "./client.key",
    autoRotate: true,
  },
});
```

### Required Changes to Core SDKs

We aim to keep core SDKs generic, but we may need to open up `ClientOptions` to allow dynamic swapping of `Credential` providers if not already supported.

<!-- Source: secrets-management.md -->

## Secrets Management

### Secret Types

| Type           | Examples               | Storage             | Rotation         |
| -------------- | ---------------------- | ------------------- | ---------------- |
| Infrastructure | AWS keys, DB passwords | AWS Secrets Manager | 90 days          |
| Application    | API keys, JWT secrets  | AWS Secrets Manager | 90 days          |
| Customer       | mTLS certs, API keys   | Encrypted in DB     | Customer-managed |
| CI/CD          | Deploy keys, tokens    | GitHub Secrets      | 90 days          |
| Developer      | Personal tokens        | 1Password           | 90 days          |

### Secret Storage

#### AWS Secrets Manager

Primary secret store for all production secrets.

```hcl
# Terraform secret creation
resource "aws_secretsmanager_secret" "db_password" {
  name = "temporal-cloud/${var.environment}/db-password"

  tags = {
    Environment = var.environment
    ManagedBy   = "terraform"
    Rotation    = "enabled"
  }
}

resource "aws_secretsmanager_secret_rotation" "db_password" {
  secret_id           = aws_secretsmanager_secret.db_password.id
  rotation_lambda_arn = aws_lambda_function.secret_rotation.arn

  rotation_rules {
    automatically_after_days = 90
  }
}
```

#### Secret Naming Convention

```
temporal-cloud/{environment}/{service}/{secret-name}

Examples:
temporal-cloud/prod/api/jwt-signing-key
temporal-cloud/prod/db/master-password
temporal-cloud/prod/stripe/api-key
temporal-cloud/staging/api/jwt-signing-key
```

#### Kubernetes Integration

Use External Secrets Operator to sync secrets to K8s:

```yaml
apiVersion: external-secrets.io/v1beta1
kind: ExternalSecret
metadata:
  name: db-credentials
  namespace: cloud-platform
spec:
  refreshInterval: 1h
  secretStoreRef:
    name: aws-secrets-manager
    kind: ClusterSecretStore
  target:
    name: db-credentials
    creationPolicy: Owner
  data:
    - secretKey: password
      remoteRef:
        key: temporal-cloud/prod/db/master-password
```

### Secret Categories

#### Database Credentials

| Secret          | Used By           | Rotation       |
| --------------- | ----------------- | -------------- |
| Master password | Migrations, admin | 90 days (auto) |
| App read-write  | Application pods  | 90 days (auto) |
| App read-only   | Reporting         | 90 days (auto) |

#### API Keys

| Secret            | Used By         | Rotation                |
| ----------------- | --------------- | ----------------------- |
| Stripe API key    | Billing service | Manual (Stripe rotates) |
| SendGrid API key  | Email service   | 90 days                 |
| PagerDuty API key | Alerting        | 90 days                 |
| Datadog API key   | Monitoring      | 90 days                 |

#### Signing Keys

| Secret            | Used By        | Rotation                |
| ----------------- | -------------- | ----------------------- |
| JWT signing key   | Auth service   | 180 days (with overlap) |
| Webhook signing   | Event delivery | 180 days                |
| SAML signing cert | SSO            | 1 year                  |

#### Encryption Keys

| Secret              | Used By              | Rotation            |
| ------------------- | -------------------- | ------------------- |
| KMS master key      | All encryption       | Never (AWS managed) |
| Data encryption key | DB column encryption | 1 year              |

### Secret Rotation

#### Automatic Rotation (Lambda)

```python
# rotation_lambda.py
def lambda_handler(event, context):
    secret_id = event['SecretId']
    step = event['Step']

    if step == "createSecret":
        # Generate new secret value
        new_password = generate_password()
        secrets_client.put_secret_value(
            SecretId=secret_id,
            ClientRequestToken=event['ClientRequestToken'],
            SecretString=new_password,
            VersionStages=['AWSPENDING']
        )

    elif step == "setSecret":
        # Update the service with new secret
        pending = get_secret_value(secret_id, 'AWSPENDING')
        update_database_password(pending)

    elif step == "testSecret":
        # Verify new secret works
        pending = get_secret_value(secret_id, 'AWSPENDING')
        test_database_connection(pending)

    elif step == "finishSecret":
        # Promote pending to current
        secrets_client.update_secret_version_stage(
            SecretId=secret_id,
            VersionStage='AWSCURRENT',
            MoveToVersionId=event['ClientRequestToken'],
            RemoveFromVersionId=get_current_version(secret_id)
        )
```

#### JWT Key Rotation (Overlap Period)

```go
// Maintain two active keys during rotation
type JWTKeyManager struct {
    currentKey  *rsa.PrivateKey
    previousKey *rsa.PrivateKey  // Still valid for 7 days after rotation
    currentKid  string
    previousKid string
}

func (m *JWTKeyManager) Sign(claims jwt.Claims) (string, error) {
    // Always sign with current key
    return jwt.Sign(claims, m.currentKey, m.currentKid)
}

func (m *JWTKeyManager) Verify(token string) (*jwt.Claims, error) {
    kid := extractKid(token)

    // Try current key first
    if kid == m.currentKid {
        return jwt.Verify(token, m.currentKey)
    }

    // Fall back to previous key (during overlap)
    if kid == m.previousKid && m.previousKey != nil {
        return jwt.Verify(token, m.previousKey)
    }

    return nil, errors.New("unknown key id")
}
```

### Secret Access

#### IAM Policies

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": ["secretsmanager:GetSecretValue"],
      "Resource": [
        "arn:aws:secretsmanager:*:*:secret:temporal-cloud/prod/api/*"
      ],
      "Condition": {
        "StringEquals": {
          "aws:PrincipalTag/Service": "cloud-api"
        }
      }
    }
  ]
}
```

#### Kubernetes RBAC

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: secret-reader
  namespace: cloud-platform
rules:
  - apiGroups: [""]
    resources: ["secrets"]
    resourceNames: ["db-credentials", "api-keys"]
    verbs: ["get"]
```

### Secret Injection

#### Environment Variables (Preferred for simple secrets)

```yaml
env:
  - name: DATABASE_PASSWORD
    valueFrom:
      secretKeyRef:
        name: db-credentials
        key: password
```

#### Volume Mount (For files like certs)

```yaml
volumes:
  - name: tls-certs
    secret:
      secretName: api-tls
volumeMounts:
  - name: tls-certs
    mountPath: /etc/tls
    readOnly: true
```

#### Init Container (For complex setup)

```yaml
initContainers:
  - name: fetch-secrets
    image: amazon/aws-cli
    command:
      - sh
      - -c
      - |
        aws secretsmanager get-secret-value \
          --secret-id temporal-cloud/prod/api/config \
          --query SecretString --output text > /secrets/config.json
    volumeMounts:
      - name: secrets
        mountPath: /secrets
```

### Audit & Monitoring

#### Secret Access Logging

All secret access is logged to CloudTrail:

```json
{
  "eventName": "GetSecretValue",
  "userIdentity": {
    "arn": "arn:aws:sts::123456789:assumed-role/cloud-api-role/..."
  },
  "requestParameters": {
    "secretId": "temporal-cloud/prod/db/master-password"
  },
  "responseElements": null,
  "eventTime": "2025-01-15T10:30:00Z"
}
```

#### Alerts

```yaml
alerts:
  - name: UnauthorizedSecretAccess
    condition: |
      cloudtrail.eventName == "GetSecretValue" 
      AND cloudtrail.errorCode == "AccessDenied"
    severity: critical

  - name: SecretAccessFromUnknownIP
    condition: |
      cloudtrail.eventName == "GetSecretValue"
      AND cloudtrail.sourceIPAddress NOT IN allowed_ips
    severity: high
```

### Emergency Procedures

#### Secret Compromise Response

1. **Immediate** (< 15 min)

   - Rotate compromised secret
   - Revoke all sessions using that secret
   - Enable enhanced logging

2. **Short-term** (< 1 hour)

   - Audit access logs
   - Identify scope of compromise
   - Notify affected parties

3. **Follow-up**
   - Root cause analysis
   - Update access policies
   - Improve detection

#### Break Glass Access

For emergency access when normal paths fail:

```bash
# Requires 2 approvals from security team
aws secretsmanager get-secret-value \
  --secret-id temporal-cloud/prod/db/master-password \
  --profile break-glass
```

All break-glass access triggers immediate PagerDuty alert.

### Secrets Checklist

#### Before Production

- [ ] All secrets in Secrets Manager (not env vars, not code)
- [ ] Rotation enabled for all rotatable secrets
- [ ] IAM policies follow least privilege
- [ ] Audit logging enabled
- [ ] Alerts configured
- [ ] Break-glass procedure tested
- [ ] Secret naming follows convention

<!-- Source: security-hardening.md -->

## Security Hardening

### Defense in Depth

```
┌─────────────────────────────────────────────────────────────────┐
│                        Layer 1: Edge                             │
│  WAF, DDoS Protection, CDN, Rate Limiting                       │
├─────────────────────────────────────────────────────────────────┤
│                        Layer 2: Network                          │
│  VPC, Security Groups, NACLs, Private Subnets                   │
├─────────────────────────────────────────────────────────────────┤
│                        Layer 3: Application                      │
│  Authentication, Authorization, Input Validation                 │
├─────────────────────────────────────────────────────────────────┤
│                        Layer 4: Data                             │
│  Encryption at Rest, Encryption in Transit, Key Management      │
├─────────────────────────────────────────────────────────────────┤
│                        Layer 5: Monitoring                       │
│  Logging, Alerting, Anomaly Detection                           │
└─────────────────────────────────────────────────────────────────┘
```

### Edge Security

#### AWS WAF Rules

```hcl
resource "aws_wafv2_web_acl" "main" {
  name  = "temporal-cloud-waf"
  scope = "REGIONAL"

  default_action {
    allow {}
  }

  # Rate limiting
  rule {
    name     = "RateLimitRule"
    priority = 1

    action {
      block {}
    }

    statement {
      rate_based_statement {
        limit              = 2000
        aggregate_key_type = "IP"
      }
    }

    visibility_config {
      sampled_requests_enabled   = true
      cloudwatch_metrics_enabled = true
      metric_name                = "RateLimitRule"
    }
  }

  # SQL Injection
  rule {
    name     = "SQLInjectionRule"
    priority = 2

    override_action {
      none {}
    }

    statement {
      managed_rule_group_statement {
        name        = "AWSManagedRulesSQLiRuleSet"
        vendor_name = "AWS"
      }
    }
  }

  # Known bad inputs
  rule {
    name     = "KnownBadInputsRule"
    priority = 3

    override_action {
      none {}
    }

    statement {
      managed_rule_group_statement {
        name        = "AWSManagedRulesKnownBadInputsRuleSet"
        vendor_name = "AWS"
      }
    }
  }

  # Bot control
  rule {
    name     = "BotControlRule"
    priority = 4

    override_action {
      none {}
    }

    statement {
      managed_rule_group_statement {
        name        = "AWSManagedRulesBotControlRuleSet"
        vendor_name = "AWS"
      }
    }
  }
}
```

#### DDoS Protection

```hcl
# AWS Shield Advanced
resource "aws_shield_protection" "alb" {
  name         = "temporal-cloud-alb"
  resource_arn = aws_lb.api.arn
}

# Automatic DDoS response
resource "aws_shield_protection_group" "main" {
  protection_group_id = "temporal-cloud"
  aggregation         = "MAX"
  pattern             = "BY_RESOURCE_TYPE"
  resource_type       = "APPLICATION_LOAD_BALANCER"
}
```

### Network Hardening

#### VPC Configuration

```hcl
# No default VPC
# Custom VPC with private subnets only for workloads

resource "aws_vpc" "main" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name = "temporal-cloud-prod"
  }
}

# Flow logs for network monitoring
resource "aws_flow_log" "main" {
  vpc_id          = aws_vpc.main.id
  traffic_type    = "ALL"
  iam_role_arn    = aws_iam_role.flow_log.arn
  log_destination = aws_cloudwatch_log_group.flow_log.arn
}
```

#### Security Groups (Least Privilege)

```hcl
# ALB - Only HTTPS from internet
resource "aws_security_group" "alb" {
  name   = "alb-sg"
  vpc_id = aws_vpc.main.id

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port       = 8080
    to_port         = 8080
    protocol        = "tcp"
    security_groups = [aws_security_group.app.id]
  }
}

# App - Only from ALB
resource "aws_security_group" "app" {
  name   = "app-sg"
  vpc_id = aws_vpc.main.id

  ingress {
    from_port       = 8080
    to_port         = 8080
    protocol        = "tcp"
    security_groups = [aws_security_group.alb.id]
  }

  egress {
    from_port       = 5432
    to_port         = 5432
    protocol        = "tcp"
    security_groups = [aws_security_group.db.id]
  }
}

# DB - Only from App
resource "aws_security_group" "db" {
  name   = "db-sg"
  vpc_id = aws_vpc.main.id

  ingress {
    from_port       = 5432
    to_port         = 5432
    protocol        = "tcp"
    security_groups = [aws_security_group.app.id]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = []  # No egress
  }
}
```

### Application Hardening

#### Input Validation

```go
// Validate all input
func (s *Server) CreateNamespace(ctx context.Context, req *CreateNamespaceRequest) error {
    // Length limits
    if len(req.Name) < 2 || len(req.Name) > 63 {
        return status.Error(codes.InvalidArgument, "name must be 2-63 characters")
    }

    // Character whitelist
    if !regexp.MustCompile(`^[a-z][a-z0-9-]*[a-z0-9]$`).MatchString(req.Name) {
        return status.Error(codes.InvalidArgument, "name must be lowercase alphanumeric")
    }

    // Sanitize for logging (prevent log injection)
    sanitizedName := sanitize(req.Name)
    log.Info("Creating namespace", "name", sanitizedName)

    // ... create namespace
}
```

#### SQL Injection Prevention

```go
// ALWAYS use parameterized queries
func GetOrganization(ctx context.Context, id string) (*Organization, error) {
    // ✅ Correct - parameterized
    row := db.QueryRowContext(ctx,
        "SELECT id, name FROM organizations WHERE id = $1",
        id)

    // ❌ NEVER do this
    // row := db.QueryRowContext(ctx,
    //     "SELECT id, name FROM organizations WHERE id = '" + id + "'")

    var org Organization
    err := row.Scan(&org.ID, &org.Name)
    return &org, err
}
```

#### XSS Prevention

```typescript
// React automatically escapes output
// But be careful with dangerouslySetInnerHTML

// ✅ Safe
return <div>{userInput}</div>;

// ❌ Dangerous - avoid unless absolutely necessary
return <div dangerouslySetInnerHTML={{ __html: userInput }} />;

// If you must render HTML, sanitize first
import DOMPurify from "dompurify";
return (
  <div dangerouslySetInnerHTML={{ __html: DOMPurify.sanitize(userInput) }} />
);
```

#### CSRF Protection

```go
// Use CSRF tokens for state-changing operations
func CSRFMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "GET" && r.Method != "HEAD" && r.Method != "OPTIONS" {
            token := r.Header.Get("X-CSRF-Token")
            if !validateCSRFToken(r.Context(), token) {
                http.Error(w, "Invalid CSRF token", http.StatusForbidden)
                return
            }
        }
        next.ServeHTTP(w, r)
    })
}
```

### Data Hardening

#### Encryption at Rest

```hcl
# RDS encryption
resource "aws_db_instance" "main" {
  storage_encrypted = true
  kms_key_id        = aws_kms_key.db.arn
}

# S3 encryption
resource "aws_s3_bucket_server_side_encryption_configuration" "main" {
  bucket = aws_s3_bucket.data.id

  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm     = "aws:kms"
      kms_master_key_id = aws_kms_key.s3.arn
    }
  }
}

# EBS encryption
resource "aws_ebs_encryption_by_default" "main" {
  enabled = true
}
```

#### Encryption in Transit

```yaml
# Force TLS 1.3
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    alb.ingress.kubernetes.io/ssl-policy: ELBSecurityPolicy-TLS13-1-2-2021-06
```

#### Secrets in Memory

```go
// Clear sensitive data from memory when done
func ProcessPayment(card *CreditCard) error {
    defer func() {
        // Zero out sensitive data
        for i := range card.Number {
            card.Number[i] = 0
        }
        for i := range card.CVV {
            card.CVV[i] = 0
        }
    }()

    // Process payment...
}
```

### Container Hardening

#### Dockerfile Best Practices

```dockerfile
# Use minimal base image
FROM gcr.io/distroless/static:nonroot

# Don't run as root
USER nonroot:nonroot

# No shell, minimal attack surface
# Read-only filesystem where possible
```

#### Pod Security

```yaml
apiVersion: v1
kind: Pod
spec:
  securityContext:
    runAsNonRoot: true
    runAsUser: 65534
    fsGroup: 65534
    seccompProfile:
      type: RuntimeDefault
  containers:
    - name: api
      securityContext:
        allowPrivilegeEscalation: false
        readOnlyRootFilesystem: true
        capabilities:
          drop:
            - ALL
```

### Hardening Checklist

#### Infrastructure

- [ ] No public S3 buckets
- [ ] No default VPC
- [ ] All storage encrypted
- [ ] Flow logs enabled
- [ ] WAF configured
- [ ] Shield enabled

#### Application

- [ ] Input validation on all endpoints
- [ ] Parameterized queries only
- [ ] CSRF protection
- [ ] Rate limiting
- [ ] Security headers set

#### Access

- [ ] MFA required
- [ ] Least privilege IAM
- [ ] No long-lived credentials
- [ ] Regular access reviews

#### Monitoring

- [ ] All access logged
- [ ] Anomaly detection enabled
- [ ] Alert on suspicious activity
- [ ] Regular security scans

<!-- Source: security.md -->

## Security

### Security Architecture

#### Network Security

- VPC with private subnets for all services
- Security groups with least-privilege rules
- WAF on all public endpoints
- DDoS protection via AWS Shield
- No direct internet access from private subnets

#### Data Security

- Encryption at rest: AES-256-GCM
- Encryption in transit: TLS 1.3
- Database encryption: RDS encryption
- S3 encryption: SSE-S3 or SSE-KMS

#### Access Control

- IAM roles for all AWS access (no static keys)
- K8s RBAC for cluster access
- SSO for all internal tools
- MFA required for all human access

### Authentication

#### User Authentication

- Email/password with MFA
- SAML 2.0 SSO (Okta, Azure AD, etc.)
- Session timeout: 24 hours
- Refresh token rotation

#### API Authentication

- API Keys (Bearer token)
- mTLS certificates (namespace access)
- Service account tokens

#### mTLS Requirements

- CA certificates: X.509 v3
- Key size: RSA 2048+ or ECDSA P-256+
- Validity: Max 1 year
- Certificate filters for fine-grained access

### Authorization

#### Account-Level Roles

| Role          | Permissions                   |
| ------------- | ----------------------------- |
| Account Owner | Full access including billing |
| Global Admin  | Full access except billing    |
| Finance Admin | Billing only                  |
| Developer     | Create namespaces, manage own |
| Read-Only     | View only                     |

#### Namespace-Level Permissions

| Permission      | Capabilities           |
| --------------- | ---------------------- |
| Namespace Admin | Full namespace control |
| Write           | CRUD workflows         |
| Read-Only       | View only              |

### Secrets Management

#### Storage

- AWS Secrets Manager for all secrets
- External Secrets Operator for K8s injection
- No secrets in environment variables
- No secrets in code or config files

#### Rotation

- Database credentials: 90 days
- API keys: User-managed, recommend 90 days
- Service account keys: 90 days
- TLS certificates: Before expiry

### Audit Logging

#### Logged Events

- All authentication attempts
- All authorization decisions
- All state-changing operations
- All admin actions

#### Log Format

```json
{
  "timestamp": "2025-01-01T00:00:00Z",
  "operation": "CreateNamespace",
  "actor": {
    "type": "user",
    "id": "user-123",
    "email": "user@example.com"
  },
  "resource": {
    "type": "namespace",
    "id": "ns-456"
  },
  "status": "success",
  "ip": "1.2.3.4"
}
```

#### Retention

- Hot storage: 90 days (PostgreSQL)
- Cold storage: 7 years (S3 Glacier)

### Compliance

#### SOC 2 Type II

- Status: Planned
- Controls: Security, Availability, Confidentiality
- Annual audit

#### GDPR

- Data processing agreements
- Right to erasure support
- Data export capability

#### Security Testing

- Penetration testing: Annual
- Vulnerability scanning: Weekly
- Dependency scanning: On every build

### Incident Response

#### Severity Levels

| Level | Description          | Response  |
| ----- | -------------------- | --------- |
| SEV1  | Data breach          | Immediate |
| SEV2  | Service compromise   | 1 hour    |
| SEV3  | Vulnerability found  | 24 hours  |
| SEV4  | Security improvement | 1 week    |

#### Response Process

1. Detect and alert
2. Assess severity
3. Contain the incident
4. Eradicate the threat
5. Recover services
6. Post-incident review

<!-- Source: sla.md -->

## Service Level Agreements (SLA)

### Availability

#### Definitions

- **Unavailable**: When valid API requests return 5xx errors or time out for > 1 minute continuously.
- **Maintenance**: Planned downtime with 24h notice (excluded from SLA).

#### Service Tiers

| Plan       | Availability Target        | Financial Credit                                 |
| ---------- | -------------------------- | ------------------------------------------------ |
| Essential  | 99.9%                      | 10% if < 99.9%, 25% if < 99.0%                   |
| Business   | 99.9%                      | 10% if < 99.9%, 25% if < 99.0%                   |
| Enterprise | 99.99% (with Multi-Region) | 10% if < 99.99%, 30% if < 99.9%, 100% if < 99.0% |

#### Calculation

`Availability = (Total Minutes - Downtime Minutes) / Total Minutes`

### Latency SLO

Target response times (P99) measured at the load balancer:

- **StartWorkflow**: < 200ms
- **SignalWorkflow**: < 200ms
- **QueryWorkflow**: < 500ms (excluding worker processing time)
- **History Events**: < 100ms

### Support Response Times

| Priority      | Essential       | Business        | Enterprise     |
| ------------- | --------------- | --------------- | -------------- |
| P0 (Critical) | N/A             | 1 hour          | 15 mins        |
| P1 (High)     | N/A             | 4 hours         | 1 hour         |
| P2 (Normal)   | 2 business days | 1 business day  | 4 hours        |
| P3 (Low)      | Best effort     | 2 business days | 1 business day |

### Exclusions

The SLA does not apply to:

1. Beta/Experimental features
2. Client-side errors (bad input, rate limiting)
3. Force Majeure events
4. Planned maintenance
5. Suspended accounts (due to non-payment)

### Claim Process

1. Customer must submit claim via Support Portal within 30 days.
2. Claim must include timestamps and error logs.
3. Temporal validates against server-side logs.
4. Credits applied to next invoice.

<!-- Source: state-management.md -->

## State Management

### State Types

| State Type        | Storage        | Scope              |
| ----------------- | -------------- | ------------------ |
| Terraform State   | S3 + DynamoDB  | Infrastructure     |
| Kubernetes State  | etcd (managed) | Cluster            |
| Application State | PostgreSQL     | Business data      |
| Session State     | Redis          | User sessions      |
| Workflow State    | Temporal       | Workflow execution |
| Cache State       | Redis          | Ephemeral          |

### Terraform State Management

#### Backend Configuration

```hcl
# backend.tf
terraform {
  backend "s3" {
    bucket         = "temporal-cloud-terraform-state"
    key            = "environments/production/us-east-1/terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
    dynamodb_table = "terraform-state-lock"

    # Cross-region replication for DR
    # Configured at bucket level
  }
}
```

#### State Structure

```
s3://temporal-cloud-terraform-state/
├── global/
│   ├── iam/terraform.tfstate
│   ├── route53/terraform.tfstate
│   └── budgets/terraform.tfstate
├── environments/
│   ├── production/
│   │   ├── us-east-1/terraform.tfstate
│   │   ├── eu-west-1/terraform.tfstate
│   │   └── gcp-us-central1/terraform.tfstate
│   ├── staging/
│   │   └── us-east-1/terraform.tfstate
│   └── dev/
│       └── us-east-1/terraform.tfstate
└── modules/  # No state here, just source
```

#### State Locking

```hcl
resource "aws_dynamodb_table" "terraform_lock" {
  name         = "terraform-state-lock"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }

  # Point-in-time recovery
  point_in_time_recovery {
    enabled = true
  }
}
```

#### State Operations

```bash
# List resources in state
terraform state list

# Show specific resource
terraform state show aws_instance.api

# Move resource (refactoring)
terraform state mv aws_instance.api aws_instance.cloud_api

# Remove from state (without destroying)
terraform state rm aws_instance.old_api

# Import existing resource
terraform import aws_instance.api i-1234567890abcdef0

# Refresh state
terraform refresh
```

#### State Backup

```yaml
# S3 bucket versioning for state history
resource "aws_s3_bucket_versioning" "state" {
  bucket = aws_s3_bucket.terraform_state.id
  versioning_configuration {
    status = "Enabled"
  }
}

# Cross-region replication
resource "aws_s3_bucket_replication_configuration" "state" {
  bucket = aws_s3_bucket.terraform_state.id

  rule {
    status = "Enabled"
    destination {
      bucket        = aws_s3_bucket.terraform_state_dr.arn
      storage_class = "STANDARD"
    }
  }
}
```

### Application State (PostgreSQL)

#### Schema Versioning

Track schema version in database:

```sql
CREATE TABLE schema_migrations (
    version BIGINT PRIMARY KEY,
    applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

#### Data Integrity

```sql
-- Use transactions for state changes
BEGIN;
    UPDATE subscriptions SET plan = 'business' WHERE org_id = $1;
    INSERT INTO audit_events (org_id, action, details) VALUES ($1, 'plan_upgrade', $2);
COMMIT;
```

#### Soft Deletes

Never hard delete, use soft deletes for audit trail:

```sql
ALTER TABLE organizations ADD COLUMN deleted_at TIMESTAMPTZ;

-- "Delete" organization
UPDATE organizations SET deleted_at = NOW() WHERE id = $1;

-- Query only active
SELECT * FROM organizations WHERE deleted_at IS NULL;
```

### Session State (Redis)

#### Session Structure

```json
{
  "session_id": "sess_abc123",
  "user_id": "user_456",
  "org_id": "org_789",
  "created_at": "2025-01-15T10:00:00Z",
  "expires_at": "2025-01-16T10:00:00Z",
  "ip": "1.2.3.4",
  "user_agent": "Mozilla/5.0..."
}
```

#### Session Commands

```go
// Create session
func CreateSession(ctx context.Context, userID string) (*Session, error) {
    session := &Session{
        ID:        generateSessionID(),
        UserID:    userID,
        CreatedAt: time.Now(),
        ExpiresAt: time.Now().Add(24 * time.Hour),
    }

    data, _ := json.Marshal(session)
    err := redis.Set(ctx, "session:"+session.ID, data, 24*time.Hour).Err()
    return session, err
}

// Validate session
func GetSession(ctx context.Context, sessionID string) (*Session, error) {
    data, err := redis.Get(ctx, "session:"+sessionID).Bytes()
    if err == redis.Nil {
        return nil, ErrSessionNotFound
    }

    var session Session
    json.Unmarshal(data, &session)

    if time.Now().After(session.ExpiresAt) {
        redis.Del(ctx, "session:"+sessionID)
        return nil, ErrSessionExpired
    }

    return &session, nil
}

// Revoke session
func RevokeSession(ctx context.Context, sessionID string) error {
    return redis.Del(ctx, "session:"+sessionID).Err()
}

// Revoke all sessions for user
func RevokeAllUserSessions(ctx context.Context, userID string) error {
    keys, _ := redis.Keys(ctx, "session:*").Result()
    for _, key := range keys {
        data, _ := redis.Get(ctx, key).Bytes()
        var session Session
        json.Unmarshal(data, &session)
        if session.UserID == userID {
            redis.Del(ctx, key)
        }
    }
    return nil
}
```

### Workflow State (Temporal)

Temporal handles workflow state internally. Key considerations:

#### State Persistence

```go
// Workflow state is automatically persisted after each event
func OrderWorkflow(ctx workflow.Context, order Order) error {
    // State checkpoint 1
    err := workflow.ExecuteActivity(ctx, ValidateOrder, order).Get(ctx, nil)

    // State checkpoint 2
    err = workflow.ExecuteActivity(ctx, ChargePayment, order).Get(ctx, nil)

    // State checkpoint 3
    err = workflow.ExecuteActivity(ctx, FulfillOrder, order).Get(ctx, nil)

    return nil
}
```

#### Continue-As-New

Prevent unbounded history:

```go
func LongRunningWorkflow(ctx workflow.Context, state State) error {
    for {
        // Process events...

        // Check history size
        info := workflow.GetInfo(ctx)
        if info.GetCurrentHistoryLength() > 10000 {
            return workflow.NewContinueAsNewError(ctx, LongRunningWorkflow, state)
        }
    }
}
```

### Cache State (Redis)

#### Cache Invalidation

```go
// Cache with TTL
func GetOrganization(ctx context.Context, id string) (*Organization, error) {
    // Try cache
    cacheKey := "org:" + id
    data, err := redis.Get(ctx, cacheKey).Bytes()
    if err == nil {
        var org Organization
        json.Unmarshal(data, &org)
        return &org, nil
    }

    // Cache miss - fetch from DB
    org, err := db.GetOrganization(ctx, id)
    if err != nil {
        return nil, err
    }

    // Populate cache
    data, _ = json.Marshal(org)
    redis.Set(ctx, cacheKey, data, 5*time.Minute)

    return org, nil
}

// Invalidate on update
func UpdateOrganization(ctx context.Context, org *Organization) error {
    err := db.UpdateOrganization(ctx, org)
    if err != nil {
        return err
    }

    // Invalidate cache
    redis.Del(ctx, "org:"+org.ID)

    return nil
}
```

#### Cache Stampede Prevention

```go
// Use singleflight to prevent multiple concurrent fetches
var group singleflight.Group

func GetOrganizationSafe(ctx context.Context, id string) (*Organization, error) {
    v, err, _ := group.Do(id, func() (interface{}, error) {
        return GetOrganization(ctx, id)
    })

    if err != nil {
        return nil, err
    }
    return v.(*Organization), nil
}
```

### State Recovery

#### Terraform State Recovery

```bash
# If state is corrupted, restore from S3 version
aws s3api list-object-versions \
  --bucket temporal-cloud-terraform-state \
  --prefix environments/production/us-east-1/terraform.tfstate

# Restore specific version
aws s3api get-object \
  --bucket temporal-cloud-terraform-state \
  --key environments/production/us-east-1/terraform.tfstate \
  --version-id "abc123" \
  terraform.tfstate.backup
```

#### Database State Recovery

See `dr.md` for point-in-time recovery procedures.

#### Redis State Recovery

Redis state is ephemeral. On failure:

1. Sessions require re-login
2. Cache repopulates automatically
3. Rate limit counters reset (acceptable)

<!-- Source: stripe-mapping.md -->

## Stripe Product Mapping

### Overview

Maps internal plan tiers and usage metrics to Stripe Products and Prices for billing.

### Product Structure

```
Temporal Cloud (Product)
├── Essential Plan (Price - recurring)
├── Business Plan (Price - recurring)
├── Enterprise Plan (Price - recurring)
├── Actions (Price - metered)
├── Active Storage (Price - metered)
└── Retained Storage (Price - metered)
```

### Price ID Mapping

#### Production Environment

| Item             | Stripe Price ID                | Type      | Amount       |
| ---------------- | ------------------------------ | --------- | ------------ |
| Essential Plan   | `price_essential_monthly_prod` | Recurring | $100/mo      |
| Business Plan    | `price_business_monthly_prod`  | Recurring | $500/mo      |
| Actions (per M)  | `price_actions_metered_prod`   | Metered   | $50/M        |
| Active Storage   | `price_active_storage_prod`    | Metered   | $0.042/GBh   |
| Retained Storage | `price_retained_storage_prod`  | Metered   | $0.00105/GBh |
| SCIM Add-on      | `price_scim_addon_prod`        | Recurring | $500/mo      |

#### Test Environment

| Item             | Stripe Price ID                | Type      | Amount       |
| ---------------- | ------------------------------ | --------- | ------------ |
| Essential Plan   | `price_essential_monthly_test` | Recurring | $100/mo      |
| Business Plan    | `price_business_monthly_test`  | Recurring | $500/mo      |
| Actions (per M)  | `price_actions_metered_test`   | Metered   | $50/M        |
| Active Storage   | `price_active_storage_test`    | Metered   | $0.042/GBh   |
| Retained Storage | `price_retained_storage_test`  | Metered   | $0.00105/GBh |
| SCIM Add-on      | `price_scim_addon_test`        | Recurring | $500/mo      |

### Configuration

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

### Subscription Creation

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

### Usage Reporting

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

### Webhook Handling

#### Supported Events

| Event                           | Handler                  |
| ------------------------------- | ------------------------ |
| `customer.subscription.created` | Sync subscription to DB  |
| `customer.subscription.updated` | Update plan/status       |
| `customer.subscription.deleted` | Mark as canceled         |
| `invoice.paid`                  | Update invoice status    |
| `invoice.payment_failed`        | Trigger dunning workflow |
| `invoice.finalized`             | Store invoice details    |

#### Webhook Handler

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

#### Plan Change from Stripe Portal

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

### Testing

#### Test Mode Setup

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

#### Webhook Testing

```bash
# Forward webhooks to local
stripe listen --forward-to localhost:8081/webhooks/stripe

# Trigger test events
stripe trigger invoice.paid
stripe trigger customer.subscription.updated
```

<!-- Source: support-escalation.md -->

## Support Escalation

### Support Tiers

| Tier             | Access               | Response Time   | Hours          |
| ---------------- | -------------------- | --------------- | -------------- |
| Community        | GitHub Issues        | Best effort     | Community      |
| Essential        | Email                | 2 business days | Business hours |
| Business         | Email + Chat         | 4 hours (P1)    | Business hours |
| Enterprise       | Email + Chat + Slack | 1 hour (P0)     | 24/7           |
| Mission Critical | Dedicated + Phone    | 15 min (P0)     | 24/7           |

### Priority Classification

| Priority | Description     | Example               |
| -------- | --------------- | --------------------- |
| P0       | Production down | All workflows failing |
| P1       | Major impact    | 50%+ error rate       |
| P2       | Moderate impact | Single feature broken |
| P3       | Minor impact    | UI bug, question      |

### Escalation Flow

```
┌─────────────────────────────────────────────────────────┐
│                    Customer                              │
└───────────────────────┬─────────────────────────────────┘
                        │
                        ▼
┌─────────────────────────────────────────────────────────┐
│              Tier 1: Support Agent                       │
│  • Answer common questions                               │
│  • Triage and categorize                                 │
│  • Collect diagnostic information                        │
│  • SLA: 15 min response (P0/P1)                         │
└───────────────────────┬─────────────────────────────────┘
                        │ Escalate if unresolved (30 min)
                        ▼
┌─────────────────────────────────────────────────────────┐
│              Tier 2: Support Engineer                    │
│  • Deep technical troubleshooting                        │
│  • Access to internal systems                            │
│  • Can reproduce issues                                  │
│  • SLA: Resolve or escalate (2 hours)                   │
└───────────────────────┬─────────────────────────────────┘
                        │ Escalate if code change needed
                        ▼
┌─────────────────────────────────────────────────────────┐
│              Tier 3: Engineering                         │
│  • Bug fixes                                             │
│  • Feature requests                                      │
│  • Architecture guidance                                 │
│  • SLA: Based on severity                               │
└─────────────────────────────────────────────────────────┘
```

### Ticket Lifecycle

```
┌─────────┐     ┌─────────┐     ┌─────────┐     ┌─────────┐
│  New    │────▶│  Open   │────▶│ Pending │────▶│ Solved  │
└─────────┘     └─────────┘     └─────────┘     └─────────┘
                     │               │
                     │               │ Customer unresponsive
                     │               ▼
                     │          ┌─────────┐
                     └─────────▶│ Closed  │
                                └─────────┘
```

### Information Gathering

#### Required for P0/P1

```markdown
## Ticket Template

**Account ID**:
**Namespace(s)**:
**Region**:

**Impact**:

- Start time:
- Users affected:
- Workflows affected:

**Description**:
[What is happening vs what should happen]

**Steps to Reproduce**:

1.
2.

**Error Messages**:
[Exact error text or screenshot]

**Recent Changes**:
[Any deployments, config changes in last 24h]
```

#### Diagnostic Commands

```bash
# Namespace health
temporal operator namespace describe --namespace $NS

# Recent failures
temporal workflow list --namespace $NS --query "CloseStatus='Failed'"

# Metrics snapshot
curl https://$NS.tmprl.cloud/prometheus/metrics | grep error
```

### Escalation Triggers

#### Tier 1 → Tier 2

- Unable to resolve with documentation
- Requires internal system access
- Customer requests escalation
- Approaching SLA breach

#### Tier 2 → Tier 3

- Confirmed bug requiring code change
- Infrastructure issue
- Security incident
- Feature request from Enterprise customer

#### Tier 3 → Management

- P0 unresolved > 2 hours
- Multiple customers affected
- Data loss confirmed
- Security breach

### Communication Templates

#### Initial Response (P0/P1)

```
Hi [Name],

Thank you for contacting Temporal Cloud Support.

I understand you're experiencing [issue summary]. This is being
treated as a Priority [X] issue, and I'm actively investigating.

To help me troubleshoot, please provide:
- [Specific info needed]

I will update you within [time] or sooner if I have new information.

Best,
[Agent Name]
```

#### Escalation to Engineering

```
## Engineering Escalation

**Ticket**: #12345
**Customer**: Acme Corp (Enterprise)
**Priority**: P1

**Summary**:
Namespace creation failing with timeout errors since 14:00 UTC

**Impact**:
- 50+ create requests failed
- Customer unable to onboard new teams

**Troubleshooting Done**:
1. Verified namespace quotas (OK)
2. Checked region health (OK)
3. Reviewed logs (timeout to provisioning service)

**Hypothesis**:
Provisioning service may be overloaded

**Request**:
Engineering review of provisioning service logs
```

#### Resolution

```
Hi [Name],

Great news - this issue has been resolved.

**Root Cause**: [Brief explanation]

**Resolution**: [What was done]

**Preventive Measures**: [If applicable]

Your namespaces should now be functioning normally. Please
confirm on your end and let me know if you encounter any
further issues.

Best,
[Agent Name]
```

### Knowledge Base

#### Article Structure

```markdown
# [Title]

## Symptoms

What the customer sees

## Cause

Why this happens

## Solution

Step-by-step fix

## Prevention

How to avoid in future

## Related Articles

- [Link 1]
- [Link 2]
```

#### Common Issues

| Issue               | Resolution                           |
| ------------------- | ------------------------------------ |
| Certificate expired | Guide to rotate certs                |
| Rate limited        | Explain limits, request increase     |
| Connection refused  | Check firewall, verify endpoint      |
| Slow workflows      | Review history size, Continue-As-New |

### Metrics

| Metric                   | Target    | Current |
| ------------------------ | --------- | ------- |
| First Response Time      | < 1h (P0) |         |
| Resolution Time          | < 4h (P0) |         |
| CSAT                     | > 4.5/5   |         |
| First Contact Resolution | > 60%     |         |
| Escalation Rate          | < 20%     |         |

### On-Call Support

#### Schedule

- Primary: Support Engineer (24/7 for Enterprise)
- Backup: Second Support Engineer
- Engineering Escalation: On-call Engineer

#### Handoff

At shift change:

1. Review open P0/P1 tickets
2. Summarize status and next steps
3. Introduce to customer if active incident

<!-- Source: system-optimization.md -->

## System Optimization & Performance

### Kernel Tuning (Linux)

Optimize the OS for high-throughput, low-latency workloads.

```bash
# /etc/sysctl.d/99-temporal.conf

# Network
net.core.somaxconn = 32768          # Increase backlog
net.ipv4.tcp_max_syn_backlog = 8192
net.ipv4.ip_local_port_range = 1024 65535
net.ipv4.tcp_tw_reuse = 1           # Reuse TIME_WAIT sockets

# TCP buffers (for high BDP)
net.core.rmem_max = 16777216
net.core.wmem_max = 16777216
net.ipv4.tcp_rmem = 4096 87380 16777216
net.ipv4.tcp_wmem = 4096 65536 16777216

# IO
fs.file-max = 2097152               # Open files limit
vm.swappiness = 10                  # Prefer RAM over swap
```

### Go Runtime Optimization

#### Garbage Collection (GC)

Temporal is memory-intensive.

- `GOGC=100` (Default): Balanced.
- `GOGC=200`: Trade memory for CPU (less GC). Good for History service if RAM available.
- `GOMEMLIMIT`: Set to 90% of container limit to avoid OOM kills.

#### Concurrency

- `GOMAXPROCS`: Auto-detect container limits (use `automaxprocs` library).

### Database Optimization (PostgreSQL)

#### Connection Pooling

Use **PgBouncer** at transaction level.

- App -> PgBouncer (local sidecar) -> RDS
- Reduces connection overhead on Postgres.

#### Query Tuning

- **Prepared Statements**: Always use.
- **JIT Compilation**: Enable for complex analytical queries.
- **Checkpointing**: Increase `max_wal_size` to reduce checkpoint frequency.

### Temporal-Specific Tuning

#### History Shards

- Default: 512 shards.
- Cloud Scale: 4k, 8k, or 16k shards.
- **Rule**: ~1k workflows per shard active. Too many shards = CPU overhead. Too few = Lock contention.

#### Cache Sizes

- `HistoryCacheMaxSize`: Increase to fit active working set in memory.
- Reduces DB IOPS significantly.

#### Batching

- Enable `UseTransactionBatching` in history service.
- Batches DB writes for higher throughput.

### Load Balancer Tuning

#### gRPC Tuning

- **HTTP/2 Keepalives**:
  - `PERMIT_WITHOUT_STREAM: true`
  - `MIN_TIME: 10s`
  - Prevents load balancers (ALB) from killing idle connections silently.

#### Connection Balancing

- Use L7 load balancing (Envoy/Alb) to distribute gRPC _requests_, not just connections.
- Prevents "hot" backend pods.

### Benchmarking

#### Standard Suite

Run `maru` (Temporal benchmark tool) weekly:

- **Throughput**: Max starts/sec.
- **Latency**: End-to-end workflow completion time.
- **Scalability**: Linear growth with added nodes?

#### Profile Continuous (Pprof)

- Enable `net/http/pprof`.
- Continuous Profiling (e.g., Pyroscope) to find CPU hot paths in production.

<!-- Source: tasks.md -->

## Implementation Tasks

### Phase 1: Foundation (Weeks 1-4)

#### Task 1.1: Repository Setup

- [ ] Fork temporalio/temporal to YOUR_ORG/temporal
- [ ] Fork temporalio/cloud-api to YOUR_ORG/cloud-api
- [ ] Fork temporalio/tcld to YOUR_ORG/tcld
- [ ] Fork temporalio/terraform-provider-temporalcloud
- [ ] Create YOUR_ORG/temporal-cloud-platform
- [ ] Create YOUR_ORG/temporal-cloud-console
- [ ] Create YOUR_ORG/temporal-cloud-infra
- [ ] Configure branch protection rules
- [ ] Setup upstream sync workflow
- [ ] Setup pre-commit hooks

#### Task 1.2: Infrastructure Bootstrap

- [ ] Create AWS accounts (dev, staging, prod)
- [ ] Setup Terraform state backend (S3 + DynamoDB)
- [ ] Create base VPC module
- [ ] Create EKS cluster module
- [ ] Create RDS module
- [ ] Create Redis module
- [ ] Deploy dev environment
- [ ] Setup GitHub Actions runners

#### Task 1.3: Database Schema

- [ ] Create 001_organizations.sql
- [ ] Create 002_subscriptions.sql
- [ ] Create 003_usage.sql
- [ ] Create 004_invoices.sql
- [ ] Create 005_audit.sql
- [ ] Setup migration tooling
- [ ] Run migrations in dev

#### Task 1.4: Proto Definitions

- [ ] Create organization/v1/message.proto
- [ ] Create organization/v1/service.proto
- [ ] Create subscription/v1/message.proto
- [ ] Create subscription/v1/service.proto
- [ ] Create billing/v1/message.proto
- [ ] Create billing/v1/service.proto
- [ ] Create audit/v1/message.proto
- [ ] Create audit/v1/service.proto
- [ ] Generate Go code
- [ ] Generate TypeScript code

### Phase 2: Metering (Weeks 5-8)

#### Task 2.1: Metering Interceptor

- [ ] Create common/cloud/metering/types.go
- [ ] Create common/cloud/metering/interceptor.go
- [ ] Create common/cloud/metering/collector.go
- [ ] Add unit tests
- [ ] Add integration tests

#### Task 2.2: Usage Aggregation

- [ ] Create internal/metering/aggregator.go
- [ ] Create internal/metering/store.go
- [ ] Create usage aggregation workflow
- [ ] Add unit tests
- [ ] Deploy to dev

#### Task 2.3: Storage Metering

- [ ] Create storage size query in persistence layer
- [ ] Create storage calculator
- [ ] Add to aggregation workflow
- [ ] Add unit tests

### Phase 3: Billing (Weeks 9-14)

#### Task 3.1: Stripe Integration

- [ ] Create internal/billing/stripe.go
- [ ] Create internal/billing/webhooks.go
- [ ] Setup Stripe test account
- [ ] Add unit tests
- [ ] Test webhook handling

#### Task 3.2: Invoice Generation

- [ ] Create internal/billing/calculator.go
- [ ] Create internal/billing/service.go
- [ ] Create invoice workflow
- [ ] Add unit tests
- [ ] Test invoice generation

#### Task 3.3: Quota Enforcement

- [ ] Create common/cloud/quota/enforcer.go
- [ ] Create common/cloud/quota/cache.go
- [ ] Add to interceptor chain
- [ ] Add unit tests
- [ ] Add integration tests

### Phase 4: Security (Weeks 15-20)

#### Task 4.1: SAML SSO

- [ ] Create internal/auth/saml/provider.go
- [ ] Create internal/auth/saml/handler.go
- [ ] Create common/cloud/auth/saml_claim_mapper.go
- [ ] Add unit tests
- [ ] Test with Okta
- [ ] Test with Azure AD

#### Task 4.2: SCIM

- [ ] Create internal/auth/scim/handler.go
- [ ] Create internal/auth/scim/users.go
- [ ] Create internal/auth/scim/groups.go
- [ ] Add unit tests
- [ ] Test with Okta

#### Task 4.3: Audit Logging

- [ ] Create common/cloud/audit/interceptor.go
- [ ] Create internal/audit/service.go
- [ ] Create internal/audit/store.go
- [ ] Create S3 archival job
- [ ] Add unit tests

### Phase 5: Console (Weeks 16-22)

#### Task 5.1: Project Setup

- [ ] Initialize Next.js project
- [ ] Configure Tailwind CSS
- [ ] Add shadcn/ui components
- [ ] Setup TanStack Query
- [ ] Generate API client from protos

#### Task 5.2: Authentication

- [ ] Create login page
- [ ] Create SSO callback handler
- [ ] Implement session management
- [ ] Add protected route wrapper

#### Task 5.3: Organization Pages

- [ ] Create organization list page
- [ ] Create organization detail page
- [ ] Create organization settings page
- [ ] Create member management page

#### Task 5.4: Billing Pages

- [ ] Create usage dashboard
- [ ] Create billing overview page
- [ ] Create invoice list page
- [ ] Create payment method page

#### Task 5.5: Settings Pages

- [ ] Create SSO configuration page
- [ ] Create audit log viewer
- [ ] Create API key management page

### Phase 6: Infrastructure (Weeks 20-24)

#### Task 6.1: Staging Environment

- [ ] Deploy staging VPC
- [ ] Deploy staging EKS
- [ ] Deploy staging RDS
- [ ] Deploy staging Redis
- [ ] Deploy all services
- [ ] Configure monitoring
- [ ] Configure alerting

#### Task 6.2: Production Environment

- [ ] Deploy production VPC (multi-AZ)
- [ ] Deploy production EKS (multi-AZ)
- [ ] Deploy production RDS (multi-AZ)
- [ ] Deploy production Redis (cluster)
- [ ] Configure cross-region backup
- [ ] Configure WAF
- [ ] Configure DDoS protection

#### Task 6.3: DR Setup

- [ ] Configure cross-region replication
- [ ] Create DR runbooks
- [ ] Test failover procedure
- [ ] Document recovery procedures

### Phase 7: Testing & Launch (Weeks 22-26)

#### Task 7.1: Testing

- [ ] Complete unit test coverage
- [ ] Complete integration tests
- [ ] Complete E2E tests
- [ ] Run load tests
- [ ] Run security scan
- [ ] Fix all critical issues

#### Task 7.2: Documentation

- [ ] Complete API documentation
- [ ] Complete user guides
- [ ] Complete runbooks
- [ ] Complete architecture docs

#### Task 7.3: Launch

- [ ] Beta launch (invite only)
- [ ] Gather feedback
- [ ] Fix issues
- [ ] General availability launch

<!-- Source: tcld.md -->

## tcld CLI Reference

### Overview

The `tcld` (Temporal Cloud) CLI is the primary tool for managing Temporal Cloud resources. It interacts with the Cloud Ops API.

### Installation

```bash
# Homebrew
brew install temporalio/brew/tcld

# Curl
curl -sSf https://temporal.download/tcld/install.sh | sh
```

### Authentication

```bash
# Interactive login
tcld login

# API Key (CI/CD)
export TEMPORAL_CLOUD_API_KEY=temporal_ak_...
```

### Command Structure

`tcld <resource> <action> [flags]`

### Common Commands

#### Account

```bash
tcld account get
tcld account users list
tcld account region list
```

#### Namespace

```bash
# Create
tcld namespace create \
  --name my-ns \
  --region aws-us-east-1 \
  --retention-days 30 \
  --ca-certificate cert.pem

# Update
tcld namespace update \
  --namespace my-ns \
  --retention-days 7

# List
tcld namespace list

# Certificate Management
tcld namespace certificates add ...
tcld namespace certificates remove ...
```

#### User

```bash
# Invite
tcld user invite \
  --email bob@example.com \
  --account-role developer

# Set Permissions
tcld user namespace-access set \
  --email bob@example.com \
  --namespace my-ns \
  --permission write
```

#### Request

```bash
# Async operation status
tcld request get --request-id <id>
```

### Output Formats

All list commands support JSON output for scripting.

```bash
tcld namespace list --output json | jq '.[].name'
```

### Configuration

Config file stored at `~/.config/tcld/config.yml`.

```yaml
server: saas-api.tmprl.cloud:443
disable_version_check: false
default_namespace: my-ns
```

<!-- Source: terraform.md -->

## Terraform Provider

### Overview

The `temporalio/temporalcloud` provider allows managing Temporal Cloud resources via Infrastructure as Code.

### Configuration

```hcl
terraform {
  required_providers {
    temporalcloud = {
      source  = "temporalio/temporalcloud"
      version = "0.4.0"
    }
  }
}

provider "temporalcloud" {
  api_key = var.temporal_api_key
}
```

### Resources

#### `temporalcloud_namespace`

```hcl
resource "temporalcloud_namespace" "main" {
  name           = "production-ns"
  regions        = ["aws-us-east-1", "aws-us-west-2"]
  retention_days = 30

  certificate {
    certificate = file("certs/ca.pem")
    filters = {
      common_name = "worker.prod.example.com"
    }
  }

  # High Availability
  is_global_namespace = true
}
```

#### `temporalcloud_user`

```hcl
resource "temporalcloud_user" "developer" {
  email        = "dev@example.com"
  account_role = "developer"

  namespace_access {
    namespace_id = temporalcloud_namespace.main.id
    permission   = "write"
  }
}
```

#### `temporalcloud_service_account`

```hcl
resource "temporalcloud_service_account" "ci" {
  name         = "ci-pipeline"
  account_role = "admin"
}
```

#### `temporalcloud_api_key`

```hcl
resource "temporalcloud_api_key" "ci_key" {
  name       = "ci-key"
  owner_id   = temporalcloud_service_account.ci.id
  owner_type = "service_account"
  duration   = "90d"
}
```

### Data Sources

- `temporalcloud_regions`: List supported regions
- `temporalcloud_namespace`: Get namespace details
- `temporalcloud_service_account`: Get SA details

### Best Practices

1. **State Management**: Use S3/GCS backend with locking
2. **Secret Management**: Do not output API keys to console
3. **Module Structure**: Create a module for standard namespace setup (ns + certs + rbac)

<!-- Source: testing.md -->

## Testing Strategy

### Test Pyramid

```
        ┌───────────┐
        │   E2E     │  10%
        │   Tests   │
        ├───────────┤
        │Integration│  30%
        │   Tests   │
        ├───────────┤
        │   Unit    │  60%
        │   Tests   │
        └───────────┘
```

### Test Types

| Type        | Tool                        | Location       | Coverage Target |
| ----------- | --------------------------- | -------------- | --------------- |
| Unit (Go)   | `go test`                   | `*_test.go`    | 80%             |
| Unit (TS)   | Vitest                      | `*.test.ts`    | 80%             |
| Integration | `go test -tags integration` | `integration/` | Critical paths  |
| E2E         | Playwright                  | `e2e/`         | User journeys   |
| Load        | k6                          | `load/`        | Before release  |
| Security    | Trivy, Snyk                 | CI pipeline    | All images      |

### Unit Tests

#### Go Tests

```bash
# Run all unit tests
make test

# Run with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run specific package
go test ./internal/billing/...
```

#### TypeScript Tests

```bash
# Run all tests
pnpm test

# Run with coverage
pnpm test:coverage

# Run in watch mode
pnpm test:watch
```

### Integration Tests

#### Database Tests

```go
//go:build integration

func TestCreateOrganization(t *testing.T) {
    db := setupTestDB(t)
    defer db.Close()

    store := NewOrganizationStore(db)
    org, err := store.Create(ctx, &Organization{Name: "test"})

    require.NoError(t, err)
    require.NotEmpty(t, org.ID)
}
```

#### API Tests

```go
//go:build integration

func TestBillingAPI(t *testing.T) {
    server := setupTestServer(t)
    defer server.Close()

    client := NewBillingClient(server.URL)
    invoice, err := client.GetInvoice(ctx, "inv-123")

    require.NoError(t, err)
    require.Equal(t, "paid", invoice.Status)
}
```

### E2E Tests

#### Playwright Setup

```typescript
// e2e/auth.spec.ts
import { test, expect } from "@playwright/test";

test("user can login via SSO", async ({ page }) => {
  await page.goto("/login");
  await page.click("text=Sign in with SSO");
  await page.fill("[name=email]", "test@example.com");
  await page.click("text=Continue");

  await expect(page).toHaveURL("/dashboard");
});
```

#### Critical User Journeys

- [ ] User signup and onboarding
- [ ] Create namespace
- [ ] View usage dashboard
- [ ] Upgrade subscription
- [ ] Configure SSO
- [ ] Generate API key

### Load Tests

#### k6 Script

```javascript
// load/billing.js
import http from "k6/http";
import { check, sleep } from "k6";

export const options = {
  stages: [
    { duration: "2m", target: 100 },
    { duration: "5m", target: 100 },
    { duration: "2m", target: 0 },
  ],
};

export default function () {
  const res = http.get("https://api.temporal-cloud.io/v1/usage");
  check(res, {
    "status is 200": (r) => r.status === 200,
    "response time < 200ms": (r) => r.timings.duration < 200,
  });
  sleep(1);
}
```

#### Performance Targets

| Endpoint        | P50   | P99   | Max RPS |
| --------------- | ----- | ----- | ------- |
| GET /usage      | 50ms  | 200ms | 1000    |
| POST /namespace | 100ms | 500ms | 100     |
| GET /invoices   | 100ms | 300ms | 500     |

### Security Tests

#### Container Scanning

```yaml
- name: Scan image
  uses: aquasecurity/trivy-action@master
  with:
    image-ref: "${{ env.IMAGE }}"
    severity: "CRITICAL,HIGH"
    exit-code: "1"
```

#### Dependency Scanning

```yaml
- name: Snyk scan
  uses: snyk/actions/golang@master
  with:
    command: test
```

### Test Environments

| Environment | Purpose             | Data                |
| ----------- | ------------------- | ------------------- |
| Local       | Developer testing   | Seed data           |
| CI          | Automated testing   | Ephemeral           |
| Staging     | Pre-prod validation | Sanitized prod copy |
| Production  | Live                | Real                |

### Commands

```bash
# Run all tests
make test-all

# Unit tests only
make test

# Integration tests
make integration-test

# E2E tests
pnpm test:e2e

# Load tests
k6 run load/billing.js

# Security scan
make security-scan
```

<!-- Source: upgrade-policies.md -->

## Upgrade & Update Policies

### Upgrade Types

| Type           | Scope               | Frequency | Downtime       |
| -------------- | ------------------- | --------- | -------------- |
| Patch          | Bug fixes, security | Weekly    | Zero           |
| Minor          | New features        | Monthly   | Zero           |
| Major          | Breaking changes    | Yearly    | Planned window |
| Hotfix         | Critical fix        | As needed | Zero           |
| Infrastructure | Cloud resources     | Monthly   | Zero           |

### Component Upgrade Matrix

| Component       | Current | Upgrade Path                       | Notes                |
| --------------- | ------- | ---------------------------------- | -------------------- |
| Temporal Server | 1.24.x  | 1.24 → 1.25 → 1.26                 | Sequential only      |
| PostgreSQL      | 15.x    | 15 → 16 (major requires migration) | Major = downtime     |
| Kubernetes      | 1.28.x  | 1.28 → 1.29 → 1.30                 | One minor at a time  |
| Go              | 1.22.x  | Any patch                          | Backwards compatible |
| Node.js         | 20.x    | 20 → 22 (LTS to LTS)               | Major = testing      |

### Temporal Server Upgrades

#### Pre-Upgrade Checklist

- [ ] Read release notes for breaking changes
- [ ] Check SDK compatibility matrix
- [ ] Backup database
- [ ] Notify customers (if applicable)
- [ ] Schedule maintenance window (major only)

#### Upgrade Procedure

```bash
# 1. Update Helm values
# values-production.yaml
temporal:
  server:
    image:
      tag: v1.25.0  # Updated from v1.24.2

# 2. Dry-run
helm diff upgrade temporal ./charts/temporal \
  -f values-production.yaml \
  -n temporal-system

# 3. Deploy to staging first
helm upgrade temporal ./charts/temporal \
  -f values-staging.yaml \
  -n temporal-system

# 4. Validate staging (24 hours)
# - Run test workflows
# - Check metrics
# - Verify SDK compatibility

# 5. Production rolling update
helm upgrade temporal ./charts/temporal \
  -f values-production.yaml \
  -n temporal-system
```

#### Rollback Procedure

```bash
# List revisions
helm history temporal -n temporal-system

# Rollback
helm rollback temporal 5 -n temporal-system

# Verify
kubectl get pods -n temporal-system
```

### Database Upgrades

#### Minor Version (e.g., 15.3 → 15.4)

```bash
# RDS handles automatically if auto-upgrade enabled
# Or manual:
aws rds modify-db-instance \
  --db-instance-identifier temporal-cloud-prod \
  --engine-version 15.4 \
  --apply-immediately
```

#### Major Version (e.g., 15 → 16)

**Requires planned downtime:**

1. **Announce maintenance** (1 week ahead)
2. **Create snapshot**
3. **Upgrade replica first**
4. **Test thoroughly**
5. **Upgrade primary during window**
6. **Validate application**
7. **Update connection strings if needed**

```bash
# Create pre-upgrade snapshot
aws rds create-db-snapshot \
  --db-instance-identifier temporal-cloud-prod \
  --db-snapshot-identifier pre-pg16-upgrade

# Upgrade
aws rds modify-db-instance \
  --db-instance-identifier temporal-cloud-prod \
  --engine-version 16.1 \
  --allow-major-version-upgrade \
  --apply-immediately
```

### Kubernetes Upgrades

#### EKS Upgrade

```bash
# 1. Upgrade control plane
aws eks update-cluster-version \
  --name temporal-cloud-prod \
  --kubernetes-version 1.29

# 2. Wait for completion
aws eks describe-update \
  --name temporal-cloud-prod \
  --update-id <update-id>

# 3. Upgrade node groups (rolling)
aws eks update-nodegroup-version \
  --cluster-name temporal-cloud-prod \
  --nodegroup-name general

# 4. Upgrade add-ons
aws eks update-addon \
  --cluster-name temporal-cloud-prod \
  --addon-name vpc-cni \
  --addon-version v1.15.0
```

#### GKE Upgrade

```bash
# 1. Enable maintenance window
gcloud container clusters update temporal-cloud-prod \
  --maintenance-window-start 2025-01-15T02:00:00Z \
  --maintenance-window-end 2025-01-15T06:00:00Z

# 2. Upgrade cluster
gcloud container clusters upgrade temporal-cloud-prod \
  --master --cluster-version 1.29
```

### Dependency Upgrades

#### Go Dependencies

```yaml
# Weekly automated PR via Dependabot
# Review and merge after CI passes

# Security patches: Merge immediately
# Minor updates: Merge weekly
# Major updates: Review breaking changes, test thoroughly
```

#### NPM Dependencies

```yaml
# Weekly automated PR via Renovate
# Group related packages (e.g., all React packages)

# Security patches: Merge immediately
# Minor updates: Merge weekly
# Major updates: Dedicated PR, full QA cycle
```

### Customer SDK Compatibility

#### SDK Support Policy

| Temporal Server | sdk-go | sdk-java | sdk-typescript |
| --------------- | ------ | -------- | -------------- |
| 1.24.x          | 1.26+  | 1.23+    | 1.9+           |
| 1.25.x          | 1.28+  | 1.24+    | 1.10+          |
| 1.26.x          | 1.30+  | 1.25+    | 1.11+          |

#### Customer Communication

When upgrading server:

1. **2 weeks before**: Email customers about upcoming upgrade
2. **1 week before**: Reminder with SDK compatibility notes
3. **Day of**: Status page update
4. **After**: Confirmation email

```markdown
Subject: Temporal Cloud Platform Upgrade - January 20, 2025

Dear Customer,

We will be upgrading the Temporal Cloud platform to version 1.25.0
on January 20, 2025.

**What's changing:**

- New feature: Multi-region namespace support
- Performance improvements
- Bug fixes

**Action required:**
Please ensure your workers are using SDK version 1.28.0 or later
before the upgrade.

**Timeline:**

- Upgrade will be performed between 02:00-04:00 UTC
- Zero downtime expected

Questions? Contact support@temporal.io
```

### Upgrade Windows

#### Preferred Windows

| Region  | Maintenance Window (UTC) |
| ------- | ------------------------ |
| US East | Tuesday 06:00-08:00      |
| US West | Tuesday 08:00-10:00      |
| EU      | Tuesday 02:00-04:00      |
| APAC    | Monday 18:00-20:00       |

#### Blackout Periods

No upgrades during:

- Last week of quarter (revenue recognition)
- Major holidays
- Customer-reported critical periods (Enterprise)
- Active incidents

### Upgrade Automation

#### Automated Patching

```yaml
# .github/workflows/auto-patch.yaml
name: Auto Patch
on:
  schedule:
    - cron: "0 2 * * 2" # Tuesday 2am

jobs:
  patch:
    runs-on: ubuntu-latest
    steps:
      - name: Check for patches
        run: |
          # Check for security patches
          # Create PR if available

      - name: Deploy to staging
        if: steps.check.outputs.has_patches
        run: |
          helm upgrade ...

      - name: Run tests
        run: |
          make integration-test

      - name: Notify team
        run: |
          # Slack notification for review
```

### Upgrade Documentation

Every upgrade must have:

1. **Pre-upgrade runbook**
2. **Upgrade steps**
3. **Validation checklist**
4. **Rollback procedure**
5. **Post-upgrade verification**

Location: `infra/docs/upgrades/`

<!-- Source: upstream-repos.md -->

## Upstream Repository Management

### Overview

Temporal has 190+ repositories at github.com/temporalio. This document defines which ones we fork, mirror, or ignore, and how we manage sync.

### Repository Classification

#### Category 1: FORK (Modify with Cloud Code)

We fork these and add cloud-specific code. Sync daily.

| Repo                               | Purpose          | Cloud Changes                            |
| ---------------------------------- | ---------------- | ---------------------------------------- |
| `temporal`                         | Core server      | Metering interceptors, quota enforcement |
| `cloud-api`                        | Cloud API protos | Extended with billing/org protos         |
| `tcld`                             | Cloud CLI        | Already cloud-native                     |
| `terraform-provider-temporalcloud` | Terraform        | Already cloud-native                     |
| `ui`                               | Web UI           | Embed in console, add cloud features     |

#### Category 2: MIRROR (Use as-is, track releases)

We don't modify these, but we need specific versions for compatibility.

| Repo             | Purpose       | How We Use            |
| ---------------- | ------------- | --------------------- |
| `sdk-go`         | Go SDK        | Workers, client code  |
| `sdk-java`       | Java SDK      | Customer SDK          |
| `sdk-typescript` | TS SDK        | Customer SDK, console |
| `sdk-python`     | Python SDK    | Customer SDK          |
| `sdk-dotnet`     | .NET SDK      | Customer SDK          |
| `sdk-php`        | PHP SDK       | Customer SDK          |
| `sdk-ruby`       | Ruby SDK      | Customer SDK          |
| `api`            | API protos    | Generated clients     |
| `cli`            | Temporal CLI  | Bundled for customers |
| `docker-builds`  | Docker images | Base images           |
| `helm-charts`    | Helm charts   | Deployment base       |
| `samples-*`      | SDK samples   | Customer docs         |

#### Category 3: REFERENCE (Consult, don't use directly)

| Repo            | Purpose          |
| --------------- | ---------------- |
| `documentation` | Docs source      |
| `proposals`     | Design proposals |
| `roadmap`       | Public roadmap   |
| `features`      | Feature flags    |

#### Category 4: IGNORE

Test repos, archived repos, one-off experiments.

### Fork Management

#### Directory Structure

```
github.com/YOUR_ORG/
├── temporal/                    # Fork of temporalio/temporal
├── cloud-api/                   # Fork of temporalio/cloud-api
├── tcld/                        # Fork of temporalio/tcld
├── terraform-provider/          # Fork
├── ui/                          # Fork of temporalio/ui
├── temporal-cloud-platform/     # NEW - Backend services
├── temporal-cloud-console/      # NEW - Web console
├── temporal-cloud-infra/        # NEW - Infrastructure
└── temporal-cloud-sdk-go/       # NEW - SDK extensions
```

#### Branch Structure (Forked Repos)

```
main                 ← Mirrors upstream/main (read-only)
├── cloud/main       ← Production cloud code
├── cloud/staging    ← Staging
├── cloud/develop    ← Development
└── release/v1.x     ← Release branches
```

#### Sync Automation

```yaml
# .github/workflows/sync-upstream.yaml
name: Sync Upstream
on:
  schedule:
    - cron: "0 0 * * *" # Daily at midnight UTC
  workflow_dispatch:

jobs:
  sync:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        repo: [temporal, cloud-api, tcld, ui]
    steps:
      - uses: actions/checkout@v4
        with:
          repository: YOUR_ORG/${{ matrix.repo }}
          ref: main
          fetch-depth: 0
          token: ${{ secrets.SYNC_PAT }}

      - name: Add upstream
        run: |
          git remote add upstream https://github.com/temporalio/${{ matrix.repo }}.git
          git fetch upstream

      - name: Sync main branch
        run: |
          git checkout main
          git reset --hard upstream/main
          git push origin main --force

      - name: Merge to cloud/main
        run: |
          git checkout cloud/main
          git merge main -m "Sync upstream $(date +%Y-%m-%d)" || {
            echo "CONFLICT DETECTED"
            # Create PR for manual resolution
            gh pr create \
              --title "Upstream sync conflict $(date +%Y-%m-%d)" \
              --body "Automatic merge failed. Manual resolution required." \
              --head sync-conflict-$(date +%Y%m%d) \
              --base cloud/main
            exit 1
          }
          git push origin cloud/main

      - name: Notify on failure
        if: failure()
        run: |
          curl -X POST ${{ secrets.SLACK_WEBHOOK }} \
            -d '{"text":"Upstream sync failed for ${{ matrix.repo }}"}'
```

#### Conflict Resolution Policy

1. **Never modify upstream files** - All cloud code in `common/cloud/` or new files
2. **If conflict in upstream file** - We made a mistake; refactor our code
3. **If upstream breaks our code** - Pin to last working version, open issue

#### Release Tracking

Track upstream releases for compatibility:

```yaml
# .github/dependabot.yml
version: 2
updates:
  - package-ecosystem: "gomod"
    directory: "/"
    schedule:
      interval: "weekly"
    allow:
      - dependency-name: "go.temporal.io/*"
```

### SDK Version Matrix

| Cloud Version | temporal Server | sdk-go | sdk-java | sdk-typescript |
| ------------- | --------------- | ------ | -------- | -------------- |
| 1.0.0         | 1.24.x          | 1.29.x | 1.25.x   | 1.10.x         |
| 1.1.0         | 1.25.x          | 1.30.x | 1.26.x   | 1.11.x         |

### Monitoring Upstream

#### Release Notifications

```yaml
# .github/workflows/watch-releases.yaml
name: Watch Upstream Releases
on:
  schedule:
    - cron: "0 */6 * * *" # Every 6 hours

jobs:
  check:
    runs-on: ubuntu-latest
    steps:
      - name: Check temporal releases
        run: |
          LATEST=$(gh release view --repo temporalio/temporal --json tagName -q .tagName)
          CURRENT=$(cat .upstream-versions/temporal)
          if [ "$LATEST" != "$CURRENT" ]; then
            echo "New temporal release: $LATEST"
            # Create issue for upgrade
            gh issue create \
              --title "Upgrade temporal to $LATEST" \
              --body "New upstream release available" \
              --label "upstream-upgrade"
          fi
```

#### Security Advisories

Subscribe to GitHub Security Advisories for all temporalio repos.
Route to #security-alerts Slack channel.

<!-- Source: users.md -->

## User Management

### Account-Level Roles

| Role          | Permissions                                           |
| ------------- | ----------------------------------------------------- |
| Account Owner | Full access including billing, can transfer ownership |
| Global Admin  | Full access except billing, can manage all users      |
| Finance Admin | Billing access only, otherwise read-only              |
| Developer     | Create namespaces, admin on own namespaces            |
| Read-Only     | View only, no modifications                           |

#### Role Capabilities

| Capability            | Owner | Global Admin | Finance | Developer | Read-Only |
| --------------------- | ----- | ------------ | ------- | --------- | --------- |
| Manage billing        | ✅    | ❌           | ✅      | ❌        | ❌        |
| Manage users          | ✅    | ✅           | ❌      | ❌        | ❌        |
| Create namespaces     | ✅    | ✅           | ❌      | ✅        | ❌        |
| Manage all namespaces | ✅    | ✅           | ❌      | ❌        | ❌        |
| View usage            | ✅    | ✅           | ✅      | ✅        | ✅        |
| Configure SSO         | ✅    | ✅           | ❌      | ❌        | ❌        |

### Namespace-Level Permissions

| Permission      | Capabilities                               |
| --------------- | ------------------------------------------ |
| Namespace Admin | Full namespace control, manage permissions |
| Write           | Create, update, delete workflows           |
| Read-Only       | View workflows and history                 |

#### Permission Inheritance

- Global Admin → Namespace Admin on all namespaces
- Developer → Namespace Admin on namespaces they create

### Inviting Users

#### Via Console

1. Go to Users → Invite User
2. Enter email address
3. Select account role
4. Optionally assign namespace permissions
5. Send invitation

#### Via tcld

```bash
tcld user invite \
  --email user@example.com \
  --account-role developer \
  --namespace-permission my-namespace:write
```

#### Via API

```bash
curl -X POST https://api.temporal.io/api/v1/users/invite \
  -H "Authorization: Bearer $API_KEY" \
  -d '{
    "email": "user@example.com",
    "account_role": "developer"
  }'
```

### Managing Roles

#### Update Account Role

```bash
tcld user update \
  --email user@example.com \
  --account-role global_admin
```

#### Update Namespace Permission

```bash
tcld user namespace-access set \
  --email user@example.com \
  --namespace my-namespace \
  --permission write
```

#### Remove Namespace Permission

```bash
tcld user namespace-access remove \
  --email user@example.com \
  --namespace my-namespace
```

### Deleting Users

```bash
tcld user delete --email user@example.com
```

**Effects:**

- User removed from account
- All namespace permissions revoked
- API keys deleted
- Active sessions terminated

### Account Owner

#### Best Practices

- Assign to at least 2 users
- Use personal email (not shared)
- Enable MFA

#### Transferring Ownership

Contact Temporal Support to transfer Account Owner role.

### User Groups

Groups allow bulk permission management.

#### Create Group

```bash
tcld user-group create \
  --name "Backend Team" \
  --description "Backend engineering team"
```

#### Add Members

```bash
tcld user-group member add \
  --group-id grp-123 \
  --email user@example.com
```

#### Set Namespace Permissions

```bash
tcld user-group namespace-access set \
  --group-id grp-123 \
  --namespace my-namespace \
  --permission write
```

### Service Accounts

Machine identities for automation.

#### Create Service Account

```bash
tcld service-account create \
  --name "CI Pipeline" \
  --account-role developer
```

#### Namespace-Scoped Service Account

```bash
tcld service-account create \
  --name "Worker" \
  --namespace my-namespace \
  --permission write
```

### Limits

| Limit             | Value         |
| ----------------- | ------------- |
| Users per account | 300 (default) |
| Service accounts  | 100           |
| User groups       | 50            |
| Members per group | 100           |

<!-- Source: zero-day-response.md -->

## Zero-Day & Vulnerability Response

### Vulnerability Classifications

| Severity | CVSS Score | Response Time | Examples                      |
| -------- | ---------- | ------------- | ----------------------------- |
| Critical | 9.0-10.0   | 24 hours      | RCE, auth bypass, data breach |
| High     | 7.0-8.9    | 7 days        | Privilege escalation, SQLi    |
| Medium   | 4.0-6.9    | 30 days       | XSS, information disclosure   |
| Low      | 0.1-3.9    | 90 days       | Minor issues                  |

### Zero-Day Response Process

#### Phase 1: Detection (0-1 hour)

```
┌─────────────────────────────────────────────────────────────────┐
│                    DETECTION SOURCES                             │
├─────────────────────────────────────────────────────────────────┤
│  • Vendor security advisories (Temporal, AWS, etc.)             │
│  • CVE databases (NVD, MITRE)                                   │
│  • Security mailing lists                                        │
│  • Bug bounty program                                           │
│  • Internal security scans                                       │
│  • Customer reports                                             │
│  • Threat intelligence feeds                                     │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
                    ┌─────────────────┐
                    │  Security Team  │
                    │   Notified      │
                    └─────────────────┘
```

#### Phase 2: Assessment (1-4 hours)

1. **Confirm vulnerability**

   - Reproduce in isolated environment
   - Determine affected versions
   - Assess exploitability

2. **Determine exposure**

   - Which systems are affected?
   - Is it actively exploited?
   - What data is at risk?

3. **Classify severity**
   - Use CVSS scoring
   - Consider business impact
   - Factor in exploitability

#### Phase 3: Containment (Immediate)

```go
// Emergency containment options

// Option 1: Feature flag disable
func containByFeatureFlag(feature string) {
    flags.Disable(feature)
    // Affected code path is no longer reachable
}

// Option 2: WAF rule
func containByWAF(pattern string) {
    waf.AddBlockRule(pattern)
    // Block exploit attempts at edge
}

// Option 3: Network isolation
func containByNetwork(service string) {
    // Remove from load balancer
    removeFromALB(service)
    // Block ingress
    updateSecurityGroup(service, "deny-all")
}

// Option 4: Service shutdown
func emergencyShutdown(service string) {
    kubectl("scale", "deployment", service, "--replicas=0")
}
```

#### Phase 4: Communication (Within 4 hours)

**Internal Communication:**

```
Subject: [SECURITY] Critical Vulnerability - CVE-2025-XXXX

Severity: CRITICAL
Affected: Cloud API v1.2.0-1.2.5
Status: Containment in progress

Summary:
A critical RCE vulnerability has been identified...

Actions:
- Engineering: Patch in progress
- Support: Prepare customer communications
- Leadership: Briefing at 14:00 UTC

DO NOT discuss outside this channel.
```

**External Communication (if customer impact):**

```
Subject: Security Advisory - Immediate Action Required

Dear Customer,

We have identified a security vulnerability affecting
Temporal Cloud. We have already implemented mitigations
and are deploying a permanent fix.

Impact: [Description without exploit details]

Action Required:
- Update SDK to version X.X.X
- Rotate API keys (optional but recommended)

Timeline:
- Issue discovered: [Time]
- Mitigation applied: [Time]
- Fix deployed: [Expected time]

We will provide updates every 2 hours.
```

#### Phase 5: Remediation (24-72 hours)

```bash
# Emergency patch process

# 1. Create hotfix branch
git checkout -b hotfix/CVE-2025-XXXX cloud/main

# 2. Apply fix
# ... code changes ...

# 3. Expedited review (2 senior engineers)
gh pr create --title "[SECURITY] Fix CVE-2025-XXXX"

# 4. Fast-track testing
make test-security
make test-affected-components

# 5. Emergency deployment
helm upgrade --install cloud-platform ./charts \
  --set image.tag=v1.2.6-security \
  --atomic \
  --timeout 10m

# 6. Verify fix
./scripts/verify-vulnerability-fixed.sh
```

#### Phase 6: Recovery (Post-fix)

1. **Verify remediation**

   - Confirm vulnerability is patched
   - Verify no regressions
   - Security scan passes

2. **Restore normal operations**

   - Remove containment measures
   - Re-enable disabled features
   - Resume normal monitoring

3. **Customer follow-up**
   - Send resolution notification
   - Provide details on fix
   - Offer support for any issues

#### Phase 7: Post-Mortem (Within 1 week)

```markdown
## Security Incident Post-Mortem: CVE-2025-XXXX

### Summary

[Brief description of vulnerability]

### Timeline

- 2025-01-15 08:00 - Vulnerability disclosed
- 2025-01-15 08:30 - Security team notified
- 2025-01-15 09:00 - Assessment complete
- 2025-01-15 09:30 - Containment applied
- 2025-01-15 10:00 - Customers notified
- 2025-01-15 18:00 - Patch deployed
- 2025-01-15 19:00 - Verified fixed

### Root Cause

[Technical details]

### Impact

- Systems affected: [List]
- Customers affected: [Count]
- Data exposed: [None/Details]
- Duration: [Time]

### What Went Well

- Fast detection
- Effective containment
- Clear communication

### What Could Be Better

- Earlier detection via scanning
- Faster patch deployment

### Action Items

| Action              | Owner     | Due        |
| ------------------- | --------- | ---------- |
| Add regression test | @eng      | 2025-01-22 |
| Improve scanning    | @security | 2025-01-29 |
| Update runbook      | @ops      | 2025-01-22 |
```

### Proactive Measures

#### Vulnerability Scanning

```yaml
# Daily vulnerability scan
name: Security Scan
on:
  schedule:
    - cron: "0 6 * * *"

jobs:
  scan:
    runs-on: ubuntu-latest
    steps:
      - name: Container scan
        uses: aquasecurity/trivy-action@master
        with:
          image-ref: "${{ env.IMAGE }}"
          severity: "CRITICAL,HIGH"

      - name: Dependency scan
        run: |
          snyk test --severity-threshold=high

      - name: SAST scan
        run: |
          semgrep --config=p/security-audit

      - name: Secret scan
        run: |
          gitleaks detect --source=.
```

#### Threat Intelligence

```yaml
# Monitor for new vulnerabilities
sources:
  - name: NVD
    url: https://nvd.nist.gov/feeds/json/cve/1.1/
    keywords: [temporal, grpc, kubernetes, postgresql, redis]

  - name: GitHub Advisory
    repos: [temporalio/temporal, golang/go, kubernetes/*]

  - name: Vendor Advisories
    urls:
      - https://aws.amazon.com/security/security-bulletins/
      - https://cloud.google.com/support/bulletins
```

#### Bug Bounty Program

```markdown
## Temporal Cloud Bug Bounty

### Scope

- cloud.temporal.io
- api.temporal.io
- \*.tmprl.cloud

### Rewards

| Severity | Bounty         |
| -------- | -------------- |
| Critical | $5,000-$15,000 |
| High     | $1,000-$5,000  |
| Medium   | $500-$1,000    |
| Low      | $100-$500      |

### Rules

- No automated scanning
- No DoS testing
- Report via security@temporal.io
- 90-day disclosure policy
```

### Emergency Contacts

| Role           | Name   | Phone   | Email                |
| -------------- | ------ | ------- | -------------------- |
| Security Lead  | [Name] | [Phone] | security@temporal.io |
| Engineering VP | [Name] | [Phone] | [Email]              |
| Legal          | [Name] | [Phone] | legal@temporal.io    |
| PR             | [Name] | [Phone] | pr@temporal.io       |

### Runbook Location

Detailed security runbooks:
`infra/docs/runbooks/security/`

- `zero-day-response.md`
- `data-breach-response.md`
- `ransomware-response.md`
- `insider-threat-response.md`
