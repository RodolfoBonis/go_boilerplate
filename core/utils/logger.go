package utils

import (
	"bytes"
	"context"
	"encoding/json"
	_ "fmt"
	"github.com/RodolfoBonis/go_boilerplate/core/config"
	"github.com/RodolfoBonis/go_boilerplate/core/entities"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"time"
)

var (
	Logger *CustomLogger
)

// CustomLogger é uma estrutura que encapsula um logrus.Logger e um cliente elasticsearch.Client.
type CustomLogger struct {
	logger *logrus.Logger
	client *elasticsearch.Client
}

// LogData encapsula os dados do log.
type LogData struct {
	Level   string                 `json:"level"`
	Message string                 `json:"message"`
	Time    time.Time              `json:"time"`
	JSON    map[string]interface{} `json:"json,omitempty"`
}

// InitLogger cria uma nova instância do CustomLogger.
func InitLogger() {
	elasticUrl := config.EnvElasticSearch()
	cfg := elasticsearch.Config{
		Addresses: []string{elasticUrl},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}

	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)

	Logger = &CustomLogger{
		logger: logger,
		client: client,
	}
}

// Info envia um log de informação para o Elasticsearch e o logger.
func (cl *CustomLogger) Info(message string, jsonData ...map[string]interface{}) {
	cl.sendLog("info", message, jsonData...)
	cl.logger.Info(message)
}

// Warning envia um log de aviso para o Elasticsearch e o logger.
func (cl *CustomLogger) Warning(message string, jsonData ...map[string]interface{}) {
	cl.sendLog("warning", message, jsonData...)
	cl.logger.Warn(message)
}

// Error envia um log de erro para o Elasticsearch e o logger.
func (cl *CustomLogger) Error(message string, jsonData ...map[string]interface{}) {
	cl.sendLog("error", message, jsonData...)
	cl.logger.Error(message)
}

// Success envia um log de sucesso para o Elasticsearch e o logger.
func (cl *CustomLogger) Success(message string, jsonData ...map[string]interface{}) {
	cl.sendLog("success", message, jsonData...)
	cl.logger.Info(message)
}

// sendLog envia um log para o Elasticsearch.
func (cl *CustomLogger) sendLog(level, message string, jsonData ...map[string]interface{}) {
	logData := LogData{
		Level:   level,
		Message: message,
		Time:    time.Now(),
	}

	if len(jsonData) > 0 {
		logData.JSON = jsonData[0]
	}

	document, err := json.Marshal(logData)
	if err != nil {
		cl.logger.Error("Error marshaling log data:", err)
		return
	}

	indexName := config.EnvServiceName()
	environment := config.EnvironmentConfig()

	if environment != entities.Environment.Development {
		req := esapi.IndexRequest{
			Index:      indexName,
			DocumentID: "",
			Body:       bytes.NewReader(document),
			Refresh:    "true",
		}

		_, err = req.Do(context.Background(), cl.client)

		if err != nil {
			cl.logger.Error("Error marshaling log data:", err)
			return
		}
	}
}

// SetOutput sets the logger output.
func (cl *CustomLogger) SetOutput(out io.Writer) {
	cl.logger.SetOutput(out)
}

// SetLevel sets the logger level.
func (cl *CustomLogger) SetLevel(level logrus.Level) {
	cl.logger.SetLevel(level)
}
