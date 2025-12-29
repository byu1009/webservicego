package models

type IoSatuSehatPractitioner struct {
	NoKtp			string	`gorm:"column:no_ktp;primaryKey"`
	practitioner_id	string	`gorm:"column:practitioner_id"`
	data			string	`gorm:"column:data"`
}

func (IoSatuSehatPractitioner) TableName() string {
	return "io_satu_sehat_practitioner"
}