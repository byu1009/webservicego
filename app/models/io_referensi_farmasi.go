package models

type IoReferensiFarmasi struct {
	NoResep				string	`gorm:"column:no_resep"`
	Kodebooking			string	`gorm:"column:kodebooking"`
	Tanggal				string	`gorm:"column:tanggal"`
	Prefix				string	`gorm:"column:prefix"`
	NoAntrian			string	`gorm:"column:no_antrian"`
	JenisResep			string	`gorm:"column:jenis_resep"`
	Calltime			string	`gorm:"column:calltime"`
	Status				string	`gorm:"column:status"`
	ValidasiSend		string	`gorm:"column:validasi_send"`
	Json				string	`gorm:"column:json"`
}

func (IoReferensiFarmasi) TableName() string {
	return "io_referensi_farmasi"
}