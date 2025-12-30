package registrasi

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"webservicego/app/helpers"
	"webservicego/app/models"
	"webservicego/app/utils"
	"webservicego/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetData(c *gin.Context) {
	var req struct {
		TglAwal  string `json:"tglawal" binding:"required"`
		TglAkhir string `json:"tglakhir" binding:"required"`
	}

	username := c.GetString("username")
	if username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Unauthorized",
		})
		return
	}

	// Bind JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("BindJSON error:", err)
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Parameter tglawal dan tglakhir wajib diisi atau JSON tidak valid",
		})
		return
	}

	// Parse tanggal dari JSON struct
	start, err := time.Parse("2006-01-02", req.TglAwal) // pakai req.TglAwal
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "Tanggal awal tidak valid"})
		return
	}

	end, err := time.Parse("2006-01-02", req.TglAkhir) // pakai req.TglAkhir
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "Tanggal akhir tidak valid"})
		return
	}

	var data []models.RegPeriksa
	db := config.DBConnect()

	// Query GORM
	if err := db.Where("tgl_registrasi BETWEEN ? AND ?", start.Format("2006-01-02"), end.Format("2006-01-02")).Find(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Terjadi kesalahan"})
		return
	}

	token, err := utils.GetToken(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to generate token",
		})
		return
	}

	// Response JSON
	c.JSON(http.StatusOK, gin.H{
		"code"		: 200,
		"message"	: "Data ada",
		"data"		: data,
		"token"		: token,
	})
}

