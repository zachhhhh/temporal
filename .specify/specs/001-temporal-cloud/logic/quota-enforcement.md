# Quota Enforcement Logic

## Problem Statement

Enforce plan-based quotas (actions/sec, storage) across a distributed system with minimal latency impact.

## Architecture

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Request   │────▶│   Quota     │────▶│   Temporal  │
│             │     │  Enforcer   │     │   Server    │
└─────────────┘     └─────────────┘     └─────────────┘
                          │
                          ▼
                    ┌─────────────┐
                    │   Redis     │
                    │   Cache     │
                    └─────────────┘
```

## Token Bucket Algorithm

### Concept

- Each namespace has a "bucket" of tokens
- Tokens replenish at a fixed rate (APS limit)
- Each action consumes one token
- Request rejected if bucket empty

### Implementation

```go
type TokenBucket struct {
    redis     *redis.Client
    namespace string
    limit     int64 // tokens per second
}

func (tb *TokenBucket) Allow(ctx context.Context) (bool, error) {
    now := time.Now().UnixNano()
    key := fmt.Sprintf("quota:%s", tb.namespace)

    // Lua script for atomic check-and-decrement
    script := `
        local key = KEYS[1]
        local limit = tonumber(ARGV[1])
        local now = tonumber(ARGV[2])
        local window = 1000000000 -- 1 second in nanoseconds

        -- Get current bucket state
        local bucket = redis.call('HMGET', key, 'tokens', 'last_update')
        local tokens = tonumber(bucket[1]) or limit
        local last_update = tonumber(bucket[2]) or now

        -- Calculate tokens to add based on time elapsed
        local elapsed = now - last_update
        local tokens_to_add = (elapsed / window) * limit
        tokens = math.min(limit, tokens + tokens_to_add)

        -- Try to consume a token
        if tokens >= 1 then
            tokens = tokens - 1
            redis.call('HMSET', key, 'tokens', tokens, 'last_update', now)
            redis.call('EXPIRE', key, 10)
            return 1
        else
            return 0
        end
    `

    result, err := tb.redis.Eval(ctx, script, []string{key}, tb.limit, now).Int()
    if err != nil {
        // On Redis failure, allow request (fail open)
        return true, err
    }

    return result == 1, nil
}
```

## Distributed Rate Limiting

### Challenge

Multiple frontend instances need coordinated rate limiting.

### Solution: Sliding Window with Redis

```go
type SlidingWindowLimiter struct {
    redis     *redis.Client
    namespace string
    limit     int64
    window    time.Duration
}

func (swl *SlidingWindowLimiter) Allow(ctx context.Context) (bool, int64, error) {
    now := time.Now()
    windowStart := now.Add(-swl.window).UnixNano()
    key := fmt.Sprintf("ratelimit:%s", swl.namespace)

    pipe := swl.redis.Pipeline()

    // Remove old entries
    pipe.ZRemRangeByScore(ctx, key, "0", fmt.Sprintf("%d", windowStart))

    // Count current entries
    countCmd := pipe.ZCard(ctx, key)

    // Add current request (optimistically)
    pipe.ZAdd(ctx, key, &redis.Z{
        Score:  float64(now.UnixNano()),
        Member: fmt.Sprintf("%d:%s", now.UnixNano(), uuid.New().String()),
    })

    // Set expiry
    pipe.Expire(ctx, key, swl.window*2)

    _, err := pipe.Exec(ctx)
    if err != nil {
        return true, 0, err // Fail open
    }

    count := countCmd.Val()
    if count >= swl.limit {
        // Remove the optimistically added entry
        swl.redis.ZRemRangeByRank(ctx, key, -1, -1)
        return false, swl.limit - count, nil
    }

    return true, swl.limit - count - 1, nil
}
```

## Quota Cache

### Local Cache for Performance

```go
type QuotaCache struct {
    cache    *lru.Cache
    store    QuotaStore
    ttl      time.Duration
}

type CachedQuota struct {
    Limit     int64
    ExpiresAt time.Time
}

func (qc *QuotaCache) GetLimit(ctx context.Context, namespace string) (int64, error) {
    // Check local cache
    if cached, ok := qc.cache.Get(namespace); ok {
        cq := cached.(*CachedQuota)
        if time.Now().Before(cq.ExpiresAt) {
            return cq.Limit, nil
        }
    }

    // Fetch from store
    limit, err := qc.store.GetNamespaceLimit(ctx, namespace)
    if err != nil {
        return 0, err
    }

    // Update cache
    qc.cache.Add(namespace, &CachedQuota{
        Limit:     limit,
        ExpiresAt: time.Now().Add(qc.ttl),
    })

    return limit, nil
}
```

## gRPC Interceptor

```go
func QuotaInterceptor(enforcer *QuotaEnforcer) grpc.UnaryServerInterceptor {
    return func(
        ctx context.Context,
        req interface{},
        info *grpc.UnaryServerInfo,
        handler grpc.UnaryHandler,
    ) (interface{}, error) {
        // Extract namespace from request
        namespace := extractNamespace(req)
        if namespace == "" {
            return handler(ctx, req)
        }

        // Check if this is a billable action
        if !isBillableAction(info.FullMethod) {
            return handler(ctx, req)
        }

        // Check quota
        allowed, remaining, err := enforcer.Check(ctx, namespace)
        if err != nil {
            // Log error but allow request (fail open)
            log.Warn("quota check failed", "error", err)
            return handler(ctx, req)
        }

        if !allowed {
            return nil, status.Errorf(
                codes.ResourceExhausted,
                "rate limit exceeded for namespace %s, retry after %v",
                namespace,
                time.Second,
            )
        }

        // Add remaining quota to response headers
        grpc.SetHeader(ctx, metadata.Pairs(
            "x-ratelimit-remaining", fmt.Sprintf("%d", remaining),
        ))

        return handler(ctx, req)
    }
}
```

## Priority Levels

### Request Prioritization

```go
type Priority int

const (
    PriorityCritical Priority = iota // External events, never throttled
    PriorityHigh                     // Workflow progress
    PriorityMedium                   // Visibility API
    PriorityLow                      // Cloud operations
)

func getPriority(method string) Priority {
    switch {
    case isExternalEvent(method):
        return PriorityCritical
    case isWorkflowProgress(method):
        return PriorityHigh
    case isVisibilityAPI(method):
        return PriorityMedium
    default:
        return PriorityLow
    }
}

func (e *QuotaEnforcer) Check(ctx context.Context, namespace string) (bool, int64, error) {
    priority := getPriority(getMethod(ctx))

    // Critical priority never throttled
    if priority == PriorityCritical {
        return true, -1, nil
    }

    // Check rate limit
    return e.limiter.Allow(ctx, namespace)
}
```

## Auto-Scaling Limits

### Dynamic Limit Adjustment

```go
func (e *QuotaEnforcer) AdjustLimits(ctx context.Context) error {
    // Run daily
    namespaces, err := e.store.ListNamespaces(ctx)
    if err != nil {
        return err
    }

    for _, ns := range namespaces {
        // Get 7-day usage
        usage, err := e.store.Get7DayUsage(ctx, ns.ID)
        if err != nil {
            continue
        }

        // Calculate new limit (max of default and 1.5x peak usage)
        peakAPS := usage.PeakActionsPerSecond
        newLimit := max(DefaultAPS, int64(float64(peakAPS)*1.5))

        // Update limit
        e.store.UpdateLimit(ctx, ns.ID, newLimit)
    }

    return nil
}
```
