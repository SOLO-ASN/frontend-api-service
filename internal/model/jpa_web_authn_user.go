package model

type Jpa_web_authn_user struct {
	//Model
	Id        int    `gorm:"column:id;type:bigint;primary_key" json:"id"`
	CreatedAt int    `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt int    `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt int    `gorm:"column:deleted_at" json:"deletedAt"`
	Username  string `gorm:"column:username;type:varchar(255)" json:"username"`
}

func (j *Jpa_web_authn_user) TableName() string {
	return "jpa_web_authn_user"
}
