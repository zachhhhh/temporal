# Database Operations

## Migration Strategy

### Tool: golang-migrate

```bash
# Install
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Create new migration
migrate create -ext sql -dir schema -seq add_user_groups

# Run migrations
migrate -path schema -database "$DATABASE_URL" up

# Rollback last migration
migrate -path schema -database "$DATABASE_URL" down 1
```

### Migration Naming

`{version}_{description}.up.sql` and `{version}_{description}.down.sql`

Example:

- `000001_create_organizations.up.sql`
- `000001_create_organizations.down.sql`

### Migration Rules

1. **Always reversible**: Every `up` must have a `down`
2. **Backwards compatible**: Old code must work with new schema
3. **No data loss**: Never drop columns without migration period
4. **Atomic**: Each migration is a single transaction

### Deployment Process

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Deploy    │────▶│   Migrate   │────▶│   Deploy    │
│   DB First  │     │   (auto)    │     │   App       │
└─────────────┘     └─────────────┘     └─────────────┘
```

1. Migration runs in CI/CD pipeline
2. App deployment waits for migration success
3. App must handle both old and new schema during rollout

### Breaking Changes

When schema change is not backwards compatible:

**Phase 1: Add**

```sql
-- Add new column (nullable)
ALTER TABLE users ADD COLUMN new_email VARCHAR(255);
```

**Phase 2: Migrate**

```sql
-- Backfill data
UPDATE users SET new_email = email WHERE new_email IS NULL;
```

**Phase 3: Switch**

- Deploy new app code using new column
- Old code still works with old column

**Phase 4: Remove**

```sql
-- Drop old column (after all apps updated)
ALTER TABLE users DROP COLUMN email;
ALTER TABLE users RENAME COLUMN new_email TO email;
```

## Backup & Recovery

### Automated Backups (RDS)

| Type         | Frequency  | Retention |
| ------------ | ---------- | --------- |
| Snapshot     | Daily      | 35 days   |
| WAL          | Continuous | 7 days    |
| Cross-region | Daily      | 7 days    |

### Manual Backup

```bash
# Create snapshot
aws rds create-db-snapshot \
  --db-instance-identifier temporal-cloud-prod \
  --db-snapshot-identifier manual-backup-$(date +%Y%m%d)
```

### Point-in-Time Recovery

```bash
# Restore to specific time
aws rds restore-db-instance-to-point-in-time \
  --source-db-instance-identifier temporal-cloud-prod \
  --target-db-instance-identifier temporal-cloud-pitr \
  --restore-time 2025-01-15T10:00:00Z
```

### Disaster Recovery

See `dr.md` for full DR procedures.

## Performance Tuning

### Connection Pooling

Use PgBouncer for connection pooling:

```yaml
# pgbouncer.ini
[databases]
temporal_cloud = host=rds-endpoint port=5432 dbname=temporal_cloud

[pgbouncer]
pool_mode = transaction
max_client_conn = 1000
default_pool_size = 50
```

### Index Management

```sql
-- Find unused indexes
SELECT
  schemaname || '.' || relname AS table,
  indexrelname AS index,
  pg_size_pretty(pg_relation_size(i.indexrelid)) AS index_size,
  idx_scan as index_scans
FROM pg_stat_user_indexes ui
JOIN pg_index i ON ui.indexrelid = i.indexrelid
WHERE NOT indisunique AND idx_scan < 50
ORDER BY pg_relation_size(i.indexrelid) DESC;

-- Find missing indexes
SELECT
  schemaname || '.' || relname AS table,
  seq_scan,
  seq_tup_read,
  idx_scan,
  n_live_tup
FROM pg_stat_user_tables
WHERE seq_scan > 1000 AND n_live_tup > 10000
ORDER BY seq_tup_read DESC;
```

### Query Analysis

```sql
-- Enable pg_stat_statements
CREATE EXTENSION IF NOT EXISTS pg_stat_statements;

-- Find slow queries
SELECT
  query,
  calls,
  total_time / 1000 as total_seconds,
  mean_time as avg_ms,
  rows
FROM pg_stat_statements
ORDER BY total_time DESC
LIMIT 20;
```

## Maintenance Tasks

### Vacuum

Automatic vacuum is enabled, but monitor:

```sql
-- Check vacuum stats
SELECT
  schemaname,
  relname,
  last_vacuum,
  last_autovacuum,
  n_dead_tup
FROM pg_stat_user_tables
ORDER BY n_dead_tup DESC;
```

### Analyze

Run after bulk operations:

```sql
ANALYZE usage_records;
```

### Reindex

Schedule during low traffic:

```sql
REINDEX INDEX CONCURRENTLY idx_usage_org_period;
```

## Monitoring

### Key Metrics

| Metric          | Warning   | Critical  |
| --------------- | --------- | --------- |
| CPU             | > 70%     | > 90%     |
| Connections     | > 80% max | > 95% max |
| Disk            | > 80%     | > 90%     |
| Replication Lag | > 30s     | > 60s     |
| Query Time P99  | > 500ms   | > 1s      |

### Alerts

```yaml
alerts:
  - name: DatabaseHighCPU
    expr: aws_rds_cpuutilization_average > 70
    for: 10m
    severity: warning

  - name: DatabaseConnectionsHigh
    expr: aws_rds_database_connections_average / aws_rds_database_connections_maximum > 0.8
    for: 5m
    severity: warning
```

## Access Control

### Roles

| Role       | Access                   | Used By               |
| ---------- | ------------------------ | --------------------- |
| app_rw     | Read/write to app tables | Application           |
| app_ro     | Read-only                | Reporting, Analytics  |
| migrations | Schema changes           | CI/CD pipeline        |
| admin      | Full access              | DBAs (emergency only) |

```sql
-- Create application role
CREATE ROLE app_rw;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO app_rw;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO app_rw;
```

### Auditing

Enable RDS audit logging:

```
rds.force_ssl = 1
log_connections = on
log_disconnections = on
log_statement = 'ddl'
```
