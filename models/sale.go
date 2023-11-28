package models

type SalesInput struct {
	IdProduct uint `json:"id_product"`
	IdUser    uint `json:"id_user"`
}

type Sales struct {
	ID      uint    `json:"id" gorm:"primary_key"`
	Product Product `json:"product"`
	User    User    `json:"user"`
	Status  string  `json:"status"`
}
