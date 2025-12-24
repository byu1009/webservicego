package routes

import (
	authController "webservicego/app/controllers/auth"
	antrolController "webservicego/app/controllers/jkn"
	profileController "webservicego/app/controllers/profile"
	registrasiController "webservicego/app/controllers/registrasi"
	"webservicego/app/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// ===== PUBLIC ROUTES =====
	r.POST("/login", authController.Login)
	r.POST("/auth/check-username", authController.CheckUsername)
	r.POST("/check", authController.Check)
	r.GET("/auth/login-data", authController.LoginData)

	// ===== PROTECTED ROUTES (1 GROUP MIDDLEWARE) =====
	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/profile", profileController.Show)

		protected.POST("/registrasi/getdata", registrasiController.GetData)

		// SERVICE ANTROL BPJS KESEHATAN
		protected.GET("/antrol/ref-poli", antrolController.RefPoli)
		// protected.GET("/antrol/ref-dokter", antrolController.RefDokter)
	}
}
