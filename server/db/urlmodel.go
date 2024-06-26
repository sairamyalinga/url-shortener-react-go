package connection

import "go.mongodb.org/mongo-driver/bson/primitive"

type URLStrings struct {
	URL      string             `bson:"url, omitempty"`
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	ShortID  string             `bson:"shortID,omitempty"`
	Username string             `bson:"user_name,omitempty"`
}
