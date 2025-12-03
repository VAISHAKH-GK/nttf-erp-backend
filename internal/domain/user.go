package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Email     string
	Username  string
	Password  string
	CreatedAt time.Time
	CreatedBy uuid.UUID
	UpdatedAt time.Time
	UpdatedBy uuid.UUID
}

type UserRepository interface {
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (*User, error)
}
