# Cloud Console Design

## Architecture

- **Framework**: Next.js 14 (App Router)
- **Styling**: Tailwind CSS + shadcn/ui
- **State**: TanStack Query + Zustand
- **Auth**: JWT in HttpOnly cookie
- **API**: gRPC-Web (Connect)

## Site Map

```
/login
/sso/callback
/dashboard (Redirects to first account)
/accounts/{account_id}/
├── overview               # Usage, health
├── namespaces/
│   ├── [list]
│   ├── create
│   └── {namespace_id}/
│       ├── overview       # Metrics
│       ├── workflows/     # Web UI
│       ├── settings       # Retention, CA certs
│       └── limits         # Quotas
├── members/               # Users & Roles
├── service-accounts/
├── billing/
│   ├── overview           # Current usage & estimated cost
│   ├── invoices/          # Past invoices
│   └── settings           # Payment methods, address
├── settings/
    ├── general
    ├── sso                # SAML config
    ├── audit-logs         # Export config
    └── api-keys
```

## Key Screens

### 1. Namespace List

**Columns**: Status, Name, Region, Retention, Active Workflows, Actions/sec (24h avg).
**Actions**: Create Namespace, Search.

### 2. Namespace Detail > Overview

**Charts** (Recharts):

- Success Rate (Line)
- Latency (Line)
- Actions/sec (Area)
  **Info**: Certificates (expiry warning), Region, Grpc Endpoint (copy button).

### 3. Create Namespace Wizard

**Step 1: Basics**: Name, Region selection (with map/latency hints).
**Step 2: Config**: Retention period slider (1-90 days).
**Step 3: Security**: Upload CA certificate (drag & drop).
**Step 4: Review**: Summary & "Create" button.

### 4. Billing Overview

**Header**: Current Month-to-Date Cost.
**Breakdown**:

- Plan Base Fee
- Actions Overage (Progress bar vs included amount)
- Storage (Active vs Retained)
  **Usage Chart**: Daily bar chart of actions.

### 5. SSO Configuration

**Status**: Enabled/Disabled toggle.
**Config**:

- IdP Metadata URL input.
- Manual file upload.
  **Attribute Mapping**:
- Role mapping table (IdP Group -> Cloud Role).

## UX Patterns

- **Loading**: Skeleton loaders (shadcn/ui `Skeleton`).
- **Errors**: Toast notifications (Sonner) + inline error boundaries.
- **Copying**: Click-to-copy for IDs, keys, endpoints.
- **Time**: All times in UTC by default, toggle for Local.
- **Dates**: Relative ("2 hours ago") with tooltip for absolute.

## Theming

- **Mode**: Dark/Light mode toggle (default Dark for dev tools).
- **Brand Colors**:
  - Primary: Temporal Black/White.
  - Accents: Blue for links, Green for success, Red for errors.
