package model

type WhitelistAddress struct {
	Address                           string `gorm:"column:address;type:varchar(255)" json:"address"`
	MaxCount                          int    `gorm:"column:max_count" json:"maxCount"`
	UsedCount                         int    `gorm:"column:used_count" json:"usedCount"`
	ClaimedLoyaltyPoints              int    `gorm:"column:claimed_loyalty_points" json:"claimedLoyaltyPoints"`
	CurrentPeriodClaimedLoyaltyPoints int    `gorm:"column:current_period_claimed_loyalty_points" json:"currentPeriodClaimedLoyaltyPoints"`
	CurrentPeriodMaxLoyaltyPoints     int    `gorm:"column:current_period_max_loyalty_points" json:"currentPeriodMaxLoyaltyPoints"`
}

func (w *WhitelistAddress) TableName() string {
	return "whitelistAddress"
}
