# Terraform Modules (Complete IaC)

## Module Inventory

```
terraform/
├── modules/
│   ├── aws/
│   │   ├── vpc/              # Network infrastructure
│   │   ├── eks/              # Kubernetes cluster
│   │   ├── rds/              # PostgreSQL database
│   │   ├── elasticache/      # Redis cache
│   │   ├── alb/              # Application Load Balancer
│   │   ├── waf/              # Web Application Firewall
│   │   ├── route53/          # DNS management
│   │   ├── acm/              # SSL certificates
│   │   ├── s3/               # Object storage
│   │   ├── cloudfront/       # CDN
│   │   ├── secrets-manager/  # Secrets storage
│   │   ├── kms/              # Encryption keys
│   │   ├── iam/              # IAM roles and policies
│   │   ├── cloudwatch/       # Monitoring and alarms
│   │   └── global-accelerator/ # Global routing
│   ├── gcp/
│   │   ├── vpc/              # GCP VPC
│   │   ├── gke/              # GKE cluster
│   │   ├── cloudsql/         # Cloud SQL PostgreSQL
│   │   ├── memorystore/      # Redis
│   │   └── cloud-dns/        # DNS
│   ├── kubernetes/
│   │   ├── temporal/         # Temporal server deployment
│   │   ├── cloud-platform/   # Cloud platform services
│   │   ├── monitoring/       # Prometheus, Grafana
│   │   ├── ingress/          # Ingress controller
│   │   └── cert-manager/     # Certificate management
│   └── common/
│       ├── tags/             # Standard tagging
│       └── naming/           # Naming conventions
├── environments/
│   ├── dev/
│   ├── staging/
│   └── production/
│       ├── us-east-1/
│       ├── eu-west-1/
│       └── ap-south-1/
└── global/
    ├── iam/
    ├── route53/
    └── budgets/
```

## Core Modules

### 1. VPC Module

```hcl
# modules/aws/vpc/main.tf

variable "name" {
  type = string
}

variable "cidr" {
  type    = string
  default = "10.0.0.0/16"
}

variable "azs" {
  type    = list(string)
  default = ["us-east-1a", "us-east-1b", "us-east-1c"]
}

variable "enable_nat_gateway" {
  type    = bool
  default = true
}

resource "aws_vpc" "main" {
  cidr_block           = var.cidr
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name = var.name
  }
}

resource "aws_subnet" "public" {
  count             = length(var.azs)
  vpc_id            = aws_vpc.main.id
  cidr_block        = cidrsubnet(var.cidr, 4, count.index)
  availability_zone = var.azs[count.index]

  map_public_ip_on_launch = true

  tags = {
    Name                     = "${var.name}-public-${var.azs[count.index]}"
    "kubernetes.io/role/elb" = "1"
  }
}

resource "aws_subnet" "private" {
  count             = length(var.azs)
  vpc_id            = aws_vpc.main.id
  cidr_block        = cidrsubnet(var.cidr, 4, count.index + length(var.azs))
  availability_zone = var.azs[count.index]

  tags = {
    Name                              = "${var.name}-private-${var.azs[count.index]}"
    "kubernetes.io/role/internal-elb" = "1"
  }
}

resource "aws_subnet" "database" {
  count             = length(var.azs)
  vpc_id            = aws_vpc.main.id
  cidr_block        = cidrsubnet(var.cidr, 4, count.index + 2 * length(var.azs))
  availability_zone = var.azs[count.index]

  tags = {
    Name = "${var.name}-database-${var.azs[count.index]}"
  }
}

resource "aws_internet_gateway" "main" {
  vpc_id = aws_vpc.main.id
}

resource "aws_nat_gateway" "main" {
  count         = var.enable_nat_gateway ? length(var.azs) : 0
  allocation_id = aws_eip.nat[count.index].id
  subnet_id     = aws_subnet.public[count.index].id
}

resource "aws_eip" "nat" {
  count  = var.enable_nat_gateway ? length(var.azs) : 0
  domain = "vpc"
}

# Route tables, flow logs, etc.
output "vpc_id" {
  value = aws_vpc.main.id
}

output "private_subnet_ids" {
  value = aws_subnet.private[*].id
}

output "public_subnet_ids" {
  value = aws_subnet.public[*].id
}

output "database_subnet_ids" {
  value = aws_subnet.database[*].id
}
```

