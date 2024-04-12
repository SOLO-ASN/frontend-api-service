package routers

import (
	"api-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func init() {
	group_Campaigns = append(group_Campaigns, func(rg *gin.RouterGroup) {
		campaignsGroup(rg, handler.NewCampaignsHandler())
	})
}

func campaignsGroup(rg *gin.RouterGroup, h handler.ICampaignsHandler) {
	// use jwt handler here
	rg.Use(func(c *gin.Context) {})

	//
	rg.POST("/")

	//
	rg.POST("/query", h.Query)
}
