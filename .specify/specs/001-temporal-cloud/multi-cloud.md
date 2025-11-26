# Multi-Cloud Strategy

## Overview

Temporal Cloud supports deployment across AWS and GCP to provide:

- Geographic coverage
- Vendor redundancy
- Customer preference accommodation
- Regulatory compliance (data residency)

## Cloud Coverage

### Primary Cloud: AWS

| Region      | Code           | Services       |
| ----------- | -------------- | -------------- |
| N. Virginia | us-east-1      | Full (Primary) |
| Oregon      | us-west-2      | Full           |
| Ireland     | eu-west-1      | Full           |
| Frankfurt   | eu-central-1   | Full           |
| Singapore   | ap-southeast-1 | Full           |
| Sydney      | ap-southeast-2 | Full           |
| Tokyo       | ap-northeast-1 | Full           |
| Mumbai      | ap-south-1     | Full           |

### Secondary Cloud: GCP

| Region      | Code         | Services |
| ----------- | ------------ | -------- |
| Iowa        | us-central1  | Full     |
| Oregon      | us-west1     | Full     |
| N. Virginia | us-east4     | Full     |
| Frankfurt   | europe-west3 | Full     |
| Mumbai      | asia-south1  | Full     |

## Architecture

### Control Plane

Single control plane on AWS, manages all clouds:

```
┌─────────────────────────────────────────────────────────────────┐
│                    Control Plane (AWS us-east-1)                 │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐              │
│  │  Cloud API  │  │   Billing   │  │  Provision  │              │
│  └─────────────┘  └─────────────┘  └─────────────┘              │
└───────────────────────────┬─────────────────────────────────────┘
                            │
            ┌───────────────┼───────────────┐
            ▼               ▼               ▼
      ┌──────────┐    ┌──────────┐    ┌──────────┐
      │   AWS    │    │   AWS    │    │   GCP    │
      │us-east-1 │    │eu-west-1 │    │us-central│
      └──────────┘    └──────────┘    └──────────┘
```

### Data Plane

Each cloud region runs independently:

```yaml
# Per-region components
- Temporal Server (Frontend, History, Matching, Worker)
- PostgreSQL (Regional, no cross-cloud replication by default)
- Redis (Regional)
- Load Balancer (Cloud-native: ALB/Cloud Load Balancing)
```

## Cross-Cloud Replication

### Namespace Replication Pairs

| Primary (AWS) | Secondary (GCP) | Latency |
| ------------- | --------------- | ------- |
| us-east-1     | us-central1     | ~30ms   |
| us-west-2     | us-west1        | ~20ms   |
| eu-central-1  | europe-west3    | ~10ms   |
| ap-south-1    | asia-south1     | ~5ms    |

### Replication Architecture

```
┌─────────────────────┐         ┌─────────────────────┐
│    AWS us-east-1    │         │   GCP us-central1   │
│  ┌───────────────┐  │         │  ┌───────────────┐  │
│  │   Temporal    │──┼─────────┼─▶│   Temporal    │  │
│  │   Primary     │  │  gRPC   │  │   Standby     │  │
│  └───────────────┘  │         │  └───────────────┘  │
│  ┌───────────────┐  │         │  ┌───────────────┐  │
│  │  PostgreSQL   │──┼─────────┼─▶│  PostgreSQL   │  │
│  │   Primary     │  │  Async  │  │   Replica     │  │
│  └───────────────┘  │         │  └───────────────┘  │
└─────────────────────┘         └─────────────────────┘
```

### Cross-Cloud Connectivity

**Option 1: Public Internet (Encrypted)**

- TLS 1.3 for all traffic
- IP allowlisting
- Simpler, higher latency

**Option 2: Dedicated Interconnect**

- AWS Direct Connect + GCP Partner Interconnect
- Private connectivity
- Lower latency, higher cost
- Used for high-volume replication

## Terraform Multi-Cloud

### Provider Configuration

```hcl
# providers.tf
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
    google = {
      source  = "hashicorp/google"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
  alias  = "primary"
}

provider "google" {
  project = var.gcp_project
  region  = var.gcp_region
  alias   = "secondary"
}
```

### Module Structure

```
terraform/
├── modules/
│   ├── aws/
│   │   ├── vpc/
│   │   ├── eks/
│   │   ├── rds/
│   │   └── temporal/
│   ├── gcp/
│   │   ├── vpc/
│   │   ├── gke/
│   │   ├── cloudsql/
│   │   └── temporal/
│   └── common/
│       ├── monitoring/
│       └── dns/
├── environments/
│   ├── prod-aws-us-east-1/
│   ├── prod-aws-eu-west-1/
│   ├── prod-gcp-us-central1/
│   └── staging/
```

