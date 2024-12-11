package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/Labaster/go-app/dbConf"
	"github.com/Labaster/go-app/routeActions"
)

func main() {
	fmt.Println("App started!")
	defer dbConf.CloseClient()

	PORT := os.Getenv("PORT")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/", routeActions.Home)
	app.Get("/api/todos", routeActions.GetTodos)
	app.Post("/api/addTodo", routeActions.AddTodo)
	app.Patch("/api/updateTodo/:id", routeActions.UpdateTodo)
	app.Delete("/api/deleteTodo/:id", routeActions.DeleteTodo)

	serverErr := app.Listen(":" + PORT)
	if serverErr != nil {
		log.Fatal(serverErr)
	}
}

/** LOCAL SAVES */

// type Todo struct {
// 	Id int `json:"id"`
// 	Completed bool `json:"completed"`
// 	Body string `json:"body"`
// }

// var todos = []Todo{}

// func Home(c *fiber.Ctx) error {
// 	return c.Status(200).JSON(fiber.Map{"message": "Hello, World!"})
// }

// func GetTodos(c *fiber.Ctx) error {
// 	return c.Status(200).JSON(fiber.Map{"todos": todos})
// }

// func AddTodo(c *fiber.Ctx) error {
// 	todo := &Todo{}

// 	if err := c.BodyParser(todo); err != nil {
// 		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	if todo.Body == "" {
// 		return c.Status(400).JSON(fiber.Map{"error": "Body is required"})
// 	}

// 	todo.Id = len(todos) + 1
// 	todos = append(todos, *todo)

// 	return c.Status(201).JSON(todo)
// }

// func UpdateTodo(c *fiber.Ctx) error {
// 	var id string = c.Params("id")

// 	for indx, todo := range todos {
// 		if fmt.Sprint(todo.Id) == id {
// 			todos[indx].Completed = true
// 			return c.Status(200).JSON(todos[indx])
// 		}
// 	}

// 	return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
// }

// func DeleteTodo(c *fiber.Ctx) error{
// 	var id string = c.Params("id")

// 	for indx, todo := range todos {
// 		if fmt.Sprint(todo.Id) == id {
// 			todos = append(todos[:indx], todos[indx+1:]...)
// 			return c.Status(200).JSON(fiber.Map{"deleted by ID": indx})
// 		}
// 	}

// 	return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
// }

// func main() {
// 	fmt.Println("App started!")
// 	// var myName string= "John Doe"
// 	// const myAge int = 30
// 	// mySurName := "Doe"
// 	// fmt.Printf("My name is %v %v and age %v", myName, mySurName, myAge)

// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	PORT := os.Getenv("PORT")

// 	app := fiber.New()

// 	app.Get("/", Home)
// 	app.Get("/api/todos", GetTodos)
// 	app.Post("/api/addTodo", AddTodo)
// 	app.Patch("/api/updateTodo/:id", UpdateTodo)
// 	app.Delete("/api/deleteTodo/:id", DeleteTodo)

// 	serverErr := app.Listen(":" + PORT)
// 	if serverErr != nil {
// 		log.Fatal(serverErr)
// 	}
// }