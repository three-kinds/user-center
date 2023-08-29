package gin_utils

import (
	"github.com/gin-gonic/gin"
	"github.com/three-kinds/user-center/services/bo"
)

const userKey = "user"

func SetUser(ctx *gin.Context, user *bo.UserBO) {
	ctx.Set(userKey, user)
}

func GetUser(ctx *gin.Context) *bo.UserBO {
	value, ok := ctx.Get(userKey)
	if !ok {
		panic("can not get user")
	}
	return value.(*bo.UserBO)
}
