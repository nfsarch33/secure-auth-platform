package repository

import (
	"context"
	"errors"

	"github.com/nfsarch33/secure-auth-platform/backend/internal/models"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

// UserRepository defines the interface for user persistence
//
//go:generate mockgen -destination=../mocks/repository/user_repository.go -package=mocks github.com/nfsarch33/secure-auth-platform/backend/internal/repository UserRepository
type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetByID(ctx context.Context, id string) (*models.User, error)
}
