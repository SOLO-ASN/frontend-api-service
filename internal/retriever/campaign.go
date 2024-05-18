package retriever

import (
	"api-service/internal/dbEntity/cache"
	"api-service/internal/model"
	"api-service/internal/types"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-telegram/bot"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CampaignRetriever interface {
	//Create (c context.Context, space *model.Campaign) error
	//Query (c context.Context) ([]*model.Campaign, error)
	Create(c context.Context, request types.CampaignCreateReqest) (error, string)
	Query(c context.Context, queryRequest types.CampaignQueryReqest) (*types.CampaignQueryResponse, error)
	TelegramisFollow(ctx context.Context, request types.TelegramIsFollowRequest) (string, error)
	IsComplete(c context.Context, queryRequest types.CmapaignIsCompleteRequst) (string, error)
	IsCredentialComplete(c context.Context, queryRequest types.IsCredentialCompleteRequst) (string, error)
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
	var unixTimestamp2 int64
	var unixTimestamp3 int64
	unixTimestamp := time.Now().Unix()
	campaign.SpaceID = request.SpaceID
	campaign.Name = request.Name
	campaign.Thumbnail = request.Thumbnail
	campaign.TelegramBotApi = request.TelegramBotApi
	campaign.TelegramChatId = request.TelegramChatId
	u, err := uuid.NewRandom()
	if err != nil {
		// 处理错误
		fmt.Println("生成UUID时发生错误:", err)
		return nil, "FAILED"
	}
	campaign.ID = u.String()
	campaign.Description = request.Description
	//campaign.TokenReward = request.TokenReward
	campaign.StartTime = request.StartTime
	campaign.EndTime = request.EndTime
	campaign.RewardTypes = request.RewardTypes
	campaign.CreatedAt = int(unixTimestamp)
	//campaign.DiscordRole = request.DiscordRole
	var cregroupids []string

	err = json.Unmarshal([]byte(request.CredentialGroups), &credgroups)
	if err != nil {
		fmt.Println(err)
	}
	unixTimestamp2 = time.Now().Unix()
	unixTimestamp3 = time.Now().Unix()
	for _, cregroup := range credgroups {

		u1, err := uuid.NewRandom()
		if err != nil {
			// 处理错误
			fmt.Println("生成UUID时发生错误:", err)
			return nil, "FAILED"
		}
		credentialGroup.CreatedAt = int(unixTimestamp2)
		unixTimestamp2 += 1
		credentialGroup.Description = cregroup.Description
		credentialGroup.Rewards, _ = json.Marshal(cregroup.Rewards)
		credentialGroup.ID = u1.String()
		deSession1 := cam.db.Session(&gorm.Session{})
		result1 := deSession1.Create(&credentialGroup)
		deSession1 = cam.db.Session(&gorm.Session{})
		deSession1 = deSession1.First(&credentialGroup)
		var creids []string

		unixTimestamp3 += 1
		for _, cre := range cregroup.Creds {

			u2, err := uuid.NewRandom()
			if err != nil {
				// 处理错误
				fmt.Println("生成UUID时发生错误:", err)
				return nil, "FAILED"
			}
			credential.ID = u2.String()
			credential.CreatedAt = int(unixTimestamp3)
			credential.UpdatedAt = int(unixTimestamp3)
			unixTimestamp3 += 1
			credential.CampaignId = campaign.ID

			credential.CredentialGroupId = credentialGroup.ID
			credential.Description = cre.Description
			credential.CredType = cre.CredType
			credential.Name = cre.Name
			credential.ReferenceLink = cre.ReferenceLink

			deSession2 := cam.db.Session(&gorm.Session{})
			result2 := deSession2.Create(&credential)
			var credential2 model.Cred
			deSession5 := cam.db.Session(&gorm.Session{})

			deSession5 = deSession5.Order("created_at desc")
			deSession5.First(&credential2)

			creids = append(creids, credential2.ID)

			if result2.Error != nil {
				return nil, "FAILED"
			}
		}

		credsofgourp.Ids = creids
		credentialGroup.CredentialIds, _ = json.Marshal(credsofgourp)
		deSession1.Save(&credentialGroup)
		cregroupids = append(cregroupids, credentialGroup.ID)
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

func (t *campaignRetriever) TelegramisFollow(ctx context.Context, request types.TelegramIsFollowRequest) (string, error) {
	var Campaign model.Campaign
	var user model.User
	var credentialParticipant model.CredentialParticipant
	var credential model.Cred
	var BotApi string
	var ChatId string
	var UserID string
	fmt.Println(3)
	success := "Follow Success"
	deSession := t.db.Session(&gorm.Session{})
	deSession = deSession.Model(Campaign).Where("id = ?", request.CampaignId)
	res := deSession.First(&Campaign)
	if res.Error != nil {
		return "NO_CAMPAIGN", nil
	}
	BotApi = Campaign.TelegramBotApi
	ChatId = Campaign.TelegramChatId

	deSession = t.db.Session(&gorm.Session{})
	deSession = deSession.Model(user).Where("name = ?", request.Username)
	res1 := deSession.First(&user)
	if res1.Error != nil {
		return "NO_USER", nil
	}

	UserID = user.TelegramAccountId
	intVal, _ := strconv.ParseInt(UserID, 10, 64)
	b, err := bot.New(BotApi)
	if nil != err {
		panic(err)
	}
	param := bot.GetChatMemberParams{
		ChatID: ChatId,
		UserID: intVal,
	}

	member, _ := b.GetChatMember(context.Background(), &param)
	if member.Owner == nil {
		return "NO_FOLLLOW", nil
	}
	deSession = t.db.Session(&gorm.Session{})
	deSession = deSession.Model(credential).Where("id = ?", request.CredentialId)
	res3 := deSession.First(&credential)

	if res3.Error != nil {
		return "NO_USER", nil
	}

	deSession = t.db.Session(&gorm.Session{})
	deSession = deSession.Model(credentialParticipant).Where("CredentialId = ? AND ParticipantId = ? ", credential.ID, user.ID)
	u, err := uuid.NewRandom()
	if err != nil {
		// 处理错误
		fmt.Println("生成UUID时发生错误:", err)
		return "FAILED", nil
	}
	res4 := deSession.First(&credentialParticipant)

	if res4.Error != nil {
		credParticipant := model.CredentialParticipant{
			ID:            u.String(),
			CredentialId:  credential.ID,
			ParticipantId: user.ID,
			Status:        true,
		}
		deSession = t.db.Session(&gorm.Session{})
		deSession = deSession.Model(credParticipant)
		result := deSession.Create(&credParticipant) // 通过数据的指针来创建
		if result.Error != nil {

			return "false", nil
		}
		return success, nil
	}

	credentialParticipant.Status = true

	deSession.Save(&credentialParticipant)
	return success, nil
}

func (cam *campaignRetriever) IsComplete(c context.Context, queryRequest types.CmapaignIsCompleteRequst) (string, error) {
	var campaign model.Campaign
	var Campaignparticipant model.CampaignParticipant

	var participant model.User
	var pointOfCampaign int

	var credentialGroupIds model.CredentialGroupIds
	var credentialGroup model.CredentialGroup
	var credential model.Cred
	var credentials []model.Cred
	var credentialGroupParticipant model.CredentialGroupParticipant

	var CredentialParticipant model.CredentialParticipant
	var CredentialParticipants []model.CredentialParticipant

	//var count int64
	deSession := cam.db.Session(&gorm.Session{})
	if err := deSession.First(&campaign, "id = ?", queryRequest.CampaigId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 未找到记录
			return "NOT_CAMPAIGN", nil
		}
		// 发生了其他错误
		return "NOT_CAMPAIGN", nil
	}
	json.Unmarshal(campaign.CredentialGroups, &credentialGroupIds)
	deSession = cam.db.Session(&gorm.Session{})
	deSession = deSession.Model(participant).Where("name = ?", queryRequest.Username)
	res2 := deSession.First(&participant)
	if res2.Error != nil {
		//return "NO_USER", nil
	}

	for _, credentialGroupid := range credentialGroupIds.Ids {

		deSession2 := cam.db.Session(&gorm.Session{})
		deSession2 = deSession2.Model(credentialGroupParticipant).Where("credentialGroupId= ? AND participantId= ?", credentialGroupid, participant.ID)
		res := deSession2.First(&credentialGroupParticipant)

		if res.Error != nil {
			var credentialGrouppart model.CredentialGroupParticipant
			u1, err := uuid.NewRandom()
			if err != nil {
				// 处理错误
				fmt.Println("生成UUID时发生错误:", err)
				return "FAILED", nil
			}
			credentialGrouppart = model.CredentialGroupParticipant{
				ID:                u1.String(),
				ParticipantId:     participant.ID,
				CredentialGroupId: credentialGroupid,
			}
			deSession1 := cam.db.Session(&gorm.Session{})
			deSession1.Create(&credentialGrouppart)

			deSession = cam.db.Session(&gorm.Session{})
			deSession = deSession.Model(credentialGroup).Where("id = ?", credentialGroupid)
			deSession.First(&credentialGroup)
			//credentialGroups = append(credentialGroups, credentialGroup)

			deSession = cam.db.Session(&gorm.Session{})
			deSession = deSession.Model(credential).Where("credentialGroupId = ?", credentialGroupid)
			deSession.Find(&credentials)
			for _, credential := range credentials {
				var credenParticipant model.CredentialParticipant
				deSession = cam.db.Session(&gorm.Session{})
				deSession = deSession.Model(credenParticipant).Where("credentialId = ? AND participantId= ? ", credential.ID, participant.ID)
				res3 := deSession.Find(&credenParticipant)
				if res3.Error != nil {
					return "NOT_COMPLETE", nil
				}

				if credenParticipant.Status != true {
					deSession1 := cam.db.Session(&gorm.Session{})
					credentialGrouppart.Status = false
					deSession1.Model(credentialGroupParticipant).Save(&credentialGrouppart)
					return "NOT_COMPLETE", nil
				}

				deSession1 := cam.db.Session(&gorm.Session{})
				credentialGrouppart.Status = true
				deSession1.Model(credentialGroupParticipant).Save(&credentialGrouppart)
			}

		}

		// deSession = cam.db.Session(&gorm.Session{})
		// deSession = deSession.Model(credentialGroup).Where("id = ?", credentialGroupid)
		// deSession.First(&credentialGroup)
		//credentialGroups = append(credentialGroups, credentialGroup)

		deSession = cam.db.Session(&gorm.Session{})
		deSession = deSession.Model(credential).Where("credentialGroupId = ?", credentialGroupid)
		deSession.Find(&credentials)
		for _, credential := range credentials {
			deSession = cam.db.Session(&gorm.Session{})
			deSession = deSession.Model(CredentialParticipant).Where("credentialId = ? AND participantId= ? ", credential.ID, participant.ID)
			res4 := deSession.First(&CredentialParticipant)
			if res4.Error != nil {
				return "NOT_COMPLETE", nil
			}
			for _, credParticipant := range CredentialParticipants {

				if !credParticipant.Status {

					credentialGroupParticipant.Status = false
					deSession1 := cam.db.Session(&gorm.Session{})
					deSession1.Model(credentialGroupParticipant).Save(&credentialGroupParticipant)
					return "NOT_COMPLETE", nil
				}
			}
			credentialGroupParticipant.Status = true
			var CreGroup model.CredentialGroup
			var Rewards model.Rewards
			deSession3 := cam.db.Session(&gorm.Session{})
			deSession3 = deSession3.Model(CreGroup)
			deSession3.First(&CreGroup)
			err := json.Unmarshal([]byte(CreGroup.Rewards), &Rewards)
			if err != nil {
				fmt.Println(err)
			}
			intVal, _ := strconv.ParseInt(Rewards.Points, 10, 64)
			credentialGroupParticipant.Point = int(intVal)
			pointOfCampaign += int(intVal)
			deSession1 := cam.db.Session(&gorm.Session{})
			deSession1.Save(&credentialGroupParticipant)
		}

	}
	deSession = cam.db.Session(&gorm.Session{})
	deSession = deSession.Model(Campaignparticipant).Where("campaignId=? AND participantId =? ", queryRequest.CampaigId, participant.ID)
	Campaignparticipant.Status = true
	Campaignparticipant.Point = pointOfCampaign
	deSession.Save(&Campaignparticipant)
	return "COMPLETE", nil
}
func (cam *campaignRetriever) IsCredentialComplete(c context.Context, queryRequest types.IsCredentialCompleteRequst) (string, error) {
	var credentialParticipant model.CredentialParticipant
	var participant model.User
	deSession := cam.db.Session(&gorm.Session{})
	deSession = deSession.Model(participant).Where("name = ?", queryRequest.Username)
	deSession.First(&participant)
	deSession = cam.db.Session(&gorm.Session{})
	deSession = deSession.Model(credentialParticipant).Where("CredentialId = ? AND ParticipantId= ? ", queryRequest.CredentialId, participant.ID)
	deSession.First(&credentialParticipant)
	if credentialParticipant.Status != true {
		return "NOT_COMPLETE", nil
	}
	return "COMPLETE", nil
}
