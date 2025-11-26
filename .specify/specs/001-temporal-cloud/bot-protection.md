# Bot Protection

## Threat Model

| Bot Type            | Impact             | Mitigation                           |
| ------------------- | ------------------ | ------------------------------------ |
| Credential stuffing | Account takeover   | Rate limiting, MFA, breach detection |
| Scraping            | Data exfiltration  | Rate limiting, fingerprinting        |
| DDoS                | Service disruption | WAF, Shield, auto-scaling            |
| Spam/abuse          | Resource waste     | CAPTCHA, reputation                  |
| API abuse           | Quota bypass       | API keys, rate limiting              |

## Multi-Layer Defense

```
┌─────────────────────────────────────────────────────────────────┐
│  Layer 1: Edge (Cloudflare/AWS WAF)                             │
│  - IP reputation                                                 │
│  - Known bot signatures                                          │
│  - JavaScript challenge                                          │
├─────────────────────────────────────────────────────────────────┤
│  Layer 2: Rate Limiting                                          │
│  - Per-IP limits                                                 │
│  - Per-account limits                                            │
│  - Per-endpoint limits                                           │
├─────────────────────────────────────────────────────────────────┤
│  Layer 3: Behavioral Analysis                                    │
│  - Request patterns                                              │
│  - Session anomalies                                             │
│  - Impossible travel                                             │
├─────────────────────────────────────────────────────────────────┤
│  Layer 4: Challenge/Response                                     │
│  - CAPTCHA (high-risk actions)                                  │
│  - Email verification                                            │
│  - MFA                                                          │
└─────────────────────────────────────────────────────────────────┘
```

## Rate Limiting

### Implementation

```go
type RateLimiter struct {
    redis *redis.Client
}

type RateLimitConfig struct {
    // Per IP
    IPRequestsPerMinute     int
    IPRequestsPerHour       int

    // Per account
    AccountRequestsPerMinute int
    AccountRequestsPerHour   int

    // Per endpoint
    LoginAttemptsPerMinute   int
    SignupAttemptsPerHour    int
}

func (rl *RateLimiter) Check(ctx context.Context, key string, limit int, window time.Duration) (bool, int, error) {
    now := time.Now()
    windowKey := fmt.Sprintf("%s:%d", key, now.Unix()/int64(window.Seconds()))

    pipe := rl.redis.Pipeline()
    incr := pipe.Incr(ctx, windowKey)
    pipe.Expire(ctx, windowKey, window)
    pipe.Exec(ctx)

    count := int(incr.Val())
    remaining := limit - count

    if count > limit {
        return false, 0, nil
    }

    return true, remaining, nil
}
```

### Rate Limits by Endpoint

| Endpoint              | Per IP/min | Per Account/min | Per IP/hour |
| --------------------- | ---------- | --------------- | ----------- |
| Login                 | 10         | N/A             | 100         |
| Signup                | 5          | N/A             | 20          |
| Password Reset        | 3          | 3               | 10          |
| API (authenticated)   | 1000       | 500             | 10000       |
| API (unauthenticated) | 100        | N/A             | 1000        |

### Response Headers

```
X-RateLimit-Limit: 1000
X-RateLimit-Remaining: 999
X-RateLimit-Reset: 1704067200
Retry-After: 60
```

## Bot Detection

### Device Fingerprinting

```typescript
// Client-side fingerprint collection
import FingerprintJS from "@fingerprintjs/fingerprintjs";

async function getFingerprint() {
  const fp = await FingerprintJS.load();
  const result = await fp.get();
  return result.visitorId;
}

// Send with requests
fetch("/api/action", {
  headers: {
    "X-Device-Fingerprint": await getFingerprint(),
  },
});
```

### Behavioral Signals

```go
type BehaviorAnalyzer struct {
    redis *redis.Client
}

type RequestSignals struct {
    IP            string
    UserAgent     string
    Fingerprint   string
    RequestPath   string
    RequestMethod string
    Timestamp     time.Time
    SessionID     string
}

func (ba *BehaviorAnalyzer) AnalyzeRequest(ctx context.Context, signals RequestSignals) (riskScore float64, reasons []string) {
    // Check velocity
    if ba.isVelocityTooHigh(ctx, signals.IP) {
        riskScore += 0.3
        reasons = append(reasons, "high_velocity")
    }

    // Check user agent
    if ba.isSuspiciousUserAgent(signals.UserAgent) {
        riskScore += 0.2
        reasons = append(reasons, "suspicious_ua")
    }

    // Check fingerprint consistency
    if ba.isFingerprintMismatch(ctx, signals.SessionID, signals.Fingerprint) {
        riskScore += 0.4
        reasons = append(reasons, "fingerprint_mismatch")
    }

    // Check impossible travel
    if ba.isImpossibleTravel(ctx, signals.SessionID, signals.IP) {
        riskScore += 0.5
        reasons = append(reasons, "impossible_travel")
    }

    return riskScore, reasons
}
```

