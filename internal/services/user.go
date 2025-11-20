package services

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/MagnaBit/nttf-erp-backend/internal/db/generated"
	"github.com/MagnaBit/nttf-erp-backend/internal/dto"
	"github.com/MagnaBit/nttf-erp-backend/utils"
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

func (s *UserService) Login(data dto.LoginReq, userAgent string, ip string) (string, string, error) {
	user, err := s.queries.GetUserByUsername(context.Background(), data.Username)
	if err != nil {
		return "", "", ErrInvalidCredentials
	}

	if err = utils.CompareHash(data.Password, user.Password); err != nil {
		return "", "", ErrInvalidCredentials
	}

	authToken, err := utils.GenerateJwtToken(s.jwtSecret, map[string]any{
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(2 * time.Hour).Unix(),
		"id":    user.ID,
		"email": user.Email,
	})
	if err != nil {
		return "", "", ErrTokenGeneration
	}

	sessionId, err := s.queries.InsertSession(context.Background(), generated.InsertSessionParams{UserID: user.ID, UserAgent: &userAgent, IpAddress: utils.StringToNetIpAddr(ip)})
	if err != nil {
		return "", "", ErrTokenGeneration
	}

	refreshToken := uuid.NewString()
	hashedToken := utils.HashToken(refreshToken)

	if _, err := s.queries.InsertRefreshToken(context.Background(), generated.InsertRefreshTokenParams{SessionID: sessionId, Token: &hashedToken}); err != nil {
		return "", "", ErrTokenGeneration
	}

	return authToken, refreshToken, nil
}

func (s *UserService) RefreshToken(refreshToken string) (string, string, error) {
	hashedToken := utils.HashToken(refreshToken)

	session, err := s.queries.GetRefreshTokenWithSession(context.Background(), &hashedToken)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", "", ErrInvalidRefreshToken
		}

		log.Println("DB Error: " + err.Error())
		return "", "", ErrDatabase
	}

	if session.IsRevoked == true || session.SessionExpiresAt.Time.Before(time.Now()) || session.TokenExpiresAt.Time.Before(time.Now()) {
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
	newRefreshHash := utils.HashToken(newRefresh)
	authToken, err := utils.GenerateJwtToken(s.jwtSecret, map[string]any{
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(2 * time.Hour).Unix(),
		"id":    user.ID,
		"email": user.Email,
	})

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
