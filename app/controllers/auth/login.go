package auth

import (
	"net/http"
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

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "Invalid request param",
		})
		return
	}

	db := config.DBConnect()
	var user models.User

	if err := db.Where("id = ? AND password = ?", req.ID, req.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Invalid credentials",
		})
		return
	}

	token, err := utils.GetToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Berhasil mendapatkan token",
		"token":   token,
	})
}
