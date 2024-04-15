package types

import "api-service/internal/model"

type ExploreQueryReqest struct {
	First        int      `json:"first"`
	After        int      `json:"after"`
	CredSources  []string `json:"credSources"`
	RewardTypes  []string `json:"rewardTypes"`
	Chains       []string `json:"chains"`
	Statuses     []string `json:"statuses"`
	ListType     string   `json:"listType"`
	SearchString string   `json:"searchString"`
}
type Exploredata struct {
	Campaign model.Campaign
	Space    model.Space
}
type ExploreQueryResponse struct {
	PageInfo struct {
		EndCursor   int  `json:"endCursor"`
		HasNextPage bool `json:"hasNextPage"`
	} `json:"pageInfo"`
	Explore []Exploredata
}

// type CampaignQueryResponse struct {
// 	Id                string `json:"id"`
// 	Name              string `json:"name"`
// 	Type              string `json:"type"`
// 	Status            string `json:"status"`
// 	Thumbnail         string `json:"thumbnail"`
// 	ParticipantsCount int    `json:"participantsCount"`
// }
