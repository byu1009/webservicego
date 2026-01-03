package taskid

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"
	"webservicego/app/helpers"
	"webservicego/app/models"
	"webservicego/app/utils"
	"webservicego/config"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func PostTaskid (c *gin.Context) {
	db := config.DBConnect()

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
		KodeBooking		string 	`json:"kodebooking" binding:"required"`
		Taskid			int		`json:"taskid" binding:"required"`
		// Waktu			string	`json:"waktu" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			errorMessages := map[string]string{
				"KodeBooking"	: "kodebooking wajib diisi",
				"Taskid"		: "taskid wajib diisi",
				// "Waktu"			: "waktu wajib diisi",
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

	twaktu := helpers.NowUTC()
	taskid := req.Taskid
	refKbo := req.KodeBooking
	
	// // format1 := regexp.MustCompile(`^\d{4}\/\d{2}\/\d{2}\/\d{6}$`)
	// // format2 := regexp.MustCompile(`^\d{14}$`)


	// switch {
	// 	case format1.MatchString(req.KodeBooking):
	// 		refKbo = "a/a"

	// 	case format2.MatchString(req.KodeBooking):
	// 		refKbo = "aa"

	// 	default:
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"code":    400,
	// 			"message": "Format kodebooking tidak sesuai",
	// 		})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"code":   200,
	// 	"refKbo": refKbo,
	// })

	var task models.IoAntrianTaskid

	err = db.Where("nobooking = ?", refKbo).
		First(&task).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			task = models.IoAntrianTaskid{
				Nobooking: refKbo,
			}

			if err := db.Create(&task).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    500,
					"message": "Gagal insert taskid",
					"error":   err.Error(),
				})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "Database error",
				"error":   err.Error(),
			})
			return
		}		
	}

	if taskid > 3 && taskid <= 7 {
		prevField := fmt.Sprintf("Taskid%d", taskid-1)

		v := reflect.ValueOf(&task).Elem().FieldByName(prevField)

		if v.IsValid() && !v.IsNil() {
			prevTime := v.Interface().(*time.Time)

			if !twaktu.After(*prevTime) {
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"code":    422,
					"message": fmt.Sprintf(
						"Waktu TaskID %d harus lebih besar dari TaskID %d",
						taskid, taskid-1,
					),
				})
				return
			}
		}
	}

	if taskid != 99 && task.Taskid99 != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": fmt.Sprintf(
				"Tidak dapat memproses TaskID %d. TaskID 99 telah mengunci proses.",
				taskid,
			),
		})
		return
	}

	if taskid == 99 {
		if task.Taskid3 == nil {
			c.JSON(http.StatusOK, gin.H{
				"code":		208,
				"message":	"Tidak dapat membatalkan karena pasien belum dilayani di Poliklinik",
			})
			return
		}

		if task.Taskid5 != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":		208,
				"message":	"Tidak dapat membatalkan karena pasien sudah selesai periksa Poliklinik",
			})
			return
		}

		if task.Taskid99 != nil || task.Taskid99Send != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":		208,
				"message":	"Tidak dapat membatalkan karena sudah dibatalkan",
			})
			return
		}

		err := db.Model(&models.IoAntrianTaskid{}).
			Where("nobooking = ?", refKbo).
			Updates(map[string]interface{}{
				"taskid_99":      twaktu,
			}).Error

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "Gagal update data",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":		200,
			"message":	"Pasien berhasil dibatalkan",
		})
		return
	}

	flow := []int{3, 4, 5, 6, 7}
	pos := -1
	for i, v := range flow {
		if v == taskid {
			pos = i
			break
		}
	}

	if pos == -1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "TaskID tidak valid",
		})
		return
	}

	if taskid != 3 {
		prevTaskID := flow[pos-1]
		prevTime := getTaskValue(task, prevTaskID, "time")

		// ❌ Task sebelumnya belum ada
		if prevTime == nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": fmt.Sprintf(
					"TaskID %d tidak dapat diproses sebelum TaskID %d terisi.",
					taskid, prevTaskID,
				),
			})
			return
		}

		// ❌ Waktu mundur
		if twaktu.Before(*prevTime) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": fmt.Sprintf(
					"Waktu TaskID %d tidak boleh lebih kecil dari TaskID %d.",
					taskid, prevTaskID,
				),
			})
			return
		}
	}

	sendTime := getTaskValue(task, taskid, "send")

	if sendTime != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": fmt.Sprintf(
				"TaskID %d tidak dapat diupdate karena sudah terkirim.",
				taskid,
			),
		})
		return
	}

	updateData, err := buildTaskUpdate(taskid, twaktu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	err = db.Model(&models.IoAntrianTaskid{}).
		Where("nobooking = ?", refKbo).
		Updates(updateData).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Gagal update data",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code"			: 200,
		"message"		: "Berhasil",
		"data"			: twaktu,
		"token"			: token,
	})
}

func getTaskValue(task models.IoAntrianTaskid, taskid int, mode string) *time.Time {
	switch mode {
	case "time":
		switch taskid {
		case 3:
			return task.Taskid3
		case 4:
			return task.Taskid4
		case 5:
			return task.Taskid5
		case 6:
			return task.Taskid6
		case 7:
			return task.Taskid7
		case 99:
			return task.Taskid99
		}
	case "send":
		switch taskid {
		case 3:
			return task.Taskid3Send
		case 4:
			return task.Taskid4Send
		case 5:
			return task.Taskid5Send
		case 6:
			return task.Taskid6Send
		case 7:
			return task.Taskid7Send
		case 99:
			return task.Taskid99Send
		}
	}
	return nil
}

func buildTaskUpdate(taskid int, t time.Time) (map[string]interface{}, error) {
	switch taskid {
	case 3:
		return map[string]interface{}{"taskid_3": t}, nil
	case 4:
		return map[string]interface{}{"taskid_4": t}, nil
	case 5:
		return map[string]interface{}{"taskid_5": t}, nil
	case 6:
		return map[string]interface{}{"taskid_6": t}, nil
	case 7:
		return map[string]interface{}{"taskid_7": t}, nil
	case 99:
		return map[string]interface{}{"taskid_99": t}, nil
	default:
		return nil, fmt.Errorf("taskid tidak valid")
	}
}
