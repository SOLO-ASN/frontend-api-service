package types

type TelegramIsFollowRequest struct {
	CampaignId   string `json:"campaignid"`
	CredentialId string `json:"credentialid"`
	Username     string `json:"username"`
}
