package retriever

import (
	"api-service/internal/dbEntity/cache"
	"api-service/internal/model"
	"api-service/internal/types"
	"context"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type CampaignRetriever interface {
	//Create (c context.Context, space *model.Campaign) error
	//Query (c context.Context) ([]*model.Campaign, error)
	Query(c context.Context, queryRequest types.CampaignQueryReqest) (*types.CampaignQueryResponse, error)
}

type campaignRetriever struct {
	db    *gorm.DB
	cache *cache.Cache
}

func NewCampaignRetriever(db *gorm.DB, cache *cache.Cache) CampaignRetriever {
	return &campaignRetriever{
		db:    db,
		cache: cache,
	}
}
func (cam *campaignRetriever) Create(c context.Context, campaigns []model.Campaign) error {

	err := cam.db.WithContext(c).Create(campaigns).Error
	return err

}
func (cam *campaignRetriever) Query(c context.Context, queryRequest types.CampaignQueryReqest) (*types.CampaignQueryResponse, error) {
	var campaign model.Campaign
	var Campaignparticipant model.CampaignParticipant
	var Campaignparticipants []model.CampaignParticipant
	var participant model.User
	var participants []model.User
	var credentialGroupIds model.CredentialGroupIds
	var credentialGroup model.CredentialGroup
	var credential model.Cred
	var credentials []model.Cred
	var CredentialGroupResponse model.CredentialGroupResponse
	var CredentialGroupResponses []model.CredentialGroupResponse
	var Space []model.Space

	//var count int64
	deSession := cam.db.Session(&gorm.Session{})
	if err := deSession.First(&campaign, "id = ?", queryRequest.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 未找到记录
			return nil, err
		}
		// 发生了其他错误
		return nil, err
	}
	json.Unmarshal(campaign.CredentialGroups, &credentialGroupIds)

	for _, credentialGroupid := range credentialGroupIds.Ids {
		deSession = cam.db.Session(&gorm.Session{})
		deSession = deSession.Model(credentialGroup).Where("id = ?", credentialGroupid)
		deSession.Find(&credentialGroup)
		//credentialGroups = append(credentialGroups, credentialGroup)

		deSession = cam.db.Session(&gorm.Session{})
		deSession = deSession.Model(credential).Where("credentialGroupId = ?", credentialGroupid)
		deSession.Find(&credentials)
		credentialGroup.Creds, _ = json.Marshal(credentials)
		CredentialGroupResponse = model.CredentialGroupResponse{
			CredentialGroup: credentialGroup,
			//	Credentails:     credentials,
		}
		CredentialGroupResponses = append(CredentialGroupResponses, CredentialGroupResponse)
	}

	deSession = cam.db.Session(&gorm.Session{})
	deSession = deSession.Model(Campaignparticipant).Where("campaignId = ?", campaign.ID)
	// deSession.Count(&count)
	// fmt.Println(count)
	if err := deSession.Find(&Campaignparticipants).Error; err != nil {
		// 处理错误
	}
	//fmt.Println(Campaignparticipants)
	for _, user := range Campaignparticipants {
		//fmt.Println(user.ParticipantId)
		deSession = cam.db.Session(&gorm.Session{})
		deSession = deSession.Model(participant).Where("id = ?", user.ParticipantId)
		deSession.Find(&participant)
		participants = append(participants, participant)
	}
	campaign.CredentialGroupResponses, _ = json.Marshal(CredentialGroupResponses)
	campaign.Participants, _ = json.Marshal(participants)

	deSession = cam.db.Session(&gorm.Session{})
	deSession = deSession.Model(Space).Where("id = ?", campaign.SpaceID)
	if err := deSession.Find(&Space).Error; err != nil {
		// 处理错误
	}
	campaign.Space, _ = json.Marshal(Space)
	return &types.CampaignQueryResponse{Campaign: campaign}, nil
}
