package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	BotToken string
	MongoURL string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, continuing with environment variables...")
	}

	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatal("BOT_TOKEN is missing")
	}

	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		log.Fatal("MONGO_URL is missing")
	}

	return &Config{
		BotToken: botToken,
		MongoURL: mongoURL,
	}
}
