package services

import (
	"github.com/RodolfoBonis/go_boilerplate/core/config"
	"github.com/RodolfoBonis/go_boilerplate/core/errors"
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func StartAmqpConnection() *amqp.Connection {
	connectionString := config.EnvAmqpConnection()
	connection, err := amqp.Dial(connectionString)
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ")
		os.Exit(http.StatusInternalServerError)
	}

	return connection
}

func StartChannelConnection() (*amqp.Channel, *errors.AppError) {
	connection := StartAmqpConnection()
	channel, err := connection.Channel()
	if err != nil {
		return nil, errors.ServiceError("Failed to open a channel")
	}

	return channel, nil
}
