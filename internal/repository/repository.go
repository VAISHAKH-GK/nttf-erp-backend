package repository

import (
	"github.com/Keracode/vidyarthidesk-backend/internal/db/generated"
	"github.com/Keracode/vidyarthidesk-backend/internal/domain"
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
