package models

import "time"

type RawatInapPr struct {
	NoRawat      string    `gorm:"column:no_rawat;primaryKey"`
	KdJenisPrw   string    `gorm:"column:kd_jenis_prw;primaryKey"`
	Nip          string    `gorm:"column:nip;primaryKey"`
	TglPerawatan time.Time `gorm:"column:tgl_perawatan;type:date;primaryKey"`
	JamRawat     string    `gorm:"column:jam_rawat;type:time;primaryKey"`
	StatusBayar  string    `gorm:"column:status_bayar"`
	BiayaRawat   float64   `gorm:"column:biaya_rawat"`
}

// TableName memastikan nama tabel sesuai database Khanza
func (RawatInapPr) TableName() string {
	return "rawat_inap_pr"
}