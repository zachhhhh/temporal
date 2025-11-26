# State Management

## State Types

| State Type        | Storage        | Scope              |
| ----------------- | -------------- | ------------------ |
| Terraform State   | S3 + DynamoDB  | Infrastructure     |
| Kubernetes State  | etcd (managed) | Cluster            |
| Application State | PostgreSQL     | Business data      |
| Session State     | Redis          | User sessions      |
| Workflow State    | Temporal       | Workflow execution |
| Cache State       | Redis          | Ephemeral          |

## Terraform State Management

### Backend Configuration

```hcl
# backend.tf
terraform {
  backend "s3" {
    bucket         = "temporal-cloud-terraform-state"
    key            = "environments/production/us-east-1/terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
    dynamodb_table = "terraform-state-lock"

    # Cross-region replication for DR
    # Configured at bucket level
  }
}
```

### State Structure

```
s3://temporal-cloud-terraform-state/
├── global/
│   ├── iam/terraform.tfstate
│   ├── route53/terraform.tfstate
│   └── budgets/terraform.tfstate
├── environments/
│   ├── production/
│   │   ├── us-east-1/terraform.tfstate
│   │   ├── eu-west-1/terraform.tfstate
│   │   └── gcp-us-central1/terraform.tfstate
│   ├── staging/
│   │   └── us-east-1/terraform.tfstate
│   └── dev/
│       └── us-east-1/terraform.tfstate
└── modules/  # No state here, just source
```

### State Locking

```hcl
resource "aws_dynamodb_table" "terraform_lock" {
  name         = "terraform-state-lock"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }

  # Point-in-time recovery
  point_in_time_recovery {
    enabled = true
  }
}
```

### State Operations

```bash
# List resources in state
terraform state list

# Show specific resource
terraform state show aws_instance.api

# Move resource (refactoring)
terraform state mv aws_instance.api aws_instance.cloud_api

# Remove from state (without destroying)
terraform state rm aws_instance.old_api

# Import existing resource
terraform import aws_instance.api i-1234567890abcdef0

# Refresh state
terraform refresh
```

### State Backup

```yaml
# S3 bucket versioning for state history
resource "aws_s3_bucket_versioning" "state" {
  bucket = aws_s3_bucket.terraform_state.id
  versioning_configuration {
    status = "Enabled"
  }
}

# Cross-region replication
resource "aws_s3_bucket_replication_configuration" "state" {
  bucket = aws_s3_bucket.terraform_state.id

  rule {
    status = "Enabled"
    destination {
      bucket        = aws_s3_bucket.terraform_state_dr.arn
      storage_class = "STANDARD"
    }
  }
}
```

## Application State (PostgreSQL)

### Schema Versioning

Track schema version in database:

```sql
CREATE TABLE schema_migrations (
    version BIGINT PRIMARY KEY,
    applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

### Data Integrity

```sql
-- Use transactions for state changes
BEGIN;
    UPDATE subscriptions SET plan = 'business' WHERE org_id = $1;
    INSERT INTO audit_events (org_id, action, details) VALUES ($1, 'plan_upgrade', $2);
COMMIT;
```

### Soft Deletes

Never hard delete, use soft deletes for audit trail:

```sql
ALTER TABLE organizations ADD COLUMN deleted_at TIMESTAMPTZ;

-- "Delete" organization
UPDATE organizations SET deleted_at = NOW() WHERE id = $1;

-- Query only active
SELECT * FROM organizations WHERE deleted_at IS NULL;
```

## Session State (Redis)

### Session Structure

```json
{
  "session_id": "sess_abc123",
  "user_id": "user_456",
  "org_id": "org_789",
  "created_at": "2025-01-15T10:00:00Z",
  "expires_at": "2025-01-16T10:00:00Z",
  "ip": "1.2.3.4",
  "user_agent": "Mozilla/5.0..."
}
```

### Session Commands

```go
// Create session
func CreateSession(ctx context.Context, userID string) (*Session, error) {
    session := &Session{
        ID:        generateSessionID(),
        UserID:    userID,
        CreatedAt: time.Now(),
        ExpiresAt: time.Now().Add(24 * time.Hour),
    }

    data, _ := json.Marshal(session)
    err := redis.Set(ctx, "session:"+session.ID, data, 24*time.Hour).Err()
    return session, err
}

