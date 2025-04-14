package usecase

import (
	"captcha-service/app/config/constant"
	"captcha-service/app/logger/singleton"
	steambap "captcha-service/app/utils/steambap/model"
	"captcha-service/pkg/api/v1/captcha/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"time"
)

func (capt *Captcha) GenerateCaptcha(ctx context.Context, request *models.RequestGenerateCaptcha) (models.ResponseGenerateCaptcha, error) {

	expiredTimeRedis := time.Duration(10 * time.Minute)
	attributeCaptcha := steambap.AttributeSteambap{
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
	redisCaptcha := models.RedisCaptcha{
		CaptchaID:    captcha.CaptchaID,
		CaptchaImage: captcha.CaptchaImage,
		CaptchaCode:  captcha.CaptchaCode,
	}
	responseByte, err := json.Marshal(redisCaptcha)
	if err != nil {
		singleton.Info("json marshall", zap.String("", fmt.Sprintf("unexpected error : %v", err)))
		return models.ResponseGenerateCaptcha{}, constant.ErrInternal
	}

	err = capt.repository.redisRepo.Set(ctx, "set result captcha to redis", fmt.Sprintf(captcha.CaptchaID), responseByte, expiredTimeRedis)
	if err != nil {
		singleton.Info("redis", zap.String("", fmt.Sprintf("unexpected error : %v", err)))
		return models.ResponseGenerateCaptcha{}, constant.ErrInternal
	}

	return models.ResponseGenerateCaptcha{
		CaptchaID:          captcha.CaptchaID,
		CaptchaImage:       fmt.Sprintf("data:image/jpeg;base64,%s", captcha.CaptchaImage),
		CaptchaExpiredTime: fmt.Sprintf("%v", expiredTimeRedis),
	}, nil
}

func (capt *Captcha) VerifyCaptcha(ctx context.Context, clientID string, request *models.RequestVerifyCaptcha) error {
	var (
		redisCaptcha models.RedisCaptcha
	)

	// get captcha id from redis
	redisTransaction, err := capt.repository.redisRepo.Get(ctx, "get result captcha from redis", request.CaptchaID)
	switch {
	case redisTransaction == "":
		return errors.New("captcha ID not found ")
	case err != nil:
		singleton.Info("redis", zap.String("", fmt.Sprintf("unexpected error : %v", err)))
		return constant.ErrInternal
	}

	if err = json.Unmarshal([]byte(redisTransaction), &redisCaptcha); err != nil {
		singleton.Info("json unmarshall", zap.String("", fmt.Sprintf("unexpected error : %v", err)))
		return constant.ErrInternal
	}

	if redisCaptcha.CaptchaCode != request.CaptchaCode {
		return errors.New("captcha code wrong ")
	}
	err = capt.repository.redisRepo.Del(ctx, "delete captcha Redis ", request.CaptchaID)
	if err != nil {
		singleton.Info("redis", zap.String("", fmt.Sprintf("unexpected error : %v", err)))
		return constant.ErrInternal
	}
	return nil
}
