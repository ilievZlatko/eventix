package bookings

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

func (r *Repository) Create(ctx context.Context, booking Booking) error {
	query := `
		INSERT INTO bookings (id, event_id, user_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		booking.ID,
		booking.EventID,
		booking.UserID,
		booking.CreatedAt,
		booking.UpdatedAt,
	)

	return err
}

func (r *Repository) ExistsByEventAndUser(ctx context.Context, eventID, userID string) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1
			FROM bookings
			WHERE event_id = $1 AND user_id = $2
		)
	`

	var exists bool
	err := r.db.QueryRow(ctx, query, eventID, userID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *Repository) CountByEventID(ctx context.Context, eventID string) (int, error) {
	query := `
		SELECT COUNT(*)
		FROM bookings
		WHERE event_id = $1
	`

	var count int
	err := r.db.QueryRow(ctx, query, eventID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *Repository) FindByUserID(ctx context.Context, userID string) ([]Booking, error) {
	query := `
		SELECT id, event_id, user_id, created_at, updated_at
		FROM bookings
		WHERE user_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []Booking

	for rows.Next() {
		var booking Booking
	
		err := rows.Scan(
			&booking.ID,
			&booking.EventID,
			&booking.UserID,
			&booking.CreatedAt,
			&booking.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}

	return bookings, rows.Err()
}

func (r *Repository) FindByID(ctx context.Context, id string) (*Booking, error) {
	query := `
		SELECT id, event_id, user_id, created_at, updated_at
		FROM bookings
		WHERE id = $1
	`

	var booking Booking
	err := r.db.QueryRow(ctx, query, id).Scan(
		&booking.ID,
		&booking.EventID,
		&booking.UserID,
		&booking.CreatedAt,
		&booking.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &booking, nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM bookings
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query, id)
	return err
}
