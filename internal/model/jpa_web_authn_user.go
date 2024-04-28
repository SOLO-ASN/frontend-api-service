package model

type Jpa_web_authn_user struct {
	Model
	Id       int    `gorm:"column:id;primary_key" json:"id"`
	Username string `gorm:"column:username;type:varchar(255)" json:"username"`
}

func (j *Jpa_web_authn_user) TableName() string {
	return "jpa_web_authn_user"
}
