package models

type Jenis struct {
	KdJns		string	`gorm:"column:kdjns;primaryKey"`
	Nama		string	`gorm:"column:nama"`
}

func (Jenis) TableName() string {
	return "jenis"
}