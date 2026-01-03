package jkn

import (
	"encoding/json"
	"net/http"
	"webservicego/app/services/bpjs"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func TambahAntreanFarmasi(c *gin.Context) {
	var req struct {
		KodeBooking         string 	`json:"kodebooking" binding:"required"`
		JenisResep         	string 	`json:"jenisresep" binding:"required"`
		NomorAntrean        int 	`json:"nomorantrean" binding:"required"`
		Keterangan          string 	`json:"keterangan" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {

			errorMessages := map[string]string{
				"KodeBooking"	: "kodebooking wajib diisi",
				"JenisResep"	: "jenisresep wajib diisi",
				"NomorAntrean"	: "nomorantrean wajib diisi",
				"Keterangan"	: "keterangan wajib diisi",
			}

			// Ambil error pertama saja (sesuai standar BPJS)
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

		// fallback jika bukan validation error
		c.JSON(http.StatusOK, gin.H{
			"code":    204,
			"message": "request tidak valid",
		})
		return
	}

	// 2. Load BPJS config
	cfg, err := bpjs.LoadConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	// 3. Request ke BPJS
	res, ts, err := bpjs.DoRequest(
		"POST",
		"/antrean/farmasi/add",
		req,
		cfg,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	// 4. Validasi metadata BPJS
	if res.Metadata.Code != 200 && res.Metadata.Code != 1 {
		c.JSON(http.StatusOK, gin.H{
			"code":    res.Metadata.Code,
			"message": res.Metadata.Message,
			"data":    nil,
		})
		return
	}

	// 5. Decrypt response (TANPA LZSTRING)
	plain, err := bpjs.DecryptResponse(
		res.Response,
		cfg.ConsID,
		cfg.SecretKey,
		ts,
		false,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	// 6. Decode JSON
	var data any
	if err := json.Unmarshal([]byte(plain), &data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "JSON decode error: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 7. Response sukses
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Ok",
		"data":    data,
	})
}