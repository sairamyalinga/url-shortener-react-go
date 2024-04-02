package router

import (
	"context"
	"fmt"
	"net/http"
	connection "urlShortener/server/db"
	"encoding/json"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/api/shorturl", CreateUrl).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/shorturl/{id}", RedirectUrl).Methods("GET", "OPTIONS")
	

	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//     http.ServeFile(w, r, "urlShortener/frontend/index.html")
	// })

	return router
}

func CreateUrl(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)

	var urlData map[string]string
	if err := json.NewDecoder(r.Body).Decode(&urlData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	client := connection.GetClient()
	collection := client.Database("go").Collection("urlStrings")
	
	isValidURL(urlData)
	insertURL(r.Context(), collection, urlData)

}
func isValidURL(urlObj map[string]string){

}
func insertURL(ctx context.Context, collection *mongo.Collection, object  map[string]string) {
	
	newID := primitive.NewObjectID()
	doc := connection.URLStrings{Url:object["url"], Id: newID}

	_, err := collection.InsertOne(context.TODO(), doc)
	if err!= nil{
		fmt.Println("Failed to insert the new document")

	}
	fmt.Println("Inserted New document")

}
func RedirectUrl(w http.ResponseWriter, r *http.Request) {

}

