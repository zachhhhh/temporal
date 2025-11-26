# Crash Proofing & Reliability

## Circuit Breakers

Protect the system from cascading failures.

```go
// Client-side circuit breaker (e.g., in Cloud API calling Billing)
cb := circuitbreaker.New(circuitbreaker.Options{
    Name:        "billing-service",
    MaxRequests: 100,
    Interval:    5 * time.Second,
    Timeout:     30 * time.Second,
    ReadyToTrip: func(counts circuitbreaker.Counts) bool {
        return counts.ConsecutiveFailures > 5
    },
})

result, err := cb.Execute(func() (interface{}, error) {
    return billingClient.Charge(ctx, req)
})
```

## Bulkheading

Isolate failures to specific components or customers.

### 1. Per-Namespace Isolation

- Shard heavy namespaces to dedicated Task Queues.
- In extreme cases, move noisy neighbor to dedicated History hosts (Isolation Groups).

### 2. Thread Pools

- Separate thread pools for:
  - Critical (API, Membership)
  - User (StartWorkflow, Signal)
  - Background (Replication, Retention)

## Graceful Degradation

When system is overloaded:

1. **Shed Load**: Reject low-priority traffic (e.g., `ListWorkflows`, `Query`) with `503 Service Unavailable`. Keep `StartWorkflow` and `Signal` working.
2. **Disable Features**: Turn off Visibility (search) updates if Elasticsearch is down.
3. **Increase Latency**: Slow down background replication to save IOPS for active traffic.

## Chaos Engineering

### Scheduled Drills

- **Simian Army**: Randomly kill pods (daily).
- **Network Partition**: Block traffic between AZs (monthly).
- **Latency Injection**: Add 100ms lag to DB calls (quarterly).

### "Game Days"

Quarterly manual drills:

1. Kill primary region.
2. Corrupt database WAL.
3. Expire root CA certificate.

## Recovery Oriented Computing

### Fast Restart

- Optimize startup time (< 10s).
- Lazy load caches.

### Stateless Frontends

- Frontend service must be killable at any time with zero impact (drains connections).

### Idempotency

- ALL write APIs must be idempotent.
- Clients must retry indefinitely on transient errors.
