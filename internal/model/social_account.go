package model

type SocialAccount struct {
	XAccountId          string `gorm:"column:x_account_id;type:varchar(50);NOT NULL" json:"xAccountId"`
	XAccountName        string `gorm:"column:x_account_name;type:varchar(50);NOT NULL" json:"xAccountName"`
	GithubAccountId     string `gorm:"column:github_account_id;type:varchar(50);NOT NULL" json:"githubAccountId"`
	GithubAccountName   string `gorm:"column:github_account_name;type:varchar(50);NOT NULL" json:"githubAccountName"`
	DiscordAccountId    string `gorm:"column:discord_account_id;type:varchar(50);NOT NULL" json:"discordAccountId"`
	DiscordAccountName  string `gorm:"column:discord_account_name;type:varchar(50);NOT NULL" json:"discordAccountName"`
	TelegramAccountId   string `gorm:"column:telegram_account_id;type:varchar(50);NOT NULL" json:"telegramAccountId"`
	TelegramAccountName string `gorm:"column:telegram_account_name;type:varchar(50);NOT NULL" json:"telegramAccountName"`
}
