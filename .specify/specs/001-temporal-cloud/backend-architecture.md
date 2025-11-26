# Backend Architecture

## Service Overview

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                              EDGE LAYER                                      │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐        │
│  │   WAF/CDN   │  │  API GW     │  │  Rate Limit │  │  Auth       │        │
│  └─────────────┘  └─────────────┘  └─────────────┘  └─────────────┘        │
└─────────────────────────────────────────────────────────────────────────────┘
                                      │
┌─────────────────────────────────────────────────────────────────────────────┐
│                           CLOUD PLATFORM SERVICES                            │
│  ┌──────────────────┐  ┌──────────────────┐  ┌──────────────────┐          │
│  │   Cloud API      │  │   Billing Svc    │  │   Provisioner    │          │
│  │   (gRPC/REST)    │  │                  │  │                  │          │
│  └──────────────────┘  └──────────────────┘  └──────────────────┘          │
│  ┌──────────────────┐  ┌──────────────────┐  ┌──────────────────┐          │
│  │   Identity Svc   │  │   Audit Svc      │  │   Metrics Svc    │          │
│  │                  │  │                  │  │                  │          │
│  └──────────────────┘  └──────────────────┘  └──────────────────┘          │
└─────────────────────────────────────────────────────────────────────────────┘
                                      │
┌─────────────────────────────────────────────────────────────────────────────┐
│                           TEMPORAL CLUSTER (Per Region)                      │
│  ┌──────────────────┐  ┌──────────────────┐  ┌──────────────────┐          │
│  │   Frontend       │  │   History        │  │   Matching       │          │
│  │   Service        │  │   Service        │  │   Service        │          │
│  └──────────────────┘  └──────────────────┘  └──────────────────┘          │
│  ┌──────────────────┐                                                       │
│  │   Worker         │                                                       │
│  │   Service        │                                                       │
│  └──────────────────┘                                                       │
└─────────────────────────────────────────────────────────────────────────────┘
                                      │
┌─────────────────────────────────────────────────────────────────────────────┐
│                              DATA LAYER                                      │
│  ┌──────────────────┐  ┌──────────────────┐  ┌──────────────────┐          │
│  │   PostgreSQL     │  │   Redis          │  │   Elasticsearch  │          │
│  │   (RDS)          │  │   (ElastiCache)  │  │   (OpenSearch)   │          │
│  └──────────────────┘  └──────────────────┘  └──────────────────┘          │
└─────────────────────────────────────────────────────────────────────────────┘
```

## Cloud Platform Services

### 1. Cloud API Service

**Purpose**: Primary interface for all cloud management operations.

**Tech Stack**:

- Go 1.22+
- gRPC with grpc-gateway (REST transcoding)
- Connect-RPC for browser clients

**API Endpoints**:

| RPC Method           | HTTP                          | Description          |
| -------------------- | ----------------------------- | -------------------- |
| `CreateOrganization` | POST /v1/organizations        | Create new org       |
| `GetOrganization`    | GET /v1/organizations/{id}    | Get org details      |
| `UpdateOrganization` | PATCH /v1/organizations/{id}  | Update org           |
| `CreateNamespace`    | POST /v1/namespaces           | Provision namespace  |
| `GetNamespace`       | GET /v1/namespaces/{id}       | Get namespace        |
| `UpdateNamespace`    | PATCH /v1/namespaces/{id}     | Modify settings      |
| `DeleteNamespace`    | DELETE /v1/namespaces/{id}    | Remove namespace     |
| `ListNamespaces`     | GET /v1/namespaces            | List with pagination |
| `CreateAPIKey`       | POST /v1/api-keys             | Generate key         |
| `ListAPIKeys`        | GET /v1/api-keys              | List keys            |
| `RevokeAPIKey`       | DELETE /v1/api-keys/{id}      | Revoke key           |
| `RotateAPIKey`       | POST /v1/api-keys/{id}/rotate | Rotate key           |
| `GetUsage`           | GET /v1/usage                 | Billing usage        |
| `GetInvoices`        | GET /v1/invoices              | Invoice history      |
| `InviteUser`         | POST /v1/users/invite         | Send invitation      |
| `ListUsers`          | GET /v1/users                 | List org users       |
| `UpdateUserRole`     | PATCH /v1/users/{id}/role     | Change role          |

**Request Flow**:

```
Client Request
      │
      ▼
┌─────────────────┐
│  Load Balancer  │  (ALB/NLB)
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  Auth Intercept │  Validate JWT/API Key
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  Rate Limiter   │  Check org/namespace quotas
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ Audit Intercept │  Log request metadata
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│    Handler      │  Business logic
└────────┬────────┘
         │
    ┌────┴────┐
    ▼         ▼