func PostData(c *gin.Context) {
	db := config.DBConnect()

	tx := db.Begin()
	if tx.Error != nil {
		c.JSON(500, gin.H{"message": "Gagal memulai transaksi"})
		return
	}

	// Pastikan rollback kalau panic
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // biarkan Gin recovery middleware handle
		}
	}()

	username := c.GetString("username")
	if username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Unauthorized",
		})
		return
	}

	var req struct {
		Dokter         string  `json:"dokter" binding:"required"`
		Poli           string  `json:"poli" binding:"required"`
		NoRkmMedis     string  `json:"norkmmedis" binding:"required"`
		TglPeriksa     string  `json:"tglperiksa" binding:"required,datetime=2006-01-02"`
		JamPeriksa     string  `json:"jamperiksa" binding:"required"`
		CaraBayar      string  `json:"carabayar" binding:"required"`
		NoReferensi    *string `json:"noreferensi"`
		JenisKunjungan int     `json:"jeniskunjungan" binding:"required,oneof=1 2 3 4"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if !helpers.ValidTanggal(req.TglPeriksa) {
		c.JSON(http.StatusOK, gin.H{
			"code":    204,
			"message": "Format tanggal harus Y-m-d",
		})
		return
	}

	var pasien models.Pasien
	if err := tx.First(&pasien, "no_rkm_medis = ?", req.NoRkmMedis).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, gin.H{"code": 204, "message": "Pasien tidak ditemukan"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Gagal mengambil data pasien"})
		return
	}

	umurStr, err := helpers.ToYMD(pasien.TglLahir)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Format tanggal lahir pasien tidak valid"})
		return
	}

	umurHit, err := helpers.HitungUmur(umurStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Gagal hitung umur"})
		return
	}

	umurSplit := strings.Fields(umurHit)
	if len(umurSplit) < 2 {
		c.JSON(http.StatusOK, gin.H{"message": "Format umur tidak valid"})
		return
	}
	umurInt, _ := strconv.Atoi(umurSplit[0])
	umurSatuan := umurSplit[1]

	// Ambil no_reg terbaru
	var lastNoRegStr *string
	err = tx.Model(&models.RegPeriksa{}).
		Select("no_reg").
		Where("kd_dokter = ?", req.Dokter).
		Where("kd_poli = ?", req.Poli).
		Where("tgl_registrasi = ?", req.TglPeriksa).
		Order("CAST(no_reg AS UNSIGNED) DESC").
		Limit(1).
		Scan(&lastNoRegStr).Error

	lastNoRegNum := 1
	if err == nil && lastNoRegStr != nil {
		if n, e := strconv.Atoi(*lastNoRegStr); e == nil {
			lastNoRegNum = n + 1
		}
	}
	noReg := fmt.Sprintf("%03d", lastNoRegNum)

	// Ambil no_rawat terbaru
	var lastNoRawatStr *string
	err = tx.Model(&models.RegPeriksa{}).
		Select("no_rawat").
		Where("tgl_registrasi = ?", req.TglPeriksa).
		Order("no_rawat DESC").
		Limit(1).
		Scan(&lastNoRawatStr).Error

	lastNoRawatNum := 1
	if err == nil && lastNoRawatStr != nil && len(*lastNoRawatStr) >= 6 {
		last6 := (*lastNoRawatStr)[len(*lastNoRawatStr)-6:]
		if n, e := strconv.Atoi(last6); e == nil {
			lastNoRawatNum = n + 1
		}
	}
	tglRawat := strings.ReplaceAll(req.TglPeriksa, "-", "/")
	noRawat := fmt.Sprintf("%s/%06d", tglRawat, lastNoRawatNum)

	// Status pasien
	var cnt int64
	tx.Model(&models.RegPeriksa{}).Where("no_rkm_medis = ?", req.NoRkmMedis).Count(&cnt)
	sttsDaftar := "Lama"
	if cnt < 1 {
		sttsDaftar = "Baru"
	}

	tx.Model(&models.RegPeriksa{}).
		Where("no_rkm_medis = ?", req.NoRkmMedis).
		Where("kd_poli = ?", req.Poli).
		Where("kd_dokter = ?", req.Dokter).
		Count(&cnt)
	statusPoli := "Lama"
	if cnt < 1 {
		statusPoli = "Baru"
	}

	biayaReg := 0
	stts := "Belum"

	hariPeriksa, err := helpers.TebakHari(req.TglPeriksa)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 204, "message": "Tanggal tidak valid"})
		return
	}

	// Insert reg_periksa
	reg := models.RegPeriksa{
		NoReg:         &noReg,
		NoRawat:       noRawat,
		TglRegistrasi: &req.TglPeriksa,
		JamReg:        &req.JamPeriksa,
		KodeDokter:    &req.Dokter,
		Norm:          &req.NoRkmMedis,
		KodePoli:      &req.Poli,
		PJawab:        &pasien.NamaKeluarga,
		Almtpj:        &pasien.AlamatPj,
		Hubunganpj:    &pasien.Keluarga,
		BiayaReg:      &biayaReg,
		Stts:          &stts,
		SttsDaftar:    sttsDaftar,
		StatusLanjut:  "Ralan",
		Kdpj:          req.CaraBayar,
		Umurdaftar:    &umurInt,
		Sttsumur:      &umurSatuan,
		Statusbayar:   "Belum Bayar",
		StatusPoli:    statusPoli,
	}

	if err := tx.Create(&reg).Error; err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"message": "Gagal insert registrasi", "error": err.Error()})
		return
	}

	// Query jadwal + pasien BPJS
	sql := `
	SELECT
		rp.no_rawat,
		p.no_peserta,
		p.no_ktp,
		p.no_tlp,
		mp.kd_poli_bpjs,
		IF((SELECT COUNT(*) FROM reg_periksa WHERE no_rkm_medis = p.no_rkm_medis)=0,'1','0') AS pasienbaru,
		p.no_rkm_medis,
		rp.tgl_registrasi,
		md.kd_dokter_bpjs,
		j.jam_mulai AS jammulai,
		CONCAT(DATE_FORMAT(j.jam_mulai, '%H:%i'), '-', DATE_FORMAT(j.jam_selesai, '%H:%i')) AS jampraktek
	FROM reg_periksa rp
	JOIN pasien p ON rp.no_rkm_medis = p.no_rkm_medis
	JOIN maping_dokter_dpjpvclaim md ON rp.kd_dokter = md.kd_dokter
	JOIN maping_poli_bpjs mp ON rp.kd_poli = mp.kd_poli_rs
	JOIN jadwal j ON j.kd_dokter = rp.kd_dokter
		AND j.hari_kerja = ?
	WHERE rp.no_rawat = ?
	`

	rows, err := tx.Raw(sql, hariPeriksa, noRawat).Rows()
	if err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"code": 500, "message": "Query gagal"})
		return
	}
	defer rows.Close()

	columns, _ := rows.Columns()
	var results []map[string]interface{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}
		if err := rows.Scan(valuePtrs...); err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"message": "Scan data gagal"})
			return
		}
		rowMap := make(map[string]interface{})
		for i, col := range columns {
			if b, ok := values[i].([]byte); ok {
				rowMap[col] = string(b)
			} else {
				rowMap[col] = values[i]
			}
		}
		results = append(results, rowMap)
	}

	if len(results) == 0 {
		tx.Rollback()
		c.JSON(500, gin.H{"message": "Data jadwal tidak ditemukan"})
		return
	}

	getString := func(row map[string]interface{}, key string) string {
		if val, ok := row[key]; ok && val != nil {
			switch v := val.(type) {
			case string:
				return v
			case []byte:
				return string(v)
			}
		}
		return ""
	}

	jammulaiStr := getString(results[0], "jammulai")
	estimasiMs, err := helpers.HitungEstimasiLayanan(req.TglPeriksa, jammulaiStr, noReg, config.EstimasiLayan)
	if err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	noBooking, err := helpers.GenerateNobooking(tx, req.TglPeriksa)
	if err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"message": "Gagal generate nomor rawat", "error": err.Error()})
		return
	}

	// c.JSON(200, gin.H{
	// 	"nobooking": noBooking,
	// })
	// return

	nomorReferensi := "-"
	if req.NoReferensi != nil {
		nomorReferensi = *req.NoReferensi
	}

	intFromInterface := func(v interface{}) int {
		if v == nil {
			return 0
		}
		switch val := v.(type) {
		case int:
			return val
		case int64:
			return int(val)
		case float64:
			return int(val)
		case []byte:
			i, _ := strconv.Atoi(string(val))
			return i
		case string:
			i, _ := strconv.Atoi(val)
			return i
		default:
			return 0
		}
	}

	regAntrol := models.ReferensiMobilejknBpjs{
		Nobooking:         	noBooking,
		NoRawat:           	reg.NoRawat,
		NomorKartu:        	getString(results[0], "no_peserta"),
		Nik:               	getString(results[0], "no_ktp"),
		NoHp:              	getString(results[0], "no_tlp"),
		KodePoli:          	getString(results[0], "kd_poli_bpjs"),
		PasienBaru:        	getString(results[0], "pasienbaru"),
		Norm:              	getString(results[0], "no_rkm_medis"),
		TanggalPeriksa:    	req.TglPeriksa,
		KodeDokter:        	getString(results[0], "kd_dokter_bpjs"),
		JamPraktek:        	getString(results[0], "jampraktek"),
		JenisKunjungan:    	strconv.Itoa(req.JenisKunjungan),
		NomorReferensi:    	nomorReferensi,
		NomorAntrean: 		*reg.KodePoli + "-" + *reg.NoReg,
		AngkaAntrean: 		*reg.NoReg,
		EstimasiDilayani:  	strconv.FormatInt(estimasiMs, 10),
		SisaKuotaJkn: 		intFromInterface(results[0]["sisa_kuota"]),
		KuotaJkn: 			intFromInterface(results[0]["kuota"]),
		SisKuotaNonJkn: 	intFromInterface(results[0]["sisa_kuota"]),
		KuotaNonJkn: 		intFromInterface(results[0]["kuota"]),
		Status: 			"Belum",
		Validasi: 			"0000-00-00 00-00-00",
		StatusKirim: 		"Belum",
	}

	if err := tx.Create(&regAntrol).Error; err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"message": "Gagal insert referensi mobile jkn", "error": err.Error()})
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(500, gin.H{"message": "Gagal commit transaksi"})
		return
	}

	token, err := utils.GetToken(username)
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "message": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": regAntrol,
		"token": token,
	})
}
