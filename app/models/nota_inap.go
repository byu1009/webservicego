package models

import "time"

type NotaInap struct {
	NoRawat			string		`gorm:"column:no_rawat"`
	NoNota			string		`gorm:"column:no_nota"`
	Tanggal			string		`gorm:"column:tanggal"`
	Jam				time.Time	`gorm:"column:jam"`
	UangMuka		int			`gorm:"column:Uang_Muka"`
}

func (NotaInap) TableName() string {
	return "nota_inap"
}