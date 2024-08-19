package initializers

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectToDB() {
	var err error

	Db, err = gorm.Open(sqlite.Open("db/sqlite.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}