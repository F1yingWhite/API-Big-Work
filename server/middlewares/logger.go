package middlewares

import (
	"bytes"
	"io"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logID := uuid.New().String()
		c.Set("logID", logID)
		startTime := time.Now()

		// 检查 Content-Type 头，判断请求是否包含文件
		contentType := c.GetHeader("Content-Type")
		if contentType != "" && (strings.Contains(contentType, "multipart/form-data")) {
			log.Printf("[LogID:%s] Request: %s %s", logID, c.Request.Method, c.Request.RequestURI)
		} else {
			body, _ := c.GetRawData()
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			log.Printf("[LogID:%s] Request: %s %s %s", logID, c.Request.Method, c.Request.RequestURI, string(body))
		}

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
