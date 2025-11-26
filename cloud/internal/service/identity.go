package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.temporal.io/cloud/internal/config"
	"go.temporal.io/cloud/internal/repository"
	"go.temporal.io/server/common/log"
)

// IdentityService handles identity and authentication business logic.
type IdentityService struct {
	repos     *repository.Repositories
	jwtConfig config.JWTConfig
	logger    log.Logger
}

// NewIdentityService creates a new identity service.
func NewIdentityService(repos *repository.Repositories, jwtCfg config.JWTConfig, logger log.Logger) *IdentityService {
	return &IdentityService{repos: repos, jwtConfig: jwtCfg, logger: logger}
}

// APIKeyInfo contains validated API key information.
type APIKeyInfo struct {
	ID          uuid.UUID
	OwnerID     uuid.UUID
	OwnerType   string
	Permissions []string
}

// ValidateAPIKey validates an API key hash and returns key info.
func (s *IdentityService) ValidateAPIKey(ctx context.Context, keyHash string) (*APIKeyInfo, error) {
	key, err := s.repos.APIKeys.GetByHash(ctx, keyHash)
	if err != nil {
		return nil, err
	}
	if key == nil {
		return nil, nil
	}

	if key.Disabled {
		return nil, nil
	}

	if key.ExpiresAt.Valid && key.ExpiresAt.Time.Before(time.Now()) {
		return nil, nil
	}

	// Parse permissions from JSON
	var permissions []string
	if key.Permissions != nil {
		_ = json.Unmarshal(key.Permissions, &permissions)
	}

	return &APIKeyInfo{
		ID:          key.ID,
		OwnerID:     key.OwnerID,
		OwnerType:   key.OwnerType,
		Permissions: permissions,
	}, nil
}

// ValidateToken validates a JWT token and returns claims.
func (s *IdentityService) ValidateToken(ctx context.Context, tokenString string) (map[string]any, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.jwtConfig.SecretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	// Validate issuer and audience
	if claims["iss"] != s.jwtConfig.Issuer {
		return nil, fmt.Errorf("invalid issuer")
	}
	if claims["aud"] != s.jwtConfig.Audience {
		return nil, fmt.Errorf("invalid audience")
	}

	return claims, nil
}

// GenerateTokens generates access and refresh tokens.
func (s *IdentityService) GenerateTokens(ctx context.Context, userID uuid.UUID, email string, orgID uuid.UUID, role string) (accessToken, refreshToken string, expiresAt time.Time, err error) {
	now := time.Now()
	expiresAt = now.Add(s.jwtConfig.AccessExpiry)

	// Access token
	accessClaims := jwt.MapClaims{
		"sub":    userID.String(),
		"email":  email,
		"org_id": orgID.String(),
		"role":   role,
		"iss":    s.jwtConfig.Issuer,
		"aud":    s.jwtConfig.Audience,
		"iat":    now.Unix(),
		"exp":    expiresAt.Unix(),
	}

	accessJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err = accessJWT.SignedString([]byte(s.jwtConfig.SecretKey))
	if err != nil {
		return "", "", time.Time{}, fmt.Errorf("failed to sign access token: %w", err)
	}

	// Refresh token
	refreshClaims := jwt.MapClaims{
		"sub":  userID.String(),
		"type": "refresh",
		"iss":  s.jwtConfig.Issuer,
		"iat":  now.Unix(),
		"exp":  now.Add(s.jwtConfig.RefreshExpiry).Unix(),
	}

	refreshJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err = refreshJWT.SignedString([]byte(s.jwtConfig.SecretKey))
	if err != nil {
		return "", "", time.Time{}, fmt.Errorf("failed to sign refresh token: %w", err)
	}

	return accessToken, refreshToken, expiresAt, nil
}

// CreateAPIKeyInput is the input for creating an API key.
type CreateAPIKeyInput struct {
	OwnerType   string
	OwnerID     uuid.UUID
	Name        string
	Permissions []string
	ExpiresAt   *time.Time
}

// CreateAPIKey creates a new API key.
func (s *IdentityService) CreateAPIKey(ctx context.Context, input *CreateAPIKeyInput) (*repository.APIKey, string, error) {
	// Generate key
	plaintext, prefix, hash, err := repository.GenerateAPIKey()
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate API key: %w", err)
	}

	permissionsJSON, _ := json.Marshal(input.Permissions)

	key := &repository.APIKey{
		OwnerType:   input.OwnerType,
		OwnerID:     input.OwnerID,
		KeyHash:     hash,
		KeyPrefix:   prefix,
		Name:        sql.NullString{String: input.Name, Valid: input.Name != ""},
		Permissions: permissionsJSON,
	}

	if input.ExpiresAt != nil {
		key.ExpiresAt = sql.NullTime{Time: *input.ExpiresAt, Valid: true}
	}

	if err := s.repos.APIKeys.Create(ctx, key); err != nil {
		return nil, "", fmt.Errorf("failed to create API key: %w", err)
	}

	return key, plaintext, nil
}

// GetAPIKey retrieves an API key by ID.
func (s *IdentityService) GetAPIKey(ctx context.Context, id uuid.UUID) (*repository.APIKey, error) {
	return s.repos.APIKeys.GetByID(ctx, id)
}

// ListAPIKeys lists API keys for an owner.
func (s *IdentityService) ListAPIKeys(ctx context.Context, ownerType string, ownerID uuid.UUID, limit, offset int) ([]*repository.APIKey, error) {
	return s.repos.APIKeys.ListByOwner(ctx, ownerType, ownerID, limit, offset)
}

// RevokeAPIKey revokes an API key.
func (s *IdentityService) RevokeAPIKey(ctx context.Context, id uuid.UUID) error {
	return s.repos.APIKeys.Disable(ctx, id)
}

// RotateAPIKey rotates an API key.
func (s *IdentityService) RotateAPIKey(ctx context.Context, id uuid.UUID) (*repository.APIKey, string, error) {
	// Get existing key
	existing, err := s.repos.APIKeys.GetByID(ctx, id)
	if err != nil {
		return nil, "", err
	}
	if existing == nil {
		return nil, "", fmt.Errorf("API key not found")
	}

	// Disable old key
	if err := s.repos.APIKeys.Disable(ctx, id); err != nil {
		return nil, "", fmt.Errorf("failed to disable old key: %w", err)
	}

	// Create new key with same settings
	var permissions []string
	_ = json.Unmarshal(existing.Permissions, &permissions)

	var expiresAt *time.Time
	if existing.ExpiresAt.Valid {
		expiresAt = &existing.ExpiresAt.Time
	}

	return s.CreateAPIKey(ctx, &CreateAPIKeyInput{
		OwnerType:   existing.OwnerType,
		OwnerID:     existing.OwnerID,
		Name:        existing.Name.String,
		Permissions: permissions,
		ExpiresAt:   expiresAt,
	})
}

// GetUser retrieves a user by ID.
func (s *IdentityService) GetUser(ctx context.Context, id uuid.UUID) (*repository.User, error) {
	return s.repos.Users.GetByID(ctx, id)
}

// GetUserByEmail retrieves a user by email.
func (s *IdentityService) GetUserByEmail(ctx context.Context, email string) (*repository.User, error) {
	return s.repos.Users.GetByEmail(ctx, email)
}

// CreateUser creates a new user.
func (s *IdentityService) CreateUser(ctx context.Context, email, name string) (*repository.User, error) {
	user := &repository.User{
		Email: email,
		Name:  sql.NullString{String: name, Valid: name != ""},
	}

	if err := s.repos.Users.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}
