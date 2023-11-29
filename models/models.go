package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int `gorm:"primaryKey"`
	Name     string
	Email    string
	Products []Product `gorm:"many2many:user_products;"`
}

type Product struct {
	gorm.Model
	ID         int `gorm:"primaryKey"`
	Name       string
	Sales      []Sale `gorm:"foreignkey:ProductID;references:ID"`
	Type       string `json:"type"`
	ExternalId string `json:"externtal_id"`
}

type Sale struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	UserID    int
	ProductID int
	Status    string `json:"status"`
}

type Transaction struct {
	gorm.Model
	ID     int `gorm:"primaryKey"`
	SaleID int
	Status string `json:"status"`
}

type Sales struct {
	gorm.Model
	ID      uint    `json:"id" gorm:"primary_key"`
	Product Product `gorm:"foreignKey:ProductID;references:ID" json:"product"`
	User    User    `gorm:"foreignKey:UserID;references:ID" json:"user"`
}

type SalesInput struct {
	IdProduct uint `json:"id_product"`
	IdUser    uint `json:"id_user"`
}

type TransactionReturn struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

type UserInput struct {
	Name  string `json:"name" binding:"required`
	Email string `json:"email" binding:"required`
}

type ProductInput struct {
	Name string `json:"name" binding:"required`
	Type string `json:"type" binding:"required`
}

type Resource struct {
	gorm.Model
	ID      uint    `json:"id" gorm:"primary_key"`
	Product Product `gorm:"foreignKey:ProductID;references:ID" json:"product"`
	SalesID uint
	Status  string
	IP      string
}

type AcessResource struct {
	gorm.Model
	AcessPublicKey string
	ResourceID     uint
}
