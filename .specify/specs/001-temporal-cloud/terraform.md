# Terraform Provider

## Overview

The `temporalio/temporalcloud` provider allows managing Temporal Cloud resources via Infrastructure as Code.

## Configuration

```hcl
terraform {
  required_providers {
    temporalcloud = {
      source  = "temporalio/temporalcloud"
      version = "0.4.0"
    }
  }
}

provider "temporalcloud" {
  api_key = var.temporal_api_key
}
```

## Resources

### `temporalcloud_namespace`

```hcl
resource "temporalcloud_namespace" "main" {
  name           = "production-ns"
  regions        = ["aws-us-east-1", "aws-us-west-2"]
  retention_days = 30

  certificate {
    certificate = file("certs/ca.pem")
    filters = {
      common_name = "worker.prod.example.com"
    }
  }

  # High Availability
  is_global_namespace = true
}
```

### `temporalcloud_user`

```hcl
resource "temporalcloud_user" "developer" {
  email        = "dev@example.com"
  account_role = "developer"

  namespace_access {
    namespace_id = temporalcloud_namespace.main.id
    permission   = "write"
  }
}
```

### `temporalcloud_service_account`

```hcl
resource "temporalcloud_service_account" "ci" {
  name         = "ci-pipeline"
  account_role = "admin"
}
```

### `temporalcloud_api_key`

```hcl
resource "temporalcloud_api_key" "ci_key" {
  name       = "ci-key"
  owner_id   = temporalcloud_service_account.ci.id
  owner_type = "service_account"
  duration   = "90d"
}
```

## Data Sources

- `temporalcloud_regions`: List supported regions
- `temporalcloud_namespace`: Get namespace details
- `temporalcloud_service_account`: Get SA details

## Best Practices

1. **State Management**: Use S3/GCS backend with locking
2. **Secret Management**: Do not output API keys to console
3. **Module Structure**: Create a module for standard namespace setup (ns + certs + rbac)
