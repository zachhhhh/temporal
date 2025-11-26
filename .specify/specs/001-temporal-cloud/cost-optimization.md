# Cost Optimization Strategy

## Philosophy

**Pay for value, not waste.** Optimize unit economics without sacrificing reliability.

## Compute Optimization

### 1. Spot Instances / Preemptible VMs

Use Spot instances for stateless workloads:

- **Worker Service**: 100% Spot (stateless, auto-recovering)
- **Frontend Service**: 50% Spot (behind load balancer)
- **Matching Service**: On-Demand (stateful in-memory)
- **History Service**: On-Demand (critical state)

**Savings**: ~70% on compute.

### 2. Graviton (ARM64) Migration

Migrate all services to AWS Graviton3 (m7g, r7g instances).

- Temporal Go code supports ARM64.
- **Savings**: ~20% price-performance improvement.

### 3. Kubernetes Right-Sizing

- **Vertical Pod Autoscaler (VPA)**: Recommend request/limits based on actual usage.
- **Karpenter**: Bin-pack pods onto optimal node sizes.
- **Over-provisioning**: Keep buffer small (10%).

## Storage Optimization

### 1. Tiered Storage (S3)

Move workflow history blobs to cheaper storage tiers.

- **Standard**: Hot data (0-30 days)
- **Intelligent Tiering**: Warm data (30-90 days)
- **Glacier Instant Retrieval**: Cold data (90+ days)

**Savings**: ~40-60% on long-term retention.

### 2. Database pruning

- Aggressively prune `task_executions` and completed workflow records (transfer to S3).
- Use partial indexes to reduce index size.

### 3. EBS Volume Types

- Use **gp3** instead of gp2 (20% cheaper per GB).
- Monitor IOPS usage and provision accurately.

## Networking Optimization

### 1. Keep Traffic Local

- Ensure AZ affinity: Frontend -> History in same AZ preferred.
- Cross-AZ traffic costs $0.01/GB.
- **Topology Aware Hints** in K8s.

### 2. Endpoint Services (PrivateLink)

- Avoid NAT Gateway data processing charges ($0.045/GB).
- Use VPC Endpoints for S3, DynamoDB, etc.

## Observability Cost Control

### 1. Metric Cardinality

- Drop high-cardinality labels (workflowID, runID) from metrics.
- Whitelist only essential metrics for Datadog/Prometheus.

### 2. Log Sampling

- Sample INFO logs at 10%.
- Keep ERROR/WARN at 100%.
- Use structural logging to allow backend filtering.

## Governance

### 1. Budgets & Alerts

- Set AWS Budgets per team/service.
- Alert when forecast exceeds budget by 10%.

### 2. Tagging Policy

Every resource MUST have:

- `CostCenter`
- `Service`
- `Environment`
- `Owner`

Resources without tags are auto-terminated after warning.

## Cost Dashboard

**Unit Metrics**:

- Cost per million Actions
- Cost per GB-hour Storage
- Cost per Active Workflow

If unit cost increases >10%, trigger investigation.
