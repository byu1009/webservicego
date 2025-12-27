package models

type IOAntrian struct {
	NoReferensi			string	`gorm:"column:no_referensi;primaryKey"`
	NoAntrian			string	`gorm:"column:no_antrian"`
	StatusPanggil		string	`gorm:"column:status_panggil"`
	StatusAntrian		string	`gorm:"column:status_antrian"`
	Calltime			int		`gorm:"column:calltime"`
	StatusPasien		string	`gorm:"column:status_pasien"`
	Order				int		`gorm:"column:order"`
}

func (IOAntrian) TableName() string {
	return "io_antrian"
}