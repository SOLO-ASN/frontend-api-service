package handler

import (
	"api-service/utils"

	"github.com/gin-gonic/gin"
	"net/http"
)

// checkHealthResponse
type checkHealthResponse struct {
	Status   string `json:"status"`
	Hostname string `json:"hostname"`
}

func CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, checkHealthResponse{Status: "Running", Hostname: utils.GetHostname()})
}
