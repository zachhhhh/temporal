package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// User represents a user in the database.
type User struct {
	ID            uuid.UUID
	Email         string
	Name          sql.NullString
	AvatarURL     sql.NullString
	EmailVerified bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// UserAccountRole represents a user's role in an organization.
type UserAccountRole struct {
	UserID         uuid.UUID
	OrganizationID uuid.UUID
	Role           string
	CreatedAt      time.Time
}

// UserNamespacePermission represents a user's permission on a namespace.
type UserNamespacePermission struct {
	UserID      uuid.UUID
	NamespaceID string
	Permission  string
	CreatedAt   time.Time
}

// UserInvitation represents a user invitation.
type UserInvitation struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	Email          string
	Role           string
	InvitedBy      uuid.UUID
	Token          string
	ExpiresAt      time.Time
	AcceptedAt     sql.NullTime
	CreatedAt      time.Time
}

// UserRepository handles user data access.
type UserRepository struct {
	db *PostgresDB
}

// NewUserRepository creates a new user repository.
func NewUserRepository(db *PostgresDB) *UserRepository {
	return &UserRepository{db: db}
}

// Create creates a new user.
func (r *UserRepository) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (id, email, name, avatar_url, email_verified, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	_, err := r.db.DB().ExecContext(ctx, query,
		user.ID, user.Email, user.Name, user.AvatarURL, user.EmailVerified, user.CreatedAt, user.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

// GetByID retrieves a user by ID.
func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*User, error) {
	query := `
		SELECT id, email, name, avatar_url, email_verified, created_at, updated_at
		FROM users
		WHERE id = $1
	`
	user := &User{}
	err := r.db.DB().QueryRowContext(ctx, query, id).Scan(
		&user.ID, &user.Email, &user.Name, &user.AvatarURL, &user.EmailVerified, &user.CreatedAt, &user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

// GetByEmail retrieves a user by email.
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id, email, name, avatar_url, email_verified, created_at, updated_at
		FROM users
		WHERE email = $1
	`
	user := &User{}
	err := r.db.DB().QueryRowContext(ctx, query, email).Scan(
		&user.ID, &user.Email, &user.Name, &user.AvatarURL, &user.EmailVerified, &user.CreatedAt, &user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	return user, nil
}

// Update updates a user.
func (r *UserRepository) Update(ctx context.Context, user *User) error {
	query := `
		UPDATE users
		SET email = $2, name = $3, avatar_url = $4, email_verified = $5, updated_at = $6
		WHERE id = $1
	`
	user.UpdatedAt = time.Now()
	_, err := r.db.DB().ExecContext(ctx, query,
		user.ID, user.Email, user.Name, user.AvatarURL, user.EmailVerified, user.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

// Delete deletes a user.
func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.DB().ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

// CreateInvitation creates a user invitation.
func (r *UserRepository) CreateInvitation(ctx context.Context, inv *UserInvitation) error {
	query := `
		INSERT INTO user_invitations (id, organization_id, email, role, invited_by, token, expires_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (organization_id, email) DO UPDATE SET
			role = $4, invited_by = $5, token = $6, expires_at = $7, accepted_at = NULL
	`
	if inv.ID == uuid.Nil {
		inv.ID = uuid.New()
	}
	inv.CreatedAt = time.Now()

	_, err := r.db.DB().ExecContext(ctx, query,
		inv.ID, inv.OrganizationID, inv.Email, inv.Role, inv.InvitedBy, inv.Token, inv.ExpiresAt, inv.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create invitation: %w", err)
	}
	return nil
}

// GetInvitationByToken retrieves an invitation by token.
func (r *UserRepository) GetInvitationByToken(ctx context.Context, token string) (*UserInvitation, error) {
	query := `
		SELECT id, organization_id, email, role, invited_by, token, expires_at, accepted_at, created_at
		FROM user_invitations
		WHERE token = $1
	`
	inv := &UserInvitation{}
	err := r.db.DB().QueryRowContext(ctx, query, token).Scan(
		&inv.ID, &inv.OrganizationID, &inv.Email, &inv.Role, &inv.InvitedBy, &inv.Token, &inv.ExpiresAt, &inv.AcceptedAt, &inv.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get invitation: %w", err)
	}
	return inv, nil
}

// AcceptInvitation marks an invitation as accepted.
func (r *UserRepository) AcceptInvitation(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE user_invitations SET accepted_at = $2 WHERE id = $1`
	_, err := r.db.DB().ExecContext(ctx, query, id, time.Now())
	if err != nil {
		return fmt.Errorf("failed to accept invitation: %w", err)
	}
	return nil
}

// SetAccountRole sets a user's role in an organization.
func (r *UserRepository) SetAccountRole(ctx context.Context, role *UserAccountRole) error {
	query := `
		INSERT INTO user_account_roles (user_id, organization_id, role, created_at)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (user_id, organization_id) DO UPDATE SET role = $3
	`
	role.CreatedAt = time.Now()
	_, err := r.db.DB().ExecContext(ctx, query, role.UserID, role.OrganizationID, role.Role, role.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to set account role: %w", err)
	}
	return nil
}

// GetAccountRole gets a user's role in an organization.
func (r *UserRepository) GetAccountRole(ctx context.Context, userID, orgID uuid.UUID) (*UserAccountRole, error) {
	query := `
		SELECT user_id, organization_id, role, created_at
		FROM user_account_roles
		WHERE user_id = $1 AND organization_id = $2
	`
	role := &UserAccountRole{}
	err := r.db.DB().QueryRowContext(ctx, query, userID, orgID).Scan(
		&role.UserID, &role.OrganizationID, &role.Role, &role.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get account role: %w", err)
	}
	return role, nil
}

// SetNamespacePermission sets a user's permission on a namespace.
func (r *UserRepository) SetNamespacePermission(ctx context.Context, perm *UserNamespacePermission) error {
	query := `
		INSERT INTO user_namespace_permissions (user_id, namespace_id, permission, created_at)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (user_id, namespace_id) DO UPDATE SET permission = $3
	`
	perm.CreatedAt = time.Now()
	_, err := r.db.DB().ExecContext(ctx, query, perm.UserID, perm.NamespaceID, perm.Permission, perm.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to set namespace permission: %w", err)
	}
	return nil
}

// GetNamespacePermission gets a user's permission on a namespace.
func (r *UserRepository) GetNamespacePermission(ctx context.Context, userID uuid.UUID, namespaceID string) (*UserNamespacePermission, error) {
	query := `
		SELECT user_id, namespace_id, permission, created_at
		FROM user_namespace_permissions
		WHERE user_id = $1 AND namespace_id = $2
	`
	perm := &UserNamespacePermission{}
	err := r.db.DB().QueryRowContext(ctx, query, userID, namespaceID).Scan(
		&perm.UserID, &perm.NamespaceID, &perm.Permission, &perm.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get namespace permission: %w", err)
	}
	return perm, nil
}
