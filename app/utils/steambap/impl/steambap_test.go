package impl

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttributeSteambap_Validate(t *testing.T) {
	tests := []struct {
		name      string
		attribute AttributeSteambap
		wantErr   bool
		err       error
	}{
		{
			name: "Valid attribute",
			attribute: AttributeSteambap{
				Width:  100,
				Height: 50,
				Length: 4,
				Noise:  0.5,
			},
			wantErr: false,
		},
		{
			name: "Negative width",
			attribute: AttributeSteambap{
				Width:  -100,
				Height: 50,
				Length: 4,
				Noise:  0.5,
			},
			wantErr: true,
			err:     errors.New("invalid captcha attribute: width, height, and length must be non-negative"),
		},
		{
			name: "Negative height",
			attribute: AttributeSteambap{
				Width:  100,
				Height: -50,
				Length: 4,
				Noise:  0.5,
			},
			wantErr: true,
			err:     errors.New("invalid captcha attribute: width, height, and length must be non-negative"),
		},
		{
			name: "Negative length",
			attribute: AttributeSteambap{
				Width:  100,
				Height: 50,
				Length: -4,
				Noise:  0.5,
			},
			wantErr: true,
			err:     errors.New("invalid captcha attribute: width, height, and length must be non-negative"),
		},
		{
			name: "Negative noise",
			attribute: AttributeSteambap{
				Width:  100,
				Height: 50,
				Length: 4,
				Noise:  -0.5,
			},
			wantErr: true,
			err:     errors.New("invalid captcha attribute: noise must be non-negative"),
		},
		{
			name: "Zero values",
			attribute: AttributeSteambap{
				Width:  0,
				Height: 0,
				Length: 0,
				Noise:  0,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.attribute.Validate()
			if tt.wantErr {
				assert.EqualError(t, err, tt.err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSteambapCaptcha_GenerateCaptcha(t *testing.T) {
	tests := []struct {
		name          string
		request       AttributeSteambap
		wantErr       bool
		expectedError string
	}{
		{
			name: "Valid request",
			request: AttributeSteambap{
				Width:  150,
				Height: 80,
				Length: 4,
				Noise:  1.0,
			},
			wantErr: false,
		},
		{
			name: "Default values",
			request: AttributeSteambap{
				Width:  0,
				Height: 0,
				Length: 0,
				Noise:  0,
			},
			wantErr: false,
		},
		{
			name: "Invalid attribute - negative width",
			request: AttributeSteambap{
				Width:  -100,
				Height: 80,
				Length: 4,
				Noise:  1.0,
			},
			wantErr:       true,
			expectedError: "invalid captcha attribute: width, height, and length must be non-negative",
		},
		{
			name: "Invalid attribute - negative noise",
			request: AttributeSteambap{
				Width:  150,
				Height: 80,
				Length: 4,
				Noise:  -1.0,
			},
			wantErr:       true,
			expectedError: "invalid captcha attribute: noise must be non-negative",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scg := NewSteambapCaptcha()
			result, err := scg.GenerateCaptcha(tt.request)

			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError)
				assert.Empty(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, result.CaptchaID)
				assert.NotEmpty(t, result.CaptchaImage)
				assert.NotEmpty(t, result.CaptchaCode)
			}
		})
	}
}
