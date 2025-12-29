package models

type IndustriFarmasi struct {
	KodeIndustri		string	`gorm:"column:kode_industri"`
	NamaIndustri		string	`gorm:"column:nama_industri"`
	Alamat				string	`gorm:"column:alamat"`
	Kota				string	`gorm:"column:kota"`
	NoTelp				string	`gorm:"column:no_telp"`
}

func (IndustriFarmasi) TableName() string {
	return "industri_farmasi"
}