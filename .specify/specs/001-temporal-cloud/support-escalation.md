# Support Escalation

## Support Tiers

| Tier             | Access               | Response Time   | Hours          |
| ---------------- | -------------------- | --------------- | -------------- |
| Community        | GitHub Issues        | Best effort     | Community      |
| Essential        | Email                | 2 business days | Business hours |
| Business         | Email + Chat         | 4 hours (P1)    | Business hours |
| Enterprise       | Email + Chat + Slack | 1 hour (P0)     | 24/7           |
| Mission Critical | Dedicated + Phone    | 15 min (P0)     | 24/7           |

## Priority Classification

| Priority | Description     | Example               |
| -------- | --------------- | --------------------- |
| P0       | Production down | All workflows failing |
| P1       | Major impact    | 50%+ error rate       |
| P2       | Moderate impact | Single feature broken |
| P3       | Minor impact    | UI bug, question      |

## Escalation Flow

```
┌─────────────────────────────────────────────────────────┐
│                    Customer                              │
└───────────────────────┬─────────────────────────────────┘
                        │
                        ▼
┌─────────────────────────────────────────────────────────┐
│              Tier 1: Support Agent                       │
│  • Answer common questions                               │
│  • Triage and categorize                                 │
│  • Collect diagnostic information                        │
│  • SLA: 15 min response (P0/P1)                         │
└───────────────────────┬─────────────────────────────────┘
                        │ Escalate if unresolved (30 min)
                        ▼
┌─────────────────────────────────────────────────────────┐
│              Tier 2: Support Engineer                    │
│  • Deep technical troubleshooting                        │
│  • Access to internal systems                            │
│  • Can reproduce issues                                  │
│  • SLA: Resolve or escalate (2 hours)                   │
└───────────────────────┬─────────────────────────────────┘
                        │ Escalate if code change needed
                        ▼
┌─────────────────────────────────────────────────────────┐
│              Tier 3: Engineering                         │
│  • Bug fixes                                             │
│  • Feature requests                                      │
│  • Architecture guidance                                 │
│  • SLA: Based on severity                               │
└─────────────────────────────────────────────────────────┘
```

## Ticket Lifecycle

```
┌─────────┐     ┌─────────┐     ┌─────────┐     ┌─────────┐
│  New    │────▶│  Open   │────▶│ Pending │────▶│ Solved  │
└─────────┘     └─────────┘     └─────────┘     └─────────┘
                     │               │
                     │               │ Customer unresponsive
                     │               ▼
                     │          ┌─────────┐
                     └─────────▶│ Closed  │
                                └─────────┘
```

## Information Gathering

### Required for P0/P1

```markdown
## Ticket Template

**Account ID**:
**Namespace(s)**:
**Region**:

**Impact**:

- Start time:
- Users affected:
- Workflows affected:

**Description**:
[What is happening vs what should happen]

**Steps to Reproduce**:

1.
2.

**Error Messages**:
[Exact error text or screenshot]

**Recent Changes**:
[Any deployments, config changes in last 24h]
```

### Diagnostic Commands

```bash
# Namespace health
temporal operator namespace describe --namespace $NS

# Recent failures
temporal workflow list --namespace $NS --query "CloseStatus='Failed'"

# Metrics snapshot
curl https://$NS.tmprl.cloud/prometheus/metrics | grep error
```

## Escalation Triggers

### Tier 1 → Tier 2

- Unable to resolve with documentation
- Requires internal system access
- Customer requests escalation
- Approaching SLA breach

### Tier 2 → Tier 3

- Confirmed bug requiring code change
- Infrastructure issue
- Security incident
- Feature request from Enterprise customer

### Tier 3 → Management

- P0 unresolved > 2 hours
- Multiple customers affected
- Data loss confirmed
- Security breach

## Communication Templates

### Initial Response (P0/P1)

```
Hi [Name],

Thank you for contacting Temporal Cloud Support.

I understand you're experiencing [issue summary]. This is being
treated as a Priority [X] issue, and I'm actively investigating.

To help me troubleshoot, please provide:
- [Specific info needed]

I will update you within [time] or sooner if I have new information.

Best,
[Agent Name]
```

### Escalation to Engineering

```
## Engineering Escalation

**Ticket**: #12345
**Customer**: Acme Corp (Enterprise)
**Priority**: P1

**Summary**:
Namespace creation failing with timeout errors since 14:00 UTC

**Impact**:
- 50+ create requests failed
- Customer unable to onboard new teams

**Troubleshooting Done**:
1. Verified namespace quotas (OK)
2. Checked region health (OK)
3. Reviewed logs (timeout to provisioning service)

**Hypothesis**:
Provisioning service may be overloaded

**Request**:
Engineering review of provisioning service logs
```

### Resolution

```
Hi [Name],

Great news - this issue has been resolved.

**Root Cause**: [Brief explanation]

**Resolution**: [What was done]

**Preventive Measures**: [If applicable]

Your namespaces should now be functioning normally. Please
confirm on your end and let me know if you encounter any
further issues.

Best,
[Agent Name]
```

## Knowledge Base

### Article Structure

```markdown
# [Title]

## Symptoms

What the customer sees

## Cause

Why this happens

## Solution

Step-by-step fix

## Prevention

How to avoid in future

## Related Articles

- [Link 1]
- [Link 2]
```

### Common Issues

| Issue               | Resolution                           |
| ------------------- | ------------------------------------ |
| Certificate expired | Guide to rotate certs                |
| Rate limited        | Explain limits, request increase     |
| Connection refused  | Check firewall, verify endpoint      |
| Slow workflows      | Review history size, Continue-As-New |

## Metrics

| Metric                   | Target    | Current |
| ------------------------ | --------- | ------- |
| First Response Time      | < 1h (P0) |         |
| Resolution Time          | < 4h (P0) |         |
| CSAT                     | > 4.5/5   |         |
| First Contact Resolution | > 60%     |         |
| Escalation Rate          | < 20%     |         |

## On-Call Support

### Schedule

- Primary: Support Engineer (24/7 for Enterprise)
- Backup: Second Support Engineer
- Engineering Escalation: On-call Engineer

### Handoff

At shift change:

1. Review open P0/P1 tickets
2. Summarize status and next steps
3. Introduce to customer if active incident
