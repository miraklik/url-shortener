package store

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
)

const (
	CacheDuration = 6 * time.Hour
)

func NewStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		log.Panicf("Error init Redis: %v", err)
	}

	fmt.Printf("\nRedis started successfully: pong = {%s}", pong)
	storeService.redisClient = redisClient

	return storeService
}

func SaveURLMapping(shortUrl, originalUrl, userID string) {
	if err := storeService.redisClient.Set(shortUrl, originalUrl, CacheDuration).Err(); err != nil {
		log.Panicf("Failed saving key url | ERR: %v - shortURL: %s - originalURL: %s\n", err, shortUrl, originalUrl)
	}
}

func RetrieveInitialUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(shortUrl).Result()
	if err != nil {
		log.Panicf("Failed to get url | ERR: %v - shortURL: %s", err, shortUrl)
	}

	return result
}
