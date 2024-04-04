package router

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	connection "urlShortener/server/db"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
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

func isValidURL(urlObj string) bool{

	u, err := url.Parse(urlObj)
	if err != nil {
		fmt.Println("Cannot parse the URL")
	}
    ips, err := net.DefaultResolver.LookupIPAddr(context.Background(), u.Hostname())

	if err != nil {
		return false
	}
	return len(ips)>0


}
func insertURL(ctx context.Context, collection *mongo.Collection, object  map[string]string) (connection.URLStrings) {
	
	newID := primitive.NewObjectID()
	doc := connection.URLStrings{Url:object["url"], Id: newID}

	_, err := collection.InsertOne(context.TODO(), doc)
	if err!= nil{
		fmt.Println("Failed to insert the new document")

	}
	fmt.Println("Inserted New document")

	return doc


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
	
	
	if isValidURL(urlData["url"]) {
		insertURL(r.Context(), collection, urlData)


	} else {
		log.Fatalln("Invalid URL")
	}
	// shorturl := 
	
	// return 

}

func RedirectUrl(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	shortId := params["id"]
	client := connection.GetClient()
	collection := client.Database("go").Collection("urlStrings")

	var urlData connection.URLStrings

	res := collection.FindOne(context.TODO(), bson.M{"id":shortId}).Decode(&urlData)

	if res!=nil{
		fmt.Println("No Short Url found")
	}

	http.Redirect(w, r, urlData.Url, http.StatusFound)



}

