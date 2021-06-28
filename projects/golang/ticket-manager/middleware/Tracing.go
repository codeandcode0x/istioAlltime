package middleware

import (
	tracing "github.com/codeandcode0x/traceandtrace-go"
	"github.com/gin-gonic/gin"
)

// Tracing 中间件
func Tracing() gin.HandlerFunc {
	return func(c *gin.Context) {
		// add tracing
		_, cancel := tracing.AddHttpTracing(
			"ticket-manager"+c.FullPath(),
			c.Request.Header,
			map[string]string{
				"component":      "gin-server",
				"httpMethod":     c.Request.Method,
				"httpUrl":        c.Request.URL.String(),
				"proto":          c.Request.Proto,
				"peerService":    c.HandlerName(),
				"httpStatusCode": "200",
				"spanKind":       "server",
			})
		defer cancel()

		c.Next()
		return
	}
}