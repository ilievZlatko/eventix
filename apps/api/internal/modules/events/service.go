package events

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrForbidden = errors.New("forbidden")
	ErrInvalidEventDates = errors.New("invalid event dates")
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, req CreateEventRequest, userID, role string) error {
	if role != "organizer" {
		return ErrForbidden
	}

	if !req.EndsAt.After(req.StartsAt) {
		return ErrInvalidEventDates
	}

	now := time.Now().UTC()

	event := Event{
		ID: 					uuid.NewString(),
		Title: 				req.Title,
		Description: 	req.Description,
		Location: 		req.Location,
		StartsAt: 		req.StartsAt,
		EndsAt: 			req.EndsAt,
		Capacity: 		req.Capacity,
		CreatedBy: 		userID,
		CreatedAt: 		now,
		UpdatedAt: 		now,
	}

	return s.repo.Create(ctx, event)
}

func (s *Service) FindAll(ctx context.Context) ([]Event, error) {
	return s.repo.FindAll(ctx)
}

func (s *Service) FindByID(ctx context.Context, id string) (*Event, error) {
	return s.repo.FindByID(ctx, id)
}
