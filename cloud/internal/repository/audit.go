package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/google/uuid"
)

// AuditEvent represents an audit event in the database.
type AuditEvent struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	ActorType      string
	ActorID        string
	ActorEmail     sql.NullString
	ActorName      sql.NullString
	Action         string
	Result         string
	ResourceType   string
	ResourceID     sql.NullString
	ResourceName   sql.NullString
	RequestID      sql.NullString
	IPAddress      *net.IP
	UserAgent      sql.NullString
	Method         sql.NullString
	Path           sql.NullString
	Details        json.RawMessage
	CreatedAt      time.Time
}

// AuditRepository handles audit event data access.
type AuditRepository struct {
	db *PostgresDB
}

// NewAuditRepository creates a new audit repository.
func NewAuditRepository(db *PostgresDB) *AuditRepository {
	return &AuditRepository{db: db}
}

// Create creates a new audit event.
func (r *AuditRepository) Create(ctx context.Context, event *AuditEvent) error {
	query := `
		INSERT INTO audit_events (
			id, organization_id, actor_type, actor_id, actor_email, actor_name,
			action, result, resource_type, resource_id, resource_name,
			request_id, ip_address, user_agent, method, path, details, created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)
	`
	if event.ID == uuid.Nil {
		event.ID = uuid.New()
	}
	event.CreatedAt = time.Now()

	var ipStr sql.NullString
	if event.IPAddress != nil {
		ipStr = sql.NullString{String: event.IPAddress.String(), Valid: true}
	}

	_, err := r.db.DB().ExecContext(ctx, query,
		event.ID, event.OrganizationID, event.ActorType, event.ActorID, event.ActorEmail, event.ActorName,
		event.Action, event.Result, event.ResourceType, event.ResourceID, event.ResourceName,
		event.RequestID, ipStr, event.UserAgent, event.Method, event.Path, event.Details, event.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create audit event: %w", err)
	}
	return nil
}

// GetByID retrieves an audit event by ID.
func (r *AuditRepository) GetByID(ctx context.Context, id uuid.UUID) (*AuditEvent, error) {
	query := `
		SELECT id, organization_id, actor_type, actor_id, actor_email, actor_name,
			action, result, resource_type, resource_id, resource_name,
			request_id, ip_address, user_agent, method, path, details, created_at
		FROM audit_events WHERE id = $1
	`
	event := &AuditEvent{}
	var ipStr sql.NullString
	err := r.db.DB().QueryRowContext(ctx, query, id).Scan(
		&event.ID, &event.OrganizationID, &event.ActorType, &event.ActorID, &event.ActorEmail, &event.ActorName,
		&event.Action, &event.Result, &event.ResourceType, &event.ResourceID, &event.ResourceName,
		&event.RequestID, &ipStr, &event.UserAgent, &event.Method, &event.Path, &event.Details, &event.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get audit event: %w", err)
	}
	if ipStr.Valid {
		ip := net.ParseIP(ipStr.String)
		event.IPAddress = &ip
	}
	return event, nil
}

// ListByOrganization lists audit events for an organization.
func (r *AuditRepository) ListByOrganization(ctx context.Context, orgID uuid.UUID, start, end time.Time, limit, offset int) ([]*AuditEvent, error) {
	query := `
		SELECT id, organization_id, actor_type, actor_id, actor_email, actor_name,
			action, result, resource_type, resource_id, resource_name,
			request_id, ip_address, user_agent, method, path, details, created_at
		FROM audit_events
		WHERE organization_id = $1 AND created_at >= $2 AND created_at <= $3
		ORDER BY created_at DESC LIMIT $4 OFFSET $5
	`
	rows, err := r.db.DB().QueryContext(ctx, query, orgID, start, end, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list audit events: %w", err)
	}
	defer rows.Close()

	var events []*AuditEvent
	for rows.Next() {
		event := &AuditEvent{}
		var ipStr sql.NullString
		if err := rows.Scan(
			&event.ID, &event.OrganizationID, &event.ActorType, &event.ActorID, &event.ActorEmail, &event.ActorName,
			&event.Action, &event.Result, &event.ResourceType, &event.ResourceID, &event.ResourceName,
			&event.RequestID, &ipStr, &event.UserAgent, &event.Method, &event.Path, &event.Details, &event.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan audit event: %w", err)
		}
		if ipStr.Valid {
			ip := net.ParseIP(ipStr.String)
			event.IPAddress = &ip
		}
		events = append(events, event)
	}
	return events, nil
}
