package test_utils

import (
	"github.com/three-kinds/user-center/initializers"
)

func InitOnTestDAO(tables ...any) {
	initializers.InitConfig("")
	initializers.InitDB(initializers.Config, tables...)
	initializers.InitLogger()
}

func InitOnTestService(tables ...any) {
	InitOnTestDAO(tables...)
	initializers.InitSnowflakeNode(initializers.Config)
}

func InitOnTestController(tables ...any) {
	InitOnTestService(tables...)
	initializers.InitValidators()
}
