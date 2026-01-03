package auth

import (
	"net/http"
	"webservicego/app/models"
	"webservicego/config"

	"github.com/gin-gonic/gin"
)

func CheckUsername(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
	}

	if err := c.ShouldBindJSON(&req); err != nil || req.Username == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    204,
			"message": "Username harus diisi!",
		})
		return
	}

	username := req.Username

	db := config.DBConnect()

	var user models.User
	if err := db.First(&user, "id = ?", username).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Invalid credentials",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Ok",
	})
}