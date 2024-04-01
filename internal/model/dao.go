package model

type DAO struct {
	ID         string `gorm:"column:id;type:varchar(50);primary_key" json:"id"`
	Name       string `gorm:"column:name;type:varchar(255);NOT NULL" json:"name"`
	Logo       string `gorm:"column:logo;type:varchar(255);NOT NULL" json:"logo"`
	Alias      string `gorm:"column:alias;type:varchar(255);NOT NULL" json:"alias"`
	IsVerified bool   `gorm:"column:is_verified;NOT NULL" json:"is_verified"`
}

func (d *DAO) TableName() string {
	return "dao"
}
