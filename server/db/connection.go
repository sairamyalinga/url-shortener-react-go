package connection

import (
	"context"
	"fmt"
	"time"
	"os"

	"github.com/joho/godotenv"
	_"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type urlStrings struct{
	Url string `bson:"url,omitempty"`
	Id  primitive.ObjectID  `json:_id bson:"_id`
}
var client *mongo.Client
func init() {

	if err := godotenv.Load(); err != nil {
		fmt.Println("No env file found in the directory")
	}

	// uri := os.Getenv("MONGODB_URI")

	// if uri == "" {
	// 	log.Fatal("Set the MONGODB_URI in the env")
	// }

	// client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	// if err != nil {
	// 	panic(err)
	// }

	ctx,cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return 
	}

	defer client.Disconnect(ctx)

}

func GetClient() *mongo.Client {
    return client
}
