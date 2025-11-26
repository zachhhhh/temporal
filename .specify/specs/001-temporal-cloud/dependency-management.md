# Dependency Management

## Dependency Types

| Type              | Examples              | Update Strategy                   |
| ----------------- | --------------------- | --------------------------------- |
| Upstream Temporal | temporal server, SDKs | Daily sync, careful review        |
| Infrastructure    | Terraform, Helm       | Monthly, test thoroughly          |
| Go modules        | grpc, stripe-go       | Weekly automated, review breaking |
| NPM packages      | React, Next.js        | Weekly automated, review major    |
| Security          | Any CVE               | Immediate                         |

## Go Dependencies

### go.mod Management

```bash
# Update all dependencies
go get -u ./...
go mod tidy

# Update specific dependency
go get -u github.com/stripe/stripe-go/v76

# Check for outdated
go list -u -m all
```

### Dependabot Configuration

```yaml
# .github/dependabot.yml
version: 2
updates:
  - package-ecosystem: "gomod"
    directory: "/"
    schedule:
      interval: "weekly"
    commit-message:
      prefix: "deps(go):"
    labels:
      - "dependencies"
      - "go"
    groups:
      temporal:
        patterns:
          - "go.temporal.io/*"
      grpc:
        patterns:
          - "google.golang.org/grpc*"
          - "google.golang.org/protobuf*"
```

### Version Pinning

```go
// go.mod
module github.com/YOUR_ORG/temporal-cloud-platform

go 1.22

require (
    go.temporal.io/api v1.29.0
    go.temporal.io/sdk v1.26.0
    github.com/stripe/stripe-go/v76 v76.0.0
)
```

## NPM Dependencies

### package.json Management

```bash
# Check outdated
pnpm outdated

# Update all (within semver)
pnpm update

# Update specific package
pnpm update react@latest
```

### Renovate Configuration

```json
{
  "extends": ["config:base"],
  "packageRules": [
    {
      "matchPackagePatterns": ["^@types/"],
      "groupName": "type definitions"
    },
    {
      "matchPackagePatterns": ["react", "next"],
      "groupName": "react ecosystem"
    }
  ],
  "schedule": ["after 10pm on sunday"]
}
```

## Temporal SDK Compatibility

### Version Matrix

Maintain compatibility between components:

```yaml
# version-matrix.yaml
cloud_platform: v1.2.0
temporal_server: v1.24.2
sdk_go: v1.29.0
sdk_java: v1.25.0
sdk_typescript: v1.10.0
sdk_python: v1.7.0
```

### Upgrade Process

1. **Test**: Run SDK compatibility tests
2. **Stage**: Deploy to staging
3. **Soak**: Run for 24 hours
4. **Prod**: Gradual rollout
5. **Document**: Update version matrix

## Security Vulnerabilities

### Scanning

```yaml
# .github/workflows/security-scan.yaml
name: Security Scan
on:
  push:
  schedule:
    - cron: "0 0 * * *"

jobs:
  scan:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Run Snyk
        uses: snyk/actions/golang@master
        with:
          command: test
          args: --severity-threshold=high
```

### Vulnerability Response

| Severity | Response Time | Action                         |
| -------- | ------------- | ------------------------------ |
| Critical | 24 hours      | Immediate patch, hotfix deploy |
| High     | 7 days        | Next release                   |
| Medium   | 30 days       | Regular release                |
| Low      | 90 days       | When convenient                |

### Patching Process

1. **Alert**: Snyk/Dependabot creates issue
2. **Assess**: Determine exploitability
3. **Patch**: Update dependency
4. **Test**: Run full test suite
5. **Deploy**: Based on severity
6. **Verify**: Confirm scan passes

## Vendoring

We **do not** vendor Go dependencies.

**Rationale**:

- `go mod` provides reproducibility
- Vendoring bloats repository
- Harder to audit for security

**Exception**: Fork and vendor if:

- Upstream is unmaintained
- Critical patch needed immediately

## License Compliance

### Allowed Licenses

| License        | Allowed   |
| -------------- | --------- |
| MIT            | ✅        |
| Apache 2.0     | ✅        |
| BSD 2/3-Clause | ✅        |
| ISC            | ✅        |
| MPL 2.0        | ⚠️ Review |
| GPL            | ❌        |
| AGPL           | ❌        |

### License Scanning

```bash
# Go
go-licenses check ./...

# NPM
npx license-checker --onlyAllow "MIT;Apache-2.0;BSD-2-Clause;BSD-3-Clause;ISC"
```

## Upgrade Checklist

Before upgrading major dependencies:

- [ ] Read changelog for breaking changes
- [ ] Check our usage of deprecated APIs
- [ ] Run full test suite
- [ ] Deploy to staging
- [ ] Monitor for 24 hours
- [ ] Deploy to production
- [ ] Update documentation

## Rollback

If dependency update causes issues:

```bash
# Go: Revert go.mod and go.sum
git checkout HEAD~1 -- go.mod go.sum
go mod download

# NPM: Revert package.json and lockfile
git checkout HEAD~1 -- package.json pnpm-lock.yaml
pnpm install
```
