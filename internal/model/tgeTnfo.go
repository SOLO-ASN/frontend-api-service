package model

type TgeInfo struct {
	Status      string `gorm:"column:status;type:varchar(50);NOT NULL" json:"status"`
	HasNextPage bool   `gorm:"column:hasNextPage;NOT NULL" json:"hasNextPage"`
}

func (t *TgeInfo) TableName() string {
	return "tgeInfo"
}
