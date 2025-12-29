package models

type KodeSatuan struct {
	KodeSat		string	`gorm:"column:kode_sat"`
	Satuan		string	`gorm:"column:satuan"`
}

func (KodeSatuan) TableName() string {
	return "kode_satuan"
}