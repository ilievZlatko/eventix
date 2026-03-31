package events

import "time"

type CreateEventRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	StartsAt    time.Time `json:"starts_at" binding:"required"`
	EndsAt      time.Time `json:"ends_at" binding:"required"`
	Capacity    int       `json:"capacity" binding:"required"`
}

type EventDetailsResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	StartsAt    time.Time `json:"starts_at"`
	EndsAt      time.Time `json:"ends_at"`
	Capacity    int       `json:"capacity"`
	BookedCount int       `json:"booked_count"`
	IsBooked    bool      `json:"is_booked"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
