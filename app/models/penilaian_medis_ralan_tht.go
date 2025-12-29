package models

import "time"

// ===================================================
// PENILAIAN MEDIS RAWAT JALAN THT
// ===================================================
type PenilaianMedisRalanTht struct {
	NoRawat           string    `gorm:"column:no_rawat;primaryKey"`
	Tanggal           time.Time `gorm:"column:tanggal"`
	KdDokter          string    `gorm:"column:kd_dokter"`
	Anamnesis         string    `gorm:"column:anamnesis"`
	Hubungan          string    `gorm:"column:hubungan"`
	KeluhanUtama      string    `gorm:"column:keluhan_utama"`
	Rps               string    `gorm:"column:rps"`
	Rpd               string    `gorm:"column:rpd"`
	Rpo               string    `gorm:"column:rpo"`
	Alergi            string    `gorm:"column:alergi"`
	Td                string    `gorm:"column:td"`
	Nadi              string    `gorm:"column:nadi"`
	Rr                string    `gorm:"column:rr"`
	Suhu              string    `gorm:"column:suhu"`
	Bb                string    `gorm:"column:bb"`
	Tb                string    `gorm:"column:tb"`
	Nyeri             string    `gorm:"column:nyeri"`
	StatusNutrisi     string    `gorm:"column:status_nutrisi"`
	Kondisi           string    `gorm:"column:kondisi"`
	KetLokalis        string    `gorm:"column:ket_lokalis"`
	Lab               string    `gorm:"column:lab"`
	Rad               string    `gorm:"column:rad"`
	TesPendengaran    string    `gorm:"column:tes_pendengaran"`
	Penunjang         string    `gorm:"column:penunjang"`
	Diagnosis         string    `gorm:"column:diagnosis"`
	DiagnosisBanding  string    `gorm:"column:diagnosisbanding"`
	Permasalahan      string    `gorm:"column:permasalahan"`
	Terapi            string    `gorm:"column:terapi"`
	Tindakan          string    `gorm:"column:tindakan"`
	Tatalaksana       string    `gorm:"column:tatalaksana"`
	Edukasi           string    `gorm:"column:edukasi"`
}

// Nama tabel eksplisit
func (PenilaianMedisRalanTht) TableName() string {
	return "penilaian_medis_ralan_tht"
}
