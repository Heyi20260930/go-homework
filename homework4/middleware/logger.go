
package middleware

import (
	"log"
	"time"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		log.Printf("[%s] %s %s -> %d (%v)",
			time.Now().Format("2026-10-26 11:39:05"),
			c.Request.Method, c.Request.URL.Path,
			c.Writer.Status(), latency)
	}
}