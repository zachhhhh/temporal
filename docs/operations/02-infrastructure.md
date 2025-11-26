# Infrastructure, HA, and IaC

## Reference architecture
- Multi-tenant control plane, per-tenant isolation via namespaces and dedicated queues where required; worker-plane autoscaling by queue depth and execution latency.
- Recommended stack: Kubernetes (managed), Postgres/CockroachDB for persistence, Elasticsearch/OpenSearch for visibility, object storage for archival, Envoy/Ingress for edge, CDN for UI assets, Redis/Memcached for caches.
- Enforce regional shards with failover pairs; keep blast radius small by scoping shared components (metrics/logging) per region.

## High availability and disaster recovery
- SLOs: control-plane p99 latency <500ms, UI <1s, scheduling throughput targets defined per region; RPO ≤ 5 minutes, RTO ≤ 30 minutes.
- Active/active for stateless services; active/passive or multi-primary for databases depending on engine (logical replication for Postgres, multi-region for CockroachDB).
- Backups: nightly full + 15m incrementals; cross-cloud copy (AWS→GCP→Azure) with integrity checks; quarterly restore drills.
- DR playbook: automated traffic shift via DNS/GLB; infra recreation from Terraform state; application bootstrap seeds namespace metadata and system workflows.

## Multi-cloud strategy
- Common Terraform modules parameterized per cloud; no cloud-specific logic in app manifests.
- Use managed KMS, managed databases, and managed Kafka/PubSub equivalents with abstraction layers; keep storage in cloud-native services plus S3 API-compatible layer for portability.
- Federate identity via OIDC/JWT; avoid IAM entanglement by using workload identity and short-lived credentials.

## IaC standards
- Terraform with validated modules; pre-commit policy checks (OPA/Conftest); plan/apply gated by PR approvals and automated drift detection.
- Environments: `dev`, `stage`, `prod`, `perf`, `drill`; immutable per-env variables stored in Vault; per-env state in remote backends with state locking.
- GitOps: manifests synced by ArgoCD/Flux; release channels mapped to branches/tags; progressive delivery via canary/blue-green and automated rollback on SLO violation.

## Networking and edge
- WAF + rate limiting + bot detection at edge; mTLS internally; segregate control/data planes by namespace and network policies.
- CDN for UI static assets with immutable caching; signed URLs for sensitive downloads; HSTS, TLS1.2+ only.

## Observability and logging
- Standard telemetry: OTel traces + metrics + structured logs; all services export RED + USE metrics.
- Central dashboards for SLOs, capacity, and cost; alert routing with severity mapping and quiet hours for non-P1.
- Log management: retention tiers (hot 7d, warm 30d, cold 180d), PII minimization, and deletion SLAs; audit logs immutable in object storage.

## State management
- Namespaces isolated by namespace IDs; history shards balanced automatically; visibility store sized for query latency targets.
- Archival lifecycle: hot storage → infrequent access → cold archive, with index retention policies aligned to compliance.

## Capacity, scaling, and optimization
- Horizontal autoscaling on queue depth, CPU, and p99 latency; vertical autoscaling guided by perf tests.
- Cost optimization: rightsizing weekly, spot/preemptible where safe, storage lifecycle policies, compression for history payloads.
- Performance regimen: quarterly load tests 10x expected load, chaos drills (pod/process, node, AZ, region), and dependency failure injection.
