package profile

import (
	"net/http"
	"webservicego/app/models"
	"webservicego/app/utils"
	"webservicego/config"

	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	username := c.GetString("username")
	if username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	db := config.DBConnect()

	var user models.User
	if err := db.Where("id = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	token, err := utils.GetToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code"		: 200,
		"message"	: "Ok",
		"data"		: gin.H{
			"id":            user.ID,
			"nip":           user.NIP,
			"user_access":   user.UserAccess,
			"group_access":  user.GroupAccess,
			"last_activity": user.LastActivity,
		},
		"token":         token,
	})
}
