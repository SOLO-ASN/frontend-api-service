package model

type Address struct {
	Model
	Id     int    `gorm:"column:id;primary_key" json:"id"`
	Avatar string `gorm:"column:avatar;type:varchar(255)" json:"avatar"`
}

func (a *Address) TableName() string {
	return "address"
}
