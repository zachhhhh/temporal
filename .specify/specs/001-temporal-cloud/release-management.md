# Release Management

## Versioning

### Semantic Versioning

`MAJOR.MINOR.PATCH`

- **MAJOR**: Breaking API changes
- **MINOR**: New features, backwards compatible
- **PATCH**: Bug fixes, security patches

### Version Components

| Component          | Version | Release Cadence |
| ------------------ | ------- | --------------- |
| Cloud Platform API | v1.x.x  | Monthly         |
| Cloud Console      | v1.x.x  | Weekly          |
| tcld CLI           | v1.x.x  | Monthly         |
| Terraform Provider | v0.x.x  | Monthly         |
| SDK Extensions     | v1.x.x  | As needed       |

## Release Process

### 1. Feature Freeze (T-5 days)

```bash
# Create release branch
git checkout cloud/develop
git checkout -b release/v1.2.0
git push origin release/v1.2.0
```

- No new features after this point
- Only bug fixes and polish

### 2. QA Validation (T-4 to T-2 days)

- Deploy release branch to staging
- Run full E2E test suite
- QA team performs manual testing
- Security scan

### 3. Release Candidate (T-2 days)

```bash
# Tag RC
git tag v1.2.0-rc1
git push origin v1.2.0-rc1
```

- Deploy RC to production (canary)
- Monitor metrics for 24 hours

### 4. Release (T-0)

```bash
# Merge to cloud/main
git checkout cloud/main
git merge release/v1.2.0
git tag v1.2.0
git push origin cloud/main --tags

# Update changelog
# Generate release notes
```

### 5. Post-Release

- Announce in #releases Slack
- Update documentation
- Email customers (major releases only)
- Monitor for issues

## Changelog

### Format (Keep a Changelog)

```markdown
# Changelog

## [1.2.0] - 2025-02-01

### Added

- Multi-region namespace support (#123)
- SCIM group sync (#456)

### Changed

- Improved billing dashboard performance

### Fixed

- Certificate expiry notification timing (#789)

### Security

- Updated dependencies for CVE-2025-1234
```

### Generation

```bash
# Auto-generate from commits
git log v1.1.0..v1.2.0 --pretty=format:"- %s (%h)" > CHANGELOG_DRAFT.md
```

## Rollout Strategy

### Canary Deployment

1. Deploy to 5% of traffic
2. Monitor error rates, latency for 1 hour
3. If healthy, increase to 25%
4. Wait 2 hours
5. Full rollout (100%)

```yaml
# ArgoCD rollout
apiVersion: argoproj.io/v1alpha1
kind: Rollout
spec:
  strategy:
    canary:
      steps:
        - setWeight: 5
        - pause: { duration: 1h }
        - setWeight: 25
        - pause: { duration: 2h }
        - setWeight: 100
      analysis:
        templates:
          - templateName: success-rate
        startingStep: 1
```

### Rollback Criteria

Automatic rollback if:

- Error rate > 1%
- P99 latency > 500ms
- Any P0/P1 bugs reported

```bash
# Manual rollback
helm rollback cloud-platform --namespace cloud-platform
```

## Hotfix Process

For critical production issues:

```bash
# Branch from cloud/main
git checkout cloud/main
git checkout -b hotfix/v1.2.1

# Fix the issue
# ... commit ...

# Tag and deploy immediately
git tag v1.2.1
git push origin v1.2.1

# Merge back
git checkout cloud/main && git merge hotfix/v1.2.1
git checkout cloud/develop && git merge hotfix/v1.2.1
```

## Release Checklist

### Pre-Release

- [ ] All tests passing
- [ ] Security scan clean
- [ ] Changelog updated
- [ ] Documentation updated
- [ ] Breaking changes documented
- [ ] Migration guide (if needed)
- [ ] Rollback plan confirmed

### Release

- [ ] Tag created
- [ ] Canary deployed
- [ ] Metrics monitored
- [ ] Full rollout complete

### Post-Release

- [ ] Release notes published
- [ ] Customers notified (if applicable)
- [ ] Retrospective scheduled (major releases)

## Communication

### Internal

- Slack #releases: All releases
- Slack #incidents: Hotfixes

### External

- Status page: Maintenance notices
- Email: Major releases, breaking changes
- Blog: Feature announcements
