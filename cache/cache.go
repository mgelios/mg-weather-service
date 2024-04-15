package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mg-weather-service/model"
	"time"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	err := client.Ping(context.Background()).Err()

	if err != nil {
		log.Fatal("Failed to connect with dragonfly", err)
	}
}

func putValue(key string, value string) {
	client.Set(context.Background(), "key", "value", time.Minute*2)
	println("Send record with key and value as key and value")

}

func getValue(key string) string {
	result := client.Get(context.Background(), "key")
	println(result.Val())
	return result.Val()
}

func putGeoCoding(city string, value model.GeoCoding) {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	client.Set(context.Background(), city, jsonValue, time.Minute*2)
	fmt.Println("cache value geocode was set")
}

func getGeoCodingByCity(city string) model.GeoCoding {
	rawResult := client.Get(context.Background(), city)
	var result model.GeoCoding
	if err := json.Unmarshal([]byte(rawResult.Val()), &result); err != nil {
		panic(err)
	}
	fmt.Println("fetched geocoding from cache")
	return result
}

func putOneCall(city string, value model.OneCallWeather) {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	client.Set(context.Background(), city, jsonValue, time.Minute*2)
	fmt.Println("cache value onecall was set")
}

func getOneCallByCity(city string) model.OneCallWeather {
	rawResult := client.Get(context.Background(), city)
	var result model.OneCallWeather
	if err := json.Unmarshal([]byte(rawResult.Val()), &result); err != nil {
		panic(err)
	}
	fmt.Println("fetched onecall from cache")
	return result
}
