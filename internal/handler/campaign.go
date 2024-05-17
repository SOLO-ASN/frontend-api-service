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

type ICampaignHandler interface {
	Create(c *gin.Context)
	Query(c *gin.Context)
	TelegramisFollow(c *gin.Context)
	IsComplete(c *gin.Context)
	IsCredentialComplete(c *gin.Context)
}

type campaignHandler struct {
	retriever retriever.CampaignRetriever
}

func NewCampaignHandler() ICampaignHandler {
	return &campaignHandler{
		retriever: retriever.NewCampaignRetriever(
			model.GetDb(false),
			&cache.Cache{}),
	}
}

func (h *campaignHandler) Create(c *gin.Context) {
	form := &types.CampaignCreateReqest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		response.Error(c, response.WithCodeMessage{
			Code:    http.StatusBadRequest,
			Message: "invalid request parameters",
		}, err)
		return
	}
	err, res := h.retriever.Create(c, *form)
	response.OutPut(c, response.WithCodeMessage{
		Code: 62001,
	}, res)
}

func (h *campaignHandler) Query(c *gin.Context) {
	form := &types.CampaignQueryReqest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		response.Error(c, response.WithCodeMessage{
			Code:    http.StatusBadRequest,
			Message: "invalid request parameters",
		}, err)
		return
	}

	// db handle campaign query

	res, _ := h.retriever.Query(c, *form)

	// assume we got all the data
	response.OutPut(c, response.WithCodeMessage{
		Code:    62001,
		Message: "campaign query success",
	}, res)
}

func (h *campaignHandler) TelegramisFollow(c *gin.Context) {
	form := &types.TelegramIsFollowRequest{}

	err := c.ShouldBindJSON(form)
	if err != nil {
		response.Error(c, response.WithCodeMessage{
			Code:    http.StatusBadRequest,
			Message: "invalid request parameters",
		}, err)
		return
	}

	// db handle campaign query

	res, _ := h.retriever.TelegramisFollow(c, *form)

	// assume we got all the data
	response.OutPut(c, response.WithCodeMessage{
		Code:    62001,
		Message: "query success",
	}, res)
}

func (h *campaignHandler) IsComplete(c *gin.Context) {
	form := &types.CmapaignIsCompleteRequst{}

	err := c.ShouldBindJSON(form)
	if err != nil {
		response.Error(c, response.WithCodeMessage{
			Code:    http.StatusBadRequest,
			Message: "invalid request parameters",
		}, err)
		return
	}

	// db handle campaign query

	res, _ := h.retriever.IsComplete(c, *form)

	// assume we got all the data
	response.OutPut(c, response.WithCodeMessage{
		Code:    62001,
		Message: "query success",
	}, res)
}
func (h *campaignHandler) IsCredentialComplete(c *gin.Context) {
	form := &types.IsCredentialCompleteRequst{}

	err := c.ShouldBindJSON(form)
	if err != nil {
		response.Error(c, response.WithCodeMessage{
			Code:    http.StatusBadRequest,
			Message: "invalid request parameters",
		}, err)
		return
	}

	// db handle campaign query

	res, _ := h.retriever.IsCredentialComplete(c, *form)

	// assume we got all the data
	response.OutPut(c, response.WithCodeMessage{
		Code:    62001,
		Message: "query success",
	}, res)
}
