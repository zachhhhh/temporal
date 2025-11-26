# API Versioning

## Versioning Strategy

We use **URL path versioning** for the Cloud Ops API.

```
https://api.temporal.io/api/v1/namespaces
https://api.temporal.io/api/v2/namespaces
```

## Version Lifecycle

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Alpha     │────▶│   Beta      │────▶│   Stable    │
│   (v1alpha1)│     │   (v1beta1) │     │   (v1)      │
└─────────────┘     └─────────────┘     └─────────────┘
                                               │
                                               │ (12+ months)
                                               ▼
                                        ┌─────────────┐
                                        │  Deprecated │
                                        └─────────────┘
                                               │
                                               │ (6 months)
                                               ▼
                                        ┌─────────────┐
                                        │   Sunset    │
                                        └─────────────┘
```

### Stability Levels

| Level      | Breaking Changes           | Support Period    |
| ---------- | -------------------------- | ----------------- |
| Alpha      | Any time                   | None              |
| Beta       | Major releases only        | 6 months          |
| Stable     | Never (new version needed) | 24 months minimum |
| Deprecated | None                       | 6 months          |

## Breaking vs Non-Breaking Changes

### Non-Breaking (OK)

- Adding new optional fields
- Adding new endpoints
- Adding new enum values
- Relaxing validation (e.g., increasing limit)
- Bug fixes to match documented behavior

### Breaking (Requires New Version)

- Removing or renaming fields
- Changing field types
- Removing endpoints
- Adding required fields
- Changing validation to be more strict
- Changing default values

## Proto Compatibility

### Reserved Fields

```protobuf
message Namespace {
  string id = 1;
  string name = 2;
  // Removed fields - NEVER reuse these numbers
  reserved 3, 4;
  reserved "old_field_name";

  // New fields start from next available
  string region = 5;
}
```

### Backwards Compatible Changes

```protobuf
// Original v1
message CreateNamespaceRequest {
  string name = 1;
  string region = 2;
}

// Updated v1 (compatible)
message CreateNamespaceRequest {
  string name = 1;
  string region = 2;
  optional int32 retention_days = 3; // NEW - optional
}
```

## Version Migration

### Deprecation Notice

When deprecating an API version:

1. **Announcement** (T-6 months)

   - Blog post
   - Email to all customers
   - Console banner

2. **Warning Headers** (T-6 months)

   ```
   Deprecation: true
   Sunset: Sat, 01 Jul 2025 00:00:00 GMT
   Link: <https://docs.temporal.io/migration/v1-to-v2>; rel="deprecation"
   ```

3. **Migration Guide**

   - Document all changes
   - Provide code examples
   - Offer migration tooling if complex

4. **Sunset** (T-0)
   - Return 410 Gone
   - Log attempts for customer outreach

### Migration Support

```go
// Dual-write during migration period
func CreateNamespace(ctx context.Context, req *v2.CreateNamespaceRequest) {
    // Write to v2 storage
    writeV2(req)

    // Also maintain v1 compatibility layer
    if featureFlags.IsEnabled("v1_compat") {
        v1Req := convertV2ToV1(req)
        writeV1(v1Req)
    }
}
```

## Client SDK Versioning

| SDK        | Cloud API v1 | Cloud API v2 |
| ---------- | ------------ | ------------ |
| sdk-go 1.x | ✅           | ❌           |
| sdk-go 2.x | ✅ (compat)  | ✅           |

## Deprecation Policy

### Minimum Support Periods

| API Type           | Stable Support | Deprecation Warning |
| ------------------ | -------------- | ------------------- |
| Public REST/gRPC   | 24 months      | 6 months            |
| Terraform Provider | 18 months      | 6 months            |
| CLI                | 12 months      | 3 months            |
| Internal           | 6 months       | 1 month             |

### Customer Communication

| Channel       | Timing                                            |
| ------------- | ------------------------------------------------- |
| Documentation | Immediately on deprecation                        |
| Email         | 6 months, 3 months, 1 month, 1 week before sunset |
| Console       | Warning banner on affected pages                  |
| API Response  | Deprecation header                                |

## Example: v1 to v2 Migration

### What Changed

| v1                          | v2                     | Change             |
| --------------------------- | ---------------------- | ------------------ |
| `GET /v1/namespaces`        | `GET /v2/namespaces`   | Pagination changed |
| `namespace.settings`        | `namespace.config`     | Field renamed      |
| `retention_period` (string) | `retention_days` (int) | Type changed       |

### Migration Code

```go
// Helper for clients
func MigrateV1ToV2Response(v1 *v1.Namespace) *v2.Namespace {
    return &v2.Namespace{
        Id:            v1.Id,
        Name:          v1.Name,
        Config:        convertSettings(v1.Settings),
        RetentionDays: parseRetention(v1.RetentionPeriod),
    }
}
```

## Testing Compatibility

```go
func TestBackwardsCompatibility(t *testing.T) {
    // Load v1 fixtures
    v1Fixtures := loadFixtures("testdata/v1/*.json")

    for _, fixture := range v1Fixtures {
        // Should still parse with current code
        var req v1.CreateNamespaceRequest
        err := json.Unmarshal(fixture, &req)
        require.NoError(t, err)

        // Should produce valid response
        resp, err := handler.CreateNamespace(ctx, &req)
        require.NoError(t, err)
        require.NotNil(t, resp)
    }
}
```

## Documentation

Each version has its own API reference:

- `docs.temporal.io/api/v1`
- `docs.temporal.io/api/v2`

With clear migration guide linking the two.
