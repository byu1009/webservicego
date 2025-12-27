package models

type Dokter struct {
	KdDokter		string	`gorm:"column:kd_dokter;primaryKey"`
	NmDokter		string	`gorm:"column:nm_dokter"`
	JK				string	`gorm:"column:jk"`
	TmpLahir		string	`gorm:"column:tmp_lahir"`
	TglLahir		string	`gorm:"column:tgl_lahir"`
	GolDarah		string	`gorm:"column:gol_darah"`
	Agama			string	`gorm:"column:agama"`
	AlmtTgl			string	`gorm:"column:almt_tgl"`
	NoTelp			string	`gorm:"column:no_telp"`
	Email			string	`gorm:"column:email"`
	SttsNikah		string	`gorm:"column:stts_nikah"`
	KdSps			string	`gorm:"column:kd_sps"`
	Alumni			string	`gorm:"column:alumni"`
	NoIjnPraktek	string	`gorm:"column:no_ijn_praktek"`
	Status			string	`gorm:"column:status"`
}

func (Dokter) TableName() string {
	return "dokter"
}