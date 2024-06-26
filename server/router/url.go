package router

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"

	connection "urlShortener/server/db"
	"urlShortener/server/middleware"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ul *URLProcessor) isValidURL(urlObj string) bool {

	u, err := url.Parse(urlObj)
	if err != nil {
		fmt.Println("Cannot parse the URL")
		return false 

	}
	if u.Path != "" {
		client := http.Client{
			Timeout: 5 * time.Second, 
		}
		resp, err := client.Get(urlObj)
		if err != nil {
			fmt.Println("Error performing GET request:", err)
			return false
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 399 {
			fmt.Println("Invalid response status code:", resp.StatusCode)
			return false
		}
	}
	ips, err := net.DefaultResolver.LookupIPAddr(context.Background(), u.Hostname()) // use req context
	if err != nil {
		return false
	}
	return len(ips) > 0

	
}

func insertURL(collection *mongo.Collection, object map[string]string, userName string) string {

	doc := connection.URLStrings{URL: object["url"], Username: userName}
	res, err := collection.InsertOne(context.TODO(), doc) // use req context
	if err != nil {
		fmt.Println("Failed to insert the new document", err)

	}

	insertedID := res.InsertedID.(primitive.ObjectID)
	shortID := insertedID.Hex()[18:]
	_, err = collection.UpdateOne( // this dependent db query should happen in a transaction to ensure atomicity - basic DBMS concept
		context.Background(), //huh, different contexts doesn't help. use req.context
		bson.M{"_id": insertedID},
		bson.M{"$set": bson.M{"shortID": shortID}},
	)
	if err != nil {
		fmt.Println("Failed to update document with shortID:", err)
	}
	return shortID

}

func (ul *URLProcessor) PostUrl(w http.ResponseWriter, r *http.Request) {
	var urlData map[string]string

	if err := json.NewDecoder(r.Body).Decode(&urlData); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	username := r.Context().Value(middleware.UsernameContextKey).(string)
	client := connection.GetClient()
	collection := client.Database("go").Collection("urlStrings")
	root := os.Getenv("ROOT")

	if ul.isValidURL(urlData["url"]) {
		id := insertURL(collection, urlData, username)
		shortURL := root + id
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"shortURL": shortURL})
	} else {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

}
