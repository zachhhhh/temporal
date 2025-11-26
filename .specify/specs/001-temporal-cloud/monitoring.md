# Monitoring & Alerting

## SLIs and SLOs

| Service   | SLI              | SLO     | Alert Threshold |
| --------- | ---------------- | ------- | --------------- |
| Cloud API | Availability     | 99.9%   | < 99.5% (5m)    |
| Cloud API | Latency P99      | < 200ms | > 500ms (5m)    |
| Console   | Page Load        | < 2s    | > 5s (5m)       |
| Billing   | Invoice Accuracy | 99.99%  | Any error       |
| Metering  | Lag              | < 5m    | > 10m           |

## Monitoring Stack

| Component  | Tool         | Purpose             |
| ---------- | ------------ | ------------------- |
| Metrics    | Prometheus   | Time-series metrics |
| Dashboards | Grafana      | Visualization       |
| Logs       | Loki         | Log aggregation     |
| Traces     | Jaeger       | Distributed tracing |
| Alerts     | Alertmanager | Alert routing       |
| Uptime     | Pingdom      | External monitoring |

## Key Metrics

### Application Metrics

```
# Request rate
sum(rate(http_requests_total[5m])) by (service)

# Error rate
sum(rate(http_requests_total{status=~"5.."}[5m])) / sum(rate(http_requests_total[5m]))

# Latency P99
histogram_quantile(0.99, rate(http_request_duration_seconds_bucket[5m]))
```

### Infrastructure Metrics

```
# CPU usage
avg(rate(container_cpu_usage_seconds_total[5m])) by (pod)

# Memory usage
container_memory_usage_bytes / container_spec_memory_limit_bytes

# Disk usage
(node_filesystem_size_bytes - node_filesystem_free_bytes) / node_filesystem_size_bytes
```

### Business Metrics

```
# Active organizations
count(temporal_cloud_organizations_total)

# Actions per second
sum(rate(temporal_cloud_actions_total[5m]))

# Revenue (monthly)
sum(temporal_cloud_invoice_total_cents) / 100
```

## Dashboards

| Dashboard         | Purpose                     | Audience    |
| ----------------- | --------------------------- | ----------- |
| Platform Overview | High-level health           | On-call     |
| API Performance   | Latency, errors, throughput | Engineering |
| Billing Metrics   | Revenue, usage              | Finance     |
| Infrastructure    | K8s, RDS, Redis             | Platform    |
| Customer Health   | Per-org metrics             | Support     |

## Alert Routing

| Severity      | Response Time     | Channel           |
| ------------- | ----------------- | ----------------- |
| P1 (Critical) | 15 min            | PagerDuty + Slack |
| P2 (High)     | 1 hour            | PagerDuty + Slack |
| P3 (Medium)   | 4 hours           | Slack             |
| P4 (Low)      | Next business day | Email             |

## Alert Definitions

### P1 - Critical

```yaml
- alert: APIDown
  expr: up{job="cloud-api"} == 0
  for: 1m
  labels:
    severity: critical
  annotations:
    summary: "Cloud API is down"

- alert: DatabaseDown
  expr: pg_up == 0
  for: 1m
  labels:
    severity: critical
```

### P2 - High

```yaml
- alert: HighErrorRate
  expr: rate(http_requests_total{status=~"5.."}[5m]) > 0.05
  for: 5m
  labels:
    severity: high

- alert: HighLatency
  expr: histogram_quantile(0.99, rate(http_request_duration_seconds_bucket[5m])) > 0.5
  for: 5m
  labels:
    severity: high
```

### P3 - Medium

```yaml
- alert: DiskSpaceLow
  expr: (node_filesystem_free_bytes / node_filesystem_size_bytes) < 0.2
  for: 15m
  labels:
    severity: medium

- alert: MeteringLag
  expr: temporal_cloud_metering_lag_seconds > 600
  for: 10m
  labels:
    severity: medium
```

## On-Call Rotation

### Schedule

- Primary: 1 week rotation
- Secondary: 1 week rotation (backup)
- Escalation: Engineering Manager → VP Engineering

### Responsibilities

- Acknowledge alerts within SLA
- Investigate and mitigate issues
- Document incidents
- Hand off to next on-call
