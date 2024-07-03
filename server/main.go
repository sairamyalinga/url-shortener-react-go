package main

import (
	"log"
	"net/http"
	"os"
	router "urlShortener/server/router"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No env file found in the directory")
	}
}

func main() {
	r := router.Router()
    port := ":"+ os.Getenv("PORT")
	frontend := os.Getenv("FRONTEND")
	err := http.ListenAndServe(port, 
		handlers.CORS(
			handlers.AllowedOrigins([]string{frontend}), 
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowCredentials(),
		)(r))
	if err != nil {
		log.Fatalln("Error with the server,", err)
	}
}

