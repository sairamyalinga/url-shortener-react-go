package router

import (
	connection "urlShortener/server/db"
	middleware "urlShortener/server/middleware"

	"github.com/gorilla/mux"	
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

