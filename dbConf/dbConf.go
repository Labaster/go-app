package dbConf

import (
	"context"
	"fmt"
	"log"
	"os"
	"errors"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client = nil

// Connect connects to MongoDB and returns the collection
func Connect(dbName string, collectionName string) (*mongo.Collection, error) {
    if dbName == "" || collectionName == "" {
        errMsg := "dbName and collectionName are required"
        return nil, errors.New(errMsg)
    }

    err := godotenv.Load()
    if err != nil {
        return nil, err
    }

    MONGO_URI := os.Getenv("MONGO_URI")
    if MONGO_URI == "" {
        return nil, errors.New("MONGO_URI is not set")
    }

    clientOptions := options.Client().ApplyURI(MONGO_URI)
    client, err = mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        return nil, fmt.Errorf("mongo.Connect err -->> %v", err)
    }

    err = client.Ping(context.Background(), nil)
    if err != nil {
        return nil, fmt.Errorf("client.Ping err -->> %v", err)
    }

    collection := client.Database(dbName).Collection(collectionName)
    fmt.Println("Connected to MongoDB to collection: ", collection.Name())

    return collection, nil
}

// CloseClient closes the MongoDB client connection
func CloseClient() {
    if client != nil {
        if err := client.Disconnect(context.Background()); err != nil {
            log.Printf("Error disconnecting from MongoDB: %v", err)
        } else {
            fmt.Println("Disconnected from MongoDB")
        }
    }
}