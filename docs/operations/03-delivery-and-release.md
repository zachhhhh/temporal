# Delivery, CI/CD, and Releases

## Git and issue workflow
- Trunk-based with short-lived feature branches; protected `main` with required reviews, status checks, and linear history.
- Issue lifecycle: triage within 24h, severity labels, reproduction steps, acceptance criteria, and owner assignment.
- Bugfix policy: P0 fix immediately with hotfix branch; P1 within 24h; P2/P3 scheduled into next sprint. Require test coverage for regressions.

## CI/CD standards
- Required checks: `make lint-code`, `make unit-test -tags test_dep`, static analysis, dependency/license scan, SBOM generation, OPA policy checks for IaC manifests, integration tests gated per service.
- Build artifacts signed; provenance captured (SLSA-style); container images scanned before release and at rest.
- Progressive delivery: canary 5% → 25% → 50% → 100% with automatic rollback on SLO burn or error spike; feature flags for risky paths.
- CD via GitOps: merges to `main` auto-sync to `dev`, promotion to `stage` via release PR, `prod` via approval with change ticket and rollback plan.

## Release management
- Cadence: biweekly minor, quarterly major; semantic versioning with compatibility guarantees and migration notes.
- Release checklist: changelog, migrations, config diffs, feature flag defaults, DR impact, backfill/archival impact, billing/entitlement changes, docs update.
- Upgrade policy: never skip more than one major; provide data migration steps and read-only windows if required; shadow traffic before cutover.
- Rollback policy: keep last two prod releases deployable; automatic rollback triggers on sustained SLO burn or error budget depletion.

## Testing and QA
- Pyramid: unit (fast), service-level integration, end-to-end, load/perf, chaos. All tests tagged; integration/infrastructure tests isolated behind `-tags integration`.
- Golden traces and snapshots for deterministic history workflows; fuzzing for workflow inputs; contract tests for SDK compatibility.
- Release candidates validated in `stage` with production-like data subsets; blue/green smoke tests automated via synthetic clients.

## Repo sync and upstream alignment
- Track all repos in the Temporal org via scheduled GitHub API crawler (hourly) that populates a manifest; auto-create forks/mirrors and watch for new repos.
- Upstream sync: mirror `temporalio/temporal` into this repo daily; auto-open PRs for divergences; block release if drift exceeds 72h.
- Cross-repo change orchestration: shared modules versioned; change proposals include impact matrix; orchestrated merges tested via multi-repo integration pipeline.
- Plugin and component management: maintain compatibility matrix by version; CI tests plugin set against current and next release candidates.

## Environment promotion
- Dev: feature validation; Stage: release candidate validation; Perf: load/chaos; Prod: customer traffic; DR: cold/warm standby validation.
- Promotions require evidence: test report links, SLO snapshots, rollback plan, and owner signoff.
