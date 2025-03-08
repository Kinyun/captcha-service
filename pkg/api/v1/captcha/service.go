package captcha

import (
	"captcha-service/pkg/api/v1/captcha/models"
	"context"
)

type Service interface {
	GenerateCaptcha(ctx context.Context, request *models.RequestGenerateCaptcha) (models.ResponseGenerateCaptcha, error)
	VerifyCaptcha(ctx context.Context, clientID string, request *models.RequestVerifyCaptcha) error
}
