package gin_utils

import (
	"github.com/gin-gonic/gin"
	"github.com/three-kinds/user-center/utils/service_utils/se"
)

func AbortWithError(ctx *gin.Context, err error) {
	serviceError, ok := err.(*se.ServiceError)
	if !ok {
		serviceError = se.ServerUnknownError(err.Error())
	}
	_ = ctx.Error(serviceError)

	ctx.AbortWithStatusJSON(serviceError.Code, gin.H{
		"status":  serviceError.Status,
		"message": serviceError.Message,
	})
}
