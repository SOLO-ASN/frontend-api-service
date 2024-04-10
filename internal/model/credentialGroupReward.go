package model

type CredentialGroupReward struct {
	Expression  string `gorm:"column:expression;type:varchar(255)" json:"expression"`
	Eligible    bool   `gorm:"column:eligible" json:"eligible"`
	RewardCount int    `gorm:"column:reward_count" json:"rewardCount"`
	RewardType  string `gorm:"column:reward_type;type:varchar(50)" json:"rewardType"`
}

func (c *CredentialGroupReward) TableName() string {
	return "credentialGroupReward"
}
