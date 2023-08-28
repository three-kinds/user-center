package auth_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/three-kinds/user-center/utils/frame_utils/gin_utils"
	"net/http"
)

func (c *AuthController) GetCaptcha(ctx *gin.Context) {
	b64, tb64, key, err := c.authService.GetCaptcha()
	if err != nil {
		gin_utils.AbortWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"image_base64": b64,
		"thumb_base64": tb64,
		"captcha_key":  key,
	})
}
