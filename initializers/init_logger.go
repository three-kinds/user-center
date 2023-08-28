package initializers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

var DefaultLogger *logrus.Entry

func InitLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	mode := Config.Mode
	if mode == gin.ReleaseMode {
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(logrus.DebugLevel)
	}

	DefaultLogger = logrus.WithFields(logrus.Fields{
		"name": "Default",
	})
}
