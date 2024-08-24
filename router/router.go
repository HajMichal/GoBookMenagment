package router

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	services "main.go/services/book"
)

func Router() {
	app := fiber.New()

	app.Post("/api/addBook", services.AddBook)
	app.Get("/api/getBook/:id", services.GetSingleBook)
	app.Delete("/api/removeBook/:id", services.DeleteBook)
	app.Patch("/api/updateBook/:id", services.UpdateSingleBook)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(port))
}

