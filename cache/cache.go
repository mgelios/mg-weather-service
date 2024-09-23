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

func InitCacheClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:17000",
	})
	err := client.Ping(context.Background()).Err()
	if err != nil {
		log.Fatal("Failed to connect with dragonfly", err)
	}
	return client
}

var cacheClient *redis.Client = InitCacheClient()

func PutGeoCoding(city string, value model.GeoCoding) {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	cacheClient.Set(context.Background(), city, jsonValue, time.Minute*2)
	fmt.Println("cache value geocode was set")
}

func GetGeoCodingByCity(city string) (model.GeoCoding, error) {
	rawResult, err := cacheClient.Get(context.Background(), city).Result()
	var result model.GeoCoding
	if err == nil {
		if err := json.Unmarshal([]byte(rawResult), &result); err != nil {
			panic(err)
		}
		fmt.Println("fetched geocoding from cache")
	}
	return result, err
}

func PutOneCall(city string, value model.OneCallWeather) {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	cacheClient.Set(context.Background(), city, jsonValue, time.Minute*2)
	fmt.Println("cache value onecall was set")
}

func GetOneCallByCity(city string) (model.OneCallWeather, error) {
	rawResult, err := cacheClient.Get(context.Background(), city).Result()
	var result model.OneCallWeather
	if err == nil {
		if err := json.Unmarshal([]byte(rawResult), &result); err != nil {
			panic(err)
		}
		fmt.Println("fetched onecall from cache")
	}
	return result, err
}
