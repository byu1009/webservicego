package models

type Kecematan struct {
	KdKec	string	`gorm:"column:kd_kec"`
	NmKec	string	`gorm:"column:nm_kec"`
}

func (Kecematan) TableName() string {
	return "kecematan"
}