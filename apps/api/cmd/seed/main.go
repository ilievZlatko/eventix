package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/ilievZlatko/eventix-api/internal/modules/bookings"
	"github.com/ilievZlatko/eventix-api/internal/modules/events"
	"github.com/ilievZlatko/eventix-api/internal/modules/users"
	"github.com/ilievZlatko/eventix-api/internal/platform/auth"
	"github.com/ilievZlatko/eventix-api/internal/platform/config"
	"github.com/ilievZlatko/eventix-api/internal/platform/db"
)

func main() {
	ctx := context.Background()

	cfg := config.Load()
	pool, err := db.NewPool(cfg)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer pool.Close()

	usersRepo := users.NewRepository(pool)
	eventsRepo := events.NewRepository(pool)
	bookingsRepo := bookings.NewRepository(pool)

	log.Println("Seeding users...")
	userIDs, err := seedUsers(ctx, usersRepo)
	if err != nil {
		log.Fatalf("failed to seed users: %v", err)
	}

	log.Println("Seeding events...")
	eventIDs, err := seedEvents(ctx, eventsRepo, userIDs)
	if err != nil {
		log.Fatalf("failed to seed events: %v", err)
	}

	log.Println("Seeding bookings...")
	if err := seedBookings(ctx, bookingsRepo, eventIDs, userIDs); err != nil {
		log.Fatalf("failed to seed bookings: %v", err)
	}

	log.Println("Seeding completed successfully")
}

const (
	organizer1Email = "organizer1@example.com"
	organizer2Email = "organizer2@example.com"
	attendee1Email  = "attendee1@example.com"
	attendee2Email  = "attendee2@example.com"
	goConfTitle     = "Go Conference 2026"
)

func seedUsers(ctx context.Context, repo *users.Repository) (map[string]string, error) {
	now := time.Now().UTC()

	type userSeed struct {
		ID    string
		Email string
		Role  string
	}

	seeds := []userSeed{
		{ID: uuid.NewString(), Email: organizer1Email, Role: "organizer"},
		{ID: uuid.NewString(), Email: organizer2Email, Role: "organizer"},
		{ID: uuid.NewString(), Email: attendee1Email, Role: "user"},
		{ID: uuid.NewString(), Email: attendee2Email, Role: "user"},
	}

	passwordHash, err := auth.HashPassword("Password123!")
	if err != nil {
		return nil, err
	}

	idsByEmail := make(map[string]string)

	for _, u := range seeds {
		user := users.User{
			ID:           u.ID,
			Email:        u.Email,
			PasswordHash: passwordHash,
			Role:         u.Role,
			CreatedAt:    now,
			UpdatedAt:    now,
		}

		if err := repo.Create(ctx, user); err != nil {
			return nil, err
		}

		idsByEmail[u.Email] = u.ID
		log.Printf("Created user %s (%s)", u.Email, u.Role)
	}

	return idsByEmail, nil
}

func seedEvents(ctx context.Context, repo *events.Repository, userIDs map[string]string) (map[string]string, error) {
	now := time.Now().UTC()

	type eventSeed struct {
		Title       string
		Description string
		Location    string
		StartsInH   int
		DurationH   int
		Capacity    int
		Organizer   string
	}

	baseSeeds := []eventSeed{
		{
			Title:       goConfTitle,
			Description: "A conference about Go and backend development.",
			Location:    "Berlin",
			StartsInH:   24,
			DurationH:   8,
			Capacity:    100,
			Organizer:   organizer1Email,
		},
		{
			Title:       "Frontend Meetup",
			Description: "Monthly meetup for frontend engineers.",
			Location:    "London",
			StartsInH:   48,
			DurationH:   3,
			Capacity:    50,
			Organizer:   organizer1Email,
		},
		{
			Title:       "DevOps Workshop",
			Description: "Hands-on workshop on Kubernetes and CI/CD.",
			Location:    "Online",
			StartsInH:   72,
			DurationH:   4,
			Capacity:    30,
			Organizer:   organizer2Email,
		},
	}

	eventIDs := make(map[string]string)

	// create the base events
	for _, e := range baseSeeds {
		startsAt := now.Add(time.Duration(e.StartsInH) * time.Hour)
		endsAt := startsAt.Add(time.Duration(e.DurationH) * time.Hour)

		event := events.Event{
			ID:          uuid.NewString(),
			Title:       e.Title,
			Description: e.Description,
			Location:    e.Location,
			StartsAt:    startsAt,
			EndsAt:      endsAt,
			Capacity:    e.Capacity,
			CreatedBy:   userIDs[e.Organizer],
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		if err := repo.Create(ctx, event); err != nil {
			return nil, err
		}

		eventIDs[e.Title] = event.ID
		log.Printf("Created event %s", e.Title)
	}

	// create additional events for pagination testing
	for i := 1; i <= 25; i++ {
		title := fmt.Sprintf("Sample Event %02d", i)

		startsAt := now.Add(time.Duration(24+i) * time.Hour)
		endsAt := startsAt.Add(3 * time.Hour)

		event := events.Event{
			ID:          uuid.NewString(),
			Title:       title,
			Description: fmt.Sprintf("Sample seeded event number %d", i),
			Location:    "Online",
			StartsAt:    startsAt,
			EndsAt:      endsAt,
			Capacity:    50,
			CreatedBy:   userIDs[organizer1Email],
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		if err := repo.Create(ctx, event); err != nil {
			return nil, err
		}

		eventIDs[title] = event.ID
		log.Printf("Created event %s", title)
	}

	return eventIDs, nil
}

func seedBookings(
	ctx context.Context,
	repo *bookings.Repository,
	eventIDs map[string]string,
	userIDs map[string]string,
) error {
	now := time.Now().UTC()

	// create bookings for every event for both attendees
	for title, eventID := range eventIDs {
		for _, attendeeEmail := range []string{attendee1Email, attendee2Email} {
			booking := bookings.Booking{
				ID:        uuid.NewString(),
				EventID:   eventID,
				UserID:    userIDs[attendeeEmail],
				CreatedAt: now,
			}

			if err := repo.Create(ctx, booking); err != nil {
				return err
			}

			log.Printf("Created booking for %s -> %s", attendeeEmail, title)
		}
	}

	return nil
}

