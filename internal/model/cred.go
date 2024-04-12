package model

import "time"

type Cred struct {
	ID                    string         `gorm:"column:id;type:varchar(50);primary_key" json:"id"`
	Description           string         `gorm:"column:description;type:varchar(255)" json:"description"`
	Name                  string         `gorm:"column:name;type:varchar(50);NOT NULL" json:"name"`
	Type                  string         `gorm:"column:type;type:varchar(50);NOT NULL" json:"type"`
	CredType              string         `gorm:"column:cred_type;type:varchar(50);NOT NULL" json:"credType"`
	CredSource            string         `gorm:"column:cred_source;type:varchar(50);NOT NULL" json:"credSource"`
	ReferenceLink         string         `gorm:"column:reference_link;type:varchar(255);NOT NULL" json:"referenceLink"`
	LastUpdate            int64          `gorm:"column:last_update" json:"lastUpdate"`
	LastSync              int64          `gorm:"column:last_sync" json:"lastSync"`
	SyncStatus            string         `gorm:"column:sync_status;type:varchar(50);NOT NULL" json:"syncStatus"`
	CredContractNFTHolder string         `gorm:"column:cred_contract_nft_holder;type:varchar(100)" json:"credContractNFTHolder"`
	Chain                 string         `gorm:"column:chain;type:varchar(50);NOT NULL" json:"chain"`
	Eligible              int            `gorm:"column:eligible;NOT NULL" json:"eligible"`
	Subgraph              string         `gorm:"column:subgraph;type:varchar(255)" json:"subgraph"`
	DimensionConfig       string         `gorm:"column:dimension_config;type:varchar(50);NOT NULL" json:"dimensionConfig"`
	Value                 string         `gorm:"column:value;type:varchar(255)" json:"value"`
	CredMetadata          CredMetadata   `gorm:"column:cred_metadata;type:varchar(255)" json:"CredMetadata"`
	CommonInfo            CredCommonInfo `gorm:"column:common_info;type:varchar(255)" json:"commonInfo"`
	CreatedAt             time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt             time.Time      `gorm:"column:updated_at" json:"updatedAt"`
}

func (c *Cred) TableName() string {
	return "cred"
}
