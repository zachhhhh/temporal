# Temporal Cloud Project Constitution

## Core Principles

### 1. Upstream Compatibility

- NEVER modify existing upstream Temporal files
- All cloud code MUST be additive (new files/packages only)
- Daily upstream sync MUST pass without manual intervention
- Use Go build tags to conditionally include cloud features

### 2. Code Quality

- All code MUST pass linting (`golangci-lint`, `eslint`)
- Test coverage MUST be ≥80% for new code
- All public APIs MUST have documentation
- No `TODO` or `FIXME` in merged code

### 3. Security First

- No secrets in code or config files
- All secrets via environment variables or secret managers
- All external APIs require authentication
- Audit logging for all state-changing operations

### 4. Infrastructure as Code

- ALL infrastructure defined in Terraform
- NO manual cloud console changes
- All changes via pull request with review
- Environments reproducible from code

### 5. Testing Standards

- Unit tests for all business logic
- Integration tests for all API endpoints
- E2E tests for critical user journeys
- Load tests before production release

### 6. Observability

- All services emit metrics (OpenTelemetry)
- Structured logging (JSON format)
- Distributed tracing enabled
- Alerting for all SLO violations

### 7. High Availability

- No single points of failure
- Multi-AZ deployment minimum
- Automated failover for all stateful services
- RTO < 1 hour, RPO < 5 minutes

### 8. Documentation

- README in every repository
- Architecture Decision Records (ADRs) for major decisions
- Runbooks for all operational procedures
- API documentation auto-generated from code
