package models

type DiagnosaPasien struct {
	NoRawat				string	`gorm:"column:no_rawat;primaryKey"`
	KdPenyakit			string	`gorm:"column:kd_penyakit"`
	Status				string	`gorm:"column:status"`
	Prioritas			string	`gorm:"column:prioritas"`
	StatusPenyakit		string	`gorm:"column:status_penyakit"`
}

func (DiagnosaPasien) TableName() string {
	return "diagnosa_pasien"
}