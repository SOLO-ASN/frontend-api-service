package types

import (
	"api-service/internal/model"

	"gorm.io/datatypes"
)

type CampaignQueryReqest struct {
	// First        int           `json:"first"`
	// After        string        `json:"after"`
	// Alias        string        `json:"alias"`
	// CredSources  []string      `json:"credSources"`
	// RewardTypes  []string      `json:"rewardTypes"`
	// Chains       []string      `json:"chains"`
	// Statuses     []interface{} `json:"statuses"`
	// ListType     string        `json:"listType"`
	// SearchString string        `json:"searchString"`
	Id string `json:"id" binding:""`
}
type CredentialGroupResponse struct {
	CredentialGroup model.CredentialGroup
	Credentails     []model.Cred
}
type CampaignQueryResponse struct {
	// Id                string `json:"id"`
	// Name              string `json:"name"`
	// Type              string `json:"type"`
	// Status            string `json:"status"`
	// Thumbnail         string `json:"thumbnail"`
	// ParticipantsCount int    `json:"participantsCount"`
	Campaign model.Campaign
}

type CampaignCreateReqest struct {
	Name string `json:"name"`

	RewardTypes string `json:"rewardTypes"`

	Description string `json:"description"`

	StartTime int `json:"startTime"`

	EndTime int `json:"endTime"`

	LoyaltyPoints       int            `json:"loyaltyPoints"`
	TokenRewardContract string         `json:"tokenRewardContract"`
	TokenReward         datatypes.JSON `json:"tokenReward"`
	SpaceID             string         `json:"space"`

	CredentialGroups datatypes.JSON `json:"credentialGroups"`

	DiscordRole datatypes.JSON `json:"discordRole"`

	Thumbnail string `json:"thumbnail"`

	TelegramBotApi string `json:"telegramBotApi"`
	TelegramChatId string `json:"telegramChatId"`
}

type CampaignCreateResponse struct {
	Result string `json:"result"`
}

type Cred struct {
	Description string `json:"description"`
	Name        string `json:"name"`

	CredType string `json:"credType"`

	ReferenceLink string `json:"referenceLink"`
}

type CredentialGroup struct {
	Description string `json:"description"`

	Rewards        model.Rewards `json:"rewards"`
	RewardAttrVals string        `json:"rewardAttrVals"`
	Creds          []Cred        `json:"creds"`
}
