package test_utils

import (
	"github.com/three-kinds/user-center/initializers"
)

func InitOnTestDAO(tables ...any) {
	initializers.InitConfig("")
	initializers.InitDB(initializers.Config, tables...)
	ClearTables(initializers.DB, tables...)

}

func InitOnTestService(tables ...any) {
	InitOnTestDAO(tables)
	initializers.InitSnowflakeNode(initializers.Config)
}
