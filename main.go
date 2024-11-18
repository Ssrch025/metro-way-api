package main

import (
	"log"

	"github.com/Ssrch025/metro-way-api/config"
	database "github.com/Ssrch025/metro-way-api/databases"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	cfg := config.LoadConfig()
	_, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	log.Fatal(app.Listen(":" + cfg.Port))
}
