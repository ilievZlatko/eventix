package auth

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/ilievZlatko/eventix-api/internal/modules/users"
	authutil "github.com/ilievZlatko/eventix-api/internal/platform/auth"
)

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Service struct {
	usersRepo *users.Repository
	jwtSecret string
}

func NewService(usersRepo *users.Repository, jwtSecret string) *Service {
	return &Service{
		usersRepo: usersRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *Service) Register(ctx context.Context, req RegisterRequest) error {
	email := strings.TrimSpace(req.Email)
	exists, err := s.usersRepo.ExistsByEmail(ctx, email)

	if err != nil {
		return err
	}

	if exists {
		return ErrEmailAlreadyExists
	}

	passwordHash, err := authutil.HashPassword(req.Password)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	user := users.User{
		ID:           uuid.NewString(),
		Email:        email,
		PasswordHash: passwordHash,
		Role:         req.Role,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	
	return s.usersRepo.Create(ctx, user)
}

func (s *Service) Login(ctx context.Context, req LoginRequest) (string, error) {
	email := strings.TrimSpace(strings.ToLower(req.Email))

	user, err := s.usersRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	err = authutil.VerifyPassword(req.Password, user.PasswordHash)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	token, err := GenerateToken(s.jwtSecret, user.ID, user.Email, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}