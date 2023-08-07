package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
	"time"
)

var (
	Logger *LokiService
)

// LokiService representa um serviço de envio de logs para o Loki.
type LokiService struct {
	lokiURL string
	labels  map[string]string
}

// NewLokiService cria uma nova instância do LokiService.
func NewLokiService(lokiURL string, labels map[string]string) {
	Logger = &LokiService{
		lokiURL: fmt.Sprintf("%s/loki/api/v1/push", lokiURL),
		labels:  labels,
	}
}

// Info envia um log de nível Info para o Loki.
func (ls *LokiService) Info(message string, jsonData ...map[string]interface{}) {
	ls.sendLog("info", message, jsonData...)
}

// Warning envia um log de nível Warning para o Loki.
func (ls *LokiService) Warning(message string, jsonData ...map[string]interface{}) {
	ls.sendLog("warning", message, jsonData...)
}

// Error envia um log de nível Error para o Loki.
func (ls *LokiService) Error(message string, jsonData ...map[string]interface{}) {
	ls.sendLog("error", message, jsonData...)
}

// Success envia um log de nível Success para o Loki.
func (ls *LokiService) Success(message string, jsonData ...map[string]interface{}) {
	ls.sendLog("success", message, jsonData...)
}

// sendLog envia o log para o Loki.
func (ls *LokiService) sendLog(level, message string, jsonData ...map[string]interface{}) {
	// Criação do payload no formato esperado pela API do Loki
	epochNano := fmt.Sprintf("%d", time.Now().UnixNano())

	logEntry := []interface{}{
		epochNano,
		fmt.Sprintf("%s - %s", level, message),
	}

	if len(jsonData) > 0 {
		// Adicionar o jsonData como labels ao logEntry
		addLabelsFromJSON(ls.labels, jsonData[0], "")
	}

	payload := map[string]interface{}{
		"streams": []map[string]interface{}{
			{
				"stream": ls.labels,
				"values": [][]interface{}{
					logEntry,
				},
			},
		},
	}

	logData, err := json.Marshal(payload)
	if err != nil {
		log.Errorf("Erro ao converter o log em JSON: %s", err)
		return
	}

	// Enviar o log para o Loki
	req, err := http.NewRequest("POST", ls.lokiURL, bytes.NewBuffer(logData))
	if err != nil {
		log.Errorf("Erro ao criar a requisição HTTP: %s", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Erro ao enviar o log para o Loki: %s", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		// Caso haja erro, ler o corpo da resposta HTTP
		body, _ := io.ReadAll(resp.Body)
		log.WithFields(log.Fields{
			"Status":     resp.StatusCode,
			"StatusText": resp.Status,
			"Response":   string(body),
		}).Error("Erro ao enviar o log para o Loki")
		return
	}

	log.WithFields(log.Fields{
		"Level":   level,
		"Message": message,
		"Labels":  ls.labels,
	}).Info("Log enviado com sucesso para o Loki")
}

// addLabelsFromJSON adiciona os dados de um JSON ao mapa de labels.
func addLabelsFromJSON(labels map[string]string, jsonData map[string]interface{}, prefix string) {
	for key, value := range jsonData {
		newKey := key
		if prefix != "" {
			newKey = fmt.Sprintf("%s.%s", prefix, key)
		}

		switch v := value.(type) {
		case string:
			labels[newKey] = v
		case float64:
			labels[newKey] = strconv.FormatFloat(v, 'f', -1, 64)
		case map[string]interface{}:
			addLabelsFromJSON(labels, v, newKey)
			// Adicione mais casos aqui conforme necessário para outros tipos de valor.
		}
	}
}
