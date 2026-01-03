package registrasi

import (
	"net/http"
	"time"
	"webservicego/app/helpers"
	"webservicego/app/models"
	"webservicego/app/services/bpjs"
	"webservicego/app/utils"
	"webservicego/config"

	"github.com/gin-gonic/gin"
)

func AddAntrian(c *gin.Context) {
	db := config.DBConnect()

	username := c.GetString("username")
	if username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Unauthorized",
		})
		return
	}

	var req struct {
		NoRawat         string  `json:"norawat" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code" : 500,
			"message": err.Error(),
		})
		return
	}

	var reg models.RegPeriksa

	err := db.
		Preload("MobileJknBpjs").
		Where("no_rawat = ?", req.NoRawat).
		First(&reg).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code"		: 201,
			"message"	: "Data registrasi dari no rawat " + req.NoRawat + " tidak ditemukan!",
		})
	}

	var namaPoli string
	err = db.Table("poliklinik").
		Select("poliklinik.nm_poli").
		Joins("JOIN maping_poli_bpjs ON maping_poli_bpjs.kd_poli_rs = poliklinik.kd_poli").
		Where("maping_poli_bpjs.kd_poli_bpjs = ?", reg.MobileJknBpjs.KodePoli).
		Scan(&namaPoli).Error
	if err != nil {
		namaPoli = ""
	}

	var namaDokter string
	err = db.Table("dokter").
		Select("dokter.nm_dokter").
		Joins("JOIN maping_dokter_dpjpvclaim ON maping_dokter_dpjpvclaim.kd_dokter = dokter.kd_dokter").
		Where("maping_dokter_dpjpvclaim.kd_dokter_bpjs = ?", reg.MobileJknBpjs.KodeDokter).
		Scan(&namaDokter).Error
	if err != nil {
		namaDokter = ""
	}

	jenisPasien := "NON JKN"
	if reg.Kdpj == "BPJ" {
		jenisPasien = "JKN"
	}

	pasienBaru, err := helpers.StrToIntE(reg.MobileJknBpjs.PasienBaru)
	kodeDokter, err := helpers.StrToIntE(reg.MobileJknBpjs.KodeDokter)
	jenisKunjungan, err := helpers.StrToIntE(reg.MobileJknBpjs.JenisKunjungan)
	angkaAntrean, err := helpers.StrToIntE(reg.MobileJknBpjs.AngkaAntrean)
	estimasiDilayani, err := helpers.StrToIntE(reg.MobileJknBpjs.EstimasiDilayani)

	reqAntrol := bpjs.AntrolTambahRequest{
		KodeBooking			: reg.MobileJknBpjs.Nobooking,
		JenisPasien			: jenisPasien,
		NomorKartu			: reg.MobileJknBpjs.NomorKartu,
		Nik					: reg.MobileJknBpjs.Nik,
		NoHp				: reg.MobileJknBpjs.NoHp,
		KodePoli			: reg.MobileJknBpjs.KodePoli,
		NamaPoli			: namaPoli,
		PasienBaru			: pasienBaru,
		Norm				: reg.MobileJknBpjs.Norm,
		TanggalPeriksa		: reg.MobileJknBpjs.TanggalPeriksa,
		KodeDokter			: kodeDokter,
		NamaDokter			: namaDokter,
		JamPraktek			: reg.MobileJknBpjs.JamPraktek,
		JenisKunjungan		: jenisKunjungan,
		NomorReferensi		: reg.MobileJknBpjs.NomorReferensi,
		NomorAntrean		: reg.MobileJknBpjs.NomorAntrean,
		AngkaAntrean		: angkaAntrean,
		EstimasiDilayani	: estimasiDilayani,
		SisaKuotaJkn		: reg.MobileJknBpjs.SisaKuotaJkn,
		KuotaJkn			: reg.MobileJknBpjs.KuotaJkn,
		SisaKuotaNonJkn		: reg.MobileJknBpjs.SisaKuotaNonJkn,
		KuotaNonJkn			: reg.MobileJknBpjs.SisaKuotaNonJkn,
		Keterangan			: "Peserta harap 30 menit lebih awal guna pencatatan administrasi.",
	}

	antrolData, antrolCode, antrolMsg, err := bpjs.TambahAntreanService(reqAntrol)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}
	
	if antrolCode != 200 && antrolCode != 1 {
		c.JSON(http.StatusOK, gin.H{
			"code":    antrolCode,
			"message": "antrian bpjs gagal dikirim",
			"messagebpjs": antrolMsg,
		})
		return
	}

	token, err := utils.GetToken(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"message": "Failed to generate token",
		})
		return
	}

	err = db.Model(&models.ReferensiMobilejknBpjs{}).
			Where("nobooking = ?", reqAntrol.KodeBooking).
			Updates(map[string]interface{}{
				"status":        "Checkin",
				"validasi":      time.Now().Format("2006-01-02 15:04:05"),
				"status_kirim":  "Sudah",
			}).Error

	c.JSON(http.StatusOK, gin.H{
		"code"				: 200,
		"message"			: "antrean bpjs berhasil",
		"data": gin.H{
			"registrasi"	: reqAntrol,
			"antrean"		: antrolData,
		},
		"token"				: token,
	})
}