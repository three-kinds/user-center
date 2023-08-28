package initializers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strings"
)

var DB *gorm.DB

func openDsn(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("failed to connect the database")
	}
	return db
}

func ensureDatabase(config *Configuration) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DBHost, config.DBUser, config.DBPassword, config.DBPort,
	)
	db := openDsn(dsn)
	defer func() {
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
	}()

	sql := fmt.Sprintf("CREATE DATABASE %s;", config.DBName)
	err := db.Exec(sql).Error
	if err != nil {
		cause := err.Error()
		if !strings.Contains(cause, "already exists") {
			log.Panicln("failed to create database", err)
		}
	}
}

func InitDB(config *Configuration, tables ...any) {
	ensureDatabase(config)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort,
	)
	db := openDsn(dsn)
	err := db.AutoMigrate(tables...)
	if err != nil {
		log.Panicln("auto migrate failed", err)
	}

	DB = db
}
