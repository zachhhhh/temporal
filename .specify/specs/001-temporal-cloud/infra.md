# Infrastructure

## Architecture Overview

### Multi-Region Deployment

- 3 regions: us-east-1 (primary), eu-west-1, ap-south-1
- Multi-AZ within each region
- EKS for container orchestration
- RDS PostgreSQL (Multi-AZ)
- ElastiCache Redis (Cluster mode)

## Terraform Modules

| Module     | Path                 | Purpose                   |
| ---------- | -------------------- | ------------------------- |
| vpc        | `modules/vpc`        | Network, subnets, NAT     |
| eks        | `modules/eks`        | Kubernetes cluster        |
| rds        | `modules/rds`        | PostgreSQL                |
| redis      | `modules/redis`      | ElastiCache               |
| monitoring | `modules/monitoring` | Prometheus, Grafana       |
| alb        | `modules/alb`        | Application Load Balancer |
| waf        | `modules/waf`        | Web Application Firewall  |

## Environments

| Env     | Account         | Region       | Size   |
| ------- | --------------- | ------------ | ------ |
| dev     | dev-account     | us-east-1    | Small  |
| staging | staging-account | us-east-1    | Medium |
| prod    | prod-account    | Multi-region | Large  |

## Resource Sizing

### Development

```yaml
eks:
  node_groups:
    - name: general
      instance_types: [t3.medium]
      min_size: 2
      max_size: 4
rds:
  instance_class: db.t3.medium
  storage: 100GB
redis:
  node_type: cache.t3.medium
  num_cache_nodes: 1
```

### Staging

```yaml
eks:
  node_groups:
    - name: general
      instance_types: [t3.large]
      min_size: 3
      max_size: 6
rds:
  instance_class: db.r6g.large
  storage: 500GB
  multi_az: true
redis:
  node_type: cache.r6g.large
  num_cache_nodes: 2
```

### Production

```yaml
eks:
  node_groups:
    - name: general
      instance_types: [m6i.xlarge]
      min_size: 6
      max_size: 20
    - name: temporal
      instance_types: [m6i.2xlarge]
      min_size: 3
      max_size: 10
rds:
  instance_class: db.r6g.2xlarge
  storage: 2TB
  multi_az: true
  read_replicas: 2
redis:
  node_type: cache.r6g.xlarge
  num_cache_nodes: 3
  cluster_mode: true
```

## Network Design

### VPC CIDR Allocation

| Environment | VPC CIDR     | Public Subnets | Private Subnets | Database Subnets |
| ----------- | ------------ | -------------- | --------------- | ---------------- |
| dev         | 10.0.0.0/16  | 10.0.0.0/20    | 10.0.16.0/20    | 10.0.32.0/20     |
| staging     | 10.1.0.0/16  | 10.1.0.0/20    | 10.1.16.0/20    | 10.1.32.0/20     |
| prod-us     | 10.10.0.0/16 | 10.10.0.0/20   | 10.10.16.0/20   | 10.10.32.0/20    |
| prod-eu     | 10.20.0.0/16 | 10.20.0.0/20   | 10.20.16.0/20   | 10.20.32.0/20    |
| prod-ap     | 10.30.0.0/16 | 10.30.0.0/20   | 10.30.16.0/20   | 10.30.32.0/20    |

### Security Groups

| Name     | Inbound            | Outbound   |
| -------- | ------------------ | ---------- |
| alb-sg   | 443 from 0.0.0.0/0 | All to VPC |
| eks-sg   | All from alb-sg    | All        |
| rds-sg   | 5432 from eks-sg   | None       |
| redis-sg | 6379 from eks-sg   | None       |

## Kubernetes Resources

### Namespaces

- `temporal-system` - Temporal server components
- `cloud-platform` - Cloud platform services
- `monitoring` - Prometheus, Grafana, Loki
- `ingress` - ALB Ingress Controller

### Deployments

| Service           | Replicas (Prod) | CPU | Memory |
| ----------------- | --------------- | --- | ------ |
| temporal-frontend | 3               | 2   | 4Gi    |
| temporal-history  | 3               | 4   | 8Gi    |
| temporal-matching | 3               | 2   | 4Gi    |
| temporal-worker   | 3               | 2   | 4Gi    |
| cloud-api         | 3               | 1   | 2Gi    |
| cloud-console     | 2               | 0.5 | 1Gi    |

## DNS & Certificates

### Domain Structure

```
temporal-cloud.io
├── api.temporal-cloud.io        → Cloud API
├── console.temporal-cloud.io    → Web Console
├── *.tmprl.cloud                → Namespace endpoints
└── grpc.tmprl.cloud             → gRPC endpoints
```

### Certificate Management

- AWS Certificate Manager for ALB
- cert-manager for internal TLS
- Let's Encrypt for public endpoints
