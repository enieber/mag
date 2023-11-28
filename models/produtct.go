package models

import "gorm.io/gorm"

type ProductInput struct {
	Name string `json:"name" binding:"required`
	Type string `json:"type" binding:"required`
}

type Product struct {
	gorm.Model
	ID         uint   `json:"id" gorm:"primary_key"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	ExternalId string `json:"externtal_id"`
}
