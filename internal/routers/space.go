package routers

import (
	"api-service/internal/handler"
	"github.com/gin-gonic/gin"
)

func init() {
	group_Space = append(group_Space, func(rg *gin.RouterGroup) {
		spaceGroup(rg, handler.NewSpaceHandler())
	})
}

func spaceGroup(rg *gin.RouterGroup, h handler.ISpaceHandler) {
	// use jwt handler here
	rg.Use(func(c *gin.Context) {})

	//
	rg.POST("/", h.Create)

	rg.POST("/query", h.Query)

}
