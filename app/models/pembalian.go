package models

import "time"

type Pembelian struct {
	NoFaktur   string    `gorm:"column:no_faktur;primaryKey"`
	KdSuplier  string    `gorm:"column:kd_suplier"`
	Nip        string    `gorm:"column:nip"`
	TglBeli    time.Time `gorm:"column:tgl_beli"`
	Total1     float64   `gorm:"column:total1"`
	Potongan   float64   `gorm:"column:potongan"`
	Total2     float64   `gorm:"column:total2"`
	Ppn        float64   `gorm:"column:ppn"`
	Tagihan    float64   `gorm:"column:tagihan"`
	KdBangsal  string    `gorm:"column:kd_bangsal"`
	KdRek      string    `gorm:"column:kd_rek"`
}

func (Pembelian) TableName() string {
	return "pembelian"
}
