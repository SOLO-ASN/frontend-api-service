package model

type PageInfo struct {
	EndCursor int `gorm:"column:endCursor;NOT NULL" json:"endCursor"`
	TgeTime   int `gorm:"column:TgeTime;NOT NULL" json:"TgeTime"`
}

func (p *PageInfo) TableName() string {
	return "pageInfo"
}
