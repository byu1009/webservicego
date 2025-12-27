package models

type IoAntrianFarmasi struct {
	NoReferensi			string	`gorm:"column:no_refrensi"`
	NoAntrian			string	`gorm:"column:no_antrian"`
	StatusAntrian		string	`gorm:"column:status_antrian"`
	Calltime			string	`gorm:"column:calltime"`
	StatusPasien		string	`gorm:"column:status_pasien"`
	StatusPanggil		string	`gorm:"column:status_panggil"`
	KategoriAntrian		string	`gorm:"column:kategori_antrian"`
	Tgl					string	`gorm:"column:tgl"`
}

func (IoAntrianFarmasi) TableName() string {
	return "io_antrian_farmasi"
}