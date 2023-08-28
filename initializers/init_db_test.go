package initializers

import (
	"github.com/stretchr/testify/assert"
	"github.com/three-kinds/user-center/daos/models"
	"github.com/three-kinds/user-center/utils/generic_utils/testify_addons"
	"testing"
)

func TestInitDB_Success(t *testing.T) {
	InitConfig("")
	InitDB(Config, &models.User{})
	// success
	assert.NotNil(t, DB)
}

func TestInitDB_FailedOnConnect(t *testing.T) {
	InitConfig("")
	config := &Configuration{
		DBHost:     Config.DBHost,
		DBPort:     Config.DBPort,
		DBName:     Config.DBName,
		DBUser:     "user",
		DBPassword: "123456",
	}
	testify_addons.PanicsWithValueMatch(t, "failed to connect the database", func() {
		InitDB(config, &models.User{})
	})
}

func TestEnsureDatabase_FailedOnCreateDatabase(t *testing.T) {
	InitConfig("")
	config := Configuration{}
	config = *Config
	config.DBName = "1/a?83!"
	// failed on create database
	testify_addons.PanicsWithValueMatch(t, "failed to create database", func() {
		ensureDatabase(&config)
	})
}

func TestInitDB_FailedOnAutoMigrate(t *testing.T) {
	InitConfig("")
	type tt struct {
		TT any
	}

	// failed on create database
	testify_addons.PanicsWithValueMatch(t, "auto migrate failed", func() {
		InitDB(Config, &tt{})
	})
}
