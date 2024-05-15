package model

type ChainAddress struct {
	UUID     string `gorm:"column:uuid;type:varchar(36)" json:"uuid"`
	MainAddr string `gorm:"column:main_addr;type:varchar(50);NOT NULL" json:"mainAddr"`

	//EthLike []EthChainAddress
}

type EthChainAddress struct {
	UUID     string `gorm:"column:uuid;type:varchar(36)" json:"uuid"`
	UserName string `gorm:"column:user_name;type:varchar(50);NOT"`

	BaseChain

	// todo add more info
}

type AptosChainAddress struct {
	BaseChain

	// todo add more info
}

func (b *ChainAddress) TableName() string {
	return "chain_address"
}
