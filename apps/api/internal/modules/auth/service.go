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

var ErrEmailAlreadyExists = errors.New("email already exists")

type Service struct {
	usersRepo *users.Repository
}

func NewService(usersRepo *users.Repository) *Service {
	return &Service{usersRepo: usersRepo}
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