package routers

import (
	"api-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func init() {
	group_Spaces = append(group_Spaces, func(rg *gin.RouterGroup) {
		spacesGroup(rg, handler.NewSpacesHandler())
	})
}

func spacesGroup(rg *gin.RouterGroup, h handler.ISpacesHandler) {
	// use jwt handler here
	rg.Use(func(c *gin.Context) {})

	//
	rg.POST("/", h.Create)

	rg.POST("/query", h.Query)

	rg.POST("/follow", h.Follow)

	rg.POST("/unfollow", h.UnFollow)
}
