package models

type SatuSehatMappingLokasiRalan struct {
	KdPoli                string  `gorm:"column:kd_poli;primaryKey"`
	IdOrganisasiSatuSehat string  `gorm:"column:id_organisasi_satusehat"`
	IdLokasiSatuSehat     string  `gorm:"column:id_lokasi_satusehat"`
	Longitude             float64 `gorm:"column:longitude"`
	Latitude              float64 `gorm:"column:latitude"`
	Altittude             float64 `gorm:"column:altittude"`
}

// TableName memastikan nama tabel sesuai database
func (SatuSehatMappingLokasiRalan) TableName() string {
	return "satu_sehat_mapping_lokasi_ralan"
}