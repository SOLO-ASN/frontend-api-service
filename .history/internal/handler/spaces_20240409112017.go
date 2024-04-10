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

type ISpacesHandler interface {
	Create(c *gin.Context)
	Query(c *gin.Context)
	Follow(c *gin.Context)
	UnFollow(c *gin.Context)
}

type spacesHandler struct {
	retriever retriever.SpacesRetriever
}

func NewSpacesHandler() ISpacesHandler {
	return &spacesHandler{
		retriever: retriever.NewSpacesRetriever(model.GetDb(false),
			&cache.Cache{}),
	}
}

func (s spacesHandler) Create(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s spacesHandler) Follow(c *gin.Context) {
	form := &types.FollowRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		response.Error(c, response.WithCodeMessage{
			Code:    http.StatusBadRequest,
			Message: "invalid request parameters",
		}, err)
		return
	}

	// db handle follow

	// assume we got all the data
	response.OutPut(c, response.WithCodeMessage{
		Code:    31002,
		Message: "NOT_LOGIN",
	})
}

func (s spacesHandler) UnFollow(c *gin.Context) {
	form := &types.FollowRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		response.Error(c, response.WithCodeMessage{
			Code:    http.StatusBadRequest,
			Message: "invalid request parameters",
		}, err)
		return
	}

	// db handle unfollow

	// assume we got all the data
	response.OutPut(c, response.WithCodeMessage{
		Code:    62001,
		Message: "unfollow success",
	})
}

func (s spacesHandler) Query(c *gin.Context) {
	form := &types.SpacesQueryRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		response.Error(c, response.WithCodeMessage{
			Code:    http.StatusBadRequest,
			Message: "invalid request parameters",
		}, err)
		return
	}

	// todo retrieve data from db
	res, endCursor, hasNextPage, err := s.retriever.Query(c, form.Filters, form.First, form.After)
	spacesQueryResponse := spacesQueryResponse(&res, endCursor, hasNextPage)
	// assume we got all the dataD
	response.OutPut(c, response.WithCodeMessage{
		Code:    62001,
		Message: "NOT_LOGIN",
	}, spacesQueryResponse)
}
func spacesQueryResponse(spaces *[]model.Space, endCursor int, hasNextPage bool) *types.SpacesQueryResponse {
	return &types.SpacesQueryResponse{
		PageInfo: struct {
			EndCursor   int  `json:"endCursor"`
			HasNextPage bool `json:"hasNextPage"`
		}{
			EndCursor:   endCursor,
			HasNextPage: hasNextPage,
		},
		Spaces: *spaces,
	}
}

// func spacesQueryResponse() *types.SpacesQueryResponse {
// 	return &types.SpacesQueryResponse{
// 		PageInfo: struct {
// 			EndCursor   string `json:"endCursor"`
// 			HasNextPage bool   `json:"hasNextPage"`
// 		}{
// 			EndCursor:   "6",
// 			HasNextPage: false,
// 		},
// 		List: []struct {
// 			Id                  string `json:"id"`
// 			Name                string `json:"name"`
// 			FollowersCount      int    `json:"followersCount"`
// 			ActiveCampaignCount int    `json:"activeCampaignCount"`
// 			IsVerified          bool   `json:"isVerified"`
// 			Thumbnail           string `json:"thumbnail"`
// 			TokenSymbol         string `json:"tokenSymbol"`
// 			Status              string `json:"status"`
// 			IsFollowing         bool   `json:"isFollowing"`
// 		}{
// 			{
// 				Id:                  "1",
// 				Name:                "Galxe",
// 				FollowersCount:      307977,
// 				ActiveCampaignCount: 37,
// 				IsVerified:          true,
// 				Thumbnail:           "https://d257b89266utxb.cloudfront.net/galaxy/images/avatar/0x0b495174e4baabe771c6660be65054d2672ee577-1662470151406825713.png",
// 				TokenSymbol:         "GAL",
// 				Status:              "Standard",
// 				IsFollowing:         false,
// 			},
// 			{
// 				Id:                  "2",
// 				Name:                "Moso",
// 				FollowersCount:      45901,
// 				ActiveCampaignCount: 29,
// 				IsVerified:          true,
// 				Thumbnail:           "https://cdn.galxe.com/tooljet/Moso Logo Icon.png",
// 				TokenSymbol:         "USDC",
// 				Status:              "Standard",
// 				IsFollowing:         false,
// 			},
// 			{
// 				Id:                  "46841",
// 				Name:                "io.net",
// 				FollowersCount:      217227,
// 				ActiveCampaignCount: 6,
// 				IsVerified:          true,
// 				Thumbnail:           "https://cdn.galxe.com/galaxy/e99c44c2686442d39f4369bd7490a436/.png_thumbnail.webp",
// 				TokenSymbol:         "",
// 				Status:              "Standard",
// 				IsFollowing:         false,
// 			},
// 			{
// 				Id:                  "3533",
// 				Name:                "Taiko",
// 				FollowersCount:      751830,
// 				ActiveCampaignCount: 16,
// 				IsVerified:          true,
// 				Thumbnail:           "https://cdn.galxe.com/galaxy/avatar/233f5252-6c2a-4adf-8799-a310e27a016d.png",
// 				TokenSymbol:         "",
// 				Status:              "Standard",
// 				IsFollowing:         false,
// 			},
// 			{
// 				Id:                  "38918",
// 				Name:                "Covalent",
// 				FollowersCount:      1648,
// 				ActiveCampaignCount: 5,
// 				IsVerified:          true,
// 				Thumbnail:           "https://cdn.galxe.com/galaxy/a91a51a48f7e4295983ccbcca89caad8/.jpeg_thumbnail.webp",
// 				TokenSymbol:         "CQT",
// 				Status:              "Standard",
// 				IsFollowing:         false,
// 			},
// 			{
// 				Id:                  "44941",
// 				Name:                "Vanar",
// 				FollowersCount:      61400,
// 				ActiveCampaignCount: 2,
// 				IsVerified:          true,
// 				Thumbnail:           "https://cdn.galxe.com/galaxy/b1205d453957434bbd65519766907a26/.jpeg_thumbnail.webp",
// 				TokenSymbol:         "VANRY",
// 				Status:              "Standard",
// 				IsFollowing:         false,
// 			},
// 		},
// 	}
// }
