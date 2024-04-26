package types

import "api-service/internal/model"

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
