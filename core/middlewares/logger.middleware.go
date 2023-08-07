package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/RodolfoBonis/go_boilerplate/core/config"
	"github.com/RodolfoBonis/go_boilerplate/core/entities"
	"github.com/RodolfoBonis/go_boilerplate/core/utils"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.EnvironmentConfig() != entities.Environment.Development {
			blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}

			jsonData, _ := io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonData))

			body := make(map[string]interface{})
			_ = json.Unmarshal(jsonData, &body)

			bytesBody, _ := json.MarshalIndent(body, "", "		")

			queryParams := c.Request.URL.Query()

			c.Writer = blw
			c.Next()

			statusCode := c.Writer.Status()

			logMessage := fmt.Sprintf("Requested \"%s\" - (%d)", c.FullPath(), statusCode)

			log := entities.Logs{
				Path:        c.FullPath(),
				BodyParams:  string(bytesBody),
				StatusCode:  strconv.Itoa(statusCode),
				Method:      c.Request.Method,
				QueryParams: queryParams.Encode(),
				Response:    blw.body.String(),
				Message:     logMessage,
			}

			if len(c.Errors) > 0 {
				for _, ginErr := range c.Errors {
					log.ErrorMessage = ginErr.Error()
					log.Message = ""

					logMap, _ := utils.StructToMap(log)
					utils.Logger.Error(logMessage, logMap)
				}

				c.Next()
				return
			}

			logMap, _ := utils.StructToMap(log)

			if statusCode != http.StatusOK && statusCode != http.StatusCreated {
				utils.Logger.Warning(logMessage, logMap)
				return
			}

			utils.Logger.Info(logMessage, logMap)
		}
	}
}
