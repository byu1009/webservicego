package models

type ResepDokterRacikan struct {
	NoResep				string	`gorm:"column:no_resep"`
	NoRacik				string	`gorm:"column:no_racik"`
	NamaRacik			string	`gorm:"column:nama_racik"`
	KdRacik				string	`gorm:"column:kd_racik"`
	JmlDr				string	`gorm:"column:jml_dr"`
	AturanPakai			string	`gorm:"column:aturan_pakai"`
	Keterangan			string	`gorm:"column:keterangan"`
}

func (ResepDokterRacikan) TableName()	string {
	return "resep_dokter_racikan"
}