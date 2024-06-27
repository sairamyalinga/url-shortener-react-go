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
)

type URLProcessor struct {
	db *connection.DBConnection
}

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

func (ul *URLProcessor) CreateURL(w http.ResponseWriter, r *http.Request) {
	var urlData map[string]string

	if err := json.NewDecoder(r.Body).Decode(&urlData); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	username := r.Context().Value(middleware.UsernameContextKey).(string)

	rootDomain := os.Getenv("ROOT")

	if ul.isValidURL(urlData["url"]) {
		id, err := ul.db.InsertURL(r.Context(), connection.URLStrings{Username: username, URL: urlData["url"]})

		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		shortURL := rootDomain + id
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"shortURL": shortURL})
	} else {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

}
