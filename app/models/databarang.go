package models

type DataBarang struct {
	KodeBrng			string	`gorm:"column:kode_brng"`
	NamaBrng			string	`gorm:"column:nama_brng"`
	KodeSatbesar		string	`gorm:"column:kode_satbesar"`
	KodeSat				string	`gorm:"column:kode_sat"`
	KodeKategori		string	`gorm:"column:kode_kategori"`
	KodeIndustri		string	`gorm:"column:kode_industri"`
	KodeGolongan		string	`gorm:"column:kode_golongan"`
	KdJns				string	`gorm:"column:kdjns"`
	LetakBarang			string	`gorm:"column:letak_barang"`
	Dasar				string	`gorm:"column:dasar"`
	HBeli				string	`gorm:"column:hbeli"`
	Ralan				string	`gorm:"column:ralan"`
	Kelas1				string	`gorm:"column:kelas1"`
	Kelas2				string	`gorm:"column:kelas2"`
	Kelas3				string	`gorm:"column:kelas3"`
	Utama				string	`gorm:"column:utama"`
	Vip					string	`gorm:"column:vip"`
	Vvip				string	`gorm:"column:vvip"`
	BeliLuar			string	`gorm:"column:beliluar"`
	JualBebas			string	`gorm:"column:jualbebas"`
	Karyawan			string	`gorm:"column:karyawan"`
}

func (DataBarang) TableName() string {
	return "databarang"
}