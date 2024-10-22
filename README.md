# CRUD API Project in Go (Golang)


 Welcome to my first project in Golang! This project is a simple CRUD API built with Go, which includes key features like JWT authentication, cookie-based authorization, and email sending functionality.

## Features
* **CRUD operations**
* **JWT Authentication:** 
Secure your API endpoints with JSON Web Token (JWT) authentication using the `github.com/golang-jwt/jwt/v5` package.
* **Cookie-Based Authorization:** After successful login, the server sets a secure cookie to manage session state, adding an additional layer of security.
* **Email Sending:** This project supports sending emails, providing functionality that could be used for features like user verification, password resets, etc.

## Database
SQLite: The project uses SQLite as the database, connected via the GORM library `gorm.io/gorm and gorm.io/driver/sqlite`, which provides an easy-to-use ORM (Object-Relational Mapping) for managing database interactions.

## Packages Used
* `github.com/joho/godotenv`: For managing environment variables from a .env file, allowing you to easily switch between configurations and keep sensitive data like API keys out of your code.

* `golang.org/x/crypto/bcrypt`: For securely hashing and comparing user passwords.

* `gorm.io/gorm`: An ORM that simplifies database operations and allows you to work with your SQLite database using Go structures.

* `github.com/golang-jwt/jwt/v5`: Used to handle JSON Web Tokens (JWT) for stateless authentication, ensuring that user sessions are secure and efficient.

> [!NOTE]
>
> In this project, I used more advanced or improving packages such as Fiber (for web frameworks) and Gomail (for sending emails). While these packages offer powerful features and can make development faster, I realize that I shouldn't have used them when learning a new language like Go. It's important to first focus on understanding the core concepts and the standard libraries of Golang. In hindsight, mastering the fundamentals before incorporating advanced tools would have been a better learning approach.

Next project will include more pure golang 👋