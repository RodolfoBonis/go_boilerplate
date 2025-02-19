package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.package_name}}/features/auth/di"
)

func InjectRoutes(route *gin.RouterGroup) {
	var authUC = di.AuthInjection()

	authRoute := route.Group("/auth")
	authRoute.POST("/", authUC.ValidateLogin)
	authRoute.POST("/logout", authUC.Logout)
	authRoute.POST("/refresh", authUC.RefreshAuthToken)
}
