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
	router := mux.NewRouter()
	router.HandleFunc("/{id}", RedirectUrl).Methods("GET")

	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("/signup", RegisterUser).Methods("POST")
	apiRouter.HandleFunc("/login", Signin).Methods("POST")

	authenticatedRouter := apiRouter.NewRoute().Subrouter()
	authenticatedRouter.Use(middleware.JWTMiddleware)

	ul := URLProcessor{db: &db}
	authenticatedRouter.HandleFunc("/urls", GetallURLs).Methods("GET")
	authenticatedRouter.HandleFunc("/url", ul.CreateURL).Methods("POST")
	authenticatedRouter.HandleFunc("/url", DeleteUrl).Methods("DELETE")

	return router
}

// router has nothing to do with shortURL methods
// move everything below to new folder/file

func RedirectUrl(w http.ResponseWriter, r *http.Request) { // move to a struct's method

	params := mux.Vars(r)
	shortId := params["id"]
	fmt.Println(shortId)                                         //?
	client := connection.GetClient()                             // move as a "singleton" property
	collection := client.Database("go").Collection("urlStrings") // move as a "singleton" property
	var urlData connection.URLStrings
	res := collection.FindOne(r.Context(), bson.M{"shortID": shortId}).Decode(&urlData) // use request context itself
	if res != nil {                                                                     //spaces????
		fmt.Println("No Short Url found") // incorrect case
		// https://chatgpt.com/c/bad45add-cbd2-4063-a56f-f157c8cb884b
	}

	http.Redirect(w, r, urlData.URL, http.StatusFound)

}

func GetallURLs(w http.ResponseWriter, r *http.Request) {
	// create separate file for DB queries, and make the controller methods just call them
	// (controller should not know what DB methods are doing inside - abstraction OOPS concept)

	username := r.Context().Value(middleware.UsernameContextKey).(string)
	filter := bson.D{{Key: "username", Value: username}}
	client := connection.GetClient()
	collection := client.Database("go").Collection("urlStrings")
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		http.Error(w, "Error fetching URLs", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var urls []map[string]string
	for cursor.Next(context.TODO()) {
		var urlDoc connection.URLStrings
		if err := cursor.Decode(&urlDoc); err != nil {
			http.Error(w, "Error parsing URL document", http.StatusInternalServerError)
			return
		}
		responseURL := map[string]string{
			"Url":      urlDoc.URL,
			"ShortURL": "http://localhost:5050/" + urlDoc.ShortID,
			"User":     urlDoc.Username,
			"ShortID":  urlDoc.ShortID,
		}
		urls = append(urls, responseURL)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(urls)
}

func DeleteUrl(w http.ResponseWriter, r *http.Request) {

	var requestBody map[string]string

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	shortID := requestBody["shortID"]

	filter := bson.D{{Key: "shortID", Value: shortID}}
	client := connection.GetClient()
	collection := client.Database("go").Collection("urlStrings")

	_, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		// log all the errors using fmt.print - helps in debugging
		http.Error(w, "Failed to delete", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json") // why only here?
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Document deleted successfully"}`)) // why not json encode like normal?
}
