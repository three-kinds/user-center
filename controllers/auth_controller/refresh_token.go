package auth_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/three-kinds/user-center/utils/frame_utils/gin_utils"
	"net/http"
)

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"  binding:"required"`
}

func (c *AuthController) RefreshToken(ctx *gin.Context) {
	var request RefreshTokenRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		gin_utils.AbortWithValidationError(ctx, err)
		return
	}

	accessToken, err := c.authService.RefreshToken(request.RefreshToken)
	if err != nil {
		gin_utils.AbortWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}
