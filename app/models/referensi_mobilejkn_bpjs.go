package models

type ReferensiMobilejknBpjs struct {
	Nobooking				string		`gorm:"column:nobooking"`
	NoRawat					string		`gorm:"column:no_rawat"`
	NomorKartu				string		`gorm:"column:nomorkartu"`
	Nik						string		`gorm:"column:nik"`
	NoHp					string		`gorm:"column:nohp"`
	KodePoli				string		`gorm:"column:kodepoli"`
	PasienBaru				string		`gorm:"column:pasienbaru"`
	Norm					string		`gorm:"column:norm"`
	TanggalPeriksa			string		`gorm:"column:tanggalperiksa"`
	KodeDokter				string		`gorm:"column:kodedokter"`
	JamPraktek				string		`gorm:"column:jampraktek"`
	JenisKunjungan			string		`gorm:"column:jeniskunjungan"`
	NomorReferensi			string		`gorm:"column:nomorreferensi"`
	NomorAntrean			string		`gorm:"column:nomorantrean"`
	AngkaAntrean			string		`gorm:"column:angkaantrean"`
	EstimasiDilayani		string		`gorm:"column:estimasidilayani"`
	SisaKuotaJkn			int			`gorm:"column:sisakuotajkn"`
	KuotaJkn				int			`gorm:"column:kuotajkn"`
	SisaKuotaNonJkn			int			`gorm:"column:sisakuotanonjkn"`
	KuotaNonJkn				int			`gorm:"column:kuotanonjkn"`
	Status					string		`gorm:"column:status"`
	Validasi				string		`gorm:"column:validasi"`
	StatusKirim				string		`gorm:"column:statuskirim"`
}

func (ReferensiMobilejknBpjs) TableName() string {
	return "referensi_mobilejkn_bpjs"
}