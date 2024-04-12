package model

type TokenReward struct {
	UserTokenAmount      string `gorm:"column:user_token_amount;type:varchar(50)" json:"userTokenAmount"`
	TokenAddress         string `gorm:"column:token_address;type:varchar(255)" json:"tokenAddress"`
	DepositedTokenAmount string `gorm:"column:deposited_token_amount;type:varchar(50)" json:"depositedTokenAmount"`
	TokenRewardId        int    `gorm:"column:token_reward_id" json:"tokenRewardId"`
	TokenDecimal         string `gorm:"column:token_decimal;type:varchar(50)" json:"tokenDecimal"`
	TokenLogo            string `gorm:"column:token_logo;type:varchar(255)" json:"tokenLogo"`
	TokenSymbol          string `gorm:"column:token_symbol;type:varchar(50)" json:"tokenSymbol"`
	Type                 string `gorm:"column:type;type:varchar(50)" json:"type"`
}

func (t *TokenReward) TableName() string {
	return "tokenReward"
}