### IP Reputation

```go
// Check IP against threat intelligence
func CheckIPReputation(ip string) (score float64, threats []string) {
    // Check against known bad IP lists
    if isInBotnet(ip) {
        return 1.0, []string{"botnet"}
    }

    // Check against Tor exit nodes
    if isTorExitNode(ip) {
        return 0.8, []string{"tor"}
    }

    // Check against VPN/proxy lists
    if isVPN(ip) || isProxy(ip) {
        return 0.3, []string{"vpn_proxy"}
    }

    // Check recent abuse reports
    abuseScore := getAbuseScore(ip)
    if abuseScore > 0.5 {
        return abuseScore, []string{"abuse_reports"}
    }

    return 0.0, nil
}
```

## CAPTCHA Integration

### When to Show CAPTCHA

| Trigger                             | Action       |
| ----------------------------------- | ------------ |
| 5+ failed logins                    | Show CAPTCHA |
| Suspicious fingerprint              | Show CAPTCHA |
| High-risk signup (disposable email) | Show CAPTCHA |
| Rate limit approaching              | Show CAPTCHA |
| VPN/Tor detected                    | Show CAPTCHA |

### Implementation (hCaptcha)

```typescript
// React component
import HCaptcha from "@hcaptcha/react-hcaptcha";

function LoginForm() {
  const [captchaToken, setCaptchaToken] = useState<string | null>(null);
  const [showCaptcha, setShowCaptcha] = useState(false);

  const onSubmit = async (data: LoginData) => {
    const response = await fetch("/api/login", {
      method: "POST",
      body: JSON.stringify({
        ...data,
        captcha_token: captchaToken,
      }),
    });

    if (response.status === 428) {
      // Server requires CAPTCHA
      setShowCaptcha(true);
    }
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      {/* ... form fields ... */}

      {showCaptcha && (
        <HCaptcha
          sitekey={process.env.HCAPTCHA_SITE_KEY}
          onVerify={setCaptchaToken}
        />
      )}

      <button type="submit">Login</button>
    </form>
  );
}
```

### Server-Side Verification

```go
func VerifyCaptcha(token string) (bool, error) {
    resp, err := http.PostForm("https://hcaptcha.com/siteverify", url.Values{
        "secret":   {os.Getenv("HCAPTCHA_SECRET")},
        "response": {token},
    })

    var result struct {
        Success bool `json:"success"`
    }
    json.NewDecoder(resp.Body).Decode(&result)

    return result.Success, nil
}
```

## Account Protection

### Credential Stuffing Prevention

```go
// Check password against breach databases
func IsBreachedPassword(password string) bool {
    // Use k-anonymity to check against HaveIBeenPwned
    hash := sha1.Sum([]byte(password))
    prefix := hex.EncodeToString(hash[:])[:5]
    suffix := hex.EncodeToString(hash[:])[5:]

    resp, _ := http.Get("https://api.pwnedpasswords.com/range/" + prefix)
    body, _ := io.ReadAll(resp.Body)

    return strings.Contains(string(body), strings.ToUpper(suffix))
}

// On registration/password change
if IsBreachedPassword(newPassword) {
    return errors.New("This password has been found in a data breach. Please choose a different password.")
}
```

### Account Lockout

```go
func HandleLoginAttempt(ctx context.Context, email string, success bool) error {
    key := "login_attempts:" + email

    if success {
        // Clear attempts on success
        redis.Del(ctx, key)
        return nil
    }

    // Increment failed attempts
    attempts := redis.Incr(ctx, key).Val()
    redis.Expire(ctx, key, 15*time.Minute)

    if attempts >= 5 {
        // Lock account
        lockAccount(ctx, email)
        sendAccountLockedEmail(email)
        return ErrAccountLocked
    }

    return nil
}
```

## Monitoring & Alerts

### Bot Detection Metrics

```yaml
alerts:
  - name: HighBotTraffic
    expr: |
      sum(rate(blocked_requests_total{reason="bot"}[5m])) > 100
    severity: warning

  - name: CredentialStuffingAttack
    expr: |
      sum(rate(failed_logins_total[5m])) by (ip) > 10
    severity: critical

  - name: ScrapingDetected
    expr: |
      sum(rate(requests_total[5m])) by (ip) > 1000
      AND sum(rate(requests_total[5m])) by (user_id) == 0
    severity: warning
```

### Response Playbook

| Alert               | Response                               |
| ------------------- | -------------------------------------- |
| High bot traffic    | Review WAF rules, add blocks           |
| Credential stuffing | Block IP ranges, notify affected users |
| Scraping            | Add rate limits, consider legal action |
| DDoS                | Activate Shield, scale infrastructure  |
