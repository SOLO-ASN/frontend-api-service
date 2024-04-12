package types

import "api-service/internal/model"

type CampaignsQueryReqest struct {
	First        int           `json:"first"`
	After        int           `json:"after"`
	Alias        string        `json:"alias"`
	CredSources  []string      `json:"credSources"`
	RewardTypes  []string      `json:"rewardTypes"`
	Chains       []string      `json:"chains"`
	Statuses     []interface{} `json:"statuses"`
	ListType     string        `json:"listType"`
	SearchString string        `json:"searchString"`
}
type CampaignsQueryResponse struct {
	PageInfo struct {
		EndCursor   int  `json:"endCursor"`
		HasNextPage bool `json:"hasNextPage"`
	} `json:"pageInfo"`
	Campaigns []model.Campaign
}

// type CampaignQueryResponse struct {
// 	Id                string `json:"id"`
// 	Name              string `json:"name"`
// 	Type              string `json:"type"`
// 	Status            string `json:"status"`
// 	Thumbnail         string `json:"thumbnail"`
// 	ParticipantsCount int    `json:"participantsCount"`
// }
