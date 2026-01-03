package auth

import (
	"net/http"
	"time"
	"webservicego/app/utils"
	"webservicego/config"

	"github.com/gin-gonic/gin"
)

func LoginData(c *gin.Context) {
	// 1️⃣ Ambil token dari header
	tokenHeader := c.GetHeader("Authorization")
	if tokenHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Unauthorized",
		})
		return
	}

	// 2️⃣ Decrypt token
	var payload utils.TokenPayload
	if err := utils.Decrypt(tokenHeader, &payload); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Unauthorized",
		})
		return
	}

	username := payload.Username
	expired := payload.Expired

	// 3️⃣ Validasi payload
	if username == "" || expired.IsZero() {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Unauthorized",
		})
		return
	}

	// 4️⃣ Cek expired token
	if time.Now().After(expired) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Token expired",
		})
		return
	}

	// 5️⃣ Ambil data user
	db := config.DBConnect()

	var user struct {
		ID          string `json:"id"`
		UserAccess  string `json:"user_access"`
		GroupAccess string `json:"group_access"`
	}

	if err := db.Table("io_user").
		Select("id, user_access, group_access").
		Where("id = ?", username).
		First(&user).Error; err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Unauthorized",
		})
		return
	}

	// 6️⃣ Tambahkan expired lama ke data
	responseData := gin.H{
		"id":           user.ID,
		"user_access":  user.UserAccess,
		"group_access": user.GroupAccess,
		"expired":      expired,
	}

	// 7️⃣ Generate token baru
	newToken, err := utils.GetToken(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed generate token",
		})
		return
	}

	// 8️⃣ Response
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Ok",
		"data":    responseData,
		"token":   newToken,
	})
}