# Email Infrastructure

## Email Services

| Type          | Provider         | Use Case                      |
| ------------- | ---------------- | ----------------------------- |
| Transactional | SendGrid         | Receipts, alerts, invitations |
| Marketing     | SendGrid         | Newsletters, announcements    |
| Support       | Zendesk          | Support tickets               |
| Internal      | Google Workspace | Employee email                |

## SendGrid Configuration

### Account Setup

```yaml
# Environment variables
SENDGRID_API_KEY: ${secrets.sendgrid_api_key}
SENDGRID_FROM_EMAIL: notifications@temporal-cloud.io
SENDGRID_FROM_NAME: Temporal Cloud
```

### Domain Authentication

1. **SPF Record**: Already in `dns-domains.md`
2. **DKIM**: CNAME records pointing to SendGrid
3. **Domain Verification**: TXT record

### Terraform Setup

```hcl
# terraform/modules/email/main.tf

resource "sendgrid_domain_authentication" "main" {
  domain = "temporal-cloud.io"
}

resource "sendgrid_link_branding" "main" {
  domain  = "temporal-cloud.io"
  default = true
}
```

## Email Templates

### Template Categories

| Category       | Templates                                                       |
| -------------- | --------------------------------------------------------------- |
| Authentication | Welcome, Email Verification, Password Reset, MFA Setup          |
| Billing        | Invoice, Payment Failed, Payment Reminder, Subscription Changed |
| Alerts         | Usage Warning, Quota Exceeded, Certificate Expiring             |
| Team           | User Invited, User Removed, Role Changed                        |
| System         | Maintenance Notice, Incident Update, Feature Announcement       |

### Template Structure

```
email-templates/
├── base.html                 # Base layout
├── auth/
│   ├── welcome.html
│   ├── verify-email.html
│   ├── password-reset.html
│   └── mfa-setup.html
├── billing/
│   ├── invoice.html
│   ├── payment-failed.html
│   ├── payment-reminder.html
│   └── subscription-changed.html
├── alerts/
│   ├── usage-warning.html
│   ├── quota-exceeded.html
│   └── cert-expiring.html
└── team/
    ├── user-invited.html
    ├── user-removed.html
    └── role-changed.html
```

### Base Template

```html
<!-- templates/base.html -->
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>{{.Subject}}</title>
    <style>
      body {
        font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto,
          sans-serif;
      }
      .container {
        max-width: 600px;
        margin: 0 auto;
        padding: 20px;
      }
      .header {
        background: #1a1a2e;
        color: white;
        padding: 20px;
        text-align: center;
      }
      .content {
        padding: 20px;
        background: #ffffff;
      }
      .button {
        background: #6366f1;
        color: white;
        padding: 12px 24px;
        text-decoration: none;
        border-radius: 6px;
        display: inline-block;
      }
      .footer {
        padding: 20px;
        text-align: center;
        color: #666;
        font-size: 12px;
      }
    </style>
  </head>
  <body>
    <div class="container">
      <div class="header">
        <img
          src="https://temporal-cloud.io/logo.png"
          alt="Temporal Cloud"
          width="150"
        />
      </div>
      <div class="content">{{template "content" .}}</div>
      <div class="footer">
        <p>Temporal Technologies, Inc.</p>
        <p>
          <a
            href="https://temporal-cloud.io/unsubscribe?token={{.UnsubscribeToken}}"
            >Unsubscribe</a
          >
          |
          <a href="https://temporal-cloud.io/preferences">Email Preferences</a>
        </p>
      </div>
    </div>
  </body>
</html>
```

### Example: Invoice Template