┌───────┐ ┌───────┐
│  DB   │ │ Cache │
└───────┘ └───────┘
```

**Code Structure**:

```
cloud-api/
├── cmd/
│   └── cloud-api/
│       └── main.go              # Entry point
├── internal/
│   ├── api/
│   │   ├── v1/                  # v1 API handlers
│   │   │   ├── organizations.go
│   │   │   ├── namespaces.go
│   │   │   ├── users.go
│   │   │   ├── apikeys.go
│   │   │   └── billing.go
│   │   └── interceptors/
│   │       ├── auth.go          # JWT/API key validation
│   │       ├── ratelimit.go     # Rate limiting
│   │       ├── audit.go         # Audit logging
│   │       └── recovery.go      # Panic recovery
│   ├── service/                 # Business logic layer
│   │   ├── organization_service.go
│   │   ├── namespace_service.go
│   │   ├── user_service.go
│   │   ├── apikey_service.go
│   │   └── billing_service.go
│   ├── repository/              # Data access layer
│   │   ├── organization_repo.go
│   │   ├── namespace_repo.go
│   │   ├── user_repo.go
│   │   └── postgres.go          # DB connection
│   ├── client/                  # External service clients
│   │   ├── stripe_client.go
│   │   ├── temporal_admin.go
│   │   ├── provisioner_client.go
│   │   └── identity_client.go
│   └── config/
│       └── config.go
├── pkg/
│   └── cloudapi/                # Shared types/utilities
│       ├── errors.go
│       └── pagination.go
└── proto/
    └── cloud/v1/                # Proto definitions
        ├── organizations.proto
        ├── namespaces.proto
        └── billing.proto
```

### 2. Billing Service

**Purpose**: Metering, usage aggregation, invoicing, payment processing.

**Components**:

| Component          | Responsibility                              |
| ------------------ | ------------------------------------------- |
| Metering Agent     | Collect usage events from Temporal clusters |
| Aggregation Worker | Roll up hourly → daily → monthly            |
| Invoice Generator  | Create Stripe invoices                      |
| Dunning Workflow   | Handle failed payments                      |
| Webhook Handler    | Process Stripe events                       |

**Data Flow**:

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│ Temporal Cluster│────▶│  Kafka Topic    │────▶│ Metering Agent  │
│ (usage events)  │     │ (usage.events)  │     │                 │
└─────────────────┘     └─────────────────┘     └────────┬────────┘
                                                         │
                                                         ▼
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│  Stripe API     │◀────│  Aggregation    │◀────│   PostgreSQL    │
│ (usage records) │     │  Worker         │     │ (raw usage)     │
└────────┬────────┘     └─────────────────┘     └─────────────────┘
         │
         ▼
┌─────────────────┐     ┌─────────────────┐
│ Invoice Created │────▶│  Email Service  │
│ (webhook)       │     │                 │
└─────────────────┘     └─────────────────┘
```

**Key Workflows** (Temporal-powered):

```go
// BillingCycleWorkflow - runs monthly per organization
func BillingCycleWorkflow(ctx workflow.Context, orgID string) error {
    // Step 1: Aggregate all usage for the billing period
    var usage UsageSummary
    err := workflow.ExecuteActivity(ctx, AggregateUsageActivity,
        AggregateUsageInput{OrgID: orgID, Period: "monthly"},
    ).Get(ctx, &usage)
    if err != nil {
        return err
    }

    // Step 2: Report usage to Stripe
    err = workflow.ExecuteActivity(ctx, ReportStripeUsageActivity,
        ReportUsageInput{OrgID: orgID, Usage: usage},
    ).Get(ctx, nil)
    if err != nil {
        return err
    }

    // Step 3: Wait for Stripe to finalize invoice (async)
    workflow.Sleep(ctx, 24*time.Hour)

    // Step 4: Send invoice notification
    return workflow.ExecuteActivity(ctx, SendInvoiceEmailActivity, orgID).Get(ctx, nil)
}

// DunningWorkflow - handles failed payment retry
func DunningWorkflow(ctx workflow.Context, invoiceID string) error {
    retrySchedule := []time.Duration{
        3 * 24 * time.Hour,   // Day 3
        7 * 24 * time.Hour,   // Day 7
        14 * 24 * time.Hour,  // Day 14
    }

    for attempt, delay := range retrySchedule {
        workflow.Sleep(ctx, delay)

        // Send reminder email
        _ = workflow.ExecuteActivity(ctx, SendPaymentReminderActivity,
            ReminderInput{InvoiceID: invoiceID, Attempt: attempt + 1},
        ).Get(ctx, nil)

        // Check if paid
        var paid bool
        _ = workflow.ExecuteActivity(ctx, CheckInvoicePaidActivity, invoiceID).Get(ctx, &paid)
        if paid {
            return nil
        }
    }

    // Final: Suspend account
    return workflow.ExecuteActivity(ctx, SuspendAccountActivity, invoiceID).Get(ctx, nil)
}
```