### 2. EKS Module

```hcl
# modules/aws/eks/main.tf

variable "cluster_name" {
  type = string
}

variable "cluster_version" {
  type    = string
  default = "1.29"
}

variable "vpc_id" {
  type = string
}

variable "subnet_ids" {
  type = list(string)
}

variable "node_groups" {
  type = list(object({
    name           = string
    instance_types = list(string)
    min_size       = number
    max_size       = number
    desired_size   = number
    capacity_type  = string # ON_DEMAND or SPOT
  }))
}

resource "aws_eks_cluster" "main" {
  name     = var.cluster_name
  version  = var.cluster_version
  role_arn = aws_iam_role.cluster.arn

  vpc_config {
    subnet_ids              = var.subnet_ids
    endpoint_private_access = true
    endpoint_public_access  = true
  }

  enabled_cluster_log_types = ["api", "audit", "authenticator"]

  encryption_config {
    provider {
      key_arn = aws_kms_key.eks.arn
    }
    resources = ["secrets"]
  }
}

resource "aws_eks_node_group" "main" {
  for_each = { for ng in var.node_groups : ng.name => ng }

  cluster_name    = aws_eks_cluster.main.name
  node_group_name = each.value.name
  node_role_arn   = aws_iam_role.node.arn
  subnet_ids      = var.subnet_ids
  capacity_type   = each.value.capacity_type
  instance_types  = each.value.instance_types

  scaling_config {
    min_size     = each.value.min_size
    max_size     = each.value.max_size
    desired_size = each.value.desired_size
  }

  update_config {
    max_unavailable = 1
  }
}

# Add-ons
resource "aws_eks_addon" "vpc_cni" {
  cluster_name = aws_eks_cluster.main.name
  addon_name   = "vpc-cni"
}

resource "aws_eks_addon" "coredns" {
  cluster_name = aws_eks_cluster.main.name
  addon_name   = "coredns"
}

resource "aws_eks_addon" "kube_proxy" {
  cluster_name = aws_eks_cluster.main.name
  addon_name   = "kube-proxy"
}

output "cluster_endpoint" {
  value = aws_eks_cluster.main.endpoint
}

output "cluster_ca_certificate" {
  value = aws_eks_cluster.main.certificate_authority[0].data
}
```

### 3. RDS Module

```hcl
# modules/aws/rds/main.tf

variable "identifier" {
  type = string
}

variable "engine_version" {
  type    = string
  default = "15.4"
}

variable "instance_class" {
  type = string
}

variable "allocated_storage" {
  type = number
}

variable "multi_az" {
  type    = bool
  default = true
}

variable "subnet_ids" {
  type = list(string)
}

variable "vpc_security_group_ids" {
  type = list(string)
}

variable "backup_retention_period" {
  type    = number
  default = 30
}

resource "aws_db_subnet_group" "main" {
  name       = var.identifier
  subnet_ids = var.subnet_ids
}

resource "aws_db_instance" "main" {
  identifier     = var.identifier
  engine         = "postgres"
  engine_version = var.engine_version
  instance_class = var.instance_class

  allocated_storage     = var.allocated_storage
  max_allocated_storage = var.allocated_storage * 2
  storage_type          = "gp3"
  storage_encrypted     = true
  kms_key_id            = aws_kms_key.rds.arn

  db_subnet_group_name   = aws_db_subnet_group.main.name
  vpc_security_group_ids = var.vpc_security_group_ids

  multi_az               = var.multi_az
  publicly_accessible    = false

  backup_retention_period = var.backup_retention_period
  backup_window          = "03:00-04:00"
  maintenance_window     = "Mon:04:00-Mon:05:00"

  deletion_protection = true
  skip_final_snapshot = false
  final_snapshot_identifier = "${var.identifier}-final"

  performance_insights_enabled = true
  monitoring_interval          = 60

  enabled_cloudwatch_logs_exports = ["postgresql", "upgrade"]

  username = "temporal_admin"
  password = random_password.db.result
}

resource "random_password" "db" {
  length  = 32
  special = false
}

resource "aws_secretsmanager_secret" "db" {
  name = "${var.identifier}-credentials"
}

resource "aws_secretsmanager_secret_version" "db" {
  secret_id = aws_secretsmanager_secret.db.id
  secret_string = jsonencode({
    username = aws_db_instance.main.username
    password = random_password.db.result
    host     = aws_db_instance.main.address
    port     = aws_db_instance.main.port
    database = aws_db_instance.main.db_name
  })
}

output "endpoint" {
  value = aws_db_instance.main.endpoint
}

output "secret_arn" {
  value = aws_secretsmanager_secret.db.arn
}
```

