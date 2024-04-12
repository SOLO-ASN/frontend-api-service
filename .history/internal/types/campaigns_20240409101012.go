package types

type CampaignsQueryReqest struct {
	First        int           `json:"first"`
	After        string        `json:"after"`
	Alias        string        `json:"alias"`
	CredSources  []string      `json:"credSources"`
	RewardTypes  []string      `json:"rewardTypes"`
	Chains       []string      `json:"chains"`
	Statuses     []interface{} `json:"statuses"`
	ListType     string        `json:"listType"`
	SearchString string        `json:"searchString"`
}

type CampaignsQueryResponse struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	Type              string `json:"type"`
	Status            string `json:"status"`
	Thumbnail         string `json:"thumbnail"`
	ParticipantsCount int    `json:"participantsCount"`
}
