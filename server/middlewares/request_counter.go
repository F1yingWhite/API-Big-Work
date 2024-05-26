package middlewares

import (
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

var RequestCounter int64

func RequestCounterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		atomic.AddInt64(&RequestCounter, 1)
		c.Next()
	}
}
