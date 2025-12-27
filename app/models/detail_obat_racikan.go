package models

type DetailObatRacikan struct {
	TglPerawatan	string	`gorm:"column:tgl_perawatan"`
	Jam				string	`gorm:"column:jam"`
	NoRawat			string	`gorm:"column:no_rawat"`
	NoRacik			string	`gorm:"column:no_racik"`
	KodeBrng		string	`gorm:"column:kode_brng"`
}

func (DetailObatRacikan) TableName() string {
	return "detail_obat_racikan"
}