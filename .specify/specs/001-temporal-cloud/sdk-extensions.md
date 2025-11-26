# SDK Extensions

## Overview

While standard Temporal SDKs work with Temporal Cloud, we provide wrapper libraries to improve Developer Experience (DX), specifically for certificate rotation and cloud-specific metrics.

## Go SDK Extension (`go.temporal.io/sdk/contrib/cloud`)

### 1. Automatic Certificate Rotation

Standard SDK requires restarting the client when certificates change. This extension watches the file system and hot-reloads TLS config.

```go
options := client.Options{
    HostPort: "my-ns.tmprl.cloud:443",
    Namespace: "my-ns",
    ConnectionOptions: client.ConnectionOptions{
        TLS: cloud.NewRotatingTLSConfig(cloud.TLSParams{
            CertPath: "/etc/certs/client.pem",
            KeyPath:  "/etc/certs/client.key",
            CheckInterval: 1 * time.Minute,
        }),
    },
}
```

### 2. Cloud Metrics Tagger

Automatically adds `namespace` and `region` tags to all metrics and converts them to OpenTelemetry format preferred by our observability stack.

```go
options := client.Options{
    MetricsHandler: cloud.NewMetricsHandler(cloud.MetricsParams{
        Prometheus: true,
    }),
}
```

### 3. Connection Tuner

Sets gRPC keepalive parameters optimal for Temporal Cloud load balancers to prevent connection resets.

```go
// Automatically sets:
// - KeepAliveTime: 30s
// - KeepAliveTimeout: 10s
// - PermitWithoutStream: true
client, err := cloud.Dial(ctx, options)
```

## Java SDK Extension

### 1. KeyStore Reloader

Watches the JKS/PKCS12 file and reloads the `SslContext`.

```java
Scope scope = new CloudScopeBuilder()
    .setCertPath(Paths.get("client.pem"))
    .setKeyPath(Paths.get("client.key"))
    .enableRotation(Duration.ofMinutes(1))
    .build();

WorkflowServiceStubsOptions options = WorkflowServiceStubsOptions.newBuilder()
    .setSslContext(scope.getSslContext())
    .build();
```

## TypeScript SDK Extension

### 1. MTLS Reloader

Uses `fs.watch` to update the connection options.

```typescript
const connection = await CloudConnection.connect({
  address: "my-ns.tmprl.cloud",
  tls: {
    certPath: "./client.pem",
    keyPath: "./client.key",
    autoRotate: true,
  },
});
```

## Required Changes to Core SDKs

We aim to keep core SDKs generic, but we may need to open up `ClientOptions` to allow dynamic swapping of `Credential` providers if not already supported.
