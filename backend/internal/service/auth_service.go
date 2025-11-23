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
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type AuthService struct {
	repo         repository.UserRepository
	tokenService *jwt.TokenService
}

func NewAuthService(repo repository.UserRepository, tokenService *jwt.TokenService) *AuthService {
	return &AuthService{
		repo:         repo,
		tokenService: tokenService,
	}
}

func (s *AuthService) SignUp(ctx context.Context, email, plainPassword string) (*models.User, error) {
	// Check if user exists
	existingUser, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		// Assuming nil, nil is returned if not found, otherwise it's a DB error
		// Ideally we handle errors better, but for now proceed if user found
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

func (s *AuthService) SignIn(ctx context.Context, email, plainPassword string) (*models.User, string, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, "", err // Or wrap error
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
