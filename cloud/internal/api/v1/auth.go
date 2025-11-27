package api

import (
	"encoding/json"
	"net/http"

	"go.temporal.io/cloud/internal/service"
)

// AuthHandler handles authentication HTTP endpoints.
type AuthHandler struct {
	authService *service.AuthService
}

// NewAuthHandler creates a new auth handler.
func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// HandleGoogleLogin initiates Google OAuth login.
func (h *AuthHandler) HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	authURL, state, err := h.authService.InitiateGoogleLogin(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Store state in cookie for CSRF protection
	http.SetCookie(w, &http.Cookie{
		Name:     "oauth_state",
		Value:    state,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   600, // 10 minutes
	})

	// Return auth URL
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"auth_url": authURL,
	})
}

// HandleGoogleCallback handles the Google OAuth callback.
func (h *AuthHandler) HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	// Verify state for CSRF protection
	stateCookie, err := r.Cookie("oauth_state")
	if err != nil {
		http.Error(w, "missing state cookie", http.StatusBadRequest)
		return
	}

	state := r.URL.Query().Get("state")
	if state != stateCookie.Value {
		http.Error(w, "invalid state", http.StatusBadRequest)
		return
	}

	// Get authorization code
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "missing code", http.StatusBadRequest)
		return
	}

	// Complete OAuth flow
	accessToken, refreshToken, err := h.authService.CompleteGoogleLogin(r.Context(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Clear state cookie
	http.SetCookie(w, &http.Cookie{
		Name:   "oauth_state",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	// Set auth tokens in cookies
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   86400, // 24 hours
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   2592000, // 30 days
	})

	// Redirect to console frontend
	http.Redirect(w, r, "http://localhost:5173/console/namespaces", http.StatusFound)
}

// HandleRefreshToken refreshes an access token.
func (h *AuthHandler) HandleRefreshToken(w http.ResponseWriter, r *http.Request) {
	// Get refresh token from cookie
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		http.Error(w, "missing refresh token", http.StatusUnauthorized)
		return
	}

	// Refresh access token
	accessToken, err := h.authService.RefreshAccessToken(r.Context(), cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Update access token cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   86400, // 24 hours
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
	})
}

// HandleLogout logs out the user.
func (h *AuthHandler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	// Clear auth cookies
	http.SetCookie(w, &http.Cookie{
		Name:   "access_token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	http.SetCookie(w, &http.Cookie{
		Name:   "refresh_token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
	})
}
