package model

import (
	"gorm.io/datatypes"
)

type Campaign struct {
	Model
	Name  string `gorm:"column:name;type:varchar(50);NOT NULL" json:"name"`
	Alias string `gorm:"column:alias;type:varchar(50);" json:"alias"`
	//InWatchList           bool                `gorm:"column:in_watch_list;NOT NULL" json:"inWatchList"`
	//InNewYearWatchList    bool                `gorm:"column:in_new_year_watch_list;NOT NULL" json:"inNewYearWatchList"`
	Thumbnail    string         `gorm:"column:thumbnail;type:varchar(255)" json:"thumbnail"`
	RewardTypes  string         `gorm:"column:rewardTypes;type:varchar(255)" json:"rewardTypes"`
	Type         string         `gorm:"column:type;type:varchar(255)" json:"type"`
	Gamification datatypes.JSON `gorm:"column:gamification" json:"gamification"`
	Dao          datatypes.JSON `gorm:"column:dao" json:"dao"`
	// TwitterSpace          bool                `gorm:"column:twitterSpace" json:"twitterSpace"`
	// TwitterEngagement     bool                `gorm:"column:twitterEngagement" json:"twitterEngagement"`
	// Oat                   bool                `gorm:"column:oat" json:"oat"`
	// Nft                   bool                `gorm:"column:nft" json:"nft"`
	// Ethereum              bool                `gorm:"column:ethereum" json:"ethereum"`
	// Polygon               bool                `gorm:"column:polygon" json:"polygon"`
	CredSources  string `gorm:"column:credSources;varchar(50)" json:"credSources"`
	IsBookmarked bool   `gorm:"column:isBookmarked" json:"isBookmarked"`
	NumberID     int    `gorm:"column:numberId" json:"numberID"`
	Info         string `gorm:"column:info;varchar(1500)" json:"info"`
	UseCred      bool   `gorm:"column:useCred" json:"useCred"`
	Formula      string `gorm:"column:formula" json:"formula"`
	GasType      string `gorm:"column:gasType;type:varchar(255)" json:"gasType"`

	RequirementInfo          string         `gorm:"column:requirementInfo" json:"requirementInfo"`
	Description              string         `gorm:"column:description" json:"description"`
	Chain                    string         `gorm:"column:chain;type:varchar(255)" json:"chain"`
	StartTime                int            `gorm:"column:startTime" json:"startTime"`
	Status                   string         `gorm:"column:status;type:varchar(50)" json:"status"`
	RequireEmail             bool           `gorm:"column:requireEmail" json:"requireEmail"`
	RequireUsername          bool           `gorm:"column:requireUsername" json:"requireUsername"`
	DistributionType         string         `gorm:"column:distributiontype;type:varchar(255)" json:"distributionType"`
	EndTime                  int            `gorm:"column:endTime" json:"endTime"`
	Cap                      int            `gorm:"column:cap" json:"cap"`
	LoyaltyPoints            int            `gorm:"column:loyaltyPoints" json:"loyaltyPoints"`
	TokenRewardContract      string         `gorm:"column:tokenRewardContract;type:varchar(50)" json:"tokenRewardContract"`
	TokenReward              datatypes.JSON `gorm:"column:tokenReward" json:"tokenReward"`
	RecurringType            string         `gorm:"column:recurringType;type:varchar(255)" json:"recurringType"`
	SpaceID                  string         `gorm:"column:spaceId;type:varchar(50)" json:"spaceId"`
	SmartbalanceDeposited    bool           `gorm:"column:smartbalanceDeposited" json:"smartbalanceDeposited"`
	SmartbalancePreCheck     string         `gorm:"column:smartbalancePreCheck;type:varchar(255)" json:"smartbalancePreCheck"`
	ParticipantsCount        int            `gorm:"column:participantsCount;" json:"participantsCount"`
	CredentialGroups         datatypes.JSON `gorm:"column:credentialGroups" json:"credentialGroups"`
	DiscordRole              datatypes.JSON `gorm:"column:discordRole" json:"discordRole"`
	Participants             datatypes.JSON `gorm:"-"`
	CredentialGroupResponses datatypes.JSON `gorm:"-"`
	Space                    datatypes.JSON `gorm:"-"`
	TelegramBotApi           string         `gorm:"column:telegramBotApi;type:varchar(400)"`
	TelegramChatId           string         `gorm:"column:telegramChatId;type:varchar(400)" json:"telegramChatId"`
}
type CredentialGroupResponse struct {
	CredentialGroup CredentialGroup `json:"credentialGroup"`
	//Credentails     []Cred          `json:"Credentails"`
}
type CredentialGroupIds struct {
	Ids []string `json:"ids"`
}

func (c *Campaign) TableName() string {
	return "campaign"
}
