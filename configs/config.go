package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
}

var config *Config

// NewConfig is a constructor for Config
func new() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	config = &Config{}
	config.Database.Host = os.Getenv("DATABASE_HOST")
	config.Database.Port = os.Getenv("DATABASE_PORT")
	config.Database.User = os.Getenv("DATABASE_USERNAME")
	config.Database.Password = os.Getenv("DATABASE_PASSWORD")
	config.Database.Name = os.Getenv("DATABASE_NAME")

	return config
}

func GetConfig() *Config {
	if config == nil {
		config = new()
	}

	return config
}
