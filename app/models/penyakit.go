package models

type Penyakit struct {
	KdPenyakit				string	`gorm:"column:kd_penyakit;primaryKey"`
	NmPenyakit				string	`gorm:"column:nm_penyakit"`
	CiriCiri				string	`gorm:"column:ciri_ciri"`
	Keterangan				string	`gorm:"column:keterangan"`
	KdKtg					string	`gorm:"column:KdKtg"`
}

func (Penyakit) TableName() string{
	return "penyakit"
}