package models

import "time"

type ResepObat struct {
	NoResep					string		`gorm:"column:no_resep;primaryKey"`
	TglPerawatan			string		`gorm:"column:tgl_perawatan"`
	Jam						time.Time	`gorm:"column:jam"`
	NoRawat					string		`gorm:"column:no_rawat"`
	KdDokter				string		`gorm:"column:kd_dokter"`
	TglPeresepan			string		`gorm:"column:tgl_peresepan"`
	JamPeresepan			string		`gorm:"column:jam_peresepan"`
	Status					string		`gorm:"column:status"`
	TglPenyerahan			string		`gorm:"column:tgl_penyerahan"`
	JamPenyerahan			string		`gorm:"column:jam_penyerahan"`
}

func (ResepObat) TableName() string {
	return "resep_obat"
}