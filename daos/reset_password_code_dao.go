package daos

import (
	"github.com/three-kinds/user-center/services/reset_password_code_service"
	"time"
)

type IResetPasswordCodeDAO interface {
	CreateCode(key string, userID int64, expiration time.Time) (*reset_password_code_service.ResetPasswordCodeBo, error)
	GetCodeByKey(key string) (*reset_password_code_service.ResetPasswordCodeBo, error)
	DeleteExpiredCode() error
	DeleteCodeByKey(key string) error
}
