package main

import (
	"log"
	"net/http"

	env "github.com/joho/godotenv"
	"necodeo.com/m/v2/helpers"
	"necodeo.com/m/v2/rest/routers"
)

func main() {
	// Load .env file
	err := env.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database
	helpers.InitDB()
	defer helpers.DB.Close()

	// REST Routes
	router := routers.InitRoutes()

	// Start server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
