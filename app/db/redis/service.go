package redis

import (
	"context"
	"time"
)

type RedisService interface {
	Set(ctx context.Context, action, key string, value interface{}, expiration time.Duration) (err error)
	Get(ctx context.Context, action, key string) (string, error)
	Del(ctx context.Context, action, key string) (err error)
	HMSET(ctx context.Context, action, key, index string, value map[string]interface{}, expiration time.Duration) (err error)
	HGetAll(ctx context.Context, action, key string) (data map[string]string, err error)
	HDel(ctx context.Context, action, key, index string) (err error)
}
