package steambap

import (
	"captcha-service/app/utils/steambap/impl"
)

type SteambapCaptcha interface {
	GenerateCaptcha(request impl.AttributeSteambap) (impl.CaptchaSteambap, error)
}
