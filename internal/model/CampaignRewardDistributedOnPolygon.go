package model

type CampaignRewardDistributedOnPolygon struct {
	ID    string `gorm:"type:string;primary_key" json:"id"`
	Alias string `gorm:"column:alias;type:varchar(50);" json:"alias"`
}

func (c *CampaignRewardDistributedOnPolygon) TableName() string {
	return "campaignRewardDistributedOnPolygon"
}
