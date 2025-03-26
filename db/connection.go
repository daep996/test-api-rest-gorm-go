package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN string
var DB *gorm.DB

func DBConnection() {
	DSN = os.Getenv("DSN")

	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database connected!")
	}
}
