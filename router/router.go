package router

import (
	"log"
	"os"
	"time"

	authServices "example.com/m/services/auth"
	bookServices "example.com/m/services/book"
	userServices "example.com/m/services/user"
	"example.com/m/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
)

func Router() {
	app := fiber.New(fiber.Config{
		AppName: "Library API",

	})

	app.Use(limiter.New(limiter.Config{
		Max: 10,
		Expiration: 10 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	app.Post("/api/addBook", bookServices.AddBook, utils.EmployeeAuth)
	app.Get("/api/getBooks", bookServices.GetAllBooks)
	app.Get("/api/getBook/:id", bookServices.GetSingleBook)
	app.Delete("/api/removeBook/:id", bookServices.DeleteBook, utils.AdminOnly)
	app.Patch("/api/updateBook/:id", bookServices.UpdateSingleBook)

	app.Post("/api/auth/register", authServices.Register)
	app.Post("/api/auth/signin", authServices.SignIn)

	app.Get("/api/user/:id", userServices.GetSingleUser)
	app.Get("/api/users", userServices.GetAllUsers, utils.EmployeeAuth)
	app.Delete("/api/user/:id", userServices.RemoveUser)
	app.Patch("/api/user/:id", userServices.UpdateSingleUser)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(port))
}

