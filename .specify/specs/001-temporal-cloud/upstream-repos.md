# Upstream Repository Management

## Overview

Temporal has 190+ repositories at github.com/temporalio. This document defines which ones we fork, mirror, or ignore, and how we manage sync.

## Repository Classification

### Category 1: FORK (Modify with Cloud Code)

We fork these and add cloud-specific code. Sync daily.

| Repo                               | Purpose          | Cloud Changes                            |
| ---------------------------------- | ---------------- | ---------------------------------------- |
| `temporal`                         | Core server      | Metering interceptors, quota enforcement |
| `cloud-api`                        | Cloud API protos | Extended with billing/org protos         |
| `tcld`                             | Cloud CLI        | Already cloud-native                     |
| `terraform-provider-temporalcloud` | Terraform        | Already cloud-native                     |
| `ui`                               | Web UI           | Embed in console, add cloud features     |

### Category 2: MIRROR (Use as-is, track releases)

We don't modify these, but we need specific versions for compatibility.

| Repo             | Purpose       | How We Use            |
| ---------------- | ------------- | --------------------- |
| `sdk-go`         | Go SDK        | Workers, client code  |
| `sdk-java`       | Java SDK      | Customer SDK          |
| `sdk-typescript` | TS SDK        | Customer SDK, console |
| `sdk-python`     | Python SDK    | Customer SDK          |
| `sdk-dotnet`     | .NET SDK      | Customer SDK          |
| `sdk-php`        | PHP SDK       | Customer SDK          |
| `sdk-ruby`       | Ruby SDK      | Customer SDK          |
| `api`            | API protos    | Generated clients     |
| `cli`            | Temporal CLI  | Bundled for customers |
| `docker-builds`  | Docker images | Base images           |
| `helm-charts`    | Helm charts   | Deployment base       |
| `samples-*`      | SDK samples   | Customer docs         |

### Category 3: REFERENCE (Consult, don't use directly)

| Repo            | Purpose          |
| --------------- | ---------------- |
| `documentation` | Docs source      |
| `proposals`     | Design proposals |
| `roadmap`       | Public roadmap   |
| `features`      | Feature flags    |

### Category 4: IGNORE

Test repos, archived repos, one-off experiments.

## Fork Management

### Directory Structure

```
github.com/YOUR_ORG/
├── temporal/                    # Fork of temporalio/temporal
├── cloud-api/                   # Fork of temporalio/cloud-api
├── tcld/                        # Fork of temporalio/tcld
├── terraform-provider/          # Fork
├── ui/                          # Fork of temporalio/ui
├── temporal-cloud-platform/     # NEW - Backend services
├── temporal-cloud-console/      # NEW - Web console
├── temporal-cloud-infra/        # NEW - Infrastructure
└── temporal-cloud-sdk-go/       # NEW - SDK extensions
```

### Branch Structure (Forked Repos)

```
main                 ← Mirrors upstream/main (read-only)
├── cloud/main       ← Production cloud code
├── cloud/staging    ← Staging
├── cloud/develop    ← Development
└── release/v1.x     ← Release branches
```

### Sync Automation

```yaml
# .github/workflows/sync-upstream.yaml
name: Sync Upstream
on:
  schedule:
    - cron: "0 0 * * *" # Daily at midnight UTC
  workflow_dispatch:

jobs:
  sync:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        repo: [temporal, cloud-api, tcld, ui]
    steps:
      - uses: actions/checkout@v4
        with:
          repository: YOUR_ORG/${{ matrix.repo }}
          ref: main
          fetch-depth: 0
          token: ${{ secrets.SYNC_PAT }}

      - name: Add upstream
        run: |
          git remote add upstream https://github.com/temporalio/${{ matrix.repo }}.git
          git fetch upstream

      - name: Sync main branch
        run: |
          git checkout main
          git reset --hard upstream/main
          git push origin main --force

      - name: Merge to cloud/main
        run: |
          git checkout cloud/main
          git merge main -m "Sync upstream $(date +%Y-%m-%d)" || {
            echo "CONFLICT DETECTED"
            # Create PR for manual resolution
            gh pr create \
              --title "Upstream sync conflict $(date +%Y-%m-%d)" \
              --body "Automatic merge failed. Manual resolution required." \
              --head sync-conflict-$(date +%Y%m%d) \
              --base cloud/main
            exit 1
          }
          git push origin cloud/main

      - name: Notify on failure
        if: failure()
        run: |
          curl -X POST ${{ secrets.SLACK_WEBHOOK }} \
            -d '{"text":"Upstream sync failed for ${{ matrix.repo }}"}'
```

### Conflict Resolution Policy

1. **Never modify upstream files** - All cloud code in `common/cloud/` or new files
2. **If conflict in upstream file** - We made a mistake; refactor our code
3. **If upstream breaks our code** - Pin to last working version, open issue

### Release Tracking

Track upstream releases for compatibility:

```yaml
# .github/dependabot.yml
version: 2
updates:
  - package-ecosystem: "gomod"
    directory: "/"
    schedule:
      interval: "weekly"
    allow:
      - dependency-name: "go.temporal.io/*"
```

## SDK Version Matrix

| Cloud Version | temporal Server | sdk-go | sdk-java | sdk-typescript |
| ------------- | --------------- | ------ | -------- | -------------- |
| 1.0.0         | 1.24.x          | 1.29.x | 1.25.x   | 1.10.x         |
| 1.1.0         | 1.25.x          | 1.30.x | 1.26.x   | 1.11.x         |

## Monitoring Upstream

### Release Notifications

```yaml
# .github/workflows/watch-releases.yaml
name: Watch Upstream Releases
on:
  schedule:
    - cron: "0 */6 * * *" # Every 6 hours

jobs:
  check:
    runs-on: ubuntu-latest
    steps:
      - name: Check temporal releases
        run: |
          LATEST=$(gh release view --repo temporalio/temporal --json tagName -q .tagName)
          CURRENT=$(cat .upstream-versions/temporal)
          if [ "$LATEST" != "$CURRENT" ]; then
            echo "New temporal release: $LATEST"
            # Create issue for upgrade
            gh issue create \
              --title "Upgrade temporal to $LATEST" \
              --body "New upstream release available" \
              --label "upstream-upgrade"
          fi
```

### Security Advisories

Subscribe to GitHub Security Advisories for all temporalio repos.
Route to #security-alerts Slack channel.
