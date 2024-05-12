package main

import (
	"fmt"
	"log"
	"net/http"
	router "urlShortener/server/routes"
	"github.com/gorilla/handlers"

	
)

func main() {

	r := router.Router()
	
	fmt.Println("Starting the server on port 5050")
	err := http.ListenAndServe(":5050", 
				handlers.CORS(
					handlers.AllowedOrigins([]string{"*"}),
					handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
					handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
				)(r))
	if err != nil {
		log.Fatalln("Error with the server,", err)
	}



}
