package domain

import (
	"context"
	"net/netip"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Id         uuid.UUID
	UserId     uuid.UUID
	UserAgent  string
	IpAddress  netip.Addr
	CreatedAt  time.Time
	ExpiresAt  time.Time
	LastUsedAt time.Time
}

type RefreshToken struct {
	Id        uuid.UUID
	SessionId uuid.UUID
	Token     string
	CreatedAt time.Time
	ExpiresAt time.Time
	IsRevoked bool
	RevokedAt time.Time
}

type SessionRepository interface {
	Create(ctx context.Context, session *Session) (uuid.UUID, error)
}

type RefreshTokenRepository interface {
	Create(ctx context.Context, refreshToken *RefreshToken) (uuid.UUID, error)
	GetWithSession(ctx context.Context, hashedToken string) (RefreshToken, Session, error)
	Revoke(ctx context.Context, hashToken string) error
}
