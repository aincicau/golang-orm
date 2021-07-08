package db

import (
	_ "github.com/golang-migrate/migrate/database/postgres"
	"github.com/jinzhu/gorm"
)

func InitDatabase(databaseURL string) {
	gorm.Open("postgres", databaseURL)
}
