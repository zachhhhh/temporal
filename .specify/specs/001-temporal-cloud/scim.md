# SCIM Provisioning

## Overview

System for Cross-domain Identity Management (SCIM 2.0) enables automatic user and group provisioning from your Identity Provider.

## Requirements

- SAML SSO must be configured first
- Plan: Enterprise/Mission Critical, or Business + $500/mo add-on

## Supported IdPs

- Okta
- Microsoft Entra ID (Azure AD)
- OneLogin
- Any SCIM 2.0 compliant IdP

## Endpoints

| Endpoint                         | Methods                  | Description        |
| -------------------------------- | ------------------------ | ------------------ |
| `/scim/v2/Users`                 | GET, POST, PATCH, DELETE | User management    |
| `/scim/v2/Groups`                | GET, POST, PATCH, DELETE | Group management   |
| `/scim/v2/ServiceProviderConfig` | GET                      | SCIM configuration |
| `/scim/v2/Schemas`               | GET                      | Schema definitions |

## User Attributes

| SCIM Attribute  | Temporal Attribute | Required |
| --------------- | ------------------ | -------- |
| userName        | email              | Yes      |
| displayName     | name               | No       |
| active          | enabled            | Yes      |
| emails[primary] | email              | Yes      |

## Group Attributes

| SCIM Attribute | Temporal Attribute |
| -------------- | ------------------ |
| displayName    | Group name         |
| members        | Group members      |

## Configuration

### Step 1: Enable SCIM

```bash
tcld account scim enable
```

### Step 2: Generate SCIM Token

```bash
tcld account scim token create
# Output: SCIM bearer token (shown only once)
```

### Step 3: Configure IdP

#### Okta

1. Go to Applications → Your SAML App → Provisioning
2. Enable SCIM provisioning
3. Set Base URL: `https://api.temporal.io/scim/v2`
4. Set API Token: `<scim-token>`
5. Enable: Create Users, Update User Attributes, Deactivate Users

#### Azure AD

1. Go to Enterprise App → Provisioning
2. Set Provisioning Mode: Automatic
3. Set Tenant URL: `https://api.temporal.io/scim/v2`
4. Set Secret Token: `<scim-token>`
5. Test connection and save

## Provisioning Behavior

### User Created in IdP

1. SCIM POST to `/Users`
2. User created in Temporal Cloud
3. Default role assigned (or from group mapping)

### User Updated in IdP

1. SCIM PATCH to `/Users/{id}`
2. User attributes updated
3. Role unchanged (unless group changed)

### User Deactivated in IdP

1. SCIM PATCH with `active: false`
2. User disabled in Temporal Cloud
3. Active sessions terminated

### User Deleted in IdP

1. SCIM DELETE to `/Users/{id}`
2. User removed from Temporal Cloud
3. API keys revoked

## Group Sync

### Group Created

1. SCIM POST to `/Groups`
2. User group created in Temporal Cloud

### Members Added

1. SCIM PATCH to `/Groups/{id}`
2. Users added to group
3. Namespace permissions applied

### Members Removed

1. SCIM PATCH to `/Groups/{id}`
2. Users removed from group
3. Namespace permissions revoked

## Namespace Permissions via Groups

Map IdP groups to namespace permissions:

```bash
tcld user-group namespace-access set \
  --group-id grp-123 \
  --namespace my-namespace \
  --permission write
```

## Pricing

| Plan             | SCIM           |
| ---------------- | -------------- |
| Essential        | Not available  |
| Business         | $500/mo add-on |
| Enterprise       | Included       |
| Mission Critical | Included       |

## Troubleshooting

### Provisioning Failed

1. Check SCIM token is valid
2. Verify Base URL is correct
3. Check IdP logs for errors
4. Verify user limit not reached

### User Not Synced

1. Check user is assigned to app in IdP
2. Verify required attributes are mapped
3. Check provisioning logs in IdP

### Group Permissions Not Applied

1. Verify group exists in Temporal Cloud
2. Check namespace permissions are configured
3. Verify user is member of group
