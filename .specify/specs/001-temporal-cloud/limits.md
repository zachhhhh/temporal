# System Limits

## Account Level

| Limit             | Default | Max             | How to Increase           |
| ----------------- | ------- | --------------- | ------------------------- |
| Users             | 300     | Contact support | Support ticket            |
| Namespaces        | 10      | 100 (auto)      | Auto-scales, then support |
| Service Accounts  | 100     | Contact support | Support ticket            |
| API Keys per user | 10      | 50              | Support ticket            |

## Namespace Level

| Limit                    | Default               | Notes                            |
| ------------------------ | --------------------- | -------------------------------- |
| Actions/sec (APS)        | 400                   | Auto-scales based on 7-day usage |
| Requests/sec (RPS)       | 1600                  | Auto-scales                      |
| Operations/sec (OPS)     | 3200                  | Auto-scales                      |
| Schedules RPS            | 10                    | Fixed                            |
| Visibility API RPS       | 30                    | Fixed                            |
| Nexus RPS                | Part of namespace RPS | Shared limit                     |
| Certificates             | 16 or 32KB            | Whichever first                  |
| Retention period         | 1-90 days             | Per namespace config             |
| Concurrent pollers       | 2000                  | Per task queue                   |
| Custom Search Attributes | 100                   | Per namespace                    |

## Workflow Level

| Limit                      | Value      | Notes                          |
| -------------------------- | ---------- | ------------------------------ |
| Identifier length          | 1000 bytes | Unicode may use multiple bytes |
| gRPC message size          | 4 MB       | All endpoints                  |
| Event History transaction  | 4 MB       | Non-configurable               |
| Payload size               | 2 MB       | Single request                 |
| Concurrent Activities      | 2000       | Per workflow                   |
| Concurrent Child Workflows | 2000       | Per workflow                   |
| Concurrent Signals         | 2000       | Per workflow                   |
| Signals per execution      | 10,000     | Total lifetime                 |
| In-flight Updates          | 10         | Concurrent                     |
| Total Updates              | 2000       | In history                     |
| Event History events       | 51,200     | Warning at 10,240              |
| Event History size         | 50 MB      | Warning at 10 MB               |
| Callbacks per workflow     | 32         | Total                          |
| In-flight Nexus ops        | 30         | Concurrent                     |

## Nexus Limits

| Limit                          | Value              |
| ------------------------------ | ------------------ |
| Endpoints per account          | 100                |
| Allowlist entries per endpoint | 100                |
| Operation request timeout      | 10 seconds         |
| Async operation timeout        | 24 hours (default) |

## Worker Versioning Limits

| Limit                     | Value |
| ------------------------- | ----- |
| Deployments per namespace | 100   |
| Versions per deployment   | 100   |
| Task queues per version   | 1000  |

## Rate Limit Behavior

### Throttling Priority

1. External events (Critical) - Never throttled
2. Workflow progress updates
3. Visibility API calls
4. Cloud operations

### When Throttled

- Requests are delayed, not dropped
- High-priority calls never delayed
- Workers may take longer to complete

## Increasing Limits

### Automatic Scaling

- APS, RPS, OPS auto-scale based on 7-day usage
- Never falls below default
- Scales up within minutes of increased usage

### Manual Increase

1. Open support ticket
2. Provide justification and expected usage
3. Temporal reviews and approves
4. Limit increased within 24-48 hours
