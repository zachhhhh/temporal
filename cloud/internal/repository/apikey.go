package repository

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// APIKey represents an API key in the database.
type APIKey struct {
	ID          uuid.UUID
	OwnerType   string
	OwnerID     uuid.UUID
	KeyHash     string
	KeyPrefix   string
	Name        sql.NullString
	Permissions json.RawMessage
	Disabled    bool
	ExpiresAt   sql.NullTime
	LastUsedAt  sql.NullTime
	CreatedAt   time.Time
}

// APIKeyRepository handles API key data access.
type APIKeyRepository struct {
	db *PostgresDB
}

// NewAPIKeyRepository creates a new API key repository.
func NewAPIKeyRepository(db *PostgresDB) *APIKeyRepository {
	return &APIKeyRepository{db: db}
}

// GenerateAPIKey generates a new API key and returns the plaintext key.
func GenerateAPIKey() (plaintext string, prefix string, hash string, err error) {
	// Generate 32 random bytes
	randomBytes := make([]byte, 32)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", "", "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	// Encode as base64
	encoded := base64.URLEncoding.EncodeToString(randomBytes)

	// Create prefix (first 8 chars)
	prefix = "tc_live_" + encoded[:8]

	// Full key
	plaintext = "tc_live_" + encoded

	// Hash for storage
	hashBytes := sha256.Sum256([]byte(plaintext))
	hash = hex.EncodeToString(hashBytes[:])

	return plaintext, prefix, hash, nil
}

// Create creates a new API key.
func (r *APIKeyRepository) Create(ctx context.Context, key *APIKey) error {
	query := `
		INSERT INTO api_keys (id, owner_type, owner_id, key_hash, key_prefix, name, permissions, disabled, expires_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	if key.ID == uuid.Nil {
		key.ID = uuid.New()
	}
	if key.Permissions == nil {
		key.Permissions = json.RawMessage("[]")
	}
	key.CreatedAt = time.Now()

	_, err := r.db.DB().ExecContext(ctx, query,
		key.ID, key.OwnerType, key.OwnerID, key.KeyHash, key.KeyPrefix,
		key.Name, key.Permissions, key.Disabled, key.ExpiresAt, key.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create API key: %w", err)
	}
	return nil
}

// GetByID retrieves an API key by ID.
func (r *APIKeyRepository) GetByID(ctx context.Context, id uuid.UUID) (*APIKey, error) {
	query := `
		SELECT id, owner_type, owner_id, key_hash, key_prefix, name, permissions, disabled, expires_at, last_used_at, created_at
		FROM api_keys
		WHERE id = $1
	`
	key := &APIKey{}
	err := r.db.DB().QueryRowContext(ctx, query, id).Scan(
		&key.ID, &key.OwnerType, &key.OwnerID, &key.KeyHash, &key.KeyPrefix,
		&key.Name, &key.Permissions, &key.Disabled, &key.ExpiresAt, &key.LastUsedAt, &key.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get API key: %w", err)
	}
	return key, nil
}

// GetByHash retrieves an API key by hash.
func (r *APIKeyRepository) GetByHash(ctx context.Context, hash string) (*APIKey, error) {
	query := `
		SELECT id, owner_type, owner_id, key_hash, key_prefix, name, permissions, disabled, expires_at, last_used_at, created_at
		FROM api_keys
		WHERE key_hash = $1
	`
	key := &APIKey{}
	err := r.db.DB().QueryRowContext(ctx, query, hash).Scan(
		&key.ID, &key.OwnerType, &key.OwnerID, &key.KeyHash, &key.KeyPrefix,
		&key.Name, &key.Permissions, &key.Disabled, &key.ExpiresAt, &key.LastUsedAt, &key.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get API key by hash: %w", err)
	}
	return key, nil
}

// ValidateKey validates an API key and returns the key record if valid.
func (r *APIKeyRepository) ValidateKey(ctx context.Context, plaintext string) (*APIKey, error) {
	// Hash the plaintext key
	hashBytes := sha256.Sum256([]byte(plaintext))
	hash := hex.EncodeToString(hashBytes[:])

	key, err := r.GetByHash(ctx, hash)
	if err != nil {
		return nil, err
	}
	if key == nil {
		return nil, nil
	}

	// Check if disabled
	if key.Disabled {
		return nil, nil
	}

	// Check if expired
	if key.ExpiresAt.Valid && key.ExpiresAt.Time.Before(time.Now()) {
		return nil, nil
	}

	// Update last used
	_ = r.UpdateLastUsed(ctx, key.ID)

	return key, nil
}

// UpdateLastUsed updates the last used timestamp.
func (r *APIKeyRepository) UpdateLastUsed(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE api_keys SET last_used_at = $2 WHERE id = $1`
	_, err := r.db.DB().ExecContext(ctx, query, id, time.Now())
	if err != nil {
		return fmt.Errorf("failed to update last used: %w", err)
	}
	return nil
}

// Disable disables an API key.
func (r *APIKeyRepository) Disable(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE api_keys SET disabled = true WHERE id = $1`
	_, err := r.db.DB().ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to disable API key: %w", err)
	}
	return nil
}

// Delete deletes an API key.
func (r *APIKeyRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM api_keys WHERE id = $1`
	_, err := r.db.DB().ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete API key: %w", err)
	}
	return nil
}

// ListByOwner lists API keys for an owner.
func (r *APIKeyRepository) ListByOwner(ctx context.Context, ownerType string, ownerID uuid.UUID, limit, offset int) ([]*APIKey, error) {
	query := `
		SELECT id, owner_type, owner_id, key_hash, key_prefix, name, permissions, disabled, expires_at, last_used_at, created_at
		FROM api_keys
		WHERE owner_type = $1 AND owner_id = $2
		ORDER BY created_at DESC
		LIMIT $3 OFFSET $4
	`
	rows, err := r.db.DB().QueryContext(ctx, query, ownerType, ownerID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list API keys: %w", err)
	}
	defer rows.Close()

	var keys []*APIKey
	for rows.Next() {
		key := &APIKey{}
		if err := rows.Scan(
			&key.ID, &key.OwnerType, &key.OwnerID, &key.KeyHash, &key.KeyPrefix,
			&key.Name, &key.Permissions, &key.Disabled, &key.ExpiresAt, &key.LastUsedAt, &key.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan API key: %w", err)
		}
		keys = append(keys, key)
	}
	return keys, nil
}
