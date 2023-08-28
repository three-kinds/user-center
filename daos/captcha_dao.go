package daos

import (
	"github.com/three-kinds/user-center/services/captcha_service"
	"time"
)

type ICaptchaDAO interface {
	CreateCaptcha(key string, answer string, expiration time.Time) (*captcha_service.CaptchaBo, error)
	GetCaptchaByKey(key string) (*captcha_service.CaptchaBo, error)
	DeleteExpiredCaptcha() error
	DeleteCaptchaByKey(key string) error
}
