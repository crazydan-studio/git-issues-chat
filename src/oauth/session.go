package oauth

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/oauth2"
)

// Session represents an authenticated session
type Session struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

// SessionService handles session operations
type SessionService struct {
	db *sql.DB
}

// NewSessionService creates a new SessionService
func NewSessionService(db *sql.DB) *SessionService {
	return &SessionService{db: db}
}

// CreateSession creates a new session for a user
func (ss *SessionService) CreateSession(ctx context.Context, userID int64, token string, duration time.Duration) (*Session, error) {
	session := &Session{
		UserID:    userID,
		Token:     token,
		ExpiresAt: time.Now().Add(duration),
		CreatedAt: time.Now(),
	}
	
	query := `
		INSERT INTO sessions (user_id, token, expires_at, created_at)
		VALUES (?, ?, ?, ?)
		RETURNING id
	`
	
	err := ss.db.QueryRowContext(ctx, query,
		session.UserID,
		session.Token,
		session.ExpiresAt,
		session.CreatedAt,
	).Scan(&session.ID)
	
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %v", err)
	}
	
	return session, nil
}

// GetSession retrieves a session by token
func (ss *SessionService) GetSession(ctx context.Context, token string) (*Session, error) {
	session := &Session{}
	
	query := `
		SELECT id, user_id, token, expires_at, created_at
		FROM sessions
		WHERE token = ? AND expires_at > ?
	`
	
	err := ss.db.QueryRowContext(ctx, query, token, time.Now()).Scan(
		&session.ID,
		&session.UserID,
		&session.Token,
		&session.ExpiresAt,
		&session.CreatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get session: %v", err)
	}
	
	return session, nil
}

// DeleteSession deletes a session by token
func (ss *SessionService) DeleteSession(ctx context.Context, token string) error {
	query := `DELETE FROM sessions WHERE token = ?`
	
	_, err := ss.db.ExecContext(ctx, query, token)
	if err != nil {
		return fmt.Errorf("failed to delete session: %v", err)
	}
	
	return nil
}

// DeleteExpiredSessions removes expired sessions from the database
func (ss *SessionService) DeleteExpiredSessions(ctx context.Context) error {
	query := `DELETE FROM sessions WHERE expires_at <= ?`
	
	_, err := ss.db.ExecContext(ctx, query, time.Now())
	if err != nil {
		return fmt.Errorf("failed to delete expired sessions: %v", err)
	}
	
	return nil
}