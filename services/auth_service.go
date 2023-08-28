package services

type IAuthService interface {
	ObtainToken(account string, password string) (accessToken string, refreshToken string, err error)
	RefreshToken(refreshToken string) (accessToken string, err error)
	GetCaptcha() (imageBase64 string, thumbBase64 string, captchaKey string, err error)
	RegisterUser(email string, password string, captchaKey string, captchaDots string) error
	ForgotPassword(email string, captchaKey string, captchaDots string) error
	ResetPassword(code string, newPassword string, captchaKey string, captchaDots string) error
}
