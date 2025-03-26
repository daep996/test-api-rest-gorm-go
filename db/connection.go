package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=daep password=daep dbname=go_test_crud port=5432"
var DB *gorm.DB

func DBConnection() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database connected!")
	}
}
