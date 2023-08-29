package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func JsonLogger(skipPaths ...string) gin.HandlerFunc {
	var skip map[string]struct{}

	if length := len(skipPaths); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range skipPaths {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		if _, ok := skip[path]; !ok {
			// Stop timer
			timestamp := time.Now()
			latency := timestamp.Sub(start)
			if latency > time.Minute {
				latency = latency.Truncate(time.Second)
			}

			clientIP := c.ClientIP()
			method := c.Request.Method
			statusCode := c.Writer.Status()
			bodySize := c.Writer.Size()

			if raw != "" {
				path = path + "?" + raw
			}

			logger := logrus.WithFields(logrus.Fields{
				"name":        "json_logger",
				"timestamp":   timestamp.Format(time.DateTime),
				"client_ip":   clientIP,
				"method":      method,
				"path":        path,
				"proto":       c.Request.Proto,
				"status_code": statusCode,
				"latency":     latency,
				"user_agent":  c.Request.UserAgent(),
				"body_size":   bodySize,
			})
			if statusCode < 400 {
				logger.Info("")
			} else {
				errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()
				if statusCode < 500 {
					logger.Warning(errorMessage)
				} else {
					logger.Error(errorMessage)
				}
			}
		}
	}
}
