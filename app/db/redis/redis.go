package redis

import (
	"captcha-service/app/config"
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"sync"
	"time"
)

var (
	dbConn *redis.Client
	lockDB sync.Mutex
)

func GetConnection() *redis.Client {
	if dbConn == nil {
		lockDB.Lock()
		defer lockDB.Unlock()
		dbConn = connectRedisDB()
	}
	return dbConn
}

func connectRedisDB() *redis.Client {
	db, _ := strconv.Atoi(config.GetConfig().RedisDb)
	maxRetries, _ := strconv.Atoi(config.GetConfig().RedisMaxRetries)
	maxIdleConn, _ := strconv.Atoi(config.GetConfig().RedisMaxIdleConnections)
	opt := &redis.Options{
		Addr:         config.GetConfig().RedisAddress,
		Password:     config.GetConfig().RedisPassword,
		DB:           db,
		MaxIdleConns: maxIdleConn,
		MaxRetries:   maxRetries,
		ReadTimeout:  time.Minute * 1,
		WriteTimeout: time.Minute * 3,
	}
	redisClient := redis.NewClient(opt)
	// Ping to check connection
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Panicf("got an error while connecting redis server, error: %s", err)

	}

	return redisClient
}

func NewConnectionRedis() (*redis.Client, error) {
	redisConn := connectRedisDB()
	return redisConn, nil
}
