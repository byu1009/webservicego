package models

import "time"

type Pegawai struct {
	ID             int       `gorm:"column:id;primaryKey;autoIncrement"`
	Nik            string    `gorm:"column:nik"`
	Nama           string    `gorm:"column:nama"`
	Jk             string    `gorm:"column:jk"`
	Jbtn           string    `gorm:"column:jbtn"`
	JnjJabatan     string    `gorm:"column:jnj_jabatan"`
	KodeKelompok   string    `gorm:"column:kode_kelompok"`
	KodeResiko     string    `gorm:"column:kode_resiko"`
	KodeEmergency  string    `gorm:"column:kode_emergency"`
	Departemen     string    `gorm:"column:departemen"`
	Bidang         string    `gorm:"column:bidang"`
	SttsWp         string    `gorm:"column:stts_wp"`
	SttsKerja      string    `gorm:"column:stts_kerja"`
	Npwp           string    `gorm:"column:npwp"`
	Pendidikan     string    `gorm:"column:pendidikan"`
	Gapok          float64   `gorm:"column:gapok"`
	TmpLahir       string    `gorm:"column:tmp_lahir"`
	TglLahir       time.Time `gorm:"column:tgl_lahir"`
	Alamat         string    `gorm:"column:alamat"`
	Kota           string    `gorm:"column:kota"`
	MulaiKerja     time.Time `gorm:"column:mulai_kerja"`
	MsKerja        string    `gorm:"column:ms_kerja"`
	IndexIns       string    `gorm:"column:indexins"`
	Bpd            string    `gorm:"column:bpd"`
	Rekening       string    `gorm:"column:rekening"`
	SttsAktif      string    `gorm:"column:stts_aktif"`
	WajibMasuk     string    `gorm:"column:wajibmasuk"`
	Pengurang      string    `gorm:"column:pengurang"`
	Indek          string    `gorm:"column:indek"`
	MulaiKontrak   time.Time `gorm:"column:mulai_kontrak"`
	CutiDiambil    string    `gorm:"column:cuti_diambil"`
	Dankes         string    `gorm:"column:dankes"`
	Photo          string    `gorm:"column:photo"`
	NoKtp          string    `gorm:"column:no_ktp"`
}

func (Pegawai) TableName() string {
	return "pegawai"
}