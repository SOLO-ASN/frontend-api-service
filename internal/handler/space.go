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

type ISpaceHandler interface {
	Create(c *gin.Context)
	Query(c *gin.Context)
}

type spaceHandler struct {
	retriever retriever.SpaceRetriever
}

func NewSpaceHandler() ISpaceHandler {
	return &spaceHandler{
		retriever: retriever.NewSpaceRetriever(
			model.GetDb(false),
			&cache.Cache{}),
	}
}

func (s spaceHandler) Create(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s spaceHandler) Query(c *gin.Context) {
	//TODO implement me
	form := &types.SpaceQueryRequest{}
	err := c.ShouldBindJSON(form)

	if err != nil {
		response.Error(c, response.WithCodeMessage{
			Code:    http.StatusBadRequest,
			Message: "invalid request parameters",
		}, err)
		return
	}
	// todo retrieve data from db

	res, _ := s.retriever.Query(c, form.Id)

	// assume we got all the data
	response.OutPut(c, response.WithCodeMessage{
		Code:    62001,
		Message: "NOT_LOGIN",
	}, res)
}

func spaceQueryResponseMockData() *types.SpaceQueryResponse {
	return &types.SpaceQueryResponse{
		Id:             "28",
		Name:           "Galxe",
		Alias:          "Galxe",
		IsVerified:     true,
		FollowersCount: 308016,
		FollowersRank:  "17",
		IsFollowing:    true,
		Thumbnail:      "https://d257b89266utxb.cloudfront.net/galaxy/images/avatar/0x0b495174e4baabe771c6660be65054d2672ee577-1662470151406825713.png",
		Info:           "Galxe is the leading platform for building Web3 community. With over 11 million unique users, Galxe has propelled the growth of Optimism, Polygon, Arbitrum, and more than 2900 partners with reward-based loyalty programs.",
		Categories:     []string{"NFT", "Web3", "DID", "Social", "Infrastructure"},
		Token: struct {
			Id     int    `json:"id"`
			Symbol string `json:"symbol"`
			Slug   string `json:"slug"`
		}{
			Id:     11877,
			Symbol: "GAL",
			Slug:   "galxe",
		},
		Links: struct {
			Discord   string `json:"Discord"`
			Github    string `json:"Github"`
			HomePage  string `json:"HomePage"`
			Instagram string `json:"Instagram"`
			Medium    string `json:"Medium"`
			Telegram  string `json:"Telegram"`
			TikTok    string `json:"TikTok"`
			Twitter   string `json:"Twitter"`
			YouTube   string `json:"YouTube"`
		}{
			Discord:   "https://discord.gg/galxe",
			Github:    "https://github.com/GalxeHQ",
			HomePage:  "galxe.com",
			Instagram: "https://www.instagram.com/galxehq",
			Medium:    "https://blog.galxe.com/",
			Telegram:  "t.me/Galxe",
			TikTok:    "https://www.tiktok.com/@galxehq",
			Twitter:   "https://twitter.com/Galxe",
			YouTube:   "https://www.youtube.com/@GalxeHQ",
		},
		Backers: []struct {
			Name string `json:"name"`
			Icon string `json:"icon"`
		}{
			{Name: "multicoincap", Icon: "https://pbs.twimg.com/profile_images/1532419804031852546/KMJZNS9c_normal.jpg"},
		},
	}
}
