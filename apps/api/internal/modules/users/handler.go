package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Me(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id": c.GetString("user_id"),
		"email": c.GetString("email"),
		"role": c.GetString("role"),
	})
}
