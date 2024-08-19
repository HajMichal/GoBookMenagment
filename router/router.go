package router

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"main.go/services"
)

func Router() {
	app := fiber.New()

	app.Post("/api/addBook", services.AddBook)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(port))
}

