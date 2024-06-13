package router

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	connection "urlShortener/server/db"
	middleware "urlShortener/server/middleware"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/api/signup", RegisterUser).Methods("POST")
	router.HandleFunc("/api/login", Signin).Methods("POST")

	middlerouter := router.PathPrefix("/api").Subrouter()
	middlerouter.Use(middleware.JWTMiddleware)
	middlerouter.HandleFunc("/shorturl", CreateUrl).Methods("POST")
	router.HandleFunc("/{id}", RedirectUrl).Methods("GET")

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

func insertURL(ctx context.Context, collection *mongo.Collection, object  map[string]string) (string) {
	
	doc := connection.URLStrings{Url:object["url"]}

	res, err := collection.InsertOne(context.TODO(), doc)
	if err!= nil{
		fmt.Println("Failed to insert the new document",err)

	}
	//fmt.Println("Inserted New document")
	insertedID := res.InsertedID.(primitive.ObjectID)

    // Convert ObjectID to hex string and extract last 6 characters
    shortID := insertedID.Hex()[18:] 
	_, err = collection.UpdateOne(
        context.Background(),
        bson.M{"_id": insertedID},
        bson.M{"$set": bson.M{"shortID": shortID}},
    )
	if err != nil {
		fmt.Println("Failed to update document with shortID:", err)
    }

    fmt.Println("Inserted New document with shortID:", shortID)

	return shortID
}

func CreateUrl(w http.ResponseWriter, r *http.Request) {
	
	if r.Method == "OPTIONS" {
		return
	}
	fmt.Println(r.URL);
	fmt.Println(r.Body)
	var urlData map[string]string
	if err := json.NewDecoder(r.Body).Decode(&urlData); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	client := connection.GetClient()
	collection := client.Database("go").Collection("urlStrings")
	
	
	if isValidURL(urlData["url"]) {
		fmt.Println("hi")
		id := insertURL(r.Context(), collection, urlData)
		shortURL := "http://localhost:5050/" + id
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{"shortURL": shortURL})
	} else {
		fmt.Println("bye")
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	
	
	
}

func RedirectUrl(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	shortId := params["id"]
	fmt.Println(shortId)
	client := connection.GetClient()
	collection := client.Database("go").Collection("urlStrings")

	var urlData connection.URLStrings

	res := collection.FindOne(context.TODO(), bson.M{"shortID": shortId}).Decode(&urlData)
	fmt.Println(res)

	if res!=nil{
		fmt.Println("No Short Url found")
	}

	http.Redirect(w, r, urlData.Url, http.StatusFound)



}

