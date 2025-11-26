# ADR-001: Use PostgreSQL for Cloud Platform Persistence

## Status

Accepted

## Context

The Temporal Cloud Platform needs a reliable, scalable database for storing:

- Organizations and users
- Subscriptions and billing data
- Usage records and invoices
- Audit logs
- Namespace metadata

We need to choose a database that:

- Is highly available
- Scales with our expected growth
- Has strong consistency guarantees
- Is operationally mature in cloud environments
- Supports our team's expertise

## Decision

We will use **PostgreSQL** as the primary database for the Cloud Platform, deployed on AWS RDS with Multi-AZ configuration.

Specifically:

- AWS RDS PostgreSQL 15
- Multi-AZ for high availability
- Read replicas for read-heavy workloads
- Point-in-time recovery enabled
- Encryption at rest (AES-256)

## Consequences

### Positive

- **Proven reliability**: PostgreSQL is battle-tested at scale
- **AWS managed**: RDS handles backups, patching, failover
- **Team expertise**: Team has deep PostgreSQL experience
- **Rich ecosystem**: Excellent tooling (pgAdmin, pg_dump, etc.)
- **Strong consistency**: ACID compliance for financial data
- **JSON support**: JSONB for flexible schema where needed

### Negative

- **Vertical scaling limits**: Eventually may need sharding
- **Vendor lock-in**: Tied to AWS RDS features
- **Cost**: More expensive than self-managed
- **Single-region**: Cross-region replication adds complexity

### Neutral

- Requires careful schema design upfront
- Migration tooling (golang-migrate) adds operational overhead

## Alternatives Considered

### Option A: CockroachDB

**Pros**: Distributed, automatic scaling, multi-region native
**Cons**: Less team experience, more complex operations, higher cost
**Decision**: Overkill for initial scale; revisit at 100K+ orgs

### Option B: MongoDB

**Pros**: Flexible schema, good for rapid iteration
**Cons**: Weaker consistency guarantees, team prefers SQL
**Decision**: Rejected due to consistency requirements for billing

### Option C: Self-managed PostgreSQL on EC2

**Pros**: More control, potentially cheaper
**Cons**: Operational burden, HA complexity, no managed backups
**Decision**: RDS overhead justified by reduced ops burden

## References

- [AWS RDS PostgreSQL Documentation](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_PostgreSQL.html)
- [PostgreSQL 15 Release Notes](https://www.postgresql.org/docs/15/release-15.html)
- Team discussion: 2025-01-05
