package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.package_name}}/core/config"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.package_name}}/core/errors"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.package_name}}/core/logger"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.package_name}}/core/middlewares"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.package_name}}/core/services"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.package_name}}/docs"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.package_name}}/routes"
)

func main() {
	app := gin.New()

	err := app.SetTrustedProxies([]string{})

	if err != nil {
		appError := errors.RootError(err.Error())
		logger.Log.Error(appError.Message, appError.ToMap())
		panic(err)
	}

	config.SentryConfig()

	_middleware := middlewares.NewMonitoringMiddleware()

	app.Use(_middleware.SentryMiddleware())
	app.Use(_middleware.LogMiddleware)

	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(gin.ErrorLogger())

	routes.InitializeRoutes(app)

	runPort := fmt.Sprintf(":%s", config.EnvPort())

	err = app.Run(runPort)

	if err != nil {
		appError := errors.RootError(err.Error())
		logger.Log.Error(appError.Message, appError.ToMap())
		panic(err)
	}

}

func init() {

	config.LoadEnvVars()

	logger.InitLogger()

	services.InitializeOAuthServer()

	// Use this for open connection with DataBase
	//appError := services.OpenConnection()
	//
	//if appError != nil {
	//	logger.Log.Error(appError.Message, appError.ToMap())
	//	panic(appError)
	//}

	// Use this for Run Yours migrations
	// services.RunMigrations()

	// Use this for open connection with RabbitMQ
	// services.StartAmqpConnection()

	docs.SwaggerInfo.Title = "{{cookiecutter.package_name}}"
	docs.SwaggerInfo.Description = "{{cookiecutter.application_description}}"
	versionFileName := "version.txt"
	if config.EnvironmentConfig() == entities.Environment.Production {
		versionFileName = "/version.txt"
	}

	version := "unknown"
	if content, err := os.ReadFile(versionFileName); err == nil {
		version = strings.TrimSpace(string(content))
	}
	host := "localhost"

	if config.EnvironmentConfig() == entities.Environment.Production {
		host = "{{cookiecutter.package_name}}.rodolfodebonis.com.br"
	}

	docs.SwaggerInfo.Host = host
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Version = version
	scheme := "http"

	if config.EnvironmentConfig() == entities.Environment.Production {
		scheme = "https"
	}

	docs.SwaggerInfo.Schemes = []string{scheme}
}
