package jkn

import (
	"encoding/json"
	"net/http"
	"webservicego/app/services/bpjs"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RefPasienFingerpoli(c *gin.Context) {
	// 1. Bind request
	var req struct {
		JenisIdentitas	string `json:"jenisidentitas" binding:"required"`
		NoIdentitas		string `json:"noidentitas" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {

		// ambil error validator
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, fe := range ve {
				switch fe.Field() {
				case "JenisIdentitas":
					c.JSON(http.StatusOK, gin.H{
						"code"		: 204,
						"message"	: "jenisidentitas wajib diisi",
					})
					return
				case "NoIdentitas":
					c.JSON(http.StatusOK, gin.H{
						"code"		: 204,
						"message"	: "noidentitas wajib diisi",
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
		"/antrean/getlisttask",
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