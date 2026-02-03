package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger middleware for HTTP request logging
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		method := c.Request.Method
		path := c.Request.RequestURI

		c.Next()

		statusCode := c.Writer.Status()
		duration := time.Since(startTime)

		log.Printf("[%s] %s %d %s", method, path, statusCode, duration)
	}
}
