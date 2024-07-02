package router

import (
	connection "urlShortener/server/db"
	middleware "urlShortener/server/middleware"

	"github.com/gorilla/mux"	
)

func Router() *mux.Router {
	db := connection.DBConnection{}
	db.InitializeDatabase()
	
	ul := URLProcessor{db: &db}
	um := UserMethods{db : &db}
	router := mux.NewRouter()

	router.HandleFunc("/{id}", ul.RedirectUrl).Methods("GET")

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/signup", um.RegisterUser).Methods("POST")
	apiRouter.HandleFunc("/login", um.Signin).Methods("POST")

	authenticatedRouter := apiRouter.NewRoute().Subrouter()
	authenticatedRouter.Use(middleware.JWTMiddleware)

	authenticatedRouter.HandleFunc("/urls", ul.GetallURLs).Methods("GET")
	authenticatedRouter.HandleFunc("/url", ul.CreateURL).Methods("POST")
	authenticatedRouter.HandleFunc("/url", ul.DeleteUrl).Methods("DELETE")

	return router
}

