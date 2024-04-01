package model

type CampaignParticipant struct {
	ParticipantsCount       int `gorm:"column:participants_count;NOT NULL" json:"participantsCount"`
	BountyWinnersCount      int `gorm:"column:bounty_winners_count;NOT NULL" json:"bountyWinnersCount"`
	Participants_connection participants_connection
}
type participants_connection struct {
	addresses []Address
}

func (c *CampaignParticipant) TableName() string {
	return "campaignParticipant"
}
