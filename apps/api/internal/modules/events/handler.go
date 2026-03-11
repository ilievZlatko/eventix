package events

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
	var req CreateEventRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	err := h.service.Create(
		c.Request.Context(),
		req,
		c.GetString("user_id"),
		c.GetString("role"),
	)

	if err != nil {
		switch {
		case errors.Is(err, ErrForbidden):
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		case errors.Is(err, ErrInvalidEventDates):
			c.JSON(http.StatusBadRequest, gin.H{"error": "edns_at must be after starts_at"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create an event"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
	})
}

func (h *Handler) FindAll(c *gin.Context) {
	events, err := h.service.FindAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get events"})
		return
	}

	c.JSON(http.StatusOK, events)
}

func (h *Handler) FindByID(c *gin.Context) {
	event, err := h.service.FindByID(c.Request.Context(), c.Param("id"))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get event"})
		return
	}

	c.JSON(http.StatusOK, event)
}
