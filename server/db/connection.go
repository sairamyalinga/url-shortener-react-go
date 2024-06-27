package connection

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBConnection struct {
	MongoClient *mongo.Client
}

func (c *DBConnection) init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No env file found in the directory")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	var err error
	c.MongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")).SetServerAPIOptions(serverAPI))
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
}
