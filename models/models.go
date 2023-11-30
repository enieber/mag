package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string
	Products []Product `gorm:"many2many:user_products;"`
}

type Product struct {
	gorm.Model
	Name       string
	Sales      []Sale `gorm:"foreignkey:ProductID;references:ID"`
	Type       string `json:"type"`
	ExternalId string `json:"externtal_id"`
}
type ProductReturn struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	ExternalId string `json:"externtal_id"`
}
type Sale struct {
	gorm.Model
	UserID    uint
	ProductID uint
	Status    string `json:"status"`
}

type Transaction struct {
	gorm.Model
	SaleID uint
	Status string `json:"status"`
}

type Sales struct {
	gorm.Model
	Product Product `gorm:"foreignKey:ProductID;references:ID" json:"product"`
	User    User    `gorm:"foreignKey:UserID;references:ID" json:"user"`
}

type Transactions struct {
	gorm.Model
	Sale Sale `gorm:"foreignKey:SaleID;references:ID" json:"sale"`
}

type TransactionInput struct {
	ID     uint   `json:"id"`
	Status string `json:"status"`
}

type SalesInput struct {
	IdProduct uint `json:"id_product"`
	IdUser    uint `json:"id_user"`
}

type TransactionReturn struct {
	ID     uint   `json:"id"`
	Status string `json:"status"`
}

type UserInput struct {
	Name  string `json:"name" binding:"required`
	Email string `json:"email" binding:"required`
}

type UserReturn struct {
	ID    uint   `json:"id"`
	Name  string `json:"name" binding:"required`
	Email string `json:"email" binding:"required`
}

type ProductInput struct {
	Name string `json:"name" binding:"required`
	Type string `json:"type" binding:"required`
}

type Resource struct {
	gorm.Model
	SalesID uint
	Status  string
	IP      string
}

type Resources struct {
	Sale Sale `gorm:"foreignKey:SaleID;references:ID" json:"sale"`
}

type AcessResource struct {
	gorm.Model
	AcessPublicKey string
	ResourceID     uint
}

type AcessResources struct {
	Resource Resource `gorm:"foreignKey:ResourceID;references:ID" json:"resource"`
}
