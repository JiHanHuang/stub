package info

import (
	"fmt"
	"time"

	"github.com/JiHanHuang/stub/pkg/logging"
	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func MSG() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path
		status := c.Writer.Status()
		logging.Info(fmt.Sprintf(" | %3d | %12v | %16s | %8s | %s",
			status, latency, clientIP, method, path))
	}
}
