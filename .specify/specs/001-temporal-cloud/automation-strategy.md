# Automation Strategy

## Philosophy

**Automate everything that can be automated.** Manual processes introduce latency, inconsistency, and risk.

## Automation Coverage Matrix

| Domain              | What's Automated    | Tooling                 | Trigger                 |
| ------------------- | ------------------- | ----------------------- | ----------------------- |
| **Repo Sync**       | 190+ upstream repos | GitHub Actions + Bot    | Schedule (4h) + Webhook |
| **Dependencies**    | Go/NPM updates      | Dependabot + Renovate   | Weekly PR               |
| **Security Scan**   | CVE detection       | Snyk + Trivy            | On PR + Daily           |
| **Build**           | Compile + Test      | GitHub Actions          | On push                 |
| **Deploy**          | Staging/Prod        | ArgoCD + Helm           | On merge                |
| **DB Migration**    | Schema changes      | golang-migrate          | On deploy               |
| **Cert Rotation**   | mTLS certs          | Cert-manager + Workflow | Before expiry           |
| **Secret Rotation** | API keys, DB creds  | AWS Secrets Manager     | 90-day schedule         |
| **Scaling**         | Pod/Node count      | HPA + Karpenter         | Metric threshold        |
| **Alerting**        | Incident creation   | PagerDuty               | SLO breach              |
| **Backup**          | DB snapshots        | RDS automated           | Daily                   |
| **Log Rotation**    | Archive to S3       | Fluentbit               | TTL-based               |
| **Invoice**         | Monthly billing     | Temporal Workflow       | Monthly schedule        |
| **Dunning**         | Payment retry       | Temporal Workflow       | Invoice unpaid          |

## Repository Automation

### Auto-Discovery of New Repos

```yaml
# .github/workflows/repo-discovery.yaml
name: Discover New Repos
on:
  schedule:
    - cron: "0 */6 * * *" # Every 6 hours

jobs:
  discover:
    runs-on: ubuntu-latest
    steps:
      - name: List upstream repos
        id: upstream
        run: |
          gh api orgs/temporalio/repos --paginate --jq '.[].name' > upstream.txt

      - name: List our forks
        id: ours
        run: |
          gh api orgs/YOUR_ORG/repos --paginate --jq '.[].name' > ours.txt

      - name: Find new repos
        run: |
          comm -23 <(sort upstream.txt) <(sort ours.txt) > new.txt
          if [ -s new.txt ]; then
            echo "NEW_REPOS=$(cat new.txt | tr '\n' ',')" >> $GITHUB_OUTPUT
          fi

      - name: Create fork for new repos
        if: steps.find.outputs.NEW_REPOS != ''
        run: |
          for repo in $(cat new.txt); do
            gh repo fork temporalio/$repo --org YOUR_ORG --clone=false
            # Apply standard settings
            gh api repos/YOUR_ORG/$repo -X PATCH \
              -f has_issues=false \
              -f has_wiki=false \
              -f allow_squash_merge=true
          done

      - name: Notify team
        if: steps.find.outputs.NEW_REPOS != ''
        uses: slackapi/slack-github-action@v1
        with:
          channel-id: "C123456"
          slack-message: "New upstream repos discovered: ${{ steps.find.outputs.NEW_REPOS }}"
```

### Auto-Sync with Upstream

