package response

import (
	"api-service/internal/middleware/logger"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type WithCodeMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func OutPut(c *gin.Context, cm WithCodeMessage, data ...interface{}) {
	resp := &Result{
		Code: cm.Code,
		Msg:  cm.Message,
	}
	if data == nil {
		resp.Data = &struct{}{}
	} else {
		resp.Data = data[0]
	}
	writeJSON(c, http.StatusOK, resp)
}

func Success(c *gin.Context, data ...interface{}) {
	OutPut(c, WithCodeMessage{Code: 62001, Message: "success"}, data...)
}

func Error(c *gin.Context, err WithCodeMessage, data ...interface{}) {
	OutPut(c, err, data...)
}

func writeJSON(c *gin.Context, code int, data interface{}) {
	c.Writer.WriteHeader(code)
	addContentType(c.Writer, jsonContentType)
	err := json.NewEncoder(c.Writer).Encode(data)
	if err != nil {
		logger.DefaultLogger().Warn("json encode error: ", zap.Error(err))
	}
}

var jsonContentType = "application/json"

func addContentType(w http.ResponseWriter, contentType string) {
	header := w.Header()
	header.Add("Content-Type", contentType)
}
