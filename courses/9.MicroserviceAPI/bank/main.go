package main

import (
	"log"

	"bank/app"
	"bank/logger"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	logger.Info("Starting the application")
	app.Start()
}
