package rencanakontrol

import (
	"net/http"
	"time"
	"webservicego/app/models"
	"webservicego/app/utils"
	"webservicego/config"

	"github.com/gin-gonic/gin"
)

func SurkonInternal(c *gin.Context) {
	db := config.DBConnect()

	username := c.GetString("username")
	if username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"message": "Unauthorized",
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

	var req struct {
		TanggalAwal  string `json:"tanggal_awal" binding:"required"`
		TanggalAkhir string `json:"tanggal_akhir" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"message": "request tidak valid",
		})
		return
	}

	start, err := time.Parse("2006-01-02", req.TanggalAwal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"message": "format tanggal_awal harus YYYY-MM-DD",
		})
		return
	}

	end, err := time.Parse("2006-01-02", req.TanggalAkhir)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"message": "format tanggal_akhir harus YYYY-MM-DD",
		})
		return
	}

	var results []models.BridgingSuratKontrolBpjs

	err = db.
		Preload("SepAsal").
		Where(
			"tgl_rencana BETWEEN ? AND ?",
			start,
			end,
		).
		Select(`
			no_sep,
			no_surat,
			tgl_surat,
			tgl_rencana,
			kd_dokter_bpjs,
			nm_dokter_bpjs,
			kd_poli_bpjs,
			nm_poli_bpjs
		`).
		Find(&results).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"message": err.Error(),
		})
		return
	}

	var count int64

	err = db.
		Table("bridging_surat_kontrol_bpjs").
		Where(
			"tgl_rencana BETWEEN ? AND ?",
			"2025-11-07",
			"2025-11-07",
		).
		Count(&count).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": "ok",
		"data": gin.H{
			"count": count,
			"result": results,
		},
		"token": token,
	})
}

