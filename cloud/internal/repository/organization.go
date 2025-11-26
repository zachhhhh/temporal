package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Organization represents an organization in the database.
type Organization struct {
	ID        uuid.UUID
	Name      string
	Slug      string
	Settings  json.RawMessage
	CreatedAt time.Time
	UpdatedAt time.Time
}

// OrganizationMember represents an organization member.
type OrganizationMember struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	UserID         uuid.UUID
	Role           string
	CreatedAt      time.Time
}

// OrganizationRepository handles organization data access.
type OrganizationRepository struct {
	db *PostgresDB
}

// NewOrganizationRepository creates a new organization repository.
func NewOrganizationRepository(db *PostgresDB) *OrganizationRepository {
	return &OrganizationRepository{db: db}
}

// Create creates a new organization.
func (r *OrganizationRepository) Create(ctx context.Context, org *Organization) error {
	query := `
		INSERT INTO organizations (id, name, slug, settings, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	if org.ID == uuid.Nil {
		org.ID = uuid.New()
	}
	if org.Settings == nil {
		org.Settings = json.RawMessage("{}")
	}
	now := time.Now()
	org.CreatedAt = now
	org.UpdatedAt = now

	_, err := r.db.DB().ExecContext(ctx, query,
		org.ID, org.Name, org.Slug, org.Settings, org.CreatedAt, org.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create organization: %w", err)
	}
	return nil
}

// GetByID retrieves an organization by ID.
func (r *OrganizationRepository) GetByID(ctx context.Context, id uuid.UUID) (*Organization, error) {
	query := `
		SELECT id, name, slug, settings, created_at, updated_at
		FROM organizations
		WHERE id = $1
	`
	org := &Organization{}
	err := r.db.DB().QueryRowContext(ctx, query, id).Scan(
		&org.ID, &org.Name, &org.Slug, &org.Settings, &org.CreatedAt, &org.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get organization: %w", err)
	}
	return org, nil
}

// GetBySlug retrieves an organization by slug.
func (r *OrganizationRepository) GetBySlug(ctx context.Context, slug string) (*Organization, error) {
	query := `
		SELECT id, name, slug, settings, created_at, updated_at
		FROM organizations
		WHERE slug = $1
	`
	org := &Organization{}
	err := r.db.DB().QueryRowContext(ctx, query, slug).Scan(
		&org.ID, &org.Name, &org.Slug, &org.Settings, &org.CreatedAt, &org.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get organization by slug: %w", err)
	}
	return org, nil
}

// Update updates an organization.
func (r *OrganizationRepository) Update(ctx context.Context, org *Organization) error {
	query := `
		UPDATE organizations
		SET name = $2, slug = $3, settings = $4, updated_at = $5
		WHERE id = $1
	`
	org.UpdatedAt = time.Now()
	_, err := r.db.DB().ExecContext(ctx, query,
		org.ID, org.Name, org.Slug, org.Settings, org.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to update organization: %w", err)
	}
	return nil
}

// Delete deletes an organization.
func (r *OrganizationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM organizations WHERE id = $1`
	_, err := r.db.DB().ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete organization: %w", err)
	}
	return nil
}

// List lists organizations with pagination.
func (r *OrganizationRepository) List(ctx context.Context, limit, offset int) ([]*Organization, error) {
	query := `
		SELECT id, name, slug, settings, created_at, updated_at
		FROM organizations
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.DB().QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list organizations: %w", err)
	}
	defer rows.Close()

	var orgs []*Organization
	for rows.Next() {
		org := &Organization{}
		if err := rows.Scan(
			&org.ID, &org.Name, &org.Slug, &org.Settings, &org.CreatedAt, &org.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan organization: %w", err)
		}
		orgs = append(orgs, org)
	}
	return orgs, nil
}

// ListByUserID lists organizations for a user.
func (r *OrganizationRepository) ListByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*Organization, error) {
	query := `
		SELECT o.id, o.name, o.slug, o.settings, o.created_at, o.updated_at
		FROM organizations o
		JOIN organization_members om ON o.id = om.organization_id
		WHERE om.user_id = $1
		ORDER BY o.created_at DESC
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.DB().QueryContext(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list organizations by user: %w", err)
	}
	defer rows.Close()

	var orgs []*Organization
	for rows.Next() {
		org := &Organization{}
		if err := rows.Scan(
			&org.ID, &org.Name, &org.Slug, &org.Settings, &org.CreatedAt, &org.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan organization: %w", err)
		}
		orgs = append(orgs, org)
	}
	return orgs, nil
}

// AddMember adds a member to an organization.
func (r *OrganizationRepository) AddMember(ctx context.Context, member *OrganizationMember) error {
	query := `
		INSERT INTO organization_members (id, organization_id, user_id, role, created_at)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (organization_id, user_id) DO UPDATE SET role = $4
	`
	if member.ID == uuid.Nil {
		member.ID = uuid.New()
	}
	member.CreatedAt = time.Now()

	_, err := r.db.DB().ExecContext(ctx, query,
		member.ID, member.OrganizationID, member.UserID, member.Role, member.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to add organization member: %w", err)
	}
	return nil
}

// RemoveMember removes a member from an organization.
func (r *OrganizationRepository) RemoveMember(ctx context.Context, orgID, userID uuid.UUID) error {
	query := `DELETE FROM organization_members WHERE organization_id = $1 AND user_id = $2`
	_, err := r.db.DB().ExecContext(ctx, query, orgID, userID)
	if err != nil {
		return fmt.Errorf("failed to remove organization member: %w", err)
	}
	return nil
}

// GetMember retrieves a member from an organization.
func (r *OrganizationRepository) GetMember(ctx context.Context, orgID, userID uuid.UUID) (*OrganizationMember, error) {
	query := `
		SELECT id, organization_id, user_id, role, created_at
		FROM organization_members
		WHERE organization_id = $1 AND user_id = $2
	`
	member := &OrganizationMember{}
	err := r.db.DB().QueryRowContext(ctx, query, orgID, userID).Scan(
		&member.ID, &member.OrganizationID, &member.UserID, &member.Role, &member.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get organization member: %w", err)
	}
	return member, nil
}

// ListMembers lists members of an organization.
func (r *OrganizationRepository) ListMembers(ctx context.Context, orgID uuid.UUID, limit, offset int) ([]*OrganizationMember, error) {
	query := `
		SELECT id, organization_id, user_id, role, created_at
		FROM organization_members
		WHERE organization_id = $1
		ORDER BY created_at ASC
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.DB().QueryContext(ctx, query, orgID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list organization members: %w", err)
	}
	defer rows.Close()

	var members []*OrganizationMember
	for rows.Next() {
		member := &OrganizationMember{}
		if err := rows.Scan(
			&member.ID, &member.OrganizationID, &member.UserID, &member.Role, &member.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan organization member: %w", err)
		}
		members = append(members, member)
	}
	return members, nil
}
