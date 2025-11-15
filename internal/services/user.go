package services

import (
	"context"

	"github.com/MagnaBit/nttf-erp-backend/internal/db/generated"
	"github.com/MagnaBit/nttf-erp-backend/utils"
)

type UserService struct {
	queries   *generated.Queries
	jwtSecret string
}

func NewUserService(queries *generated.Queries, jwtSecret string) *UserService {
	return &UserService{queries: queries, jwtSecret: jwtSecret}
}

func (s *UserService) Login(username string, password string) (string, error) {
	user, err := s.queries.GetUserByUsername(context.Background(), username)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	if err = utils.CompareHash(password, user.Password); err != nil {
		return "", ErrInvalidCredentials
	}

	token, err := utils.GenerateJwtToken(s.jwtSecret, map[string]any{
		"id":    user.ID,
		"email": user.Email,
	})
	if err != nil {
		return "", ErrTokenGeneration
	}

	return token, nil
}
