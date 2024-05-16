package model

type User struct {
	Model
	ChainAddress
	//ChainAddress ChainAddress `gorm:"column:chainAddress" json:"chainAddress"`
	Name   string  `gorm:"column:name;type:varchar(50);NOT NULL;unique_index;unique" json:"name"`
	Avatar string  `gorm:"column:avatar;type:varchar(255);NOT NULL" json:"avatar"`
	Email  *string `gorm:"column:email;type:varchar(50);unique" json:"email"`
	SocialAccount
}

func (u *User) TableName() string {
	return "user"
}
