# Runbooks

## Runbook Index

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

## Incident Response

### Severity Classification

| Level | Description      | Response Time     |
| ----- | ---------------- | ----------------- |
| SEV1  | Complete outage  | 15 minutes        |
| SEV2  | Partial outage   | 1 hour            |
| SEV3  | Degraded service | 4 hours           |
| SEV4  | Minor issue      | Next business day |

### Response Process

1. **Acknowledge** - Claim incident in PagerDuty
2. **Assess** - Determine severity and impact
3. **Communicate** - Update status page, notify stakeholders
4. **Mitigate** - Apply temporary fix if needed
5. **Resolve** - Implement permanent fix
6. **Review** - Post-incident review within 48 hours

---

## Database Maintenance

### Backup Verification

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

### Point-in-Time Recovery

```bash
# Restore to specific time
aws rds restore-db-instance-to-point-in-time \
  --source-db-instance-identifier temporal-cloud-prod \
  --target-db-instance-identifier temporal-cloud-pitr \
  --restore-time 2025-01-01T12:00:00Z
```

---

## Scaling

### Horizontal Scaling (Pods)

```bash
# Scale deployment
kubectl scale deployment cloud-api \
  --replicas=5 \
  -n cloud-platform

# Verify
kubectl get pods -n cloud-platform
```

### Vertical Scaling (Resources)

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

### Database Scaling

```bash
# Modify RDS instance
aws rds modify-db-instance \
  --db-instance-identifier temporal-cloud-prod \
  --db-instance-class db.r6g.2xlarge \
  --apply-immediately
```

---

## Certificate Rotation

### Check Expiration

```bash
# Check certificate expiration
openssl x509 -enddate -noout -in /path/to/cert.pem

# List all certs expiring in 30 days
kubectl get certificates -A -o json | jq '.items[] | select(.status.notAfter | fromdateiso8601 < (now + 2592000))'
```

### Rotate Certificate

```bash
# Trigger cert-manager renewal
kubectl delete secret <cert-secret-name> -n <namespace>

# Verify new certificate
kubectl get certificate <cert-name> -n <namespace>
```

---

## Deployment

### Standard Deployment

```bash
# Deploy to staging
helm upgrade --install cloud-platform ./charts/cloud-platform \
  --namespace cloud-platform \
  --set image.tag=v1.2.3 \
  -f values-staging.yaml

# Verify deployment
kubectl rollout status deployment/cloud-api -n cloud-platform
```

### Canary Deployment

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

## Rollback

### Helm Rollback

```bash
# List releases
helm history cloud-platform -n cloud-platform

# Rollback to previous
helm rollback cloud-platform 1 -n cloud-platform

# Rollback to specific revision
helm rollback cloud-platform 5 -n cloud-platform
```

### Kubernetes Rollback

```bash
# Rollback deployment
kubectl rollout undo deployment/cloud-api -n cloud-platform

# Rollback to specific revision
kubectl rollout undo deployment/cloud-api --to-revision=2 -n cloud-platform
```

---

## Secret Rotation

### Database Credentials

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

### API Keys

```bash
# Rotate via API
curl -X POST https://api.temporal-cloud.io/v1/api-keys/rotate \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -d '{"key_id": "key-123"}'
```
