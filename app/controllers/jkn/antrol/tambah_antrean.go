package antrol

import (
	"net/http"
	"webservicego/app/services/bpjs"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func TambahAntrean(c *gin.Context) {
	var req bpjs.AntrolTambahRequest

	// 1. Bind & validasi
	if err := c.ShouldBindJSON(&req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {

			errorMessages := map[string]string{
				"KodeBooking":      "kodebooking wajib diisi",
				"JenisPasien":      "jenispasien wajib diisi",
				"NomorKartu":       "nomorkartu wajib diisi",
				"Nik":              "nik wajib diisi",
				"NoHp":             "nohp wajib diisi",
				"KodePoli":         "kodepoli wajib diisi",
				"NamaPoli":         "namapoli wajib diisi",
				"PasienBaru":       "pasienbaru wajib diisi",
				"Norm":             "norm wajib diisi",
				"TanggalPeriksa":   "tanggalperiksa wajib diisi",
				"KodeDokter":       "kodedokter wajib diisi",
				"NamaDokter":       "namadokter wajib diisi",
				"JamPraktek":       "jampraktek wajib diisi",
				"JenisKunjungan":   "jeniskunjungan wajib diisi",
				"NomorReferensi":   "nomorreferensi wajib diisi",
				"NomorAntrean":     "nomorantrean wajib diisi",
				"AngkaAntrean":     "angkaantrean wajib diisi",
				"EstimasiDilayani": "estimasidilayani wajib diisi",
				"SisaKuotaJkn":     "sisakuotajkn wajib diisi",
				"KuotaJkn":         "kuotajkn wajib diisi",
				"SisaKuotaNonJkn":  "sisakuotanonjkn wajib diisi",
				"KuotaNonJkn":      "kuotanonjkn wajib diisi",
				"Keterangan":       "keterangan wajib diisi",
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

	data, code, message, err := bpjs.TambahAntreanService(req)

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
		"message": "Pendaftaran antrean berhasil",
		"data":    data,
	})
}