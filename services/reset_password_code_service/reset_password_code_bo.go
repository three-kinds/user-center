package reset_password_code_service

import "time"

type ResetPasswordCodeBo struct {
	Key        string
	UserID     int64
	Expiration time.Time
}
