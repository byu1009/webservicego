package models

type IoSatuSehatImagingStudy struct {
	Acsn					string	`gorm:"column:acsn;primaryKey"`
	NoOrder					string	`gorm:"column:noorder"`
	KdJenisPrw				string	`gorm:"column:kd_jenis_prw"`
	IdImagingStudy			string	`gorm:"column:id_imaging_study"`
	StudyId					string	`gorm:"column:study_id"`
	StudyUuid				string	`gorm:"column:study_uuid"`
	PatientId				string	`gorm:"column:patient_id"`
}

func (IoSatuSehatImagingStudy) TableName() string {
	return "io_satu_sehat_imaging_study"
}