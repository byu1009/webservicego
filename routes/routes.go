package routes

import (
	authController "webservicego/app/controllers/auth"
	profileController "webservicego/app/controllers/profile"
	"webservicego/app/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// ===== PUBLIC ROUTES =====
	r.POST("/login", authController.Login)

	// ===== PROTECTED ROUTES (1 GROUP MIDDLEWARE) =====
	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/profile", profileController.Show)
	}
}
