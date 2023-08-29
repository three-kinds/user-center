package user_management_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/three-kinds/user-center/controllers/vo"
	"github.com/three-kinds/user-center/utils/frame_utils/gin_utils"
	"net/http"
)

type ListUserRequest struct {
	Page        int   `json:"page" binding:"required"`
	Size        int   `json:"size" binding:"required"`
	IsActive    *bool `json:"is_active"`
	IsStaff     *bool `json:"is_staff"`
	IsSuperuser *bool `json:"is_superuser"`
}

func (c *UserManagementController) ListUsers(ctx *gin.Context) {
	var request ListUserRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		gin_utils.AbortWithValidationError(ctx, err)
		return
	}

	total, users, err := c.userManagementService.ListUsers(request.Page, request.Size, request.IsActive, request.IsSuperuser)
	if err != nil {
		gin_utils.AbortWithError(ctx, err)
		return
	}

	userList := make([]*vo.UserVO, len(users))
	for i, user := range users {
		userList[i] = (*vo.UserVO)(user)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"total":      total,
		"entry_list": userList,
	})
}
