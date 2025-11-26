# High Availability

## Redundancy Design

### Component Redundancy

| Component          | Replicas | Spread         | Failover  |
| ------------------ | -------- | -------------- | --------- |
| Temporal Frontend  | 3        | Multi-AZ       | Automatic |
| Temporal History   | 3        | Multi-AZ       | Automatic |
| Temporal Matching  | 3        | Multi-AZ       | Automatic |
| Cloud Platform API | 3        | Multi-AZ       | Automatic |
| Cloud Console      | 2        | Multi-AZ       | Automatic |
| PostgreSQL         | 2        | Multi-AZ (RDS) | Automatic |
| Redis              | 3        | Cluster mode   | Automatic |

### Pod Anti-Affinity

```yaml
affinity:
  podAntiAffinity:
    requiredDuringSchedulingIgnoredDuringExecution:
      - labelSelector:
          matchLabels:
            app: temporal-frontend
        topologyKey: topology.kubernetes.io/zone
```

## Failure Scenarios

| Scenario       | Detection        | Recovery            | RTO   |
| -------------- | ---------------- | ------------------- | ----- |
| Pod failure    | K8s health check | Auto-restart        | < 30s |
| Node failure   | K8s node monitor | Reschedule pods     | < 2m  |
| AZ failure     | ALB health check | Route to healthy AZ | < 1m  |
| Region failure | Route53 health   | DNS failover        | < 5m  |
| DB failure     | RDS monitoring   | Auto-failover       | < 2m  |
| Redis failure  | ElastiCache      | Replica promotion   | < 1m  |

## Health Checks

### Kubernetes Probes

```yaml
livenessProbe:
  httpGet:
    path: /health
    port: 8080
  initialDelaySeconds: 30
  periodSeconds: 10
  failureThreshold: 3

readinessProbe:
  httpGet:
    path: /ready
    port: 8080
  initialDelaySeconds: 5
  periodSeconds: 5
  failureThreshold: 3
```

### ALB Health Checks

- Path: `/health`
- Interval: 10 seconds
- Timeout: 5 seconds
- Healthy threshold: 2
- Unhealthy threshold: 3

## Load Balancing

### External Traffic

- Route53 latency-based routing
- ALB with cross-zone load balancing
- Connection draining: 30 seconds

### Internal Traffic

- Kubernetes Service (ClusterIP)
- gRPC load balancing via headless service

## Namespace High Availability

### Same-Region Replication

- Data replicated within region
- Automatic failover between AZs
- 99.99% SLA

### Multi-Region Replication

- Async replication to secondary region
- Manual or automatic failover
- RPO: < 5 minutes
- RTO: < 15 minutes

### Failover Process

1. Detect primary region failure
2. Promote secondary to primary
3. Update DNS records
4. Notify affected customers
5. Begin data reconciliation
