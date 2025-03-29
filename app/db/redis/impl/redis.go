package impl

import (
	"captcha-service/app/config/constant"
	"captcha-service/app/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	weekHour = 168
)

// RDB represents RDB struct
type RDB struct {
	client *redis.Client
}

// NewDB returns a new Redis DB instance
func NewDB(c *redis.Client) *RDB {
	return &RDB{client: c}
}

type options struct {
	method  string
	action  string
	expired time.Duration
}

func internalLog(ctx context.Context, key string, value []byte, opt options, err error) {
	status := constant.Success
	if err != nil {
		status = constant.Error
	}
	info := fmt.Sprintf("%s %s", opt.action, key)
	if opt.expired > 0 {
		info = fmt.Sprintf("%s %s: %d", opt.action, key, opt.expired)
	}
	models.RedisLog(ctx, time.Now(), opt.method, status, info, value, err)
}

func (r *RDB) Set(ctx context.Context, action, key string, value interface{}, expiration time.Duration) (err error) {
	defer func() {
		switch value.(type) {
		case string:
			internalLog(ctx, key, []byte(value.(string)), options{method: "Set", action: action, expired: expiration}, err)
		case bool:
			if value == true {
				internalLog(ctx, key, []byte("1"), options{method: "Set", action: action, expired: expiration}, err)
			} else {
				internalLog(ctx, key, []byte("0"), options{method: "Set", action: action, expired: expiration}, err)
			}
		}
	}()

	if expiration == 0 {
		expiration = weekHour * time.Hour
	}

	_, err = r.client.Set(ctx, key, value, expiration).Result()
	return
}

func (r *RDB) Get(ctx context.Context, action, key string) (data string, err error) {
	defer func() {
		internalLog(ctx, key, []byte(data), options{method: "Get", action: action}, err)
	}()

	data, err = r.client.Get(ctx, key).Result()
	return
}

func (r *RDB) Del(ctx context.Context, action, key string) (err error) {
	defer func() {
		internalLog(ctx, key, nil, options{method: "Del", action: action}, err)
	}()

	_, err = r.client.Del(ctx, key).Result()
	return
}

func (r *RDB) HMSET(ctx context.Context, action, key, index string, value map[string]interface{}, expiration time.Duration) (err error) {
	defer func() {
		internalLog(ctx, key, value[index].([]byte), options{method: "HMSET", action: action}, err)
	}()

	if expiration == 0 {
		expiration = 36 * time.Hour
	}

	if _, err = r.client.HMSet(ctx, key, value).Result(); err != nil {
		return
	}
	_, err = r.client.Expire(ctx, key, expiration).Result()
	return
}

func (r *RDB) HGetAll(ctx context.Context, action, key string) (data map[string]string, err error) {
	defer func() {
		byteData, _ := json.Marshal(data)
		internalLog(ctx, key, json.RawMessage(string(byteData)), options{method: "HGetAll", action: action}, err)
	}()

	data, err = r.client.HGetAll(ctx, key).Result()
	return
}

func (r *RDB) HDel(ctx context.Context, action, key, index string) (err error) {
	defer func() {
		internalLog(ctx, key, nil, options{method: "HDel", action: action}, err)
	}()

	_, err = r.client.HDel(ctx, key, index).Result()
	return
}
