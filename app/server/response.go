package server

import (
	"captcha-service/app/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ResponseOK(c echo.Context, msg string, r interface{}) error {
	return c.JSON(http.StatusOK, models.CommonResponse{
		Error:   false,
		Message: msg,
		Data:    r,
	})
}

func ResponseNoContent(c echo.Context, msg string) error {
	return c.JSON(http.StatusNoContent, models.CommonResponse{
		Error:   false,
		Message: msg,
	})
}

func ResponseBadRequest(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, models.CommonResponse{
		Error:   true,
		Message: err.Error(),
	})
}

func ResponseNotFound(c echo.Context, err error) error {
	return c.JSON(http.StatusNotFound, models.CommonResponse{
		Error:   true,
		Message: err.Error(),
	})
}

func ResponseFail(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, models.CommonResponse{
		Error:   true,
		Message: err.Error(),
	})
}

func ResponseStatusServiceUnavailable(c echo.Context, msg string, data interface{}) error {
	return c.JSON(http.StatusServiceUnavailable, models.CommonResponse{
		Error:   true,
		Status:  http.StatusServiceUnavailable,
		Data:    data,
		Message: msg,
	})
}
