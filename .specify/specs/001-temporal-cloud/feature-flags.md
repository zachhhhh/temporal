# Feature Flags

## Overview

Feature flags enable safe rollout of new features, A/B testing, and quick rollback without deployment.

## Flag Types

| Type        | Purpose             | Example                          |
| ----------- | ------------------- | -------------------------------- |
| Release     | Gate new features   | `enable_multi_region_namespaces` |
| Ops         | Operational toggles | `enable_rate_limiting`           |
| Experiment  | A/B testing         | `new_billing_ui_variant`         |
| Kill Switch | Emergency disable   | `disable_namespace_creation`     |

## Flag Definition

```yaml
# flags/enable_scim.yaml
name: enable_scim
description: Enable SCIM provisioning for enterprise customers
type: release
owner: platform-team
default: false
targeting:
  - match:
      plan: enterprise
    value: true
  - match:
      org_id:
        - org-beta-1
        - org-beta-2
    value: true
```

## Targeting Rules

### By Plan

```yaml
targeting:
  - match:
      plan: [enterprise, business]
    value: true
```

### By Organization

```yaml
targeting:
  - match:
      org_id: org-123
    value: true
```

### By Percentage

```yaml
targeting:
  - match:
      percentage: 10 # 10% of all users
    value: true
```

### By User Email Domain

```yaml
targeting:
  - match:
      email_domain: temporal.io
    value: true
```

## Backend Usage

### Go SDK

```go
import "github.com/YOUR_ORG/temporal-cloud-platform/pkg/flags"

func CreateNamespace(ctx context.Context, req *CreateNamespaceRequest) error {
    // Check feature flag
    if flags.IsEnabled(ctx, "enable_multi_region") {
        // New multi-region logic
        return createMultiRegionNamespace(ctx, req)
    }

    // Existing single-region logic
    return createSingleRegionNamespace(ctx, req)
}
```

### Flag Client

```go
type FlagClient interface {
    IsEnabled(ctx context.Context, flagName string) bool
    GetVariant(ctx context.Context, flagName string) string
    GetValue(ctx context.Context, flagName string, defaultValue interface{}) interface{}
}

// Implementation with caching
type cachedFlagClient struct {
    store  FlagStore
    cache  *lru.Cache
    ttl    time.Duration
}

func (c *cachedFlagClient) IsEnabled(ctx context.Context, flagName string) bool {
    // Extract targeting context
    org := GetOrgFromContext(ctx)
    user := GetUserFromContext(ctx)

    // Check cache
    cacheKey := fmt.Sprintf("%s:%s:%s", flagName, org.ID, user.ID)
    if cached, ok := c.cache.Get(cacheKey); ok {
        return cached.(bool)
    }

    // Evaluate flag
    result := c.evaluate(flagName, EvalContext{
        OrgID:   org.ID,
        Plan:    org.Plan,
        UserID:  user.ID,
        Email:   user.Email,
    })

    c.cache.Add(cacheKey, result)
    return result
}
```

## Frontend Usage

### React Hook

```typescript
import { useFeatureFlag } from "@/hooks/useFeatureFlags";

function BillingPage() {
  const showNewBillingUI = useFeatureFlag("new_billing_ui");

  if (showNewBillingUI) {
    return <NewBillingDashboard />;
  }

  return <LegacyBillingDashboard />;
}
```

### Flag Provider

```typescript
// Fetch flags on app load
export function FlagProvider({ children }) {
  const { data: flags } = useQuery({
    queryKey: ["feature-flags"],
    queryFn: () => api.getFlags(),
    staleTime: 5 * 60 * 1000, // 5 minutes
  });

  return <FlagContext.Provider value={flags}>{children}</FlagContext.Provider>;
}
```

## Flag Lifecycle

### 1. Create Flag

```bash
# Add flag definition
cat > flags/enable_new_feature.yaml << EOF
name: enable_new_feature
description: New feature description
type: release
owner: @engineer
default: false
EOF

# Deploy flag
make deploy-flags
```

### 2. Implement Behind Flag

```go
if flags.IsEnabled(ctx, "enable_new_feature") {
    newFeatureLogic()
} else {
    existingLogic()
}
```

### 3. Gradual Rollout

```yaml
# Week 1: Internal
targeting:
  - match: { email_domain: temporal.io }
    value: true

# Week 2: Beta customers
targeting:
  - match: { email_domain: temporal.io }
    value: true
  - match: { org_id: [beta-org-1, beta-org-2] }
    value: true

# Week 3: 10% of all
targeting:
  - match: { percentage: 10 }
    value: true

# Week 4: 100%
default: true
```

### 4. Remove Flag

After 100% rollout and monitoring:

1. Remove flag checks from code
2. Delete flag definition
3. Mark as deprecated (7 days)
4. Delete flag

## Operations

### Emergency Kill Switch

```bash
# Disable feature immediately
curl -X PUT https://config.internal/flags/enable_new_billing \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -d '{"default": false, "targeting": []}'
```

### View Flag Status

```bash
# List all flags
curl https://config.internal/flags

# Get specific flag
curl https://config.internal/flags/enable_scim
```

### Audit Trail

All flag changes are logged:

```json
{
  "timestamp": "2025-01-01T12:00:00Z",
  "action": "update",
  "flag": "enable_scim",
  "actor": "admin@temporal.io",
  "before": { "default": false },
  "after": { "default": true }
}
```

## Best Practices

1. **Short-lived flags**: Remove within 30 days of 100% rollout
2. **Clear naming**: `enable_`, `disable_`, `show_`, `use_`
3. **Document**: Always include description and owner
4. **Test both paths**: Unit tests for flag on and off
5. **Monitor**: Add metrics for flag evaluation

## Metrics

```go
// Track flag usage
flagEvaluations.WithLabelValues(flagName, result).Inc()
```

Dashboard:

- Flag evaluation count by flag
- Flag distribution (true vs false)
- Flags enabled over time
