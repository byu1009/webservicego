package models

type JnsPerawatanRadiologi struct {
	KdJenisPrw					string	`gorm:"column:kd_jenis_perawatan"`
	NmPerawatan					string	`gorm:"column:nm_perawatan"`
	Bhp							string	`gorm:"column:bhp"`
	TarifPerujuk				string	`gorm:"column:tarif_perujuk"`
	TarifTindakanDokter			string	`gorm:"column:tarif_tindakan_dokter"`
	TarifTindakanPetugas		string	`gorm:"column:tarif_tindakan_petugas"`
	Kso							string	`gorm:"column:kso"`
	Menejemen					string	`gorm:"column:menejemen"`
	TotalByr					string	`gorm:"column:total_byr"`
	KdPj						string	`gorm:"column:kd_pj"`
	Status						string	`gorm:"column:status"`
	Kelas						string	`gorm:"column:kelas"`
}

func (JnsPerawatanRadiologi) TableName() string {
	return "jns_perawatan_radiologi"
}