package httpserver

import (
	"github.com/gin-gonic/gin"
	_ "registration/docs"
	"registration/internal/app"
	"registration/pkg/logger"
)

func appRouter(r *gin.RouterGroup, a app.App, logs logger.Logger) {
	r.Use(panicMiddleware(logs))
	r.Use(logMiddleware(logs))

	r.POST("register", addCompanyWithOwner(a))
	r.GET("companies/industries", getIndustriesMap(a))
}
