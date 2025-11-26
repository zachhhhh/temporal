# QA Process

## QA Environments

| Environment | Purpose             | Data                | Access         |
| ----------- | ------------------- | ------------------- | -------------- |
| Dev         | Developer testing   | Seed data           | All engineers  |
| QA          | QA team testing     | Sanitized prod copy | QA + Engineers |
| Staging     | Pre-prod validation | Sanitized prod copy | All            |
| Prod        | Production          | Real                | Restricted     |

## QA Workflow

### Feature Development

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Dev       │────▶│   QA        │────▶│   Staging   │
│   Testing   │     │   Testing   │     │   Testing   │
└─────────────┘     └─────────────┘     └─────────────┘
     │                    │                    │
     ▼                    ▼                    ▼
  Unit Tests         Functional          E2E + Load
  Integration        Exploratory         UAT
```

### QA Entry Criteria

Before QA testing begins:

- [ ] Unit tests passing (>80% coverage)
- [ ] Integration tests passing
- [ ] Code review approved
- [ ] Feature deployed to QA environment
- [ ] Test data prepared
- [ ] Acceptance criteria documented

### QA Exit Criteria

Before promotion to staging:

- [ ] All test cases passed
- [ ] No P0/P1 bugs open
- [ ] Performance benchmarks met
- [ ] Security checklist completed
- [ ] Documentation reviewed

## Test Types

### 1. Functional Testing

Manual and automated testing of features.

**Test Case Template**:

```markdown
## TC-001: Create Namespace

**Preconditions**:

- User logged in with admin role
- No namespace with name "test-ns" exists

**Steps**:

1. Navigate to Namespaces
2. Click "Create Namespace"
3. Enter name: "test-ns"
4. Select region: "aws-us-east-1"
5. Click "Create"

**Expected Result**:

- Namespace created successfully
- Appears in namespace list
- Correct region shown

**Priority**: High
**Automated**: Yes
```

### 2. Exploratory Testing

Unscripted testing to find edge cases.

**Session Template**:

```markdown
## Exploratory Session: Billing Flow

**Charter**: Explore billing edge cases around plan upgrades mid-cycle

**Time Box**: 2 hours

**Areas to Explore**:

- Upgrade with pending invoice
- Upgrade then immediate downgrade
- Upgrade with low credit balance

**Findings**:

- [Bug] Proration calculation off by one day
- [Observation] Slow response when loading invoices
```

### 3. Regression Testing

Ensure existing features still work after changes.

**Regression Suite**:

- Core Workflows: 50 cases (automated)
- Billing: 30 cases (automated)
- Authentication: 25 cases (automated)
- Console UI: 100 cases (80% automated)

### 4. Performance Testing

```bash
# Run load test
k6 run --vus 100 --duration 10m load/api-test.js

# Expected Results
# - P99 latency < 200ms
# - Error rate < 0.1%
# - No memory leaks
```

### 5. Security Testing

- **SAST**: Static analysis on every PR (Semgrep)
- **DAST**: Weekly dynamic scanning (OWASP ZAP)
- **Dependency Scan**: On every build (Snyk)
- **Penetration Test**: Annual (third party)

## Test Data Management

### Seed Data

```sql
-- QA seed data
INSERT INTO organizations (id, name, slug) VALUES
  ('qa-org-1', 'QA Test Org 1', 'qa-test-1'),
  ('qa-org-2', 'QA Test Org 2', 'qa-test-2');

INSERT INTO users (id, email, name) VALUES
  ('qa-user-1', 'qa1@test.temporal.io', 'QA User 1'),
  ('qa-user-2', 'qa2@test.temporal.io', 'QA User 2');
```

### Data Sanitization

For staging/QA from production:

```bash
# Sanitize production data
pg_dump prod_db | \
  sed 's/[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}/redacted@example.com/g' | \
  psql qa_db
```

## Bug Reporting

### From QA Testing

```markdown
## Bug: [Short Title]

**Found In**: QA Environment
**Build**: v1.2.0-rc1
**Test Case**: TC-042

**Steps to Reproduce**:

1. ...

**Expected**: ...
**Actual**: ...

**Attachments**:

- Screenshot
- Console logs
- Network trace

**Severity**: P2
**Assigned To**: @engineer
```

## QA Metrics

### Quality Metrics

| Metric               | Target | Measurement                  |
| -------------------- | ------ | ---------------------------- |
| Test Coverage        | >80%   | Codecov                      |
| Bug Escape Rate      | <5%    | Bugs in prod / total bugs    |
| Automation Rate      | >70%   | Automated / total test cases |
| Regression Pass Rate | >99%   | Passing / total regression   |

### Release Quality Gate

Release blocked if:

- Test coverage < 80%
- Any P0/P1 bugs open
- Security scan has critical findings
- Performance regression > 10%

## Tools

| Purpose           | Tool              |
| ----------------- | ----------------- |
| Test Management   | TestRail / Linear |
| Automated Testing | Playwright, k6    |
| Bug Tracking      | Linear            |
| Security Scanning | Snyk, Semgrep     |
| Performance       | k6, Grafana       |
