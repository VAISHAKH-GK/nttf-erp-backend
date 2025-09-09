package services

import (
	"context"

	"github.com/MagnaBit/nttf-erp-backend/internal/db/generated"
)

type UserService struct {
	queries *generated.Queries
}

func NewUserService(queries *generated.Queries) *UserService {
	return &UserService{queries: queries}
}

func (s *UserService) GetUsers(username string) (generated.User, error) {
	user, err := s.queries.GetUserByUsername(context.Background(), username)
	if err != nil {
		return user, err
	}

	return user, nil
}
