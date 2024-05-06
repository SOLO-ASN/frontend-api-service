package routers

import (
	"api-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func init() {
	group_Images = append(group_Images, func(rg *gin.RouterGroup) {
		imagesGroup(rg, handler.NewImagesHandler())
	})
}

func imagesGroup(rg *gin.RouterGroup, h handler.IImagesHandler) {
	// use jwt handler here
	rg.Use(func(c *gin.Context) {})

	//
	//rg.POST("/", h.Create)

	rg.POST("/upload", h.Upload)

}
