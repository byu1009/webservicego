package models

type BridgingSep struct {
	NoSep					string	`gorm:"column:no_sep;primaryKey"`
	NoRawat					string	`gorm:"column:no_rawat"`
	TglSep					string	`gorm:"column:tgl_sep"`
	TglRujukan				string	`gorm:"column:tglrujukan"`
	NoRujukan				string	`gorm:"column:no_rujukan"`
	KdPpkRujukan			string	`gorm:"column:kdppkrujukan"`
	NmPpkRujukan			string	`gorm:"column:nmppkrujukan"`
	KdPpkPelayanan			string	`gorm:"column:kdppkpelayanan"`
	NmPpkPelayanan			string	`gorm:"column:nmppkpelayanan"`
	JnsPelayanan			string	`gorm:"column:jnspelayanan"`
	Catatan					string	`gorm:"column:catatan"`
	Diagawal				string	`gorm:"column:diagawal"`
	NmDiagnosaAwal			string	`gorm:"column:nmdiagnosaawal"`
	KdPoliTujuan			string	`gorm:"column:kdpolitujuan"`
	NmPoliTujuan			string	`gorm:"column:nmpolitujuan"`
	KlsRawat				string	`gorm:"column:klsrawat"`
	KlsNaik					string	`gorm:"column:klsnaik"`
	Pembiayaan				string	`gorm:"column:pembiayaan"`
	PjNaikKelas				string	`gorm:"column:pjnaikkelas"`
	LakaLantas				string	`gorm:"column:lakalantas"`
	User					string	`gorm:"column:user"`
	Nomr					string	`gorm:"column:nomr"`
	NamaPasien				string	`gorm:"column:nama_pasien"`
	TanggalLahir			string	`gorm:"column:tanggal_lahir"`
	Peserta					string	`gorm:"column:peserta"`
	Jkel					string	`gorm:"column:jkel"`
	NoKartu					string	`gorm:"column:no_kartu"`
	TglPulang				string	`gorm:"column:tglpulang"`
	AsalRujukan				string	`gorm:"column:asal_rujukan"`
	Eksekutif				string	`gorm:"column:eksekutif"`
	Cob						string	`gorm:"column:cob"`
	NoTelep					string	`gorm:"column:notelep"`
	Katarak					string	`gorm:"column:katarak"`
	TglKkl					string	`gorm:"column:tglkkl"`
	KeteranganKkl			string	`gorm:"column:keterangankkl"`
	Suplesi					string	`gorm:"column:suplesi"`
	NoSepSuplesi			string	`gorm:"column:no_sep_suplesi"`
	KdProp					string	`gorm:"column:kdprop"`
	NmProp					string	`gorm:"column:nmprop"`
	KdKab					string	`gorm:"column:kdkab"`
	NmKab					string	`gorm:"column:nmkab"`
	KdKec					string	`gorm:"column:kdkec"`
	NmKec					string	`gorm:"column:nmkec"`
	NoSkdp					string	`gorm:"column:noskdp"`
	KdDpjp					string	`gorm:"column:kddpjp"`
	NmDpjp					string	`gorm:"column:nmdpjp"`
	TujuanKunjungan			string	`gorm:"column:tujuankunjungan"`
	FalgProsedur			string	`gorm:"column:flasprosedur"`
	Penunjang				string	`gorm:"column:penunjang"`
	AsesmenPelayanan		string	`gorm:"column:asesmenpelayanan"`
	KdDpjpPelayanan			string	`gorm:"column:kddpjppelayanan"`
	NmDpjpPelayanan			string	`gorm:"column:nmdpjppelayanan"`
}

func (BridgingSep) TableName() string {
	return "bridging_sep"
}
