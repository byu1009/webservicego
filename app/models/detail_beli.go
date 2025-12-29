package models

type DetailBeli struct {
	NoFaktur			string	`gorm:"column:no_faktur;primaryKey"`
	KodeBrng			string	`gorm:"column:kode_brng"`
	KodeSat				string	`gorm:"column:kode_sat"`
	Jumlah				string	`gorm:"column:jumlah"`
	HBeli				string	`gorm:"column:hbeli"`
	Subtotal			string	`gorm:"column:subtotal"`
	Dis					string	`gorm:"column:dis"`
	BesarDis			string	`gorm:"column:besar_dis"`
	Total				string	`gorm:"column:total"`
	NoBatch				string	`gorm:"column:no_batch"`
	Jumlah2				string	`gorm:"columm:jumlah2"`
	Kadaluarsa			string	`gorm:"column:kadaluarsa"`
}

func (DetailBeli) TableName() string {
	return "detail_beli"
}