package registrasi

import (
	"errors"
	"net/http"
	"strconv"
	"time"
	"webservicego/app/helpers"
	"webservicego/app/models"
	"webservicego/app/services/bpjs"
	"webservicego/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BatalAntrian(c *gin.Context) {
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
		NoRawat			string  `json:"norawat" binding:"required"`
		KetBatal        string  `json:"ketbatal" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code"		: 400,
			"message"	: err.Error(),
		})
		return
	}

	var reg models.RegPeriksa
	err := db.
			Where("no_rawat = ?", req.NoRawat).
			Preload("MobileJknBpjs", "status != ?", "Batal").
			First(&reg).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code"		: 201,
			"message"	: "Data registrasi dari no rawat " + req.NoRawat + " tidak ditemukan!",
		})
	}

	var cek models.IoAntrianTaskid
	err = db.
		First(&cek, "nobooking = ?", reg.MobileJknBpjs.Nobooking).
		Error

	// 1️⃣ Data tidak ditemukan
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, gin.H{
				"code":    204,
				"message": "Data referensi taskid tidak ditemukan",
			})
			return
		}

		// error lain (DB error, dll)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"message": err.Error(),
		})
		return
	}

	// 2️⃣ Data ada, tapi taskid_3 belum ada
	if cek.Taskid3 == nil {
		ts := helpers.NowMillisWIB()

		reqUpdateWaktu := bpjs.AntrolUpdateWaktuRequest{
			KodeBooking: reg.MobileJknBpjs.Nobooking,
			TaskId: strconv.Itoa(99),
			Waktu: ts,
		}

		antrolData, antrolCode, antrolMsg, err := bpjs.UpdateWaktuAntreanService(reqUpdateWaktu)
		_= antrolData

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
				"message": antrolMsg,
			})
			return
		}

		err = db.Model(&models.ReferensiMobilejknBpjs{}).
			Where("nobooking = ?", reqUpdateWaktu.KodeBooking).
			Updates(map[string]interface{}{
				"status":        "Batal",
				"validasi":      time.Now().Format("2006-01-02 15:04:05"),
				"status_kirim":  "Sudah",
			}).Error

		c.JSON(http.StatusOK, gin.H{
			"code":    antrolCode,
			"message": "Pendaftaran " + reqUpdateWaktu.KodeBooking + " berhasil dibatalkan",
		})
		return
	}

	reqBatal := bpjs.AntrolBatalRequest{
		KodeBooking: reg.MobileJknBpjs.Nobooking,
		Keterangan: req.KetBatal,
	}

	antrolData, antrolCode, antrolMsg, err := bpjs.BatalAntreanService(reqBatal)
	_= antrolData

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
			"message": antrolMsg,
		})
		return
	}

	err = db.Model(&models.ReferensiMobilejknBpjs{}).
		Where("nobooking = ?", reqBatal.KodeBooking).
		Updates(map[string]interface{}{
			"status":        "Batal",
			"validasi":      time.Now().Format("2006-01-02 15:04:05"),
			"status_kirim":  "Sudah",
		}).Error

	c.JSON(http.StatusOK, gin.H{
		"code":    antrolCode,
		"message": "Pendaftaran " + reqBatal.KodeBooking + " berhasil dibatalkan",
	})
}