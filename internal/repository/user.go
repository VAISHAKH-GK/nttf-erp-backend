package repository

import (
	"context"
	"errors"

	"github.com/Keracode/vidyarthidesk-backend/internal/db/generated"
	"github.com/Keracode/vidyarthidesk-backend/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type userRepository struct {
	queries *generated.Queries
}

func newUserRepository(queries *generated.Queries) domain.UserRepository {
	return &userRepository{queries: queries}
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}

		return nil, domain.ErrDatabase
	}

	return toDomain(user), nil
}

func (r *userRepository) GetUserById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user, err := r.queries.GetUserById(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}

		return nil, domain.ErrDatabase
	}

	return toDomain(user), nil
}

func toDomain(user generated.User) *domain.User {
	return &domain.User{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Password:  user.Password,
		CreatedAt: user.CreatedAt.Time,
		CreatedBy: user.CreatedBy,
		UpdatedAt: user.UpdatedAt.Time,
		UpdatedBy: user.UpdatedBy,
	}
}
