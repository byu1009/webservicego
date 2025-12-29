package models

import "time"

type PenilaianAwalKeperawatanIgd struct {
	NoRawat            string    `gorm:"column:no_rawat;primaryKey"`
	Tanggal            time.Time `gorm:"column:tanggal"`
	Informasi          string    `gorm:"column:informasi"`
	KeluhanUtama       string    `gorm:"column:keluhan_utama"`
	Rpd                string    `gorm:"column:rpd"`
	Rpo                string    `gorm:"column:rpo"`
	StatusKehamilan    string    `gorm:"column:status_kehamilan"`
	Gravida            string    `gorm:"column:gravida"`
	Para               string    `gorm:"column:para"`
	Abortus            string    `gorm:"column:abortus"`
	Hpht               string    `gorm:"column:hpht"`
	Tekanan            string    `gorm:"column:tekanan"`
	Pupil              string    `gorm:"column:pupil"`
	Neurosensorik      string    `gorm:"column:neurosensorik"`
	Integumen          string    `gorm:"column:integumen"`
	Turgor             string    `gorm:"column:turgor"`
	Edema              string    `gorm:"column:edema"`
	Mukosa             string    `gorm:"column:mukosa"`
	Perdarahan         string    `gorm:"column:perdarahan"`
	JumlahPerdarahan   string    `gorm:"column:jumlah_perdarahan"`
	WarnaPerdarahan    string    `gorm:"column:warna_perdarahan"`
	Intoksikasi        string    `gorm:"column:intoksikasi"`
	Bab                string    `gorm:"column:bab"`
	Xbab               string    `gorm:"column:xbab"`
	Kbab               string    `gorm:"column:kbab"`
	Wbab               string    `gorm:"column:wbab"`
	Bak                string    `gorm:"column:bak"`
	Xbak               string    `gorm:"column:xbak"`
	Wbak               string    `gorm:"column:wbak"`
	Lbak               string    `gorm:"column:lbak"`
	Psikologis         string    `gorm:"column:psikologis"`
	Jiwa               string    `gorm:"column:jiwa"`
	Perilaku           string    `gorm:"column:perilaku"`
	Dilaporkan         string    `gorm:"column:dilaporkan"`
	Sebutkan           string    `gorm:"column:sebutkan"`
	Hubungan           string    `gorm:"column:hubungan"`
	TinggalDengan      string    `gorm:"column:tinggal_dengan"`
	KetTinggal         string    `gorm:"column:ket_tinggal"`
	Budaya             string    `gorm:"column:budaya"`
	KetBudaya          string    `gorm:"column:ket_budaya"`
	PendidikanPj       string    `gorm:"column:pendidikan_pj"`
	KetPendidikanPj    string    `gorm:"column:ket_pendidikan_pj"`
	Edukasi            string    `gorm:"column:edukasi"`
	KetEdukasi         string    `gorm:"column:ket_edukasi"`
	Kemampuan          string    `gorm:"column:kemampuan"`
	Aktifitas          string    `gorm:"column:aktifitas"`
	AlatBantu          string    `gorm:"column:alat_bantu"`
	KetBantu           string    `gorm:"column:ket_bantu"`
	Nyeri              string    `gorm:"column:nyeri"`
	Provokes           string    `gorm:"column:provokes"`
	KetProvokes        string    `gorm:"column:ket_provokes"`
	Quality            string    `gorm:"column:quality"`
	KetQuality         string    `gorm:"column:ket_quality"`
	Lokasi             string    `gorm:"column:lokasi"`
	Menyebar           string    `gorm:"column:menyebar"`
	SkalaNyeri         string    `gorm:"column:skala_nyeri"`
	Durasi             string    `gorm:"column:durasi"`
	NyeriHilang        string    `gorm:"column:nyeri_hilang"`
	KetNyeri           string    `gorm:"column:ket_nyeri"`
	PadaDokter         string    `gorm:"column:pada_dokter"`
	KetDokter          string    `gorm:"column:ket_dokter"`
	BerjalanA          string    `gorm:"column:berjalan_a"`
	BerjalanB          string    `gorm:"column:berjalan_b"`
	BerjalanC          string    `gorm:"column:berjalan_c"`
	Hasil              string    `gorm:"column:hasil"`
	Lapor              string    `gorm:"column:lapor"`
	KetLapor           string    `gorm:"column:ket_lapor"`
	Rencana            string    `gorm:"column:rencana"`
	Nip                string    `gorm:"column:nip"`
}

func (PenilaianAwalKeperawatanIgd) TableName() string {
	return "penilaian_awal_keperawatan_igd"
}