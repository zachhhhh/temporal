# Upgrade & Update Policies

## Upgrade Types

| Type           | Scope               | Frequency | Downtime       |
| -------------- | ------------------- | --------- | -------------- |
| Patch          | Bug fixes, security | Weekly    | Zero           |
| Minor          | New features        | Monthly   | Zero           |
| Major          | Breaking changes    | Yearly    | Planned window |
| Hotfix         | Critical fix        | As needed | Zero           |
| Infrastructure | Cloud resources     | Monthly   | Zero           |

## Component Upgrade Matrix

| Component       | Current | Upgrade Path                       | Notes                |
| --------------- | ------- | ---------------------------------- | -------------------- |
| Temporal Server | 1.24.x  | 1.24 → 1.25 → 1.26                 | Sequential only      |
| PostgreSQL      | 15.x    | 15 → 16 (major requires migration) | Major = downtime     |
| Kubernetes      | 1.28.x  | 1.28 → 1.29 → 1.30                 | One minor at a time  |
| Go              | 1.22.x  | Any patch                          | Backwards compatible |
| Node.js         | 20.x    | 20 → 22 (LTS to LTS)               | Major = testing      |

## Temporal Server Upgrades

### Pre-Upgrade Checklist

- [ ] Read release notes for breaking changes
- [ ] Check SDK compatibility matrix
- [ ] Backup database
- [ ] Notify customers (if applicable)
- [ ] Schedule maintenance window (major only)

### Upgrade Procedure

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

### Rollback Procedure

```bash
# List revisions
helm history temporal -n temporal-system

# Rollback
helm rollback temporal 5 -n temporal-system

# Verify
kubectl get pods -n temporal-system
```

## Database Upgrades

### Minor Version (e.g., 15.3 → 15.4)

```bash
# RDS handles automatically if auto-upgrade enabled
# Or manual:
aws rds modify-db-instance \
  --db-instance-identifier temporal-cloud-prod \
  --engine-version 15.4 \
  --apply-immediately
```

### Major Version (e.g., 15 → 16)

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

## Kubernetes Upgrades

### EKS Upgrade

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

### GKE Upgrade

```bash
# 1. Enable maintenance window
gcloud container clusters update temporal-cloud-prod \
  --maintenance-window-start 2025-01-15T02:00:00Z \
  --maintenance-window-end 2025-01-15T06:00:00Z

# 2. Upgrade cluster
gcloud container clusters upgrade temporal-cloud-prod \
  --master --cluster-version 1.29
```

## Dependency Upgrades

### Go Dependencies

```yaml
# Weekly automated PR via Dependabot
# Review and merge after CI passes

# Security patches: Merge immediately
# Minor updates: Merge weekly
# Major updates: Review breaking changes, test thoroughly
```

### NPM Dependencies

```yaml
# Weekly automated PR via Renovate
# Group related packages (e.g., all React packages)

# Security patches: Merge immediately
# Minor updates: Merge weekly
# Major updates: Dedicated PR, full QA cycle
```

## Customer SDK Compatibility

### SDK Support Policy

| Temporal Server | sdk-go | sdk-java | sdk-typescript |
| --------------- | ------ | -------- | -------------- |
| 1.24.x          | 1.26+  | 1.23+    | 1.9+           |
| 1.25.x          | 1.28+  | 1.24+    | 1.10+          |
| 1.26.x          | 1.30+  | 1.25+    | 1.11+          |

### Customer Communication

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

## Upgrade Windows

### Preferred Windows

| Region  | Maintenance Window (UTC) |
| ------- | ------------------------ |
| US East | Tuesday 06:00-08:00      |
| US West | Tuesday 08:00-10:00      |
| EU      | Tuesday 02:00-04:00      |
| APAC    | Monday 18:00-20:00       |

### Blackout Periods

No upgrades during:

- Last week of quarter (revenue recognition)
- Major holidays
- Customer-reported critical periods (Enterprise)
- Active incidents

## Upgrade Automation

### Automated Patching

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

## Upgrade Documentation

Every upgrade must have:

1. **Pre-upgrade runbook**
2. **Upgrade steps**
3. **Validation checklist**
4. **Rollback procedure**
5. **Post-upgrade verification**

Location: `infra/docs/upgrades/`
