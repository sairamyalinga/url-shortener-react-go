package connection

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	UserID   primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `json:"user_name" validate:"required,min=4,max=30"`
	Password string             `json:"password" validate:"required,min=8,max=30"`
}

var UserExistErr string = "username already exists"
func (db *DBConnection) InsertUser(ctx context.Context, user User) (error) {
	
	var existingUser User 
	err := db.userCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&existingUser)

	if err == nil {	
		return fmt.Errorf(UserExistErr)

	} else if err != mongo.ErrNoDocuments {
		// http.Error(w, "Error checking username availability", http.StatusInternalServerError) //bad response format
		return fmt.Errorf("failed to signup. Try again : %v", err)
	}
	_, err = db.userCollection.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to insert the new document: %v", err)
	}	
	return nil
}

func (db *DBConnection) FindUserbyName(ctx context.Context, username User) (error) {
	
}