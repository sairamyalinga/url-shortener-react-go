package main

import (
	"fmt"
	"log"
	"net/http"
	router "urlShortener/server/routes"

	
)

func main() {

	r := router.Router()
	fmt.Println("Starting the server on port 5050")
	err := http.ListenAndServe(":5050", r)
	if err != nil {
		log.Fatalln("Error with the server,", err)
	}



}
