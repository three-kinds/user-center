package profile_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/three-kinds/user-center/controllers/vo"
	"github.com/three-kinds/user-center/utils/frame_utils/gin_utils"
	"net/http"
)

func (c *ProfileController) GetProfile(ctx *gin.Context) {
	userBO := gin_utils.GetUser(ctx)
	ctx.JSON(http.StatusOK, vo.TUserBO2UserVO(userBO))
}
