package models

import "time"

// ===================================================
// PENILAIAN MEDIS RAWAT JALAN OBGYN / KANDUNGAN
// ===================================================
type PenilaianMedisRalanObgyn struct {
	NoRawat      string    `gorm:"column:no_rawat;primaryKey"`
	Tanggal      time.Time `gorm:"column:tanggal"`
	KdDokter     string    `gorm:"column:kd_dokter"`
	Anamnesis    string    `gorm:"column:anamnesis"`
	Hubungan     string    `gorm:"column:hubungan"`
	KeluhanUtama string    `gorm:"column:keluhan_utama"`
	Rps          string    `gorm:"column:rps"`
	Rpd          string    `gorm:"column:rpd"`
	Rpk          string    `gorm:"column:rpk"`
	Rpo          string    `gorm:"column:rpo"`
	Alergi       string    `gorm:"column:alergi"`
	Keadaan      string    `gorm:"column:keadaan"`
	Gcs          string    `gorm:"column:gcs"`
	Kesadaran    string    `gorm:"column:kesadaran"`
	Td           string    `gorm:"column:td"`
	Nadi         string    `gorm:"column:nadi"`
	Rr           string    `gorm:"column:rr"`
	Suhu         string    `gorm:"column:suhu"`
	Spo          string    `gorm:"column:spo"`
	Bb           string    `gorm:"column:bb"`
	Tb           string    `gorm:"column:tb"`
	Kepala       string    `gorm:"column:kepala"`
	Mata         string    `gorm:"column:mata"`
	Gigi         string    `gorm:"column:gigi"`
	Tht          string    `gorm:"column:tht"`
	Thoraks      string    `gorm:"column:thoraks"`
	Abdomen      string    `gorm:"column:abdomen"`
	Genital      string    `gorm:"column:genital"`
	Ekstremitas  string    `gorm:"column:ekstremitas"`
	Kulit        string    `gorm:"column:kulit"`
	KetFisik     string    `gorm:"column:ket_fisik"`
	Tfu          string    `gorm:"column:tfu"`
	Tbj          string    `gorm:"column:tbj"`
	His          string    `gorm:"column:his"`
	Kontraksi    string    `gorm:"column:kontraksi"`
	Djj          string    `gorm:"column:djj"`
	Inspeksi     string    `gorm:"column:inspeksi"`
	Inspekulo    string    `gorm:"column:inspekulo"`
	Vt           string    `gorm:"column:vt"`
	Rt           string    `gorm:"column:rt"`
	Ultra        string    `gorm:"column:ultra"`
	Kardio       string    `gorm:"column:kardio"`
	Lab          string    `gorm:"column:lab"`
	Diagnosis    string    `gorm:"column:diagnosis"`
	Tata         string    `gorm:"column:tata"`
	Konsul       string    `gorm:"column:konsul"`
}

// Nama tabel eksplisit
func (PenilaianMedisRalanObgyn) TableName() string {
	return "penilaian_medis_ralan_kandungan"
}
