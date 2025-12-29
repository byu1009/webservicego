package models

import "time"

// ===================================================
// PENILAIAN MEDIS RAWAT JALAN PENYAKIT DALAM
// ===================================================
type PenilaianMedisRalanPenyakitDalam struct {
	NoRawat                string    `gorm:"column:no_rawat;primaryKey"`
	Tanggal                time.Time `gorm:"column:tanggal"`
	KdDokter               string    `gorm:"column:kd_dokter"`
	Anamnesis              string    `gorm:"column:anamnesis"`
	Hubungan               string    `gorm:"column:hubungan"`
	KeluhanUtama           string    `gorm:"column:keluhan_utama"`
	Rps                    string    `gorm:"column:rps"`
	Rpd                    string    `gorm:"column:rpd"`
	Rpo                    string    `gorm:"column:rpo"`
	Alergi                 string    `gorm:"column:alergi"`
	Kondisi                string    `gorm:"column:kondisi"`
	Status                 string    `gorm:"column:status"`
	Td                     string    `gorm:"column:td"`
	Nadi                   string    `gorm:"column:nadi"`
	Suhu                   string    `gorm:"column:suhu"`
	Rr                     string    `gorm:"column:rr"`
	Bb                     string    `gorm:"column:bb"`
	Nyeri                  string    `gorm:"column:nyeri"`
	Gcs                    string    `gorm:"column:gcs"`
	Kepala                 string    `gorm:"column:kepala"`
	KeteranganKepala        string    `gorm:"column:keterangan_kepala"`
	Thoraks                string    `gorm:"column:thoraks"`
	KeteranganThorak        string    `gorm:"column:keterangan_thorak"`
	Abdomen                string    `gorm:"column:abdomen"`
	KeteranganAbdomen       string    `gorm:"column:keterangan_abdomen"`
	Ekstremitas             string    `gorm:"column:ekstremitas"`
	KeteranganEkstremitas   string    `gorm:"column:keterangan_ekstremitas"`
	Lainnya                string    `gorm:"column:lainnya"`
	Lab                    string    `gorm:"column:lab"`
	Rad                    string    `gorm:"column:rad"`
	PenunjangLain           string    `gorm:"column:penunjanglain"`
	Diagnosis               string    `gorm:"column:diagnosis"`
	Diagnosis2              string    `gorm:"column:diagnosis2"`
	Permasalahan            string    `gorm:"column:permasalahan"`
	Terapi                  string    `gorm:"column:terapi"`
	Tindakan                string    `gorm:"column:tindakan"`
	Edukasi                 string    `gorm:"column:edukasi"`
}

// Nama tabel eksplisit
func (PenilaianMedisRalanPenyakitDalam) TableName() string {
	return "penilaian_medis_ralan_penyakit_dalam"
}
