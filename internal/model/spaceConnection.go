package model

type SapceConnection struct {
	Model
	TotalCount int      `gorm:"column:name;NOT NULL" json:"totalCount"`
	PageInfo   PageInfo `gorm:"embedded;embeddedPrefix:pageInfo;NOT NULL" json:"pageInfo"`
	//Space      []Space  `gorm:"embedded;embeddedPrefix:space;NOT NULL" json:"space"`
}

func (s *SapceConnection) TableName() string {
	return "sapceConnection"
}
