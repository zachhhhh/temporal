# CI/CD Pipeline

## Pipeline Overview

```
┌─────────┐   ┌─────────┐   ┌─────────┐   ┌─────────┐   ┌─────────┐
│  Lint   │──▶│  Test   │──▶│  Build  │──▶│ Deploy  │──▶│ Verify  │
│         │   │         │   │         │   │ Staging │   │         │
└─────────┘   └─────────┘   └─────────┘   └─────────┘   └─────────┘
                                                │
                                                ▼
                                          ┌─────────┐
                                          │ Deploy  │
                                          │  Prod   │
                                          └─────────┘
```

## Workflow Files

### ci.yaml

```yaml
name: CI
on:
  push:
    branches: [cloud/main, cloud/develop]
  pull_request:
    branches: [cloud/main]

jobs:
  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: "1.22"
      - run: make lint

  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: "1.22"
      - run: make test

  build:
    needs: [lint, test]
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - run: make build
      - uses: docker/build-push-action@v5
        with:
          push: ${{ github.ref == 'refs/heads/cloud/main' }}
          tags: ${{ env.REGISTRY }}/${{ env.IMAGE }}:${{ github.sha }}
```

### deploy-staging.yaml

```yaml
name: Deploy Staging
on:
  push:
    branches: [cloud/develop]

jobs:
  deploy:
    runs-on: ubuntu-latest
    environment: staging
    steps:
      - uses: actions/checkout@v4
      - uses: aws-actions/configure-aws-credentials@v4
      - run: |
          helm upgrade --install cloud-platform ./charts/cloud-platform \
            --namespace cloud-platform \
            --set image.tag=${{ github.sha }}
```

### deploy-prod.yaml

```yaml
name: Deploy Production
on:
  workflow_dispatch:
    inputs:
      version:
        description: "Version to deploy"
        required: true

jobs:
  deploy:
    runs-on: ubuntu-latest
    environment: production
    steps:
      - uses: actions/checkout@v4
      - uses: aws-actions/configure-aws-credentials@v4
      - run: |
          helm upgrade --install cloud-platform ./charts/cloud-platform \
            --namespace cloud-platform \
            --set image.tag=${{ inputs.version }}
```

### sync-upstream.yaml

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
          fetch-depth: 0
      - run: |
          git remote add upstream https://github.com/temporalio/temporal.git
          git fetch upstream
          git checkout cloud/main
          git merge upstream/main -m "Sync upstream"
          git push origin cloud/main
```

## Environments

| Environment | Trigger               | Approval    | URL                       |
| ----------- | --------------------- | ----------- | ------------------------- |
| Development | Push to feature/\*    | None        | dev.temporal-cloud.io     |
| Staging     | Push to cloud/develop | None        | staging.temporal-cloud.io |
| Production  | Manual                | 2 approvers | temporal-cloud.io         |

## Deployment Strategy

### Rolling Update

```yaml
strategy:
  type: RollingUpdate
  rollingUpdate:
    maxSurge: 1
    maxUnavailable: 0
```

### Canary (Future)

- Deploy to 10% of traffic
- Monitor error rates
- Gradually increase to 100%
- Automatic rollback on errors

## Rollback Procedure

### Automatic Rollback

```yaml
# Helm rollback on failed deployment
- run: |
    helm upgrade ... || helm rollback cloud-platform
```

### Manual Rollback

```bash
# List releases
helm history cloud-platform -n cloud-platform

# Rollback to previous
helm rollback cloud-platform 1 -n cloud-platform
```

## Secrets Management

### GitHub Secrets

- `AWS_ACCESS_KEY_ID`
- `AWS_SECRET_ACCESS_KEY`
- `DOCKER_REGISTRY_TOKEN`
- `SLACK_WEBHOOK_URL`

### Runtime Secrets

- Stored in AWS Secrets Manager
- Injected via External Secrets Operator
- Rotated every 90 days
