package bo

import "time"

type ResetPasswordCodeBo struct {
	Key        string
	UserID     int64
	Expiration time.Time
}
