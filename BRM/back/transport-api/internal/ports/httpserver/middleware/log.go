package middleware

import (
	"github.com/gin-gonic/gin"
	"transport-api/pkg/logger"
)

func Log(logs logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		logs.Info(logger.Fields{
			"path":   c.Request.URL.Path,
			"method": c.Request.Method,
			"status": c.Writer.Status(),
		}, "")
	}
}
