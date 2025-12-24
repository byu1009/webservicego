package auth

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"webservicego/app/models"
	"webservicego/app/utils"
	"webservicego/config"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req struct {
		ID       string `json:"id" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Bind JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		// Mengembalikan error custom jika field tidak sesuai
		c.JSON(http.StatusOK, gin.H{
			"code"		: 400,
			"message"	: "Invalid request param, pastikan field yang berlaku ada",
		})
		return
	}

	db := config.DBConnect()
	var user models.User

	if err := db.Where("id = ? AND password = ?", req.ID, req.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code"		: 401,
			"message"	: "Invalid credentials",
		})
		return
	}

	// Ambil waktu expired token dari env
	expMin, _ := strconv.Atoi(os.Getenv("TOKEN_EXPIRED_MINUTES"))

	accessPayload := map[string]interface{}{
		"username"	: user.ID,
		"expired"	: time.Now().Add(time.Minute * time.Duration(expMin)),
		"type"		: "access",
	}

	accessToken, err := utils.Encrypt(accessPayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code"		: 500,
			"message"	: "Failed to generate access token",
			"error"		: err.Error(),
		})
		return
	}

	// Response HANYA access token
	c.JSON(http.StatusOK, gin.H{
		"code"		: 200,
		"message"	: "Berhasil mendapatkan token",
		"token"		: accessToken,
	})
}
