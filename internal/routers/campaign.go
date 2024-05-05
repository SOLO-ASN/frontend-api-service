package routers

import (
	"api-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func init() {
	group_Campaign = append(group_Campaign, func(rg *gin.RouterGroup) {
		campaignGroup(rg, handler.NewCampaignHandler())
	})
}

func campaignGroup(rg *gin.RouterGroup, h handler.ICampaignHandler) {
	// use jwt handler here
	rg.Use(func(c *gin.Context) {})

	//
	rg.POST("/")

	//
	rg.POST("/query", h.Query)
	rg.POST("/create", h.Create)
}
