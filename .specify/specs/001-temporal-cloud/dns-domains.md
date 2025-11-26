# DNS & Domain Management

## Domain Structure

### Primary Domains

| Domain              | Purpose                      | Registrar         | DNS Provider |
| ------------------- | ---------------------------- | ----------------- | ------------ |
| `temporal-cloud.io` | Marketing, docs              | Namecheap/Route53 | Route53      |
| `tmprl.cloud`       | Customer namespace endpoints | Route53           | Route53      |

### Subdomain Hierarchy

```
temporal-cloud.io
├── www.temporal-cloud.io          → Marketing site (CloudFront)
├── docs.temporal-cloud.io         → Documentation (CloudFront)
├── console.temporal-cloud.io      → Web Console (ALB)
├── api.temporal-cloud.io          → REST/gRPC API (ALB)
├── status.temporal-cloud.io       → Status page (Atlassian)
├── support.temporal-cloud.io      → Support portal (Zendesk)
└── billing.temporal-cloud.io      → Billing portal (Stripe)

tmprl.cloud
├── <namespace>.<account>.<region>.tmprl.cloud  → Namespace endpoints
│   Example: prod.acme.us-east-1.tmprl.cloud
├── metrics.<region>.tmprl.cloud                → Prometheus endpoint
└── grpc.<region>.tmprl.cloud                   → gRPC endpoint
```

## Route53 Configuration

### Hosted Zones

```hcl
# terraform/modules/dns/main.tf

resource "aws_route53_zone" "primary" {
  name = "temporal-cloud.io"

  tags = {
    Environment = "production"
    ManagedBy   = "terraform"
  }
}

resource "aws_route53_zone" "customer" {
  name = "tmprl.cloud"

  tags = {
    Environment = "production"
    ManagedBy   = "terraform"
  }
}
```

### DNS Records

```hcl
# Console (CloudFront)
resource "aws_route53_record" "console" {
  zone_id = aws_route53_zone.primary.zone_id
  name    = "console.temporal-cloud.io"
  type    = "A"

  alias {
    name                   = aws_cloudfront_distribution.console.domain_name
    zone_id                = aws_cloudfront_distribution.console.hosted_zone_id
    evaluate_target_health = false
  }
}

# API (ALB with latency routing)
resource "aws_route53_record" "api_us" {
  zone_id        = aws_route53_zone.primary.zone_id
  name           = "api.temporal-cloud.io"
  type           = "A"
  set_identifier = "us-east-1"

  alias {
    name                   = aws_lb.api_us.dns_name
    zone_id                = aws_lb.api_us.zone_id
    evaluate_target_health = true
  }

  latency_routing_policy {
    region = "us-east-1"
  }
}

resource "aws_route53_record" "api_eu" {
  zone_id        = aws_route53_zone.primary.zone_id
  name           = "api.temporal-cloud.io"
  type           = "A"
  set_identifier = "eu-west-1"

  alias {
    name                   = aws_lb.api_eu.dns_name
    zone_id                = aws_lb.api_eu.zone_id
    evaluate_target_health = true
  }

  latency_routing_policy {
    region = "eu-west-1"
  }
}

# Wildcard for customer namespaces
resource "aws_route53_record" "namespace_wildcard" {
  zone_id = aws_route53_zone.customer.zone_id
  name    = "*.us-east-1.tmprl.cloud"
  type    = "A"

  alias {
    name                   = aws_lb.namespace_us.dns_name
    zone_id                = aws_lb.namespace_us.zone_id
    evaluate_target_health = true
  }
}
```

### Health Checks

```hcl
resource "aws_route53_health_check" "api_us" {
  fqdn              = "api-us-east-1.temporal-cloud.io"
  port              = 443
  type              = "HTTPS"
  resource_path     = "/health"
  failure_threshold = 3
  request_interval  = 10

  tags = {
    Name = "api-us-east-1-health"
  }
}

# Failover routing
resource "aws_route53_record" "api_failover_primary" {
  zone_id         = aws_route53_zone.primary.zone_id
  name            = "api.temporal-cloud.io"
  type            = "A"
  set_identifier  = "primary"
  health_check_id = aws_route53_health_check.api_us.id

  failover_routing_policy {
    type = "PRIMARY"
  }

  alias {
    name    = aws_lb.api_us.dns_name
    zone_id = aws_lb.api_us.zone_id
  }
}
```

## Dynamic DNS for Namespaces

When a namespace is created, DNS is provisioned automatically:

```go
func (p *Provisioner) CreateNamespaceDNS(ctx context.Context, ns *Namespace) error {
    // Construct FQDN
    fqdn := fmt.Sprintf("%s.%s.%s.tmprl.cloud",
        ns.Name,
        ns.AccountSlug,
        ns.Region)

    // Create Route53 record
    input := &route53.ChangeResourceRecordSetsInput{
        HostedZoneId: aws.String(p.customerZoneID),
        ChangeBatch: &route53.ChangeBatch{
            Changes: []*route53.Change{
                {
                    Action: aws.String("CREATE"),
                    ResourceRecordSet: &route53.ResourceRecordSet{
                        Name: aws.String(fqdn),
                        Type: aws.String("CNAME"),
                        TTL:  aws.Int64(300),
                        ResourceRecords: []*route53.ResourceRecord{
                            {Value: aws.String(p.regionalEndpoints[ns.Region])},
                        },
                    },
                },
            },
        },
    }

    _, err := p.route53.ChangeResourceRecordSets(ctx, input)
    return err
}
```

