package models

type SatuSehatPatient struct {
	NoKtp     string `gorm:"column:no_ktp;primaryKey"`
	PatientID string `gorm:"column:patient_id"`
}

// TableName memastikan nama tabel sesuai database
func (SatuSehatPatient) TableName() string {
	return "io_satu_sehat_patient"
}