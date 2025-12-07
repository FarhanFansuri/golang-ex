package main

import (
	"golang-ex/internal/config"
	"golang-ex/internal/utils"
	"strings"

	"github.com/joho/godotenv"

	"golang-ex/internal/models"

	"github.com/gin-gonic/gin"

	"golang-ex/internal/repository"
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
		users = []models.User{}
		repository := repository.NewUserRepository(config.DB)
		users, err = repository.GetUser()
		if err != nil {
			c.JSON(500, gin.H{
				"status":  "error",
				"message": "Failed to retrieve users",
			})
			return
		}
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "users retrieved successfully",
			"data":    users,
		})

	})

	r.POST("/register", func(c *gin.Context) {
		var data models.User
		c.BindJSON(&data)
		repository := repository.NewUserRepository(config.DB)
		createdUser, err := data.Create(repository.DB, data)
		if err != nil {
			c.JSON(500, gin.H{
				"status":  "error",
				"message": "Failed to create user",
			})
			return
		}
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "login endpoint",
			"data":    createdUser,
		})
	})
	r.POST("/login", func(c *gin.Context) {
		var data models.User
		c.BindJSON(&data)
		// Generate JWT token
		token, err := utils.GenerateJWT(data)
		if err != nil {
			c.JSON(500, gin.H{
				"status":  "error",
				"message": "Failed to generate token",
				"error":   err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "login endpoint",
			"token":   token,
		})
	})

	r.POST("/validation", func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(401, gin.H{
				"status":  "error",
				"message": "Missing Authorization header",
			})
			return
		}

		// Format harus: Bearer <token>
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(401, gin.H{
				"status":  "error",
				"message": "Invalid Authorization header format",
			})
			return
		}

		// Ambil token tanpa prefix "Bearer "
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		isValid, err := utils.ValidateJWT(tokenString)
		if err != nil || !isValid {
			c.JSON(401, gin.H{
				"status":  "error",
				"message": "Invalid token",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"status":  "success",
			"message": "Token is valid",
		})
	})

	// Start server
	r.Run(":3000")
}
