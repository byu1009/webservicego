package models

import "time"

type PenilaianAwalKeperawatanRalanBayi struct {
	NoRawat                string    `gorm:"column:no_rawat;primaryKey"`
	Tanggal                time.Time `gorm:"column:tanggal"`
	Informasi              string    `gorm:"column:informasi"`
	Td                     string    `gorm:"column:td"`
	Nadi                   string    `gorm:"column:nadi"`
	Rr                     string    `gorm:"column:rr"`
	Suhu                   string    `gorm:"column:suhu"`
	Gcs                    string    `gorm:"column:gcs"`
	Bb                     string    `gorm:"column:bb"`
	Tb                     string    `gorm:"column:tb"`
	Lp                     string    `gorm:"column:lp"`
	Lk                     string    `gorm:"column:lk"`
	Ld                     string    `gorm:"column:ld"`
	KeluhanUtama            string    `gorm:"column:keluhan_utama"`
	Rpd                    string    `gorm:"column:rpd"`
	Rpk                    string    `gorm:"column:rpk"`
	Rpo                    string    `gorm:"column:rpo"`
	Alergi                 string    `gorm:"column:alergi"`
	Anakke                 string    `gorm:"column:anakke"`
	DariSaudara            string    `gorm:"column:darisaudara"`
	CaraLahir              string    `gorm:"column:caralahir"`
	KetCaraLahir           string    `gorm:"column:ket_caralahir"`
	UmurKelahiran          string    `gorm:"column:umurkelahiran"`
	KelainanBawaan         string    `gorm:"column:kelainanbawaan"`
	KetKelainanBawaan      string    `gorm:"column:ket_kelainan_bawaan"`
	UsiaTengkurap          string    `gorm:"column:usiatengkurap"`
	UsiaDuduk               string    `gorm:"column:usiaduduk"`
	UsiaBerdiri            string    `gorm:"column:usiaberdiri"`
	UsiaGigiPertama        string    `gorm:"column:usiagigipertama"`
	UsiaBerjalan           string    `gorm:"column:usiaberjalan"`
	UsiaBicara             string    `gorm:"column:usiabicara"`
	UsiaMembaca            string    `gorm:"column:usiamembaca"`
	UsiaMenulis            string    `gorm:"column:usiamenulis"`
	GangguanEmosi          string    `gorm:"column:gangguanemosi"`
	AlatBantu              string    `gorm:"column:alat_bantu"`
	KetBantu               string    `gorm:"column:ket_bantu"`
	Prothesa               string    `gorm:"column:prothesa"`
	KetPro                 string    `gorm:"column:ket_pro"`
	Adl                    string    `gorm:"column:adl"`
	StatusPsiko            string    `gorm:"column:status_psiko"`
	KetPsiko               string    `gorm:"column:ket_psiko"`
	HubKeluarga            string    `gorm:"column:hub_keluarga"`
	Pengasuh               string    `gorm:"column:pengasuh"`
	KetPengasuh            string    `gorm:"column:ket_pengasuh"`
	Ekonomi                string    `gorm:"column:ekonomi"`
	Budaya                 string    `gorm:"column:budaya"`
	KetBudaya              string    `gorm:"column:ket_budaya"`
	Edukasi                string    `gorm:"column:edukasi"`
	KetEdukasi             string    `gorm:"column:ket_edukasi"`
	BerjalanA              string    `gorm:"column:berjalan_a"`
	BerjalanB              string    `gorm:"column:berjalan_b"`
	BerjalanC              string    `gorm:"column:berjalan_c"`
	Hasil                  string    `gorm:"column:hasil"`
	Lapor                  string    `gorm:"column:lapor"`
	KetLapor               string    `gorm:"column:ket_lapor"`
	Sg1                    string    `gorm:"column:sg1"`
	Nilai1                 string    `gorm:"column:nilai1"`
	Sg2                    string    `gorm:"column:sg2"`
	Nilai2                 string    `gorm:"column:nilai2"`
	Sg3                    string    `gorm:"column:sg3"`
	Nilai3                 string    `gorm:"column:nilai3"`
	Sg4                    string    `gorm:"column:sg4"`
	Nilai4                 string    `gorm:"column:nilai4"`
	TotalHasil             string    `gorm:"column:total_hasil"`
	Wajah                  string    `gorm:"column:wajah"`
	NilaiWajah             string    `gorm:"column:nilaiwajah"`
	Kaki                   string    `gorm:"column:kaki"`
	NilaiKaki              string    `gorm:"column:nilaikaki"`
	Aktifitas              string    `gorm:"column:aktifitas"`
	NilaiAktifitas         string    `gorm:"column:nilaiaktifitas"`
	Menangis               string    `gorm:"column:menangis"`
	NilaiMenangis          string    `gorm:"column:nilaimenangis"`
	Bersuara               string    `gorm:"column:bersuara"`
	NilaiBersuara          string    `gorm:"column:nilaibersuara"`
	HasilNyeri             string    `gorm:"column:hasilnyeri"`
	Nyeri                  string    `gorm:"column:nyeri"`
	Lokasi                 string    `gorm:"column:lokasi"`
	Durasi                 string    `gorm:"column:durasi"`
	Frekuensi              string    `gorm:"column:frekuensi"`
	NyeriHilang            string    `gorm:"column:nyeri_hilang"`
	KetNyeri               string    `gorm:"column:ket_nyeri"`
	PadaDokter             string    `gorm:"column:pada_dokter"`
	KetDokter              string    `gorm:"column:ket_dokter"`
	Rencana                string    `gorm:"column:rencana"`
	Nip                    string    `gorm:"column:nip"`
}

func (PenilaianAwalKeperawatanRalanBayi) TableName() string {
	return "penilaian_awal_keperawatan_ralan_bayi"
}