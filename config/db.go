package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB connection variables
var (
	MongoClient    *mongo.Client
	BookCollection *mongo.Collection
)

// ConnectDB establishes connection with MongoDB
func ConnectDB() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://harshakumara1998030944:reqlHarsha321@cluster0.70hsm08.mongodb.net/")

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	// Set client and collection
	MongoClient = client
	BookCollection = client.Database("bookstore").Collection("books")
}
