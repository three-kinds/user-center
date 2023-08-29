package user_management_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/three-kinds/user-center/services/user_management_service"
	"github.com/three-kinds/user-center/utils/frame_utils/gin_utils"
	"net/http"
)

type CreateUserRequest struct {
	Email       string `json:"email" binding:"required,email"`
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	IsSuperuser bool   `json:"is_superuser" binding:"required"`
}

func (c *UserManagementController) CreateUser(ctx *gin.Context) {
	var request CreateUserRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		gin_utils.AbortWithValidationError(ctx, err)
		return
	}

	user, err := c.userManagementService.CreateUser(&user_management_service.CreateUserBO{
		Email:       request.Email,
		Username:    request.Username,
		Password:    request.Password,
		IsSuperuser: request.IsSuperuser,
	})
	if err != nil {
		gin_utils.AbortWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": user.ID,
	})
}
