# Plugin Management

## Overview

Temporal allows plugins (interceptors, custom code). In Cloud, we must manage these strictly for security and stability.

## Supported Plugin Types

1. **Interceptors (Client-side)**

   - Run in customer's SDK.
   - **Management**: Distribute via SDK extensions (see `sdk-extensions.md`).
   - **Safety**: Safe (customer's risk).

2. **Remote Data Converter (Codec Server)**

   - Decrypts payloads for display in Console.
   - **Architecture**: Browser calls customer's HTTP endpoint directly.
   - **Management**: Configured per-namespace in Console settings.
   - **Safety**: Cloud never sees unencrypted data.

3. **Search Attributes (Server-side)**
   - Custom index fields.
   - **Limit**: 100 per namespace.
   - **Types**: Strict validation (Keyword, Int, etc.).

## Prohibited Plugins

- **Custom Workflow Logic (Server-side)**: No user code runs on Temporal servers.
- **Custom Interceptors (Server-side)**: Only standard Cloud interceptors allowed (Metering, Auth).

## Plugin Update Strategy

### SDK Plugins

- Versioned via standard package managers (npm, maven, go mod).
- Deprecation warnings in SDK logs.

### Codec Server

- Customer responsible for maintaining availability.
- Console handles connection errors gracefully ("Unable to decode payload").

## Internal Plugins (Cloud Platform)

We use internal plugins for extensibility:

```go
// Platform plugin interface
type CloudPlugin interface {
    OnNamespaceCreated(ctx context.Context, ns *Namespace) error
    OnNamespaceDeleted(ctx context.Context, ns *Namespace) error
}

// Implementation example: Datadog Metrics
type DatadogPlugin struct {}

func (p *DatadogPlugin) OnNamespaceCreated(ctx context.Context, ns *Namespace) error {
    // Register new tag in Datadog
    return nil
}
```

Managed via dependency injection at startup.
