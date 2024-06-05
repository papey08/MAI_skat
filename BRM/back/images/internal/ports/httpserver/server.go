package httpserver

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"images/internal/app"
	"images/pkg/logger"
	"net/http"
)

func New(addr string, originAddrs []string, a app.App, logs logger.Logger) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(corsMiddleware(originAddrs))
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("api/v1")
	appRouter(api, a, logs)

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
