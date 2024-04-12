package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Model is the base model, like gorm.Model
type Model struct {
	ID        string     `gorm:"type:string;primary_key;auto_increment" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
}

type BaseChain struct {
	Address   string `gorm:"column:address;type:varchar(66);NOT NULL" json:"address"`
	ChainID   uint64 `gorm:"column:chain_id;type:int(11);NOT NULL" json:"chainId"`
	ChainName string `gorm:"column:chain_name;type:varchar(50);NOT NULL" json:"chainName"`
}

func (m *Model) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.New().String()
	return nil
}
