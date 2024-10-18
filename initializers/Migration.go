package initializers

import "example.com/m/models"


func Migration() {
	Db.AutoMigrate(&models.Book{})
	Db.AutoMigrate(&models.User{})
}