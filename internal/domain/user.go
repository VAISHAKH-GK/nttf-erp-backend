package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Email     string
	Name      string
	Password  string
	CreatedAt time.Time
	CreatedBy uuid.UUID
	UpdatedAt time.Time
	UpdatedBy uuid.UUID
}

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (*User, error)
}
