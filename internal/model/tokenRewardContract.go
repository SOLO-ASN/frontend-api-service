package model

type TokenRewardContract struct {
	ID      string `gorm:"column:id;type:varchar(50);primary_key" json:"id"`
	Address string `gorm:"column:address;type:varchar(255);NOT NULL" json:"address"`
	Chain   string `gorm:"column:chain;type:varchar(50);NOT NULL" json:"chain"`
}

func (t *TokenRewardContract) TableName() string {
	return "tokenRewardContract"
}
