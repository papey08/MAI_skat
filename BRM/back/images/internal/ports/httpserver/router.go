package httpserver

import (
	"github.com/gin-gonic/gin"
	_ "images/docs"
	"images/internal/app"
	"images/pkg/logger"
)

func appRouter(r *gin.RouterGroup, a app.App, logs logger.Logger) {
	r.Use(panicMiddleware(logs))
	r.Use(logMiddleware(logs))

	r.POST("images", handleAddImage(a))
	r.GET("images/:id", handleGetImage(a))
}
