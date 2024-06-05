package httpserver

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"transport-api/internal/app"
	"transport-api/internal/ports/httpserver/middleware"
	"transport-api/pkg/logger"
	"transport-api/pkg/tokenizer"
)

func New(addr string, originAddrs []string, a app.App, tkn tokenizer.Tokenizer, logs logger.Logger) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(middleware.Cors(originAddrs))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("api/v1")
	appRouter(api, a, tkn, logs)

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
