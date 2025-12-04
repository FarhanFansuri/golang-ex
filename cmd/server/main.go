package main

import (
	"golang-ex/internal/config"

	"github.com/joho/godotenv"
)

func main() {
	// Application entry point
	err := godotenv.Load("../../.env")
	if err != nil {
		panic("Error loading .env file")
	}

	// Database connection would be initialized here
	config.ConnectDatabaseMySQL()
}
