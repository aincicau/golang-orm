package db

import (
	"fmt"

	_ "github.com/golang-migrate/migrate/database/postgres"
	"github.com/jinzhu/gorm"
)

const dbErrorMessage = "Error connecting to Database"

var db *gorm.DB

func InitDatabase(databaseURL string) {
	db, err := gorm.Open("postgres", databaseURL)
	if err != nil {
		fmt.Println(dbErrorMessage)
	}

	fmt.Println(db)
}

func GetDB() *gorm.DB {
	return db
}
