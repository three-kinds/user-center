package user_management_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/three-kinds/user-center/controllers/vo"
	"github.com/three-kinds/user-center/services/bo"
	"github.com/three-kinds/user-center/utils/frame_utils/gin_utils"
	"mime/multipart"
	"net/http"
)

type PartialUpdateUserRequest struct {
	Email       *string               `form:"email" binding:"omitempty,email"`
	Username    *string               `form:"username" binding:"omitempty"`
	Password    *string               `form:"password" binding:"omitempty"`
	IsActive    *bool                 `form:"is_active" binding:"omitempty"`
	IsStaff     *bool                 `form:"is_staff" binding:"omitempty"`
	IsSuperuser *bool                 `form:"is_superuser" binding:"omitempty"`
	Nickname    *string               `form:"nickname"`
	PhoneNumber *string               `form:"phone_number" binding:"omitempty,phone_number"`
	Avatar      *multipart.FileHeader `form:"avatar"`
}

func (c *UserManagementController) PartialUpdateUser(ctx *gin.Context) {
	var requestURI gin_utils.Int64IDRequest
	err := ctx.ShouldBindUri(&requestURI)
	if err != nil {
		gin_utils.AbortWithValidationError(ctx, err)
		return
	}
	var request PartialUpdateUserRequest
	err = ctx.ShouldBind(&request)
	if err != nil {
		gin_utils.AbortWithValidationError(ctx, err)
		return
	}

	user, err := c.userManagementService.GetUserByID(requestURI.ID)
	if err != nil {
		gin_utils.AbortWithError(ctx, err)
		return
	}
	// avatar
	var avatar *string
	if request.Avatar != nil {
		relativePath, err := gin_utils.ReceiveUploadedImage(ctx, request.Avatar, fmt.Sprintf("/avatar/%d", user.ID))
		if err != nil {
			gin_utils.AbortWithError(ctx, err)
			return
		}
		avatar = &relativePath
	}

	updatedUser, err := c.userManagementService.PartialUpdateUser(user.ID, &bo.UpdateUserBO{
		Email:       request.Email,
		Username:    request.Username,
		Password:    request.Password,
		IsActive:    request.IsActive,
		IsStaff:     request.IsStaff,
		IsSuperuser: request.IsSuperuser,
		Nickname:    request.Nickname,
		PhoneNumber: request.PhoneNumber,
		Avatar:      avatar,
	})
	if err != nil {
		gin_utils.AbortWithError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, vo.TUserBO2UserVO(updatedUser))
}
