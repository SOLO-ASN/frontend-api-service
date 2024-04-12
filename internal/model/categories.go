package model

type Categories struct {
	Nft  string `gorm:"column:nft;type:varchar(50);" json:"nft"`
	Web3 string `gorm:"column:web3;type:varchar(50);" json:"web3"`
}

func (c *Categories) TableName() string {
	return "categories"
}
