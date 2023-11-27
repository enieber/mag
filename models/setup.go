package models

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbUrl := os.Getenv("DATABASE_URL")
	database, err := gorm.Open(sqlite.Open(dbUrl), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Book{})
	if err != nil {
		return
	}

	DB = database
}
