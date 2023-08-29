package profile_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/three-kinds/user-center/utils/frame_utils/gin_utils"
	"net/http"
)

type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required,max=16,min=4"`
	NewPassword string `json:"new_password" binding:"required,max=16,min=4"`
}

func (c *ProfileController) UpdatePassword(ctx *gin.Context) {
	var request UpdatePasswordRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		gin_utils.AbortWithValidationError(ctx, err)
		return
	}

	user := gin_utils.GetUser(ctx)

	err = c.profileService.UpdatePassword(user.ID, request.OldPassword, request.NewPassword)
	if err != nil {
		gin_utils.AbortWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
