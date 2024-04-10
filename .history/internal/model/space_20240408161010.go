package model

import "gorm.io/datatypes"

type Space struct {
	Model
	ID         string `gorm:"column:id;type:varchar(50);primary_key" json:"id"`
	Name       string `gorm:"column:name;type:varchar(50);NOT NULL" json:"name"`
	Owner      string `gorm:"column:owner;type:varchar(50);NOT NULL" json:"owner"`
	Thumbnail  string `gorm:"column:thumbnail;type:varchar(255);" json:"thumbnail"`
	Alias      string `gorm:"column:alias;type:varchar(50);" json:"alias"` //别名是否可以不要
	IsVerified bool   `gorm:"column:is_verified" json:"isVerified"`
	Info       string `gorm:"column:info;type:varchar(255);" json:"info"`
	//TgeInfo          []TgeInfo        `json:"tgeInfo"`
	Links            datatypes.JSON   `gorm:"column:links" json:"links"`
	Status           string           `gorm:"column:status;type:varchar(50)" json:"status"`
	FollowersCount   int              `gorm:"column:followers_count" json:"followersCount"`
	Backers          datatypes.JSON   `gorm:"column:backers" json:"backers"`
	TokenID          string           `gorm:"column:tokenID;type:varchar(50);" json:"tokenID"` //用于索引token
	Token            Token            `json:"token"`
	DiscordGuildID   string           `gorm:"column:discord_guild_id;type:varchar(50)" json:"discordGuildID"`
	DiscordGuildInfo DiscordGuildInfo `gorm:"embedded;embeddedPrefix:discord_guild_info_" json:"discordGuildInfo"`
	Banner           string           `gorm:"column:banner;type:varchar(255)" json:"banner"`
	SeoImage         string           `gorm:"column:seo_image;type:varchar(255)" json:"seoImage"`
	Categories       []string         `gorm:"column:categories" json:"categories"`
}

func (s *Space) TableName() string {
	return "space"
}
