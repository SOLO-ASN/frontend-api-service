package handler

import (
	"net/http"

	"api-service/internal/dbEntity/cache"
	"api-service/internal/model"
	"api-service/internal/response"
	"api-service/internal/retriever"
	"api-service/internal/types"

	"github.com/gin-gonic/gin"
)

type IImagesHandler interface {
	Create(c *gin.Context)
	Upload(c *gin.Context)
}

type imagesHandler struct {
	retriever retriever.ImagesRetriever
}

func NewImagesHandler() IImagesHandler {
	return &imagesHandler{
		retriever: retriever.NewImagesRetriever(
			model.GetDb(false),
			&cache.Cache{}),
	}
}

func (s imagesHandler) Create(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s imagesHandler) Upload(c *gin.Context) {
	//TODO implement me
	form := &types.ImageUploadRequest{}
	err := c.ShouldBindJSON(form)

	uploadPath := "/home/node/picture"
	if err != nil {
		response.Error(c, response.WithCodeMessage{
			Code:    http.StatusBadRequest,
			Message: "invalid request parameters",
		}, err)
		return
	}

	res, _ := s.retriever.Upload(c, *form, uploadPath)

	// assume we got all the data
	response.OutPut(c, response.WithCodeMessage{
		Code:    62001,
		Message: "SUCCESSED",
	}, res)

}
