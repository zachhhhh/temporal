# Environment Configuration

## Environment Hierarchy

```
┌─────────────────────────────────────────────────────────────────┐
│                         Production                               │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐              │
│  │  us-east-1  │  │  eu-west-1  │  │ ap-south-1  │              │
│  └─────────────┘  └─────────────┘  └─────────────┘              │
└─────────────────────────────────────────────────────────────────┘
                              ▲
┌─────────────────────────────────────────────────────────────────┐
│                         Staging                                  │
│  ┌─────────────┐                                                │
│  │  us-east-1  │  (Production mirror, sanitized data)          │
│  └─────────────┘                                                │
└─────────────────────────────────────────────────────────────────┘
                              ▲
┌─────────────────────────────────────────────────────────────────┐
│                           QA                                     │
│  ┌─────────────┐                                                │
│  │  us-east-1  │  (Testing, seed data)                         │
│  └─────────────┘                                                │
└─────────────────────────────────────────────────────────────────┘
                              ▲
┌─────────────────────────────────────────────────────────────────┐
│                      Development                                 │
│  ┌─────────────┐                                                │
│  │   Local     │  (Docker Compose)                              │
│  └─────────────┘                                                │
└─────────────────────────────────────────────────────────────────┘
```

## Environment Definitions

| Environment | Purpose             | Data           | Access      | Deploy              |
| ----------- | ------------------- | -------------- | ----------- | ------------------- |
| Local       | Developer machine   | Seed           | Developer   | Manual              |
| Dev         | Shared development  | Seed           | Engineering | On push             |
| QA          | QA testing          | Sanitized      | QA + Eng    | On merge to develop |
| Staging     | Pre-prod validation | Sanitized prod | All         | On merge to main    |
| Production  | Live service        | Real           | Restricted  | Manual approval     |

## Configuration Management

### Configuration Sources (Priority Order)

1. **Environment Variables** (highest priority)
2. **Config Files** (per-environment)
3. **Defaults** (in code)

### Configuration Structure

```yaml
# config/base.yaml (defaults)
server:
  port: 8080
  read_timeout: 30s
  write_timeout: 30s

database:
  max_connections: 100
  idle_connections: 10

logging:
  level: info
  format: json

features:
  enable_scim: false
  enable_multi_region: false
```

```yaml
# config/production.yaml (overrides)
server:
  port: 8080

database:
  max_connections: 500
  idle_connections: 50

logging:
  level: warn

features:
  enable_scim: true
  enable_multi_region: true
```

### Environment Variables

```bash
# Naming convention: TEMPORAL_CLOUD_{SECTION}_{KEY}
TEMPORAL_CLOUD_SERVER_PORT=8080
TEMPORAL_CLOUD_DATABASE_HOST=prod-db.internal
TEMPORAL_CLOUD_DATABASE_MAX_CONNECTIONS=500
TEMPORAL_CLOUD_LOGGING_LEVEL=warn
```

### Go Configuration Loading

```go
type Config struct {
    Server   ServerConfig   `yaml:"server"`
    Database DatabaseConfig `yaml:"database"`
    Logging  LoggingConfig  `yaml:"logging"`
    Features FeatureConfig  `yaml:"features"`
}

func LoadConfig(env string) (*Config, error) {
    // Load base config
    cfg := loadYAML("config/base.yaml")

    // Overlay environment-specific config
    envConfig := loadYAML(fmt.Sprintf("config/%s.yaml", env))
    mergeConfig(cfg, envConfig)

    // Override with environment variables
    applyEnvVars(cfg)

    // Validate
    if err := cfg.Validate(); err != nil {
        return nil, err
    }

    return cfg, nil
}
```

## Environment-Specific Settings

### Database

| Setting         | Local     | Dev    | QA    | Staging    | Prod    |
| --------------- | --------- | ------ | ----- | ---------- | ------- |
| Host            | localhost | dev-db | qa-db | staging-db | prod-db |
| Max Connections | 10        | 50     | 100   | 200        | 500     |
| SSL             | false     | true   | true  | true       | true    |
| Read Replicas   | 0         | 0      | 0     | 1          | 3       |

### Caching

