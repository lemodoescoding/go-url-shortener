package store

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx = context.Background()
)

const cacheDuration = 6 * time.Hour

func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}

func SaveURLMapping(shortURL string, originalURL string, userID string) {
	err := storeService.redisClient.Set(ctx, shortURL, originalURL, cacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortURL: %s - originalURL: %s", err, shortURL, originalURL))
	}
}

func RetreiveInitialURL(shortURL string) (string, error) {
	result, err := storeService.redisClient.Get(ctx, shortURL).Result()
	if err == redis.Nil {
		return "", err
	} else if err != nil  {
		return "", nil
	}
	// if err != nil {
	// 	panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortURL))
	// }
	
	return result, nil
}
