package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func InitializeClient() error {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:9443",
		Password: "",
		DB:       0,
	})

	ctx := context.TODO()

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Println("error connecting to redis")
		return err
	}

	return nil
}
