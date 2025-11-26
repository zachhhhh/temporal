# Certificate Management

## mTLS Authentication

Temporal Cloud uses mutual TLS (mTLS) for namespace authentication. Clients must present a valid certificate signed by a CA registered with the namespace.

## Certificate Requirements

### CA Certificates

- Format: X.509 v3
- Encoding: PEM
- Key size: RSA 2048+ or ECDSA P-256+
- Max per namespace: 16 certificates or 32KB total

### End-Entity Certificates

- Must be signed by registered CA
- Validity: Recommended max 1 year
- Must include Extended Key Usage: Client Authentication

## Certificate Operations

| Operation               | Console | tcld | API | Terraform |
| ----------------------- | ------- | ---- | --- | --------- |
| Add CA certificate      | ✅      | ✅   | ✅  | ✅        |
| Remove CA certificate   | ✅      | ✅   | ✅  | ✅        |
| List certificates       | ✅      | ✅   | ✅  | ✅        |
| Set certificate filters | ✅      | ✅   | ✅  | ✅        |

## Certificate Filters

Fine-grained access control based on certificate attributes.

### Filter Types

| Filter     | Description          | Example              |
| ---------- | -------------------- | -------------------- |
| Subject CN | Common Name match    | `*.example.com`      |
| Subject OU | Organizational Unit  | `engineering`        |
| SAN DNS    | DNS Subject Alt Name | `worker.example.com` |

### Filter Configuration

```json
{
  "filters": [
    {
      "type": "subject_cn",
      "value": "*.prod.example.com"
    },
    {
      "type": "subject_ou",
      "value": "production"
    }
  ]
}
```

## Generating Certificates

### Using OpenSSL

```bash
# Generate CA key and certificate
openssl genrsa -out ca.key 4096
openssl req -new -x509 -days 365 -key ca.key -out ca.pem \
  -subj "/CN=Temporal CA/O=MyOrg"

# Generate client key and CSR
openssl genrsa -out client.key 4096
openssl req -new -key client.key -out client.csr \
  -subj "/CN=temporal-worker/O=MyOrg"

# Sign client certificate
openssl x509 -req -days 365 -in client.csr \
  -CA ca.pem -CAkey ca.key -CAcreateserial \
  -out client.pem
```

### Using tcld

```bash
tcld generate-certificates \
  --namespace my-namespace \
  --ca-cert ca.pem \
  --output-dir ./certs
```

## Certificate Rotation

### Process

1. Generate new CA certificate
2. Add new CA to namespace (both old and new active)
3. Update all workers with new client certificates
4. Remove old CA certificate
5. Monitor for authentication failures

### Zero-Downtime Rotation

```bash
# Step 1: Add new CA (keep old)
tcld namespace certificates add \
  --namespace my-namespace \
  --ca-certificate new-ca.pem

# Step 2: Update workers (deploy with new certs)

# Step 3: Remove old CA
tcld namespace certificates remove \
  --namespace my-namespace \
  --ca-certificate-fingerprint <old-fingerprint>
```

## Expiration Monitoring

### Check Expiration

```bash
# Check certificate expiration
openssl x509 -enddate -noout -in ca.pem

# Output: notAfter=Jan  1 00:00:00 2026 GMT
```

### Alerts

- 30 days before expiry: Warning notification
- 7 days before expiry: Critical notification
- On expiry: Certificate rejected

## Troubleshooting

### Certificate Rejected

1. Verify CA is registered with namespace
2. Check certificate not expired
3. Verify certificate chain is complete
4. Check certificate filters match

### Connection Failed

```bash
# Test connection with certificate
grpcurl -cert client.pem -key client.key -cacert ca.pem \
  my-namespace.tmprl.cloud:443 \
  temporal.api.workflowservice.v1.WorkflowService/GetSystemInfo
```
