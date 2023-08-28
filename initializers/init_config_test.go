package initializers

import (
	"github.com/three-kinds/user-center/utils/generic_utils/testify_addons"
	"os"
	"testing"
)

func TestInitConfig_FailedOnGetPath(t *testing.T) {
	baseDir := os.Getenv("BASE_DIR")
	_ = os.Setenv("BASE_DIR", "")
	testify_addons.PanicsWithValueMatch(t, "please set config file path or set BASE_DIR", func() {
		InitConfig("")
	})
	_ = os.Setenv("BASE_DIR", baseDir)
}

func TestInitConfig_FailedOnRead(t *testing.T) {
	testify_addons.PanicsWithValueMatch(t, "failed to load config file", func() {
		InitConfig("./test_assets")
	})
}

func TestInitConfig_FailedOnUnmarshal(t *testing.T) {
	testify_addons.PanicsWithValueMatch(t, "failed to unmarshal config", func() {
		InitConfig("./failed_test_assets")
	})
}
