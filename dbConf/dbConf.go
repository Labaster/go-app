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

func Connect(dbName string, collectionName string) (*mongo.Collection, error) {
	if dbName == "" || collectionName == "" {
		errMsg := "dbName and collectionName are required!"
		log.Fatal(errMsg)
		return nil, errors.New(errMsg)
	}

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	MONGO_URI := os.Getenv("MONGO_URI")

	clientOptions := options.Client().ApplyURI(MONGO_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("mongo.Connect err -->>", err)
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("client.Ping err -->>", err)
		return nil, err
	}

	collection := client.Database(dbName).Collection(collectionName)

	fmt.Println("Connected to MongoDB to collection: ", collection.Name())

	return collection, nil
}