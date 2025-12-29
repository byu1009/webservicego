package models

type Deposit struct {
	NoDeposit				string	`gorm:"column:no_deposit;primaryKey"`
	NoRawat					string	`gorm:"column:no_rawat"`
	TglDeposit				string	`gorm:"column:tgl_deposit"`
	NamaBayar				string	`gorm:"column:nama_bayar"`
	BesarPpn				string	`gorm:"column:besar_ppn"`
	BesarDeposit			string	`gorm:"column:besar_deposit"`
	Nip						string	`gorm:"column:nip"`
	Keterangan				string	`gorm:"column:keterangan"`
}

func (Deposit) TableName() string {
	return "deposit"
}