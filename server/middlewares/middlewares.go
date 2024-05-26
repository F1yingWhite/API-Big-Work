package middlewares

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/juju/ratelimit"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logID := uuid.New().String()
		c.Set("logID", logID)
		startTime := time.Now()
		body, _ := c.GetRawData()
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		log.Printf("[LogID:%s] Request: %s", logID, string(body))
		c.Next()
		elapsedTime := time.Since(startTime)
		log.Printf("[LogID:%s] %s %s %d %v %s",
			logID,
			c.Request.Method,
			c.Request.RequestURI,
			c.Writer.Status(),
			elapsedTime,
			c.Request.UserAgent(),
		)
	}
}

func RateLimitMiddleware(fillInterval time.Duration, cap, quantum int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(fillInterval, cap, quantum)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			c.String(http.StatusForbidden, "rate limit...")
			c.Abort()
			return
		}
		c.Next()
	}
}

var RequestCounter int64

func RequestCounterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		atomic.AddInt64(&RequestCounter, 1)
		c.Next()
	}
}
