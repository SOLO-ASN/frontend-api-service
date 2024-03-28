package model

type User struct {
	Model
	Name         string       `gorm:"column:name;type:varchar(50);NOT NULL" json:"name"`
	Avatar       string       `gorm:"column:avatar;type:varchar(255);NOT NULL" json:"avatar"`
	UUID         string       `gorm:"column:uuid;type:varchar(36)" json:"uuid"`
	Email        string       `gorm:"column:email;type:varchar(50);NOT NULL" json:"email"`
	XAccount     string       `gorm:"column:x_account;type:varchar(50);NOT NULL" json:"xAccount"`
	ChainAddress ChainAddress `gorm:"column:chain_address;type:json;NOT NULL" json:"chainAddress"`
}

func (u *User) TableName() string {
	return "user"
}
