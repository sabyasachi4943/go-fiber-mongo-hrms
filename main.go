package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client
	Db
}

var mg MongoInstance

const dbName = "fiber-hrms"
const mongoURI = "mongodb://localhost:27017" + dbName

type Employee struct {
	ID string
	Nmae string
	Salary float64
	Age float64
}

func Connect() error{
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	context.WithTimeout(context.Background(), 30*time.Second)
}

func main(){

	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/employee", func(c *fiber.Ctx) error{

	})
	app.Post("/employee")
	app.Put("/employee/:id")
	app.Delete("/employee/:id")

}