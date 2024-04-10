package model

type ExprReward struct {
	Arithmetics       []string `gorm:"column:arithmetics;type:varchar(255)[]" json:"arithmetics"`
	ArithmeticFormula string   `gorm:"column:arithmetic_formula;type:varchar(255)" json:"arithmeticFormula"`
	RewardType        string   `gorm:"column:reward_type;type:varchar(255)" json:"rewardType"`
	RewardCount       int      `gorm:"column:reward_count" json:"rewardCount"`
	RewardVal         string   `gorm:"column:reward_val;type:varchar(255)" json:"rewardVal"`
}

func (e *ExprReward) TableName() string {
	return "exprReward"
}
