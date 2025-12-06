package models

import (
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
