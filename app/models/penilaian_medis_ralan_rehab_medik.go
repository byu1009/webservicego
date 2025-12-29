package models

import "time"

// ===================================================
// PENILAIAN MEDIS RAWAT JALAN REHABILITASI MEDIK
// ===================================================
type PenilaianMedisRalanRehabMedik struct {
	NoRawat               string    `gorm:"column:no_rawat;primaryKey"`
	Tanggal               time.Time `gorm:"column:tanggal"`
	KdDokter              string    `gorm:"column:kd_dokter"`
	Anamnesis             string    `gorm:"column:anamnesis"`
	Hubungan              string    `gorm:"column:hubungan"`
	KeluhanUtama          string    `gorm:"column:keluhan_utama"`
	Rps                   string    `gorm:"column:rps"`
	Rpd                   string    `gorm:"column:rpd"`
	Alergi                string    `gorm:"column:alergi"`
	Kesadaran             string    `gorm:"column:kesadaran"`
	Nyeri                 string    `gorm:"column:nyeri"`
	SkalaNyeri            string    `gorm:"column:skala_nyeri"`
	Td                    string    `gorm:"column:td"`
	Nadi                  string    `gorm:"column:nadi"`
	Suhu                  string    `gorm:"column:suhu"`
	Rr                    string    `gorm:"column:rr"`
	Bb                    string    `gorm:"column:bb"`
	Kepala                string    `gorm:"column:kepala"`
	KeteranganKepala      string    `gorm:"column:keterangan_kepala"`
	Thoraks               string    `gorm:"column:thoraks"`
	KeteranganThoraks     string    `gorm:"column:keterangan_thoraks"`
	Abdomen               string    `gorm:"column:abdomen"`
	KeteranganAbdomen     string    `gorm:"column:keterangan_abdomen"`
	Ekstremitas            string    `gorm:"column:ekstremitas"`
	KeteranganEkstremitas  string    `gorm:"column:keterangan_ekstremitas"`
	Columna               string    `gorm:"column:columna"`
	KeteranganColumna     string    `gorm:"column:keterangan_columna"`
	Muskulos              string    `gorm:"column:muskulos"`
	KeteranganMuskulos    string    `gorm:"column:keterangan_muskulos"`
	Lainnya               string    `gorm:"column:lainnya"`
	ResikoJatuh            string    `gorm:"column:resiko_jatuh"`
	ResikoNutrisional      string    `gorm:"column:resiko_nutrisional"`
	KebutuhanFungsional    string    `gorm:"column:kebutuhan_fungsional"`
	DiagnosaMedis          string    `gorm:"column:diagnosa_medis"`
	DiagnosaFungsi         string    `gorm:"column:diagnosa_fungsi"`
	PenunjangLain          string    `gorm:"column:penunjang_lain"`
	Fisio                  string    `gorm:"column:fisio"`
	Okupasi                string    `gorm:"column:okupasi"`
	Wicara                 string    `gorm:"column:wicara"`
	Akupuntur              string    `gorm:"column:akupuntur"`
	TataLain               string    `gorm:"column:tatalain"`
	FrekuensiTerapi        string    `gorm:"column:frekuensi_terapi"`
	Fisioterapi            string    `gorm:"column:fisioterapi"`
	TerapiOkupasi          string    `gorm:"column:terapi_okupasi"`
	TerapiWicara           string    `gorm:"column:terapi_wicara"`
	TerapiAkupuntur        string    `gorm:"column:terapi_akupuntur"`
	TerapiLainnya          string    `gorm:"column:terapi_lainnya"`
	Edukasi                string    `gorm:"column:edukasi"`
}

// Nama tabel eksplisit
func (PenilaianMedisRalanRehabMedik) TableName() string {
	return "penilaian_medis_ralan_rehab_medik"
}
