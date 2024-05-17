package model

type CredentialGroupParticipant struct {
	ID                string `gorm:"column:id;type:varchar(255);primary_key" json:"id"`
	CredentialGroupId string `gorm:"column:credentialGroupId;type:varchar(255)" json:"credentialGroupId"`
	ParticipantId     string `gorm:"column:participantId;type:varchar(255)" json:"participantId"`
	Status            bool   `gorm:"column:status;type:bool" json:"status"`
	Point             int    `gorm:"column:point;" json:"point"`
}

func (c *CredentialGroupParticipant) TableName() string {
	return "credentialGroupParticipant"
}
