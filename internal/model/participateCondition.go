package model

type ParticipateCondition struct {
	ConditionalFormula string `gorm:"column:conditional_formula;type:varchar(255);NOT NULL" json:"conditionalFormula"`
	Eligible           bool   `gorm:"column:eligible;NOT NULL" json:"eligible"`
	//ExprEntity         ExprEntity `gorm:"foreignkey:ExprEntityID" json:"ExprEntity"`
}

func (p *ParticipateCondition) TableName() string {
	return "participateCondition"
}