| Setting      | Local     | Dev       | QA       | Staging       | Prod       |
| ------------ | --------- | --------- | -------- | ------------- | ---------- |
| Redis Host   | localhost | dev-redis | qa-redis | staging-redis | prod-redis |
| TTL          | 60s       | 60s       | 300s     | 300s          | 300s       |
| Cluster Mode | false     | false     | false    | true          | true       |

### External Services

| Service   | Local     | Dev       | QA        | Staging   | Prod      |
| --------- | --------- | --------- | --------- | --------- | --------- |
| Stripe    | Test mode | Test mode | Test mode | Test mode | Live mode |
| SendGrid  | Sandbox   | Sandbox   | Sandbox   | Live      | Live      |
| PagerDuty | Disabled  | Disabled  | Disabled  | Test      | Live      |

### Feature Flags

| Feature        | Local | Dev  | QA   | Staging | Prod     |
| -------------- | ----- | ---- | ---- | ------- | -------- |
| SCIM           | true  | true | true | true    | Per plan |
| Multi-region   | true  | true | true | true    | Per plan |
| New Billing UI | true  | true | true | 50%     | 10%      |

## Kubernetes ConfigMaps

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: cloud-platform-config
  namespace: cloud-platform
data:
  TEMPORAL_CLOUD_ENV: production
  TEMPORAL_CLOUD_REGION: us-east-1
  TEMPORAL_CLOUD_LOG_LEVEL: warn
  TEMPORAL_CLOUD_METRICS_ENABLED: "true"
```

### Per-Environment Kustomization

```yaml
# kustomization.yaml (base)
resources:
  - deployment.yaml
  - service.yaml
  - configmap.yaml

# overlays/production/kustomization.yaml
resources:
  - ../../base
patchesStrategicMerge:
  - deployment-patch.yaml
configMapGenerator:
  - name: cloud-platform-config
    behavior: merge
    literals:
      - TEMPORAL_CLOUD_ENV=production
      - TEMPORAL_CLOUD_REPLICAS=3
```

## Environment Promotion

### Promotion Flow

```
feature/* → develop → staging → production
    │          │          │          │
    ▼          ▼          ▼          ▼
  Local       Dev        QA      Staging → Prod
```

### Promotion Checklist

**Dev → QA**

- [ ] Unit tests pass
- [ ] Integration tests pass
- [ ] Code review approved

**QA → Staging**

- [ ] QA sign-off
- [ ] No P0/P1 bugs
- [ ] Performance benchmarks pass

**Staging → Production**

- [ ] 24h soak in staging
- [ ] Security scan clean
- [ ] Rollback plan confirmed
- [ ] On-call notified
- [ ] 2 approvals

## Environment Isolation

### Network Isolation

```hcl
# Each environment in separate VPC
module "vpc_prod" {
  source = "./modules/vpc"
  cidr   = "10.0.0.0/16"
  env    = "production"
}

module "vpc_staging" {
  source = "./modules/vpc"
  cidr   = "10.1.0.0/16"
  env    = "staging"
}

# No VPC peering between prod and non-prod
```

### Data Isolation

- Production data NEVER in non-prod environments
- Staging uses sanitized copy (PII removed)
- QA/Dev use synthetic seed data

### Access Isolation

| Environment | Access                                           |
| ----------- | ------------------------------------------------ |
| Production  | SRE, On-call (emergency), Senior Eng (read-only) |
| Staging     | All Engineering                                  |
| QA          | All Engineering, QA                              |
| Dev         | All Engineering                                  |

## Environment Debugging

### Local Development

```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f api

# Access database
psql postgres://temporal:temporal@localhost:5432/temporal_cloud
```

### Remote Environments

```bash
# Port forward to pod
kubectl port-forward svc/cloud-api 8080:8080 -n cloud-platform

# View logs
kubectl logs -f deployment/cloud-api -n cloud-platform

# Exec into pod
kubectl exec -it deployment/cloud-api -n cloud-platform -- /bin/sh
```

## Environment Health Checks

### Health Endpoints

| Endpoint   | Purpose                  |
| ---------- | ------------------------ |
| `/health`  | Basic liveness           |
| `/ready`   | Readiness (dependencies) |
| `/metrics` | Prometheus metrics       |

### Environment Dashboard

Each environment has a health dashboard showing:

- Service status (up/down)
- Error rates
- Latency percentiles
- Database connections
- Cache hit rate
- Queue depth
