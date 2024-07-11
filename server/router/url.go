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

	utils "urlShortener/server/common"
	connection "urlShortener/server/db"
	"urlShortener/server/middleware"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type URLProcessor struct {
	db *connection.DBConnection
}

func (ul *URLProcessor) isValidURL(ctx context.Context, urlObj string) bool {
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

	ips, err := net.DefaultResolver.LookupIPAddr(ctx, u.Hostname()) 
	if err != nil {
		return false
	}
	return len(ips) > 0

}

func (ul *URLProcessor) CreateURL(w http.ResponseWriter, r *http.Request) {
	var urlData map[string]string
	if err := json.NewDecoder(r.Body).Decode(&urlData); err != nil {
		fmt.Println(err)
		utils.SendJSONResponse(w, nil, http.StatusBadRequest, "Invalid request body")
		return
	}
	username := r.Context().Value(middleware.UsernameContextKey).(string)
	rootDomain := os.Getenv("ROOT")

	if ul.isValidURL(r.Context(), urlData["url"]) {
		found, _ := ul.db.FindURLByUsername(r.Context(), username, urlData["url"])
		if found {
			utils.SendJSONResponse(w, nil, http.StatusConflict, "Already Exists")
			return
		}
		id, err := ul.db.InsertURL(r.Context(), connection.ShortURL{Username: username, URL: urlData["url"]})
		if err != nil {
			utils.SendJSONResponse(w, nil, http.StatusInternalServerError, "Internal error")
			return
		}
		shortURL := rootDomain + id
		utils.SendJSONResponse(w, map[string]string{"shortURL": shortURL}, http.StatusCreated, "Created shortURL")

	} else {
		utils.SendJSONResponse(w, nil, http.StatusBadRequest, "Invalid URL")
		return
	}

}

func (ul *URLProcessor) RedirectUrl(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	urlRecord, err := ul.db.GetURLByID(r.Context(), params["id"])

	if err == mongo.ErrNoDocuments {
		utils.SendJSONResponse(w, nil, http.StatusNotFound, "URL not found")
		return
	}
	if err != nil {
		utils.SendJSONResponse(w, nil, http.StatusInternalServerError, "Error fetching URL document")
		return
	}

	http.Redirect(w, r, urlRecord.URL, http.StatusFound)

}

func (ul *URLProcessor) GetallURLs(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value(middleware.UsernameContextKey).(string)
	urls, err := ul.db.GetAllURLsByUsername(r.Context(), username)
	if err != nil {

		utils.SendJSONResponse(w, nil, http.StatusInternalServerError, "Error fetching URLs")

	}
	utils.SendJSONResponse(w, urls, http.StatusOK, "Fetched URLs")

}

func (ul *URLProcessor) DeleteUrl(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]string
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	shortID := requestBody["shortID"]
	err := ul.db.DeleteURLByID(r.Context(), shortID)
	if err != nil {
		fmt.Printf("Failed to delete: %s", err)
		utils.SendJSONResponse(w, nil, http.StatusInternalServerError, "Failed to delete")
		return
	}
	utils.SendJSONResponse(w, nil, http.StatusOK, "Document Deleted Successfully")

}
