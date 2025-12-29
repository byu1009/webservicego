package models

type IoSatuSehatPatient struct {
	NoKtp		string	`gorm:"column:no_ktp"`
	PatientId	string	`gorm:"column:patient_id"`
}

func (IoSatuSehatPatient) TableName() string {
	return "io_satu_sehat_patient"
}