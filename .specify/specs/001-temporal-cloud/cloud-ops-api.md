# Cloud Ops API

## Overview

The Cloud Ops API is a gRPC service for managing Temporal Cloud resources programmatically. It is the foundation for the Console, CLI (`tcld`), and Terraform provider.

## Service Definition

### Endpoint

`saas-api.tmprl.cloud:443`

### Authentication

- Bearer Token (API Key)
- mTLS (Service Identity)

## Proto Definitions

### Service: `AccountService`

Manage account-level settings and users.

- `GetAccount`
- `UpdateAccount`
- `ListUsers`
- `InviteUser`
- `UpdateUser`
- `DeleteUser`

### Service: `NamespaceService`

Manage namespaces.

- `CreateNamespace`
- `GetNamespace`
- `UpdateNamespace`
- `DeleteNamespace`
- `ListNamespaces`
- `FailoverNamespace`
- `AddNamespaceRegion`

### Service: `AccessService`

Manage permissions and API keys.

- `CreateAPIKey`
- `ListAPIKeys`
- `UpdateAPIKey`
- `RotateAPIKey`
- `DeleteAPIKey`

### Service: `UsageService`

Retrieve usage data.

- `GetUsageSummary` (Daily/Monthly)
- `GetRequestUsage` (Detailed)

## Rate Limiting

| Scope            | Limit    | Burst |
| ---------------- | -------- | ----- |
| Per Account      | 200 RPS  | 300   |
| Per User         | 20 RPS   | 40    |
| Read Operations  | 1000 RPS | 2000  |
| Write Operations | 50 RPS   | 100   |

## Error Handling

Standard gRPC error codes:

- `INVALID_ARGUMENT` (400): Validation failed
- `UNAUTHENTICATED` (401): Invalid/missing token
- `PERMISSION_DENIED` (403): Insufficient role
- `NOT_FOUND` (404): Resource doesn't exist
- `RESOURCE_EXHAUSTED` (429): Rate limit exceeded
- `UNAVAILABLE` (503): Maintenance or outage

## Idempotency

All write operations support `request_id` field.

- Clients MUST generate a UUID for `request_id`.
- Server stores result of operations for 24 hours.
- Retrying with same `request_id` returns the original result without re-executing.
