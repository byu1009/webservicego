package models

type DetailNotaJalan struct {
	NoRawat			string	`gorm:"column:no_rawat"`
	NamaBayar		string	`gorm:"column:nama_bayar"`
	BesarPpn		string	`gorm:"column:besarppn"`
	BesarBayar		string	`gorm:"column:besar_bayar"`
}

func (DetailNotaJalan) TableName() string {
	return "detail_nota_jalan"
}