package models

type Jadwal struct {
	KdDokter			string	`gorm:"column:kd_dokter"`
	HariKerja			string	`gorm:"column:hari_kerja"`
	JamMulai			string	`gorm:"column:jam_mulai"`
	JamSelesai			string	`gorm:"column:jam_selesai"`
	KdPoli				string	`gorm:"column:kd_poli"`
	Kuota				string	`gorm:"column:kuota"`
}

func (Jadwal) TableName() string {
	return "jadwal"
}