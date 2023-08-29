package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/three-kinds/user-center/utils/frame_utils/gin_utils"
	"github.com/three-kinds/user-center/utils/service_utils/se"
)

func IsSuperuser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := gin_utils.GetUser(ctx)
		if !user.IsSuperuser {
			gin_utils.AbortWithError(ctx, se.ForbiddenError("not superuser"))
			return
		}

		ctx.Next()
	}
}
