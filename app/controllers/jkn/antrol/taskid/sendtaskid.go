package taskid

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"
	"webservicego/app/helpers"
	"webservicego/app/services/bpjs"
	"webservicego/app/utils"
	"webservicego/config"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func SendTaskid (c *gin.Context) {
	db := config.DBConnect()
	_= db

	username := c.GetString("username")
	if username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Unauthorized",
		})
		return
	}

	token, err := utils.GetToken(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"message": "Failed to generate token",
		})
		return
	}

	var req struct {
		Kbooking		string 	`json:"kodebooking" binding:"required"`
		Taskid			int 	`json:"taskid" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			errorMessages := map[string]string{
				"Kbooking"		: "kodebooking wajib diisi",
				"Taskid"		: "taskid wajib diisi",
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

	refKbo, err := helpers.CekNoRef(req.Kbooking)
	if err != nil {
		// handle error
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"message": err.Error(),
		})
		return
	}

	var result map[string]interface{}

	validTask := map[int]bool{
		1:  true,
		2:  true,
		3:  true,
		4:  true,
		5:  true,
		6:  true,
		7:  true,
		99: true,
	}

	if !validTask[req.Taskid] {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "taskid tidak valid",
		})
		return
	}

	col := "taskid_" + strconv.Itoa(req.Taskid)
	
	err = db.
		Table("io_antrian_taskid").
		Where("nobooking = ?", refKbo).
		Where(col + " IS NOT NULL").
		Select(col + " AS waktu").
		Take(&result).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"message": "Data tidak ditemukan " + err.Error(),
		})
		return
	}

	waktuRaw, ok := result["waktu"]
	if !ok || waktuRaw == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "data tidak ditemukan",
			"data":    nil,
			"token":   token,
		})
		return
	}

	waktu, ok := waktuRaw.(time.Time)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "format waktu tidak valid",
		})
		return
	}

	reqTaskid := bpjs.AntrolUpdateWaktuRequest{
		KodeBooking: req.Kbooking,
		TaskId:      strconv.Itoa(req.Taskid),
		Waktu:       helpers.ToJakartaMillis(waktu),
	}

	if req.Taskid == 5 {
		var jenisResep sql.NullString

		err = db.
			Table("io_referensi_farmasi").
			Where("kodebooking = ?", req.Kbooking).
			Select("jenis_resep").
			Take(&jenisResep).Error

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "Data farmasi tidak ditemukan: " + err.Error(),
			})
			return
		}

		// DEFAULT VALUE JIKA NULL
		if jenisResep.Valid {
			reqTaskid.JenisResep = jenisResep.String
		} else {
			reqTaskid.JenisResep = "Tidak ada"
		}
	}

	// antrolData, antrolCode, antrolMsg, err := bpjs.UpdateWaktuAntreanService(reqTaskid)

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"code":    500,
	// 		"message": err.Error(),
	// 	})
	// 	return
	// }
	
	// if antrolCode != 200 && antrolCode != 1 {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code":    antrolCode,
	// 		"message": "Taskid gagal dikirim " + antrolMsg,
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"code":			200,
		"message":		"Ok",
		"data":			reqTaskid,
		"token":		token,
	})
}