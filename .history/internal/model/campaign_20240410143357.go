package model

import (
	"time"

	"gorm.io/datatypes"
)

type Campaign struct {
	Model
	Name string `gorm:"column:name;type:varchar(50);NOT NULL" json:"name"`
	//InWatchList           bool                `gorm:"column:in_watch_list;NOT NULL" json:"inWatchList"`
	//InNewYearWatchList    bool                `gorm:"column:in_new_year_watch_list;NOT NULL" json:"inNewYearWatchList"`
	Thumbnail             string         `gorm:"column:thumbnail;type:varchar(255);NOT NULL" json:"thumbnail"`
	RewardTypes           string         `gorm:"column:rewardTypes;type:varchar(255)" json:"rewardTypes"`
	Type                  string         `gorm:"column:type;type:varchar(255);NOT NULL" json:"type"`
	Gamification          datatypes.JSON `gorm:"column:gamification" json:"gamification"`
	Dao                   datatypes.JSON `gorm:"column:dao" json:"dao"`
	TwitterSpace          bool
	CredSources           string              `gorm:"column:credSources;varchar(50)" json:"credSources"`
	IsBookmarked          bool                `gorm:"column:isBookmarked;NOT NULL" json:"isBookmarked"`
	NumberID              int                 `gorm:"column:numberId;NOT NULL" json:"numberID"`
	Info                  string              `gorm:"column:info" json:"info"`
	UseCred               bool                `gorm:"column:useCred;NOT NULL" json:"useCred"`
	Formula               string              `gorm:"column:formula" json:"formula"`
	GasType               string              `gorm:"column:gasType;type:varchar(255);NOT NULL" json:"gasType"`
	CreatedAt             string              `gorm:"column:createdAt" json:"createdAt"`
	RequirementInfo       string              `gorm:"column:requirementInfo" json:"requirementInfo"`
	Description           string              `gorm:"column:description" json:"description"`
	Chain                 string              `gorm:"column:chain;type:varchar(255);NOT NULL" json:"chain"`
	StartTime             time.Time           `gorm:"column:startTime;NOT NULL" json:"startTime"`
	Status                string              `gorm:"column:status;type:varchar(255);NOT NULL" json:"status"`
	RequireEmail          bool                `gorm:"column:requireEmail;NOT NULL" json:"requireEmail"`
	RequireUsername       bool                `gorm:"column:requireUsername;NOT NULL" json:"requireUsername"`
	DistributionType      string              `gorm:"column:distribution_type;type:varchar(255);NOT NULL" json:"distributionType"`
	EndTime               time.Time           `gorm:"column:endtime;NOT NULL" json:"endTime"`
	Cap                   int                 `gorm:"column:cap;NOT NULL" json:"cap"`
	LoyaltyPoints         int                 `gorm:"column:loyaltyPoints" json:"loyaltyPoints"`
	TokenRewardContract   TokenRewardContract `gorm:"column:tokenRewardContract" json:"tokenRewardContract"`
	RecurringType         string              `gorm:"column:recurringType;type:varchar(255)" json:"recurringType"`
	SpaceID               string              `gorm:"column:spaceId;type:varchar(50)" json:"spaceId"`
	SmartbalanceDeposited bool                `gorm:"column:smartbalanceDeposited" json:"smartbalanceDeposited"`
	SmartbalancePreCheck  string              `gorm:"column:smartbalancePreCheck;type:varchar(255);NOT NULL" json:"smartbalancePreCheck"`
}

func (c *Campaign) TableName() string {
	return "campaign"
}
