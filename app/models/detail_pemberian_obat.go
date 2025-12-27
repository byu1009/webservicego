package models

type DetailPemberianObat struct {
	TglPerawatan string  `gorm:"column:tgl_perawatan;type:date" json:"tgl_perawatan"`
	Jam          string  `gorm:"column:jam;type:time" json:"jam"`
	NoRawat      string  `gorm:"column:no_rawat;type:varchar(50)" json:"no_rawat"`
	KodeBrng     string  `gorm:"column:kode_brng;type:varchar(50)" json:"kode_brng"`
	NoBatch      string  `gorm:"column:no_batch;type:varchar(50)" json:"no_batch"`
	NoFaktur     string  `gorm:"column:no_faktur;type:varchar(50)" json:"no_faktur"`

	HBeli       float64 `gorm:"column:h_beli" json:"h_beli"`
	BiayaObat   float64 `gorm:"column:biaya_obat" json:"biaya_obat"`
	Jml         float64 `gorm:"column:jml" json:"jml"`
	Embalase    float64 `gorm:"column:embalase" json:"embalase"`
	Tuslah      float64 `gorm:"column:tuslah" json:"tuslah"`
	Total       float64 `gorm:"column:total" json:"total"`
	Status      string  `gorm:"column:status" json:"status"`
	KdBangsal   string  `gorm:"column:kd_bangsal" json:"kd_bangsal"`

	// RELATIONS
	DataBarang *DataBarang `gorm:"foreignKey:KodeBrng;references:KodeBrng" json:"databarang,omitempty"`
	Bangsal    *Bangsal    `gorm:"foreignKey:KdBangsal;references:KdBangsal" json:"bangsal,omitempty"`
}

func (DetailPemberianObat) TableName() string {
	return "detail_pemberian_obat"
}
