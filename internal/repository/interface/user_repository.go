package repository

import "golang-ex/internal/models"

type UserRepositoryInterface interface {
	// Create
	CreateUser(user models.User) error
	// Read
	GetUserByID(id uint) (models.User, error)
	GetUser() ([]models.User, error)
}
