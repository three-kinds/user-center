package user_management_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/three-kinds/user-center/utils/frame_utils/gin_utils"
	"net/http"
)

func (c *UserManagementController) DeleteUser(ctx *gin.Context) {
	var request gin_utils.Int64IDRequest
	err := ctx.ShouldBindUri(&request)
	if err != nil {
		gin_utils.AbortWithValidationError(ctx, err)
		return
	}

	err = c.userManagementService.DeleteUser(request.ID)
	if err != nil {
		gin_utils.AbortWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})

}
