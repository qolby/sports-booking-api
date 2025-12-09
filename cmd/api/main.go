package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/qolby/sports-booking-api/internal/config"
	"github.com/qolby/sports-booking-api/internal/database"
	"github.com/qolby/sports-booking-api/internal/handlers"
	"github.com/qolby/sports-booking-api/internal/middleware"
	"github.com/qolby/sports-booking-api/internal/services"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to database
	if err := database.Connect(cfg); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"success": false,
				"message": "An error occurred",
				"error":   err.Error(),
			})
		},
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	// Initialize services
	db := database.GetDB()
	authService := services.NewAuthService(db, cfg)
	fieldService := services.NewFieldService(db)
	bookingService := services.NewBookingService(db)
	paymentService := services.NewPaymentService(db)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	fieldHandler := handlers.NewFieldHandler(fieldService)
	bookingHandler := handlers.NewBookingHandler(bookingService)
	paymentHandler := handlers.NewPaymentHandler(paymentService)

	// Routes
	setupRoutes(app, cfg, authHandler, fieldHandler, bookingHandler, paymentHandler)

	// Start server
	port := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("Server starting on port %s", port)
	if err := app.Listen(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupRoutes(
	app *fiber.App,
	cfg *config.Config,
	authHandler *handlers.AuthHandler,
	fieldHandler *handlers.FieldHandler,
	bookingHandler *handlers.BookingHandler,
	paymentHandler *handlers.PaymentHandler,
) {
	// API v1
	api := app.Group("/api/v1")

	// Health check
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "API is running",
		})
	})

	// Auth routes (public)
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	// Field routes
	fields := api.Group("/fields")
	fields.Get("/", fieldHandler.GetAllFields)    // Public
	fields.Get("/:id", fieldHandler.GetFieldByID) // Public

	// Protected field routes (admin only)
	fields.Post("/",
		middleware.AuthRequired(cfg),
		middleware.AdminOnly(),
		fieldHandler.CreateField,
	)
	fields.Put("/:id",
		middleware.AuthRequired(cfg),
		middleware.AdminOnly(),
		fieldHandler.UpdateField,
	)
	fields.Delete("/:id",
		middleware.AuthRequired(cfg),
		middleware.AdminOnly(),
		fieldHandler.DeleteField,
	)

	// Booking routes (authenticated users)
	bookings := api.Group("/bookings", middleware.AuthRequired(cfg))
	bookings.Post("/", bookingHandler.CreateBooking)
	bookings.Get("/", bookingHandler.GetUserBookings)
	bookings.Get("/:id", bookingHandler.GetBookingByID)

	// Payment routes (authenticated users)
	payments := api.Group("/payments", middleware.AuthRequired(cfg))
	payments.Post("/", paymentHandler.ProcessPayment)
}
