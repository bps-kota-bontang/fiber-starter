package main

import (
	"fmt"

	"log"
	"os"

	"github.com/bps-kota-bontang/fiber-starter/di"

	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv("APP_ENV")
	if env != "production" {
		// Load environment variables from .env only in non-production environments
		err := godotenv.Load()
		if err != nil {
			log.Println("Warning: No .env file found, using system environment variables")
		}
	}

	app, err := di.InitializeApp()
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
