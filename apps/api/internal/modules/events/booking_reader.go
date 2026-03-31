package events

import "context"

type BookingReader interface {
	CountByEventID(ctx context.Context, eventID string) (int, error)
	ExistsByEventAndUser(ctx context.Context, eventID, userID string) (bool, error)
}
