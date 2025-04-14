package steambap

import (
	steambap "captcha-service/app/utils/steambap/model"
)

type SteambapCaptcha interface {
	GenerateCaptcha(request steambap.AttributeSteambap) (steambap.CaptchaSteambap, error)
}