### 4. Route53 Module

```hcl
# modules/aws/route53/main.tf

variable "domain_name" {
  type = string
}

variable "records" {
  type = list(object({
    name    = string
    type    = string
    ttl     = number
    records = list(string)
  }))
  default = []
}

variable "alias_records" {
  type = list(object({
    name                   = string
    type                   = string
    alias_name             = string
    alias_zone_id          = string
    evaluate_target_health = bool
  }))
  default = []
}

resource "aws_route53_zone" "main" {
  name = var.domain_name
}

resource "aws_route53_record" "standard" {
  for_each = { for r in var.records : "${r.name}-${r.type}" => r }

  zone_id = aws_route53_zone.main.zone_id
  name    = each.value.name
  type    = each.value.type
  ttl     = each.value.ttl
  records = each.value.records
}

resource "aws_route53_record" "alias" {
  for_each = { for r in var.alias_records : "${r.name}-${r.type}" => r }

  zone_id = aws_route53_zone.main.zone_id
  name    = each.value.name
  type    = each.value.type

  alias {
    name                   = each.value.alias_name
    zone_id                = each.value.alias_zone_id
    evaluate_target_health = each.value.evaluate_target_health
  }
}

output "zone_id" {
  value = aws_route53_zone.main.zone_id
}

output "name_servers" {
  value = aws_route53_zone.main.name_servers
}
```

### 5. CloudFront Module

```hcl
# modules/aws/cloudfront/main.tf

variable "domain_names" {
  type = list(string)
}

variable "origin_domain" {
  type = string
}

variable "certificate_arn" {
  type = string
}

variable "s3_bucket_id" {
  type    = string
  default = null
}

resource "aws_cloudfront_distribution" "main" {
  enabled             = true
  is_ipv6_enabled     = true
  default_root_object = "index.html"
  aliases             = var.domain_names
  price_class         = "PriceClass_All"

  origin {
    domain_name = var.origin_domain
    origin_id   = "primary"

    custom_origin_config {
      http_port              = 80
      https_port             = 443
      origin_protocol_policy = "https-only"
      origin_ssl_protocols   = ["TLSv1.2"]
    }
  }

  default_cache_behavior {
    allowed_methods        = ["GET", "HEAD", "OPTIONS"]
    cached_methods         = ["GET", "HEAD"]
    target_origin_id       = "primary"
    viewer_protocol_policy = "redirect-to-https"
    compress               = true

    cache_policy_id          = aws_cloudfront_cache_policy.main.id
    origin_request_policy_id = aws_cloudfront_origin_request_policy.main.id
  }

  # SPA handling
  custom_error_response {
    error_code         = 404
    response_code      = 200
    response_page_path = "/index.html"
  }

  viewer_certificate {
    acm_certificate_arn      = var.certificate_arn
    ssl_support_method       = "sni-only"
    minimum_protocol_version = "TLSv1.2_2021"
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }
}

resource "aws_cloudfront_cache_policy" "main" {
  name        = "temporal-cloud-cache"
  min_ttl     = 1
  default_ttl = 86400
  max_ttl     = 31536000

  parameters_in_cache_key_and_forwarded_to_origin {
    cookies_config {
      cookie_behavior = "none"
    }
    headers_config {
      header_behavior = "none"
    }
    query_strings_config {
      query_string_behavior = "none"
    }
  }
}

output "distribution_id" {
  value = aws_cloudfront_distribution.main.id
}

output "domain_name" {
  value = aws_cloudfront_distribution.main.domain_name
}
```

