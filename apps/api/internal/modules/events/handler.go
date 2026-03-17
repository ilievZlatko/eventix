package events

import (
	"errors"
	"fmt"
	"math"
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
	page := 1
	limit := 10

	if p := c.Query("page"); p != "" {
		fmt.Sscanf(p, "%d", &page)
	}

	if l := c.Query("limit"); l != "" {
		fmt.Sscanf(l, "%d", &limit)
	}

	if page < 1 {
		page = 1
	}
	
	if limit < 1 {
		limit = 10
	}

	if limit > 100 {
		limit = 100
	}

	offset := (page - 1) * limit
	
	events, total, err := h.service.FindAll(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get events"})
		return
	}

	totalPages := 0
	if total > 0 {
		totalPages = int(math.Ceil(float64(total) / float64(limit)))
	}

	c.JSON(http.StatusOK, gin.H{
		"data": events,
		"meta": gin.H{
			"page": page,
			"limit": limit,
			"total": total,
			"total_pages": totalPages,
		},
	})
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
