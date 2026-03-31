package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ilievZlatko/eventix-api/internal/modules/auth"
	"github.com/ilievZlatko/eventix-api/internal/modules/bookings"
	"github.com/ilievZlatko/eventix-api/internal/modules/events"
	"github.com/ilievZlatko/eventix-api/internal/modules/users"
	"github.com/ilievZlatko/eventix-api/internal/platform/config"
	"github.com/ilievZlatko/eventix-api/internal/platform/db"
	"github.com/ilievZlatko/eventix-api/internal/platform/middleware"
)

func main() {
	// LOAD CONFIG
	cfg := config.Load()
	pool, err := db.NewPool(cfg)

	if err != nil {
		log.Fatal(err)
	}
	
	defer pool.Close()

	// USERS MODULE
	usersRepo := users.NewRepository(pool)
	usersHandler := users.NewHandler()

	// AUTH MODULE
	authService := auth.NewService(usersRepo, cfg.JWTSecret)
	authHandler := auth.NewHandler(authService)

	// EVENTS MODULE
	eventsRepo := events.NewRepository(pool)
	bookingsRepo := bookings.NewRepository(pool)

	eventsService := events.NewService(eventsRepo, bookingsRepo)
	eventsHandler := events.NewHandler(eventsService)

	// BOOKINGS MODULE
	bookingsService := bookings.NewService(bookingsRepo, eventsRepo)
	bookingsHandler := bookings.NewHandler(bookingsService)
	

	// CREATE ROUTER
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: 		[]string{"http://localhost:5173"},
		AllowMethods: 		[]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: 		[]string{"Content-Type", "Authorization", "Origin", "Accept"},
		AllowCredentials: true,
		MaxAge: 					time.Hour * 12,
	}))

	api := router.Group("/api")
	v1 := api.Group("/v1")
	protected := v1.Group("/")

	// HEALTH CHECK ROUTE
	v1.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	// AUTH ROUTES
	authGroup := v1.Group("/auth")
	authGroup.POST("/register", authHandler.Register)
	authGroup.POST("/login", authHandler.Login)

	protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	protected.GET("/me", usersHandler.Me)

	// EVENTS ROUTES
	v1.GET("/events", eventsHandler.FindAll)
	v1.GET("/events/:id", eventsHandler.FindByID)

	protected.POST("/events", eventsHandler.Create)
	protected.POST("/events/:id/bookings", bookingsHandler.Create)

	// BOOKINGS ROUTES
	protected.GET("/bookings", bookingsHandler.FindMyBookings)
	protected.DELETE("/bookings/:id", bookingsHandler.Cancel)

	// START SERVER
	log.Printf("API running on port: %s", cfg.AppPort)

	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}
