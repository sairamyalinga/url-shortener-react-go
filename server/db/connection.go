package connection

import (
	"context"
	"fmt"
	"os"
	"time"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type URLStrings struct {
	Url string             `bson:"url, omitempty"`
	Id primitive.ObjectID `bson:"_id,omitempty"`
	ShortID string         `bson:"shortID,omitempty"`
	UserName string 		`json:"user_name"`
}

type User struct {
	UserId primitive.ObjectID `bson:"_id,omitempty"`
	UserName string `json:"user_name" validate:"required,min=4,max=30"`
	Password string `json:"password" validate:"required,min=8,max=30"`

}

var client *mongo.Client

func init() {

	if err := godotenv.Load(); err != nil {
		fmt.Println("No env file found in the directory")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")).SetServerAPIOptions(serverAPI))
	// client = mclient
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	  }
	  fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	//defer client.Disconnect(ctx)

}

func GetClient() *mongo.Client {
	return client
}
