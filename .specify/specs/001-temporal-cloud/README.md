# Temporal Cloud Platform Specification

> Complete specification for building a managed Temporal Cloud service.

## Quick Links

| Document                       | Purpose                               |
| ------------------------------ | ------------------------------------- |
| [spec.md](spec.md)             | Vision, requirements, success metrics |
| [plan.md](plan.md)             | Technical architecture and phases     |
| [tasks.md](tasks.md)           | Implementation task breakdown         |
| [quickstart.md](quickstart.md) | Developer setup guide                 |

## Document Index

### 📋 Core Specifications

| Document                                           | Description                                  |
| -------------------------------------------------- | -------------------------------------------- |
| [spec.md](spec.md)                                 | Functional and non-functional requirements   |
| [plan.md](plan.md)                                 | Architecture, technology stack, phases       |
| [backend-architecture.md](backend-architecture.md) | Service architecture, API flows, data models |
| [tasks.md](tasks.md)                               | Detailed implementation tasks                |
| [data-model.md](data-model.md)                     | Database schema (14 tables)                  |
| [quickstart.md](quickstart.md)                     | Local development setup                      |

### 🏗️ Infrastructure

| Document                                           | Description                            |
| -------------------------------------------------- | -------------------------------------- |
| [infra.md](infra.md)                               | Cloud architecture overview            |
| [terraform-modules.md](terraform-modules.md)       | Complete IaC module inventory          |
| [dns-domains.md](dns-domains.md)                   | DNS, domains, Route53, certificates    |
| [email-infrastructure.md](email-infrastructure.md) | Email (SendGrid), templates, workflows |
| [ha.md](ha.md)                                     | High availability design               |
| [dr.md](dr.md)                                     | Disaster recovery procedures           |
| [regions.md](regions.md)                           | Supported AWS/GCP regions              |
| [bootstrap.md](bootstrap.md)                       | Day 0 cluster bootstrap                |
| [capacity-planning.md](capacity-planning.md)       | Scaling and resource planning          |
| [multi-cloud.md](multi-cloud.md)                   | AWS + GCP multi-cloud strategy         |
| [state-management.md](state-management.md)         | Terraform, app, session state          |
| [environment-config.md](environment-config.md)     | Environment configuration              |

### 🔄 Operations

| Document                                         | Description                   |
| ------------------------------------------------ | ----------------------------- |
| [cicd.md](cicd.md)                               | CI/CD pipeline configuration  |
| [monitoring.md](monitoring.md)                   | Observability and alerting    |
| [runbooks.md](runbooks.md)                       | Operational procedures        |
| [incident-management.md](incident-management.md) | Incident response process     |
| [database-operations.md](database-operations.md) | DB maintenance and migrations |

### 🔐 Security & Access

| Document                                             | Description                 |
| ---------------------------------------------------- | --------------------------- |
| [security.md](security.md)                           | Security architecture       |
| [security-hardening.md](security-hardening.md)       | Defense in depth, hardening |
| [certificates.md](certificates.md)                   | mTLS certificate management |
| [api-keys.md](api-keys.md)                           | API key lifecycle           |
| [secrets-management.md](secrets-management.md)       | Secrets storage & rotation  |
| [saml.md](saml.md)                                   | SAML SSO integration        |
| [scim.md](scim.md)                                   | SCIM user provisioning      |
| [users.md](users.md)                                 | User roles and permissions  |
| [compliance-automation.md](compliance-automation.md) | SOC 2 automation            |
| [audit-logs.md](audit-logs.md)                       | Audit logging               |
| [bot-protection.md](bot-protection.md)               | Bot & abuse prevention      |
| [zero-day-response.md](zero-day-response.md)         | Vulnerability response      |

### 💰 Billing & Payments

