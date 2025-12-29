package models

type GolonganBarang struct {
	Kode		string	`gorm:"column:kode"`
	Nama		string	`gorm:"column:nama"`
}

func (GolonganBarang) TableName() string {
	return "golongan_barang"
}