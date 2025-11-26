# Security, Privacy, and Resilience

## Identity and access
- SSO everywhere (OIDC/SAML); short-lived tokens; mTLS between services; per-namespace RBAC with least privilege.
- Admin actions require MFA + Just-In-Time elevation with session recording; audit all control-plane writes.

## Secrets and key management
- Root keys in cloud KMS; Vault for dynamic secrets; no long-lived credentials. Secrets injected via sidecars or CSI; no secrets in images.
- Rotation: keys quarterly or on incident; tokens daily; database creds every 24h via Vault leases.
- Encryption: TLS1.2+ in transit; at-rest encryption with KMS; double encryption for backups; envelope encryption for payloads with customer-managed keys when required.

## Secure SDLC and zero-day handling
- Mandatory dependency scanning (SCA) and container scanning per build; weekly vulnerability backlog triage with SLA by severity (P0: 24h, P1: 72h, P2: sprint).
- Zero-day runbook: freeze risky rollouts, apply upstream patches, build and ship hotfix, add WAF rules, increase monitoring, post-incident review with hardening tasks.
- Supply chain: verify signatures (sigstore/cosign), build in isolated builders, generate SBOMs, and store attestations; disallow unsigned artifacts.

## Bot/DDoS fraud controls
- WAF rules for volumetric and application-layer attacks; rate limits per IP/tenant; CAPTCHA or WebAuthn challenges for suspicious flows.
- Anomaly detection on auth/payment events; circuit breakers for abusive traffic; shaded throttling per tenant to protect shared control-plane capacity.

## Data protection and privacy
- Data classification with tagging; PII minimization and field-level encryption where needed; retention policies aligned to contracts.
- Right-to-be-forgotten automation; access logging for sensitive tables; export controls for cross-border data movement.

## Monitoring and incident response
- Security alerts into central SIEM; runbooks with on-call rotations; P0 ≤ 15m acknowledgment; tabletop exercises quarterly.
- Forensic readiness: immutable logs, time-synced hosts, memory/disk capture playbooks.

## Compliance posture
- Map controls to SOC2/ISO27001; policy-as-code enforcement (CIS benchmarks, kube policies).
- Customer isolation guarantees documented; penetration tests twice yearly; red-team once yearly.
