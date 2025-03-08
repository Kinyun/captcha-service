package logging

import (
	"captcha-service/app/config/constant"
	"captcha-service/app/logger/singleton"
	"captcha-service/app/server/request"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"time"
)

func Logging() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqID := request.ID()
			c.Set("requestID", reqID)
			defer func(now time.Time) {
				message := constant.LLvlAccess
				fields := []zap.Field{
					zap.String("at", now.Format(constant.DateFormatWithTime)),
					zap.String("method", c.Request().Method),
					zap.Int("code", c.Response().Status),
					zap.String("uri", c.Request().URL.String()),
					zap.String("ip", c.RealIP()),
					zap.String("host", c.Request().Host),
					zap.String("user_agent", c.Request().UserAgent()),
					zap.String("header", fmt.Sprintf("%s", c.Request().Header)),
				}
				singleton.WithRequestID(reqID).Info(message, fields...)
			}(time.Now())
			return next(c)
		}
	}
}
