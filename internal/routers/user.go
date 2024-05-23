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
	//rg.Use(func(c *gin.Context) {c.Set("uuid", "68decb8d-944c-403a-b010-a342b0d9d5c3")})

	// check username duplicate
	rg.POST("/check", h.CheckDuplicate)
	// add new user
	rg.POST("", h.Create)
	// update user's social account
	rg.POST("/update/socialAccount", h.UpdateSocialAccountById)
	// update user's email
	rg.POST("/update/email", h.UpdateEmailById)
	// update user
	rg.POST("/update/address", h.UpdateAddressById)
	// delete user
	rg.DELETE("/delete", h.DeleteById)
	// get user
	rg.POST("/info", h.GetById)
	// get user by name
	rg.POST("/info/:name", h.GetByName)
	// check twitter account
	rg.POST("/checkTwitterAccount", h.CheckTwitterAccount)
	// send email
	rg.POST("/email/sendCode", h.SendCode)
	// verify email
	rg.POST("/email/verifyCode", h.VerifyCode)
	// parse fido list
	rg.POST("/parseFidoList", h.ParseFidoList)
}
