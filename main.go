package main

import (
	"main.go/initializers"
	"main.go/router"
)

func main() {
	
	initializers.LoadEnv()
	initializers.ConnectToDB()
	
	router.Router()
}







