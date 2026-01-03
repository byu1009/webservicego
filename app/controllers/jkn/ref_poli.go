package jkn

import (
	"encoding/json"
	"net/http"
	"webservicego/app/services/bpjs"

	"github.com/gin-gonic/gin"
)

func RefPoli(c *gin.Context) {
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
	res, ts, err := bpjs.DoRequest("GET", "/ref/poli", nil, cfg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	// 3. Validasi metadata BPJS
	if res.Metadata.Code != 1 {
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