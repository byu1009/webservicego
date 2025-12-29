package models

type TemplateLaboratorium struct {
	IDTemplate        int64   `gorm:"column:id_template;primaryKey;autoIncrement"`
	KdJenisPrw        string  `gorm:"column:kd_jenis_prw"`
	Pemeriksaan       string  `gorm:"column:Pemeriksaan"`
	Satuan            string  `gorm:"column:satuan"`
	NilaiRujukanLD    string  `gorm:"column:nilai_rujukan_ld"`
	NilaiRujukanLA    string  `gorm:"column:nilai_rujukan_la"`
	NilaiRujukanPD    string  `gorm:"column:nilai_rujukan_pd"`
	NilaiRujukanPA    string  `gorm:"column:nilai_rujukan_pa"`
	BagianRS          float64 `gorm:"column:bagian_rs"`
	BHP               float64 `gorm:"column:bhp"`
	BagianPerujuk     float64 `gorm:"column:bagian_perujuk"`
	BagianDokter      float64 `gorm:"column:bagian_dokter"`
	BagianLaborat     float64 `gorm:"column:bagian_laborat"`
	KSO               float64 `gorm:"column:kso"`
	Menejemen         float64 `gorm:"column:menejemen"`
	BiayaItem         float64 `gorm:"column:biaya_item"`
	Urut              int     `gorm:"column:urut"`
}

// TableName override nama tabel
func (TemplateLaboratorium) TableName() string {
	return "template_laboratorium"
}