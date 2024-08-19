package initializers

import (
	"main.go/models"
)


func Migration() {
	Db.AutoMigrate(&models.Book{})
	Db.AutoMigrate(&models.User{})
}