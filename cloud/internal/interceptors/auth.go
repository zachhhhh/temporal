// Package interceptors provides gRPC interceptors for the Cloud API.
package interceptors

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"connectrpc.com/connect"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.temporal.io/cloud/internal/service"
	"go.temporal.io/server/common/log"
	"go.temporal.io/server/common/log/tag"
)

// AuthContext keys for storing auth info in context.
type authContextKey struct{}

// AuthInfo contains authenticated user/service account information.
type AuthInfo struct {
	UserID         uuid.UUID
	Email          string
	OrganizationID uuid.UUID
	Role           string
	Permissions    []string
	IsAPIKey       bool
	APIKeyID       uuid.UUID
}

// AuthInterceptor handles authentication for gRPC requests.
type AuthInterceptor struct {
	identityService *service.IdentityService
	logger          log.Logger
}

// NewAuthInterceptor creates a new auth interceptor.
func NewAuthInterceptor(identityService *service.IdentityService, logger log.Logger) *AuthInterceptor {
	return &AuthInterceptor{
		identityService: identityService,
		logger:          logger,
	}
}

// WrapUnary implements connect.Interceptor.
func (i *AuthInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		// Skip auth for health checks
		if strings.HasPrefix(req.Spec().Procedure, "/grpc.health") {
			return next(ctx, req)
		}

		// Get authorization header
		authHeader := req.Header().Get("Authorization")
		if authHeader == "" {
			return nil, connect.NewError(connect.CodeUnauthenticated, nil)
		}

		// Parse token
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return nil, connect.NewError(connect.CodeUnauthenticated, nil)
		}
		token := parts[1]

		var authInfo *AuthInfo
		var err error

		// Check if it's an API key (starts with tc_)
		if strings.HasPrefix(token, "tc_") {
			authInfo, err = i.validateAPIKey(ctx, token)
		} else {
			authInfo, err = i.validateJWT(ctx, token)
		}

		if err != nil {
			i.logger.Warn("Authentication failed", tag.Error(err))
			return nil, connect.NewError(connect.CodeUnauthenticated, nil)
		}

		// Add auth info to context
		ctx = context.WithValue(ctx, authContextKey{}, authInfo)

		return next(ctx, req)
	}
}

// WrapStreamingClient implements connect.Interceptor.
func (i *AuthInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return next
}

// WrapStreamingHandler implements connect.Interceptor.
func (i *AuthInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return next
}

func (i *AuthInterceptor) validateAPIKey(ctx context.Context, key string) (*AuthInfo, error) {
	// Hash the key
	hashBytes := sha256.Sum256([]byte(key))
	hash := hex.EncodeToString(hashBytes[:])

	// Validate with identity service
	apiKey, err := i.identityService.ValidateAPIKey(ctx, hash)
	if err != nil {
		return nil, err
	}
	if apiKey == nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, nil)
	}

	return &AuthInfo{
		UserID:      apiKey.OwnerID,
		IsAPIKey:    true,
		APIKeyID:    apiKey.ID,
		Permissions: apiKey.Permissions,
	}, nil
}

func (i *AuthInterceptor) validateJWT(ctx context.Context, tokenString string) (*AuthInfo, error) {
	claims, err := i.identityService.ValidateToken(ctx, tokenString)
	if err != nil {
		return nil, err
	}

	userID, _ := uuid.Parse(claims["sub"].(string))
	orgID, _ := uuid.Parse(claims["org_id"].(string))

	return &AuthInfo{
		UserID:         userID,
		Email:          claims["email"].(string),
		OrganizationID: orgID,
		Role:           claims["role"].(string),
		IsAPIKey:       false,
	}, nil
}

// GetAuthInfo retrieves auth info from context.
func GetAuthInfo(ctx context.Context) *AuthInfo {
	if info, ok := ctx.Value(authContextKey{}).(*AuthInfo); ok {
		return info
	}
	return nil
}

// JWTClaims represents JWT claims.
type JWTClaims struct {
	jwt.RegisteredClaims
	Email          string `json:"email"`
	OrganizationID string `json:"org_id"`
	Role           string `json:"role"`
}
