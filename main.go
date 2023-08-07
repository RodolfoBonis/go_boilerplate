package main

import (
	"fmt"
	"github.com/RodolfoBonis/go_boilerplate/core/config"
	"github.com/RodolfoBonis/go_boilerplate/core/middlewares"
	"github.com/RodolfoBonis/go_boilerplate/core/services"
	"github.com/RodolfoBonis/go_boilerplate/core/utils"
	"github.com/RodolfoBonis/go_boilerplate/docs"
	"github.com/RodolfoBonis/go_boilerplate/routes"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	app := gin.New()

	err := app.SetTrustedProxies([]string{})

	if err != nil {
		log.Fatal(err)
	}

	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(gin.ErrorLogger())

	app.Use(middlewares.LogMiddleware())

	routes.InitializeRoutes(app)

	runPort := fmt.Sprintf(":%s", config.EnvPort())

	err = app.Run(runPort)

	if err != nil {
		panic(err)
	}

}

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
	})

	labels := map[string]string{
		"source":      config.EnvServiceName(),
		"environment": config.EnvironmentConfig(),
	}

	utils.NewLokiService(config.EnvGrafana(), labels)

	config.LoadEnvVars()

	services.InitializeOAuthServer()

	docs.SwaggerInfo.Title = "Internal API"
	docs.SwaggerInfo.Description = "A Internal API to management internal applications"
	docs.SwaggerInfo.Version = "0.0.1"
	docs.SwaggerInfo.Host = fmt.Sprintf("http://localhost:%s", config.EnvPort())
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
