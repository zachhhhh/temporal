# Data Model

## Entity Relationship Diagram

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

## Schema Definitions

### 001_organizations.sql

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

### 002_subscriptions.sql

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

### 003_usage.sql

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

### 004_invoices.sql

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

### 005_audit.sql

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

### 006_users.sql

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

### 007_service_accounts.sql

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

### 008_api_keys.sql

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

### 009_namespaces.sql

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

### 010_connectivity.sql

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

### 011_nexus.sql

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

### 012_export.sql

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

### 013_credits.sql

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

### 014_user_groups.sql

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
