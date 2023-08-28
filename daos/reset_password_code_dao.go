package daos

import (
	"github.com/three-kinds/user-center/services/bo"
	"time"
)

type IResetPasswordCodeDAO interface {
	CreateCode(key string, userID int64, expiration time.Time) (*bo.ResetPasswordCodeBo, error)
	GetCodeByKey(key string) (*bo.ResetPasswordCodeBo, error)
	DeleteExpiredCode() error
	DeleteCodeByKey(key string) error
}
