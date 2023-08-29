package profile_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/three-kinds/user-center/controllers/vo"
	"github.com/three-kinds/user-center/services/bo"
	"github.com/three-kinds/user-center/utils/frame_utils/gin_utils"
	"mime/multipart"
	"net/http"
)

type PartialUpdateProfileRequest struct {
	Nickname    *string               `form:"nickname"`
	PhoneNumber *string               `form:"phone_number" binding:"omitempty,phone_number"`
	Avatar      *multipart.FileHeader `form:"avatar"`
}

func (c *ProfileController) PartialUpdateProfile(ctx *gin.Context) {
	var request PartialUpdateProfileRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		gin_utils.AbortWithValidationError(ctx, err)
		return
	}
	user := gin_utils.GetUser(ctx)

	var avatar *string
	if request.Avatar != nil {
		relativePath, err := gin_utils.ReceiveUploadedImage(ctx, request.Avatar, fmt.Sprintf("/avatar/%d", user.ID))
		if err != nil {
			gin_utils.AbortWithError(ctx, err)
			return
		}
		avatar = &relativePath
	}

	updatedUser, err := c.profileService.PartialUpdateProfile(user.ID, &bo.UpdateProfileBO{
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
