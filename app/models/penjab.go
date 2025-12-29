package models

type Penjab struct {
	KdPj					string	`gorm:"column:kd_pj;primaryKey"`
	PngJawab				string	`gorm:"column:png_jawab"`
	NamaPerusahaan			string	`gorm:"column:nama_perusahaan"`
	AlamatAsuransi			string	`gorm:"column:alamat_asuransi"`
	NoTelp					string	`gorm:"column:no_telp"`
	Attn					string	`gorm:"column:attn"`
	Status					string	`gorm:"column:status"`
}

func (Penjab) TableName() string {
	return "penjab"
}