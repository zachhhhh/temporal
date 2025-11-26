# Bug Triage & Issue Management

## Issue Sources

1. **Customer Support** - Zendesk tickets escalated to engineering
2. **Internal** - Found by team during development/testing
3. **Monitoring** - Automated alerts that indicate bugs
4. **Security** - Vulnerability reports

## Issue Classification

### Severity

| Severity | Description                             | Response | Resolution |
| -------- | --------------------------------------- | -------- | ---------- |
| P0       | Production down, data loss              | 15 min   | 4 hours    |
| P1       | Major feature broken, workaround exists | 1 hour   | 24 hours   |
| P2       | Feature partially broken                | 4 hours  | 1 week     |
| P3       | Minor issue, cosmetic                   | 1 day    | 1 month    |
| P4       | Nice to have, improvement               | 1 week   | Backlog    |

### Type Labels

| Label               | Description                    |
| ------------------- | ------------------------------ |
| `bug`               | Something is broken            |
| `security`          | Security vulnerability         |
| `performance`       | Performance degradation        |
| `regression`        | Previously working, now broken |
| `customer-reported` | From customer support          |

### Component Labels

| Label        | Owner Team     |
| ------------ | -------------- |
| `billing`    | Platform       |
| `auth`       | Platform       |
| `namespaces` | Platform       |
| `console`    | Frontend       |
| `api`        | Platform       |
| `infra`      | Infrastructure |

## Triage Process

### Daily Triage Meeting (15 min)

**Attendees**: On-call engineer, Tech Lead, PM (optional)

**Agenda**:

1. Review new issues (5 min)
2. Assign severity and owner (5 min)
3. Review P0/P1 progress (5 min)

### Triage Workflow

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   New       │────▶│   Triage    │────▶│  Assigned   │
│   Issue     │     │   Meeting   │     │             │
└─────────────┘     └─────────────┘     └─────────────┘
                                               │
                    ┌─────────────┐             │
                    │   Closed    │◀────────────┤
                    │  (Invalid)  │             │
                    └─────────────┘             ▼
                                        ┌─────────────┐
┌─────────────┐     ┌─────────────┐     │    In       │
│   Closed    │◀────│   Review    │◀────│  Progress   │
│  (Resolved) │     │             │     │             │
└─────────────┘     └─────────────┘     └─────────────┘
```

## Issue Template

```markdown
## Bug Report

### Description

[Clear description of the issue]

### Steps to Reproduce

1.
2.
3.

### Expected Behavior

[What should happen]

### Actual Behavior

[What actually happens]

### Environment

- Account ID:
- Namespace:
- Region:
- SDK Version:
- Browser (if UI):

### Logs/Screenshots

[Attach relevant logs or screenshots]

### Impact

[Number of customers affected, revenue impact if known]
```

## Bug Fix Workflow

### 1. Reproduce

```bash
# Create test case that reproduces the bug
func TestBug123_NamespaceCreationFails(t *testing.T) {
    // This should fail until bug is fixed
}
```

### 2. Root Cause Analysis

Document in the issue:

- What went wrong
- Why it wasn't caught earlier
- Related code/commits

### 3. Fix

- Branch from `cloud/develop`
- Implement fix
- Add regression test
- Update any affected documentation

### 4. Review

- Code review required
- QA verification required for P0/P1
- Security review if security-related

### 5. Deploy

- P0/P1: Hotfix process (immediate)
- P2/P3: Next regular release
- P4: When convenient

## Escalation Path

```
On-Call Engineer
      │
      ▼
Tech Lead (after 30 min for P0)
      │
      ▼
Engineering Manager (after 1 hour for P0)
      │
      ▼
VP Engineering (after 2 hours for P0)
```

## Metrics

### Bug Metrics (Monthly Review)

- **MTTR** (Mean Time To Resolve): Target < 24h for P1
- **Bug Escape Rate**: Bugs found in production vs testing
- **Regression Rate**: Bugs that reoccur
- **Customer-Reported Ratio**: % of bugs from customers

### SLA Compliance

Track resolution time against severity SLAs.
Alert if approaching SLA breach.

## Post-Mortem

Required for all P0 and P1 bugs.

### Template

```markdown
## Post-Mortem: [Title]

### Summary

[1-2 sentence summary]

### Timeline

- HH:MM - Issue detected
- HH:MM - Investigation started
- HH:MM - Root cause identified
- HH:MM - Fix deployed
- HH:MM - Confirmed resolved

### Root Cause

[Technical explanation]

### Impact

- Duration: X hours
- Customers affected: N
- Revenue impact: $X

### What Went Well

-

### What Went Wrong

-

### Action Items

- [ ] [Action 1] - Owner - Due Date
- [ ] [Action 2] - Owner - Due Date
```
