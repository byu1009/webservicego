package models

type Kelurahan struct {
	KdKel	string	`gorm:"column:kd_kel"`
	NmKel	string	`gorm:"column:nm_kel"`
}

func (Kelurahan) TableName() string {
	return "kelurahan"
}