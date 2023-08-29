package di

import (
	"github.com/gin-gonic/gin"
	"github.com/three-kinds/user-center/controllers/auth_controller"
	"github.com/three-kinds/user-center/controllers/profile_controller"
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

func RegisterProfileControllerRouter(rg *gin.RouterGroup) {
	controller := profile_controller.NewProfileController(NewProfileService())

	router := rg.Group("profile")
	router.Use(middlewares.TokenValidator(NewUserService()))

	router.GET("", controller.GetProfile)
	router.PATCH("", controller.PartialUpdateProfile)
	router.PUT("/password", controller.UpdatePassword)
}
