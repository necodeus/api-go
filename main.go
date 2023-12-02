package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"necodeo.com/m/v2/db"
	"necodeo.com/m/v2/routers"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDB()
	defer db.DB.Close()

	router := routers.InitRoutes()
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
