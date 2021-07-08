package db

import (
	"fmt"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
)

const dbErrorMessage = "Error connecting to Database"

var db *gorm.DB

func InitDatabase(databaseURL string, version uint) {
	migrateConnection, err := migrate.New("file://db/migrate", databaseURL)
	if err != nil {
		fmt.Println("Error Creating Tables")
		return
	}

	currentVersion, _, _ := migrateConnection.Version()

	if version != currentVersion {
		err := migrateConnection.Migrate(version)

		if err != nil {
			fmt.Println("Error Creating Tables")
			return
		}
	}

	migrateConnection.Close()

	db, err := gorm.Open("postgres", databaseURL)
	if err != nil {
		fmt.Println(dbErrorMessage)
	}

	fmt.Println(db)
}

func GetDB() *gorm.DB {
	return db
}