### 3. Provisioner Service

**Purpose**: Creates, updates, and deletes Temporal namespaces across clusters.

**Operations**:

| Operation         | Steps                                                                                                                                                                                                                |
| ----------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| CreateNamespace   | 1. Select target cluster (region, capacity)<br>2. Generate mTLS certificates<br>3. Call Temporal RegisterNamespace API<br>4. Configure retention, limits, search attrs<br>5. Create DNS record<br>6. Update Cloud DB |
| UpdateNamespace   | 1. Validate change (e.g., retention increase only)<br>2. Call Temporal UpdateNamespace API<br>3. Update Cloud DB<br>4. Emit audit event                                                                              |
| DeleteNamespace   | 1. Check deletion protection<br>2. Drain active workflows (optional)<br>3. Call Temporal DeprecateNamespace API<br>4. Remove DNS<br>5. Archive to S3<br>6. Update Cloud DB                                           |
| FailoverNamespace | 1. Verify standby ready<br>2. Fence primary<br>3. Promote standby<br>4. Update DNS<br>5. Verify traffic switched                                                                                                     |

**State Machine**:

```
     ┌──────────────────────────────────────────────────────┐
     │                                                      │
     ▼                                                      │
┌─────────┐    ┌──────────────┐    ┌────────┐    ┌─────────┴───┐
│ PENDING │───▶│ PROVISIONING │───▶│ ACTIVE │───▶│  DELETING   │
└─────────┘    └──────┬───────┘    └────┬───┘    └──────┬──────┘
                      │                 │               │
                      │                 │               ▼
                      │                 │         ┌──────────┐
                      │                 │         │ DELETED  │
                      │                 │         └──────────┘
                      │                 │
                      ▼                 ▼
               ┌──────────┐      ┌───────────┐
               │  FAILED  │      │ SUSPENDED │
               └──────────┘      └───────────┘
```

**Cluster Selection Algorithm**:

```go
func SelectCluster(region string, plan PlanTier) (*Cluster, error) {
    clusters := GetClustersInRegion(region)

    // Filter by plan requirements
    if plan == PlanEnterprise {
        clusters = FilterDedicated(clusters)
    }

    // Sort by available capacity
    sort.Slice(clusters, func(i, j int) bool {
        return clusters[i].AvailableSlots > clusters[j].AvailableSlots
    })

    // Return cluster with most capacity
    if len(clusters) > 0 && clusters[0].AvailableSlots > 0 {
        return clusters[0], nil
    }

    return nil, ErrNoCapacity
}
```

### 4. Identity Service

**Purpose**: User authentication, SSO integration, service account management.

**Capabilities**:

- SAML 2.0 SSO (Okta, Azure AD, Google Workspace, OneLogin)
- SCIM 2.0 provisioning (user/group sync)
- API key lifecycle management
- JWT token issuance and validation

**Authentication Flows**:

**SAML SSO Flow**:

```
User                    Console                 IdP                  Identity Svc
  │                        │                      │                        │
  │──── Login Click ──────▶│                      │                        │
  │                        │──── SAML Request ───▶│                        │
  │                        │                      │                        │
  │◀─────────────────── IdP Login Page ──────────│                        │
  │                        │                      │                        │
  │──── Credentials ──────▶│                      │                        │
  │                        │                      │                        │
  │◀───────────────── SAML Assertion ────────────│                        │
  │                        │                      │                        │
  │                        │──────────────── Validate Assertion ─────────▶│
  │                        │                      │                        │
  │                        │◀───────────────── JWT Token ─────────────────│
  │◀─── Redirect + JWT ────│                      │                        │
```

**API Key Flow**:

```
Client                   Cloud API               Identity Svc
  │                          │                        │
  │── Request + API Key ────▶│                        │
  │                          │── Validate Key ───────▶│
  │                          │                        │── Check: exists, not revoked,
  │                          │                        │   not expired, permissions
  │                          │◀─── Claims ───────────│
  │                          │                        │
  │◀──── Response ──────────│                        │
```

### 5. Audit Service

**Purpose**: Immutable record of all administrative actions for compliance.

**Event Schema**:

