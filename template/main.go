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

	newRelicConfig := config.NewRelicConfig()

	_middleware := middlewares.NewMonitoringMiddleware(newRelicConfig)

	app.Use(_middleware.NewRelicMiddleware())
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

	docs.SwaggerInfo.Title = "Go API Boilerplate"
	docs.SwaggerInfo.Description = "A Boilerplate to create go services using gin gonic"
	docs.SwaggerInfo.Version = "0.0.1"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", config.EnvPort())
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
