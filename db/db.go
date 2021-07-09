package db

import (
	"fmt"
	"golangorm/config"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
)

const dbErrorMessage = "Error connecting to Database"

var db *gorm.DB

func InitDatabase() {
	migrateConnection, err := migrate.New("file://db/migrate", config.GetConfig().Database.URL)
	if err != nil {
		fmt.Println("Error Creating Tables")
		return
	}

	currentVersion, _, _ := migrateConnection.Version()

	if config.GetConfig().Database.Version != currentVersion {
		err := migrateConnection.Migrate(config.GetConfig().Database.Version)

		if err != nil {
			fmt.Println("Error Creating Tables")
			return
		}
	}

	migrateConnection.Close()

	db, err = gorm.Open("postgres", config.GetConfig().Database.URL)
	if err != nil {
		fmt.Println(dbErrorMessage)
	}

	fmt.Println(db)
	db.LogMode(config.GetConfig().Database.LogMode)
}

func GetDB() *gorm.DB {
	return db
}
