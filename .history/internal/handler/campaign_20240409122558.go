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
	//TODO implement me
	panic("implement me")
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

	res := campaignQueryMockData()

	// assume we got all the data
	response.OutPut(c, response.WithCodeMessage{
		Code:    62001,
		Message: "campaign query success",
	}, res)
}
func campaignQuery(spaces *[]model.Space, endCursor int, hasNextPage bool) types.CampaignQueryResponse {
	return types.CampaignQueryResponse{}
}

// func campaignQueryMockData() []types.CampaignQueryResponse {
// 	return []types.CampaignQueryResponse{
// 		{
// 			Id:                "GCK5JUUjFn",
// 			Name:              "Project Galaxy Meme Contest #1 Winner",
// 			Type:              "Oat",
// 			Status:            "Expired",
// 			Thumbnail:         "https://cdn.galxe.com/galaxy/assets/galaxyspace/1653583425186120168.png",
// 			ParticipantsCount: 3,
// 		},
// 		{
// 			Id:                "GC4SjtTJ2f",
// 			Name:              "Galxe Radio Episode 60 Feat. Laika AI",
// 			Type:              "Oat",
// 			Status:            "Active",
// 			Thumbnail:         "https://cdn.galxe.com/galaxy/galxe/adef1dc4-97e1-4247-b929-502af976edba.png",
// 			ParticipantsCount: 993,
// 		},
// 		{
// 			Id:                "GC2D5kRt7L",
// 			Name:              "Community Art Contest",
// 			Type:              "Art Contest",
// 			Status:            "Ended",
// 			Thumbnail:         "https://cdn.galxe.com/galaxy/assets/galaxyspace/1653583425186120168.png",
// 			ParticipantsCount: 532,
// 		},
// 		{
// 			Id:                "GC4G7fJk8Z",
// 			Name:              "Galxe SDK Tutorial Completion",
// 			Type:              "Tutorial",
// 			Status:            "Active",
// 			Thumbnail:         "https://cdn.galxe.com/galaxy/assets/galaxyspace/1653583425186120168.png",
// 			ParticipantsCount: 780,
// 		},
// 		{
// 			Id:                "GC5H8iKl0Q",
// 			Name:              "Beta Testers for New Platform Features",
// 			Type:              "Beta Testing",
// 			Status:            "Recruiting",
// 			Thumbnail:         "https://cdn.galxe.com/galaxy/assets/galaxyspace/1653583425186120168.png",
// 			ParticipantsCount: 320,
// 		},
// 	}
// }
