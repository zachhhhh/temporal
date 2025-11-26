# Log Management

## Architecture

```
[Pod] -> [Stdout] -> [FluentBit] -> [Kinesis] -> [OpenSearch/Loki] -> [S3 Archive]
```

## Log Levels

| Level     | Meaning                                   | Example                              |
| --------- | ----------------------------------------- | ------------------------------------ |
| **ERROR** | Action failed, human intervention needed. | DB connection lost.                  |
| **WARN**  | Action failed, handled gracefully.        | Workflow retried.                    |
| **INFO**  | Major lifecycle events.                   | Workflow started, Namespace created. |
| **DEBUG** | Granular logic flow (Sampled).            | Request received, state changed.     |

## Structural Logging

JSON format mandatory for parsing.

```json
{
  "level": "info",
  "ts": "2025-01-01T12:00:00Z",
  "caller": "history/handler.go:123",
  "msg": "Processing task",
  "service": "history",
  "shard_id": 42,
  "namespace_id": "ns-123",
  "workflow_id": "wf-abc",
  "trace_id": "1a2b3c"
}
```

## Sensitive Data Redaction

- **PII**: Email, Names, IP (hash or mask).
- **Secrets**: API Keys, Passwords (REDACTED).
- **Payloads**: Workflow inputs/results (NEVER log full payloads).

## Retention Policies

| Log Type       | Hot Storage (Searchable) | Cold Storage (Archive) |
| -------------- | ------------------------ | ---------------------- |
| Access/Audit   | 90 days                  | 7 years                |
| Error/Warn     | 30 days                  | 1 year                 |
| Info (Sampled) | 7 days                   | 90 days                |
| Debug          | 3 days                   | 30 days                |

## Cost Control

### Indexing Rules

- **Index**: `level`, `service`, `trace_id`, `namespace_id`, `error`.
- **Do Not Index**: `stack_trace`, free-text `msg`.

### Sampling

- Use dynamic sampling based on volume. If 1000 logs/sec from one workflow, sample at 0.1%.

## Access Control

- **Devs**: Read access to non-PII logs.
- **SRE**: Full access.
- **Customers**: Access ONLY to their namespace's logs via Export API (not direct access).