```yaml
# .github/workflows/upstream-sync.yaml
name: Sync Upstream
on:
  schedule:
    - cron: "0 */4 * * *" # Every 4 hours
  workflow_dispatch:

jobs:
  sync:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        repo: [temporal, temporal-ui, sdk-go, sdk-java] # Key repos
    steps:
      - uses: actions/checkout@v4
        with:
          repository: YOUR_ORG/${{ matrix.repo }}
          fetch-depth: 0

      - name: Add upstream
        run: git remote add upstream https://github.com/temporalio/${{ matrix.repo }}.git

      - name: Fetch upstream
        run: git fetch upstream

      - name: Merge upstream
        id: merge
        run: |
          git checkout main
          if git merge upstream/main --no-edit; then
            echo "CONFLICT=false" >> $GITHUB_OUTPUT
          else
            echo "CONFLICT=true" >> $GITHUB_OUTPUT
            git merge --abort
          fi

      - name: Push if clean
        if: steps.merge.outputs.CONFLICT == 'false'
        run: git push origin main

      - name: Create conflict PR
        if: steps.merge.outputs.CONFLICT == 'true'
        run: |
          git checkout -b sync/upstream-$(date +%Y%m%d)
          git merge upstream/main --no-commit || true
          git add -A
          git commit -m "WIP: Sync upstream (conflicts)"
          git push origin HEAD
          gh pr create --title "⚠️ Upstream sync conflict" \
            --body "Manual resolution required" \
            --label "upstream-sync,needs-attention"
```

## CI/CD Automation

### Automated Testing Pipeline

```yaml
# Every PR gets:
# 1. Lint (golangci-lint, eslint)
# 2. Unit tests
# 3. Integration tests (against test Temporal cluster)
# 4. Security scan
# 5. Coverage check (must not decrease)

name: CI
on: [push, pull_request]

jobs:
  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: golangci/golangci-lint-action@v3

  test:
    runs-on: ubuntu-latest
    services:
      postgres:
        image: postgres:15
      redis:
        image: redis:7
      temporal:
        image: temporalio/auto-setup:latest
    steps:
      - run: make test
      - run: make integration-test

  security:
    runs-on: ubuntu-latest
    steps:
      - uses: snyk/actions/golang@master
      - uses: aquasecurity/trivy-action@master
```

### Automated Deployment

```yaml
# ArgoCD Application
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: cloud-platform
spec:
  source:
    repoURL: https://github.com/YOUR_ORG/cloud-platform
    path: deploy/helm
    targetRevision: HEAD
    helm:
      valueFiles:
        - values-production.yaml
  destination:
    server: https://kubernetes.default.svc
    namespace: cloud-platform
  syncPolicy:
    automated:
      prune: true
      selfHeal: true
    syncOptions:
      - CreateNamespace=true
```

## Infrastructure Automation

### Auto-Scaling

```yaml
# Horizontal Pod Autoscaler
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: cloud-api
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: cloud-api
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
          name: http_requests_per_second
        target:
          type: AverageValue
          averageValue: 1000
  behavior:
    scaleUp:
      stabilizationWindowSeconds: 60
      policies:
        - type: Percent
          value: 100
          periodSeconds: 60
    scaleDown:
      stabilizationWindowSeconds: 300
```

### Node Auto-Provisioning (Karpenter)

```yaml
apiVersion: karpenter.sh/v1alpha5
kind: Provisioner
metadata:
  name: default
spec:
  requirements:
    - key: karpenter.sh/capacity-type
      operator: In
      values: ["spot", "on-demand"]
    - key: kubernetes.io/arch
      operator: In
      values: ["arm64"] # Graviton for cost savings
  limits:
    resources:
      cpu: 1000
  ttlSecondsAfterEmpty: 60
  ttlSecondsUntilExpired: 604800 # 7 days
```

## Security Automation

### Automated Secret Rotation

```go
// Temporal workflow for rotating database credentials
func SecretRotationWorkflow(ctx workflow.Context, secretName string) error {
    // 1. Generate new credentials
    var newCreds Credentials
    err := workflow.ExecuteActivity(ctx, GenerateCredentialsActivity).Get(ctx, &newCreds)

    // 2. Update in database (dual-write period)
    err = workflow.ExecuteActivity(ctx, AddDBUserActivity, newCreds).Get(ctx, nil)

    // 3. Update in Secrets Manager
    err = workflow.ExecuteActivity(ctx, UpdateSecretActivity, secretName, newCreds).Get(ctx, nil)

    // 4. Wait for apps to pick up new secret
    workflow.Sleep(ctx, 5*time.Minute)

    // 5. Remove old credentials
    err = workflow.ExecuteActivity(ctx, RemoveOldDBUserActivity).Get(ctx, nil)

    // 6. Schedule next rotation
    return workflow.NewContinueAsNewError(ctx, SecretRotationWorkflow, secretName)
}
```

