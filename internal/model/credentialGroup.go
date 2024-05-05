package model

import "gorm.io/datatypes"

type CredentialGroup struct {
	Model
	Description    string         `gorm:"column:description;type:varchar(2000)" json:"description"`
	CredentialIds  datatypes.JSON `gorm:"column:credentialIds" json:"credentialIds"`
	Rewards        datatypes.JSON `gorm:"column:rewards" json:"rewards"`
	RewardAttrVals string         `gorm:"column:rewardAttrVals;type:varchar(55)" json:"rewardAttrVals"`
	Creds          datatypes.JSON `json:"creds" gorm:"-"`
}
type CredentialIds struct {
	Ids []string `json:"ids"`
}
type Rewards struct {
	IsToken bool   `json:"isToken"`
	IsPoint bool   `json:"isPoint"`
	IsRole  bool   `json:"isRole"`
	Points  string `json:"points"`
}

func (c *CredentialGroup) TableName() string {
	return "credentialGroup"
}
