package models

type DetailNotaInap struct {
	NoRawat				string	`gorm:"column:no_rawat"`
	NamaBayar			string	`gorm:"column:nama_bayar"`
	BesarPpn			string	`gorm:"column:besarppn"`
	BesarBayar			string	`gorm:"column:besar_bayar"`
}

func (DetailNotaInap) TableName() string {
	return "detail_nota_inap"
}