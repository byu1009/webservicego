package models

type AkunBayar struct {
	NamaBayar		string	`gorm:"column:nama_bayar;primaryKey"`
	KdRek			string	`gorm:"column:kd_rek"`
	Ppn				string	`gorm:"column:ppn"`
}

func (AkunBayar) TableName() string {
	return "akun_bayar"
}