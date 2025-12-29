package models

type JnsPerawatanInap struct {
	KdJenisPrw				string	`gorm:"column:kd_jenis_prw"`
	NmPerawatan				string	`gorm:"column:nm_perawatan"`
	KdKategori				string	`gorm:"column:kd_kategori"`
	Bhp						string	`gorm:"column:bhp"`
	TarifTindakanDr			string	`gorm:"tarif_tindakandr"`
	TarifTindakanPr			string	`gorm:"tarif_tindakanpr"`
	Kso						string	`gorm:"kso"`
	Menejemen				string	`gorm:"menejemen"`
	KdPj					string	`gorm:"kd_pj"`
	KdBangsal				string	`gorm:"kd_bangsal"`
	Kelas					string	`gorm:"kelas"`
	Status					string	`gorm:"status"`
	TotalByrDr				string	`gorm:"total_byrdr"`
	TotalByrPr				string	`gorm:"total_byrpr"`
	TotalByrDrPr			string	`gorm:"total_byrDrPr"`
}

func (JnsPerawatanInap) TableName() string {
	return "jns_perawatan_inap"
}
