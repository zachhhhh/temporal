package service

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.temporal.io/cloud/internal/config"
	"go.temporal.io/cloud/internal/repository"
	"go.temporal.io/server/common/log"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthService struct {
	repos        *repository.Repositories
	jwtConfig    config.JWTConfig
	googleConfig *oauth2.Config
	logger       log.Logger
}

func NewAuthService(repos *repository.Repositories, jwtCfg config.JWTConfig, googleCfg *oauth2.Config, logger log.Logger) *AuthService {
	return &AuthService{repos: repos, jwtConfig: jwtCfg, googleConfig: googleCfg, logger: logger}
}

func (s *AuthService) InitiateGoogleLogin(ctx context.Context) (string, string, error) {
	state, err := generateRandomString(32)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate state: %w", err)
	}
	return s.googleConfig.AuthCodeURL(state, oauth2.AccessTypeOffline), state, nil
}

func (s *AuthService) CompleteGoogleLogin(ctx context.Context, code string) (string, string, error) {
	token, err := s.googleConfig.Exchange(ctx, code)
	if err != nil {
		return "", "", fmt.Errorf("failed to exchange code: %w", err)
	}
	userInfo, err := s.getGoogleUserInfo(ctx, token.AccessToken)
	if err != nil {
		return "", "", fmt.Errorf("failed to get user info: %w", err)
	}
	user, err := s.repos.Users.GetByEmail(ctx, userInfo.Email)
	if err != nil {
		return "", "", fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		userID := uuid.New()
		user = &repository.User{
			ID: userID, Email: userInfo.Email,
			Name: sql.NullString{String: userInfo.Name, Valid: userInfo.Name != ""},
			AvatarURL: sql.NullString{String: userInfo.AvatarURL, Valid: userInfo.AvatarURL != ""},
			EmailVerified: true,
		}
		if err := s.repos.Users.Create(ctx, user); err != nil {
			return "", "", fmt.Errorf("failed to create user: %w", err)
		}
	}
	accessToken, err := s.generateAccessToken(user.ID.String(), user.Email)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate access token: %w", err)
	}
	refreshToken, err := s.generateRefreshToken(user.ID.String())
	if err != nil {
		return "", "", fmt.Errorf("failed to generate refresh token: %w", err)
	}
	return accessToken, refreshToken, nil
}

type GoogleUserInfo struct {
	Email, Name, AvatarURL string
}

func (s *AuthService) getGoogleUserInfo(ctx context.Context, accessToken string) (*GoogleUserInfo, error) {
	req, _ := http.NewRequestWithContext(ctx, "GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get user info: %s", string(body))
	}
	var data map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&data)
	userInfo := &GoogleUserInfo{Email: data["email"].(string), Name: data["name"].(string)}
	if picture, ok := data["picture"].(string); ok {
		userInfo.AvatarURL = picture
	}
	return userInfo, nil
}

func (s *AuthService) generateAccessToken(userID, email string) (string, error) {
	claims := jwt.MapClaims{"sub": userID, "email": email, "iss": s.jwtConfig.Issuer, "aud": s.jwtConfig.Audience, "exp": time.Now().Add(time.Hour * 24).Unix(), "iat": time.Now().Unix()}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(s.jwtConfig.SecretKey))
}

func (s *AuthService) generateRefreshToken(userID string) (string, error) {
	claims := jwt.MapClaims{"sub": userID, "iss": s.jwtConfig.Issuer, "aud": s.jwtConfig.Audience, "exp": time.Now().Add(time.Hour * 24 * 30).Unix(), "iat": time.Now().Unix()}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(s.jwtConfig.SecretKey))
}

func (s *AuthService) RefreshAccessToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := s.validateToken(refreshToken)
	if err != nil {
		return "", err
	}
	userID := claims["sub"].(string)
	user, err := s.repos.Users.GetByID(ctx, uuid.MustParse(userID))
	if err != nil || user == nil {
		return "", fmt.Errorf("user not found")
	}
	return s.generateAccessToken(user.ID.String(), user.Email)
}

func (s *AuthService) validateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(s.jwtConfig.SecretKey), nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return token.Claims.(jwt.MapClaims), nil
}

func generateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func InitGoogleOAuthConfig(clientID, clientSecret, redirectURL string) *oauth2.Config {
	return &oauth2.Config{
		ClientID: clientID, ClientSecret: clientSecret,
		RedirectURL: redirectURL + "/auth/google/callback",
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}
}
