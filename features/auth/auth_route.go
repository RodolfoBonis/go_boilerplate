package auth

import (
	"github.com/RodolfoBonis/go_boilerplate/features/auth/di"
	"github.com/gin-gonic/gin"
)

func InjectRoutes(route *gin.RouterGroup) {
	var authUC = di.AuthInjection()

	authRoute := route.Group("/auth")
	authRoute.POST("/", authUC.ValidateLogin)
	authRoute.POST("/logout", authUC.Logout)
	authRoute.POST("/refresh", authUC.RefreshAuthToken)
}