### Automated Compliance Checks

```yaml
# Daily compliance scan
name: Compliance
on:
  schedule:
    - cron: "0 0 * * *"

jobs:
  check:
    runs-on: ubuntu-latest
    steps:
      - name: Check S3 encryption
        run: |
          aws s3api list-buckets --query 'Buckets[].Name' --output text | \
          while read bucket; do
            encryption=$(aws s3api get-bucket-encryption --bucket $bucket 2>/dev/null || echo "NONE")
            if [ "$encryption" == "NONE" ]; then
              echo "::error::Bucket $bucket is not encrypted"
              exit 1
            fi
          done

      - name: Check RDS encryption
        run: |
          aws rds describe-db-instances --query 'DBInstances[?StorageEncrypted==`false`].DBInstanceIdentifier' | \
          jq -e 'length == 0' || exit 1

      - name: Check public access
        run: |
          # No public S3 buckets
          # No public RDS instances
          # No 0.0.0.0/0 security group rules
```

## Operational Automation

### Automated Incident Response

```yaml
# PagerDuty + Runbook automation
- trigger: SLO breach (availability < 99.9%)
  actions:
    - page on-call
    - create incident ticket
    - start status page update
    - gather diagnostic data automatically

- trigger: High error rate
  actions:
    - check recent deployments
    - auto-rollback if deployment < 30 min old
    - page on-call if not deployment-related
```

### Automated Backup Verification

```go
// Weekly workflow to verify backups are restorable
func BackupVerificationWorkflow(ctx workflow.Context) error {
    // 1. Get latest backup
    var backup BackupInfo
    err := workflow.ExecuteActivity(ctx, GetLatestBackupActivity).Get(ctx, &backup)

    // 2. Restore to test instance
    err = workflow.ExecuteActivity(ctx, RestoreToTestActivity, backup).Get(ctx, nil)

    // 3. Run validation queries
    var valid bool
    err = workflow.ExecuteActivity(ctx, ValidateRestoredDataActivity).Get(ctx, &valid)

    // 4. Cleanup test instance
    err = workflow.ExecuteActivity(ctx, CleanupTestInstanceActivity).Get(ctx, nil)

    // 5. Report results
    if !valid {
        return workflow.ExecuteActivity(ctx, AlertBackupFailureActivity).Get(ctx, nil)
    }

    return nil
}
```

## Billing Automation

### Automated Usage Metering

```go
// Real-time usage tracking
func MeteringWorkflow(ctx workflow.Context, orgID string) error {
    for {
        // Aggregate last hour's usage
        var usage HourlyUsage
        err := workflow.ExecuteActivity(ctx, AggregateHourlyUsageActivity, orgID).Get(ctx, &usage)

        // Report to Stripe
        err = workflow.ExecuteActivity(ctx, ReportStripeUsageActivity, usage).Get(ctx, nil)

        // Check quota
        var overQuota bool
        err = workflow.ExecuteActivity(ctx, CheckQuotaActivity, orgID).Get(ctx, &overQuota)

        if overQuota {
            err = workflow.ExecuteActivity(ctx, SendQuotaWarningActivity, orgID).Get(ctx, nil)
        }

        // Wait for next hour
        workflow.Sleep(ctx, 1*time.Hour)
    }
}
```

## Automation Metrics

Track automation effectiveness:

| Metric                        | Target   | Description               |
| ----------------------------- | -------- | ------------------------- |
| `automation_coverage_percent` | > 95%    | % of operations automated |
| `manual_intervention_count`   | < 5/week | Manual fixes required     |
| `mean_time_to_deploy`         | < 15 min | Code merge to production  |
| `auto_rollback_success_rate`  | > 99%    | Successful auto-rollbacks |
| `secret_rotation_failures`    | 0        | Failed rotations          |
