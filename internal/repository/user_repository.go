package repository

import (
	"golang-ex/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

// GORM CRUD
func (r *UserRepository) GetUser() ([]models.User, error) {
	var users []models.User
	r.DB.Find(&users)
	return users, nil
}
