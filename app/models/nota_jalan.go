package models

import "time"

type NotaJalan struct {
	NoRawat		string		`gorm:"column:no_rawat"`
	NoNota		string		`gorm:"column:no_nota"`
	Tanggal		string		`gorm:"column:tanggal"`
	Jam			time.Time	`gorm:"column:jam"`
}

func (NotaJalan) TableName() string {
	return "nota_jalan"
}