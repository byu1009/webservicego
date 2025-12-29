package models

type KamarInap struct {
	NoRawat			string	`gorm:"column:no_rawat"`
	TglMasuk		string	`gorm:"column:tgl_masuk"`
	JamMasuk		string	`gorm:"column:jam_masuk"`
}

func (KamarInap) TableName() string {
	return "kamar_inap"
}