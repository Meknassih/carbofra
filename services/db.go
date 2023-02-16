package services

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Open_client() *mongo.Client {
	mongo_uri := os.Getenv("MONGO_URI")
	if mongo_uri == "" {
		mongo_uri = "mongodb://localhost:27017"
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongo_uri))
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB")

	return client
}

func Close_client(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
