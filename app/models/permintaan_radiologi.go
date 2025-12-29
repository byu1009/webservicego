package models

import "time"

type PermintaanRadiologi struct {
	NoOrder            string    `gorm:"column:noorder;primaryKey"`
	NoRawat            string    `gorm:"column:no_rawat"`
	TglPermintaan      time.Time `gorm:"column:tgl_permintaan;type:date"`
	JamPermintaan      string    `gorm:"column:jam_permintaan;type:time"`
	TglSampel          time.Time `gorm:"column:tgl_sampel;type:date"`
	JamSampel          string    `gorm:"column:jam_sampel;type:time"`
	TglHasil           time.Time `gorm:"column:tgl_hasil;type:date"`
	JamHasil           string    `gorm:"column:jam_hasil;type:time"`
	DokterPerujuk      string    `gorm:"column:dokter_perujuk"`
	Status             string    `gorm:"column:status"`
	InformasiTambahan  string    `gorm:"column:informasi_tambahan"`
	DiagnosaKlinis     string    `gorm:"column:diagnosa_klinis"`
}

// TableName memastikan nama tabel sesuai database
func (PermintaanRadiologi) TableName() string {
	return "permintaan_radiologi"
}
