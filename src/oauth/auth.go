package oauth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/webui-dev/go-webui/v2"
	"golang.org/x/oauth2"
)

// AuthHandler handles OAuth authentication flows
type AuthHandler struct {
	oauthManager   *OAuthManager
	userService    *UserService
	tokenService   *TokenService
	sessionService *SessionService
}

// NewAuthHandler creates a new AuthHandler
func NewAuthHandler(oauthManager *OAuthManager, userService *UserService, tokenService *TokenService, sessionService *SessionService) *AuthHandler {
	return &AuthHandler{
		oauthManager:   oauthManager,
		userService:    userService,
		tokenService:   tokenService,
		sessionService: sessionService,
	}
}

// GenerateState generates a random state string for OAuth flow
func (ah *AuthHandler) GenerateState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	
	return base64.URLEncoding.EncodeToString(b), nil
}

// HandleLogin initiates the OAuth login flow
func (ah *AuthHandler) HandleLogin(providerName string) (string, error) {
	state, err := ah.GenerateState()
	if err != nil {
		return "", fmt.Errorf("failed to generate state: %v", err)
	}
	
	authURL, err := ah.oauthManager.GetAuthURL(providerName, state)
	if err != nil {
		return "", fmt.Errorf("failed to get auth URL: %v", err)
	}
	
	return authURL, nil
}

// HandleCallback handles the OAuth callback
func (ah *AuthHandler) HandleCallback(ctx context.Context, providerName, code, state string) (*User, error) {
	// Exchange code for token
	token, err := ah.oauthManager.ExchangeCode(ctx, providerName, code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code: %v", err)
	}
	
	// Get user info
	profile, err := ah.oauthManager.GetUserInfo(ctx, providerName, token)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %v", err)
	}
	
	// Authenticate user
	user, err := ah.userService.AuthenticateUser(ctx, profile, token, providerName)
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate user: %v", err)
	}
	
	// Save token
	err = ah.tokenService.SaveToken(ctx, user.ID, token)
	if err != nil {
		return nil, fmt.Errorf("failed to save token: %v", err)
	}
	
	return user, nil
}

// CreateSessionForUser creates a session for an authenticated user
func (ah *AuthHandler) CreateSessionForUser(ctx context.Context, user *User) (string, error) {
	// Generate session token
	token, err := ah.GenerateState()
	if err != nil {
		return "", fmt.Errorf("failed to generate session token: %v", err)
	}
	
	// Create session
	_, err = ah.sessionService.CreateSession(ctx, user.ID, token, 24*time.Hour)
	if err != nil {
		return "", fmt.Errorf("failed to create session: %v", err)
	}
	
	return token, nil
}

// ValidateSession validates a session token
func (ah *AuthHandler) ValidateSession(ctx context.Context, token string) (*User, error) {
	// Get session
	session, err := ah.sessionService.GetSession(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("failed to get session: %v", err)
	}
	
	if session == nil {
		return nil, fmt.Errorf("invalid session")
	}
	
	// Get user
	// Note: This would require a GetUserByID method in UserService
	// For now, we'll return a minimal user object
	user := &User{
		ID: session.UserID,
	}
	
	return user, nil
}

// Logout logs out a user by deleting their session
func (ah *AuthHandler) Logout(ctx context.Context, token string) error {
	return ah.sessionService.DeleteSession(ctx, token)
}