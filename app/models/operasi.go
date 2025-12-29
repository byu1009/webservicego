package models

type Operasi struct {
	NoRawat						string	`gorm:"column:no_rawat"`
	TglOperasi					string	`gorm:"column:tgl_operasi"`
	KodePaket					string	`gorm:"column:kode_paket"`
	Kategori					string	`gorm:"column:kategori"`
	Operator1					string	`gorm:"column:operator1"`
	Operator2					string	`gorm:"column:operator2"`
	Operator3					string	`gorm:"column:operator3"`
	AsistenOperator1			string	`gorm:"column:asisten_operator1"`
	AsistenOperator2			string	`gorm:"column:asisten_operator2"`
	AsistenOperator3			string	`gorm:"column:asisten_operator3"`
	Instrumen					string	`gorm:"column:instrumen"`
	DokterAnak					string	`gorm:"column:dokter_anak"`
	PerawatResusitas			string	`gorm:"column:perawat_resusitas"`
	DokterAnestesi				string	`gorm:"column:dokter_anestesi"`
	DokterAnestesi2				string	`gorm:"column:dokter_anestesi2"`
	Bidan						string	`gorm:"column:bidan"`
	Bidan2						string	`gorm:"column:bidan2"`
	Bidan3						string	`gorm:"column:bidan3"`
	PerawatLuar					string	`gorm:"column:perawat_luar"`
	SewaOk						string	`gorm:"column:sewa_ok"`
	Alat						string	`gorm:"column:alat"`
	Akomodasi					string	`gorm:"column:akomodasi"`
	BagianRs					string	`gorm:"column:bagian_rs"`
	Omloop						string	`gorm:"column:omloop"`
	Omloop2						string	`gorm:"column:omloop2"`
	Omloop3						string	`gorm:"column:omloop3"`
	Omloop4						string	`gorm:"column:omloop4"`
	Omloop5						string	`gorm:"column:omloop5"`
	Sarpras						string	`gorm:"column:sarpras"`
	DokterPjAnak				string	`gorm:"column:dokter_pjanak"`
	DokterUmum					string	`gorm:"column:dokter_umum"`
	Status						string	`gorm:"column:status"`
	KdPj						string	`gorm:"column:kd_pj"`
	Kelas						string	`gorm:"column:kelas"`
}

func (Operasi) TableName() string {
	return "operasi"
}