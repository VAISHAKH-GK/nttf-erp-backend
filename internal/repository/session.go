package repository

import (
	"context"
	"errors"

	"github.com/MagnaBit/nttf-erp-backend/internal/db/generated"
	"github.com/MagnaBit/nttf-erp-backend/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type sessionRepository struct {
	queries *generated.Queries
}

func newSessionRepository(queries *generated.Queries) domain.SessionRepository {
	return &sessionRepository{queries: queries}
}

func (r *sessionRepository) Create(ctx context.Context, session *domain.Session) (uuid.UUID, error) {
	if sessionId, err := r.queries.InsertSession(ctx, generated.InsertSessionParams{
		UserID:    session.UserID,
		UserAgent: &session.UserAgent,
		IpAddress: &session.IpAddress,
	}); err == nil {
		session.Id = sessionId
		return sessionId, nil
	}

	return uuid.Nil, domain.ErrDatabase
}

type refreshTokenRepository struct {
	queries *generated.Queries
}

func newRefreshToken(queries *generated.Queries) domain.RefreshTokenRepository {
	return &refreshTokenRepository{queries: queries}
}

func (r *refreshTokenRepository) Create(ctx context.Context, refreshToken *domain.RefreshToken) (uuid.UUID, error) {
	if refreshTokenId, err := r.queries.InsertRefreshToken(ctx, generated.InsertRefreshTokenParams{
		SessionID: refreshToken.SessionID,
		Token:     &refreshToken.Token,
	}); err == nil {
		refreshToken.Id = refreshTokenId
		return refreshTokenId, nil
	}

	return uuid.Nil, domain.ErrDatabase
}

func (r *refreshTokenRepository) GetWithSession(ctx context.Context, hashedToken string) (domain.RefreshToken, domain.Session, error) {
	result, err := r.queries.GetRefreshTokenWithSession(ctx, &hashedToken)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.RefreshToken{}, domain.Session{}, domain.ErrInvalidRefreshToken
		}

		return domain.RefreshToken{}, domain.Session{}, domain.ErrDatabase
	}

	refreshToken := domain.RefreshToken{
		Token:     *result.Token,
		ExpiresAt: result.TokenExpiresAt.Time,
		IsRevoked: result.IsRevoked,
		SessionID: result.SessionID,
	}

	session := domain.Session{
		Id:        result.SessionID,
		UserID:    result.UserID,
		ExpiresAt: result.SessionExpiresAt.Time,
	}

	return refreshToken, session, nil
}

func (r *refreshTokenRepository) Revoke(ctx context.Context, token string) error {
	err := r.queries.RevokeRefreshToken(ctx, &token)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.ErrInvalidRefreshToken
		}

		return domain.ErrDatabase
	}

	return nil
}