| Document                                       | Description                  |
| ---------------------------------------------- | ---------------------------- |
| [pricing.md](pricing.md)                       | Plan tiers and pricing       |
| [actions-detailed.md](actions-detailed.md)     | Billable actions reference   |
| [limits.md](limits.md)                         | System limits and quotas     |
| [stripe-mapping.md](stripe-mapping.md)         | Stripe product configuration |
| [payment-collection.md](payment-collection.md) | Dunning and collections      |

### 🔌 APIs & Tools

| Document                                   | Description               |
| ------------------------------------------ | ------------------------- |
| [cloud-ops-api.md](cloud-ops-api.md)       | gRPC API reference        |
| [tcld.md](tcld.md)                         | CLI command reference     |
| [terraform.md](terraform.md)               | Terraform provider        |
| [metrics-endpoint.md](metrics-endpoint.md) | Prometheus metrics export |
| [sdk-extensions.md](sdk-extensions.md)     | SDK helper libraries      |
| [api-versioning.md](api-versioning.md)     | API versioning policy     |

### 🎨 Frontend

| Document                               | Description          |
| -------------------------------------- | -------------------- |
| [console-design.md](console-design.md) | UI/UX specifications |
| [namespaces.md](namespaces.md)         | Namespace management |

### 📦 Development Process

| Document                                             | Description                      |
| ---------------------------------------------------- | -------------------------------- |
| [git.md](git.md)                                     | Git workflow and branch strategy |
| [upstream-repos.md](upstream-repos.md)               | Managing 190+ temporalio repos   |
| [repo-automation.md](repo-automation.md)             | Automated repo sync & discovery  |
| [automation-strategy.md](automation-strategy.md)     | End-to-end automation blueprint  |
| [release-management.md](release-management.md)       | Release process                  |
| [upgrade-policies.md](upgrade-policies.md)           | Update & upgrade policies        |
| [testing.md](testing.md)                             | Testing strategy                 |
| [qa-process.md](qa-process.md)                       | QA workflow                      |
| [bug-triage.md](bug-triage.md)                       | Issue management                 |
| [feature-flags.md](feature-flags.md)                 | Feature flag system              |
| [dependency-management.md](dependency-management.md) | Dependency updates               |
| [plugin-management.md](plugin-management.md)         | Plugin safety & architecture     |

### 🚀 Optimization & Reliability

| Document                                         | Description                     |
| ------------------------------------------------ | ------------------------------- |
| [cost-optimization.md](cost-optimization.md)     | Compute, storage, network costs |
| [system-optimization.md](system-optimization.md) | Kernel, runtime, DB tuning      |
| [crash-proofing.md](crash-proofing.md)           | Circuit breakers, chaos eng     |
| [log-management.md](log-management.md)           | Centralized logging & retention |
| [cdn.md](cdn.md)                                 | Edge caching & global routing   |

### 🤝 Customer & Support

| Document                                         | Description                  |
| ------------------------------------------------ | ---------------------------- |
| [customer-onboarding.md](customer-onboarding.md) | Signup to first workflow     |
| [sla.md](sla.md)                                 | Service level agreements     |
| [support-escalation.md](support-escalation.md)   | Support tiers and escalation |

### 🧠 Business Logic

| Document                                                           | Description                  |
| ------------------------------------------------------------------ | ---------------------------- |
| [logic/billing-reconciliation.md](logic/billing-reconciliation.md) | Billing workflow algorithms  |
| [logic/quota-enforcement.md](logic/quota-enforcement.md)           | Rate limiting implementation |
| [logic/namespace-failover.md](logic/namespace-failover.md)         | Failover state machine       |

### 📝 Architecture Decisions

| Document                                                                       | Description            |
| ------------------------------------------------------------------------------ | ---------------------- |
| [adr/README.md](adr/README.md)                                                 | ADR index and template |
| [adr/ADR-001-postgresql-persistence.md](adr/ADR-001-postgresql-persistence.md) | Database choice        |

### 📜 Governance

| Document                                               | Description        |
| ------------------------------------------------------ | ------------------ |
| [../memory/constitution.md](../memory/constitution.md) | Project principles |

---

## End-to-End Coverage

