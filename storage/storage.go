package storage

import (
	"context"
	"fmt"
	"log"
	"mg-weather-service/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initMongoClient() *mongo.Client {
	clientOption := options.Client().ApplyURI("mongodb://localhost:19000")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	return client
}

var mongo_client *mongo.Client = initMongoClient()

func PutGeoCodingRecord(record model.GeoCoding) {
	collection := mongo_client.Database("openweather").Collection("geocoding")
	res, err := collection.InsertOne(context.Background(), record)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted document: %v \r\n", res)
}

func GetGeoCodingRecordByCity(city string) model.GeoCoding {
	collection := mongo_client.Database("openweather").Collection("geocoding")
	result := model.GeoCoding{}
	filter := bson.D{{"quered_city", city}}
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	fmt.Printf("Error: %v \r\n", err)
	return result
}
