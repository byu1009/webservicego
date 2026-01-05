package taskid

import (
	"net/http"
	"webservicego/app/utils"
	"webservicego/config"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetDataTaskid(c *gin.Context) {
	db := config.DBConnect()

	username := c.GetString("username")
	if username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
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
		Tanggal		string 	`json:"tanggal" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			errorMessages := map[string]string{
				"Tanggal"	: "tanggal wajib diisi",
			}

			fe := ve[0]
			msg, ok := errorMessages[fe.Field()]
			if !ok {
				msg = "request tidak valid"
			}

			c.JSON(http.StatusOK, gin.H{
				"code":    204,
				"message": msg,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    204,
			"message": "request tidak valid",
		})
		return
	}

	var result []map[string]interface{}

	if config.Sistem.AddFarmasi == "YA" {
		err = db.
			Table("referensi_mobilejkn_bpjs r").
			// Joins("JOIN io_antrian t ON t.no_referensi = r.nobooking").
			Joins("JOIN io_antrian t ON t.no_referensi = r.no_rawat").
			Joins("JOIN io_referensi_farmasi rf ON r.nobooking = rf.kodebooking").
			Where("r.tanggalperiksa = ?", req.Tanggal).
			Select("t.*, r.nobooking").
			Scan(&result).Error

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"message": err.Error(),
			})
			return
		}
	} else {
		err = db.
			Table("referensi_mobilejkn_bpjs r").
			// Joins("JOIN io_antrian t ON t.no_referensi = r.nobooking").
			Joins("JOIN io_antrian t ON t.no_referensi = r.no_rawat").
			Where("r.tanggalperiksa = ?", req.Tanggal).
			Select("t.*, r.nobooking").
			Scan(&result).Error

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":		200,
		"message":	"Ok",
		"data":		gin.H{
			"total":	len(result),
			"result":	result,
		},
		"token":	token,
	})
}