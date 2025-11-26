# Namespace Management

## Overview

A Namespace is the unit of isolation in Temporal Cloud. Each namespace has its own:

- Workflow executions
- Task queues
- Search attributes
- Certificates
- Retention period

## Namespace Properties

| Property            | Description                   | Configurable     |
| ------------------- | ----------------------------- | ---------------- |
| Name                | Unique identifier             | At creation only |
| Region              | Cloud region                  | At creation only |
| Retention           | History retention (1-90 days) | Yes              |
| Certificates        | mTLS CA certificates          | Yes              |
| Search Attributes   | Custom search attributes      | Yes              |
| Tags                | Key-value metadata            | Yes              |
| Deletion Protection | Prevent accidental delete     | Yes              |

## Operations

| Operation     | Console | tcld | API | Terraform |
| ------------- | ------- | ---- | --- | --------- |
| Create        | ✅      | ✅   | ✅  | ✅        |
| Update        | ✅      | ✅   | ✅  | ✅        |
| Delete        | ✅      | ✅   | ✅  | ✅        |
| Failover (HA) | ✅      | ✅   | ✅  | ❌        |

## Creating a Namespace

### Via Console

1. Go to Namespaces → Create Namespace
2. Enter name and select region
3. Configure retention period
4. Add certificates
5. Click Create

### Via tcld

```bash
tcld namespace create \
  --name my-namespace \
  --region aws-us-east-1 \
  --retention-days 7 \
  --ca-certificate ca.pem
```

### Via Terraform

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

## Namespace Naming

### Rules

- 2-63 characters
- Lowercase letters, numbers, hyphens
- Must start with letter
- Must end with letter or number
- Globally unique (account-scoped)

### Best Practices

- Include environment: `myapp-prod`, `myapp-staging`
- Include team: `payments-prod`, `orders-prod`
- Avoid generic names: `test`, `dev`

## Custom Search Attributes

### Built-in Attributes

- `WorkflowId`
- `WorkflowType`
- `StartTime`
- `CloseTime`
- `ExecutionStatus`

### Adding Custom Attributes

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

### Attribute Types

| Type        | Description        | Example       |
| ----------- | ------------------ | ------------- |
| Keyword     | Exact match string | `CustomerId`  |
| Text        | Full-text search   | `Description` |
| Int         | Integer            | `RetryCount`  |
| Double      | Decimal            | `OrderTotal`  |
| Bool        | Boolean            | `IsVIP`       |
| Datetime    | Timestamp          | `DueDate`     |
| KeywordList | List of keywords   | `Tags`        |

## Retention Period

- Minimum: 1 day
- Maximum: 90 days
- Default: 7 days

### Changing Retention

```bash
tcld namespace update \
  --namespace my-namespace \
  --retention-days 30
```

**Note**: Reducing retention does not immediately delete old data.

## Deletion Protection

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

## Tags

Key-value metadata for organization and billing.

```bash
tcld namespace update \
  --namespace my-namespace \
  --tag environment=production \
  --tag team=payments \
  --tag cost-center=12345
```

### Tag Limits

- Max 50 tags per namespace
- Key: 1-128 characters
- Value: 0-256 characters

## Endpoints

### gRPC Endpoint

```
<namespace>.<account-id>.tmprl.cloud:443
```

### Regional Endpoint

```
<region>.region.tmprl.cloud:443
```

## High Availability

### Enable Multi-Region

```bash
tcld namespace update \
  --namespace my-namespace \
  --add-region aws-us-west-2
```

### Failover

```bash
tcld namespace failover \
  --namespace my-namespace \
  --target-region aws-us-west-2
```
