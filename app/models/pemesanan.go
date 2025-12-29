package models

import "time"

type Pemesanan struct {
	NoFaktur  string    `gorm:"column:no_faktur;primaryKey"`
	KdSuplier string    `gorm:"column:kd_suplier"`
	Nip       string    `gorm:"column:nip"`
	TglPesan  time.Time `gorm:"column:tgl_pesan"`
	TglFaktur time.Time `gorm:"column:tgl_faktur"`
	TglTempo  time.Time `gorm:"column:tgl_tempo"`
	Total1    float64   `gorm:"column:total1"`
	Potongan  float64   `gorm:"column:potongan"`
	Total2    float64   `gorm:"column:total2"`
	Ppn       float64   `gorm:"column:ppn"`
	Tagihan   float64   `gorm:"column:tagihan"`
	KdBangsal string    `gorm:"column:kd_bangsal"`
	Materai   float64   `gorm:"column:materai"`
}

func (Pemesanan) TableName() string {
	return "pemesanan"
}
