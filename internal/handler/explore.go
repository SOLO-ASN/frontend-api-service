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

type IExploreHandler interface {
	Create(c *gin.Context)
	Query(c *gin.Context)
}

type ExploreHandler struct {
	retriever retriever.ExploreRetriever
}

func NewExploreHandler() IExploreHandler {
	return &ExploreHandler{
		retriever: retriever.NewExploreRetriever(
			model.GetDb(false),
			&cache.Cache{}),
	}
}

func (h *ExploreHandler) Create(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *ExploreHandler) Query(c *gin.Context) {
	form := &types.ExploreQueryReqest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		response.Error(c, response.WithCodeMessage{
			Code:    http.StatusBadRequest,
			Message: "invalid request parameters",
		}, err)
		return
	}

	// db handle Explore query

	res, endCursor, hasNextPage, err := h.retriever.Query(c, *form, form.First, form.After)
	exploreResponse := exploreQuery(res, endCursor, hasNextPage)

	// assume we got all the data
	response.OutPut(c, response.WithCodeMessage{
		Code:    62001,
		Message: "Explore query success",
	}, exploreResponse)
}
func exploreQuery(exlopres *[]types.Exploredata, endCursor int, hasNextPage bool) types.ExploreQueryResponse {
	return types.ExploreQueryResponse{
		PageInfo: struct {
			EndCursor   int  `json:"endCursor"`
			HasNextPage bool `json:"hasNextPage"`
		}{
			EndCursor:   endCursor,
			HasNextPage: hasNextPage,
		},
		Explore: *exlopres,
	}
}
