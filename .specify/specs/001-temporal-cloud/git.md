# Git Management

## Repository Strategy

### Forked Repositories

| Repo                             | Upstream                                    | Purpose                          |
| -------------------------------- | ------------------------------------------- | -------------------------------- |
| temporal                         | temporalio/temporal                         | Core server + cloud interceptors |
| cloud-api                        | temporalio/cloud-api                        | Cloud API protos                 |
| tcld                             | temporalio/tcld                             | Cloud CLI                        |
| terraform-provider-temporalcloud | temporalio/terraform-provider-temporalcloud | Terraform                        |

### New Repositories

| Repo                    | Purpose                |
| ----------------------- | ---------------------- |
| temporal-cloud-platform | Backend services       |
| temporal-cloud-console  | Web UI                 |
| temporal-cloud-infra    | Infrastructure as Code |

## Branch Strategy

```
main                    ← Mirror of upstream (auto-sync)
├── cloud/main          ← Production releases
├── cloud/staging       ← Staging environment
├── cloud/develop       ← Development integration
└── feature/*           ← Feature branches
```

## Sync Policy

### Golden Rule

**NEVER modify existing upstream files. Only ADD new files/packages.**

### Allowed Changes (temporal fork)

```
✅ ADD: common/cloud/           (new package)
✅ ADD: common/cloud/metering/
✅ ADD: common/cloud/quota/
✅ ADD: common/cloud/audit/
❌ MODIFY: common/config/config.go
❌ MODIFY: service/frontend/fx.go
```

### Protected Paths

```yaml
# .github/protected-paths.yaml
temporal:
  protected:
    - "api/"
    - "cmd/server/"
    - "common/authorization/authorizer.go"
    - "common/config/config.go"
    - "service/frontend/workflow_handler.go"
    - "service/history/"
    - "service/matching/"
  allowed_additions:
    - "common/cloud/"
```

## Sync Automation

### Daily Sync Workflow

```yaml
name: Sync Upstream
on:
  schedule:
    - cron: "0 0 * * *"
  workflow_dispatch:

jobs:
  sync:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
        with:
          ref: cloud/main
          fetch-depth: 0
          token: ${{ secrets.GH_PAT }}

      - name: Setup Git
        run: |
          git config user.name "Sync Bot"
          git config user.email "bot@your-org.com"

      - name: Sync
        run: |
          git remote add upstream https://github.com/temporalio/temporal.git || true
          git fetch upstream
          git merge upstream/main --no-edit || exit 1
          git push origin cloud/main
```

## Pre-Commit Hooks

### Installation

```bash
# Install pre-commit
pip install pre-commit
pre-commit install
```

### Configuration

```yaml
# .pre-commit-config.yaml
repos:
  - repo: local
    hooks:
      - id: check-protected-paths
        name: Check Protected Paths
        entry: scripts/check-protected-paths.sh
        language: script
        pass_filenames: true

      - id: go-fmt
        name: Go Format
        entry: gofmt -w
        language: system
        files: \.go$

      - id: go-lint
        name: Go Lint
        entry: golangci-lint run
        language: system
        files: \.go$
```

### Protected Path Check Script

```bash
#!/bin/bash
# scripts/check-protected-paths.sh

PROTECTED=(
  "common/config/config.go"
  "common/authorization/authorizer.go"
  "service/frontend/fx.go"
  "service/history/"
  "service/matching/"
)

for file in "$@"; do
  for p in "${PROTECTED[@]}"; do
    if [[ "$file" == "$p"* ]]; then
      echo "❌ BLOCKED: Cannot modify protected path: $file"
      echo "   Add cloud code to common/cloud/ instead"
      exit 1
    fi
  done
done
```

## Branch Protection Rules

### cloud/main

- Require pull request reviews: 2
- Require status checks: lint, test, build
- Require branches to be up to date
- No force pushes
- No deletions

### cloud/develop

- Require pull request reviews: 1
- Require status checks: lint, test
- Allow force pushes (with lease)

## Commit Conventions

### Format

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Types

- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation
- `style`: Formatting
- `refactor`: Code restructuring
- `test`: Tests
- `chore`: Maintenance

### Example

```
feat(billing): add invoice generation workflow

Implements the monthly invoice generation workflow that:
- Calculates usage from metering data
- Generates line items for actions and storage
- Creates Stripe invoice

Closes #123
```
