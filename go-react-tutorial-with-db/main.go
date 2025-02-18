// package main is the entry point of the program, it is the root package
package main

// import the fmt package which is used to print to the console
import (
	// context is used to handle the context of the request
	"context"

	// fmt is used to print to the console
	"fmt"

	// log is used to print to the console
	"log"

	// os is used to get the environment variables
	"os"

	// fiber is a web framework for Go, similar to express.js in Node.js for the backend
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	// filesystem is used to serve the static files

	// ui is used to serve the static files

	// godotenv is used to load the environment variables from the .env file
	"github.com/joho/godotenv"

	// bson is a binary JSON format used for MongoDB, similar to JSON but more efficient
	"go.mongodb.org/mongo-driver/bson"

	// primitive is used to represent the primitive types in MongoDB
	"go.mongodb.org/mongo-driver/bson/primitive"

	// mongo is used to connect to the MongoDB database
	"go.mongodb.org/mongo-driver/mongo"

	// mongo/options is used to connect to the MongoDB database
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Structs are used to create custom types in Go

// The todo struct is used to represent a todo in the database
type Todo struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool `json:"completed"`
	Body string `json:"body"`
}

var collection *mongo.Collection

// main function is the entry point of the program
func main() {

	// fmt.Println is used to print to the console
	fmt.Println("Hello, World!")

	if os.Getenv("ENV") != "production" {
		// Load the environment variables
		err := godotenv.Load(".env")

		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// load the environment variables
	err := godotenv.Load(".env")

	// if the environment variables are not loaded, print an error
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get the port from the environment variables
	MONGODB_URI := os.Getenv("MONGODB_URI")

	// create the client options which are used to connect to the database
	clientOptions := options.Client().ApplyURI(MONGODB_URI)

	// connect to the database using the client options
	// context.Background() is used to create a new context that is not cancelled for the duration of the connection
	// think of it as a container that holds the connection to the database
	client, err := mongo.Connect(context.Background(), clientOptions)

	// if there is an error, print the error
	if err != nil {
		log.Fatal(err)
	}

	// When the main function is finished, disconnect from the database
	defer client.Disconnect(context.Background())

	// check if the connection to the database is successful
	err = client.Ping(context.Background(), nil)

	// if there is an error, print the error
	if err != nil {
		log.Fatal(err)
	}

	// print a message to the console if the connection to the database is successful
	fmt.Println("Connected to MongoDB ATLAS!")

	// get the collection from the database
	collection = client.Database("golang_db").Collection("todos")
	
	// create a new fiber app
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// index, err := fs.Sub(Index, "dist")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// app.Use("/", filesystem.New(filesystem.Config{
	// 	Root: http.FS(index),
	// 	Index: "index.html",
	// 	Browse: false,
	// }))

	// get the todos from the collection
	app.Get("/api/todos", getTodos)

	// create a new todo
	app.Post("/api/todos", createTodo)

	// update a todo
	app.Patch("/api/todos/:id", updateTodo)

	// delete a todo
	app.Delete("/api/todos/:id", deleteTodo)

	// get the port from the environment variables
	port := os.Getenv("PORT")

	// if the port is not set, set it to 5000
	if port == "" {
		port = "5000"
	}

	if os.Getenv("ENV") == "production" {
		app.Static("/", "../ui/dist")
	}

	// listen to the port
	log.Fatal(app.Listen("0.0.0.0:" + port))
}

// getTodos returns all the todos in the collection
func getTodos(c *fiber.Ctx) error {
	var todos []Todo

	// find all the todos in the collection, the bson.M{} is the filter, which is empty in this case
	// the end result is a cursor, which is used to iterate over the todos
	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return err
	}

	// defer is used to close the cursor after the function is finished
	defer cursor.Close(context.Background())

	// iterate over the todos and decode them into the Todo struct
	for cursor.Next(context.Background()) {
		// decode the todo into the Todo struct
		var todo Todo

		// if there is an error, return the error	
		if err := cursor.Decode(&todo); err != nil {
			return err
		}

		todos = append(todos, todo)
	}
	return c.JSON(todos)
}

// createTodo creates a new todo in the collection
func createTodo(c *fiber.Ctx) error {
	todo := new(Todo)

	if err := c.BodyParser(&todo); err != nil {
		return err
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Body is required"})
	}

	insertResult, err := collection.InsertOne(context.Background(), todo)

	if err != nil {
		return err
	}

	todo.ID = insertResult.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(todo)
}

// updateTodo updates a todo in the collection
func updateTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": true}}

	_, err = collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"message": "Todo updated"})
}

// deleteTodo deletes a todo in the collection
func deleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	} 
	
	filter := bson.M{"_id": objectID}

	_, err = collection.DeleteOne(context.Background(), filter)
	
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"message": "Todo deleted"})
}