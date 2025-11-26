# Secrets Management

## Secret Types

| Type           | Examples               | Storage             | Rotation         |
| -------------- | ---------------------- | ------------------- | ---------------- |
| Infrastructure | AWS keys, DB passwords | AWS Secrets Manager | 90 days          |
| Application    | API keys, JWT secrets  | AWS Secrets Manager | 90 days          |
| Customer       | mTLS certs, API keys   | Encrypted in DB     | Customer-managed |
| CI/CD          | Deploy keys, tokens    | GitHub Secrets      | 90 days          |
| Developer      | Personal tokens        | 1Password           | 90 days          |

## Secret Storage

### AWS Secrets Manager

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

### Secret Naming Convention

```
temporal-cloud/{environment}/{service}/{secret-name}

Examples:
temporal-cloud/prod/api/jwt-signing-key
temporal-cloud/prod/db/master-password
temporal-cloud/prod/stripe/api-key
temporal-cloud/staging/api/jwt-signing-key
```

### Kubernetes Integration

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

## Secret Categories

### Database Credentials

| Secret          | Used By           | Rotation       |
| --------------- | ----------------- | -------------- |
| Master password | Migrations, admin | 90 days (auto) |
| App read-write  | Application pods  | 90 days (auto) |
| App read-only   | Reporting         | 90 days (auto) |

### API Keys

| Secret            | Used By         | Rotation                |
| ----------------- | --------------- | ----------------------- |
| Stripe API key    | Billing service | Manual (Stripe rotates) |
| SendGrid API key  | Email service   | 90 days                 |
| PagerDuty API key | Alerting        | 90 days                 |
| Datadog API key   | Monitoring      | 90 days                 |

### Signing Keys

| Secret            | Used By        | Rotation                |
| ----------------- | -------------- | ----------------------- |
| JWT signing key   | Auth service   | 180 days (with overlap) |
| Webhook signing   | Event delivery | 180 days                |
| SAML signing cert | SSO            | 1 year                  |

### Encryption Keys

| Secret              | Used By              | Rotation            |
| ------------------- | -------------------- | ------------------- |
| KMS master key      | All encryption       | Never (AWS managed) |
| Data encryption key | DB column encryption | 1 year              |

## Secret Rotation

### Automatic Rotation (Lambda)

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

### JWT Key Rotation (Overlap Period)

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

## Secret Access

### IAM Policies

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

### Kubernetes RBAC

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

## Secret Injection

### Environment Variables (Preferred for simple secrets)

```yaml
env:
  - name: DATABASE_PASSWORD
    valueFrom:
      secretKeyRef:
        name: db-credentials
        key: password
```

### Volume Mount (For files like certs)

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

### Init Container (For complex setup)

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

## Audit & Monitoring

### Secret Access Logging

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

### Alerts

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

## Emergency Procedures

### Secret Compromise Response

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

### Break Glass Access

For emergency access when normal paths fail:

```bash
# Requires 2 approvals from security team
aws secretsmanager get-secret-value \
  --secret-id temporal-cloud/prod/db/master-password \
  --profile break-glass
```

All break-glass access triggers immediate PagerDuty alert.

## Secrets Checklist

### Before Production

- [ ] All secrets in Secrets Manager (not env vars, not code)
- [ ] Rotation enabled for all rotatable secrets
- [ ] IAM policies follow least privilege
- [ ] Audit logging enabled
- [ ] Alerts configured
- [ ] Break-glass procedure tested
- [ ] Secret naming follows convention
