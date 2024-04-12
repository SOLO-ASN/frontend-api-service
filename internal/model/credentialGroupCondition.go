package model

type CredentialGroupCondition struct {
	Expression      string `gorm:"column:expression;type:varchar(255)" json:"expression"`
	Eligible        bool   `gorm:"column:eligible" json:"eligible"`
	EligibleAddress string `gorm:"column:eligible_address;type:varchar(255)" json:"eligibleAddress"`
}

func (c *CredentialGroupCondition) TableName() string {
	return "credentialGroupCondition"
}
