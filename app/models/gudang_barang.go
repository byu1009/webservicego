package models

type GudangBarang struct {
	KodeBrng		string	`gorm:"column:kode_brng"`
	KdBangsal		string	`gorm:"column:kd_bangsal"`
	Stok			string	`gorm:"column:stok"`
	NoBatch			string	`gorm:"column:no_batch"`
	NoFaktur		string	`gorm:"column:no_faktur"`
}

func (GudangBarang) TableName() string {
	return  "gudang_barang"
}