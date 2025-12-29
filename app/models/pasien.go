package models

type Pasien struct {
	NoRkmMedis       string `gorm:"column:no_rkm_medis;primaryKey"`
	NmPasien         string `gorm:"column:nm_pasien"`
	NoKtp            string `gorm:"column:no_ktp"`
	Jk               string `gorm:"column:jk"`
	TmpLahir         string `gorm:"column:tmp_lahir"`
	TglLahir         string `gorm:"column:tgl_lahir"`
	NmIbu            string `gorm:"column:nm_ibu"`
	Alamat           string `gorm:"column:alamat"`
	GolDarah         string `gorm:"column:gol_darah"`
	Pekerjaan        string `gorm:"column:pekerjaan"`
	SttsNikah        string `gorm:"column:stts_nikah"`
	Agama            string `gorm:"column:agama"`
	TglDaftar        string `gorm:"column:tgl_daftar"`
	NoTlp            string `gorm:"column:no_tlp"`
	Umur             string `gorm:"column:umur"`
	Pnd              string `gorm:"column:pnd"`
	Keluarga         string `gorm:"column:keluarga"`
	NamaKeluarga     string `gorm:"column:namakeluarga"`
	KdPj             string `gorm:"column:kd_pj"`
	NoPeserta        string `gorm:"column:no_peserta"`
	KdKel            string `gorm:"column:kd_kel"`
	KdKec            string `gorm:"column:kd_kec"`
	KdKab            string `gorm:"column:kd_kab"`
	PekerjaanPj      string `gorm:"column:pekerjaanpj"`
	AlamatPj         string `gorm:"column:alamatpj"`
	KelurahanPj      string `gorm:"column:kelurahanpj"`
	KecamatanPj      string `gorm:"column:kecamatanpj"`
	KabupatenPj      string `gorm:"column:kabupatenpj"`
	PerusahaanPasien string `gorm:"column:perusahaan_pasien"`
	SukuBangsa       string `gorm:"column:suku_bangsa"`
	BahasaPasien     string `gorm:"column:bahasa_pasien"`
	CacatFisik       string `gorm:"column:cacat_fisik"`
	Email            string `gorm:"column:email"`
	Nip              string `gorm:"column:nip"`
	KdProp           string `gorm:"column:kd_prop"`
	PropinsiPj       string `gorm:"column:propinsipj"`
}

func (Pasien) TableName() string {
	return "pasien"
}
