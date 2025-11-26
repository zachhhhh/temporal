# Development Environment Configuration

terraform {
  required_version = ">= 1.6"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 3.0"
    }
  }

  backend "s3" {
    bucket         = "temporal-cloud-terraform-state"
    key            = "dev/terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
    dynamodb_table = "temporal-cloud-terraform-locks"
  }
}

provider "aws" {
  region = var.region

  default_tags {
    tags = {
      Project     = "temporal-cloud"
      Environment = "dev"
      ManagedBy   = "terraform"
    }
  }
}

variable "region" {
  description = "AWS region"
  type        = string
  default     = "us-east-1"
}

locals {
  environment        = "dev"
  availability_zones = ["${var.region}a", "${var.region}b", "${var.region}c"]
}

# VPC
module "vpc" {
  source = "../../modules/vpc"

  environment        = local.environment
  region             = var.region
  vpc_cidr           = "10.0.0.0/16"
  availability_zones = local.availability_zones
}

# EKS Cluster
module "eks" {
  source = "../../modules/eks"

  environment        = local.environment
  cluster_name       = "temporal-cloud-dev"
  cluster_version    = "1.29"
  vpc_id             = module.vpc.vpc_id
  private_subnet_ids = module.vpc.private_subnet_ids

  node_groups = {
    general = {
      instance_types = ["t3.medium"]
      min_size       = 2
      max_size       = 4
      desired_size   = 2
      disk_size      = 50
    }
  }
}

# RDS PostgreSQL
module "rds" {
  source = "../../modules/rds"

  environment                = local.environment
  identifier                 = "temporal-cloud-dev"
  vpc_id                     = module.vpc.vpc_id
  database_subnet_ids        = module.vpc.database_subnet_ids
  allowed_security_group_ids = [module.eks.cluster_security_group_id]
  instance_class             = "db.t3.medium"
  allocated_storage          = 100
  multi_az                   = false
}

# Outputs
output "vpc_id" {
  value = module.vpc.vpc_id
}

output "eks_cluster_endpoint" {
  value = module.eks.cluster_endpoint
}

output "rds_endpoint" {
  value = module.rds.endpoint
}

output "rds_secret_arn" {
  value = module.rds.secret_arn
}
