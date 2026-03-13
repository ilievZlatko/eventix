package bookings

import "time"

type Booking struct {
	ID        string    `json:"id"`
	EventID   string    `json:"event_id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BookingWithEvent struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Event     EventInfo `json:"event"`
}

type EventInfo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Location    string    `json:"location"`
	StartsAt    time.Time `json:"starts_at"`
}
