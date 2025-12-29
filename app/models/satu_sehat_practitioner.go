package models

type SatuSehatPractitioner struct {
	NoKtp          string `gorm:"column:no_ktp;primaryKey"`
	PractitionerID string `gorm:"column:practitioner_id"`
}

// TableName memastikan nama tabel sesuai database
func (SatuSehatPractitioner) TableName() string {
	return "io_satu_sehat_practitioner"
}