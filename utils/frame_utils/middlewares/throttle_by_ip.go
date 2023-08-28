package middlewares

import (
	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"github.com/three-kinds/user-center/utils/frame_utils/gin_utils"
	"github.com/three-kinds/user-center/utils/service_utils/se"
	"time"
)

func ThrottleByIP(limitPerHour uint) gin.HandlerFunc {
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Hour,
		Limit: limitPerHour,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: func(c *gin.Context, info ratelimit.Info) {
			gin_utils.AbortWithError(c, se.ThrottledError("IP受限"))
		},
		KeyFunc: func(c *gin.Context) string {
			return c.ClientIP()
		},
	})

	return mw
}
