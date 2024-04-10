package model

import (
	"time"

	"gorm.io/datatypes"
)

type Campaign struct {
	Model
	ID                    string              `gorm:"column:id;type:varchar(50);primary_key" json:"id"`
	Name                  string              `gorm:"column:name;type:varchar(255);NOT NULL" json:"name"`
	InWatchList           bool                `gorm:"column:in_watch_list;NOT NULL" json:"inWatchList"`
	InNewYearWatchList    bool                `gorm:"column:in_new_year_watch_list;NOT NULL" json:"inNewYearWatchList"`
	Thumbnail             string              `gorm:"column:thumbnail;type:varchar(255);NOT NULL" json:"thumbnail"`
	RewardTypes           string              `gorm:"column:rewardTypes;type:varchar(255)" json:"rewardTypes"`
	Type                  string              `gorm:"column:type;type:varchar(255);NOT NULL" json:"type"`
	Gamification          datatypes.JSON      `gorm:"column:gamification" json:"gamification"`
	Dao                   datatypes.JSON      `gorm:"column:dao" json:"dao"`
	IsBookmarked          bool                `gorm:"column:is_bookmarked;NOT NULL" json:"isBookmarked"`
	NumberID              int                 `gorm:"column:number_id;NOT NULL" json:"numberID"`
	Info                  string              `gorm:"column:info" json:"info"`
	UseCred               bool                `gorm:"column:use_cred;NOT NULL" json:"useCred"`
	Formula               string              `gorm:"column:formula" json:"formula"`
	GasType               string              `gorm:"column:gas_type;type:varchar(255);NOT NULL" json:"gasType"`
	CreatedAt             string              `gorm:"column:created_at" json:"createdAt"`
	RequirementInfo       string              `gorm:"column:requirement_info" json:"requirementInfo"`
	Description           string              `gorm:"column:description" json:"description"`
	Chain                 string              `gorm:"column:chain;type:varchar(255);NOT NULL" json:"chain"`
	StartTime             time.Time           `gorm:"column:start_time;NOT NULL" json:"startTime"`
	Status                string              `gorm:"column:status;type:varchar(255);NOT NULL" json:"status"`
	RequireEmail          bool                `gorm:"column:require_email;NOT NULL" json:"requireEmail"`
	RequireUsername       bool                `gorm:"column:require_username;NOT NULL" json:"requireUsername"`
	DistributionType      string              `gorm:"column:distribution_type;type:varchar(255);NOT NULL" json:"distributionType"`
	EndTime               time.Time           `gorm:"column:end_time;NOT NULL" json:"endTime"`
	Cap                   int                 `gorm:"column:cap;NOT NULL" json:"cap"`
	LoyaltyPoints         int                 `gorm:"column:loyalty_points;NOT NULL" json:"loyaltyPoints"`
	TokenRewardContract   TokenRewardContract `gorm:"column:token_reward_contract" json:"tokenRewardContract"`
	RecurringType         string              `gorm:"column:recurring_type;type:varchar(255);NOT NULL" json:"recurringType"`
	SpaceID               string              `gorm:"column:space_id;type:varchar(50);NOT NULL" json:"spaceId"`
	SmartbalanceDeposited bool                `gorm:"column:smartbalance_deposited;NOT NULL" json:"smartbalanceDeposited"`
	SmartbalancePreCheck  string              `gorm:"column:smartbalance_pre_check;type:varchar(255);NOT NULL" json:"smartbalancePreCheck"`
}

func (c *Campaign) TableName() string {
	return "campaign"
}
