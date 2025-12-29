package models

type SatuSehatMappingRadiologi struct {
	KdJenisPrw     string `gorm:"column:kd_jenis_prw;primaryKey"`
	Code           string `gorm:"column:code"`
	System         string `gorm:"column:system"`
	Display        string `gorm:"column:display"`
	SampelCode     string `gorm:"column:sampel_code"`
	SampelSystem   string `gorm:"column:sampel_system"`
	SampelDisplay  string `gorm:"column:sampel_display"`
}

// TableName memastikan nama tabel sesuai database
func (SatuSehatMappingRadiologi) TableName() string {
	return "satu_sehat_mapping_radiologi"
}