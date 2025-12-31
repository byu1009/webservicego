package bpjs

type AntrolTambahRequest struct {
	KodeBooking         string `json:"kodebooking" binding:"required"`
	JenisPasien         string `json:"jenispasien" binding:"required"`
	NomorKartu          string `json:"nomorkartu" binding:"required"`
	Nik                 string `json:"nik" binding:"required"`
	NoHp                string `json:"nohp" binding:"required"`
	KodePoli            string `json:"kodepoli" binding:"required"`
	NamaPoli            string `json:"namapoli" binding:"required"`
	PasienBaru          int    `json:"pasienbaru" binding:"oneof=0 1"`
	Norm                string `json:"norm" binding:"required"`
	TanggalPeriksa      string `json:"tanggalperiksa" binding:"required"`
	KodeDokter          int    `json:"kodedokter" binding:"required"`
	NamaDokter          string `json:"namadokter" binding:"required"`
	JamPraktek          string `json:"jampraktek" binding:"required"`
	JenisKunjungan      int    `json:"jeniskunjungan" binding:"required"`
	NomorReferensi      string `json:"nomorreferensi" binding:"required"`
	NomorAntrean        string `json:"nomorantrean" binding:"required"`
	AngkaAntrean        int    `json:"angkaantrean" binding:"required"`
	EstimasiDilayani    int  `json:"estimasidilayani" binding:"required"`
	SisaKuotaJkn        int    `json:"sisakuotajkn" binding:"required"`
	KuotaJkn            int    `json:"kuotajkn" binding:"required"`
	SisaKuotaNonJkn     int    `json:"sisakuotanonjkn" binding:"required"`
	KuotaNonJkn         int    `json:"kuotanonjkn" binding:"required"`
	Keterangan          string `json:"keterangan" binding:"required"`
}