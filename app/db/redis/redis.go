package redis

import (
	"captcha-service/app/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"sync"
)

var (
	dbConn *redis.Client
	lockDB sync.Mutex
)

func GetConnection() *redis.Client {
	if dbConn == nil {
		lockDB.Lock()
		defer lockDB.Unlock()
		if dbConn == nil {
			var err error
			dbConn, err = connectRedisDB()
			if err != nil {
				log.Fatalf("Failed to connect to Redis: %v", err)
			}
		}
	}
	return dbConn
}

func connectRedisDB() (*redis.Client, error) {
	db, _ := strconv.Atoi(config.GetConfig().RedisDb)
	maxRetries, _ := strconv.Atoi(config.GetConfig().RedisMaxRetries)
	maxIdleConn, _ := strconv.Atoi(config.GetConfig().RedisMaxIdleConnections)
	client := redis.NewClient(&redis.Options{
		Addr:         config.GetConfig().RedisAddress,
		Password:     config.GetConfig().RedisPassword,
		DB:           db,
		MaxRetries:   maxRetries,
		MaxIdleConns: maxIdleConn,
		ReadTimeout:  5,
		WriteTimeout: 5,
	})
	// Ping to check connection
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Panicf("got an error while connecting redis server, error: %s", err)

	}
	fmt.Println("Pong:", pong)

	return client, nil
}
