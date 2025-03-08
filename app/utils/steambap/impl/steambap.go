package impl

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/steambap/captcha"
	"image/color"
)

const (
	DefaultCharPreset   = "0123456789"
	DefaultNoise        = 1.0
	DefaultCurvedNumber = 2
	DefaultTextLength   = 4
	DefaultWidth        = 150 // Default width
	DefaultHeight       = 80  // Default height
)

var (
	DefaultBackgroundColor = color.White
)

var ErrInvalidAttribute = errors.New("invalid captcha attribute")

type AttributeSteambap struct {
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Length int     `json:"length"`
	Noise  float64 `json:"noise"`
}

func (a *AttributeSteambap) Validate() error {
	if a.Width < 0 || a.Height < 0 || a.Length < 0 {
		return fmt.Errorf("%w: width, height, and length must be non-negative", ErrInvalidAttribute)
	}
	if a.Noise < 0 {
		return fmt.Errorf("%w: noise must be non-negative", ErrInvalidAttribute)
	}
	return nil
}

type CaptchaSteambap struct {
	CaptchaID    string `json:"captcha_id"`
	CaptchaCode  string `json:"captcha_code"`
	CaptchaImage string `json:"captcha_image"`
}

type SteambapCaptcha struct{}

func NewSteambapCaptcha() *SteambapCaptcha {
	return &SteambapCaptcha{}
}

func (scg *SteambapCaptcha) GenerateCaptcha(request AttributeSteambap) (CaptchaSteambap, error) {

	if err := request.Validate(); err != nil {
		return CaptchaSteambap{}, err
	}

	captchaID, err := uuid.NewV4()
	if err != nil {
		return CaptchaSteambap{}, fmt.Errorf("failed to generate captcha ID: %w", err)
	}

	buff := bytes.Buffer{}

	width := request.Width
	if width == 0 {
		width = DefaultWidth
	}
	height := request.Height
	if height == 0 {
		height = DefaultHeight
	}
	length := request.Length
	if length == 0 {
		length = DefaultTextLength
	}
	noise := request.Noise
	if noise == 0 {
		noise = DefaultNoise
	}

	data, err := captcha.New(width, height, func(options *captcha.Options) {
		options.CharPreset = DefaultCharPreset
		options.CurveNumber = DefaultCurvedNumber
		options.TextLength = length
		options.Noise = noise
		options.BackgroundColor = DefaultBackgroundColor
	})
	if err != nil {
		return CaptchaSteambap{}, fmt.Errorf("failed to create captcha: %w", err)
	}

	if err := data.WriteJPG(&buff, nil); err != nil {
		return CaptchaSteambap{}, fmt.Errorf("failed to write captcha image: %w", err) // Lebih spesifik.
	}

	base64Image := base64.StdEncoding.EncodeToString(buff.Bytes())

	return CaptchaSteambap{
		CaptchaID:    captchaID.String(),
		CaptchaImage: base64Image,
		CaptchaCode:  data.Text,
	}, nil
}
