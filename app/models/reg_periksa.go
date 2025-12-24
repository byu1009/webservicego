package models

type RegPeriksa struct {
	NoReg			*string		`gorm:"column:no_reg"`
	NoRawat			string		`gorm:"column:no_rawat;primaryKey"`
	TglRegistrasi	*string		`gorm:"column:tgl_registrasi"`
	JamReg			*string		`gorm:"column:jam_reg"`
	KodeDokter		*string		`gorm:"column:kd_dokter"`
	Norm			*string		`gorm:"column:no_rkm_medis"`
	KodePoli		*string		`gorm:"column:kd_poli"`
	PJawab			*string		`gorm:"column:p_jawab"`
	Almtpj			*string		`gorm:"column:almt_pj"`
	Hubunganpj		*string		`gorm:"column:hubunganpj"`
	BiayaReg	 	*int			`gorm:"column:biaya_reg"`
	Stts			*string		`gorm:"column:stts"`
	SttsDaftar		string		`gorm:"column:stts_daftar"`
	StatusLanjut	string		`gorm:"column:status_lanjut"`
	Kdpj			string		`gorm:"column:kd_pj"`
	Umurdaftar		*int			`gorm:"column:umurdaftar"`
	Sttsumur		*string		`gorm:"column:sttsumur"`
	Statusbayar		string		`gorm:"column:status_bayar"`
	StatusPoli		string		`gorm:"column:status_poli"`
}

func (RegPeriksa) TableName() string {
	return "reg_periksa"
}