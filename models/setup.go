package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		SkipDefaultTransaction: false,
	})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.Debug().AutoMigrate(
		&User{},
		&Product{},
		&Sale{},
		&Transaction{},
		&Resource{},
		&AcessResource{})
	if err != nil {
		panic("Failed to migration!")
	}

	DB = database
}
