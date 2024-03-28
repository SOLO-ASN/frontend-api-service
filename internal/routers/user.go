package routers

import (
	"api-service/internal/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	group_UserId = append(group_UserId, func(rg *gin.RouterGroup) {
		userIdGroup(rg, handler.NewUserHandler())
	})
}

func userIdGroup(rg *gin.RouterGroup, h handler.IUserHandler) {
	// todo if you want to add jwt auth, you can add it here.
	// like this: rg.Use(jwtMiddleware.MiddlewareFunc())
	rg.Use(func(c *gin.Context) {})

	// test test test
	// todo delete me
	rg.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// add new user
	rg.POST("/user", h.Create)
	// update user
	rg.PUT("/user/:id", h.UpdateById)
	// delete user
	rg.DELETE("/user/:id", h.DeleteById)
	// get user
	rg.GET("/user/:id", h.GetById)
}
