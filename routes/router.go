package routes

import (
	"github.com/RodolfoBonis/go_boilerplate/core/health"
	"github.com/RodolfoBonis/go_boilerplate/features/auth"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(router *gin.Engine) {

	root := router.Group("/v1")

	root.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	health.InjectRoute(root)
	auth.InjectRoutes(root)
}
