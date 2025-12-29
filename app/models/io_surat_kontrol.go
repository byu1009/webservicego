package models

type IoSuratKontrol struct {
	NoSurat					string	`gorm:"column:no_surat"`
	NoRkmMedis				string	`gorm:"column:no_rkm_medis"`
	Alasan					string	`gorm:"column:alasan"`
	TanggalKontrol			string	`gorm:"column:tanggal_kontrol"`
	KdDokter				string	`gorm:"column:kd_dokter"`
	KdPoli					string	`gorm:"column:kd_poli"`
	Status					string	`gorm:"column:status"`
}

func (IoSuratKontrol) TableName() string {
	return "io_surat_kontrol"
}