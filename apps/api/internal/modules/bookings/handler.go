package bookings

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(c *gin.Context) {
	eventID := c.Param("id")
	userID := c.GetString("user_id")

	err := h.service.Create(c.Request.Context(), eventID, userID)
	if err != nil {
		switch {
		case errors.Is(err, ErrBookingAlreadyExists):
			c.JSON(http.StatusConflict, gin.H{"error": "Booking already exists"})
		case errors.Is(err, ErrEventFull):
			c.JSON(http.StatusBadRequest, gin.H{"error": "Event is full"})
		case errors.Is(err, pgx.ErrNoRows):
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Booking created successfully"})
}

func (h *Handler) FindMyBookings(c *gin.Context) {
	userID := c.GetString("user_id")

	bookings, err := h.service.FindByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get bookings"})
		return
	}
	c.JSON(http.StatusOK, bookings)
}

func (h *Handler) Cancel(c *gin.Context) {
	bookingID := c.Param("id")
	userID := c.GetString("user_id")

	err := h.service.Cancel(c.Request.Context(), bookingID, userID)
	if err != nil {
		switch {
		case errors.Is(err, ErrBookingNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		case errors.Is(err, ErrForbidden):
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel booking"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking cancelled successfully"})
}
