package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"name"`
	Email    string `json:"email" gorm:"email"`
	Password string `json:"password" gorm:"password"`
}

func (User) TableName() string {
	return "users"
}
