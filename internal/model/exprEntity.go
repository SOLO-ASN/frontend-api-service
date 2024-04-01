package model

type ExprEntity struct {
	Cred            Cred     `gorm:"column:cred;type:varchar(255)" json:"cred"`
	Attrs           []string `gorm:"column:attrs;type:varchar(255)[]" json:"attrs"`
	AttrFormula     string   `gorm:"column:attr_formula;type:varchar(255)" json:"attrFormula"`
	Eligible        bool     `gorm:"column:eligible" json:"eligible"`
	EligibleAddress string   `gorm:"column:eligible_address;type:varchar(255)" json:"eligibleAddress"`
}

func (e *ExprEntity) TableName() string {
	return "exprEntity"
}
