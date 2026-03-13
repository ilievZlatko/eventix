package bookings

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/ilievZlatko/eventix-api/internal/modules/events"
	"github.com/jackc/pgx/v5"
)

var (
	ErrBookingAlreadyExists = errors.New("booking already exists")
	ErrEventFull = errors.New("event is full")
	ErrBookingNotFound = errors.New("booking not found")
	ErrForbidden = errors.New("forbidden")
)

type Service struct {
	repo 			 *Repository
	eventsRepo *events.Repository
}

func NewService(repo *Repository, eventsRepo *events.Repository) *Service {
	return &Service{
		repo: 			repo,
		eventsRepo: eventsRepo,
	}
}

func (s *Service) Create(ctx context.Context, eventID, userID string) error {
	_, err := s.eventsRepo.FindByID(ctx, eventID)
	if err != nil {
		return err
	}

	exists, err := s.repo.ExistsByEventAndUser(ctx, eventID, userID)
	if err != nil {
		return err
	}
	if exists {
		return ErrBookingAlreadyExists
	}

	event, err := s.eventsRepo.FindByID(ctx, eventID)
	if err != nil {
		return err
	}

	count, err := s.repo.CountByEventID(ctx, eventID)
	if err != nil {
		return err
	}

	if count >= event.Capacity {
		return ErrEventFull
	}

	booking := Booking{
		ID: uuid.NewString(),
		EventID: eventID,
		UserID: userID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	return s.repo.Create(ctx, booking)
}

func (s *Service) FindByUserID(ctx context.Context, userID string) ([]BookingWithEvent, error) {
	return s.repo.FindByUserID(ctx, userID)
}

func (s *Service) Cancel(ctx context.Context, bookingID, userID string) error {
	booking, err := s.repo.FindByID(ctx, bookingID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrBookingNotFound
		}
		return err
	}

	if booking.UserID != userID {
		return ErrForbidden
	}
	
	return s.repo.Delete(ctx, bookingID)
}
