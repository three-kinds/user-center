package auth_service

import (
	"github.com/three-kinds/user-center/services/captcha_service"
	"github.com/three-kinds/user-center/services/reset_password_code_service"
	"github.com/three-kinds/user-center/services/user_service"
	"github.com/three-kinds/user-center/utils/service_utils/email_util"
	"github.com/three-kinds/user-center/utils/service_utils/jwt_utils"
	"github.com/three-kinds/user-center/utils/service_utils/se"
)

type AuthServiceImpl struct {
	userService              user_service.IUserService
	captchaService           captcha_service.ICaptchaService
	resetPasswordCodeService reset_password_code_service.IResetPasswordCodeService
	emailUtil                email_util.IEmailUtil
}

func (s *AuthServiceImpl) ObtainToken(account string, password string) (accessToken string, refreshToken string, err error) {
	user, err := s.userService.LoginByAccountAndPassword(account, password)
	if err != nil {
		return
	}

	accessToken, err = jwt_utils.CreateAccessToken(user.ID)
	if err != nil {
		return
	}

	refreshToken, err = jwt_utils.CreateRefreshToken(user.ID)
	if err != nil {
		return
	}

	return
}

func (s *AuthServiceImpl) RefreshToken(refreshToken string) (accessToken string, err error) {
	userID, err := jwt_utils.ValidateRefreshToken(refreshToken)
	if err != nil {
		return
	}

	user, err := s.userService.GetActiveUserByID(userID)
	if err != nil {
		return
	}

	accessToken, err = jwt_utils.CreateAccessToken(user.ID)
	if err != nil {
		return
	}

	return
}

func (s *AuthServiceImpl) GetCaptcha() (string, string, string, error) {
	return s.captchaService.GetCaptcha()
}

func (s *AuthServiceImpl) RegisterUser(email string, password string, captchaKey string, captchaAnswer string) error {
	ok, err := s.captchaService.ValidateCaptcha(captchaKey, captchaAnswer)
	if err != nil {
		return err
	}
	if !ok {
		return se.ValidationError("captcha mismatch")
	}

	return s.userService.RegisterUserByEmailPassword(email, password)
}

func (s *AuthServiceImpl) ForgotPassword(email string, captchaKey string, captchaAnswer string) error {
	ok, err := s.captchaService.ValidateCaptcha(captchaKey, captchaAnswer)
	if err != nil {
		return err
	}
	if !ok {
		return se.ValidationError("captcha mismatch")
	}

	user, err := s.userService.GetActiveUserByEmail(email)
	if err != nil {
		return err
	}

	code, err := s.resetPasswordCodeService.CreateCode(user.ID)
	if err != nil {
		return err
	}

	return s.emailUtil.SendResetPasswordEmail(user.Email, code.Key)
}

func (s *AuthServiceImpl) ResetPassword(codeKey string, newPassword string, captchaKey string, captchaAnswer string) error {
	ok, err := s.captchaService.ValidateCaptcha(captchaKey, captchaAnswer)
	if err != nil {
		return err
	}
	if !ok {
		return se.ValidationError("captcha mismatch")
	}

	userID, err := s.resetPasswordCodeService.ValidateCode(codeKey)
	if err != nil {
		return err
	}

	user, err := s.userService.GetActiveUserByID(userID)
	if err != nil {
		return err
	}

	return s.userService.ResetPassword(user.ID, newPassword)
}

func NewAuthServiceImpl(
	userService user_service.IUserService,
	captchaService captcha_service.ICaptchaService,
	resetPasswordCodeService reset_password_code_service.IResetPasswordCodeService,
	emailUtil email_util.IEmailUtil,
) *AuthServiceImpl {
	return &AuthServiceImpl{
		userService:              userService,
		captchaService:           captchaService,
		resetPasswordCodeService: resetPasswordCodeService,
		emailUtil:                emailUtil,
	}
}
