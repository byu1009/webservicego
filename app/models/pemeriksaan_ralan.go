package models

import "time"

type PemeriksaanRalan struct {
	NoRawat        string    `gorm:"column:no_rawat;primaryKey"`
	TglPerawatan   time.Time `gorm:"column:tgl_perawatan"`
	JamRawat       string    `gorm:"column:jam_rawat"`
	SuhuTubuh      string    `gorm:"column:suhu_tubuh"`
	Tensi          string    `gorm:"column:tensi"`
	Nadi           string    `gorm:"column:nadi"`
	Respirasi      string    `gorm:"column:respirasi"`
	Tinggi         string    `gorm:"column:tinggi"`
	Berat          string    `gorm:"column:berat"`
	Spo2           string    `gorm:"column:spo2"`
	Gcs            string    `gorm:"column:gcs"`
	Kesadaran      string    `gorm:"column:kesadaran"`
	Keluhan        string    `gorm:"column:keluhan"`
	Pemeriksaan    string    `gorm:"column:pemeriksaan"`
	Alergi         string    `gorm:"column:alergi"`
	LingkarPerut   string    `gorm:"column:lingkar_perut"`
	Rtl            string    `gorm:"column:rtl"`
	Penilaian      string    `gorm:"column:penilaian"`
	Instruksi      string    `gorm:"column:instruksi"`
	Evaluasi       string    `gorm:"column:evaluasi"`
	Nip            string    `gorm:"column:nip"`
}

func (PemeriksaanRalan) TableName() string {
	return "pemeriksaan_ralan"
}
