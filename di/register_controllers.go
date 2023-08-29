package di

import (
	"github.com/gin-gonic/gin"
	"github.com/three-kinds/user-center/controllers/auth_controller"
	"github.com/three-kinds/user-center/controllers/profile_controller"
	"github.com/three-kinds/user-center/controllers/user_management_controller"
	"github.com/three-kinds/user-center/initializers"
	"github.com/three-kinds/user-center/utils/frame_utils/middlewares"
	"net/http"
)

func RegisterOperationRouter(rg *gin.RouterGroup) {
	rg.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"status": "OK"})
	})
}

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

func RegisterUserManagementControllerRouter(rg *gin.RouterGroup) {
	controller := user_management_controller.NewUserManagementController(NewUserManagementService())

	router := rg.Group("user_management")
	router.Use(middlewares.TokenValidator(NewUserService()), middlewares.IsSuperuser())

	router.GET("", controller.ListUsers)
	router.POST("", controller.CreateUser)

	router.GET("/:id", controller.GetUser)
	router.PATCH("/:id", controller.PartialUpdateUser)
	router.DELETE("/:id", controller.DeleteUser)
}
