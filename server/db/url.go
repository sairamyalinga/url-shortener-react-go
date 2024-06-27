package connection

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShortURL struct {
	URL      string             `bson:"url, omitempty"`
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	ShortID  string             `bson:"shortID,omitempty"`
	Username string             `bson:"user_name,omitempty"`
}

func (db *DBConnection) InsertURL(ctx context.Context, document ShortURL) (string, error) {
	res, err := db.urlCollection.InsertOne(ctx, document) // use req context
	if err != nil {
		return "", fmt.Errorf("Failed to insert the new document: %v", err)

	}

	insertedID := res.InsertedID.(primitive.ObjectID)
	shortID := insertedID.Hex()[18:]
	_, err = db.urlCollection.UpdateOne( // this dependent db query should happen in a transaction to ensure atomicity - basic DBMS concept
		ctx, //huh, different contexts doesn't help. use req.context
		bson.M{"_id": insertedID},
		bson.M{"$set": bson.M{"shortID": shortID}},
	)
	if err != nil {
		return "", fmt.Errorf("Failed to update document with shortID: %v", err)
	}
	return shortID, nil

}

func (db *DBConnection) GetURLByID(ctx context.Context, id string) (*ShortURL, error) {
	var urlData ShortURL
	res := db.urlCollection.FindOne(ctx, bson.M{"shortID": id}).Decode(&urlData) // use request context itself
	if res != nil {                                                              //spaces????
		// fmt.Println("No Short Url found") // incorrect case
		// https://chatgpt.com/c/bad45add-cbd2-4063-a56f-f157c8cb884b
		return nil, fmt.Errorf(res.Error())
	}

	return &urlData, nil
}
