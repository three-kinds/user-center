package auth_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/three-kinds/user-center/daos"
	"github.com/three-kinds/user-center/initializers"
	"github.com/three-kinds/user-center/services/auth_service"
	"github.com/three-kinds/user-center/services/captcha_service"
	"github.com/three-kinds/user-center/services/reset_password_code_service"
	"github.com/three-kinds/user-center/services/user_service"
	"github.com/three-kinds/user-center/utils/frame_utils/test_utils"
	"github.com/three-kinds/user-center/utils/service_utils/email_util"
	"github.com/three-kinds/user-center/utils/service_utils/se"
	"go.uber.org/mock/gomock"
	"testing"
)

var authController *AuthController
var userService user_service.IUserService
var captchaService captcha_service.ICaptchaService
var captchaDAO daos.ICaptchaDAO
var resetPasswordCodeService reset_password_code_service.IResetPasswordCodeService

func init() {
	test_utils.InitOnTestController()
	captchaDAO = daos.NewCaptchaDAOImpl(initializers.DB)
	captchaService = captcha_service.NewCaptchaServiceImpl(captchaDAO)
	userService = user_service.NewUserServiceImpl(daos.NewUserDAOImpl(initializers.DB))
	resetPasswordCodeService = reset_password_code_service.NewResetPasswordCodeServiceImpl(daos.NewResetPasswordCodeDAOImpl(initializers.DB))
	authController = NewAuthController(
		auth_service.NewAuthServiceImpl(
			userService,
			captchaService,
			resetPasswordCodeService,
			email_util.NewEmailUtilImpl(),
		),
	)
}

func TestAuthController_ForgotPassword(t *testing.T) {
	test_utils.ClearTables()

	email := "1@xx.com"
	_ = userService.RegisterUserByEmailPassword(email, "password")
	_, _, key, _ := captchaService.GetCaptcha()
	captchaBO, _ := captchaDAO.GetCaptchaByKey(key)

	code, rd := test_utils.PrepareTestController(
		authController.ForgotPassword, "/", "/",
		gin.H{
			"email":          email,
			"captcha_key":    captchaBO.Key,
			"captcha_answer": test_utils.RebuildCaptchaAnswer(&captchaBO.Answer),
		},
		"",
	)
	assert.Equal(t, 200, code)
	assert.Equal(t, 0, len(rd))

	// validation error
	code, rd = test_utils.PrepareTestController(
		authController.ForgotPassword, "/", "/",
		gin.H{
			"email":          "email",
			"captcha_key":    captchaBO.Key,
			"captcha_answer": test_utils.RebuildCaptchaAnswer(&captchaBO.Answer),
		},
		"",
	)
	te := se.ValidationError("")
	assert.Equal(t, te.Code, code)
	assert.Equal(t, te.Status, rd["status"])
}

func TestAuthController_ForgotPassword_WithMock(t *testing.T) {
	test_utils.ClearTables()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authServiceMock := auth_service.NewMockIAuthService(ctrl)
	controller := NewAuthController(authServiceMock)

	email := "1@xx.com"
	_ = userService.RegisterUserByEmailPassword(email, "password")
	_, _, key, _ := captchaService.GetCaptcha()
	captchaBO, _ := captchaDAO.GetCaptchaByKey(key)

	mockErr := se.ClientKnownError("mock error")
	authServiceMock.EXPECT().ForgotPassword(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockErr)
	code, rd := test_utils.PrepareTestController(
		controller.ForgotPassword, "/", "/",
		gin.H{
			"email":          email,
			"captcha_key":    captchaBO.Key,
			"captcha_answer": test_utils.RebuildCaptchaAnswer(&captchaBO.Answer),
		},
		"",
	)
	assert.Equal(t, mockErr.Code, code)
	assert.Equal(t, mockErr.Status, rd["status"])
}
