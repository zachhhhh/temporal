package interceptors

import (
	"context"
	"encoding/json"
	"net"
	"strings"
	"time"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	"go.temporal.io/cloud/internal/service"
	"go.temporal.io/server/common/log"
)

// AuditInterceptor logs all API requests for audit purposes.
type AuditInterceptor struct {
	auditService *service.AuditService
	logger       log.Logger
}

// NewAuditInterceptor creates a new audit interceptor.
func NewAuditInterceptor(auditService *service.AuditService, logger log.Logger) *AuditInterceptor {
	return &AuditInterceptor{
		auditService: auditService,
		logger:       logger,
	}
}

// WrapUnary implements connect.Interceptor.
func (i *AuditInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		startTime := time.Now()

		// Execute the request
		resp, err := next(ctx, req)

		// Log the audit event asynchronously
		go i.logAuditEvent(ctx, req, resp, err, startTime)

		return resp, err
	}
}

// WrapStreamingClient implements connect.Interceptor.
func (i *AuditInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return next
}

// WrapStreamingHandler implements connect.Interceptor.
func (i *AuditInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return next
}

func (i *AuditInterceptor) logAuditEvent(ctx context.Context, req connect.AnyRequest, resp connect.AnyResponse, err error, startTime time.Time) {
	// Skip health checks
	if strings.Contains(req.Spec().Procedure, "health") {
		return
	}

	authInfo := GetAuthInfo(ctx)
	if authInfo == nil {
		return
	}

	// Determine result
	result := "success"
	if err != nil {
		if connect.CodeOf(err) == connect.CodePermissionDenied {
			result = "denied"
		} else {
			result = "failure"
		}
	}

	// Extract resource info from procedure name
	procedure := req.Spec().Procedure
	parts := strings.Split(procedure, "/")
	action := ""
	if len(parts) > 0 {
		action = parts[len(parts)-1]
	}

	// Get IP address
	var ipAddr *net.IP
	if ip := req.Header().Get("X-Forwarded-For"); ip != "" {
		parsed := net.ParseIP(strings.Split(ip, ",")[0])
		ipAddr = &parsed
	}

	event := &service.AuditEventInput{
		OrganizationID: authInfo.OrganizationID,
		ActorType:      i.getActorType(authInfo),
		ActorID:        authInfo.UserID.String(),
		ActorEmail:     authInfo.Email,
		Action:         action,
		Result:         result,
		ResourceType:   i.extractResourceType(procedure),
		RequestID:      req.Header().Get("X-Request-ID"),
		IPAddress:      ipAddr,
		UserAgent:      req.Header().Get("User-Agent"),
		Method:         req.Spec().Procedure,
		Duration:       time.Since(startTime),
	}

	// Add request details (sanitized)
	if details := i.sanitizeRequest(req.Any()); details != nil {
		event.Details = details
	}

	if err := i.auditService.LogEvent(context.Background(), event); err != nil {
		i.logger.Warn("Failed to log audit event")
	}
}

func (i *AuditInterceptor) getActorType(authInfo *AuthInfo) string {
	if authInfo.IsAPIKey {
		return "service_account"
	}
	return "user"
}

func (i *AuditInterceptor) extractResourceType(procedure string) string {
	if strings.Contains(procedure, "Organization") {
		return "organization"
	}
	if strings.Contains(procedure, "Namespace") {
		return "namespace"
	}
	if strings.Contains(procedure, "User") {
		return "user"
	}
	if strings.Contains(procedure, "APIKey") {
		return "api_key"
	}
	if strings.Contains(procedure, "Subscription") || strings.Contains(procedure, "Billing") {
		return "billing"
	}
	return "unknown"
}

func (i *AuditInterceptor) sanitizeRequest(req any) json.RawMessage {
	// Convert to JSON and remove sensitive fields
	data, err := json.Marshal(req)
	if err != nil {
		return nil
	}

	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		return nil
	}

	// Remove sensitive fields
	sensitiveFields := []string{"password", "secret", "token", "key", "certificate"}
	for _, field := range sensitiveFields {
		delete(m, field)
	}

	result, _ := json.Marshal(m)
	return result
}

// RequestIDKey is the context key for request ID.
type requestIDKey struct{}

// WithRequestID adds a request ID to the context.
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey{}, requestID)
}

// GetRequestID retrieves the request ID from context.
func GetRequestID(ctx context.Context) string {
	if id, ok := ctx.Value(requestIDKey{}).(string); ok {
		return id
	}
	return uuid.New().String()
}
