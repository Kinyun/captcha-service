package api

import (
	"captcha-service/app/config"
	"captcha-service/app/server"
	"captcha-service/pkg/api/v1/captcha/routes"
	"captcha-service/pkg/api/v1/captcha/usecase"
	"github.com/labstack/echo/v4"
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

func Start() {

	//// Load ENV from file
	//err := godotenv.Load()
	//if err != nil {
	//	fmt.Print("unable to load .env file: ", err)
	//}

	// Init the Echo Framework
	e := server.InitEcho()

	// Init for routing group
	captchaV1 := e.Group("/v1")
	captchaV1.Use(maintenanceMode)

	routes.NewHTTP(usecase.Initialize(), captchaV1)

	// Start the server
	server.Start(e)
}
