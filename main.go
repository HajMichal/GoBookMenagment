package main

import (
	"example.com/m/initializers"
	"example.com/m/router"
)

func main() {
	
	initializers.LoadEnv()
	initializers.ConnectToDB()
	// initializers.Migration()
	
	router.Router()
}







