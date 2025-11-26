# Security

## Security Architecture

### Network Security

- VPC with private subnets for all services
- Security groups with least-privilege rules
- WAF on all public endpoints
- DDoS protection via AWS Shield
- No direct internet access from private subnets

### Data Security

- Encryption at rest: AES-256-GCM
- Encryption in transit: TLS 1.3
- Database encryption: RDS encryption
- S3 encryption: SSE-S3 or SSE-KMS

### Access Control

- IAM roles for all AWS access (no static keys)
- K8s RBAC for cluster access
- SSO for all internal tools
- MFA required for all human access

## Authentication

### User Authentication

- Email/password with MFA
- SAML 2.0 SSO (Okta, Azure AD, etc.)
- Session timeout: 24 hours
- Refresh token rotation

### API Authentication

- API Keys (Bearer token)
- mTLS certificates (namespace access)
- Service account tokens

### mTLS Requirements

- CA certificates: X.509 v3
- Key size: RSA 2048+ or ECDSA P-256+
- Validity: Max 1 year
- Certificate filters for fine-grained access

## Authorization

### Account-Level Roles

| Role          | Permissions                   |
| ------------- | ----------------------------- |
| Account Owner | Full access including billing |
| Global Admin  | Full access except billing    |
| Finance Admin | Billing only                  |
| Developer     | Create namespaces, manage own |
| Read-Only     | View only                     |

### Namespace-Level Permissions

| Permission      | Capabilities           |
| --------------- | ---------------------- |
| Namespace Admin | Full namespace control |
| Write           | CRUD workflows         |
| Read-Only       | View only              |

## Secrets Management

### Storage

- AWS Secrets Manager for all secrets
- External Secrets Operator for K8s injection
- No secrets in environment variables
- No secrets in code or config files

### Rotation

- Database credentials: 90 days
- API keys: User-managed, recommend 90 days
- Service account keys: 90 days
- TLS certificates: Before expiry

## Audit Logging

### Logged Events

- All authentication attempts
- All authorization decisions
- All state-changing operations
- All admin actions

### Log Format

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

### Retention

- Hot storage: 90 days (PostgreSQL)
- Cold storage: 7 years (S3 Glacier)

## Compliance

### SOC 2 Type II

- Status: Planned
- Controls: Security, Availability, Confidentiality
- Annual audit

### GDPR

- Data processing agreements
- Right to erasure support
- Data export capability

### Security Testing

- Penetration testing: Annual
- Vulnerability scanning: Weekly
- Dependency scanning: On every build

## Incident Response

### Severity Levels

| Level | Description          | Response  |
| ----- | -------------------- | --------- |
| SEV1  | Data breach          | Immediate |
| SEV2  | Service compromise   | 1 hour    |
| SEV3  | Vulnerability found  | 24 hours  |
| SEV4  | Security improvement | 1 week    |

### Response Process

1. Detect and alert
2. Assess severity
3. Contain the incident
4. Eradicate the threat
5. Recover services
6. Post-incident review
