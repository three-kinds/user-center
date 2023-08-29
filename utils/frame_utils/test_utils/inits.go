package test_utils

import (
	"github.com/three-kinds/user-center/daos/models"
	"github.com/three-kinds/user-center/initializers"
)

func InitOnTestDAO() {
	initializers.InitConfig("")
	initializers.InitDB(initializers.Config, &models.User{}, &models.ResetPasswordCode{}, &models.Captcha{})
	initializers.InitLogger()
}

func InitOnTestService() {
	InitOnTestDAO()
	initializers.InitSnowflakeNode(initializers.Config)
}

func InitOnTestController() {
	InitOnTestService()
	initializers.InitValidators()
}
