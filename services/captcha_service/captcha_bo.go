package captcha_service

import "time"

type CaptchaBo struct {
	Key        string
	Answer     string
	Expiration time.Time
}