// Validate session
func GetSession(ctx context.Context, sessionID string) (*Session, error) {
    data, err := redis.Get(ctx, "session:"+sessionID).Bytes()
    if err == redis.Nil {
        return nil, ErrSessionNotFound
    }

    var session Session
    json.Unmarshal(data, &session)

    if time.Now().After(session.ExpiresAt) {
        redis.Del(ctx, "session:"+sessionID)
        return nil, ErrSessionExpired
    }

    return &session, nil
}

// Revoke session
func RevokeSession(ctx context.Context, sessionID string) error {
    return redis.Del(ctx, "session:"+sessionID).Err()
}

// Revoke all sessions for user
func RevokeAllUserSessions(ctx context.Context, userID string) error {
    keys, _ := redis.Keys(ctx, "session:*").Result()
    for _, key := range keys {
        data, _ := redis.Get(ctx, key).Bytes()
        var session Session
        json.Unmarshal(data, &session)
        if session.UserID == userID {
            redis.Del(ctx, key)
        }
    }
    return nil
}
```

## Workflow State (Temporal)

Temporal handles workflow state internally. Key considerations:

### State Persistence

```go
// Workflow state is automatically persisted after each event
func OrderWorkflow(ctx workflow.Context, order Order) error {
    // State checkpoint 1
    err := workflow.ExecuteActivity(ctx, ValidateOrder, order).Get(ctx, nil)

    // State checkpoint 2
    err = workflow.ExecuteActivity(ctx, ChargePayment, order).Get(ctx, nil)

    // State checkpoint 3
    err = workflow.ExecuteActivity(ctx, FulfillOrder, order).Get(ctx, nil)

    return nil
}
```

### Continue-As-New

Prevent unbounded history:

```go
func LongRunningWorkflow(ctx workflow.Context, state State) error {
    for {
        // Process events...

        // Check history size
        info := workflow.GetInfo(ctx)
        if info.GetCurrentHistoryLength() > 10000 {
            return workflow.NewContinueAsNewError(ctx, LongRunningWorkflow, state)
        }
    }
}
```

## Cache State (Redis)

### Cache Invalidation

```go
// Cache with TTL
func GetOrganization(ctx context.Context, id string) (*Organization, error) {
    // Try cache
    cacheKey := "org:" + id
    data, err := redis.Get(ctx, cacheKey).Bytes()
    if err == nil {
        var org Organization
        json.Unmarshal(data, &org)
        return &org, nil
    }

    // Cache miss - fetch from DB
    org, err := db.GetOrganization(ctx, id)
    if err != nil {
        return nil, err
    }

    // Populate cache
    data, _ = json.Marshal(org)
    redis.Set(ctx, cacheKey, data, 5*time.Minute)

    return org, nil
}

// Invalidate on update
func UpdateOrganization(ctx context.Context, org *Organization) error {
    err := db.UpdateOrganization(ctx, org)
    if err != nil {
        return err
    }

    // Invalidate cache
    redis.Del(ctx, "org:"+org.ID)

    return nil
}
```

### Cache Stampede Prevention

```go
// Use singleflight to prevent multiple concurrent fetches
var group singleflight.Group

func GetOrganizationSafe(ctx context.Context, id string) (*Organization, error) {
    v, err, _ := group.Do(id, func() (interface{}, error) {
        return GetOrganization(ctx, id)
    })

    if err != nil {
        return nil, err
    }
    return v.(*Organization), nil
}
```

## State Recovery

### Terraform State Recovery

```bash
# If state is corrupted, restore from S3 version
aws s3api list-object-versions \
  --bucket temporal-cloud-terraform-state \
  --prefix environments/production/us-east-1/terraform.tfstate

# Restore specific version
aws s3api get-object \
  --bucket temporal-cloud-terraform-state \
  --key environments/production/us-east-1/terraform.tfstate \
  --version-id "abc123" \
  terraform.tfstate.backup
```

### Database State Recovery

See `dr.md` for point-in-time recovery procedures.

### Redis State Recovery

Redis state is ephemeral. On failure:

1. Sessions require re-login
2. Cache repopulates automatically
3. Rate limit counters reset (acceptable)
