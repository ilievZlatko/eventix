package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilievZlatko/eventix-api/internal/modules/auth"
	"github.com/ilievZlatko/eventix-api/internal/modules/users"
	"github.com/ilievZlatko/eventix-api/internal/platform/config"
	"github.com/ilievZlatko/eventix-api/internal/platform/db"
)

func main() {
	// LOAD CONFIG
	config := config.Load()
	pool, err := db.NewPool(config)

	if err != nil {
		log.Fatal(err)
	}
	
	defer pool.Close()

	// CREATE CONFIGURATIONS
	usersRepo := users.NewRepository(pool)
	authService := auth.NewService(usersRepo)
	authHandler := auth.NewHandler(authService)

	// CREATE ROUTER
	router := gin.Default()
	api := router.Group("/api")
	v1 := api.Group("/v1")

	// HEALTH CHECK ROUTE
	v1.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	// AUTH ROUTES
	authGroup := v1.Group("/auth")
	authGroup.POST("/register", authHandler.Register)


	// START SERVER
	log.Printf("API running on port: %s", config.AppPort)
	if err := router.Run(":" + config.AppPort); err != nil {
		log.Fatal(err)
	}
}
