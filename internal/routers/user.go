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
	rg.Use(func(c *gin.Context) { c.Set("uuid", "8dd97aca-a279-4438-b0b0-588601ffcd6e") })

	// check username duplicate
	rg.POST("/check", h.CheckDuplicate)
	// add new user
	rg.POST("", h.Create)
	// update user
	rg.POST("/update/socialAccount", h.UpdateSocialAccountById)
	// update user
	rg.POST("/update/address", h.UpdateAddressById)
	// delete user
	rg.DELETE("/delete", h.DeleteById)
	// get user
	rg.POST("/info", h.GetById)
}
