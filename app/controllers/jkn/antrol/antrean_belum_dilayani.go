package antrol

import (
	"encoding/json"
	"net/http"
	"webservicego/app/services/bpjs"

	"github.com/gin-gonic/gin"
)

func AntreanBelumDilayani(c *gin.Context) {
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
		"GET",
		"/antrean/pendaftaran/aktif",
		nil,
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
