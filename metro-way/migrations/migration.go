package main

import (
	"fmt"
	"log"

	"github.com/Ssrch025/metro-way-api/config"
	"github.com/Ssrch025/metro-way-api/database"
	"github.com/Ssrch025/metro-way-api/metro-way/domain"
)

func main() {
	cfg := config.LoadConfig()
	db := database.Connect(cfg.DatabaseURL)
	err := db.AutoMigrate(&domain.Booking{})
	if err != nil {
		log.Fatal(err)
		fmt.Println("Failed to migrate to database")
	}
	fmt.Println("Migrate to database successfully")
}
