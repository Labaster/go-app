package routeActions

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Labaster/go-app/dbConf"
	"github.com/Labaster/go-app/structures"
)

var mongoConnInst *mongo.Collection = nil

func init() {
    mongoConnInst = getConn()
}

func getConn() *mongo.Collection {
    if mongoConnInst != nil {
        return mongoConnInst
    }
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

    if mongoConnInst == nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to connect to database"})
    }

    cursor, err := mongoConnInst.Find(context.Background(), bson.M{})
    if err != nil {
        log.Println("GetTodos err -->> ", err)
        return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch todos"})
    }

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

    defer cursor.Close(context.Background())

    return c.JSON(todos)
}

func AddTodo(c *fiber.Ctx) error {
    todo := new(structures.Todo) // -->> todo := &structures.Todo{}

    if err := c.BodyParser(todo); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": err.Error(), 
            "message": "Failed to parse request body",
        })
    }

    // todo.Id = primitive.NewObjectID()
    insertResult, err := mongoConnInst.InsertOne(context.Background(), todo)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to insert todo"})
    }

    todo.Id = insertResult.InsertedID.(primitive.ObjectID)

    fmt.Println("Inserted todo with ID: ", todo.Id)

    return c.Status(201).JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
    objectId, err := primitive.ObjectIDFromHex(id)

    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
    }

    filter := bson.M{"_id": objectId}
    update := bson.M{"$set": bson.M{"completed": true}}

    _, err = mongoConnInst.UpdateOne(context.Background(), filter, update)

    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to update todo"})
    }

    return c.Status(200).JSON(fiber.Map{"message": "Todo updated"})
}

func DeleteTodo(c *fiber.Ctx) error{
	id := c.Params("id")
    objectId, err := primitive.ObjectIDFromHex(id)

    if err != nil { 
        return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
    }

    _, err = mongoConnInst.DeleteOne(context.Background(), bson.M{"_id": objectId})

    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to delete todo"})
    }

    return c.Status(200).JSON(fiber.Map{"message": "Todo deleted"})
}