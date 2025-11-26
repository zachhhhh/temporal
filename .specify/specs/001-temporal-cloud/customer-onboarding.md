# Customer Onboarding

## Onboarding Flow

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Signup    │────▶│   Verify    │────▶│   Setup     │
│   Form      │     │   Email     │     │   Account   │
└─────────────┘     └─────────────┘     └─────────────┘
                                               │
                                               ▼
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   First     │◀────│   Create    │◀────│   Select    │
│   Workflow  │     │   Namespace │     │   Plan      │
└─────────────┘     └─────────────┘     └─────────────┘
```

## Step 1: Signup

### Form Fields

- Email (required)
- Full Name (required)
- Company Name (required)
- Password (8+ chars, 1 uppercase, 1 number)

### Validations

- Email not already registered
- Email domain not blocked (disposable emails)
- Password strength check

### Actions

1. Create user record (unverified)
2. Create organization record
3. Send verification email
4. Log signup event to analytics

## Step 2: Email Verification

### Email Content

```
Subject: Verify your Temporal Cloud account

Hi {{name}},

Click the link below to verify your email:
{{verification_link}}

This link expires in 24 hours.
```

### Verification Link

`https://cloud.temporal.io/verify?token={{jwt_token}}`

Token contains: user_id, email, expiry

### Actions

1. Mark user as verified
2. Redirect to account setup

## Step 3: Account Setup

### Organization Details

- Organization name (pre-filled)
- Organization slug (auto-generated, editable)
- Industry (dropdown)
- Company size (dropdown)

### Actions

1. Update organization record
2. Create default settings
3. Proceed to plan selection

## Step 4: Plan Selection

### Plan Display

```
┌─────────────────────────────────────────────────────────┐
│  Essential           Business           Enterprise      │
│  $100/mo             $500/mo            Contact Sales   │
│                                                         │
│  ✓ 1M actions        ✓ 2.5M actions     ✓ Custom        │
│  ✓ 1 GB active       ✓ 2.5 GB active    ✓ Custom        │
│  ✓ Email support     ✓ Chat support     ✓ Dedicated     │
│                      ✓ SSO              ✓ SCIM          │
│                                                         │
│  [Start Free Trial]  [Start Free Trial] [Contact Sales] │
└─────────────────────────────────────────────────────────┘
```

### Free Trial

- 14 days on selected plan
- No credit card required
- Full feature access

### Actions

1. Create subscription (trial status)
2. Set trial_ends_at = now + 14 days
3. Schedule trial expiry reminder workflow

## Step 5: Create First Namespace

### Guided Wizard

```
Step 1 of 3: Name Your Namespace

  Namespace Name: [my-first-namespace]

  This will be part of your endpoint:
  my-first-namespace.abc123.tmprl.cloud

  [Next]
```

```
Step 2 of 3: Select Region

  ○ US East (Virginia)     - Recommended
  ○ US West (Oregon)
  ○ EU West (Ireland)
  ○ Asia Pacific (Singapore)

  [Next]
```

```
Step 3 of 3: Security

  Upload CA Certificate (optional)
  [Drop file here or click to upload]

  You can also generate certificates using our CLI:
  $ tcld generate-certificates --namespace my-first-namespace

  [Skip for now]  [Create Namespace]
```

### Actions

1. Create namespace via Cloud Ops API
2. Wait for namespace to be ready (~30s)
3. Show success with connection details

## Step 6: First Workflow (Getting Started)

### Connection Details

```
Your namespace is ready!

Endpoint: my-first-namespace.abc123.tmprl.cloud:443
Namespace: my-first-namespace

Quick Start:
$ temporal workflow start \
    --address my-first-namespace.abc123.tmprl.cloud:443 \
    --namespace my-first-namespace \
    --task-queue my-queue \
    --type MyWorkflow
```

### Code Examples (Tabs)

- Go
- Java
- TypeScript
- Python

### Next Steps Checklist

- [ ] Download and install SDK
- [ ] Run your first workflow
- [ ] Invite team members
- [ ] Set up billing (before trial ends)

## Onboarding Workflow (Backend)

```go
func CustomerOnboardingWorkflow(ctx workflow.Context, input OnboardingInput) error {
    // Step 1: Create user and org
    var user User
    workflow.ExecuteActivity(ctx, CreateUser, input).Get(ctx, &user)

    // Step 2: Send verification email
    workflow.ExecuteActivity(ctx, SendVerificationEmail, user.Email)

    // Step 3: Wait for verification (up to 7 days)
    verified := workflow.GetSignalChannel(ctx, "email-verified")

    selector := workflow.NewSelector(ctx)
    selector.AddReceive(verified, func(c workflow.ReceiveChannel, more bool) {
        // Continue onboarding
    })
    selector.AddFuture(workflow.NewTimer(ctx, 7*24*time.Hour), func(f workflow.Future) {
        // Send reminder or expire
    })
    selector.Select(ctx)

    // Step 4: Wait for first namespace
    workflow.ExecuteActivity(ctx, WaitForFirstNamespace, user.OrgID)

    // Step 5: Send welcome email with tips
    workflow.ExecuteActivity(ctx, SendWelcomeEmail, user.Email)

    // Step 6: Schedule trial reminders
    workflow.ExecuteActivity(ctx, ScheduleTrialReminders, input.TrialEndsAt)

    return nil
}
```

## Trial Reminders

| Day | Email                                                 |
| --- | ----------------------------------------------------- |
| 7   | "Your trial is halfway through"                       |
| 12  | "2 days left - add payment method"                    |
| 14  | "Trial ended - add payment to continue"               |
| 17  | "Account suspended - data will be deleted in 30 days" |

## Metrics

- Signup → Verified: Target 80%
- Verified → First Namespace: Target 70%
- First Namespace → First Workflow: Target 60%
- Trial → Paid Conversion: Target 25%
