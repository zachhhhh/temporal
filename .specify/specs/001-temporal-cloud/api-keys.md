# API Key Management

## Overview

API keys provide programmatic access to Temporal Cloud without mTLS certificates. They are tied to a user or service account identity.

## Key Types

| Type                    | Owner           | Use Case                 |
| ----------------------- | --------------- | ------------------------ |
| User API Key            | User            | Personal automation, CLI |
| Service Account API Key | Service Account | CI/CD, Workers           |

## Key Lifecycle

### Create

```bash
# Via tcld
tcld apikey create \
  --name "CI Pipeline" \
  --duration 90d

# Output: API key (shown only once)
# temporal_ak_xxxxxxxxxxxxxxxxxxxx
```

### List

```bash
tcld apikey list
```

### Disable/Enable

```bash
tcld apikey disable --id key-123
tcld apikey enable --id key-123
```

### Delete

```bash
tcld apikey delete --id key-123
```

### Rotate

```bash
tcld apikey rotate --id key-123
# Returns new key, old key disabled after 24h
```

## Permissions

API keys inherit permissions from their owner:

- User API Key → User's account role + namespace permissions
- Service Account API Key → Service account's role + permissions

## Security Best Practices

1. **Short expiration**: Set 90-day expiration
2. **Least privilege**: Use service accounts with minimal permissions
3. **Rotate regularly**: Rotate keys every 90 days
4. **Monitor usage**: Review last_used_at regularly
5. **Disable unused**: Disable keys not used in 30+ days

## Using API Keys

### Temporal CLI

```bash
temporal workflow list \
  --api-key temporal_ak_xxxx \
  --address my-namespace.tmprl.cloud:443
```

### SDKs

```go
// Go SDK
client, err := client.Dial(client.Options{
    HostPort:  "my-namespace.tmprl.cloud:443",
    Namespace: "my-namespace",
    Credentials: client.NewAPIKeyStaticCredentials("temporal_ak_xxxx"),
})
```

### tcld

```bash
tcld --api-key temporal_ak_xxxx namespace list
```

### Cloud Ops API

```bash
curl -H "Authorization: Bearer temporal_ak_xxxx" \
  https://api.temporal.io/api/v1/namespaces
```

## Schema

```sql
CREATE TABLE api_keys (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    owner_type VARCHAR(50) NOT NULL,  -- 'user' or 'service_account'
    owner_id UUID NOT NULL,
    key_hash VARCHAR(255) NOT NULL,   -- SHA-256 hash
    key_prefix VARCHAR(10) NOT NULL,  -- First 10 chars for identification
    name VARCHAR(255),
    expires_at TIMESTAMPTZ,
    disabled BOOLEAN DEFAULT FALSE,
    last_used_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_api_keys_owner ON api_keys(owner_type, owner_id);
CREATE INDEX idx_api_keys_prefix ON api_keys(key_prefix);
```

## Rate Limits

| Scope       | Limit            |
| ----------- | ---------------- |
| Per API key | 20 requests/sec  |
| Per account | 200 requests/sec |

## Audit Events

All API key operations are logged:

- `CreateAPIKey`
- `DeleteAPIKey`
- `UpdateAPIKey` (enable/disable)
- `RotateAPIKey`
