package models

type KategoriBarang struct {
	Kode	string	`gorm:"column:kode"`
	Nama	string	`gorm:"column:nama"`
}

func (KategoriBarang) TableName() string {
	return "kategori_barang"
}