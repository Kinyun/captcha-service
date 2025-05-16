package routes

import (
	appModel "captcha-service/app/models"
	"captcha-service/app/server"
	"captcha-service/pkg/api/v1/captcha"
	"captcha-service/pkg/api/v1/captcha/models"
	"context"
	"github.com/labstack/echo/v4"
)

type HTTP struct {
	svc captcha.Service
}

func NewHTTP(svc captcha.Service, err *echo.Group) {
	h := HTTP{svc: svc}

	g := err.Group("")
	g.GET("/image", h.generateCaptcha)
	g.POST("/verify", h.verifyCaptcha)
}

func (h *HTTP) generateCaptcha(c echo.Context) error {
	var (
		requestID = c.Get("requestID").(string)
		request   = new(models.RequestGenerateCaptcha)
	)

	if err := c.Bind(request); err != nil {
		return server.ResponseFail(c, err)
	}

	if err := request.Validate(); err != nil {
		return server.ResponseFail(c, err)
	}

	ctx := appModel.NewContext(context.Background(), requestID)

	result, err := h.svc.GenerateCaptcha(ctx, request)
	if err != nil {
		return server.ResponseFail(c, err)
	}
	return server.ResponseOK(c, "success", result)
}

func (h *HTTP) verifyCaptcha(c echo.Context) error {
	var (
		requestID = c.Get("requestID").(string)
		clientID  = c.Param("client_id")
		request   = new(models.RequestVerifyCaptcha)
	)

	if err := c.Bind(request); err != nil {
		return server.ResponseFail(c, err)
	}

	if err := request.Validate(); err != nil {
		return server.ResponseFail(c, err)
	}

	ctx := appModel.NewContext(context.Background(), requestID)

	err := h.svc.VerifyCaptcha(ctx, clientID, request)
	if err != nil {
		return server.ResponseFail(c, err)
	}

	return server.ResponseNoContent(c, "success")
}