```go
type AuditEvent struct {
    ID            string                 `json:"id"`
    Timestamp     time.Time              `json:"timestamp"`

    // Actor
    ActorType     string                 `json:"actor_type"`     // user, service_account, system
    ActorID       string                 `json:"actor_id"`
    ActorEmail    string                 `json:"actor_email,omitempty"`

    // Action
    Action        string                 `json:"action"`         // namespace.create, user.invite
    Result        string                 `json:"result"`         // success, failure

    // Resource
    ResourceType  string                 `json:"resource_type"`  // namespace, user, api_key
    ResourceID    string                 `json:"resource_id"`
    ResourceName  string                 `json:"resource_name"`

    // Context
    OrganizationID string                `json:"organization_id"`
    RequestID      string                `json:"request_id"`
    IPAddress      string                `json:"ip_address"`
    UserAgent      string                `json:"user_agent"`

    // Details (action-specific)
    Details        map[string]any        `json:"details,omitempty"`
}
```

**Storage Strategy**:

- **Hot** (0-90 days): PostgreSQL, full query support
- **Warm** (90-365 days): S3 + Athena, query on demand
- **Cold** (1-7 years): S3 Glacier, compliance retention

### 6. Metrics Service

**Purpose**: Expose customer namespace metrics for Prometheus scraping.

**Architecture**:

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│ Temporal Server │────▶│ Internal Prom   │────▶│ Metrics Service │
│ (all namespaces)│     │ (full metrics)  │     │ (filtered)      │
└─────────────────┘     └─────────────────┘     └────────┬────────┘
                                                         │
                                                         │ /metrics?namespace=ns-123
                                                         ▼
                                                ┌─────────────────┐
                                                │ Customer Prom   │
                                                │ (their metrics) │
                                                └─────────────────┘
```

**Endpoint**: `GET https://metrics.{region}.tmprl.cloud/prometheus/api/v1/query`

**Available Metrics**:

- `temporal_workflow_started_total`
- `temporal_workflow_completed_total`
- `temporal_workflow_failed_total`
- `temporal_activity_execution_latency_seconds`
- `temporal_schedule_action_total`

## Temporal Cluster Services

### Frontend Service

- **Role**: API gateway for all Temporal operations
- **Scaling**: Horizontal (stateless), add replicas
- **Key Functions**: Request authentication, rate limiting, routing to History

### History Service

- **Role**: Workflow execution engine, event sourcing
- **Scaling**: Shard-based (512 to 16k shards)
- **Key Functions**: State machine execution, timer management, replication

### Matching Service

- **Role**: Task queue management and dispatch
- **Scaling**: Horizontal per task queue
- **Key Functions**: Task routing, sync match, worker polling

### Worker Service

- **Role**: Internal system workflows
- **Key Functions**: Retention, archival, replication, batch operations

## Data Models

### Entity Relationships

```
┌─────────────────┐
│  Organization   │
│─────────────────│
│ id              │
│ name            │
│ slug            │
│ plan_tier       │
│ stripe_cust_id  │
└────────┬────────┘
         │
    ┌────┴────┬──────────────┐
    │         │              │
    ▼         ▼              ▼
┌────────┐ ┌────────┐  ┌────────────┐
│Namespace│ │ User   │  │Subscription│
│────────│ │────────│  │────────────│
│id      │ │id      │  │id          │
│name    │ │email   │  │plan        │
│region  │ │role    │  │status      │
│cluster │ │org_id  │  │stripe_sub  │
│org_id  │ └────────┘  └────────────┘
└────────┘
    │
    ▼
┌────────────┐
│  APIKey    │
│────────────│
│id          │
│namespace_id│
│key_hash    │
│permissions │
└────────────┘
```

### Key Tables

See `data-model.md` for complete DDL. Summary:

| Table         | Purpose            | Indexes                                 |
| ------------- | ------------------ | --------------------------------------- |
| organizations | Org metadata       | slug (unique), stripe_customer_id       |
| users         | User accounts      | email (unique), org_id                  |
| namespaces    | Namespace registry | name+org_id (unique), region            |
| api_keys      | API credentials    | key_prefix, namespace_id                |
| subscriptions | Billing plans      | org_id (unique), stripe_subscription_id |
| usage_records | Metering data      | org_id+period, namespace_id+period      |
| audit_events  | Audit log          | org_id+timestamp, actor_id+timestamp    |

## API Authentication & Authorization

### Authentication Methods

| Method  | Header        | Format               | Use Case            |
| ------- | ------------- | -------------------- | ------------------- |
| API Key | Authorization | `Bearer tc_live_...` | Programmatic access |
| JWT     | Authorization | `Bearer eyJhbG...`   | Console sessions    |
| mTLS    | Client Cert   | X.509                | Worker connections  |

