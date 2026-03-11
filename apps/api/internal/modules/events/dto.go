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
