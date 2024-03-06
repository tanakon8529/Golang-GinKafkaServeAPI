package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"ginapi-gateway/settings"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// GetRedisClient returns a Redis client based on the server name
func GetRedisClient() *redis.Client {
	// Load environment variables
	config := settings.LoadEnv(".env")
	host := config.RedisHost
	port := config.RedisPort
	password := config.RedisPassword

	// Create a Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       0,
	})

	// Declare the err variable to capture errors from Ping
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("Failed to connect to Redis: %v\n", err)
		os.Exit(1)
	}

	return redisClient
}

// ValidAccessToken checks if the access token is valid
func ValidAccessToken(token string) error {
	redisClient := GetRedisClient()
	log.Printf("ValidAccessToken: token: %s\n", token)

	_, err := redisClient.Get(ctx, token).Result()
	if err == redis.Nil {
		log.Printf("ValidAccessToken: invalid access token: %s\n", token)
		return fmt.Errorf("invalid access token")
	} else if err != nil {
		log.Printf("ValidAccessToken: redis error: %v\n", err)
		return fmt.Errorf("redis error: %v", err)
	}

	return nil
}

// StoreAccessToken stores an access token in Redis
// example tokenExpireMinutes : if need to expire in 1 hour, tokenExpireMinutes = 60
func StoreAccessToken(accessToken string, tokenExpireMinutes time.Duration) error {
	redisClient := GetRedisClient()
	log.Printf("StoreAccessToken: accessToken: %s\n", accessToken)

	// Store the access token in Redis
	err := redisClient.Set(ctx, accessToken, "valid", tokenExpireMinutes*time.Minute).Err()
	if err != nil {
		log.Printf("StoreAccessToken: redis error: %v\n", err)
		return fmt.Errorf("redis error: %v", err)
	}
	return nil
}
