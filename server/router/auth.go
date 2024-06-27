package router

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	_ "net/url"
	"os"
	"time"
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

// move the methods to belong to a struct (lets call it User {})
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()
	var user connection.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
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
	client := connection.GetClient()                       // init this client ONLY once when initialising the struct User{client: connection.GetClient()}
	collection := client.Database("go").Collection("user") // move to struct

	var existuser connection.User //existuser???? is that even a word
	err = collection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&existuser)

	if err == nil {
		http.Error(w, "Username already exists", http.StatusConflict) // errors are not json encoded response
		return

	} else if err != mongo.ErrNoDocuments {
		http.Error(w, "Error checking username availability", http.StatusInternalServerError) //bad response format
		return
	}

	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println("Failed to insert the new document", err)

	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"Alert": "SignUp Success!"}) //alert??? => message
	// S U caps? make message cleaner /formal

}

func Signin(w http.ResponseWriter, r *http.Request) {

	var logindetails connection.User //use snake_case or camelCase consistently - logindetails is neither!

	if err := json.NewDecoder(r.Body).Decode(&logindetails); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	client := connection.GetClient()
	collection := client.Database("go").Collection("user")

	var user connection.User
	err := collection.FindOne(context.TODO(), bson.M{"username": logindetails.Username}).Decode(&user)

	if err != nil {
		http.Error(w, "UserName doesn't exists. Please, SignUp.", http.StatusNotFound) // who even writes like this
		//https://chatgpt.com/c/bad45add-cbd2-4063-a56f-f157c8cb884b
		// better say incorrect credentials to prevent "user enumeration" vulnerability
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(logindetails.Password))
	if err != nil {
		http.Error(w, "Error comparing password hashes", http.StatusInternalServerError) // huh? you should say incorrect credentials
	}

	token, err := generateJWT(user.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"JWTtoken": token}) //JWTToken???? is that a good key? why not just call it "token"?
}