This specification covers the complete lifecycle:

```
┌─────────────────────────────────────────────────────────────────────┐
│                         CUSTOMER JOURNEY                             │
├─────────────────────────────────────────────────────────────────────┤
│  Signup → Verify → Trial → Subscribe → Use → Pay → Support → Renew │
│    ↓        ↓       ↓         ↓        ↓     ↓       ↓         ↓   │
│  [onboarding] [saml]  [pricing] [stripe] [limits] [dunning] [sla] │
└─────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────┐
│                         ENGINEERING JOURNEY                          │
├─────────────────────────────────────────────────────────────────────┤
│  Plan → Code → Test → Review → Deploy → Monitor → Debug → Improve   │
│   ↓      ↓      ↓       ↓        ↓        ↓        ↓        ↓      │
│ [spec] [git] [testing] [qa]   [cicd]  [monitor] [bugs]  [release] │
└─────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────┐
│                         INFRASTRUCTURE JOURNEY                       │
├─────────────────────────────────────────────────────────────────────┤
│  Bootstrap → Deploy → Scale → Monitor → Incident → Recover          │
│     ↓          ↓       ↓        ↓          ↓          ↓            │
│ [bootstrap] [infra] [capacity] [monitor] [incident]  [dr]          │
└─────────────────────────────────────────────────────────────────────┘
```

## Getting Started

1. **Read the spec**: Start with [spec.md](spec.md) for requirements
2. **Understand the plan**: Review [plan.md](plan.md) for architecture
3. **Set up locally**: Follow [quickstart.md](quickstart.md)
4. **Pick a task**: Check [tasks.md](tasks.md) for work items

## Document Count

**Total: 75 specification documents**

- Core: 6 (includes backend-architecture)
- Infrastructure: 12 (+3: terraform-modules, dns-domains, email)
- Operations: 5
- Security: 12
- Billing: 5
- APIs: 6
- Frontend: 2
- Process: 12 (includes automation-strategy)
- Optimization: 5
- Customer: 3
- Logic: 3
- ADR: 2
- Governance: 1

## Complete Coverage

This spec covers **everything** from start to receiving payment:

| Phase          | Coverage                                                |
| -------------- | ------------------------------------------------------- |
| Backend        | `backend-architecture.md` (services, APIs, data models) |
| IaC            | `terraform-modules.md` (all Terraform modules)          |
| DNS/Domains    | `dns-domains.md` (Route53, certs, dynamic DNS)          |
| Email          | `email-infrastructure.md` (SendGrid, templates)         |
| Automation     | `automation-strategy.md` (CI/CD, scaling, rotation)     |
| Repo Sync      | `repo-automation.md`, `upstream-repos.md` (190+ repos)  |
| Git/Process    | `git.md`, `plugin-management.md`                        |
| Development    | `quickstart.md`, `dependency-management.md`             |
| Optimization   | `cost-optimization.md`, `system-optimization.md`        |
| Reliability    | `crash-proofing.md`, `log-management.md`, `cdn.md`      |
| Testing/QA     | `testing.md`, `qa-process.md`                           |
| Deployment     | `cicd.md`, `release-management.md`                      |
| Infrastructure | `infra.md`, `multi-cloud.md`, `bootstrap.md`            |
| State          | `state-management.md` (Terraform, app, session)         |
| Environments   | `environment-config.md`                                 |
| HA/DR          | `ha.md`, `dr.md`, `capacity-planning.md`                |
| Security       | `security-hardening.md`, `secrets-management.md`        |
| Bot/Abuse      | `bot-protection.md`                                     |
| Zero-Day       | `zero-day-response.md`                                  |
| Upgrades       | `upgrade-policies.md`                                   |
| Billing        | `pricing.md`, `stripe-mapping.md`                       |
| Payments       | `payment-collection.md` (dunning)                       |
| Onboarding     | `customer-onboarding.md`                                |
| Support        | `support-escalation.md`, `sla.md`                       |
