package usecase

import (
	redis2 "captcha-service/app/db/redis"
	impl2 "captcha-service/app/db/redis/impl"
	"captcha-service/app/utils/steambap"
	"captcha-service/app/utils/steambap/impl"
	"github.com/redis/go-redis/v9"
)

type Connection struct {
	redis *redis.Client
}

type Repository struct {
	steambap  steambap.SteambapCaptcha
	redisRepo redis2.RedisService
}

type Captcha struct {
	conn       Connection
	repository Repository
}

func NewCaptchaRepository(redisConn *redis.Client) Repository {
	return Repository{
		steambap:  impl.NewSteambapCaptcha(),
		redisRepo: impl2.NewDB(redisConn),
	}
}

func New(redisConn *redis.Client, repository Repository) *Captcha {
	return &Captcha{
		conn: Connection{
			redis: redisConn,
		},
		repository: repository,
	}
}

func Initialize(redisConn *redis.Client) *Captcha {
	return New(redisConn, NewCaptchaRepository(redisConn))
}
