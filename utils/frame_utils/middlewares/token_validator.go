package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/three-kinds/user-center/services/user_service"
	"github.com/three-kinds/user-center/utils/frame_utils/gin_utils"
	"github.com/three-kinds/user-center/utils/service_utils/jwt_utils"
	"github.com/three-kinds/user-center/utils/service_utils/se"
	"strings"
)

func TokenValidator(userService user_service.IUserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var accessToken string
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			accessToken = fields[1]
		}

		if accessToken == "" {
			gin_utils.AbortWithError(ctx, se.InvalidTokenError("token is empty"))
			return
		}

		userID, err := jwt_utils.ValidateAccessToken(accessToken)
		if err != nil {
			gin_utils.AbortWithError(ctx, err)
			return
		}

		user, err := userService.GetActiveUserByID(userID)
		if err != nil {
			gin_utils.AbortWithError(ctx, err)
			return
		}

		gin_utils.SetUser(ctx, user)
		ctx.Next()
	}
}
