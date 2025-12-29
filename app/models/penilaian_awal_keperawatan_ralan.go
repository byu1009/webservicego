package models

import "time"

type PenilaianAwalKeperawatanRalan struct {
	NoRawat        string    `gorm:"column:no_rawat;primaryKey"`
	Tanggal        time.Time `gorm:"column:tanggal"`
	Informasi      string    `gorm:"column:informasi"`
	Td             string    `gorm:"column:td"`
	Nadi           string    `gorm:"column:nadi"`
	Rr             string    `gorm:"column:rr"`
	Suhu           string    `gorm:"column:suhu"`
	Gcs            string    `gorm:"column:gcs"`
	Bb             string    `gorm:"column:bb"`
	Tb             string    `gorm:"column:tb"`
	Bmi            string    `gorm:"column:bmi"`
	KeluhanUtama   string    `gorm:"column:keluhan_utama"`
	Rpd            string    `gorm:"column:rpd"`
	Rpk            string    `gorm:"column:rpk"`
	Rpo            string    `gorm:"column:rpo"`
	Alergi         string    `gorm:"column:alergi"`
	AlatBantu      string    `gorm:"column:alat_bantu"`
	KetBantu       string    `gorm:"column:ket_bantu"`
	Prothesa       string    `gorm:"column:prothesa"`
	KetPro         string    `gorm:"column:ket_pro"`
	Adl            string    `gorm:"column:adl"`
	StatusPsiko    string    `gorm:"column:status_psiko"`
	KetPsiko       string    `gorm:"column:ket_psiko"`
	HubKeluarga    string    `gorm:"column:hub_keluarga"`
	TinggalDengan  string    `gorm:"column:tinggal_dengan"`
	KetTinggal     string    `gorm:"column:ket_tinggal"`
	Ekonomi        string    `gorm:"column:ekonomi"`
	Budaya         string    `gorm:"column:budaya"`
	KetBudaya      string    `gorm:"column:ket_budaya"`
	Edukasi        string    `gorm:"column:edukasi"`
	KetEdukasi     string    `gorm:"column:ket_edukasi"`
	BerjalanA      string    `gorm:"column:berjalan_a"`
	BerjalanB      string    `gorm:"column:berjalan_b"`
	BerjalanC      string    `gorm:"column:berjalan_c"`
	Hasil          string    `gorm:"column:hasil"`
	Lapor          string    `gorm:"column:lapor"`
	KetLapor       string    `gorm:"column:ket_lapor"`
	Sg1            string    `gorm:"column:sg1"`
	Nilai1         string    `gorm:"column:nilai1"`
	Sg2            string    `gorm:"column:sg2"`
	Nilai2         string    `gorm:"column:nilai2"`
	TotalHasil     string    `gorm:"column:total_hasil"`
	Nyeri          string    `gorm:"column:nyeri"`
	Provokes       string    `gorm:"column:provokes"`
	KetProvokes    string    `gorm:"column:ket_provokes"`
	Quality        string    `gorm:"column:quality"`
	KetQuality     string    `gorm:"column:ket_quality"`
	Lokasi         string    `gorm:"column:lokasi"`
	Menyebar       string    `gorm:"column:menyebar"`
	SkalaNyeri     string    `gorm:"column:skala_nyeri"`
	Durasi         string    `gorm:"column:durasi"`
	NyeriHilang    string    `gorm:"column:nyeri_hilang"`
	KetNyeri       string    `gorm:"column:ket_nyeri"`
	PadaDokter     string    `gorm:"column:pada_dokter"`
	KetDokter      string    `gorm:"column:ket_dokter"`
	Rencana        string    `gorm:"column:rencana"`
	Nip            string    `gorm:"column:nip"`
}

func (PenilaianAwalKeperawatanRalan) TableName() string {
	return "penilaian_awal_keperawatan_ralan"
}