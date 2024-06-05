package httpserver

import (
	"auth/internal/app"
	"auth/pkg/logger"
	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func New(addr string, originAddrs []string, a app.App, logs logger.Logger) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(corsMiddleware(originAddrs))

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("api/v1/auth")
	appRouter(api, a, logs)

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
