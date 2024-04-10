package model

type Categories struct {
	Nft  string `gorm:"column:id;type:varchar(50);primary_key" json:"id"`
	Web3 string `gorm:"column:name;type:varchar(255);NOT NULL" json:"name"`
}

func (d *DAO) TableName() string {
	return "dao"
}
