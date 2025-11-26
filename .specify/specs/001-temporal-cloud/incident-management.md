# Incident Management

## Incident Lifecycle

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Detect    │────▶│   Respond   │────▶│   Mitigate  │
│             │     │             │     │             │
└─────────────┘     └─────────────┘     └─────────────┘
                                               │
                                               ▼
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Improve   │◀────│   Review    │◀────│   Resolve   │
│             │     │             │     │             │
└─────────────┘     └─────────────┘     └─────────────┘
```

## Severity Levels

| Level | Description         | Examples                           | Response Time |
| ----- | ------------------- | ---------------------------------- | ------------- |
| SEV1  | Complete outage     | All namespaces down, data loss     | 15 min        |
| SEV2  | Major degradation   | Region down, 50%+ error rate       | 30 min        |
| SEV3  | Partial degradation | Single feature broken, 10%+ errors | 2 hours       |
| SEV4  | Minor issue         | UI bug, minor latency increase     | 24 hours      |

## Detection

### Automatic (PagerDuty)

```yaml
# Alert rules that trigger incidents
alerts:
  - name: HighErrorRate
    condition: error_rate > 5%
    duration: 5m
    severity: SEV2

  - name: ServiceDown
    condition: up == 0
    duration: 1m
    severity: SEV1

  - name: HighLatency
    condition: p99_latency > 1s
    duration: 10m
    severity: SEV3
```

### Manual (Slack)

```
/incident create
Title: Users unable to create namespaces
Severity: SEV2
Description: Multiple customers reporting 500 errors
```

## Response

### Incident Commander (IC)

The first responder becomes IC until handoff.

**IC Responsibilities**:

- Coordinate response
- Communicate status
- Make decisions
- Delegate tasks

### Communication Channels

| Channel        | Purpose                        |
| -------------- | ------------------------------ |
| #incident-{id} | Real-time coordination         |
| #incidents     | All incident notifications     |
| Status Page    | External communication         |
| Email          | Customer notification (SEV1/2) |

### Status Page Updates

```
[Investigating] We are investigating issues with namespace creation.

[Identified] We have identified the root cause and are implementing a fix.

[Monitoring] A fix has been deployed. We are monitoring the situation.

[Resolved] This incident has been resolved.
```

## Roles

### During Incident

| Role                | Responsibility                 |
| ------------------- | ------------------------------ |
| Incident Commander  | Overall coordination           |
| Technical Lead      | Technical investigation        |
| Communications Lead | Status updates, customer comms |
| Scribe              | Document timeline, actions     |

### Escalation

```
On-Call Engineer (0 min)
      ↓
Tech Lead (15 min for SEV1)
      ↓
Engineering Manager (30 min for SEV1)
      ↓
VP Engineering (60 min for SEV1)
      ↓
CTO (customer-impacting SEV1)
```

## Response Procedures

### SEV1 Response

```
1. [0 min] Alert fires → On-call paged
2. [5 min] Acknowledge → Join #incident channel
3. [10 min] Assess impact → Confirm severity
4. [15 min] Assemble team → Page additional help if needed
5. [20 min] Initial status update → Post to status page
6. [30 min] Identify root cause OR escalate
7. [Ongoing] Updates every 15 minutes
```

### Common Actions

**Rollback**:

```bash
# Helm rollback
helm rollback cloud-platform -n cloud-platform

# Verify
kubectl get pods -n cloud-platform
```

**Feature Flag Disable**:

```bash
# Disable problematic feature
curl -X PUT https://config.internal/flags/new-billing-flow \
  -d '{"enabled": false}'
```

**Scale Up**:

```bash
# Scale replicas
kubectl scale deployment cloud-api --replicas=10 -n cloud-platform
```

**Traffic Shift**:

```bash
# Shift traffic to healthy region
aws route53 change-resource-record-sets ...
```

## Communication Templates

### Internal (Slack)

```
🚨 *INCIDENT STARTED* 🚨

*Incident ID*: INC-2025-001
*Severity*: SEV1
*Title*: API returning 500 errors
*IC*: @oncall-engineer
*Channel*: #incident-inc-2025-001

*Impact*: All API requests failing
*Customers Affected*: All

*Current Status*: Investigating
```

### External (Customer Email)

```
Subject: [Temporal Cloud Incident] Service Degradation

Dear Customer,

We are currently experiencing service issues affecting
Temporal Cloud. Our team is actively investigating.

Impact: [Description]
Start Time: [Time UTC]

We will provide updates every 30 minutes.

Current Status: [Status]

For real-time updates, visit: status.temporal.io

Temporal Cloud Team
```

## Post-Incident

### Timeline (Required for SEV1/2)

Within 24 hours:

- [ ] Detailed timeline documented
- [ ] Root cause identified
- [ ] Impact quantified

### Post-Mortem (Required for SEV1/2)

Within 5 business days:

- [ ] Post-mortem document completed
- [ ] Review meeting held
- [ ] Action items assigned

### Post-Mortem Template

```markdown
# Post-Mortem: INC-2025-001

## Summary

One-paragraph summary of what happened.

## Impact

- Duration: 45 minutes
- Customers affected: 1,234
- Namespaces affected: 5,678
- Revenue impact: $X,XXX
- SLA impact: Yes/No

## Timeline (all times UTC)

- 14:00 - Deployment started
- 14:05 - Error rate increased
- 14:10 - Alert fired, on-call paged
- 14:15 - Investigation started
- 14:25 - Root cause identified
- 14:30 - Rollback initiated
- 14:35 - Rollback complete
- 14:45 - Confirmed resolved

## Root Cause

Technical explanation of what went wrong.

## Contributing Factors

- Factor 1
- Factor 2

## What Went Well

- Fast detection (5 min)
- Clear runbook followed

## What Went Wrong

- Missing test case for edge condition
- Deployment during peak hours

## Action Items

| Action              | Owner | Due        |
| ------------------- | ----- | ---------- |
| Add regression test | @eng1 | 2025-02-01 |
| Update runbook      | @eng2 | 2025-02-01 |
| Add circuit breaker | @eng3 | 2025-02-15 |

## Lessons Learned

Key takeaways for the team.
```

## Metrics

| Metric                          | Target          |
| ------------------------------- | --------------- |
| MTTA (Mean Time To Acknowledge) | < 5 min         |
| MTTD (Mean Time To Detect)      | < 5 min         |
| MTTR (Mean Time To Resolve)     | < 1 hour (SEV1) |
| Post-mortem completion rate     | 100% (SEV1/2)   |
| Action item completion rate     | 100%            |
