package router

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	_"net/url"
	connection "urlShortener/server/db"

	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson"
	_"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)


func RegisterUser(w http.ResponseWriter, r *http.Request){

	validate := validator.New()
	var user connection.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := validate.Struct(user)
	if err != nil{
		for _, err := range err.(validator.ValidationErrors){
 		http.Error(w, fmt.Sprintf("Field %s failed validation: %s", err.Field(), err.Tag()), http.StatusBadRequest)
        return
    }
}
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	user.Password = string(bytes)
	client := connection.GetClient()
	collection := client.Database("go").Collection("user")

	var existuser connection.User
	err = collection.FindOne(context.TODO(), bson.M{"username": user.UserName}).Decode(&existuser)

	if err == nil {
		// fmt.Println("UserName Already Exists!")
		http.Error(w, "Username already exists", http.StatusConflict)
		return

	} else if err != mongo.ErrNoDocuments {
		// An unexpected error occurred
		http.Error(w, "Error checking username availability", http.StatusInternalServerError)
		return
	}

	_, err = collection.InsertOne(context.TODO(), user)
	if err!= nil{
		fmt.Println("Failed to insert the new document",err)

	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"Alert": "SignUp Success!" })

}