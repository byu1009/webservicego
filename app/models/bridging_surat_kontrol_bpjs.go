package models

type BridgingSuratKontrolBpjs struct {
	NoSurat      string `gorm:"column:no_surat;primaryKey;type:varchar(50)" json:"no_surat"`
	NoSep        string `gorm:"column:no_sep;type:varchar(50)" json:"no_sep"`
	TglSurat     string `gorm:"column:tgl_surat;type:date" json:"tgl_surat"`
	TglRencana   string `gorm:"column:tgl_rencana;type:date" json:"tgl_rencana"`
	KdDokterBpjs string `gorm:"column:kd_dokter_bpjs;type:varchar(50)" json:"kd_dokter_bpjs"`
	NmDokterBpjs string `gorm:"column:nm_dokter_bpjs;type:varchar(255)" json:"nm_dokter_bpjs"`
	KdPoliBpjs   string `gorm:"column:kd_poli_bpjs;type:varchar(50)" json:"kd_poli_bpjs"`
	NmPoliBpjs   string `gorm:"column:nm_poli_bpjs;type:varchar(255)" json:"nm_poli_bpjs"`

	// RELASI hasOne
	SepAsal *BridgingSep `gorm:"foreignKey:NoSep;references:NoSep" json:"sep_asal,omitempty"`
}

func (BridgingSuratKontrolBpjs) TableName() string {
	return "bridging_surat_kontrol_bpjs"
}
