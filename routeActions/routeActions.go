package routeActions

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Labaster/go-app/dbConf"
	"github.com/Labaster/go-app/structures"
)

var mongoConn = getConn()

func getConn() *mongo.Collection {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	MONGO_DB_NAME := os.Getenv("MONGO_DB_NAME")
	MONGO_COLLECTION_NAME := os.Getenv("MONGO_COLLECTION_NAME")


	mongoConn, err := dbConf.Connect(MONGO_DB_NAME, MONGO_COLLECTION_NAME)

	if err != nil {
		log.Fatal("getConn err -->>", err)
		return nil
	}

	return mongoConn
}

func Home(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"message": "Hello, World!"})
}

func GetTodos(c *fiber.Ctx) error {
	var todos []structures.Todo

	cursor, err := mongoConn.Find(context.Background(), bson.M{})

	defer cursor.Close(context.Background())

	if err != nil {
		log.Fatal("GetTodos err -->>", err)
		return c.Status(500).JSON(fiber.Map{"err": err})
	}

	for cursor.Next(context.Background()) {
		var todo structures.Todo
		if err := cursor.Decode(&todo); err != nil {
			log.Fatal("GetTodos cursor.Next err -->>", err)
			return c.Status(500).JSON(fiber.Map{"err": err})
		}
		todos = append(todos, todo)
	}

	return c.Status(200).JSON(todos)
}

// func AddTodo(c *fiber.Ctx) error {
	
// }

// func UpdateTodo(c *fiber.Ctx) error {
	
// }

// func DeleteTodo(c *fiber.Ctx) error{
	
// }