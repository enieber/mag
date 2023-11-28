package models

type Transaction struct {
	ID           uint `json:"id" gorm:"primary_key"`
	Sales        Sales
	StatusChange string
}

type TransactionReturn struct {
	Id     uint   `json:"id"`
	Status string `json:"status"`
}
