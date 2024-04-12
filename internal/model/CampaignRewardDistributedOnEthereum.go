package model

type CampaignRewardDistributedOnEthereum struct {
	ID    string `gorm:"type:string;primary_key" json:"id"`
	Alias string `gorm:"column:alias;type:varchar(50);" json:"alias"`
}

func (c *CampaignRewardDistributedOnEthereum) TableName() string {
	return "campaignRewardDistributedOnEthereum"
}