### Authorization Model

```go
type Permission string

const (
    PermNamespaceRead   Permission = "namespace:read"
    PermNamespaceWrite  Permission = "namespace:write"
    PermNamespaceAdmin  Permission = "namespace:admin"
    PermOrgRead         Permission = "org:read"
    PermOrgWrite        Permission = "org:write"
    PermOrgAdmin        Permission = "org:admin"
    PermBillingRead     Permission = "billing:read"
    PermBillingWrite    Permission = "billing:write"
)

type Role struct {
    Name        string
    Permissions []Permission
}

var Roles = map[string]Role{
    "owner":     {Permissions: []Permission{/* all */}},
    "admin":     {Permissions: []Permission{PermOrgWrite, PermNamespaceAdmin, PermBillingRead}},
    "developer": {Permissions: []Permission{PermOrgRead, PermNamespaceWrite}},
    "read-only": {Permissions: []Permission{PermOrgRead, PermNamespaceRead}},
}
```

## Error Handling

### gRPC Status Codes

| Code                  | When                   | Client Action      |
| --------------------- | ---------------------- | ------------------ |
| `INVALID_ARGUMENT`    | Validation failed      | Fix request        |
| `NOT_FOUND`           | Resource doesn't exist | Check ID           |
| `ALREADY_EXISTS`      | Duplicate create       | Use existing       |
| `PERMISSION_DENIED`   | Authz failed           | Check permissions  |
| `UNAUTHENTICATED`     | Auth failed            | Re-authenticate    |
| `RESOURCE_EXHAUSTED`  | Rate/quota exceeded    | Backoff + retry    |
| `FAILED_PRECONDITION` | State conflict         | Refresh + retry    |
| `UNAVAILABLE`         | Transient error        | Retry with backoff |
| `INTERNAL`            | Server bug             | Contact support    |

### Error Response Schema

```json
{
  "code": "INVALID_ARGUMENT",
  "message": "namespace name must be 2-63 characters",
  "details": [
    {
      "@type": "type.googleapis.com/google.rpc.BadRequest",
      "field_violations": [
        {
          "field": "name",
          "description": "must be between 2 and 63 characters"
        }
      ]
    }
  ],
  "request_id": "req_abc123"
}
```

## Scalability Architecture

### Horizontal Scaling Matrix

| Service           | Strategy         | Trigger                    |
| ----------------- | ---------------- | -------------------------- |
| Cloud API         | Add replicas     | CPU > 70%, RPS > threshold |
| Billing           | Partition by org | Queue depth > 1000         |
| Provisioner       | Add workers      | Pending tasks > 100        |
| Identity          | Add replicas     | Latency p99 > 200ms        |
| Temporal Frontend | Add replicas     | Connection count           |
| Temporal History  | Add shards       | Shard latency p99 > 500ms  |
| Temporal Matching | Add partitions   | Task backlog > 10k         |

### Database Scaling

- **Read Replicas**: Route read-only queries (ListNamespaces, GetUsage)
- **Connection Pooling**: PgBouncer in transaction mode (6000 connections → 100 DB connections)
- **Partitioning**: Time-based for audit_events, usage_records

## Observability

### Distributed Tracing

All services propagate trace context:

```go
func (s *NamespaceService) Create(ctx context.Context, req *CreateNamespaceRequest) (*Namespace, error) {
    ctx, span := tracer.Start(ctx, "NamespaceService.Create")
    defer span.End()

    span.SetAttributes(
        attribute.String("org_id", req.OrgID),
        attribute.String("namespace.name", req.Name),
        attribute.String("namespace.region", req.Region),
    )

    // ... implementation

    if err != nil {
        span.RecordError(err)
        span.SetStatus(codes.Error, err.Error())
    }

    return ns, err
}
```

### Key Metrics

```yaml
# Cloud API
cloud_api_request_duration_seconds{method, status, org_id}
cloud_api_request_total{method, status}
cloud_api_active_connections{}

# Billing
billing_usage_records_processed_total{status}
billing_invoice_generated_total{status}
billing_stripe_api_latency_seconds{operation}

# Provisioner
provisioner_namespace_operations_total{operation, status}
provisioner_operation_duration_seconds{operation}
provisioner_queue_depth{}
```

### Health Checks

Every service exposes:

| Endpoint   | Purpose    | Behavior                     |
| ---------- | ---------- | ---------------------------- |
| `/health`  | Liveness   | 200 if process running       |
| `/ready`   | Readiness  | 200 if dependencies healthy  |
| `/metrics` | Prometheus | Metrics in Prometheus format |
