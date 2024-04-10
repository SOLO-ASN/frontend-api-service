package types

type SpacesQueryRequest struct {
	First         int    `json:"first"`
	After         string `json:"after"`
	Filter        string `json:"filter"`
	SearchString  string `json:"searchString"`
	SpaceListType string `json:"spaceListType"`
	VerifiedOnly  bool   `json:"verifiedOnly"`
}

List []struct {
	Filter                  string `json:"id"`
	Name                string `json:"name"`
	
} `json:"list"`


type SpacesQueryResponse struct {
	PageInfo struct {
		EndCursor   string `json:"endCursor"`
		HasNextPage bool   `json:"hasNextPage"`
	} `json:"pageInfo"`
	List []struct {
		Id                  string `json:"id"`
		Name                string `json:"name"`
		FollowersCount      int    `json:"followersCount"`
		ActiveCampaignCount int    `json:"activeCampaignCount"`
		IsVerified          bool   `json:"isVerified"`
		Thumbnail           string `json:"thumbnail"`
		TokenSymbol         string `json:"tokenSymbol"`
		Status              string `json:"status"`
		IsFollowing         bool   `json:"isFollowing"`
	} `json:"list"`
}

type FollowRequest struct {
	Id int `json:"id"`
}
