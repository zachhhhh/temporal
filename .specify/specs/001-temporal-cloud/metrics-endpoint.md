# Metrics Endpoint

## Overview

The Metrics Endpoint allows customers to scrape Prometheus-compatible metrics for their own namespaces. This enables them to build custom alerts and dashboards in their own monitoring systems (Datadog, Grafana Cloud, etc.).

## Configuration

### 1. Generate Certificate

Metrics are protected by mTLS. You must generate a dedicated client certificate for scraping.

```bash
tcld namespace metrics-cert create \
  --namespace my-ns \
  --cert metrics-client.pem
```

### 2. Endpoint URL

`https://<namespace>.<account>.tmprl.cloud/prometheus/metrics`

## Scrape Configuration

### Prometheus (`prometheus.yml`)

```yaml
scrape_configs:
  - job_name: "temporal-cloud"
    scrape_interval: 1m
    scheme: https
    static_configs:
      - targets: ["my-ns.a1b2c.tmprl.cloud:443"]
    metrics_path: /prometheus/metrics
    tls_config:
      cert_file: /path/to/metrics-client.pem
      key_file: /path/to/metrics-client.key
      insecure_skip_verify: false
```

### Datadog

Use the OpenMetrics integration.

```yaml
# conf.d/openmetrics.d/conf.yaml
instances:
  - prometheus_url: https://my-ns.a1b2c.tmprl.cloud/prometheus/metrics
    namespace: temporal_cloud
    metrics:
      - temporal_cloud*
    tls_cert: /path/to/cert.pem
    tls_private_key: /path/to/key.pem
```

## Available Metrics

### Workflow Metrics

- `temporal_cloud_workflow_start_count`: Counter
- `temporal_cloud_workflow_success_count`: Counter
- `temporal_cloud_workflow_failed_count`: Counter
- `temporal_cloud_workflow_latency_seconds`: Histogram

### Activity Metrics

- `temporal_cloud_activity_start_count`: Counter
- `temporal_cloud_activity_failed_count`: Counter
- `temporal_cloud_activity_latency_seconds`: Histogram

### Task Queue Metrics

- `temporal_cloud_task_queue_backlog_count`: Gauge
- `temporal_cloud_task_queue_latency_seconds`: Histogram (Schedule-to-Start)

### Resource Metrics

- `temporal_cloud_action_count`: Counter (Billable actions)
- `temporal_cloud_storage_active_bytes`: Gauge
- `temporal_cloud_storage_retained_bytes`: Gauge

## Labels

All metrics include:

- `namespace`
- `operation` (e.g., StartWorkflowExecution)
- `task_queue`
- `workflow_type`
- `activity_type`
- `status` (for counters)

## Cardinality Limits

To protect the system, high-cardinality labels (like Workflow ID) are **NOT** included in this endpoint.
