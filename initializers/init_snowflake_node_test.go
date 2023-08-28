package initializers

import (
	"github.com/stretchr/testify/assert"
	"github.com/three-kinds/user-center/utils/generic_utils/testify_addons"
	"testing"
)

func TestInitSnowflakeNode_Success(t *testing.T) {
	InitConfig("")
	InitSnowflakeNode(Config)
	// success
	assert.NotNil(t, SnowflakeNode)
}

func TestInitSnowflakeNode_Failed(t *testing.T) {
	InitConfig("")
	config := Config
	config.SnowflakeNode = -1

	testify_addons.PanicsWithValueMatch(t, "failed to init snowflake node", func() {
		InitSnowflakeNode(config)
	})
}
