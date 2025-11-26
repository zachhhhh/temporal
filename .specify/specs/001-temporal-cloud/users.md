# User Management

## Account-Level Roles

| Role          | Permissions                                           |
| ------------- | ----------------------------------------------------- |
| Account Owner | Full access including billing, can transfer ownership |
| Global Admin  | Full access except billing, can manage all users      |
| Finance Admin | Billing access only, otherwise read-only              |
| Developer     | Create namespaces, admin on own namespaces            |
| Read-Only     | View only, no modifications                           |

### Role Capabilities

| Capability            | Owner | Global Admin | Finance | Developer | Read-Only |
| --------------------- | ----- | ------------ | ------- | --------- | --------- |
| Manage billing        | ✅    | ❌           | ✅      | ❌        | ❌        |
| Manage users          | ✅    | ✅           | ❌      | ❌        | ❌        |
| Create namespaces     | ✅    | ✅           | ❌      | ✅        | ❌        |
| Manage all namespaces | ✅    | ✅           | ❌      | ❌        | ❌        |
| View usage            | ✅    | ✅           | ✅      | ✅        | ✅        |
| Configure SSO         | ✅    | ✅           | ❌      | ❌        | ❌        |

## Namespace-Level Permissions

| Permission      | Capabilities                               |
| --------------- | ------------------------------------------ |
| Namespace Admin | Full namespace control, manage permissions |
| Write           | Create, update, delete workflows           |
| Read-Only       | View workflows and history                 |

### Permission Inheritance

- Global Admin → Namespace Admin on all namespaces
- Developer → Namespace Admin on namespaces they create

## Inviting Users

### Via Console

1. Go to Users → Invite User
2. Enter email address
3. Select account role
4. Optionally assign namespace permissions
5. Send invitation

### Via tcld

```bash
tcld user invite \
  --email user@example.com \
  --account-role developer \
  --namespace-permission my-namespace:write
```

### Via API

```bash
curl -X POST https://api.temporal.io/api/v1/users/invite \
  -H "Authorization: Bearer $API_KEY" \
  -d '{
    "email": "user@example.com",
    "account_role": "developer"
  }'
```

## Managing Roles

### Update Account Role

```bash
tcld user update \
  --email user@example.com \
  --account-role global_admin
```

### Update Namespace Permission

```bash
tcld user namespace-access set \
  --email user@example.com \
  --namespace my-namespace \
  --permission write
```

### Remove Namespace Permission

```bash
tcld user namespace-access remove \
  --email user@example.com \
  --namespace my-namespace
```

## Deleting Users

```bash
tcld user delete --email user@example.com
```

**Effects:**

- User removed from account
- All namespace permissions revoked
- API keys deleted
- Active sessions terminated

## Account Owner

### Best Practices

- Assign to at least 2 users
- Use personal email (not shared)
- Enable MFA

### Transferring Ownership

Contact Temporal Support to transfer Account Owner role.

## User Groups

Groups allow bulk permission management.

### Create Group

```bash
tcld user-group create \
  --name "Backend Team" \
  --description "Backend engineering team"
```

### Add Members

```bash
tcld user-group member add \
  --group-id grp-123 \
  --email user@example.com
```

### Set Namespace Permissions

```bash
tcld user-group namespace-access set \
  --group-id grp-123 \
  --namespace my-namespace \
  --permission write
```

## Service Accounts

Machine identities for automation.

### Create Service Account

```bash
tcld service-account create \
  --name "CI Pipeline" \
  --account-role developer
```

### Namespace-Scoped Service Account

```bash
tcld service-account create \
  --name "Worker" \
  --namespace my-namespace \
  --permission write
```

## Limits

| Limit             | Value         |
| ----------------- | ------------- |
| Users per account | 300 (default) |
| Service accounts  | 100           |
| User groups       | 50            |
| Members per group | 100           |
