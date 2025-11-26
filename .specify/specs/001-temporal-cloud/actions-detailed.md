# Billable Actions - Complete List

## Workflow Actions

| Action                   | Count | Notes                                |
| ------------------------ | ----- | ------------------------------------ |
| Workflow started         | 1     | Via client, Continue-As-New, Child   |
| Workflow reset           | 1     | Actions before reset still count     |
| Timer started            | 1     | Includes implicit SDK timers         |
| Search Attribute upsert  | 1     | Per UpsertSearchAttributes call      |
| Signal sent              | 1     | From client or workflow              |
| Query received           | 1     | Including UI stack trace             |
| Update received          | 1     | Successful or rejected               |
| Side Effect recorded     | 1     | Mutable: only on change              |
| Workflow options updated | 1     | Callback attach, versioning override |

### Non-Billable Workflow Operations

- De-duplicated workflow starts (same Workflow ID)
- De-duplicated updates (same Update ID)
- Search attributes at workflow start
- TemporalChangeVersion search attribute

## Child Workflow Actions

| Action               | Count | Notes                      |
| -------------------- | ----- | -------------------------- |
| Start Child Workflow | 2     | Intent (1) + Execution (1) |

## Activity Actions

| Action               | Count | Notes                    |
| -------------------- | ----- | ------------------------ |
| Activity started     | 1     | Each attempt             |
| Activity retried     | 1     | Each retry attempt       |
| Local Activity batch | 1     | All in one Workflow Task |
| Activity Heartbeat   | 1     | Only if reaches server   |

### Local Activity Details

- All Local Activities in one Workflow Task = 1 Action
- Each Workflow Task heartbeat = 1 additional Action
- Retries after heartbeat = 1 Action (capped at 100)

### Heartbeat Throttling

- SDKs throttle heartbeats (default: 80% of timeout)
- Only heartbeats reaching server are billed
- Local Activities don't have heartbeats

## Schedule Actions

| Action             | Count | Notes                       |
| ------------------ | ----- | --------------------------- |
| Schedule execution | 3     | 2 (schedule) + 1 (workflow) |

## Export Actions

| Action            | Count | Notes                |
| ----------------- | ----- | -------------------- |
| Workflow exported | 1     | Per workflow history |

## Nexus Actions

| Action              | Namespace | Count          |
| ------------------- | --------- | -------------- |
| Operation scheduled | Caller    | 1              |
| Operation canceled  | Caller    | 1              |
| Handler primitives  | Handler   | Normal billing |

### Nexus Notes

- Retries of Nexus Operation itself: Not billed
- Underlying Activities/Workflows: Normal billing on handler namespace

## Action Estimation

### In UI

- Workflow history shows "Billable Actions" column
- Summary at top of workflow view
- Experimental feature - may not include all actions

### Not Included in UI Estimate

- Query
- Activity Heartbeats
- Rejected Updates
- Export
- Schedule overhead

## Cost Optimization Tips

1. **Batch operations**: Use Local Activities for small, fast operations
2. **Reduce signals**: Combine multiple signals into one
3. **Optimize timers**: Use longer durations when possible
4. **Heartbeat wisely**: Increase heartbeat timeout to reduce frequency
5. **Use Continue-As-New**: Prevent unbounded history growth
