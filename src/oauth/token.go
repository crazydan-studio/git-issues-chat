package oauth

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/oauth2"
)

// TokenService handles OAuth token operations
type TokenService struct {
	db *sql.DB
}

// NewTokenService creates a new TokenService
func NewTokenService(db *sql.DB) *TokenService {
	return &TokenService{db: db}
}

// SaveToken saves an OAuth token to the database
func (ts *TokenService) SaveToken(ctx context.Context, userID int64, token *oauth2.Token) error {
	query := `
		INSERT INTO oauth_tokens (user_id, access_token, refresh_token, expiry, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
		ON CONFLICT(user_id) DO UPDATE SET
		access_token = excluded.access_token,
		refresh_token = excluded.refresh_token,
		expiry = excluded.expiry,
		updated_at = excluded.updated_at
	`
	
	_, err := ts.db.ExecContext(ctx, query,
		userID,
		token.AccessToken,
		token.RefreshToken,
		token.Expiry,
		time.Now(),
		time.Now(),
	)
	
	if err != nil {
		return fmt.Errorf("failed to save token: %v", err)
	}
	
	return nil
}

// GetToken retrieves an OAuth token for a user
func (ts *TokenService) GetToken(ctx context.Context, userID int64) (*oauth2.Token, error) {
	token := &oauth2.Token{}
	
	query := `
		SELECT access_token, refresh_token, expiry
		FROM oauth_tokens
		WHERE user_id = ?
	`
	
	err := ts.db.QueryRowContext(ctx, query, userID).Scan(
		&token.AccessToken,
		&token.RefreshToken,
		&token.Expiry,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get token: %v", err)
	}
	
	return token, nil
}

// DeleteToken deletes an OAuth token for a user
func (ts *TokenService) DeleteToken(ctx context.Context, userID int64) error {
	query := `DELETE FROM oauth_tokens WHERE user_id = ?`
	
	_, err := ts.db.ExecContext(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("failed to delete token: %v", err)
	}
	
	return nil
}