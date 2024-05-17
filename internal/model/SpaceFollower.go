package model

type SpaceFollower struct {
	ID            int    `gorm:"type:string;aouto_increment;primary_key" json:"id"`
	SpaceId       string `gorm:"column:spaceId;type:varchar(255)" json:"spaceId"`
	ParticipantId string `gorm:"column:participantId;type:varchar(255)" json:"participantId"`
	IsFollowing   bool   `gorm:"column:isFollowing;" json:"isFollowing"`

	// ParticipantsCount       int `gorm:"column:participants_count;NOT NULL" json:"participantsCount"`
	// BountyWinnersCount      int `gorm:"column:bounty_winners_count;NOT NULL" json:"bountyWinnersCount"`
	// Participants_connection participants_connection
}
type SpaceFollowerAddresses struct {
	addresses []Address
}

func (c *SpaceFollower) TableName() string {
	return "spaceFollower"
}
