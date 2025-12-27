package models

type Bangsal struct {
	KdBangsal		string	`gorm:"column:kd_bangsal"`
	NmBangsal		string	`gorm:"column:nm_bangsal"`
	Status			string	`gorm:"column:status"`
}

func (Bangsal) TableName() string {
	return "bangsal"
}