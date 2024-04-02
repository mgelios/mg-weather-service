package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

var mongo_client *mongo.Client

func init() {
	clientOption := options.Client().ApplyURI("mongodb://localhost:19000")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	mongo_client, err := mongo.Connect(ctx, clientOption)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = mongo_client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func putRecord() {
	collection := mongo_client.Database("test").Collection("people")
	person := Person{Name: "John", Age: 30, Email: "john@example.com"}
	res, err := collection.InsertOne(context.Background(), person)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted document: %v \r\n", res)

}

func getRecrod() {
	collection := mongo_client.Database("test").Collection("people")
	result := Person{}
	filter := bson.D{{"name", "John"}}
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	fmt.Printf("Error: %v \r\n", err)
	fmt.Printf("Value: %v \r\n", result)
}
