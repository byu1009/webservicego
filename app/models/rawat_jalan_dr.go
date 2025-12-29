package models

import "time"

type RawatJalanDr struct {
	NoRawat         string    `gorm:"column:no_rawat;primaryKey"`
	KdJenisPrw      string    `gorm:"column:kd_jenis_prw;primaryKey"`
	KdDokter        string    `gorm:"column:kd_dokter;primaryKey"`
	TglPerawatan    time.Time `gorm:"column:tgl_perawatan;type:date;primaryKey"`
	JamRawat        string    `gorm:"column:jam_rawat;type:time;primaryKey"`
	Material        float64   `gorm:"column:material"`
	Bhp             float64   `gorm:"column:bhp"`
	TarifTindakanDr float64   `gorm:"column:tarif_tindakandr"`
	Kso             float64   `gorm:"column:kso"`
	Menejemen       float64   `gorm:"column:menejemen"`
	BiayaRawat      float64   `gorm:"column:biaya_rawat"`
	SttsBayar       string    `gorm:"column:stts_bayar"`
}

// TableName memastikan nama tabel sesuai database
func (RawatJalanDr) TableName() string {
	return "rawat_jl_dr"
}