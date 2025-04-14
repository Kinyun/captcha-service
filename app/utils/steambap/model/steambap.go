package model

import (
	"errors"
	"fmt"
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
