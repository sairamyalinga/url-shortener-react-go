package connection

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	UserID   primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `json:"user_name" validate:"required,min=4,max=30"`
	Password string             `json:"password" validate:"required,min=8,max=30"`
}
