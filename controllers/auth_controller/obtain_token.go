package auth_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/three-kinds/user-center/utils/frame_utils/gin_utils"
	"net/http"
)

type ObtainTokenRequest struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (c *AuthController) ObtainToken(ctx *gin.Context) {
	var request ObtainTokenRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		gin_utils.AbortWithValidationError(ctx, err)
		return
	}

	accessToken, refreshToken, err := c.authService.ObtainToken(request.Account, request.Password)
	if err != nil {
		gin_utils.AbortWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
