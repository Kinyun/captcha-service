package config

import (
	"github.com/caarlos0/env"
	"log"
	"sync"
)

type Configurations struct {
	UnderMaintenance        bool   `env:"UNDER_MAINTENANCE"`
	HTTPServerTimeOut       int    `env:"HTTP_SERVER_TIMEOUT"`
	HTTPPort                string `env:"HTTP_PORT"`
	JwtKey                  string `env:"JWT_KEY"`
	RedisAddress            string `env:"REDIS_ADDRESS"`
	RedisPassword           string `env:"REDIS_PASSWORD"`
	RedisDb                 string `env:"REDIS_DB"`
	RedisMaxIdleConnections string `env:"REDIS_MAX_IDLE"`
	RedisMaxRetries         string `env:"REDIS_MAX_RETRIES"`
	ChatBotUsername         string `env:"CHAT_BOT_USERNAME"`
	ChatBotPassword         string `env:"CHAT_BOT_PASSWORD"`
}

var (
	configuration Configurations
	mutex         sync.Once
)

func GetConfig() Configurations {
	mutex.Do(func() {
		configuration = newConfig()
	})

	return configuration
}

func newConfig() Configurations {
	var cfg = Configurations{}
	if err := env.Parse(&cfg); err != nil {
		log.Panic(err.Error())
	}

	return cfg
}
