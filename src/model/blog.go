package model

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title       string `json:"title" gorm:"title"`
	Description string `json:"description" gorm:"description"`
	Image       string `json:"image" gorm:"image"`
}

func (Blog) TableName() string {
	return "blogs"
}
