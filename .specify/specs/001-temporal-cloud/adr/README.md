# Architecture Decision Records

## What is an ADR?

An Architecture Decision Record (ADR) captures an important architectural decision made along with its context and consequences.

## When to Write an ADR

- New technology choice
- Major design pattern
- Significant refactoring
- Integration approach
- Security decisions
- Breaking changes

## ADR Template

```markdown
# ADR-NNN: [Title]

## Status

[Proposed | Accepted | Deprecated | Superseded by ADR-XXX]

## Context

What is the issue that we're seeing that is motivating this decision?

## Decision

What is the change that we're proposing and/or doing?

## Consequences

What becomes easier or harder because of this change?

### Positive

-

### Negative

-

### Neutral

-

## Alternatives Considered

### Option A

Description and trade-offs

### Option B

Description and trade-offs

## References

- [Link to relevant docs]
- [Link to discussion]
```

## Process

1. **Propose**: Create ADR in `proposed/` folder
2. **Review**: Team reviews in PR
3. **Accept**: Move to `accepted/` folder
4. **Implement**: Reference ADR in code
5. **Supersede**: When replaced, update status

## Naming Convention

`ADR-{number}-{short-title}.md`

Example: `ADR-001-use-postgresql-for-persistence.md`

## Index

| ADR | Title                              | Status   | Date       |
| --- | ---------------------------------- | -------- | ---------- |
| 001 | Use PostgreSQL for persistence     | Accepted | 2025-01-01 |
| 002 | Adopt gRPC for Cloud Ops API       | Accepted | 2025-01-15 |
| 003 | Multi-region namespace replication | Accepted | 2025-02-01 |
