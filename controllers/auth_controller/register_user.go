package auth_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/three-kinds/user-center/utils/frame_utils/gin_utils"
	"net/http"
)

type RegisterUserRequest struct {
	Email         string `json:"email" binding:"required,email"`
	Password      string `json:"password" binding:"required"`
	CaptchaKey    string `json:"captcha_key" binding:"required,len=32"`
	CaptchaAnswer string `json:"captcha_answer" binding:"required"`
}

func (c *AuthController) RegisterUser(ctx *gin.Context) {
	var request RegisterUserRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		gin_utils.AbortWithValidationError(ctx, err)
		return
	}

	err = c.authService.RegisterUser(request.Email, request.Password, request.CaptchaKey, request.CaptchaAnswer)
	if err != nil {
		gin_utils.AbortWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
