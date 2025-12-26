package jkn

import (
	"encoding/json"
	"net/http"
	"strconv"
	"webservicego/app/services/bpjs"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func RefDokter(c *gin.Context) {
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
	res, ts, err := bpjs.DoRequest("GET", "/ref/dokter", nil, cfg)
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

func RefFingerpoli(c *gin.Context) {
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
	res, ts, err := bpjs.DoRequest("GET", "/ref/poli/fp", nil, cfg)
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

func TambahAntrean(c *gin.Context) {
	var req struct {
		KodeBooking         string 	`json:"kodebooking" binding:"required"`
		JenisPasien         string 	`json:"jenispasien" binding:"required"`
		NomorKartu          string 	`json:"nomorkartu" binding:"required"`
		Nik                 string 	`json:"nik" binding:"required"`
		NoHp                string 	`json:"nohp" binding:"required"`
		KodePoli            string 	`json:"kodepoli" binding:"required"`
		NamaPoli            string 	`json:"namapoli" binding:"required"`
		PasienBaru			int		`json:"pasienbaru" binding:"oneof=0 1"`
		Norm                string 	`json:"norm" binding:"required"`
		TanggalPeriksa      string 	`json:"tanggalperiksa" binding:"required"`
		KodeDokter          int    	`json:"kodedokter" binding:"required"`
		NamaDokter          string 	`json:"namadokter" binding:"required"`
		JamPraktek          string 	`json:"jampraktek" binding:"required"`
		JenisKunjungan      int   	`json:"jeniskunjungan" binding:"required"`
		NomorReferensi      string 	`json:"nomorreferensi" binding:"required"`
		NomorAntrean        string 	`json:"nomorantrean" binding:"required"`
		AngkaAntrean        int    	`json:"angkaantrean" binding:"required"`
		EstimasiDilayani    int64  	`json:"estimasidilayani" binding:"required"`
		SisaKuotaJkn        int    	`json:"sisakuotajkn" binding:"required"`
		KuotaJkn            int    	`json:"kuotajkn" binding:"required"`
		SisaKuotaNonJkn     int    	`json:"sisakuotanonjkn" binding:"required"`
		KuotaNonJkn         int    	`json:"kuotanonjkn" binding:"required"`
		Keterangan          string 	`json:"keterangan" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {

			errorMessages := map[string]string{
				"KodeBooking":      "kodebooking wajib diisi",
				"JenisPasien":      "jenispasien wajib diisi",
				"NomorKartu":       "jeniskartu wajib diisi",
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
		"/antrean/add",
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

func UpdateWaktuAntrean(c *gin.Context) {
	var req struct {
		KodeBooking string  `json:"kodebooking" binding:"required"`
		Taskid      int     `json:"taskid" binding:"required"`
		Waktu       int64   `json:"waktu" binding:"required"`
		JenisResep  *string `json:"jenisresep"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {

		// Validation error (required, type, dll)
		if ve, ok := err.(validator.ValidationErrors); ok {

			errorMessages := map[string]string{
				"KodeBooking": "kodebooking wajib diisi",
				"Taskid":      "taskid wajib diisi",
				"Waktu":       "waktu wajib diisi",
			}

			fe := ve[0] // ambil error pertama
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

		// selain validation error (type mismatch, overflow, dll)
		c.JSON(http.StatusOK, gin.H{
			"code":    204,
			"message": "request tidak valid",
		})
		return
	}

	if req.JenisResep != nil {

		switch *req.JenisResep {
		case "Tidak ada", "Racikan", "Non racikan":
			// valid → lanjut
		default:
			c.JSON(http.StatusOK, gin.H{
				"code":    204,
				"message": "jenisresep harus: Tidak ada, Racikan, atau Non racikan",
			})
			return
		}
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
		"/antrean/updatewaktu",
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

func BatalAntrean(c *gin.Context) {
	var req struct {
		KodeBooking		string  `json:"kodebooking" binding:"required"`
		Keterangan      string  `json:"keterangan" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {

		// Validation error (required, type, dll)
		if ve, ok := err.(validator.ValidationErrors); ok {

			errorMessages := map[string]string{
				"KodeBooking"	: "kodebooking wajib diisi",
				"Keterangan"	: "keterangan wajib diisi",
			}

			fe := ve[0] // ambil error pertama
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

		// selain validation error (type mismatch, overflow, dll)
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
		"/antrean/batal",
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

func ListTaskid(c *gin.Context) {
	// 1. Bind request
	var req struct {
		KodeBooking string `json:"kodebooking" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "kodebooking wajib diisi",
			"data":    nil,
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

func AntreanPerTanggal(c *gin.Context) {
	// 1. Bind request
	var req struct {
		Tanggal string `json:"tanggal" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "tanggal wajib diisi",
			"data":    nil,
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
		"GET",
		"/antrean/pendaftaran/tanggal/" + req.Tanggal,
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

func AntreanPerKodebooking(c *gin.Context) {
	// 1. Bind request
	var req struct {
		KodeBooking string `json:"kodebooking" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "kodebooking wajib diisi",
			"data":    nil,
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
		"GET",
		"/antrean/pendaftaran/kodebooking/" + req.KodeBooking,
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

func AntreanBelumDilayaniDetail(c *gin.Context) {
	// 1. Bind request
	var req struct {
		KodePoli   string `json:"kodepoli" binding:"required"`
		KodeDokter int    `json:"kodedokter" binding:"required"`
		Hari       int    `json:"hari" binding:"required"`
		JamPraktek string `json:"jampraktek" binding:"required"`
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
				case "KodeDokter":
					c.JSON(http.StatusOK, gin.H{
						"code"		: 204,
						"message"	: "kodedokter wajib diisi",
					})
					return
				case "Hari":
					c.JSON(http.StatusOK, gin.H{
						"code"		: 204,
						"message"	: "hari wajib diisi",
					})
					return
				case "JamPraktek":
					c.JSON(http.StatusOK, gin.H{
						"code"		: 204,
						"message"	: "jampraktek wajib diisi",
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
		"GET",
		"/antrean/pendaftaran/kodepoli/" + req.KodePoli + "/kodedokter/" + strconv.Itoa(req.KodeDokter) + "/hari/" + strconv.Itoa(req.Hari) + "/jampraktek/" + req.JamPraktek,
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