```html
<!-- templates/billing/invoice.html -->
{{define "content"}}
<h1>Invoice #{{.InvoiceNumber}}</h1>

<p>Hi {{.CustomerName}},</p>

<p>Your invoice for {{.BillingPeriod}} is ready.</p>

<table style="width: 100%; border-collapse: collapse;">
  <tr style="background: #f5f5f5;">
    <th style="padding: 10px; text-align: left;">Description</th>
    <th style="padding: 10px; text-align: right;">Amount</th>
  </tr>
  {{range .LineItems}}
  <tr>
    <td style="padding: 10px; border-bottom: 1px solid #eee;">
      {{.Description}}
    </td>
    <td
      style="padding: 10px; border-bottom: 1px solid #eee; text-align: right;"
    >
      ${{.Amount}}
    </td>
  </tr>
  {{end}}
  <tr style="font-weight: bold;">
    <td style="padding: 10px;">Total</td>
    <td style="padding: 10px; text-align: right;">${{.Total}}</td>
  </tr>
</table>

<p style="margin-top: 20px;">
  <a href="{{.InvoiceURL}}" class="button">View Invoice</a>
</p>

<p style="color: #666; font-size: 14px;">
  This invoice will be automatically charged to your payment method on file.
</p>
{{end}}
```

## Email Service Implementation

### Go Email Client

```go
// internal/email/client.go
package email

import (
    "github.com/sendgrid/sendgrid-go"
    "github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Client struct {
    sg        *sendgrid.Client
    fromEmail string
    fromName  string
    templates *template.Template
}

func NewClient(apiKey, fromEmail, fromName string) *Client {
    return &Client{
        sg:        sendgrid.NewSendClient(apiKey),
        fromEmail: fromEmail,
        fromName:  fromName,
        templates: loadTemplates(),
    }
}

func (c *Client) SendEmail(ctx context.Context, req SendEmailRequest) error {
    // Render template
    var body bytes.Buffer
    if err := c.templates.ExecuteTemplate(&body, req.Template, req.Data); err != nil {
        return fmt.Errorf("render template: %w", err)
    }

    // Build message
    from := mail.NewEmail(c.fromName, c.fromEmail)
    to := mail.NewEmail(req.ToName, req.ToEmail)
    message := mail.NewSingleEmail(from, req.Subject, to, "", body.String())

    // Add tracking
    message.SetTrackingSettings(&mail.TrackingSettings{
        ClickTracking: &mail.ClickTrackingSetting{Enable: true},
        OpenTracking:  &mail.OpenTrackingSetting{Enable: true},
    })

    // Send
    response, err := c.sg.SendWithContext(ctx, message)
    if err != nil {
        return fmt.Errorf("send email: %w", err)
    }

    if response.StatusCode >= 400 {
        return fmt.Errorf("sendgrid error: %d %s", response.StatusCode, response.Body)
    }

    return nil
}

type SendEmailRequest struct {
    ToEmail  string
    ToName   string
    Subject  string
    Template string
    Data     any
}
```

### Email Types

```go
// Predefined email senders
func (c *Client) SendWelcomeEmail(ctx context.Context, user *User) error {
    return c.SendEmail(ctx, SendEmailRequest{
        ToEmail:  user.Email,
        ToName:   user.Name,
        Subject:  "Welcome to Temporal Cloud",
        Template: "auth/welcome.html",
        Data: map[string]any{
            "Name":        user.Name,
            "ConsoleURL":  "https://console.temporal-cloud.io",
            "DocsURL":     "https://docs.temporal.io",
        },
    })
}

func (c *Client) SendInvoiceEmail(ctx context.Context, invoice *Invoice) error {
    return c.SendEmail(ctx, SendEmailRequest{
        ToEmail:  invoice.CustomerEmail,
        ToName:   invoice.CustomerName,
        Subject:  fmt.Sprintf("Invoice #%s from Temporal Cloud", invoice.Number),
        Template: "billing/invoice.html",
        Data:     invoice,
    })
}

func (c *Client) SendPaymentFailedEmail(ctx context.Context, invoice *Invoice) error {
    return c.SendEmail(ctx, SendEmailRequest{
        ToEmail:  invoice.CustomerEmail,
        ToName:   invoice.CustomerName,
        Subject:  "Payment Failed - Action Required",
        Template: "billing/payment-failed.html",
        Data: map[string]any{
            "InvoiceNumber": invoice.Number,
            "Amount":        invoice.Total,
            "UpdateURL":     "https://console.temporal-cloud.io/billing/payment",
        },
    })
}
```

