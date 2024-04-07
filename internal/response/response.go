package response

import (
	"api-service/internal/middleware/logger"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type ErrMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Success(c *gin.Context, data interface{}) {
	resp := &Result{
		Code: 62001,
		Msg:  "success",
	}
	if data == nil {
		resp.Data = &struct{}{}
	} else {
		resp.Data = data
	}

	writeJSON(c, http.StatusOK, resp)
}

func Error(c *gin.Context, err ErrMessage, data ...interface{}) {
	resp := &Result{
		Code: err.Code,
		Msg:  err.Message,
	}
	if data == nil {
		resp.Data = &struct{}{}
	} else {
		resp.Data = data
	}
	writeJSON(c, http.StatusOK, resp)
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
