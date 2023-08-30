package main

import (
	"fmt"
	"github.com/RodolfoBonis/go_boilerplate/core/config"
	"github.com/RodolfoBonis/go_boilerplate/core/services"
	"github.com/RodolfoBonis/go_boilerplate/core/utils"
	"github.com/RodolfoBonis/go_boilerplate/docs"
	"github.com/RodolfoBonis/go_boilerplate/routes"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmgin"
	"time"
)

func main() {
	app := gin.New()

	err := app.SetTrustedProxies([]string{})

	if err != nil {
		log.Fatal(err)
	}

	app.Use(apmgin.Middleware(app))

	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(gin.ErrorLogger())

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

	config.LoadEnvVars()

	utils.InitLogger()

	services.InitializeOAuthServer()

	// Use this for open connection with DataBase
	// services.OpenConnection()

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
