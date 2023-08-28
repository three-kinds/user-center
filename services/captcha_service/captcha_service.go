package captcha_service

type ICaptchaService interface {
	GetCaptcha() (b64 string, tb64 string, key string, err error)
	ValidateCaptcha(key string, answer string) (bool, error)
}
