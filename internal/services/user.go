package services

import (
	"context"

	"github.com/MagnaBit/nttf-erp-backend/internal/db/generated"
	"github.com/MagnaBit/nttf-erp-backend/internal/dto"
	"github.com/MagnaBit/nttf-erp-backend/utils"
	"github.com/google/uuid"
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
	hashedToken, err := utils.HashPassword(refreshToken)
	if err != nil {
		return "", "", ErrTokenGeneration
	}

	if _, err := s.queries.InsertRefreshToken(context.Background(), generated.InsertRefreshTokenParams{SessionID: sessionId, Token: &hashedToken}); err == nil {
		return "", "", ErrTokenGeneration
	}

	return authToken, refreshToken, nil
}
