package models

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	MaxWidth  = 1000
	MinWidth  = 80
	MaxHeight = 1000
	MinHeight = 50
	MaxNoise  = 10
	MinNoise  = 1.0
	MaxLength = 8
	MinLength = 4
)

type RequestGenerateCaptcha struct {
	Width  int     `json:"width" query:"width"`
	Height int     `json:"height" query:"height"`
	Length int     `json:"length" query:"length"`
	Noise  float64 `json:"noise" query:"noise"`
}
type ResponseGenerateCaptcha struct {
	CaptchaID          string `json:"captcha_id"`
	CaptchaImage       string `json:"captcha_image"`
	CaptchaExpiredTime string `json:"captcha_expired_time"`
}

type RedisCaptcha struct {
	CaptchaID    string `json:"captcha_id"`
	CaptchaImage string `json:"captcha_image"`
	CaptchaCode  string `json:"captcha_code"`
}

type RequestVerifyCaptcha struct {
	CaptchaID   string `json:"captcha_id"`
	CaptchaCode string `json:"captcha_code"`
}

func (v *RequestGenerateCaptcha) Validate() error {
	if v.Width < MinWidth {
		return errors.New("Minimum Width is 80px ")
	}

	if v.Width > MaxWidth {
		return errors.New("Maximum Width is 1000px ")
	}

	if v.Height < MinHeight {
		return errors.New("Minimum Height is 50px ")
	}

	if v.Height > MaxHeight {
		return errors.New("Maximum Height is 1000px ")
	}

	if v.Length > MaxLength {
		return errors.New("Maximum Length is 8 ")
	}

	if v.Length < MinLength {
		return errors.New("Minimum Length is 4 ")
	}

	if v.Noise > MaxNoise {
		return errors.New("Maximum Noise is 10 ")
	}

	if v.Noise < MinNoise {
		return errors.New("Minimum Noise is 1.0 ")
	}

	return validation.ValidateStruct(v,
		validation.Field(&v.Width, validation.Required),
		validation.Field(&v.Height, validation.Required),
		validation.Field(&v.Length, validation.Required),
		validation.Field(&v.Noise, validation.Required),
	)
}

func (v *RequestVerifyCaptcha) Validate() error {
	return validation.ValidateStruct(v,
		validation.Field(&v.CaptchaID, validation.Required),
		validation.Field(&v.CaptchaCode, validation.Required),
	)
}
