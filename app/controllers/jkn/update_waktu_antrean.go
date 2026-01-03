package jkn

import (
	"net/http"
	"webservicego/app/services/bpjs"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func UpdateWaktuAntrean(c *gin.Context) {
	var req bpjs.AntrolUpdateWaktuRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {

			errorMessages := map[string]string{
				"KodeBooking": "kodebooking wajib diisi",
				"Taskid":      "taskid wajib diisi",
				"Waktu":       "waktu wajib diisi",
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

	if req.JenisResep != "" {
		switch req.JenisResep {
		case "Tidak ada", "Racikan", "Non racikan":
			// valid
		default:
			c.JSON(http.StatusOK, gin.H{
				"code":    204,
				"message": "jenisresep harus: Tidak ada, Racikan, atau Non racikan",
			})
			return
		}
	}

	data, code, message, err := bpjs.UpdateWaktuAntreanService(req)

	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	if code != 200 && code != 1 {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Update waktu " + req.TaskId + " berhasil",
		"data":    data,
	})
}