package model

import "time"

type Cred struct {
	Model
	CredentialGroupId string `gorm:"column:credentialGroupId;type:varchar(55)" json:"credentialGroupId"`
	Description       string `gorm:"column:description;type:varchar(2000)" json:"description"`
	Name              string `gorm:"column:name;type:varchar(50);" json:"name"`
	CampaignId        string `gorm:"column:campaignId;type:varchar(55);" json:"campaignId"`
	Type              string `gorm:"column:type;type:varchar(50);" json:"type"`
	CredType          string `gorm:"column:cred_type;type:varchar(50);" json:"credType"`
	//CredSource            string `gorm:"column:cred_source;type:varchar(50); " json:"credSource"`
	ReferenceLink string    `gorm:"column:reference_link;type:varchar(255);" json:"referenceLink"`
	LastUpdate    time.Time `gorm:"column:last_update" json:"lastUpdate"`
	//LastSync              int64  `gorm:"column:last_sync" json:"lastSync"`
	//SyncStatus            string `gorm:"column:sync_status;type:varchar(50); " json:"syncStatus"`
	//CredContractNFTHolder string `gorm:"column:cred_contract_nft_holder;type:varchar(100)" json:"credContractNFTHolder"`
	Chain string `gorm:"column:chain;type:varchar(50);" json:"chain"`
	//Eligible              int    `gorm:"column:eligible; " json:"eligible"`
	//Subgraph              string `gorm:"column:subgraph;type:varchar(255)" json:"subgraph"`
	//DimensionConfig       string `gorm:"column:dimension_config;type:varchar(50); " json:"dimensionConfig"`
	//Value                 string `gorm:"column:value;type:varchar(255)" json:"value"`
	//CredMetadata          CredMetadata
	//CommonInfo            CredCommonInfo
}

func (c *Cred) TableName() string {
	return "Cred"
}
