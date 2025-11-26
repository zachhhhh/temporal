# Automated Repository Sync & Management

## Overview

We must maintain synchronization with 190+ repositories in `temporalio` organization, including automatically discovering and syncing new ones.

## Architecture

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│  GitHub     │────▶│   Sync      │────▶│   Our       │
│  API        │     │   Bot       │     │   Org       │
└─────────────┘     └─────────────┘     └─────────────┘
```

## Automated Discovery

### Sync Bot Job

Runs daily (or triggered by webhook):

```go
func DiscoverAndSync(ctx context.Context) error {
    // 1. List all upstream repos
    upstreamRepos, err := github.ListOrgRepos("temporalio")

    // 2. List our repos
    ourRepos, err := github.ListOrgRepos("YOUR_ORG")

    // 3. Find new repos
    for _, upstream := range upstreamRepos {
        if !contains(ourRepos, upstream.Name) {
            // New repo found!
            handleNewRepo(ctx, upstream)
        } else {
            // Existing repo - check for updates
            syncRepo(ctx, upstream)
        }
    }

    return nil
}
```

### Handling New Repos

When a new repo appears in `temporalio`:

1. **Classify**: Determine repo type (Go SDK, Java SDK, Core, Tool, etc.) based on languages and topics.
2. **Fork/Mirror**:
   - If it's a core component we modify -> **Fork**
   - If it's a dependency/SDK -> **Mirror**
3. **Configure**: Apply standard branch protection, CI workflows, and team access.
4. **Notify**: Alert #engineering-ops about the new repo.

## Sync Strategy

### 1. Mirror Repos (Read-Only)

For repos we don't modify (e.g. `sdk-java`, `samples-go`):

```yaml
# .github/workflows/mirror-sync.yaml
name: Mirror Sync
on:
  schedule:
    - cron: "0 */4 * * *" # Every 4 hours

jobs:
  sync:
    runs-on: ubuntu-latest
    steps:
      - name: Sync
        run: |
          git remote add upstream https://github.com/temporalio/${REPO_NAME}.git
          git fetch upstream
          git checkout main
          git merge upstream/main --ff-only
          git push origin main
          # Sync tags too
          git fetch upstream --tags
          git push origin --tags
```

### 2. Forked Repos (Modified)

For repos we patch (e.g. `temporal`, `ui`):

```bash
# Logic:
# 1. Fetch upstream/main
# 2. Rebase cloud/main onto upstream/main? NO.
# 3. Merge upstream/main into cloud/main? YES.

git checkout cloud/main
git merge upstream/main -m "Sync upstream"

if conflict:
    create_pr_for_manual_resolution()
else:
    push origin cloud/main
    trigger_build()
```

## Release Synchronization

### Automatic Release Detection

When `temporalio/temporal` tags `v1.25.0`:

1. **Detect**: Webhook or polling sees new tag.
2. **Branch**: Create `release/v1.25.0` from `upstream/v1.25.0`.
3. **Apply Patches**: Cherry-pick our cloud-specific commits (if structured as patches) OR merge `cloud/main` features.
4. **Build**: Trigger release build.
5. **Test**: Run integration tests.
6. **Promote**: If tests pass, mark as `cloud-v1.25.0` ready for staging.

## Conflict Management

### Prevention

- **Strict Isolation**: Cloud code goes in `common/cloud/` directory.
- **Interfaces**: Use Go interfaces to inject cloud logic, avoiding changes to core files.
- **Interceptors**: Use gRPC interceptors instead of modifying handlers.

### Resolution

If auto-sync fails:

1. **Alert**: "Sync failed for `temporal`: Merge conflict in `service/history/workflow.go`"
2. **PR**: Bot creates a PR with the conflict markers.
3. **Block**: Deployments blocked until resolved.

## Tooling

### Repo Config as Code

Manage settings for all 190+ repos centrally:

```yaml
# repos.yaml
defaults:
  has_issues: false
  has_projects: false
  has_wiki: false
  allow_squash_merge: true
  allow_merge_commit: false
  allow_rebase_merge: false

repos:
  temporal:
    has_issues: true
    protected_branches:
      main:
        required_reviews: 1
      cloud/main:
        required_reviews: 2
```

Run `github-settings-sync` tool to apply.

## Metrics

- **Sync Latency**: Time from upstream commit to our repo (Target: < 1h)
- **Sync Failures**: Number of manual interventions required (Target: < 1/week)
- **Repo Coverage**: % of upstream repos mirrored (Target: 100%)
