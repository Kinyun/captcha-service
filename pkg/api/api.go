package api

import (
	"captcha-service/app/config"
	redis2 "captcha-service/app/db/redis"
	"captcha-service/app/server"
	"captcha-service/pkg/api/v1/captcha/routes"
	"captcha-service/pkg/api/v1/captcha/usecase"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"log"
)

// function as the maintenance switcher
func maintenanceMode(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if config.GetConfig().UnderMaintenance {
			return server.ResponseStatusServiceUnavailable(c, "service is under maintenance", nil)
		}
		return next(c)
	}
}

func initDatabase() *redis.Client {

	redisConn, err := redis2.NewConnectionRedis()
	if err != nil {
		log.Fatal(err)
	}

	return redisConn
}

func Start() {

	//// Load ENV from file
	//err := godotenv.Load()
	//if err != nil {
	//	fmt.Print("unable to load .env file: ", err)
	//}

	redisConn := initDatabase()

	// Init the Echo Framework
	e := server.InitEcho()

	// Init for routing group
	captchaV1 := e.Group("/v1")
	captchaV1.Use(maintenanceMode)

	routes.NewHTTP(usecase.Initialize(redisConn), captchaV1)

	// Start the server
	server.Start(e)
}
