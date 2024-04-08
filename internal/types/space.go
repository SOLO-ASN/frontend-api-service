package types

type SpaceCreateRequest struct{}

type SpaceQueryRequest struct {
	Alias string `json:"alias" binding:""`
}

type SpaceQueryResponse struct {
	Id             string   `json:"id"`
	Name           string   `json:"name"`
	Alias          string   `json:"alias"`
	IsVerified     bool     `json:"isVerified"`
	FollowersCount int      `json:"followersCount"`
	FollowersRank  string   `json:"followersRank"`
	IsFollowing    bool     `json:"isFollowing"`
	Thumbnail      string   `json:"thumbnail"`
	Info           string   `json:"info"`
	Categories     []string `json:"categories"`
	Token          struct {
		Id     int    `json:"id"`
		Symbol string `json:"symbol"`
		Slug   string `json:"slug"`
	} `json:"token"`
	Links struct {
		Discord   string `json:"Discord"`
		Github    string `json:"Github"`
		HomePage  string `json:"HomePage"`
		Instagram string `json:"Instagram"`
		Medium    string `json:"Medium"`
		Telegram  string `json:"Telegram"`
		TikTok    string `json:"TikTok"`
		Twitter   string `json:"Twitter"`
		YouTube   string `json:"YouTube"`
	} `json:"links"`
	Backers []struct {
		Name string `json:"name"`
		Icon string `json:"icon"`
	} `json:"backers"`
}
