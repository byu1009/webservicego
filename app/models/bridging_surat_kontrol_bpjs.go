package models

import "time"

type BridgingSuratKontrolBpjs struct {
	NoSurat      string `gorm:"column:no_surat;primaryKey"`
	NoSep        string `gorm:"column:no_sep"`
	TglSurat     time.Time `gorm:"column:tgl_surat"`
	TglRencana   time.Time `gorm:"column:tgl_rencana"`
	KdDokterBpjs string
	NmDokterBpjs string
	KdPoliBpjs   string
	NmPoliBpjs   string

	// RELASI hasOne
	// SepAsal *BridgingSep `gorm:"foreignKey:NoSep;references:NoSep" json:"sep_asal,omitempty"`
	SepAsal *BridgingSep `gorm:"foreignKey:NoSkdp;references:NoSurat"`
}

func (BridgingSuratKontrolBpjs) TableName() string {
	return "bridging_surat_kontrol_bpjs"
}
