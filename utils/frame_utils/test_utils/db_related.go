package test_utils

import (
	"github.com/three-kinds/user-center/utils/generic_utils/gorm_addons"
	"log"
)

func ClearTables(db gorm_addons.IDB, tables ...any) {
	for _, table := range tables {
		result := db.Where("id != ?", "0").Delete(table)
		if result.Error != nil {
			log.Panicln("ClearTable failed", result.Error)
		}
	}
}
