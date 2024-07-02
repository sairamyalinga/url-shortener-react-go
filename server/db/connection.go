package connection

import (
	"context"
	"fmt"
	"os"
	"time"

	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBConnection struct {
	mongoClient   *mongo.Client
	database      *mongo.Database
	urlCollection *mongo.Collection
	userCollection *mongo.Collection
}


func (c *DBConnection) InitializeDatabase() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	var err error
	c.mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")).SetServerAPIOptions(serverAPI))
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}

	c.database = c.mongoClient.Database("go")
	c.urlCollection = c.database.Collection("urlStrings")
	c.userCollection = c.database.Collection("user")

}
