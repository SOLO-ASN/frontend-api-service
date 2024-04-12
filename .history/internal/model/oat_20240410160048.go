package model

type Oat struct {
	Model
}

func (o *Oat) TableName() string {
	return "oat"
}
