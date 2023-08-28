package bo

import "time"

type CaptchaBo struct {
	Key        string
	Answer     string
	Expiration time.Time
}
