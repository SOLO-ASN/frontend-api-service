package model

import "gorm.io/datatypes"

type Space struct {
	Model
	Name                string `gorm:"column:name;type:varchar(50);NOT NULL" json:"name"`
	Owner               string `gorm:"column:owner;type:varchar(50);NOT NULL" json:"owner"`
	Thumbnail           string `gorm:"column:thumbnail;type:varchar(255);" json:"thumbnail"`
	Alias               string `gorm:"column:alias;type:varchar(50);" json:"alias"`
	IsVerified          bool   `gorm:"column:isVerified" json:"isVerified"`
	Info                string `gorm:"column:info;type:varchar(1500);" json:"info"`
	ActiveCampaignCount int    `gorm:"column:followersCount" json:"activeCampaignCount"`
	//TgeInfo          []TgeInfo        `json:"tgeInfo"`
	Links            datatypes.JSON `gorm:"column:links" json:"links"`
	Status           string         `gorm:"column:status;type:varchar(50)" json:"status"`
	FollowersCount   int            `gorm:"column:followersCount" json:"followersCount"`
	Backers          datatypes.JSON `gorm:"column:backers" json:"backers"`
	TokenID          string         `gorm:"column:tokenid;type:varchar(50);" json:"tokenid"` //用于索引token.;创建space时要注意
	Token            Token
	DiscordGuildID   string           `gorm:"column:discordGuildID;type:varchar(50)" json:"discordGuildID"`
	DiscordGuildInfo DiscordGuildInfo `gorm:"embedded;embeddedPrefix:discordGuildInfo" json:"discordGuildInfo"`
	Banner           string           `gorm:"column:banner;type:varchar(255)" json:"banner"`
	SeoImage         string           `gorm:"column:seoImage;type:varchar(255)" json:"seoImage"`
	Categories       datatypes.JSON   `gorm:"column:categories" json:"categories"`
	IsFollowing      bool
}

func (s *Space) TableName() string {
	return "space"
}
