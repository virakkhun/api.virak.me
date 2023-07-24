package database

import (
	"log"
	"os"

	"api.virak.me/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInitializer() {
	dsn := os.Getenv("DSN")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if utils.IsNotNil(err) {
		log.Fatal("Failed to connect to DB")
	}

	DB = db
}
