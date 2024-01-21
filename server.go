package main

import (
	"log"
	"net/http"

	env "github.com/joho/godotenv"
	"necodeo.com/m/v2/helpers"
	"necodeo.com/m/v2/rest/routers"
)

func main() {
	err := env.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	helpers.InitDB()
	defer helpers.DB.Close()

	router := routers.InitRoutes()

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
