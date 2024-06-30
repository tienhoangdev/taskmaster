package main

import (
	"golang_project_base/config"
	"golang_project_base/routers"

	_ "github.com/joho/godotenv/autoload"
	"os"
)

func main() {
	// Initialize the database
	config.InitDB()

	// Set up the router
	r := routers.SetupRouter()

	// Run the server
	port := os.Getenv("PORT")
	r.Run(":" + port)
}
