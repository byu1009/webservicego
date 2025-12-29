package models

type Kabupaten struct {
	KdKab	string	`gorm:"column:kd_kab"`
	NmKab	string	`gorm:"column:nm_kab"`
}

func (Kabupaten) TableName() string {
	return "kabupaten"
}