## Environment Configuration

### Production US-East-1

```hcl
# environments/production/us-east-1/main.tf

module "vpc" {
  source = "../../../modules/aws/vpc"
  name   = "temporal-cloud-prod-us"
  cidr   = "10.10.0.0/16"
  azs    = ["us-east-1a", "us-east-1b", "us-east-1c"]
}

module "eks" {
  source       = "../../../modules/aws/eks"
  cluster_name = "temporal-cloud-prod-us"
  vpc_id       = module.vpc.vpc_id
  subnet_ids   = module.vpc.private_subnet_ids

  node_groups = [
    {
      name           = "general"
      instance_types = ["m6i.xlarge"]
      min_size       = 6
      max_size       = 20
      desired_size   = 6
      capacity_type  = "ON_DEMAND"
    },
    {
      name           = "temporal"
      instance_types = ["m6i.2xlarge"]
      min_size       = 3
      max_size       = 10
      desired_size   = 3
      capacity_type  = "ON_DEMAND"
    },
    {
      name           = "spot"
      instance_types = ["m6i.xlarge", "m5.xlarge"]
      min_size       = 0
      max_size       = 20
      desired_size   = 3
      capacity_type  = "SPOT"
    }
  ]
}

module "rds" {
  source            = "../../../modules/aws/rds"
  identifier        = "temporal-cloud-prod-us"
  instance_class    = "db.r6g.2xlarge"
  allocated_storage = 2000
  multi_az          = true
  subnet_ids        = module.vpc.database_subnet_ids
  vpc_security_group_ids = [module.security_groups.rds_sg_id]
}

module "redis" {
  source         = "../../../modules/aws/elasticache"
  cluster_id     = "temporal-cloud-prod-us"
  node_type      = "cache.r6g.xlarge"
  num_cache_nodes = 3
  subnet_ids     = module.vpc.private_subnet_ids
}

module "dns" {
  source      = "../../../modules/aws/route53"
  domain_name = "temporal-cloud.io"

  alias_records = [
    {
      name                   = "api"
      type                   = "A"
      alias_name             = module.alb.dns_name
      alias_zone_id          = module.alb.zone_id
      evaluate_target_health = true
    },
    {
      name                   = "console"
      type                   = "A"
      alias_name             = module.cloudfront.domain_name
      alias_zone_id          = "Z2FDTNDATAQYW2" # CloudFront zone
      evaluate_target_health = false
    }
  ]
}
```

## Module Checklist

| Module                                   | AWS | GCP | Status   |
| ---------------------------------------- | --- | --- | -------- |
| VPC/Network                              | ✅  | ✅  | Complete |
| Kubernetes (EKS/GKE)                     | ✅  | ✅  | Complete |
| Database (RDS/CloudSQL)                  | ✅  | ✅  | Complete |
| Cache (ElastiCache/Memorystore)          | ✅  | ✅  | Complete |
| Load Balancer                            | ✅  | ✅  | Complete |
| DNS (Route53/Cloud DNS)                  | ✅  | ✅  | Complete |
| CDN (CloudFront/Cloud CDN)               | ✅  | ✅  | Complete |
| WAF                                      | ✅  | ✅  | Complete |
| Secrets (Secrets Manager/Secret Manager) | ✅  | ✅  | Complete |
| KMS                                      | ✅  | ✅  | Complete |
| IAM                                      | ✅  | ✅  | Complete |
| Monitoring                               | ✅  | ✅  | Complete |
| S3/GCS                                   | ✅  | ✅  | Complete |
