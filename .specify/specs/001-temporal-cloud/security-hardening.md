# Security Hardening

## Defense in Depth

```
┌─────────────────────────────────────────────────────────────────┐
│                        Layer 1: Edge                             │
│  WAF, DDoS Protection, CDN, Rate Limiting                       │
├─────────────────────────────────────────────────────────────────┤
│                        Layer 2: Network                          │
│  VPC, Security Groups, NACLs, Private Subnets                   │
├─────────────────────────────────────────────────────────────────┤
│                        Layer 3: Application                      │
│  Authentication, Authorization, Input Validation                 │
├─────────────────────────────────────────────────────────────────┤
│                        Layer 4: Data                             │
│  Encryption at Rest, Encryption in Transit, Key Management      │
├─────────────────────────────────────────────────────────────────┤
│                        Layer 5: Monitoring                       │
│  Logging, Alerting, Anomaly Detection                           │
└─────────────────────────────────────────────────────────────────┘
```

## Edge Security

### AWS WAF Rules

```hcl
resource "aws_wafv2_web_acl" "main" {
  name  = "temporal-cloud-waf"
  scope = "REGIONAL"

  default_action {
    allow {}
  }

  # Rate limiting
  rule {
    name     = "RateLimitRule"
    priority = 1

    action {
      block {}
    }

    statement {
      rate_based_statement {
        limit              = 2000
        aggregate_key_type = "IP"
      }
    }

    visibility_config {
      sampled_requests_enabled   = true
      cloudwatch_metrics_enabled = true
      metric_name                = "RateLimitRule"
    }
  }

  # SQL Injection
  rule {
    name     = "SQLInjectionRule"
    priority = 2

    override_action {
      none {}
    }

    statement {
      managed_rule_group_statement {
        name        = "AWSManagedRulesSQLiRuleSet"
        vendor_name = "AWS"
      }
    }
  }

  # Known bad inputs
  rule {
    name     = "KnownBadInputsRule"
    priority = 3

    override_action {
      none {}
    }

    statement {
      managed_rule_group_statement {
        name        = "AWSManagedRulesKnownBadInputsRuleSet"
        vendor_name = "AWS"
      }
    }
  }

  # Bot control
  rule {
    name     = "BotControlRule"
    priority = 4

    override_action {
      none {}
    }

    statement {
      managed_rule_group_statement {
        name        = "AWSManagedRulesBotControlRuleSet"
        vendor_name = "AWS"
      }
    }
  }
}
```

### DDoS Protection

```hcl
# AWS Shield Advanced
resource "aws_shield_protection" "alb" {
  name         = "temporal-cloud-alb"
  resource_arn = aws_lb.api.arn
}

# Automatic DDoS response
resource "aws_shield_protection_group" "main" {
  protection_group_id = "temporal-cloud"
  aggregation         = "MAX"
  pattern             = "BY_RESOURCE_TYPE"
  resource_type       = "APPLICATION_LOAD_BALANCER"
}
```

## Network Hardening

### VPC Configuration

```hcl
# No default VPC
# Custom VPC with private subnets only for workloads

resource "aws_vpc" "main" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name = "temporal-cloud-prod"
  }
}

# Flow logs for network monitoring
resource "aws_flow_log" "main" {
  vpc_id          = aws_vpc.main.id
  traffic_type    = "ALL"
  iam_role_arn    = aws_iam_role.flow_log.arn
  log_destination = aws_cloudwatch_log_group.flow_log.arn
}
```

### Security Groups (Least Privilege)

```hcl
# ALB - Only HTTPS from internet
resource "aws_security_group" "alb" {
  name   = "alb-sg"
  vpc_id = aws_vpc.main.id

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port       = 8080
    to_port         = 8080
    protocol        = "tcp"
    security_groups = [aws_security_group.app.id]
  }
}

# App - Only from ALB
resource "aws_security_group" "app" {
  name   = "app-sg"
  vpc_id = aws_vpc.main.id

  ingress {
    from_port       = 8080
    to_port         = 8080
    protocol        = "tcp"
    security_groups = [aws_security_group.alb.id]
  }

  egress {
    from_port       = 5432
    to_port         = 5432
    protocol        = "tcp"
    security_groups = [aws_security_group.db.id]
  }
}

# DB - Only from App
resource "aws_security_group" "db" {
  name   = "db-sg"
  vpc_id = aws_vpc.main.id

  ingress {
    from_port       = 5432
    to_port         = 5432
    protocol        = "tcp"
    security_groups = [aws_security_group.app.id]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = []  # No egress
  }
}
```

## Application Hardening

### Input Validation

