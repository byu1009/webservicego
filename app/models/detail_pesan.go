package models

type DetailPesan struct {
	NoFaktur			string	`gorm:"column:no_faktur;primaryKey"`
	KodeBrng			string	`gorm:"column:kode_brng"`
	KodeSat				string	`gorm:"column:kode_sat"`
	Jumlah				string	`gorm:"column:jumlah"`
	HPesan				string	`gorm:"column:h_pesan"`
	Subtotal			string	`gorm:"column:subtotal"`
	Dis					string	`gorm:"column:dis"`
	BesarDis			string	`gorm:"column:besardis"`
	Total				string	`gorm:"column:total"`
	NoBatch				string	`gorm:"column:no_batch"`
	Jumlah2				string	`gorm:"column:jumlah2"`
	Kadaluarsa			string	`gorm:"column:kadaluarsa"`
}

func (DetailPesan) TableName() string {
	return "detail_pesan"
}