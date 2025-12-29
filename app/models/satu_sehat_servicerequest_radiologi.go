package models

type SatuSehatServiceRequestRadiologi struct {
	NoOrder          string `gorm:"column:noorder;primaryKey"`
	KdJenisPrw       string `gorm:"column:kd_jenis_prw"`
	IDServiceRequest string `gorm:"column:id_servicerequest"`
}

// TableName memastikan nama tabel sesuai database
func (SatuSehatServiceRequestRadiologi) TableName() string {
	return "satu_sehat_servicerequest_radiologi"
}