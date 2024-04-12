package model

type TgeInfo struct {
	Status  string `gorm:"column:status;type:varchar(50);NOT NULL" json:"status"`
	TgeTime int    `gorm:"column:TgeTime;NOT NULL" json:"TgeTime"`
}

func (t *TgeInfo) TableName() string {
	return "tgeInfo"
}
