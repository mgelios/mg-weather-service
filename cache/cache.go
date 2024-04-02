package cache

import (
	"context"
	"log"
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
	client.Set(context.Background(), "key", "value", time.Second*5)
	println("Send record with key and value as key and value")

}

func getValue(key string) string {
	result := client.Get(context.Background(), "key")
	println(result.Val())
	return result.Val()
}
