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
	ActiveCampaignCount int    `gorm:"column:activeCampaignCount" json:"activeCampaignCount"`
	//TgeInfo          []TgeInfo        `json:"tgeInfo"`
	Links            datatypes.JSON `gorm:"column:links" json:"links"`
	Status           string         `gorm:"column:status;type:varchar(50)" json:"status"`
	Backers          datatypes.JSON `gorm:"column:backers" json:"backers"`
	TokenID          string         `gorm:"column:tokenid;type:varchar(50);" json:"tokenid"` //用于索引token.;创建space时要注意
	Token            Token
	DiscordGuildID   string         `gorm:"column:discordGuildID;type:varchar(50)" json:"discordGuildID"`
	Discordguildinfo datatypes.JSON `gorm:"colum:discordguildinfo" json:"discordguildinfo"`
	Banner           string         `gorm:"column:banner;type:varchar(255)" json:"banner"`
	SeoImage         string         `gorm:"column:seoImage;type:varchar(255)" json:"seoImage"`
	Categories       datatypes.JSON `gorm:"column:categories" json:"categories"`
	IsFollowing      bool           `json:"isFollowing" gorm:"-"`
	Followers        int            `gorm:"column:followers" json:"followers"`
	IsOwner          bool           `json:"isOwner" gorm:"-"`
}

func (s *Space) TableName() string {
	return "space"
}
