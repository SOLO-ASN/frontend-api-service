package model

type CredentialParticipant struct {
	ID            int    `gorm:"column:id;type:varchar(255);primary_key" json:"id"`
	CredentialId  string `gorm:"column:credentialId;type:varchar(255)" json:"credentialId"`
	ParticipantId string `gorm:"column:participantId;type:varchar(255)" json:"participantId"`
	Status        bool   `gorm:"column:status;type:bool" json:"status"`
	Point         int    `gorm:"column:point;" json:"point"`
}

func (c *CredentialParticipant) TableName() string {
	return "credentialParticipant"
}
