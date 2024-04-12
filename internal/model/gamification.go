package model

type Gamification struct {
	ID   string `gorm:"column:id;type:varchar(50);primary_key"`
	Type string `gorm:"column:type;type:varchar(255);NOT NULL"`
}

func (g *Gamification) TableName() string {
	return "gamification"
}
