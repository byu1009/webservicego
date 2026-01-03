package jkn

import (
	"net/http"
	"webservicego/app/services/bpjs"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BatalAntrean(c *gin.Context) {
	// 1. Validasi request
	var req bpjs.AntrolBatalRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		if ve, ok := err.(validator.ValidationErrors); ok {
			errorMessages := map[string]string{
				"KodeBooking": "kodebooking wajib diisi",
				"Keterangan":  "keterangan wajib diisi",
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

	data, code, message, err := bpjs.BatalAntreanService(req)

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
		"message": "Antran " + req.KodeBooking + " berhasil dibatalkan",
		"data":    data,
	})
}