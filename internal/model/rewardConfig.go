package model

type RewardConfig struct {
	ID                   string                 `gorm:"column:id;type:varchar(50);primary_key" json:"id"`
	ParticipateCondition []ParticipateCondition `gorm:"foreignKey:ParticipateConditionID;NOT NULL" json:"ParticipateCondition"`

	ExprEntity []ExprEntity `gorm:"foreignKey:ExprEntityID;NOT NULL" json:"ExprEntity"`

	ExprReward         []ExprReward `gorm:"foreignKey:ExprRewardID;NOT NULL" json:"ExprReward"`
	ConditionalFormula string       `gorm:"column:conditional_formula;type:varchar(255);NOT NULL;default:ALL" json:"conditionalFormula"`
	Description        string       `gorm:"column:description" json:"description"`
	Eligible           bool         `gorm:"column:eligible;NOT NULL" json:"eligible"`
	RewardAttrVals     []string     `gorm:"column:reward_attr_vals;type:varchar(255)[]" json:"rewardAttrVals"`
}

func (r *RewardConfig) TableName() string {
	return "rewardConfig"
}
