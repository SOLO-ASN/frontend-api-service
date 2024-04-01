package model

type ChainAddress struct {
	Model

	UUID     string `gorm:"column:uuid;type:varchar(36)" json:"uuid"`
	MainAddr string `gorm:"column:main_addr;type:varchar(50);NOT NULL" json:"mainAddr"`

	EthLike []EthChainAddress
	Aptos   []AptosChainAddress
}

type EthChainAddress struct {
	Model

	BaseChain

	// todo add more info
}

type AptosChainAddress struct {
	Model

	BaseChain

	// todo add more info
}

func (b *ChainAddress) TableName() string {
	return "chainAddress"
}
