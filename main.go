package main

import (
	"fmt"
	"log"

	"github.com/Ssrch025/metro-way-api/config"
	"github.com/Ssrch025/metro-way-api/database"
	"github.com/Ssrch025/metro-way-api/metro-way/handlers"
	"github.com/Ssrch025/metro-way-api/metro-way/repositories"
	"github.com/Ssrch025/metro-way-api/metro-way/usecases"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadConfig()
	db := database.Connect(cfg.DatabaseURL)

	// Init Repo and Usecase
	// Booking
	bookingRepo := repositories.NewBookingReposiotry(db)
	bookingUsecase := usecases.NewBookingUsecase(bookingRepo)

	// Start Fiber
	app := fiber.New()

	// Init Handler
	// Check Health
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	// Booking
	handlers.NewBookingHandler(app, bookingUsecase)

	// Start Server
	fmt.Printf("Server is running on port %s\n", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}
