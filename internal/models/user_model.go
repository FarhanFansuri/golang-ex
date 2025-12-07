package models

import (
	"golang-ex/internal/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       string `json:"id" gorm:"type:char(36);primaryKey"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return err
}

func (u *User) Create(tx *gorm.DB, data User) (User, error) {
	hashedPassword, err := utils.HashPassword(data.Password)
	if err != nil {
		return User{}, err
	}
	data.Password = hashedPassword
	err = tx.Create(&data).Error
	if err != nil {
		return User{}, err
	}
	return data, nil
}

func (u *User) GetAll(tx *gorm.DB) ([]User, error) {
	var users []User
	err := tx.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
