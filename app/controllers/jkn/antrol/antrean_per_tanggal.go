package antrol

import (
	"net/http"
	"webservicego/app/services/bpjs"

	"github.com/gin-gonic/gin"
)

func AntreanPerTanggal(c *gin.Context) {
	var req bpjs.AntrolPerTanggalRequest

	if err := c.ShouldBindJSON(&req); err != nil || req.Tanggal == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "tanggal wajib diisi",
			"data":    nil,
		})
		return
	}

	data, err := bpjs.AntreanPerTanggalService(req)
	
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    204,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Ok",
		"data":    data,
	})
}