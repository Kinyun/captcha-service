package models

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
