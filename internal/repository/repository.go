package repository

import (
	"github.com/MagnaBit/nttf-erp-backend/internal/db/generated"
	"github.com/MagnaBit/nttf-erp-backend/internal/domain"
)

type Repositories struct {
	User         domain.UserRepository
	Session      domain.SessionRepository
	RefreshToken domain.RefreshTokenRepository
}

func NewRepositories(queries *generated.Queries) *Repositories {
	return &Repositories{
		User:         newUserRepository(queries),
		Session:      newSessionRepository(queries),
		RefreshToken: newRefreshToken(queries),
	}
}
