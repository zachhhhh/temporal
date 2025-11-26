# Namespace Failover Logic

## Problem Statement

Fail over a namespace from one region to another without data loss, maintaining workflow consistency.

## Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                    Primary Region                            │
│  ┌─────────────┐     ┌─────────────┐     ┌─────────────┐   │
│  │   Workers   │────▶│  Temporal   │────▶│  Database   │   │
│  └─────────────┘     └─────────────┘     └─────────────┘   │
│                             │                               │
│                             │ Replication                   │
│                             ▼                               │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                   Secondary Region                           │
│  ┌─────────────┐     ┌─────────────┐     ┌─────────────┐   │
│  │   Workers   │────▶│  Temporal   │────▶│  Database   │   │
│  │  (standby)  │     │  (standby)  │     │  (replica)  │   │
│  └─────────────┘     └─────────────┘     └─────────────┘   │
└─────────────────────────────────────────────────────────────┘
```

## Failover State Machine

```
┌─────────┐    Initiate    ┌─────────────┐    Fence    ┌─────────┐
│ Active  │───────────────▶│  Fencing    │────────────▶│ Syncing │
└─────────┘                └─────────────┘             └─────────┘
                                                            │
                                                            │ Sync Complete
                                                            ▼
┌─────────┐    Complete    ┌─────────────┐    Promote  ┌─────────┐
│ Active  │◀───────────────│  Switching  │◀────────────│ Syncing │
│ (new)   │                └─────────────┘             └─────────┘
```

## Failover Workflow

```go
func NamespaceFailoverWorkflow(ctx workflow.Context, input FailoverInput) error {
    logger := workflow.GetLogger(ctx)

    // Step 1: Validate failover request
    var validation ValidationResult
    err := workflow.ExecuteActivity(ctx, ValidateFailover, input).Get(ctx, &validation)
    if err != nil || !validation.Valid {
        return fmt.Errorf("failover validation failed: %v", validation.Reason)
    }

    // Step 2: Fence the primary region
    // This prevents new writes to the primary
    logger.Info("Fencing primary region", "region", input.SourceRegion)
    err = workflow.ExecuteActivity(ctx, FencePrimaryRegion, FenceInput{
        NamespaceID: input.NamespaceID,
        Region:      input.SourceRegion,
    }).Get(ctx, nil)
    if err != nil {
        return fmt.Errorf("failed to fence primary: %w", err)
    }

    // Step 3: Wait for replication to catch up
    logger.Info("Waiting for replication sync")
    var syncStatus SyncStatus
    for {
        err = workflow.ExecuteActivity(ctx, CheckReplicationLag, input.NamespaceID).Get(ctx, &syncStatus)
        if err != nil {
            return fmt.Errorf("failed to check replication: %w", err)
        }

        if syncStatus.LagSeconds < 1 {
            break
        }

        logger.Info("Replication lag", "seconds", syncStatus.LagSeconds)
        workflow.Sleep(ctx, 5*time.Second)
    }

    // Step 4: Promote secondary to primary
    logger.Info("Promoting secondary region", "region", input.TargetRegion)
    err = workflow.ExecuteActivity(ctx, PromoteSecondary, PromoteInput{
        NamespaceID: input.NamespaceID,
        Region:      input.TargetRegion,
    }).Get(ctx, nil)
    if err != nil {
        // Attempt rollback
        workflow.ExecuteActivity(ctx, UnfencePrimary, input.SourceRegion)
        return fmt.Errorf("failed to promote secondary: %w", err)
    }

    // Step 5: Update DNS/routing
    logger.Info("Updating DNS routing")
    err = workflow.ExecuteActivity(ctx, UpdateDNSRouting, DNSInput{
        NamespaceID: input.NamespaceID,
        NewRegion:   input.TargetRegion,
    }).Get(ctx, nil)
    if err != nil {
        return fmt.Errorf("failed to update DNS: %w", err)
    }

    // Step 6: Notify customers
    workflow.ExecuteActivity(ctx, NotifyFailoverComplete, input)

    // Step 7: Demote old primary to secondary
    workflow.ExecuteActivity(ctx, DemoteToSecondary, DemoteInput{
        NamespaceID: input.NamespaceID,
        Region:      input.SourceRegion,
    })

    logger.Info("Failover complete")
    return nil
}
```

## Fencing Mechanism

### Purpose

Prevent split-brain by ensuring only one region accepts writes.

### Implementation

```go
func FencePrimaryRegion(ctx context.Context, input FenceInput) error {
    // Step 1: Set namespace to read-only mode
    err := setNamespaceReadOnly(ctx, input.NamespaceID, input.Region)
    if err != nil {
        return err
    }

    // Step 2: Wait for in-flight operations to complete
    timeout := time.After(30 * time.Second)
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-timeout:
            return fmt.Errorf("timeout waiting for in-flight operations")
        case <-ticker.C:
            count, err := getInFlightOperations(ctx, input.NamespaceID)
            if err != nil {
                return err
            }
            if count == 0 {
                return nil
            }
        }
    }
}
```

## Replication Sync Check

```go
func CheckReplicationLag(ctx context.Context, namespaceID string) (SyncStatus, error) {
    // Get last written position in primary
    primaryPos, err := getPrimaryPosition(ctx, namespaceID)
    if err != nil {
        return SyncStatus{}, err
    }

    // Get last replicated position in secondary
    secondaryPos, err := getSecondaryPosition(ctx, namespaceID)
    if err != nil {
        return SyncStatus{}, err
    }

    // Calculate lag
    lag := primaryPos.Timestamp.Sub(secondaryPos.Timestamp)

    return SyncStatus{
        PrimaryPosition:   primaryPos,
        SecondaryPosition: secondaryPos,
        LagSeconds:        int(lag.Seconds()),
        InSync:            lag < time.Second,
    }, nil
}
```

## DNS Update

```go
func UpdateDNSRouting(ctx context.Context, input DNSInput) error {
    // Get namespace endpoint
    endpoint := fmt.Sprintf("%s.tmprl.cloud", input.NamespaceID)

    // Update Route53 record to point to new region
    _, err := route53Client.ChangeResourceRecordSets(ctx, &route53.ChangeResourceRecordSetsInput{
        HostedZoneId: aws.String(hostedZoneID),
        ChangeBatch: &types.ChangeBatch{
            Changes: []types.Change{
                {
                    Action: types.ChangeActionUpsert,
                    ResourceRecordSet: &types.ResourceRecordSet{
                        Name: aws.String(endpoint),
                        Type: types.RRTypeA,
                        AliasTarget: &types.AliasTarget{
                            DNSName:              aws.String(regionEndpoints[input.NewRegion]),
                            HostedZoneId:         aws.String(regionHostedZones[input.NewRegion]),
                            EvaluateTargetHealth: aws.Bool(true),
                        },
                    },
                },
            },
        },
    })

    return err
}
```

## Rollback Procedure

If failover fails mid-way:

```go
func RollbackFailover(ctx workflow.Context, input FailoverInput) error {
    // Step 1: Unfence primary (if fenced)
    workflow.ExecuteActivity(ctx, UnfencePrimary, input.SourceRegion)

    // Step 2: Revert DNS (if changed)
    workflow.ExecuteActivity(ctx, UpdateDNSRouting, DNSInput{
        NamespaceID: input.NamespaceID,
        NewRegion:   input.SourceRegion, // Back to original
    })

    // Step 3: Demote secondary (if promoted)
    workflow.ExecuteActivity(ctx, DemoteToSecondary, DemoteInput{
        NamespaceID: input.NamespaceID,
        Region:      input.TargetRegion,
    })

    // Step 4: Notify of rollback
    workflow.ExecuteActivity(ctx, NotifyFailoverRollback, input)

    return nil
}
```

## Timing Expectations

| Phase            | Expected Duration            |
| ---------------- | ---------------------------- |
| Validation       | < 5 seconds                  |
| Fencing          | < 30 seconds                 |
| Replication sync | < 5 minutes (depends on lag) |
| Promotion        | < 10 seconds                 |
| DNS update       | < 60 seconds (propagation)   |
| **Total RTO**    | **< 15 minutes**             |

## Monitoring

### Metrics

- `failover_duration_seconds`
- `replication_lag_seconds`
- `failover_success_total`
- `failover_failure_total`

### Alerts

- Replication lag > 5 minutes
- Failover duration > 30 minutes
- Failover failure
