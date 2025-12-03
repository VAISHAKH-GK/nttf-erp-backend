package repository

import (
	"github.com/MagnaBit/nttf-erp-backend/internal/db/generated"
	"github.com/MagnaBit/nttf-erp-backend/internal/domain"
)

type Repositories struct {
	User domain.UserRepository
}

func NewRepositories(queries *generated.Queries) *Repositories {
	return &Repositories{
		User: newUserRepository(queries),
	}
}
