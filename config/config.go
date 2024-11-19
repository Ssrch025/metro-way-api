package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	return &Config{
		Port:        os.Getenv("PORT"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}
