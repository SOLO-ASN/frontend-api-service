package types

import "api-service/internal/model"

type SpacesQueryRequest struct {
	First         int      `json:"first"`
	After         int      `json:"after"`
	Filters       []string `json:"filters"`
	SearchString  string   `json:"searchString"`
	SpaceListType string   `json:"spaceListType"`
	VerifiedOnly  bool     `json:"verifiedOnly"`
}

// Filters []struct {
// 	Filter string `json:"id"`
// 	Name   string `json:"name"`
// } `json:"filters"`

// Filters{
// 	{
// 		Filter:RwardType
// 		Name: Twitter Space
// 	},
// 	{
// 		Filter:RwardType
// 		Name: Quiz
// 	}
// }

type SpacesQueryResponse struct {
	PageInfo struct {
		EndCursor   int  `json:"endCursor"`
		HasNextPage bool `json:"hasNextPage"`
	} `json:"pageInfo"`
	// List []struct {
	// 	Id                  string `json:"id"`
	// 	Name                string `json:"name"`
	// 	FollowersCount      int    `json:"followersCount"`
	// 	ActiveCampaignCount int    `json:"activeCampaignCount"`
	// 	IsVerified          bool   `json:"isVerified"`
	// 	Thumbnail           string `json:"thumbnail"`
	// 	TokenSymbol         string `json:"tokenSymbol"`
	// 	Status              string `json:"status"`
	// 	IsFollowing         bool   `json:"isFollowing"`
	// } `json:"list"`
	Spaces []model.Space
}

type FollowRequest struct {
	Id int `json:"id"`
}
