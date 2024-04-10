package model

type Token struct {
	ID     int    `gorm:"column:id;primary_key" json:"id"`
	Symbol string `gorm:"column:symbol;type:varchar(10)" json:"symbol"`
	Slug   string `gorm:"column:slug;type:varchar(50)" json:"slug"`
}

func (t *Token) TableName() string {
	return "token"
}
