// Main package is the entry point for the application
package main

// fmt is a package that contains functions for formatting input and output, as well as scanning input, and printing output
import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	// Format is: ID int `json:"id"` 
	// This is a struct tag, it is used to specify the JSON key name for the field
	// ID is the value of the field, and int is the type of the field, and json is the key name in the JSON
	// So the format is: field_name type `json:"key_name"`
	ID int `json:"id"`
	Completed bool `json:"completed"`
	Body string `json:"body"`
}

func main(){
	// Println is a function that prints a string to the console
	fmt.Println("Hello, World!!")

	// App is a function that creates a new Fiber app, which is a web framework for Go that is used to create APIs
	// Fiber is similar to Express.js in Node.js
	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	todos := []Todo{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Hello, World!"})
	})

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	// Create a new todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		// & is used to create a pointer to the todo struct
		// A pointer is a variable that stores the memory address of another variable
		// So todo is a pointer to a Todo struct
		// &Todo{} is a pointer to a Todo struct
		todo := &Todo{} // {id: 0, completed: false, body: ""}

		// BodyParser is a function that parses the body of the request and unmarshals it into the todo struct
		if err := c.BodyParser(todo); err != nil {
			return err
		}

		// If the todo body is empty, return a 400 status code and a message
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
		}

		// Set the ID of the todo
		// len(todos) is the length of the todos array
		// +1 is the next number in the sequence
		todo.ID = len(todos) + 1

		// Append the todo to the todos array
		todos = append(todos, *todo)

		// Pointers explained:
		// var x int =  5 // 0x14000112000 (memory address)
		// var pointer *int = &x // 0x14000112000 (memory address is the same as x)
		// fmt.Println(pointer) // 0x14000112000
		// fmt.Println(*pointer) // 5 (* is used to get the value of the variable at the memory address)

		// Return a 201 status code and the todo
		return c.Status(201).JSON(todo)

	})

	// Update a todo
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		// Go For Loop:
		// Note: Go does not have a while loop, it only has a for loop
		for i, todo := range todos {
			// fmt.Sprint(todo.ID) is used to convert the ID to a string because id is a string
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})

	// Delete a todo
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			// fmt.Sprint(todo.ID) is used to convert the ID to a string because id is a string
			if fmt.Sprint(todo.ID) == id {
				// append(todos[:i], todos[i+1:]...) is used to remove the todo from the todos array
				// 1 2 3 4 5
				// Say we want to remove 3
				// append(todos[:i], todos[i+1:]...) will remove 3 because it takes the elements before index i and the elements after index i+1
				// So it will be 1 2 4 5
				// This is a way to remove an element from an array in Go
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"message": "Todo deleted"})
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})

	// Listen is a function that listens for incoming requests on the specified port
	// log.Fatal is a function that logs an error and then terminates the program
	// ":4000" is the port number that the server will listen on
	// The app is listening for incoming requests on port 4000, and has logged the error if it occurs
	log.Fatal(app.Listen(":" + PORT))
}