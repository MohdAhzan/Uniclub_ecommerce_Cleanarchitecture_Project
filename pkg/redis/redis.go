package redis

import (
	"context"
	"fmt"

	"github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/config"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
}

func InitializeClient(cfg config.Config) (*redis.Client, error) {

  redisAddr:=fmt.Sprintf("%s:%s",cfg.REDIS_HOST,cfg.REDIS_PORT)
  fmt.Println("redisADDRESS IN ENV",redisAddr)
	redisClient := redis.NewClient(&redis.Options{
		Addr:    redisAddr,
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
