package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Redis struct{
	
}


func InitializeClient() (*redis.Client, error) {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.TODO()

	status, err := redisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Println("error connecting to redis")
		return nil, err
	}
	fmt.Print(status, "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
	return redisClient, nil
}
