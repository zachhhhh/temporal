# Day 0 Bootstrap Sequence

## Problem Statement

The cloud platform runs on Temporal, but it also manages Temporal. This creates a circular dependency that must be resolved during initial setup.

## Architecture

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

## Bootstrap Phases

### Phase 0: Prerequisites

- [ ] AWS accounts created (dev, staging, prod)
- [ ] Domain registered (temporal-cloud.io)
- [ ] SSL certificates provisioned
- [ ] Terraform state backend configured

### Phase 1: Global Infrastructure

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

### Phase 2: Seed Cluster

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

### Phase 3: Cloud Platform Database

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

### Phase 4: Cloud Platform Services

```bash
# 4.1 Deploy cloud platform to seed cluster
cd terraform/seed
terraform apply -target=module.cloud_platform

# 4.2 Verify health
curl https://api.temporal-cloud.io/health
```

### Phase 5: First Customer Region

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

### Phase 6: Verification

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

## Disaster Recovery for Seed Cluster

The seed cluster is critical infrastructure. Special DR considerations:

### Backup Strategy

- Database: Continuous WAL + daily snapshots
- Workflow state: Replicated to S3
- Configuration: Git (infrastructure as code)

### Recovery Procedure

1. Provision new seed cluster from Terraform
2. Restore database from backup
3. Workflows will resume from last checkpoint
4. Update DNS to new seed cluster

### Seed Cluster Monitoring

- 24/7 alerting
- Separate from customer monitoring
- Direct PagerDuty escalation

## Security Considerations

### Seed Cluster Access

- No customer data
- Restricted to platform team
- Separate AWS account
- VPN-only access

### Secrets

- Stored in AWS Secrets Manager
- Rotated every 30 days
- Accessed via IAM roles only

## Runbook Location

Detailed bootstrap runbook:
`temporal-cloud-infra/docs/runbooks/bootstrap.md`
