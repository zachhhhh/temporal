# Disaster Recovery

## Recovery Objectives

| Metric                         | Target    |
| ------------------------------ | --------- |
| RPO (Recovery Point Objective) | 5 minutes |
| RTO (Recovery Time Objective)  | 1 hour    |

## Backup Strategy

### Database Backups

| Data                | Frequency  | Retention | Location        |
| ------------------- | ---------- | --------- | --------------- |
| PostgreSQL WAL      | Continuous | 35 days   | Cross-region S3 |
| PostgreSQL Snapshot | Daily      | 90 days   | Cross-region S3 |
| Redis Snapshot      | Hourly     | 7 days    | Same region S3  |

### Configuration Backups

| Data                 | Frequency | Retention | Location        |
| -------------------- | --------- | --------- | --------------- |
| Terraform state      | On change | Forever   | S3 + versioning |
| Kubernetes manifests | On change | Forever   | Git             |
| Secrets              | On change | 90 days   | AWS Backup      |

### Audit Log Archive

| Data         | Frequency | Retention | Location                |
| ------------ | --------- | --------- | ----------------------- |
| Audit events | Real-time | 7 years   | Cross-region S3 Glacier |

## Recovery Procedures

### Scenario 1: Data Corruption

1. Identify corruption scope and time
2. Stop writes to affected tables
3. Restore from point-in-time backup
4. Validate data integrity
5. Resume operations
6. Post-incident review

**RTO**: 30 minutes  
**RPO**: 5 minutes

### Scenario 2: Single Region Failure

1. Confirm region is unavailable
2. Trigger DNS failover to secondary
3. Verify secondary region health
4. Notify customers of degraded service
5. Begin data sync when primary recovers
6. Failback when primary is stable

**RTO**: 15 minutes  
**RPO**: 5 minutes

### Scenario 3: Complete Platform Failure

1. Activate incident response team
2. Provision new infrastructure from Terraform
3. Restore database from latest backup
4. Deploy applications from CI/CD
5. Validate all services
6. Update DNS to new infrastructure
7. Comprehensive post-mortem

**RTO**: 4 hours  
**RPO**: 5 minutes

## DR Testing

### Test Schedule

| Test Type          | Frequency | Duration |
| ------------------ | --------- | -------- |
| Backup restoration | Monthly   | 2 hours  |
| Failover drill     | Quarterly | 4 hours  |
| Full DR exercise   | Annually  | 1 day    |

### Test Checklist

- [ ] Verify backup integrity
- [ ] Test restore procedure
- [ ] Validate data consistency
- [ ] Test failover automation
- [ ] Measure actual RTO/RPO
- [ ] Document lessons learned

## Runbook Location

All detailed runbooks are in:
`temporal-cloud-infra/docs/runbooks/`

- `disaster-recovery.md`
- `database-restore.md`
- `region-failover.md`
- `incident-response.md`
