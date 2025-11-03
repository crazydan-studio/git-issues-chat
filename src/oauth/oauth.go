package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// OAuthProvider represents an OAuth provider configuration
type OAuthProvider struct {
	Name         string
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Scopes       []string
	Endpoint     oauth2.Endpoint
}

// UserProfile represents a user's profile information
type UserProfile struct {
	ID        string `json:"id"`
	Username  string `json:"login"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
}

// OAuthManager manages OAuth authentication flows
type OAuthManager struct {
	providers map[string]*OAuthProvider
	configs   map[string]*oauth2.Config
}

// NewOAuthManager creates a new OAuthManager
func NewOAuthManager() *OAuthManager {
	return &OAuthManager{
		providers: make(map[string]*OAuthProvider),
		configs:   make(map[string]*oauth2.Config),
	}
}

// RegisterProvider registers a new OAuth provider
func (om *OAuthManager) RegisterProvider(provider *OAuthProvider) {
	om.providers[provider.Name] = provider
	
	// Create oauth2 config
	config := &oauth2.Config{
		ClientID:     provider.ClientID,
		ClientSecret: provider.ClientSecret,
		RedirectURL:  provider.RedirectURL,
		Scopes:       provider.Scopes,
		Endpoint:     provider.Endpoint,
	}
	
	om.configs[provider.Name] = config
}

// GetAuthURL generates the authorization URL for a provider
func (om *OAuthManager) GetAuthURL(providerName, state string) (string, error) {
	config, exists := om.configs[providerName]
	if !exists {
		return "", fmt.Errorf("provider %s not found", providerName)
	}
	
	return config.AuthCodeURL(state, oauth2.AccessTypeOffline), nil
}

// ExchangeCode exchanges an authorization code for an access token
func (om *OAuthManager) ExchangeCode(ctx context.Context, providerName, code string) (*oauth2.Token, error) {
	config, exists := om.configs[providerName]
	if !exists {
		return nil, fmt.Errorf("provider %s not found", providerName)
	}
	
	return config.Exchange(ctx, code)
}

// GetUserInfo fetches user information using an access token
func (om *OAuthManager) GetUserInfo(ctx context.Context, providerName string, token *oauth2.Token) (*UserProfile, error) {
	config, exists := om.configs[providerName]
	if !exists {
		return nil, fmt.Errorf("provider %s not found", providerName)
	}
	
	client := config.Client(ctx, token)
	
	// GitHub user info endpoint
	var apiURL string
	switch providerName {
	case "github":
		apiURL = "https://api.github.com/user"
	default:
		return nil, fmt.Errorf("unsupported provider: %s", providerName)
	}
	
	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user info: %v", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to fetch user info: %s, %s", resp.Status, string(body))
	}
	
	var profile UserProfile
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, fmt.Errorf("failed to decode user info: %v", err)
	}
	
	return &profile, nil
}

// DefaultProviders returns default OAuth provider configurations
func DefaultProviders() map[string]*OAuthProvider {
	return map[string]*OAuthProvider{
		"github": {
			Name:         "github",
			ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
			ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
			RedirectURL:  os.Getenv("GITHUB_REDIRECT_URL"),
			Scopes:       []string{"user:email"},
			Endpoint:     github.Endpoint,
		},
	}
}