# Runbooks and Operational Playbooks

## Incident response
- P0 triggers: control-plane outage, data loss risk, auth failure impacting >5% tenants, billing double-charge risk. Page SRE + Domain owner; assemble bridge; update status page within 15 minutes.
- Workflow: stabilize (traffic routing, rate limiting), mitigate (rollback/canary halt, feature flag disable), communicate (customers + internal), recover, document, and drive postmortem with action items.

## Disaster recovery
- Trigger: region down >15 minutes or data durability threat. Steps: freeze writes, promote replica in paired region, restore backups if needed, reconfigure DNS/GLB, validate namespace metadata and histories, lift traffic gradually, run data consistency checks, and backfill visibility indices.
- Post-restore: confirm RPO/RTO met; reconcile billing events; re-enable background jobs and archival.

## Upgrade and rollback
- Pre-checks: schema migrations dry-run in stage; backup snapshots; capacity headroom; compatibility tests for SDKs and plugins.
- Rollout: canary per shard/namespace; monitor p99 latency, error rate, queue depth; pause on budget burn; automatic rollback ready with previous release artifacts.
- Post-checks: smoke tests, data drift checks, perf regression guardrails; mark release ready when error budget impact ≤1%.

## Scaling and performance
- Scale-out triggers: queue depth > target for 5m, CPU >70%, p99 latency breach. Actions: increase worker deployments, adjust shard counts, and tune visibility store resources.
- Performance tuning: optimize persistence configs (batch sizes, cache), enable compression, adjust history retention. Run load tests after changes.

## Cost optimization
- Weekly rightsizing report; shut down idle perf/test clusters nightly; storage lifecycle rules; choose spot/preemptible nodes for stateless workers with PodDisruptionBudgets.
- Validate cost-impacting changes via canary; maintain dashboards for cost per workflow execution.

## Logging and observability hygiene
- Ensure logs are structured and sampled; drop noisy fields; enforce retention tiers; redact sensitive data at source.
- Alerts tied to SLOs; false-positive review weekly; run synthetic checks for UI and API.

## CDN and edge
- Cache UI assets aggressively with immutable hashes; purge on release; use signed URLs for private content; monitor cache hit rate and TLS error spikes.

## Plugin and dependency management
- Maintain compatibility matrix; run nightly CI against supported plugin versions; auto-create issues for failures.
- For new plugins: security review, load test, feature flags, rollout via staged tenants, and observability baked in.