## Email Workflows (Temporal)

### Drip Campaign Workflow

```go
// Onboarding email sequence
func OnboardingEmailWorkflow(ctx workflow.Context, userID string) error {
    // Day 0: Welcome email
    _ = workflow.ExecuteActivity(ctx, SendWelcomeEmailActivity, userID).Get(ctx, nil)

    // Day 1: Getting started guide
    workflow.Sleep(ctx, 24*time.Hour)
    _ = workflow.ExecuteActivity(ctx, SendGettingStartedEmailActivity, userID).Get(ctx, nil)

    // Day 3: First workflow tips
    workflow.Sleep(ctx, 48*time.Hour)
    _ = workflow.ExecuteActivity(ctx, SendWorkflowTipsEmailActivity, userID).Get(ctx, nil)

    // Day 7: Check in
    workflow.Sleep(ctx, 96*time.Hour)

    // Only send if no workflows created
    var hasWorkflows bool
    _ = workflow.ExecuteActivity(ctx, CheckUserHasWorkflowsActivity, userID).Get(ctx, &hasWorkflows)

    if !hasWorkflows {
        _ = workflow.ExecuteActivity(ctx, SendNeedHelpEmailActivity, userID).Get(ctx, nil)
    }

    return nil
}
```

### Payment Reminder Workflow

```go
func PaymentReminderWorkflow(ctx workflow.Context, invoiceID string) error {
    schedule := []struct {
        delay   time.Duration
        urgency string
    }{
        {3 * 24 * time.Hour, "gentle"},
        {7 * 24 * time.Hour, "reminder"},
        {10 * 24 * time.Hour, "urgent"},
        {14 * 24 * time.Hour, "final"},
    }

    for _, step := range schedule {
        workflow.Sleep(ctx, step.delay)

        // Check if paid
        var paid bool
        _ = workflow.ExecuteActivity(ctx, CheckInvoicePaidActivity, invoiceID).Get(ctx, &paid)
        if paid {
            return nil
        }

        // Send reminder
        _ = workflow.ExecuteActivity(ctx, SendPaymentReminderActivity, invoiceID, step.urgency).Get(ctx, nil)
    }

    return nil
}
```

## Email Deliverability

### Monitoring

```yaml
# Metrics to track
- delivery_rate # Target: > 99%
- open_rate # Target: > 20%
- click_rate # Target: > 5%
- bounce_rate # Target: < 2%
- spam_complaint_rate # Target: < 0.1%
```

### Bounce Handling

```go
// Webhook handler for SendGrid events
func HandleSendGridWebhook(events []SendGridEvent) {
    for _, event := range events {
        switch event.Type {
        case "bounce":
            // Mark email as invalid
            markEmailInvalid(event.Email, event.Reason)

        case "spamreport":
            // Unsubscribe user
            unsubscribeUser(event.Email)
            log.Warn("Spam report", "email", event.Email)

        case "unsubscribe":
            unsubscribeUser(event.Email)
        }
    }
}
```

### Suppression List

Never email:

- Bounced addresses (hard bounce)
- Spam reporters
- Unsubscribed users
- Invalid domains

## AWS SES (Backup)

```hcl
# Backup email provider
resource "aws_ses_domain_identity" "main" {
  domain = "temporal-cloud.io"
}

resource "aws_ses_domain_dkim" "main" {
  domain = aws_ses_domain_identity.main.domain
}

# Configuration set for tracking
resource "aws_ses_configuration_set" "main" {
  name = "temporal-cloud"
}
```
