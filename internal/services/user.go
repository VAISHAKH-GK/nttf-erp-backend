package services

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/MagnaBit/nttf-erp-backend/internal/db/generated"
	"github.com/MagnaBit/nttf-erp-backend/internal/dto"
	"github.com/MagnaBit/nttf-erp-backend/pkg/hash"
	"github.com/MagnaBit/nttf-erp-backend/pkg/ip"
	"github.com/MagnaBit/nttf-erp-backend/pkg/jwt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type UserService struct {
	queries   *generated.Queries
	jwtSecret string
}

func NewUserService(queries *generated.Queries, jwtSecret string) *UserService {
	return &UserService{queries: queries, jwtSecret: jwtSecret}
}

func (s *UserService) Login(data dto.LoginReq, userAgent string, ipAddr string) (string, string, error) {
	user, err := s.queries.GetUserByUsername(context.Background(), data.Username)
	if err != nil {
		return "", "", ErrInvalidCredentials
	}

	if err = hash.CompareHash(data.Password, user.Password); err != nil {
		return "", "", ErrInvalidCredentials
	}

	authToken, err := jwt.GenerateJwtToken(s.jwtSecret, jwt.Claims{
		UserId:   user.ID,
		Email:    user.Email,
		IssuedAt: time.Now(),
		Expiry:   time.Now().Add(2 * time.Hour),
	})
	if err != nil {
		return "", "", ErrTokenGeneration
	}

	sessionId, err := s.queries.InsertSession(context.Background(), generated.InsertSessionParams{UserID: user.ID, UserAgent: &userAgent, IpAddress: ip.StringToNetIpAddr(ipAddr)})
	if err != nil {
		return "", "", ErrTokenGeneration
	}

	refreshToken := uuid.NewString()
	hashedToken := hash.HashToken(refreshToken)

	if _, err := s.queries.InsertRefreshToken(context.Background(), generated.InsertRefreshTokenParams{SessionID: sessionId, Token: &hashedToken}); err != nil {
		return "", "", ErrTokenGeneration
	}

	return authToken, refreshToken, nil
}

func (s *UserService) RefreshToken(refreshToken string) (string, string, error) {
	hashedToken := hash.HashToken(refreshToken)

	session, err := s.queries.GetRefreshTokenWithSession(context.Background(), &hashedToken)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", "", ErrInvalidRefreshToken
		}

		log.Println("DB Error: " + err.Error())
		return "", "", ErrDatabase
	}

	if session.IsRevoked || session.SessionExpiresAt.Time.Before(time.Now()) || session.TokenExpiresAt.Time.Before(time.Now()) {
		return "", "", ErrInvalidRefreshToken
	}

	user, err := s.queries.GetUserById(context.Background(), session.UserID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", "", ErrInvalidRefreshToken
		}

		log.Println("DB Error: " + err.Error())
		return "", "", ErrDatabase
	}

	newRefresh := uuid.NewString()
	newRefreshHash := hash.HashToken(newRefresh)
	authToken, err := jwt.GenerateJwtToken(s.jwtSecret, jwt.Claims{
		UserId:   user.ID,
		Email:    user.Email,
		IssuedAt: time.Now(),
		Expiry:   time.Now().Add(2 * time.Hour),
	})
	if err != nil {
		return "", "", ErrTokenGeneration
	}

	_, err = s.queries.InsertRefreshToken(context.Background(), generated.InsertRefreshTokenParams{SessionID: session.SessionID, Token: &newRefreshHash})
	if err != nil {
		return "", "", ErrTokenGeneration
	}

	err = s.queries.RevokeRefreshToken(context.Background(), &hashedToken)
	if err != nil {
		log.Println("DB Error: " + err.Error())
		return "", "", ErrDatabase
	}

	return authToken, newRefresh, nil
}
