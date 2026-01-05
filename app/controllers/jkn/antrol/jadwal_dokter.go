package antrol

import (
	"encoding/json"
	"net/http"
	"webservicego/app/services/bpjs"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func JadwalDokter(c *gin.Context) {
	var req struct {
		KodePoli string `json:"kodepoli" binding:"required"`
		Tanggal  string `json:"tanggal" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {

		// ambil error validator
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, fe := range ve {
				switch fe.Field() {
				case "KodePoli":
					c.JSON(http.StatusOK, gin.H{
						"code"		: 204,
						"message"	: "kodepoli wajib diisi",
					})
					return
				case "Tanggal":
					c.JSON(http.StatusOK, gin.H{
						"code"		: 204,
						"message"	: "tanggal wajib diisi",
					})
					return
				}
			}
		}

		// fallback error
		c.JSON(http.StatusOK, gin.H{
			"code"		: 204,
			"message"	: "request tidak valid",
		})
		return
	}

	// 1. Load config
	cfg, err := bpjs.LoadConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	// 2. Request ke BPJS
	res, ts, err := bpjs.DoRequest("GET", "/jadwaldokter/kodepoli/" + req.KodePoli + "/tanggal/" + req.Tanggal, nil, cfg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	// 3. Validasi metadata BPJS
	if res.Metadata.Code != 200 {
		c.JSON(http.StatusOK, gin.H{
			"code":    res.Metadata.Code,
			"message": res.Metadata.Message,
			"data":    nil,
		})
		return
	}

	// 4. Decrypt response (PAKAI LZSTRING)
	plain, err := bpjs.DecryptResponse(
		res.Response,
		cfg.ConsID,
		cfg.SecretKey,
		ts,
		true,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	// 5. Decode JSON hasil decrypt
	var data any
	if err := json.Unmarshal([]byte(plain), &data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "JSON decode error: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 6. Response sukses
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Ok",
		"data":    data,
	})
}