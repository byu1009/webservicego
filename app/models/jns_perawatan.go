package models

type JnsPerawatan struct {
	KdJenisPrw				string	`gorm:"column:kd_jenis_prw"`
	NmPerawatan				string	`gorm:"column:nm_perawtan"`
	KdKategori				string	`gorm:"column:kd_kategori"`
	Bhp						string	`gorm:"column:bhp"`
	TarifTindakanDr			string	`gorm:"column:tarif_tindakandr"`
	TarifTindakanPr			string	`gorm:"column:tarif_tindakanpr"`
	Kso						string	`gorm:"column:kso"`
	Menejemen				string	`gorm:"column:menejemen"`
	KdPj					string	`gorm:"column:kd_pj"`
	KdPoli					string	`gorm:"column:kd_poli"`
	Status					string	`gorm:"column:status"`
	TotalByrDr				string	`gorm:"column:total_byrdr"`
	TotalByrPr				string	`gorm:"column:total_byrpr"`
	TotalByrDrPr			string	`gorm:"column:total_byrdrpr"`
}

func (JnsPerawatan) TableName() string {
	return "jns_perawatan"
}