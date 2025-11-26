# Technical Implementation Plan

## 1. Repository Structure

### 1.1 Repositories

| ID  | Name                             | Type | Purpose                          |
| --- | -------------------------------- | ---- | -------------------------------- |
| R1  | temporal                         | FORK | Core server + cloud interceptors |
| R2  | cloud-api                        | FORK | Cloud API proto definitions      |
| R3  | tcld                             | FORK | Cloud CLI                        |
| R4  | terraform-provider-temporalcloud | FORK | Terraform provider               |
| R5  | temporal-cloud-platform          | NEW  | Backend services                 |
| R6  | temporal-cloud-console           | NEW  | Web UI                           |
| R7  | temporal-cloud-infra             | NEW  | Infrastructure as Code           |

### 1.2 Branch Strategy

```
main              ← Protected, requires PR + review
├── cloud/main    ← Cloud production branch
├── cloud/staging ← Staging environment
├── cloud/develop ← Development integration
└── feature/*     ← Feature branches
```

## 2. Technology Stack

### 2.1 Backend

| Component     | Technology         | Version   |
| ------------- | ------------------ | --------- |
| Language      | Go                 | 1.22+     |
| API Protocol  | gRPC + Connect     | -         |
| Database      | PostgreSQL         | 15        |
| Cache         | Redis              | 7         |
| Message Queue | Temporal Workflows | -         |
| Billing       | Stripe             | API v2024 |
| SAML          | crewjam/saml       | 0.4.x     |
| SCIM          | elimity-com/scim   | 2.x       |

### 2.2 Frontend

| Component  | Technology     | Version |
| ---------- | -------------- | ------- |
| Framework  | Next.js        | 14      |
| UI Library | React          | 18      |
| Styling    | Tailwind CSS   | 3.4     |
| Components | shadcn/ui      | latest  |
| State      | TanStack Query | 5       |
| API Client | Connect-Web    | latest  |

### 2.3 Infrastructure

| Component               | Technology           | Version |
| ----------------------- | -------------------- | ------- |
| IaC                     | Terraform            | 1.6+    |
| Container Orchestration | Kubernetes           | 1.28+   |
| Package Manager         | Helm                 | 3.13+   |
| CI/CD                   | GitHub Actions       | -       |
| Secrets                 | AWS Secrets Manager  | -       |
| Monitoring              | Prometheus + Grafana | -       |
| Logging                 | Loki                 | -       |
| Tracing                 | Jaeger               | -       |

## 3. Infrastructure Architecture

### 3.1 Multi-Region Topology

```
                    ┌─────────────────────┐
                    │   Global Load       │
                    │   Balancer (DNS)    │
                    └──────────┬──────────┘
           ┌───────────────────┼───────────────────┐
           ▼                   ▼                   ▼
    ┌─────────────┐     ┌─────────────┐     ┌─────────────┐
    │  us-east-1  │     │  eu-west-1  │     │ ap-south-1  │
    │  (Primary)  │     │ (Secondary) │     │ (Secondary) │
    └─────────────┘     └─────────────┘     └─────────────┘
```

### 3.2 Single Region Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                         VPC                                  │
│  ┌─────────────────────────────────────────────────────┐    │
│  │                 Public Subnets                       │    │
│  │  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐  │    │
│  │  │   ALB/NLB   │  │   ALB/NLB   │  │   ALB/NLB   │  │    │
│  │  │    AZ-a     │  │    AZ-b     │  │    AZ-c     │  │    │
│  │  └─────────────┘  └─────────────┘  └─────────────┘  │    │
│  └─────────────────────────────────────────────────────┘    │
│  ┌─────────────────────────────────────────────────────┐    │
│  │                 Private Subnets                      │    │
│  │  ┌─────────────────────────────────────────────┐    │    │
│  │  │              EKS Cluster                     │    │    │
│  │  │  ┌─────────┐ ┌─────────┐ ┌─────────────┐    │    │    │
│  │  │  │Temporal │ │ Cloud   │ │   Cloud     │    │    │    │
│  │  │  │ Server  │ │Platform │ │  Console    │    │    │    │
│  │  │  └─────────┘ └─────────┘ └─────────────┘    │    │    │
│  │  └─────────────────────────────────────────────┘    │    │
│  └─────────────────────────────────────────────────────┘    │
│  ┌─────────────────────────────────────────────────────┐    │
│  │                 Database Subnets                     │    │
│  │  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐  │    │
│  │  │  RDS (Pri)  │  │ RDS (Read)  │  │   Redis     │  │    │
│  │  │    AZ-a     │  │    AZ-b     │  │  Cluster    │  │    │
│  │  └─────────────┘  └─────────────┘  └─────────────┘  │    │
│  └─────────────────────────────────────────────────────┘    │
└─────────────────────────────────────────────────────────────┘
```

## 4. Implementation Phases

### Phase 1: Foundation (Weeks 1-4)

- Fork repos, setup branches
- Create new repos (platform, console, infra)
- Setup sync automation
- Create DB schemas
- Create org/billing protos

### Phase 2: Metering (Weeks 5-8)

- Metering interceptor
- Action collector
- Storage calculator
- Usage aggregation workflow

### Phase 3: Billing (Weeks 9-14)

- Stripe integration
- Invoice generation
- Quota enforcement
- Payment processing

### Phase 4: Security (Weeks 15-20)

- SAML SSO
- SCIM provisioning
- Audit logging
- API keys

### Phase 5: Console (Weeks 16-22)

- Project setup
- Auth pages
- Organization pages
- Billing pages
- Settings pages

### Phase 6: Infrastructure (Weeks 20-24)

- Staging environment
- Production environment
- DR setup
- Monitoring & alerting

### Phase 7: Testing & Launch (Weeks 22-26)

- Complete test coverage
- Load testing
- Security audit
- Beta launch
- GA launch
