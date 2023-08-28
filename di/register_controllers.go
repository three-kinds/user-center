package di

import (
	"github.com/gin-gonic/gin"
	"github.com/three-kinds/user-center/controllers/auth_controller"
	"github.com/three-kinds/user-center/initializers"
	"github.com/three-kinds/user-center/utils/frame_utils/middlewares"
)

func RegisterAuthControllerRouter(rg *gin.RouterGroup) {
	controller := auth_controller.NewAuthController(NewAuthService())

	router := rg.Group("auth")
	router.Use(middlewares.ThrottleByIP(initializers.Config.ThrottleByAnonymousIP))

	router.POST("/jwt", controller.ObtainToken)
	router.POST("/jwt/refresh", controller.RefreshToken)
	router.POST("/captcha", controller.GetCaptcha)
	router.POST("/user-registration", controller.RegisterUser)
	router.POST("/password-forgotten", controller.ForgotPassword)
	router.POST("/password-reset", controller.ResetPassword)
}
