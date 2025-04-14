package impl

import (
	"bytes"
	steambap "captcha-service/app/utils/steambap/model"
	"encoding/base64"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/steambap/captcha"
)

type SteambapCaptcha struct{}

func NewSteambapCaptcha() *SteambapCaptcha {
	return &SteambapCaptcha{}
}

func (scg *SteambapCaptcha) GenerateCaptcha(request steambap.AttributeSteambap) (steambap.CaptchaSteambap, error) {

	if err := request.Validate(); err != nil {
		return steambap.CaptchaSteambap{}, err
	}

	captchaID, err := uuid.NewV4()
	if err != nil {
		return steambap.CaptchaSteambap{}, fmt.Errorf("failed to generate captcha ID: %w", err)
	}

	buff := bytes.Buffer{}

	width := request.Width
	if width == 0 {
		width = steambap.DefaultWidth
	}
	height := request.Height
	if height == 0 {
		height = steambap.DefaultHeight
	}
	length := request.Length
	if length == 0 {
		length = steambap.DefaultTextLength
	}
	noise := request.Noise
	if noise == 0 {
		noise = steambap.DefaultNoise
	}

	data, err := captcha.New(width, height, func(options *captcha.Options) {
		options.CharPreset = steambap.DefaultCharPreset
		options.CurveNumber = steambap.DefaultCurvedNumber
		options.TextLength = length
		options.Noise = noise
		options.BackgroundColor = steambap.DefaultBackgroundColor
	})
	if err != nil {
		return steambap.CaptchaSteambap{}, fmt.Errorf("failed to create captcha: %w", err)
	}

	if err := data.WriteJPG(&buff, nil); err != nil {
		return steambap.CaptchaSteambap{}, fmt.Errorf("failed to write captcha image: %w", err) // Lebih spesifik.
	}

	base64Image := base64.StdEncoding.EncodeToString(buff.Bytes())

	return steambap.CaptchaSteambap{
		CaptchaID:    captchaID.String(),
		CaptchaImage: base64Image,
		CaptchaCode:  data.Text,
	}, nil
}
