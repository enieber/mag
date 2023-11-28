package models

import "gorm.io/gorm"

type UserInput struct {
	Name  string `json:"name" binding:"required`
	Email string `json:"email" binding:"required`
}

type User struct {
	gorm.Model
	ID         uint   `json:"id" gorm:"primary_key"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	ExternalId string `json:"externtal_id"`
}
