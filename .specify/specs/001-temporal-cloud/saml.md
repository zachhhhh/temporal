# SAML SSO

## Overview

SAML 2.0 Single Sign-On allows users to authenticate via their organization's Identity Provider (IdP).

## Supported IdPs

- Microsoft Entra ID (Azure AD)
- Okta
- OneLogin
- Any SAML 2.0 compliant IdP

## Configuration

### Step 1: Create SAML Application in IdP

#### Okta

1. Create new SAML 2.0 application
2. Set ACS URL: `https://login.tmprl.cloud/saml/acs/<account-id>`
3. Set Entity ID: `urn:auth0:prod-tmprl:saml-<account-id>`
4. Configure attribute mappings

#### Azure AD

1. Create Enterprise Application
2. Set up SAML SSO
3. Set Reply URL: `https://login.tmprl.cloud/saml/acs/<account-id>`
4. Set Identifier: `urn:auth0:prod-tmprl:saml-<account-id>`

### Step 2: Configure Temporal Cloud

```bash
# Via tcld
tcld account saml configure \
  --idp-metadata-url "https://your-idp.com/metadata.xml"

# Or upload metadata file
tcld account saml configure \
  --idp-metadata-file ./idp-metadata.xml
```

### Step 3: Enable SAML

```bash
tcld account saml enable
```

## Attribute Mapping

| SAML Attribute | Temporal Attribute | Required |
| -------------- | ------------------ | -------- |
| email          | User email         | Yes      |
| firstName      | First name         | No       |
| lastName       | Last name          | No       |
| groups         | Role mapping       | No       |

## Group-to-Role Mapping

Map IdP groups to Temporal Cloud roles:

```json
{
  "group_mappings": [
    {
      "idp_group": "temporal-admins",
      "account_role": "global_admin"
    },
    {
      "idp_group": "temporal-developers",
      "account_role": "developer"
    },
    {
      "idp_group": "temporal-readonly",
      "account_role": "read_only"
    }
  ]
}
```

## Just-in-Time Provisioning

When enabled, users are automatically created on first login:

- User created with mapped role
- Email from SAML assertion
- Name from SAML attributes (if provided)

```bash
tcld account saml configure \
  --jit-provisioning enabled \
  --default-role developer
```

## Pricing

| Plan             | SAML SSO     |
| ---------------- | ------------ |
| Essential        | Not included |
| Business         | Included     |
| Enterprise       | Included     |
| Mission Critical | Included     |

## Troubleshooting

### Login Failed

1. Verify IdP metadata is correct
2. Check ACS URL matches exactly
3. Verify Entity ID matches
4. Check SAML assertion contains email

### User Not Created

1. Verify JIT provisioning is enabled
2. Check email attribute is mapped
3. Verify user limit not reached

### Role Not Assigned

1. Check group mappings configuration
2. Verify groups attribute is sent
3. Check group name matches exactly

## Security Considerations

1. **Require SAML**: Disable password login after SAML setup
2. **Session timeout**: Configure appropriate session duration
3. **Group sync**: Use SCIM for real-time group sync
4. **Audit**: Monitor SAML login events
