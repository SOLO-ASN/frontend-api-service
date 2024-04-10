package model

import "time"

type CredCommonInfo struct {
	ParticipateEndTime time.Time `gorm:"column:participate_end_time;NOT NULL" json:"participateEndTime"`
	ModificationInfo   string    `gorm:"column:modification_info;type:varchar(255)" json:"modificationInfo"`
	Typename           string    `gorm:"column:__typename;type:varchar(50);NOT NULL" json:"__typename"`
}

func (c *CredCommonInfo) TableName() string {
	return "commonInfo"
}
