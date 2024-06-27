package router

import (
	"context"
	"encoding/json"
	"fmt"

	"net/http"

	connection "urlShortener/server/db"
	middleware "urlShortener/server/middleware"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func Router() *mux.Router {
	db := connection.DBConnection{}
	ul := URLProcessor{db: &db}
	router := mux.NewRouter()

	router.HandleFunc("/{id}", ul.RedirectUrl).Methods("GET")

	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("/signup", RegisterUser).Methods("POST")
	apiRouter.HandleFunc("/login", Signin).Methods("POST")

	authenticatedRouter := apiRouter.NewRoute().Subrouter()
	authenticatedRouter.Use(middleware.JWTMiddleware)

	authenticatedRouter.HandleFunc("/urls", ul.GetallURLs).Methods("GET")
	authenticatedRouter.HandleFunc("/url", ul.CreateURL).Methods("POST")
	authenticatedRouter.HandleFunc("/url", ul.DeleteUrl).Methods("DELETE")

	return router
}

// router has nothing to do with shortURL methods
// move everything below to new folder/file
