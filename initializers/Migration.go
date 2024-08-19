package initializers

import "main.go/models"


func Migration() {
	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.User{})
}