# Temporal Cloud Platform

This directory contains the Temporal Cloud managed service platform, built as an additive layer on top of the open-source Temporal server.

## Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                      Cloud Console                          │
│                    (Next.js + React)                        │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                       Cloud API                             │
│              (gRPC + Connect-RPC Gateway)                   │
├─────────────┬─────────────┬─────────────┬──────────────────┤
│  Billing    │  Identity   │ Provisioner │     Audit        │
│  Service    │  Service    │   Service   │    Service       │
└─────────────┴─────────────┴─────────────┴──────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                    Temporal Clusters                        │
│              (Multi-region, Multi-tenant)                   │
└─────────────────────────────────────────────────────────────┘
```

## Quick Start

### Prerequisites

- Go 1.22+
- Node.js 20+
- Docker & Docker Compose
- PostgreSQL 15+ (or use Docker)

### Local Development

```bash
# Start infrastructure (PostgreSQL, Redis, Temporal)
docker-compose -f docker-compose.dev.yaml up -d postgresql redis temporal temporal-ui

# Run database migrations
cd schema/migrations
for f in *.up.sql; do psql -h localhost -U temporal -d temporal_cloud -f "$f"; done

# Generate proto code (requires buf)
cd api && buf generate

# Start the Cloud API service
go run ./cmd/cloud-api

# Start the Cloud Console (in another terminal)
cd console
npm install
npm run dev
```

### Access URLs

| Service       | URL                   |
| ------------- | --------------------- |
| Cloud Console | http://localhost:3000 |
| Cloud API     | http://localhost:8081 |
| Temporal UI   | http://localhost:8080 |
| Prometheus    | http://localhost:9090 |
| Grafana       | http://localhost:3001 |

## Directory Structure

```
cloud/
├── api/                    # Proto definitions and generated code
│   └── cloud/v1/          # Cloud API v1 protos
├── cmd/                    # Application entry points
│   └── cloud-api/         # Cloud API server
├── console/               # Next.js frontend
├── infra/                 # Terraform infrastructure
│   ├── modules/           # Reusable Terraform modules
│   └── environments/      # Environment configurations
├── internal/              # Internal packages
│   ├── api/v1/           # API handlers
│   ├── config/           # Configuration
│   ├── interceptors/     # gRPC interceptors
│   ├── repository/       # Data access layer
│   ├── service/          # Business logic
│   └── workflows/        # Temporal workflows
└── schema/               # Database schema
    └── migrations/       # SQL migrations
```

## Services

- **Cloud API**: Main API gateway for all cloud operations
- **Billing Service**: Usage metering, invoicing, Stripe integration
- **Identity Service**: Authentication, API keys, SAML SSO, SCIM
- **Provisioner Service**: Namespace lifecycle management
- **Audit Service**: Compliance and security logging

## API Overview

### Organization Service

- Create, update, delete organizations
- Manage members and roles
- Configure SAML SSO

### Namespace Service

- Provision and manage namespaces
- Configure retention, search attributes
- Manage certificates and certificate filters
- Failover between regions

### Billing Service

- View subscription and usage
- Manage payment methods
- Access invoices and credits

### Identity Service

- Create and manage API keys
- Service account management
- SAML login flows

### Audit Service

- Query audit events
- Export audit logs

## Infrastructure

Terraform modules for AWS deployment:

- **VPC**: Multi-AZ networking with public/private/database subnets
- **EKS**: Kubernetes cluster for running services
- **RDS**: PostgreSQL for cloud platform data
- **ElastiCache**: Redis for caching and rate limiting

## Development

See [CONTRIBUTING.md](../CONTRIBUTING.md) for development guidelines.

## Project Constitution

This project follows the principles defined in [.specify/memory/constitution.md](../.specify/memory/constitution.md):

- All cloud code is additive - no modifications to upstream Temporal files
- High code quality standards with comprehensive testing
- Security-first approach with audit logging
- Infrastructure as Code using Terraform
