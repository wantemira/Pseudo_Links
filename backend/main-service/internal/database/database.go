// Package database provides Redis database connection management
package database

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var client *redis.Client

// Init initializes Redis client with configuration from environment
func Init(log *logrus.Logger) *redis.Client {
	if err := godotenv.Load(); err != nil {
		log.Warnf("Warning: Error loading .env file: %v", err)
	}

	host := getEnv("REDIS_HOST", "localhost")
	port := getEnv("REDIS_PORT", "6379")
	password := getEnv("REDIS_PASSWORD", "")
	db, _ := strconv.Atoi(getEnv("REDIS_DB", "0"))

	log.Infof("Connecting to Redis at %s:%s", host, port)

	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       db,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Errorf("Error connecting to Redis: %v", err)
	}
	log.Infof("Redis connected successfully")
	return client

}
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// GetClient returns Redis client instance
func GetClient(logger *logrus.Logger) *redis.Client {
	if client == nil {
		client = Init(logger)
	}
	return client
}
