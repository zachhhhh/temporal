# Quickstart Guide

## Prerequisites

```bash
# Install required tools
brew install go node@20 docker kubectl helm terraform
brew install bufbuild/buf/buf golangci-lint gh
npm install -g pnpm

# Verify versions
go version      # 1.22+
node --version  # 20.x
docker --version
kubectl version --client
helm version
terraform version
```

## 1. Clone Repositories

```bash
mkdir -p ~/temporal-cloud && cd ~/temporal-cloud

# Clone all repos
gh repo clone YOUR_ORG/temporal
gh repo clone YOUR_ORG/cloud-api
gh repo clone YOUR_ORG/temporal-cloud-platform
gh repo clone YOUR_ORG/temporal-cloud-console
gh repo clone YOUR_ORG/temporal-cloud-infra

# Setup upstream remotes for forks
cd temporal && git remote add upstream https://github.com/temporalio/temporal.git
cd ../cloud-api && git remote add upstream https://github.com/temporalio/cloud-api.git
```

## 2. Start Local Environment

### Docker Compose

```yaml
# docker-compose.dev.yaml
version: "3.8"

services:
  postgresql:
    image: postgres:15
    environment:
      POSTGRES_USER: temporal
      POSTGRES_PASSWORD: temporal
      POSTGRES_DB: temporal_cloud
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

  redis:
    image: redis:7
    ports:
      - "6379:6379"

  temporal:
    image: temporalio/auto-setup:latest
    ports:
      - "7233:7233"
    environment:
      - DB=postgresql
      - DB_PORT=5432
      - POSTGRES_USER=temporal
      - POSTGRES_PWD=temporal
      - POSTGRES_SEEDS=postgresql
    depends_on:
      - postgresql

  temporal-ui:
    image: temporalio/ui:latest
    ports:
      - "8080:8080"
    environment:
      - TEMPORAL_ADDRESS=temporal:7233
    depends_on:
      - temporal

volumes:
  postgres_data:
```

```bash
# Start infrastructure
docker-compose -f docker-compose.dev.yaml up -d

# Verify services
docker-compose ps
```

## 3. Run Migrations

```bash
cd temporal-cloud-platform

# Install golang-migrate
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Run migrations
migrate -path schema -database "postgres://temporal:temporal@localhost:5432/temporal_cloud?sslmode=disable" up
```

## 4. Start Backend

```bash
cd temporal-cloud-platform

# Install dependencies
go mod download

# Run the server
go run cmd/server/main.go
```

## 5. Start Frontend

```bash
cd temporal-cloud-console

# Install dependencies
pnpm install

# Start dev server
pnpm dev
```

## 6. Access Services

| Service       | URL                   |
| ------------- | --------------------- |
| Temporal UI   | http://localhost:8080 |
| Cloud Console | http://localhost:3000 |
| Cloud API     | http://localhost:8081 |
| PostgreSQL    | localhost:5432        |
| Redis         | localhost:6379        |

## 7. Run Tests

```bash
# Backend tests
cd temporal-cloud-platform
make test

# Frontend tests
cd temporal-cloud-console
pnpm test

# Integration tests
make integration-test
```

## 8. Generate Code

```bash
# Generate proto code
cd cloud-api
buf generate

# Generate TypeScript types
cd temporal-cloud-console
pnpm generate
```

## Common Issues

### Port Already in Use

```bash
# Find process using port
lsof -i :5432

# Kill process
kill -9 <PID>
```

### Database Connection Failed

```bash
# Check PostgreSQL is running
docker-compose ps postgresql

# Check logs
docker-compose logs postgresql
```

### Proto Generation Failed

```bash
# Ensure buf is installed
buf --version

# Clear cache and retry
buf mod update
buf generate
```
