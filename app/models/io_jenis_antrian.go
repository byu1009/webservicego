package models

type IoJenisAntrian struct {
	Id				int		`gorm:"column:id"`
	JenisAntrian	string	`gorm:"column:jenis_antrian"`
	Prefix			string	`gorm:"column:prefix"`
}

func (IoJenisAntrian) TableName() string {
	return "io_jenis_antrian"
}