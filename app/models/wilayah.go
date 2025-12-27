package models

type Wilayah struct {
	Kode		string 	`gorm:"column:kode;primaryKey"`
	Nama		string	`gorm:"column:nama"`
}

func (Wilayah) TableName() string {
	return "wilayah"
}