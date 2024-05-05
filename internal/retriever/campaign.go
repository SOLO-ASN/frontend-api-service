package retriever

import (
	"api-service/internal/dbEntity/cache"
	"api-service/internal/model"
	"api-service/internal/types"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type CampaignRetriever interface {
	//Create (c context.Context, space *model.Campaign) error
	//Query (c context.Context) ([]*model.Campaign, error)
	Create(c context.Context, request types.CampaignCreateReqest) (error, string)
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

type CregroupofCampaign struct {
	Id      int    `json:"id"`
	GroupId string `json:"groupId"`
}

func (cam *campaignRetriever) Create(c context.Context, request types.CampaignCreateReqest) (error, string) {
	var campaign model.Campaign
	var credential model.Cred
	var credentialGroup model.CredentialGroup
	var CregroupofCampaign model.CredentialGroupIds
	var credsofgourp model.CredentialIds
	var credgroups []types.CredentialGroup
	unixTimestamp := time.Now().Unix()
	campaign.SpaceID = request.SpaceID
	campaign.Name = request.Name
	data := request.Name
	h := sha256.New()
	h.Write([]byte(data))
	sha256Hash := hex.EncodeToString(h.Sum(nil))
	campaign.ID = sha256Hash
	campaign.Description = request.Description
	//campaign.TokenReward = request.TokenReward
	campaign.StartTime = request.StartTime
	campaign.EndTime = request.EndTime
	campaign.RewardTypes = request.RewardTypes
	campaign.CreatedAt = int(unixTimestamp)
	//campaign.DiscordRole = request.DiscordRole
	var cregroupids []string

	err := json.Unmarshal([]byte(request.CredentialGroups), &credgroups)
	if err != nil {
		fmt.Println(err)
	}
	for _, cregroup := range credgroups {
		data = cregroup.Description + campaign.ID
		h.Write([]byte(data))
		sha256Hash = hex.EncodeToString(h.Sum(nil))
		credentialGroup.CreatedAt = int(unixTimestamp)
		credentialGroup.Description = cregroup.Description
		credentialGroup.Rewards, _ = json.Marshal(cregroup.Rewards)
		credentialGroup.ID = sha256Hash
		var creids []string
		for _, cre := range cregroup.Creds {
			credential.CampaignId = campaign.ID
			credential.CredentialGroupId = credentialGroup.ID
			credential.Description = cre.Description
			credential.CredType = cre.CredType
			credential.Name = cre.Name
			credential.ReferenceLink = cre.ReferenceLink

			creids = append(creids, credentialGroup.ID)

			deSession2 := cam.db.Session(&gorm.Session{})
			result2 := deSession2.Create(&credential)
			if result2.Error != nil {
				return nil, "FAILED"
			}
		}
		credsofgourp.Ids = creids
		credentialGroup.Creds, _ = json.Marshal(credsofgourp)
		cregroupids = append(cregroupids, credentialGroup.ID)
		deSession1 := cam.db.Session(&gorm.Session{})
		result1 := deSession1.Create(&credentialGroup)

		if result1.Error != nil {
			return nil, "FAILED"
		}
	}
	CregroupofCampaign.Ids = cregroupids
	campaign.CredentialGroups, _ = json.Marshal(CregroupofCampaign)
	deSession := cam.db.Session(&gorm.Session{})
	result := deSession.Create(&campaign)
	if result.Error != nil {
		return nil, "FAILED"
	}
	return nil, "SECCESSED"

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
