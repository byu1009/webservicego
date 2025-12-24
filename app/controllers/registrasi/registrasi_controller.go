package registrasi

import (
	"fmt"
	"net/http"
	"time"
	"webservicego/app/models"
	"webservicego/app/utils"
	"webservicego/config"

	"github.com/gin-gonic/gin"
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