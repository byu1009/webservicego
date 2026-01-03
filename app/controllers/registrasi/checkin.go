package registrasi

import (
	"net/http"
	"time"
	"webservicego/app/helpers"
	"webservicego/app/models"
	"webservicego/app/services/bpjs"
	"webservicego/config"

	"github.com/gin-gonic/gin"
)

func Checkin(c *gin.Context) {
	db := config.DBConnect()

	username := c.GetString("username")
	if username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Unauthorized",
		})
		return
	}

	// ===== 1️⃣ Bind JSON request =====
	var req struct {
		NoRawat string `json:"norawat" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	// ===== 2️⃣ Generate nomor booking BPJS =====
	noref, err := helpers.CekNoRef(req.NoRawat)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": "Gagal cek nomor referensi: " + err.Error(),
		})
		return
	}

	// ===== 3️⃣ Ambil data registrasi =====
	var reg models.RegPeriksa
	if err := db.Preload("MobileJknBpjs", "status != ?", "Batal").
		Where("no_rawat = ?", req.NoRawat).
		First(&reg).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    202,
			"message": "Data registrasi tidak ditemukan",
		})
		return
	}

	// ===== 4️⃣ Tebak hari kerja =====
	if reg.TglRegistrasi == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    203,
			"message": "TglRegistrasi kosong",
		})
		return
	}

	hariKerja, err := helpers.TebakHari(*reg.TglRegistrasi)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    204,
			"message": "Format tanggal tidak valid",
		})
		return
	}

	// ===== 5️⃣ Ambil jadwal dokter =====
	var jadwal models.Jadwal
	if err := db.Where("kd_dokter = ? AND kd_poli = ? AND hari_kerja = ?", reg.KodeDokter, reg.KodePoli, hariKerja).
		First(&jadwal).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    205,
			"message": "Jadwal dokter tidak ditemukan",
		})
		return
	}

	// ===== 6️⃣ Hitung waktu checkin =====
	now := time.Now().In(time.FixedZone("WIB", 7*3600))
	waktuCheckin := now

	// parse jam mulai jadwal
	jamMulai, err := time.Parse("15:04:05", jadwal.JamMulai)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    206,
			"message": "Format jamMulai jadwal tidak valid",
		})
		return
	}

	// gabung tanggal sekarang + jamMulai
	waktuJadwal := time.Date(now.Year(), now.Month(), now.Day(),
		jamMulai.Hour(), jamMulai.Minute(), jamMulai.Second(), 0, now.Location())

	if now.Before(waktuJadwal) {
		waktuCheckin = waktuJadwal
	}

	// ===== 7️⃣ Konversi ke millisecond =====
	waktuMillis := waktuCheckin.UnixMilli()

	// ===== 8️⃣ Cek taskid_3 sudah ada =====
	var cekTaskid models.IoAntrianTaskid
	if err := db.Where("nobooking = ? AND taskid_3 IS NOT NULL", noref).First(&cekTaskid).Error; err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    204,
			"message": "Pasien sudah checkin",
		})
		return
	}

	// ===== 9️⃣ Siapkan payload taskid 3 =====
	reqUpdate := bpjs.AntrolUpdateWaktuRequest{
		KodeBooking: noref,
		TaskId:      "3",
		Waktu:       waktuMillis,
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Taskid 3 siap dikirim",
		"payload": reqUpdate,
	})
}
