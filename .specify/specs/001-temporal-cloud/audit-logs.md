# Audit Logging

## Overview

Audit logs provide a record of all administrative actions in Temporal Cloud for compliance and security monitoring.

## Supported Events

### Account

| Event                 | Description                         |
| --------------------- | ----------------------------------- |
| ChangeAccountPlanType | Plan upgrade/downgrade              |
| UpdateAccountAPI      | Configure audit logs, observability |

### API Keys

| Event        | Description              |
| ------------ | ------------------------ |
| CreateAPIKey | API key created          |
| DeleteAPIKey | API key deleted          |
| UpdateAPIKey | API key enabled/disabled |

### Namespace

| Event                          | Description                |
| ------------------------------ | -------------------------- |
| CreateNamespaceAPI             | Namespace created          |
| DeleteNamespaceAPI             | Namespace deleted          |
| UpdateNamespaceAPI             | Namespace settings changed |
| FailoverNamespacesAPI          | HA namespace failover      |
| RenameCustomSearchAttributeAPI | Search attribute renamed   |

### Users

| Event                     | Description                  |
| ------------------------- | ---------------------------- |
| CreateUserAPI             | User created                 |
| DeleteUserAPI             | User deleted                 |
| InviteUsersAPI            | User invited                 |
| UpdateUserAPI             | User role changed            |
| SetUserNamespaceAccessAPI | Namespace permission changed |

### Service Accounts

| Event                      | Description             |
| -------------------------- | ----------------------- |
| CreateServiceAccount       | Service account created |
| DeleteServiceAccount       | Service account deleted |
| UpdateServiceAccount       | Service account updated |
| CreateServiceAccountAPIKey | API key created         |

### Nexus

| Event               | Description      |
| ------------------- | ---------------- |
| CreateNexusEndpoint | Endpoint created |
| DeleteNexusEndpoint | Endpoint deleted |
| UpdateNexusEndpoint | Endpoint updated |

### Connectivity

| Event                  | Description  |
| ---------------------- | ------------ |
| CreateConnectivityRule | Rule created |
| DeleteConnectivityRule | Rule deleted |

## Log Format

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

## Export Configuration

### Supported Sinks

- AWS S3
- GCP Cloud Storage
- Datadog
- Splunk

### Configure via Console

1. Go to Settings → Audit Logs
2. Click "Add Sink"
3. Select sink type
4. Configure credentials
5. Test connection
6. Enable

### Configure via tcld

```bash
tcld account audit-log-sink create \
  --type s3 \
  --bucket my-audit-logs \
  --region us-east-1 \
  --role-arn arn:aws:iam::123456789:role/temporal-audit
```

## Retention

| Storage   | Retention | Purpose      |
| --------- | --------- | ------------ |
| Hot (API) | 90 days   | Quick access |
| Cold (S3) | 7 years   | Compliance   |

## Querying Logs

### Via Console

1. Go to Settings → Audit Logs
2. Filter by date, operation, user
3. Export to CSV

### Via API

```bash
curl -H "Authorization: Bearer $API_KEY" \
  "https://api.temporal.io/api/v1/audit-logs?start_time=2025-01-01&end_time=2025-01-31"
```

## Compliance

### SOC 2

- All state-changing operations logged
- Immutable storage
- 7-year retention

### GDPR

- User actions logged
- Data access logged
- Export capability

## Best Practices

1. **Export to SIEM**: Send logs to your security monitoring system
2. **Alert on anomalies**: Set up alerts for unusual activity
3. **Regular review**: Review logs weekly for security
4. **Retain exports**: Keep exported logs beyond 90 days
