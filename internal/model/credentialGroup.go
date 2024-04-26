package model

import "gorm.io/datatypes"

type CredentialGroup struct {
	Model
	Description    string         `gorm:"column:description;type:varchar(2000)" json:"description"`
	CredentialIds  CredentialIds  `gorm:"column:credentialIds;type:json" json:"credentialIds"`
	Rewards        datatypes.JSON `gorm:"column:rewards" json:"rewards"`
	RewardAttrVals string         `gorm:"column:rewardAttrVals;type:varchar(55)" json:"rewardAttrVals"`
	Creds          datatypes.JSON `json:"creds"`
}
type CredentialIds struct {
	Ids []string `json:"ids"`
}

func (c *CredentialGroup) TableName() string {
	return "credentialGroup"
}
