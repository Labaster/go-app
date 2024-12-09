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

func Home(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"message": "Hello, World!"})
}

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

func GetTodos(c *fiber.Ctx) error {
    var todos []structures.Todo

    mongoConn := getConn()
    if mongoConn == nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to connect to database"})
    }

    cursor, err := mongoConn.Find(context.Background(), bson.M{})
    if err != nil {
        log.Println("GetTodos err -->> ", err)
        return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch todos"})
    }
    defer cursor.Close(context.Background())

    for cursor.Next(context.Background()) {
        var todo structures.Todo
        if err := cursor.Decode(&todo); err != nil {
            log.Println("Cursor decode err -->> ", err)
            return c.Status(500).JSON(fiber.Map{"error": "Failed to decode todo"})
        }
        todos = append(todos, todo)
    }

    if err := cursor.Err(); err != nil {
        log.Println("Cursor err -->> ", err)
        return c.Status(500).JSON(fiber.Map{"error": "Cursor error"})
    }

    return c.JSON(todos)
}

// func AddTodo(c *fiber.Ctx) error {
	
// }

// func UpdateTodo(c *fiber.Ctx) error {
	
// }

// func DeleteTodo(c *fiber.Ctx) error{
	
// }