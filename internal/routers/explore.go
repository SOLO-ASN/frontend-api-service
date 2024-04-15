package routers

import (
	"api-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func init() {
	group_Explore = append(group_Explore, func(rg *gin.RouterGroup) {
		ExploreGroup(rg, handler.NewExploreHandler())
	})
}

func ExploreGroup(rg *gin.RouterGroup, h handler.IExploreHandler) {
	// use jwt handler here
	rg.Use(func(c *gin.Context) {})

	//
	rg.POST("/")

	//
	rg.POST("/query", h.Query)
}
