package handler

import "github.com/gin-gonic/gin"

type IUserHandler interface {
	Create(c *gin.Context)
	UpdateById(c *gin.Context)
	GetById(c *gin.Context)
	DeleteById(c *gin.Context)
}

func NewUserHandler() IUserHandler {
	return &userHandler{}
}

type userHandler struct{}

func (u userHandler) Create(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u userHandler) UpdateById(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u userHandler) GetById(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u userHandler) DeleteById(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
