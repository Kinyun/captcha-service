package usecase

import (
	"captcha-service/app/utils/steambap"
	"captcha-service/app/utils/steambap/impl"
)

type Connection struct {
}

type Repository struct {
	steambap steambap.SteambapCaptcha
}

type Captcha struct {
	conn       Connection
	repository Repository
}

func NewCaptchaRepository() Repository {
	return Repository{
		steambap: impl.NewSteambapCaptcha(),
	}
}

func New(repository Repository) *Captcha {
	return &Captcha{
		conn:       Connection{},
		repository: repository,
	}
}

func Initialize() *Captcha {
	return New(NewCaptchaRepository())
}
