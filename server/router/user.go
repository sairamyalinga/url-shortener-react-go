package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	utils "urlShortener/server/common"
	connection "urlShortener/server/db"

	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

type UserMethods struct {
	db *connection.DBConnection
}

func (um *UserMethods) RegisterUser(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()
	var user connection.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println(err)
		utils.SendJSONResponse(w, nil, http.StatusBadRequest, "Invalid request body")
		return	
	}

	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			utils.SendJSONResponse(w, nil,  http.StatusBadRequest, fmt.Sprintf("Field %s failed validation: %s", err.Field(), err.Tag()))
			return
		}
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.SendJSONResponse(w, nil,  http.StatusInternalServerError, "Error hashing password")
		return
	}

	user.Password = string(bytes)
	err = um.db.InsertUser(r.Context(), user)
	if err != nil {
		if err.Error() == connection.UserExistErr {
			utils.SendJSONResponse(w, nil, http.StatusBadRequest, "Already exists!")
			return
		}
		utils.SendJSONResponse(w, nil, http.StatusInternalServerError, "Internal error")
		return
	}
	utils.SendJSONResponse(w, nil, http.StatusCreated, "Signup Success!")
	
}

func (um *UserMethods) Signin(w http.ResponseWriter, r *http.Request) {

	var loginDetails connection.User 

	if err := json.NewDecoder(r.Body).Decode(&loginDetails); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := um.db.FindUserbyName(r.Context(), loginDetails.Username)
	if err!= nil {
		utils.SendJSONResponse(w, nil, http.StatusBadRequest, "Invalid Credentials")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDetails.Password))
	if err != nil {
		utils.SendJSONResponse(w, nil, http.StatusBadRequest, "Invalid Credentials")
		return
	}

	token, err := generateJWT(user.Username)
	if err != nil {
		utils.SendJSONResponse(w, nil, http.StatusInternalServerError, "Error generating JWT")
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.SendJSONResponse(w, map[string]string{"token": token}, http.StatusOK, "JWT generated")
}
