package auth

import (
	"net/http"
	"time"
	"webservicego/app/utils"

	"github.com/gin-gonic/gin"
)

func Check(c *gin.Context) {
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

	// 4️⃣ Cek expired
	if time.Now().After(expired) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Token expired",
		})
		return
	}

	// 5️⃣ Token masih valid
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Ok check",
		"expired": expired.Format("2006-01-02 15:04:05"),
	})
}