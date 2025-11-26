# Content Delivery Network (CDN)

## Strategy

Temporal Cloud is largely API-based (dynamic), but static assets and global routing benefit from CDN.

## Static Assets

### Cloud Console

- **HTML/JS/CSS**: Cached at edge (Cloudflare/CloudFront).
- **TTL**: 1 year (immutable hashes).
- **Invalidation**: On deployment.

### Documentation

- Cached globally.
- Stale-while-revalidate strategy.

## Dynamic Routing (Global Accelerator)

Accelerate gRPC/API traffic using AWS Global Accelerator or Cloudflare Spectrum.

1. **Anycast IP**: User connects to nearest edge POP.
2. **Backbone**: Traffic traverses provider's private fiber, not public internet.
3. **Origin**: Handoff to nearest regional ALB.

**Benefit**: ~30% latency reduction for cross-continent API calls.

## Security at Edge

- **WAF**: Block threats at edge before they hit origin.
- **DDoS**: Absorb volumetric attacks.
- **TLS Termination**: Offload crypto handshake to edge.

## Configuration

### AWS CloudFront

```hcl
resource "aws_cloudfront_distribution" "console" {
  origin {
    domain_name = aws_s3_bucket.frontend.bucket_regional_domain_name
    origin_id   = "s3-console"
  }

  # SPA routing
  custom_error_response {
    error_code         = 404
    response_code      = 200
    response_page_path = "/index.html"
  }
}
```

### AWS Global Accelerator

```hcl
resource "aws_globalaccelerator_accelerator" "api" {
  name            = "temporal-cloud-api"
  ip_address_type = "IPV4"
  enabled         = true
}

resource "aws_globalaccelerator_listener" "api" {
  accelerator_arn = aws_globalaccelerator_accelerator.api.id
  protocol        = "TCP"

  port_range {
    from_port = 443
    to_port   = 443
  }
}
```
