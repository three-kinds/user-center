package test_utils

import (
	"github.com/three-kinds/user-center/daos/models"
	"github.com/three-kinds/user-center/initializers"
	"log"
)

func ClearTables() {
	tables := []any{&models.User{}, &models.ResetPasswordCode{}, &models.Captcha{}}
	db := initializers.DB
	for _, table := range tables {
		result := db.Where("id != ?", "0").Delete(table)
		if result.Error != nil {
			log.Panicln("ClearTable failed", result.Error)
		}
	}
}
