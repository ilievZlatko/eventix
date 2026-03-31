package events

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrForbidden       = errors.New("forbidden")
	ErrInvalidEventDates = errors.New("invalid event dates")
)

type Service struct {
	repo          *Repository
	bookingReader BookingReader
}

func NewService(repo *Repository, bookingReader BookingReader) *Service {
	return &Service{
		repo:          repo,
		bookingReader: bookingReader,
	}
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

func (s *Service) FindAll(ctx context.Context, limit, offset int) ([]Event, int, error) {
	events, err := s.repo.FindAll(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	return events, total, nil
}

func (s *Service) FindByID(ctx context.Context, id, userID string) (*EventDetailsResponse, error) {
	event, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	bookedCount, err := s.bookingReader.CountByEventID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	isBooked := false
	if userID != "" {
		isBooked, err = s.bookingReader.ExistsByEventAndUser(ctx, id, userID)
		if err != nil {
			return nil, err
		}
	}

	return &EventDetailsResponse{
		ID: 					event.ID,
		Title: 				event.Title,
		Description: 	event.Description,
		Location: 		event.Location,
		StartsAt: 		event.StartsAt,
		EndsAt: 			event.EndsAt,
		Capacity: 		event.Capacity,
		BookedCount: 	bookedCount,
		IsBooked: 		isBooked,
		CreatedBy: 		event.CreatedBy,
		CreatedAt: 		event.CreatedAt,
		UpdatedAt: 		event.UpdatedAt,
	}, nil
}