### Unified Resource Abstraction

```hcl
# modules/common/kubernetes-cluster/main.tf
module "aws_cluster" {
  count  = var.cloud == "aws" ? 1 : 0
  source = "../../aws/eks"
  # ...
}

module "gcp_cluster" {
  count  = var.cloud == "gcp" ? 1 : 0
  source = "../../gcp/gke"
  # ...
}

output "cluster_endpoint" {
  value = var.cloud == "aws" ? module.aws_cluster[0].endpoint : module.gcp_cluster[0].endpoint
}
```

## Cloud-Agnostic Components

### Kubernetes (EKS/GKE)

Same Helm charts, different values:

```yaml
# values-aws.yaml
ingress:
  class: alb
  annotations:
    alb.ingress.kubernetes.io/scheme: internet-facing

# values-gcp.yaml
ingress:
  class: gce
  annotations:
    kubernetes.io/ingress.class: gce
```

### Database (RDS/Cloud SQL)

Same schema, different provisioning:

```hcl
# AWS
resource "aws_db_instance" "temporal" {
  engine         = "postgres"
  engine_version = "15"
  instance_class = "db.r6g.xlarge"
}

# GCP
resource "google_sql_database_instance" "temporal" {
  database_version = "POSTGRES_15"
  settings {
    tier = "db-custom-4-16384"
  }
}
```

### Object Storage (S3/GCS)

```go
// Cloud-agnostic storage interface
type ObjectStore interface {
    Put(ctx context.Context, key string, data []byte) error
    Get(ctx context.Context, key string) ([]byte, error)
    Delete(ctx context.Context, key string) error
}

// Factory
func NewObjectStore(cloud, bucket string) ObjectStore {
    switch cloud {
    case "aws":
        return &S3Store{bucket: bucket}
    case "gcp":
        return &GCSStore{bucket: bucket}
    }
}
```

## DNS & Traffic Management

### Global Load Balancing

```
                    ┌─────────────────────┐
                    │   Route53 / Cloud   │
                    │       DNS           │
                    └──────────┬──────────┘
                               │ Latency-based routing
           ┌───────────────────┼───────────────────┐
           ▼                   ▼                   ▼
    ┌─────────────┐     ┌─────────────┐     ┌─────────────┐
    │ AWS ALB     │     │ AWS ALB     │     │ GCP LB      │
    │ us-east-1   │     │ eu-west-1   │     │ us-central1 │
    └─────────────┘     └─────────────┘     └─────────────┘
```

### Failover Configuration

```hcl
resource "aws_route53_health_check" "aws_region" {
  fqdn              = "us-east-1.temporal-cloud.io"
  port              = 443
  type              = "HTTPS"
  resource_path     = "/health"
  failure_threshold = 3
  request_interval  = 10
}

resource "aws_route53_record" "api" {
  zone_id = aws_route53_zone.main.zone_id
  name    = "api.temporal-cloud.io"
  type    = "A"

  set_identifier = "aws-us-east-1"

  alias {
    name    = aws_lb.api.dns_name
    zone_id = aws_lb.api.zone_id
  }

  latency_routing_policy {
    region = "us-east-1"
  }

  health_check_id = aws_route53_health_check.aws_region.id
}
```

## Cost Optimization

### Cloud Pricing Comparison

| Resource         | AWS (us-east-1) | GCP (us-central1) |
| ---------------- | --------------- | ----------------- |
| Compute (4 vCPU) | $0.17/hr        | $0.15/hr          |
| PostgreSQL       | $0.46/hr        | $0.42/hr          |
| Egress (per GB)  | $0.09           | $0.12             |
| Load Balancer    | $0.025/hr       | $0.025/hr         |

### Multi-Cloud Cost Strategy

1. **Committed Use**: 1-year reserved instances on primary cloud
2. **Spot/Preemptible**: For non-critical workloads
3. **Right-sizing**: Regular review of instance sizes
4. **Egress Optimization**: Minimize cross-cloud data transfer

## Monitoring Across Clouds

### Unified Observability

```yaml
# Single Grafana instance with multiple data sources
datasources:
  - name: Prometheus-AWS
    type: prometheus
    url: https://prometheus.aws.temporal-cloud.io

  - name: Prometheus-GCP
    type: prometheus
    url: https://prometheus.gcp.temporal-cloud.io

  - name: CloudWatch
    type: cloudwatch

  - name: Stackdriver
    type: stackdriver
```

### Cross-Cloud Alerting

```yaml
# Alert on cross-cloud replication lag
alert: CrossCloudReplicationLag
expr: |
  temporal_replication_lag_seconds{
    source_cloud="aws",
    target_cloud="gcp"
  } > 60
severity: warning
```
