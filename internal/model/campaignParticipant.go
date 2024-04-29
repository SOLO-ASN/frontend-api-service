package model

type CampaignParticipant struct {
	ID            int    `gorm:"type:string;aouto_increment;primary_key" json:"id"`
	CampaignId    string `gorm:"column:campaignId;type:varchar(255)" json:"campaignId"`
	ParticipantId string `gorm:"column:participantId;type:varchar(255)" json:"participantId"`
	Status        string `gorm:"column:status;type:varchar(55)" json:"status"`
	Point         int    `gorm:"column:point;" json:"point"`
	// ParticipantsCount       int `gorm:"column:participants_count;NOT NULL" json:"participantsCount"`
	// BountyWinnersCount      int `gorm:"column:bounty_winners_count;NOT NULL" json:"bountyWinnersCount"`
	// Participants_connection participants_connection
}
type participants_connection struct {
	addresses []Address
}

func (c *CampaignParticipant) TableName() string {
	return "campaignParticipant"
}
