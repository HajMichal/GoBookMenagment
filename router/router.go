package router

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	authServices "main.go/services/auth"
	bookServices "main.go/services/book"
	userServices "main.go/services/user"
)

func Router() {
	app := fiber.New()

	app.Post("/api/addBook", bookServices.AddBook)
	app.Get("/api/getBook/:id", bookServices.GetSingleBook)
	app.Delete("/api/removeBook/:id", bookServices.DeleteBook)
	app.Patch("/api/updateBook/:id", bookServices.UpdateSingleBook)

	app.Post("/api/auth/register", authServices.Register)
	app.Post("/api/auth/signin", authServices.SignIn)

	app.Get("/api/user/:id", userServices.GetSingleUser)
	app.Get("/api/users", userServices.GetAllUsers)
	app.Delete("/api/user/:id", userServices.RemoveUser)
	app.Patch("/api/user/:id", userServices.UpdateSingleUser)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(port))
}

