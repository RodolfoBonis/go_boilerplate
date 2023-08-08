package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InjectRoute(route *gin.RouterGroup) {
	route.GET("/health_check", healthCheck)
}

// healthCheck godoc
// @Summary Health Check
// @Schemes
// @Description Check if This service is healthy
// @Tags HealthCheck
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} utils.HttpError
// @Failure 401 {object} utils.HttpError
// @Failure 403 {object} utils.HttpError
// @Failure 409 {object} utils.HttpError
// @Failure 500 {object} utils.HttpError
// @Router /health_check [get]
func healthCheck(context *gin.Context) {
	context.String(http.StatusOK, "This Service is Healthy")
}