## SSL/TLS Certificates

### AWS Certificate Manager (ACM)

```hcl
# Wildcard cert for primary domain
resource "aws_acm_certificate" "primary" {
  domain_name       = "temporal-cloud.io"
  validation_method = "DNS"

  subject_alternative_names = [
    "*.temporal-cloud.io"
  ]

  lifecycle {
    create_before_destroy = true
  }
}

# Wildcard cert for customer domain
resource "aws_acm_certificate" "customer" {
  domain_name       = "tmprl.cloud"
  validation_method = "DNS"

  subject_alternative_names = [
    "*.tmprl.cloud",
    "*.us-east-1.tmprl.cloud",
    "*.eu-west-1.tmprl.cloud",
    "*.ap-south-1.tmprl.cloud"
  ]
}

# DNS validation
resource "aws_route53_record" "cert_validation" {
  for_each = {
    for dvo in aws_acm_certificate.primary.domain_validation_options : dvo.domain_name => {
      name   = dvo.resource_record_name
      record = dvo.resource_record_value
      type   = dvo.resource_record_type
    }
  }

  zone_id = aws_route53_zone.primary.zone_id
  name    = each.value.name
  type    = each.value.type
  records = [each.value.record]
  ttl     = 60
}
```

### Certificate Renewal

ACM certificates auto-renew. For cert-manager:

```yaml
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: letsencrypt-prod
spec:
  acme:
    server: https://acme-v02.api.letsencrypt.org/directory
    email: platform@temporal.io
    privateKeySecretRef:
      name: letsencrypt-prod-key
    solvers:
      - dns01:
          route53:
            region: us-east-1
            hostedZoneID: Z1234567890
```

## Domain Security

### DNSSEC

```hcl
resource "aws_route53_key_signing_key" "primary" {
  hosted_zone_id             = aws_route53_zone.primary.zone_id
  key_management_service_arn = aws_kms_key.dnssec.arn
  name                       = "temporal-cloud-ksk"
}

resource "aws_route53_hosted_zone_dnssec" "primary" {
  hosted_zone_id = aws_route53_zone.primary.zone_id

  depends_on = [aws_route53_key_signing_key.primary]
}
```

### CAA Records

```hcl
resource "aws_route53_record" "caa" {
  zone_id = aws_route53_zone.primary.zone_id
  name    = "temporal-cloud.io"
  type    = "CAA"
  ttl     = 3600

  records = [
    "0 issue \"amazon.com\"",
    "0 issue \"letsencrypt.org\"",
    "0 iodef \"mailto:security@temporal.io\""
  ]
}
```

### SPF, DKIM, DMARC (for email)

```hcl
# SPF
resource "aws_route53_record" "spf" {
  zone_id = aws_route53_zone.primary.zone_id
  name    = "temporal-cloud.io"
  type    = "TXT"
  ttl     = 3600
  records = ["v=spf1 include:amazonses.com include:sendgrid.net ~all"]
}

# DKIM (SendGrid)
resource "aws_route53_record" "dkim" {
  zone_id = aws_route53_zone.primary.zone_id
  name    = "s1._domainkey.temporal-cloud.io"
  type    = "CNAME"
  ttl     = 3600
  records = ["s1.domainkey.u12345.wl.sendgrid.net"]
}

# DMARC
resource "aws_route53_record" "dmarc" {
  zone_id = aws_route53_zone.primary.zone_id
  name    = "_dmarc.temporal-cloud.io"
  type    = "TXT"
  ttl     = 3600
  records = ["v=DMARC1; p=quarantine; rua=mailto:dmarc@temporal.io"]
}
```

## DNS Monitoring

### CloudWatch Alarms

```hcl
resource "aws_cloudwatch_metric_alarm" "health_check_failed" {
  alarm_name          = "dns-health-check-failed"
  comparison_operator = "LessThanThreshold"
  evaluation_periods  = 2
  metric_name         = "HealthCheckStatus"
  namespace           = "AWS/Route53"
  period              = 60
  statistic           = "Minimum"
  threshold           = 1

  dimensions = {
    HealthCheckId = aws_route53_health_check.api_us.id
  }

  alarm_actions = [aws_sns_topic.alerts.arn]
}
```

## Runbook: DNS Propagation Issues

1. Check Route53 record exists: `aws route53 list-resource-record-sets --hosted-zone-id Z123`
2. Check TTL hasn't expired: `dig +trace api.temporal-cloud.io`
3. Verify health check status in Route53 console
4. Check if change is still PENDING: `aws route53 get-change --id /change/C123`
