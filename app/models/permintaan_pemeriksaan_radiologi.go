package models

type PermintaanPemeriksaanRadiologi struct {
	NoOrder     string `gorm:"column:noorder;primaryKey"`
	KdJenisPrw  string `gorm:"column:kd_jenis_prw;primaryKey"`
	Status      string `gorm:"column:status"`
}

// Nama tabel eksplisit
func (PermintaanPemeriksaanRadiologi) TableName() string {
	return "permintaan_pemeriksaan_radiologi"
}