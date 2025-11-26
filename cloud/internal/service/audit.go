package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"net"
	"time"

	"github.com/google/uuid"
	"go.temporal.io/cloud/internal/repository"
	"go.temporal.io/server/common/log"
)

// AuditService handles audit logging business logic.
type AuditService struct {
	repos  *repository.Repositories
	logger log.Logger
}

// NewAuditService creates a new audit service.
func NewAuditService(repos *repository.Repositories, logger log.Logger) *AuditService {
	return &AuditService{repos: repos, logger: logger}
}

// AuditEventInput is the input for logging an audit event.
type AuditEventInput struct {
	OrganizationID uuid.UUID
	ActorType      string
	ActorID        string
	ActorEmail     string
	ActorName      string
	Action         string
	Result         string
	ResourceType   string
	ResourceID     string
	ResourceName   string
	RequestID      string
	IPAddress      *net.IP
	UserAgent      string
	Method         string
	Path           string
	Details        json.RawMessage
	Duration       time.Duration
}

// LogEvent logs an audit event.
func (s *AuditService) LogEvent(ctx context.Context, input *AuditEventInput) error {
	event := &repository.AuditEvent{
		OrganizationID: input.OrganizationID,
		ActorType:      input.ActorType,
		ActorID:        input.ActorID,
		ActorEmail:     sql.NullString{String: input.ActorEmail, Valid: input.ActorEmail != ""},
		ActorName:      sql.NullString{String: input.ActorName, Valid: input.ActorName != ""},
		Action:         input.Action,
		Result:         input.Result,
		ResourceType:   input.ResourceType,
		ResourceID:     sql.NullString{String: input.ResourceID, Valid: input.ResourceID != ""},
		ResourceName:   sql.NullString{String: input.ResourceName, Valid: input.ResourceName != ""},
		RequestID:      sql.NullString{String: input.RequestID, Valid: input.RequestID != ""},
		IPAddress:      input.IPAddress,
		UserAgent:      sql.NullString{String: input.UserAgent, Valid: input.UserAgent != ""},
		Method:         sql.NullString{String: input.Method, Valid: input.Method != ""},
		Path:           sql.NullString{String: input.Path, Valid: input.Path != ""},
		Details:        input.Details,
	}

	return s.repos.Audit.Create(ctx, event)
}

// GetEvent retrieves an audit event by ID.
func (s *AuditService) GetEvent(ctx context.Context, id uuid.UUID) (*repository.AuditEvent, error) {
	return s.repos.Audit.GetByID(ctx, id)
}

// ListEventsInput is the input for listing audit events.
type ListEventsInput struct {
	OrganizationID uuid.UUID
	StartTime      time.Time
	EndTime        time.Time
	ActorID        string
	Action         string
	ResourceType   string
	ResourceID     string
	Result         string
	Limit          int
	Offset         int
}

// ListEvents lists audit events for an organization.
func (s *AuditService) ListEvents(ctx context.Context, input *ListEventsInput) ([]*repository.AuditEvent, error) {
	if input.Limit == 0 {
		input.Limit = 100
	}
	if input.Limit > 1000 {
		input.Limit = 1000
	}

	if input.StartTime.IsZero() {
		input.StartTime = time.Now().AddDate(0, 0, -90) // Default to 90 days
	}
	if input.EndTime.IsZero() {
		input.EndTime = time.Now()
	}

	return s.repos.Audit.ListByOrganization(ctx, input.OrganizationID, input.StartTime, input.EndTime, input.Limit, input.Offset)
}

// Common audit actions
const (
	AuditActionOrganizationCreate = "organization.create"
	AuditActionOrganizationUpdate = "organization.update"
	AuditActionOrganizationDelete = "organization.delete"
	AuditActionUserInvite         = "user.invite"
	AuditActionUserRemove         = "user.remove"
	AuditActionUserRoleUpdate     = "user.role.update"
	AuditActionNamespaceCreate    = "namespace.create"
	AuditActionNamespaceUpdate    = "namespace.update"
	AuditActionNamespaceDelete    = "namespace.delete"
	AuditActionAPIKeyCreate       = "apikey.create"
	AuditActionAPIKeyRevoke       = "apikey.revoke"
	AuditActionAPIKeyRotate       = "apikey.rotate"
	AuditActionSubscriptionUpdate = "subscription.update"
	AuditActionSAMLConfigure      = "saml.configure"
	AuditActionSCIMConfigure      = "scim.configure"
)
