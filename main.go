package main

import (
	"log"

	"github.com/Ssrch025/metro-way-api/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadConfig()
	// db := database.Connect(cfg.DatabaseURL)
	app := fiber.New()
	log.Fatal(app.Listen(":" + cfg.Port))
}
