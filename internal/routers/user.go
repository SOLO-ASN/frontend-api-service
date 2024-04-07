package routers

import (
	"api-service/internal/handler"
	"github.com/gin-gonic/gin"
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

	// add new user
	rg.POST("/user", h.Create)
	// update user
	rg.PUT("/user/:id", h.UpdateById)
	// delete user
	rg.DELETE("/user/:id", h.DeleteById)
	// get user
	rg.GET("/user/:id", h.GetById)
}
