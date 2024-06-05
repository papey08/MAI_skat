package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"transport-api/internal/model"
	"transport-api/pkg/logger"
)

func Panic(logs logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logs.Error(logger.Fields{
					"path":   c.Request.URL.Path,
					"method": c.Request.Method,
				}, fmt.Sprintf("panic: %v", r))

				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"data": nil, "error": model.ErrCoreError.Error()})
			}
		}()
		c.Next()
	}
}
