package main

import (
	"golang-ex/internal/config"

	"github.com/joho/godotenv"

	"golang-ex/internal/models"

	"github.com/gin-gonic/gin"
)

var (
	users []models.User
)

func main() {
	// Application entry point
	err := godotenv.Load("../../.env")
	if err != nil {
		panic("Error loading .env file")
	}

	r := gin.Default()
	// Database connection would be initialized here
	config.ConnectDatabaseMySQL()

	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, users)
	})
	// Start server
	r.Run(":3000")
}
