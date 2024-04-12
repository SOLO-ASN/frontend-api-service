package model

type CampaignRequirementTwitterEngagement struct {
	ID    string `gorm:"type:string;primary_key" json:"id"`
	Alias string `gorm:"column:alias;type:varchar(50);" json:"alias"`
}

func (c *CampaignRequirementTwitterEngagement) TableName() string {
	return "campaignRequirementTwitterEngagement"
}
