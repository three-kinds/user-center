package di

import (
	"github.com/three-kinds/user-center/services/auth_service"
	"github.com/three-kinds/user-center/services/captcha_service"
	"github.com/three-kinds/user-center/services/profile_service"
	"github.com/three-kinds/user-center/services/reset_password_code_service"
	"github.com/three-kinds/user-center/services/user_management_service"
	"github.com/three-kinds/user-center/services/user_service"
)

// 单体

func NewUserService() user_service.IUserService {
	return user_service.NewUserServiceImpl(NewUserDAO())
}

func NewCaptchaService() captcha_service.ICaptchaService {
	return captcha_service.NewCaptchaServiceImpl(NewCaptchaDAO())
}

func NewResetPasswordService() reset_password_code_service.IResetPasswordCodeService {
	return reset_password_code_service.NewResetPasswordCodeServiceImpl(NewResetPasswordCodeDAO())
}

// 复合

func NewAuthService() auth_service.IAuthService {
	return auth_service.NewAuthServiceImpl(NewUserService(), NewCaptchaService(), NewResetPasswordService(), NewEmailUtil())
}

func NewProfileService() profile_service.IProfileService {
	return profile_service.NewProfileServiceImpl(NewUserDAO())
}

func NewUserManagementService() user_management_service.IUserManagementService {
	return user_management_service.NewUserManagementServiceImpl(NewUserDAO())
}
