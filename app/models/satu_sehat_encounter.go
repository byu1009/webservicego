package models

type SatuSehatEncounter struct {
	NoRawat     string `gorm:"column:no_rawat;primaryKey"`
	IdEncounter string `gorm:"column:id_encounter"`
}

// TableName memastikan nama tabel sesuai database
func (SatuSehatEncounter) TableName() string {
	return "satu_sehat_encounter"
}