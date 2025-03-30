package usecase

import (
	"captcha-service/app/config/constant"
	"captcha-service/app/utils/steambap/impl"
	"captcha-service/pkg/api/v1/captcha/models"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

func (capt *Captcha) GenerateCaptcha(ctx context.Context, request *models.RequestGenerateCaptcha) (models.ResponseGenerateCaptcha, error) {

	expiredTimeRedis := time.Duration(10 * time.Minute)
	attributeCaptcha := impl.AttributeSteambap{
		Width:  request.Width,
		Height: request.Height,
		Length: request.Length,
		Noise:  request.Noise,
	}

	captcha, err := capt.repository.steambap.GenerateCaptcha(attributeCaptcha)
	if err != nil {
		//log.Printf("err : %v", err)
		return models.ResponseGenerateCaptcha{}, err
	}

	responseByte, _ := json.Marshal(captcha)

	err = capt.repository.redisRepo.Set(ctx, "set result captcha to redis", fmt.Sprintf(captcha.CaptchaID), responseByte, expiredTimeRedis)
	if err != nil {
		return models.ResponseGenerateCaptcha{}, constant.ErrInternal
	}

	return models.ResponseGenerateCaptcha{
		CaptchaID:          captcha.CaptchaID,
		CaptchaImage:       fmt.Sprintf("data:image/jpeg;base64,%s", captcha.CaptchaImage),
		CaptchaExpiredTime: fmt.Sprintf("%v", expiredTimeRedis),
	}, nil
}

func (capt *Captcha) VerifyCaptcha(ctx context.Context, clientID string, request *models.RequestVerifyCaptcha) error {

	return nil
}
