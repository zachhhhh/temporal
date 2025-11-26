# Capacity Planning

## Capacity Model

### Key Metrics

| Metric           | Unit  | Current | 6mo Forecast | 12mo Forecast |
| ---------------- | ----- | ------- | ------------ | ------------- |
| Organizations    | count | 1,000   | 2,500        | 5,000         |
| Namespaces       | count | 10,000  | 25,000       | 50,000        |
| Actions/sec      | rate  | 100K    | 300K         | 1M            |
| Active Storage   | GB    | 500     | 1,500        | 5,000         |
| Retained Storage | TB    | 10      | 30           | 100           |

### Resource Mapping

| Workload   | Primary Bottleneck | Scaling Factor                    |
| ---------- | ------------------ | --------------------------------- |
| Workflows  | History service    | 1 pod per 1K concurrent workflows |
| Activities | Matching service   | 1 pod per 10K activities/sec      |
| Storage    | PostgreSQL IOPS    | Scale vertically + read replicas  |
| Search     | Elasticsearch      | 1 node per 10K queries/sec        |

## Infrastructure Sizing

### Per-Region Requirements

#### Small (< 10K namespaces)

```yaml
eks:
  nodes: 10 x m6i.xlarge
temporal:
  frontend: 3 replicas
  history: 3 replicas
  matching: 3 replicas
  worker: 2 replicas
database:
  instance: db.r6g.xlarge
  storage: 500GB
  iops: 3000
redis:
  instance: cache.r6g.large
  nodes: 3
```

#### Medium (10K-50K namespaces)

```yaml
eks:
  nodes: 25 x m6i.2xlarge
temporal:
  frontend: 6 replicas
  history: 9 replicas
  matching: 6 replicas
  worker: 4 replicas
database:
  instance: db.r6g.2xlarge
  storage: 2TB
  iops: 10000
  read_replicas: 2
redis:
  instance: cache.r6g.xlarge
  nodes: 6
```

#### Large (50K+ namespaces)

```yaml
eks:
  nodes: 50 x m6i.4xlarge
temporal:
  frontend: 12 replicas
  history: 18 replicas
  matching: 12 replicas
  worker: 8 replicas
database:
  instance: db.r6g.4xlarge
  storage: 10TB
  iops: 30000
  read_replicas: 4
redis:
  instance: cache.r6g.2xlarge
  nodes: 9
```

## Scaling Triggers

### Automatic Scaling

```yaml
# HPA configuration
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: temporal-frontend
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: temporal-frontend
  minReplicas: 3
  maxReplicas: 20
  metrics:
    - type: Resource
      resource:
        name: cpu
        target:
          type: Utilization
          averageUtilization: 70
    - type: Pods
      pods:
        metric:
          name: grpc_requests_per_second
        target:
          type: AverageValue
          averageValue: 1000
```

### Manual Scaling Thresholds

| Metric               | Warning | Action           |
| -------------------- | ------- | ---------------- |
| CPU > 70% sustained  | Alert   | Scale pods       |
| Memory > 80%         | Alert   | Scale pods       |
| DB CPU > 70%         | Alert   | Scale instance   |
| DB connections > 80% | Alert   | Scale instance   |
| Disk > 80%           | Alert   | Increase storage |
| Queue depth > 1000   | Alert   | Scale workers    |

## Capacity Review

### Weekly Review

1. Review resource utilization dashboards
2. Check growth trends
3. Identify hotspots
4. Update forecasts

### Monthly Review

1. Compare actual vs forecast
2. Adjust resource allocations
3. Plan next quarter infrastructure
4. Budget review

### Quarterly Planning

1. Update 6/12 month forecasts
2. Plan major infrastructure changes
3. Reserved instance purchasing
4. Budget approval

## Cost Optimization

### Reserved Instances

| Resource    | On-Demand | 1yr Reserved | 3yr Reserved | Recommendation      |
| ----------- | --------- | ------------ | ------------ | ------------------- |
| EKS nodes   | $0.192/hr | $0.120/hr    | $0.080/hr    | 1yr for stable load |
| RDS         | $0.456/hr | $0.285/hr    | $0.190/hr    | 1yr for prod        |
| ElastiCache | $0.156/hr | $0.097/hr    | $0.065/hr    | 1yr for prod        |

### Right-Sizing

Weekly job to identify:

- Underutilized instances (< 20% CPU)
- Oversized storage
- Unused resources

```bash
# Generate right-sizing report
aws compute-optimizer get-ec2-instance-recommendations \
  --output json > rightsizing-report.json
```

## Disaster Capacity

### Reserve Capacity

Maintain 50% headroom in secondary regions for failover.

| Region       | Primary Capacity | Reserve Capacity |
| ------------ | ---------------- | ---------------- |
| us-east-1    | 100%             | 0%               |
| us-west-2    | 50%              | 50% (failover)   |
| eu-west-1    | 100%             | 0%               |
| eu-central-1 | 50%              | 50% (failover)   |

### Burst Capacity

Use spot instances for non-critical workloads:

- Batch processing
- Analytics
- Development environments

## Monitoring

### Capacity Dashboard

Key panels:

1. Resource utilization by service
2. Growth trends (7d, 30d, 90d)
3. Forecast vs actual
4. Cost per customer
5. Scaling events

### Alerts

```yaml
alerts:
  - name: CapacityWarning
    condition: |
      avg(cpu_utilization) > 70% for 1h
      OR avg(memory_utilization) > 80% for 30m
      OR disk_utilization > 80%
    severity: warning

  - name: CapacityCritical
    condition: |
      avg(cpu_utilization) > 90% for 15m
      OR avg(memory_utilization) > 95% for 15m
      OR disk_utilization > 95%
    severity: critical
```

## Runbook: Emergency Scaling

When capacity limits are hit:

1. **Immediate** (< 5 min)

   - Scale pods via HPA override
   - Enable autoscaling if disabled

2. **Short-term** (< 1 hour)

   - Add nodes to cluster
   - Scale database read replicas

3. **Medium-term** (< 1 day)
   - Scale database instance
   - Increase storage
   - Add new region if needed
