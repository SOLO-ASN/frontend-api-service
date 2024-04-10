package model

type DiscordGuildInfo struct {
	Icon string `gorm:"column:icon;type:varchar(50);NOT NULL" json:"icon"`
	ID   string `gorm:"column:id;type:varchar(50);primary_key" json:"id"`
	Link string `gorm:"column:link;type:varchar(255);NOT NULL" json:"link"`
	Name string `gorm:"column:name;type:varchar(50);NOT NULL" json:"name"`
}

func (d *DiscordGuildInfo) TableName() string {
	return "discordGuildInfo"
}
