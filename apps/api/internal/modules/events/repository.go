package events

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, event Event) error {
	query := `
		INSERT INTO events (
			id, title, description, location,
			starts_at, ends_at, capacity, created_by,
			created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		event.ID,
		event.Title,
		event.Description,
		event.Location,
		event.StartsAt,
		event.EndsAt,
		event.Capacity,
		event.CreatedBy,
		event.CreatedAt,
		event.UpdatedAt,
	)

	return err
}

func (r *Repository) FindAll(ctx context.Context) ([]Event, error) {
	query := `
		SELECT id, title, description, location, starts_at, ends_at,
					 capacity, created_by, created_at, updated_at
		FROM events
		ORDER BY starts_at ASC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event

		err := rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&event.Location,
			&event.StartsAt,
			&event.EndsAt,
			&event.Capacity,
			&event.CreatedBy,
			&event.CreatedAt,
			&event.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, rows.Err()
}

func (r *Repository) FindByID(ctx context.Context, id string) (*Event, error) {
	query := `
		SELECT id, title, description, location, starts_at, ends_at,
					 capacity, created_by, created_at, updated_at
		FROM events
		WHERE id = $1
	`

	var event Event

	err := r.db.QueryRow(ctx, query, id).Scan(
		&event.ID,
		&event.Title,
		&event.Description,
		&event.Location,
		&event.StartsAt,
		&event.EndsAt,
		&event.Capacity,
		&event.CreatedBy,
		&event.CreatedAt,
		&event.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &event, nil
}
