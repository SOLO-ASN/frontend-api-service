package model

type CampaignRequirementTwitterSapce struct {
	ID    string `gorm:"type:string;primary_key" json:"id"`
	Alias string `gorm:"column:alias;type:varchar(50);" json:"alias"`
}

func (c *CampaignRequirementTwitterSapce) TableName() string {
	return "campaignRequirementTwitterSapce"
}
