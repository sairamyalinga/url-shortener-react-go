package connection

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShortURL struct {
	URL      string             `bson:"url, omitempty"`
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	ShortID  string             `bson:"shortID,omitempty"`
	Username string             `bson:"user_name,omitempty"`
}

func (db *DBConnection) InsertURL(ctx context.Context, document ShortURL) (string, error) {
	callback := func(sesctx mongo.SessionContext) (interface{}, error) {
		res, err := db.urlCollection.InsertOne(sesctx, document)
		if err != nil {
			return "", fmt.Errorf("failed to insert the new document: %v", err)
		}

		insertedID := res.InsertedID.(primitive.ObjectID)
		shortID := insertedID.Hex()[18:]
		_, err = db.urlCollection.UpdateOne(
			sesctx,
			bson.M{"_id": insertedID},
			bson.M{"$set": bson.M{"shortID": shortID}},
		)
		if err != nil {
			return "", fmt.Errorf("failed to update document with shortID: %v", err)
		}

		return shortID, nil
	}

	session, err := db.mongoClient.StartSession()
	if err != nil {
		return "", err
	}
	defer session.EndSession(ctx)
	result, err := session.WithTransaction(ctx, callback)
	if err != nil {
		return "", err
	}
	return result.(string), nil
}

func (db *DBConnection) GetURLByID(ctx context.Context, id string) (*ShortURL, error) {
	var urlData ShortURL
	res := db.urlCollection.FindOne(ctx, bson.M{"shortID": id}).Decode(&urlData) 
	if res != nil {
		return nil, fmt.Errorf(res.Error())
	}

	return &urlData, nil
}

func (db *DBConnection) GetAllURLsByUsername(ctx context.Context, username string) ([]ShortURL, error) {
	filter := bson.D{{Key: "username", Value: username}}

	cursor, err := db.urlCollection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf(cursor.Err().Error())
	}
	defer cursor.Close(ctx)

	var urls []ShortURL
	for cursor.Next(ctx) {
		var urlDoc ShortURL
		if err := cursor.Decode(&urlDoc); err != nil {
			return nil, fmt.Errorf(err.Error())
		}

		urls = append(urls, urlDoc)
	}
	return urls, nil
}

func (db *DBConnection) DeleteURLByID(ctx context.Context, ID string) error {
	filter := bson.D{{Key: "shortID", Value: ID}}
	_, err := db.urlCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}