```go
// Validate all input
func (s *Server) CreateNamespace(ctx context.Context, req *CreateNamespaceRequest) error {
    // Length limits
    if len(req.Name) < 2 || len(req.Name) > 63 {
        return status.Error(codes.InvalidArgument, "name must be 2-63 characters")
    }

    // Character whitelist
    if !regexp.MustCompile(`^[a-z][a-z0-9-]*[a-z0-9]$`).MatchString(req.Name) {
        return status.Error(codes.InvalidArgument, "name must be lowercase alphanumeric")
    }

    // Sanitize for logging (prevent log injection)
    sanitizedName := sanitize(req.Name)
    log.Info("Creating namespace", "name", sanitizedName)

    // ... create namespace
}
```

### SQL Injection Prevention

```go
// ALWAYS use parameterized queries
func GetOrganization(ctx context.Context, id string) (*Organization, error) {
    // ✅ Correct - parameterized
    row := db.QueryRowContext(ctx,
        "SELECT id, name FROM organizations WHERE id = $1",
        id)

    // ❌ NEVER do this
    // row := db.QueryRowContext(ctx,
    //     "SELECT id, name FROM organizations WHERE id = '" + id + "'")

    var org Organization
    err := row.Scan(&org.ID, &org.Name)
    return &org, err
}
```

### XSS Prevention

```typescript
// React automatically escapes output
// But be careful with dangerouslySetInnerHTML

// ✅ Safe
return <div>{userInput}</div>;

// ❌ Dangerous - avoid unless absolutely necessary
return <div dangerouslySetInnerHTML={{ __html: userInput }} />;

// If you must render HTML, sanitize first
import DOMPurify from "dompurify";
return (
  <div dangerouslySetInnerHTML={{ __html: DOMPurify.sanitize(userInput) }} />
);
```

### CSRF Protection

```go
// Use CSRF tokens for state-changing operations
func CSRFMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "GET" && r.Method != "HEAD" && r.Method != "OPTIONS" {
            token := r.Header.Get("X-CSRF-Token")
            if !validateCSRFToken(r.Context(), token) {
                http.Error(w, "Invalid CSRF token", http.StatusForbidden)
                return
            }
        }
        next.ServeHTTP(w, r)
    })
}
```

## Data Hardening

### Encryption at Rest

```hcl
# RDS encryption
resource "aws_db_instance" "main" {
  storage_encrypted = true
  kms_key_id        = aws_kms_key.db.arn
}

# S3 encryption
resource "aws_s3_bucket_server_side_encryption_configuration" "main" {
  bucket = aws_s3_bucket.data.id

  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm     = "aws:kms"
      kms_master_key_id = aws_kms_key.s3.arn
    }
  }
}

# EBS encryption
resource "aws_ebs_encryption_by_default" "main" {
  enabled = true
}
```

### Encryption in Transit

```yaml
# Force TLS 1.3
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    alb.ingress.kubernetes.io/ssl-policy: ELBSecurityPolicy-TLS13-1-2-2021-06
```

### Secrets in Memory

```go
// Clear sensitive data from memory when done
func ProcessPayment(card *CreditCard) error {
    defer func() {
        // Zero out sensitive data
        for i := range card.Number {
            card.Number[i] = 0
        }
        for i := range card.CVV {
            card.CVV[i] = 0
        }
    }()

    // Process payment...
}
```

## Container Hardening

### Dockerfile Best Practices

```dockerfile
# Use minimal base image
FROM gcr.io/distroless/static:nonroot

# Don't run as root
USER nonroot:nonroot

# No shell, minimal attack surface
# Read-only filesystem where possible
```

### Pod Security

```yaml
apiVersion: v1
kind: Pod
spec:
  securityContext:
    runAsNonRoot: true
    runAsUser: 65534
    fsGroup: 65534
    seccompProfile:
      type: RuntimeDefault
  containers:
    - name: api
      securityContext:
        allowPrivilegeEscalation: false
        readOnlyRootFilesystem: true
        capabilities:
          drop:
            - ALL
```

## Hardening Checklist

### Infrastructure

- [ ] No public S3 buckets
- [ ] No default VPC
- [ ] All storage encrypted
- [ ] Flow logs enabled
- [ ] WAF configured
- [ ] Shield enabled

### Application

- [ ] Input validation on all endpoints
- [ ] Parameterized queries only
- [ ] CSRF protection
- [ ] Rate limiting
- [ ] Security headers set

### Access

- [ ] MFA required
- [ ] Least privilege IAM
- [ ] No long-lived credentials
- [ ] Regular access reviews

### Monitoring

- [ ] All access logged
- [ ] Anomaly detection enabled
- [ ] Alert on suspicious activity
- [ ] Regular security scans
