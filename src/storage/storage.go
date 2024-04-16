package storage

import (
	"context"
	"fmt"
	"log"
	"mg-weather-service/src/model"
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
	client.Database("openweather").CreateCollection(ctx, "geocoding")
	if err != nil {
		fmt.Println("Error while creating geocoding")
	}
	client.Database("openweather").CreateCollection(ctx, "onecall")
	if err != nil {
		fmt.Println("Error while creating onecall")
	}
	return client
}

// func disconnectMongoClient() {
// 	if err = mongo_client.Disconnect(ctx); err != nil {
// 		panic(err)
// 	}
// }

var mongo_client *mongo.Client = initMongoClient()

func UpdateGeoCodingRecord(id string, record model.GeoCoding) {
	collection := mongo_client.Database("openweather").Collection("geocoding")
	_, err := collection.UpdateByID(context.Background(), id, record)
	if err != nil {
		log.Fatal(err)
	}
}

func PutGeoCodingRecord(record model.GeoCoding) {
	collection := mongo_client.Database("openweather").Collection("geocoding")
	_, err := collection.InsertOne(context.Background(), record)
	if err != nil {
		log.Fatal(err)
	}
}

func GetGeoCodingRecordByCity(city string) (model.GeoCoding, error) {
	collection := mongo_client.Database("openweather").Collection("geocoding")
	result := model.GeoCoding{}
	filter := bson.D{{"queredcity", city}}
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	return result, err
}

func UpdateOneCallRecord(id string, record model.OneCallWeather) {
	collection := mongo_client.Database("openweather").Collection("onecall")
	_, err := collection.UpdateByID(context.Background(), id, record)
	if err != nil {
		log.Fatal(err)
	}
}

func PutOneCallRecord(record model.OneCallWeather) {
	collection := mongo_client.Database("openweather").Collection("onecall")
	_, err := collection.InsertOne(context.Background(), record)
	if err != nil {
		log.Fatal(err)
	}
}

func GetOneCallRecordByCity(city string) (model.OneCallWeather, error) {
	collection := mongo_client.Database("openweather").Collection("onecall")
	result := model.OneCallWeather{}
	filter := bson.D{{"city", city}}
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	return result, err
}
