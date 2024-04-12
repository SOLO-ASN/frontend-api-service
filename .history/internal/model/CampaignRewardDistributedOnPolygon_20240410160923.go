package model

type CampaignRewardDistributedOnEthereum struct {
	Model
}

func (c *CampaignRewardDistributedOnEthereum) TableName() string {
	return "campaignRewardDistributedOnEthereum"
}
