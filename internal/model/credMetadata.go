package model

type CredMetadata struct {
	VisitLink string `gorm:"column:visit_link;type:varchar(255)" json:"visitLink"`
	Twitter   string `gorm:"column:twitter;type:varchar(255)" json:"twitter"`
	Worldcoin string `gorm:"column:worldcoin;type:varchar(255)" json:"worldcoin"`
}

func (c *CredMetadata) TableName() string {
	return "credMetadata"
}
