package daos

import (
	"github.com/three-kinds/user-center/services/bo"
	"time"
)

type ICaptchaDAO interface {
	CreateCaptcha(key string, answer string, expiration time.Time) (*bo.CaptchaBo, error)
	GetCaptchaByKey(key string) (*bo.CaptchaBo, error)
	DeleteExpiredCaptcha() error
	DeleteCaptchaByKey(key string) error
}
