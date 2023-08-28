package auth_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/three-kinds/user-center/utils/frame_utils/gin_utils"
	"net/http"
)

type ResetPasswordRequest struct {
	Code          string `json:"code" binding:"required"`
	NewPassword   string `json:"new_password" binding:"required"`
	CaptchaKey    string `json:"captcha_key" binding:"required,len=32"`
	CaptchaAnswer string `json:"captcha_answer" binding:"required"`
}

func (c *AuthController) ResetPassword(ctx *gin.Context) {
	var request ResetPasswordRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		gin_utils.AbortWithValidationError(ctx, err)
		return
	}

	err = c.authService.ResetPassword(request.Code, request.NewPassword, request.CaptchaKey, request.CaptchaAnswer)
	if err != nil {
		gin_utils.AbortWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
