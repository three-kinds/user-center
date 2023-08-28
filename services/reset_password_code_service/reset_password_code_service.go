package reset_password_code_service

type IResetPasswordCodeService interface {
	CreateCode(userID int64) (*ResetPasswordCodeBo, error)
	ValidateCode(codeKey string) (int64, error)
}
