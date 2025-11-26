# Testing Strategy

## Test Pyramid

```
        ┌───────────┐
        │   E2E     │  10%
        │   Tests   │
        ├───────────┤
        │Integration│  30%
        │   Tests   │
        ├───────────┤
        │   Unit    │  60%
        │   Tests   │
        └───────────┘
```

## Test Types

| Type        | Tool                        | Location       | Coverage Target |
| ----------- | --------------------------- | -------------- | --------------- |
| Unit (Go)   | `go test`                   | `*_test.go`    | 80%             |
| Unit (TS)   | Vitest                      | `*.test.ts`    | 80%             |
| Integration | `go test -tags integration` | `integration/` | Critical paths  |
| E2E         | Playwright                  | `e2e/`         | User journeys   |
| Load        | k6                          | `load/`        | Before release  |
| Security    | Trivy, Snyk                 | CI pipeline    | All images      |

## Unit Tests

### Go Tests

```bash
# Run all unit tests
make test

# Run with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run specific package
go test ./internal/billing/...
```

### TypeScript Tests

```bash
# Run all tests
pnpm test

# Run with coverage
pnpm test:coverage

# Run in watch mode
pnpm test:watch
```

## Integration Tests

### Database Tests

```go
//go:build integration

func TestCreateOrganization(t *testing.T) {
    db := setupTestDB(t)
    defer db.Close()

    store := NewOrganizationStore(db)
    org, err := store.Create(ctx, &Organization{Name: "test"})

    require.NoError(t, err)
    require.NotEmpty(t, org.ID)
}
```

### API Tests

```go
//go:build integration

func TestBillingAPI(t *testing.T) {
    server := setupTestServer(t)
    defer server.Close()

    client := NewBillingClient(server.URL)
    invoice, err := client.GetInvoice(ctx, "inv-123")

    require.NoError(t, err)
    require.Equal(t, "paid", invoice.Status)
}
```

## E2E Tests

### Playwright Setup

```typescript
// e2e/auth.spec.ts
import { test, expect } from "@playwright/test";

test("user can login via SSO", async ({ page }) => {
  await page.goto("/login");
  await page.click("text=Sign in with SSO");
  await page.fill("[name=email]", "test@example.com");
  await page.click("text=Continue");

  await expect(page).toHaveURL("/dashboard");
});
```

### Critical User Journeys

- [ ] User signup and onboarding
- [ ] Create namespace
- [ ] View usage dashboard
- [ ] Upgrade subscription
- [ ] Configure SSO
- [ ] Generate API key

## Load Tests

### k6 Script

```javascript
// load/billing.js
import http from "k6/http";
import { check, sleep } from "k6";

export const options = {
  stages: [
    { duration: "2m", target: 100 },
    { duration: "5m", target: 100 },
    { duration: "2m", target: 0 },
  ],
};

export default function () {
  const res = http.get("https://api.temporal-cloud.io/v1/usage");
  check(res, {
    "status is 200": (r) => r.status === 200,
    "response time < 200ms": (r) => r.timings.duration < 200,
  });
  sleep(1);
}
```

### Performance Targets

| Endpoint        | P50   | P99   | Max RPS |
| --------------- | ----- | ----- | ------- |
| GET /usage      | 50ms  | 200ms | 1000    |
| POST /namespace | 100ms | 500ms | 100     |
| GET /invoices   | 100ms | 300ms | 500     |

## Security Tests

### Container Scanning

```yaml
- name: Scan image
  uses: aquasecurity/trivy-action@master
  with:
    image-ref: "${{ env.IMAGE }}"
    severity: "CRITICAL,HIGH"
    exit-code: "1"
```

### Dependency Scanning

```yaml
- name: Snyk scan
  uses: snyk/actions/golang@master
  with:
    command: test
```

## Test Environments

| Environment | Purpose             | Data                |
| ----------- | ------------------- | ------------------- |
| Local       | Developer testing   | Seed data           |
| CI          | Automated testing   | Ephemeral           |
| Staging     | Pre-prod validation | Sanitized prod copy |
| Production  | Live                | Real                |

## Commands

```bash
# Run all tests
make test-all

# Unit tests only
make test

# Integration tests
make integration-test

# E2E tests
pnpm test:e2e

# Load tests
k6 run load/billing.js

# Security scan
make security-scan
```
