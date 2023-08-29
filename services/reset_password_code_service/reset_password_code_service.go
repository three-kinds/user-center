package reset_password_code_service

import "github.com/three-kinds/user-center/services/bo"

type IResetPasswordCodeService interface {
	CreateCode(userID int64) (*bo.ResetPasswordCodeBo, error)
	ValidateCode(codeKey string) (int64, error)
}
