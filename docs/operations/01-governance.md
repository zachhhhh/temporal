# Governance and Program Management

## SPSECKIT framing
- **Scope:** Define service boundaries (API surfaces, data stores, control planes), compliance targets, and SLAs (API: 99.95%, UI: 99.9%, control plane: 99.99% for orchestration metadata).
- **Plan:** Roadmap per quarter with exit criteria; risk register with owners; dependency mapping across infra, security, billing, and ecosystem repos.
- **Secure:** Embed threat modeling, dependency scanning, and policy-as-code gates in CI; require security signoff before production deploys.
- **Execute:** Automation-first delivery (IaC, GitHub Actions, Make targets); progressive rollouts with auto-rollback hooks.
- **Check:** SLO dashboards with error budgets; continuous verification post-release; audit trails for infra and app changes.
- **Keep improving:** Blameless postmortems with action items tracked to closure; cost and performance reviews each sprint.
- **Integrate/Transition:** Release governance with compatibility guarantees, migration guides, and downstream repo sync automation.

## Ownership and cadence
- Single accountable owner per domain: Infra, Security, SRE, Data, Billing, UI/UX, Developer Experience, Release, Compliance.
- Weekly program sync: status, risk burndown, release gates, and dependency review across the ~190 Temporal org repos.
- Monthly architecture review: changes to storage, control-plane APIs, or SDK contracts require signoff and migration notes.
- Quarterly disaster recovery simulation and cost-optimization review.

## Documentation rules
- All runbooks and policies live in `docs/operations`; changes require PR with owner approval.
- Link every production change to a ticket with acceptance criteria and roll-forward/rollback steps.
- Track deviations from standards with time-bound exceptions and explicit risk acceptance.

## Tooling governance
- Standard toolchain: Terraform for IaC, GitHub Actions for CI/CD, ArgoCD or Flux for GitOps CD, Vault for secrets, KMS for key roots, Datadog/Prometheus/OTel for observability, S3/GCS/Azure Blob for durable backups.
- No new third-party services without security and cost review; capture data residency and PII impact assessments.

## Decision flow
- Proposals captured as design docs with SLO impact, cost estimate, failure-mode analysis, and migration plan.
- Approvals: Domain owner + Security for data-impacting changes; Release + SRE for rollout strategies; Billing for monetization-impacting features.

## Metrics and success criteria
- Error budget burn (<5% monthly), change failure rate (<10%), MTTR (<30m P1, <2h P2), release on-time rate (>90%), cost per workflow execution trending downward QoQ.
