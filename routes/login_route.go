package routes

import (
	"github.com/RodolfoBonis/go_boilerplate/features/login/di"
	"github.com/gin-gonic/gin"
)

func loginRoutes(route *gin.RouterGroup) {
	var loginUC = di.LoginInjection()

	loginRoute := route.Group("/login")
	loginRoute.POST("/", loginUC.ValidateLogin)
}
