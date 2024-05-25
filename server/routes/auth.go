package router

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	_ "net/url"
	"time"
	"os"
	connection "urlShortener/server/db"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte(os.Getenv("JWT_KEY"))

func generateJWT(username string) (string, error) {

	claims := jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}
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

func Signin(w http.ResponseWriter, r *http.Request){

	var logindetails connection.User

	if err := json.NewDecoder(r.Body).Decode(&logindetails); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	client := connection.GetClient()
	collection := client.Database("go").Collection("user")

	var user connection.User
	err := collection.FindOne(context.TODO(), bson.M{"username": logindetails.UserName}).Decode(&user)

	if err != nil{
        http.Error(w, "UserName doesn't exists. Please, SignUp.", http.StatusNotFound)
	}
    
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(logindetails.Password))
	if err != nil{
		http.Error(w, "Error comparing password hashes", http.StatusInternalServerError)
	}

	token, err := generateJWT(user.UserName)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"JWTtoken": token})
}

