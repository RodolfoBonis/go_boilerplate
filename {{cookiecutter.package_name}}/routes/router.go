package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.package_name}}/core/health"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.package_name}}/features/auth"
)

func InitializeRoutes(router *gin.Engine) {

	root := router.Group("/v1")

	root.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	health.InjectRoute(root)
	auth.InjectRoutes(root)
}
