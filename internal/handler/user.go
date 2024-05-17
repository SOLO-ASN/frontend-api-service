package handler

import (
	"fmt"
	"net/http"
	"strings"

	"api-service/internal/dbEntity/cache"
	"api-service/internal/middleware/logger"
	"api-service/internal/model"
	"api-service/internal/response"
	"api-service/internal/retriever"
	"api-service/internal/types"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type IUserHandler interface {
	CheckDuplicate(c *gin.Context)
	Create(c *gin.Context)
	UpdateById(c *gin.Context)
	GetById(c *gin.Context)
	GetByName(c *gin.Context)
	UpdateSocialAccountById(c *gin.Context)
	UpdateEmailById(c *gin.Context)
	UpdateAddressById(c *gin.Context)
	DeleteById(c *gin.Context)
	CheckTwitterAccount(c *gin.Context)
}

func NewUserHandler() IUserHandler {
	return &userHandler{
		retriever: retriever.NewUserRetriever(
			model.GetDb(false),
			cache.Cache{},
		),
	}
}

type userHandler struct {
	retriever retriever.UserRetriever
}

func (u *userHandler) UpdateSocialAccountById(c *gin.Context) {
	//TODO implement me
	form := &types.UpdateSocialAccountRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		response.Error(c, response.WithCodeMessage{
			Code:    http.StatusBadRequest,
			Message: "invalid request parameters",
		}, err)
		return
	}
	// check login
	// todo add jwt checker
	if form.UserName == "" {
		response.Error(c, response.WithCodeMessage{
			Code:    http.StatusBadRequest,
			Message: "NOT_LOGIN",
		})
		return
	}

	//
	sAccount := &model.User{
		SocialAccount: model.SocialAccount{
			XAccountId:          form.XAccount.Id,
			XAccountName:        form.XAccount.Name,
			GithubAccountId:     form.GithubAccount.Id,
			GithubAccountName:   form.GithubAccount.Name,
			DiscordAccountId:    form.DiscordAccount.Id,
			DiscordAccountName:  form.DiscordAccount.Name,
			TelegramAccountId:   form.TelegramAccount.Id,
			TelegramAccountName: form.TelegramAccount.Name,
		}}
	err = u.retriever.UpdateSocialAccountById(c, form.UserName, sAccount)
	if err != nil {
		logger.DefaultLogger().Error("Update error: ", zap.Error(err))
		response.Error(c, response.WithCodeMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	response.Success(c, gin.H{"status": "update success"})
}

func (u *userHandler) UpdateEmailById(c *gin.Context) {
	// todo implement me
	form := &types.UpdateEmailRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		response.Error(c, response.WithCodeMessage{
			Code:    http.StatusBadRequest,
			Message: "invalid request parameters",
		}, err)
		return
	}

	// todo verify address and verification code

	// update email
	e := &model.User{
		Model: model.Model{
			ID: c.GetString("uuid"),
		},
		Email: &form.Email,
	}
	err = u.retriever.UpdateEmailById(c, e)
	if err != nil {
		response.Error(c, response.WithCodeMessage{
			Code:    http.StatusInternalServerError,
			Message: "update email error",
		})
	}

	response.Success(c, gin.H{"status": "update success"})

}

func (u *userHandler) UpdateAddressById(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u *userHandler) CheckDuplicate(c *gin.Context) {
	form := &types.CheckDuplicateRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		response.Error(c, response.WithCodeMessage{
			Code:    http.StatusBadRequest,
			Message: "invalid request parameters",
		}, err)
		return
	}
	user := &model.User{}
	user.Name = form.Name

	logger.DefaultLogger().Info("check user name duplicated", zap.String("name", form.Name))

	response.Success(c, gin.H{
		"duplicated_name": u.retriever.CheckDuplicateName(c, user),
	})
}

func (u *userHandler) Create(c *gin.Context) {
	//TODO implement me
	form := &types.CreateUserRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.DefaultLogger().Warn("BindJSON error: ", zap.Error(err))
		c.JSON(32001, "invalid params") // todo refactor
		return
	}
	
	user := &model.User{}
	user.Name = form.Name
	user.Avatar = form.Avatar
	if form.Email == "" {
		user.Email = nil
	}

	err = u.retriever.Create(c, user)
	if err != nil {
		logger.DefaultLogger().Error("Create error: ", zap.Error(err))
		if strings.Contains(err.Error(), "Error 1062") {
			response.Error(c, response.WithCodeMessage{ // todo refactor
				Code:    31062,
				Message: "duplicated user name",
			})
		} else {
			response.Error(c, response.WithCodeMessage{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		return
	}

	response.Success(c, gin.H{"uuid": user.ID})
}

func (u *userHandler) UpdateById(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u *userHandler) GetById(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u *userHandler) GetByName(c *gin.Context) {
	name := c.Param("name")

	user, err := u.retriever.GetByName(c, name)
	if err != nil {
		response.Error(c, response.WithCodeMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	response.Success(c, gin.H{
		"addressInfo": user,
	})
}

func (u *userHandler) DeleteById(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u *userHandler) CheckTwitterAccount(c *gin.Context) {
	// use mock data
	response.OutPut(c, response.WithCodeMessage{
		Code:    62001,
		Message: "LOGIN",
	}, &types.CheckTwitterAccountResponse{
		CheckTwitterAccount: &types.TwitterAccount{
			TwitterUserID:   "897858341040410624",
			TwitterUserName: "XUserName",
		},
		Verified: true,
	})
}
