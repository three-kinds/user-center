package auth_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/three-kinds/user-center/utils/frame_utils/gin_utils"
	"net/http"
)

type ForgotPasswordRequest struct {
	Email         string `json:"email" binding:"required,email"`
	CaptchaKey    string `json:"captcha_key" binding:"required,len=32"`
	CaptchaAnswer string `json:"captcha_answer" binding:"required"`
}

func (c *AuthController) ForgotPassword(ctx *gin.Context) {
	var request ForgotPasswordRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		gin_utils.AbortWithValidationError(ctx, err)
		return
	}

	err = c.authService.ForgotPassword(request.Email, request.CaptchaKey, request.CaptchaAnswer)
	if err != nil {
		gin_utils.AbortWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
