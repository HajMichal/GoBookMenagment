package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	

	app.Get("/", func(c fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON("You pinged app")
	})
	app.Post("/api/addBook", AddBook)

	// _, err = db.Exec("INSERT INTO `Userlist` VALUES (AUTO_INCREMENT, 'Michał Haj','somehashedPwd', 5, 5)")
	// if err != nil {
	// 	panic(err.Error())
	// }

	log.Fatal(app.Listen(":3000"))
}

func DB() *sql.DB {
	username := os.Getenv("USERNAME")
	password:= os.Getenv("PASSWORD")
	ip:= os.Getenv("DB_IP")
	port:= os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open("mysql",fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, ip, port, dbName ))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	return db
}


func AddBook(c fiber.Ctx) error {
	book := new(AddBookStruct)

	if err := c.Bind().Body(book); err != nil {
		return err
	}
	
	if book.Title == "" || book.ISBN == "" || book.Author == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	return c.Status(http.StatusOK).JSON(book)
}
