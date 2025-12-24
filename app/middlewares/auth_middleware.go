package middlewares

import (
	"net/http"
	"time"
	"webservicego/app/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		var payload utils.TokenPayload
		if err := utils.Decrypt(token, &payload); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "Invalid token",
			})
			c.Abort()
			return
		}

		if payload.Username == "" || payload.Expired.IsZero() {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "Invalid token payload",
			})
			c.Abort()
			return
		}

		if time.Now().After(payload.Expired) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "Token expired",
			})
			c.Abort()
			return
		}

		c.Set("username", payload.Username)
		c.Next()
	}
}
