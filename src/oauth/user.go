package oauth

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/oauth2"
)

// User represents a user in the system
type User struct {
	ID           int64          `json:"id"`
	Provider     string         `json:"provider"`
	ProviderID   string         `json:"provider_id"`
	Username     string         `json:"username"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	AvatarURL    string         `json:"avatar_url"`
	AccessToken  string         `json:"access_token"`
	RefreshToken string         `json:"refresh_token"`
	Expiry       time.Time      `json:"expiry"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

// UserService handles user-related operations
type UserService struct {
	db *sql.DB
}

// NewUserService creates a new UserService
func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

// CreateUser creates a new user in the database
func (us *UserService) CreateUser(ctx context.Context, profile *UserProfile, token *oauth2.Token, provider string) (*User, error) {
	user := &User{
		Provider:     provider,
		ProviderID:   profile.ID,
		Username:     profile.Username,
		Name:         profile.Name,
		Email:        profile.Email,
		AvatarURL:    profile.AvatarURL,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	
	query := `
		INSERT INTO users (provider, provider_id, username, name, email, avatar_url, access_token, refresh_token, expiry, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		RETURNING id
	`
	
	err := us.db.QueryRowContext(ctx, query,
		user.Provider,
		user.ProviderID,
		user.Username,
		user.Name,
		user.Email,
		user.AvatarURL,
		user.AccessToken,
		user.RefreshToken,
		user.Expiry,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.ID)
	
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}
	
	return user, nil
}

// GetUserByProviderID retrieves a user by provider and provider ID
func (us *UserService) GetUserByProviderID(ctx context.Context, provider, providerID string) (*User, error) {
	user := &User{}
	
	query := `
		SELECT id, provider, provider_id, username, name, email, avatar_url, access_token, refresh_token, expiry, created_at, updated_at
		FROM users
		WHERE provider = ? AND provider_id = ?
	`
	
	err := us.db.QueryRowContext(ctx, query, provider, providerID).Scan(
		&user.ID,
		&user.Provider,
		&user.ProviderID,
		&user.Username,
		&user.Name,
		&user.Email,
		&user.AvatarURL,
		&user.AccessToken,
		&user.RefreshToken,
		&user.Expiry,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user: %v", err)
	}
	
	return user, nil
}

// UpdateUserToken updates a user's OAuth tokens
func (us *UserService) UpdateUserToken(ctx context.Context, userID int64, token *oauth2.Token) error {
	query := `
		UPDATE users
		SET access_token = ?, refresh_token = ?, expiry = ?, updated_at = ?
		WHERE id = ?
	`
	
	_, err := us.db.ExecContext(ctx, query,
		token.AccessToken,
		token.RefreshToken,
		token.Expiry,
		time.Now(),
		userID,
	)
	
	if err != nil {
		return fmt.Errorf("failed to update user token: %v", err)
	}
	
	return nil
}

// AuthenticateUser authenticates a user with OAuth
func (us *UserService) AuthenticateUser(ctx context.Context, profile *UserProfile, token *oauth2.Token, provider string) (*User, error) {
	// Check if user already exists
	user, err := us.GetUserByProviderID(ctx, provider, profile.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to check user existence: %v", err)
	}
	
	// If user doesn't exist, create new user
	if user == nil {
		user, err = us.CreateUser(ctx, profile, token, provider)
		if err != nil {
			return nil, fmt.Errorf("failed to create user: %v", err)
		}
	} else {
		// If user exists, update tokens
		err = us.UpdateUserToken(ctx, user.ID, token)
		if err != nil {
			return nil, fmt.Errorf("failed to update user token: %v", err)
		}
		
		// Update user info
		user.Username = profile.Username
		user.Name = profile.Name
		user.Email = profile.Email
		user.AvatarURL = profile.AvatarURL
	}
	
	return user, nil
}