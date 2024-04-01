package model

type PageInfo struct {
	EndCursor int `gorm:"column:endCursor;NOT NULL" json:"endCursor"`
}

func (p *PageInfo) TableName() string {
	return "pageInfo"
}
