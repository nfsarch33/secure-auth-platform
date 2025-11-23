package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/models"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/repository"
	"github.com/nfsarch33/secure-auth-platform/backend/pkg/jwt"
	"github.com/nfsarch33/secure-auth-platform/backend/pkg/password"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

// AuthService defines the interface for authentication logic
//
//go:generate mockgen -destination=../mocks/service/auth_service.go -package=mocks github.com/nfsarch33/secure-auth-platform/backend/internal/service AuthService
type AuthService interface {
	SignUp(ctx context.Context, email, plainPassword string) (*models.User, error)
	SignIn(ctx context.Context, email, plainPassword string) (*models.User, string, error)
}

type AuthServiceImpl struct {
	repo         repository.UserRepository
	tokenService *jwt.TokenService
}

// Ensure AuthServiceImpl implements AuthService
var _ AuthService = &AuthServiceImpl{}

func NewAuthService(repo repository.UserRepository, tokenService *jwt.TokenService) *AuthServiceImpl {
	return &AuthServiceImpl{
		repo:         repo,
		tokenService: tokenService,
	}
}

func (s *AuthServiceImpl) SignUp(ctx context.Context, email, plainPassword string) (*models.User, error) {
	// Check if user exists
	existingUser, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		// Assuming nil, nil is returned if not found, otherwise it's a DB error
	}
	if existingUser != nil {
		return nil, ErrUserAlreadyExists
	}

	hashedPassword, err := password.HashPassword(plainPassword)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: hashedPassword,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthServiceImpl) SignIn(ctx context.Context, email, plainPassword string) (*models.User, string, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, "", err
	}
	if user == nil {
		return nil, "", ErrInvalidCredentials
	}

	match, err := password.CheckPassword(plainPassword, user.PasswordHash)
	if err != nil {
		return nil, "", err
	}
	if !match {
		return nil, "", ErrInvalidCredentials
	}

	token, err := s.tokenService.GenerateToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
