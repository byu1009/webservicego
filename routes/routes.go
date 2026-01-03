package routes

import (
	authController "webservicego/app/controllers/auth"
	antrolController "webservicego/app/controllers/jkn"
	taskidController "webservicego/app/controllers/jkn/taskid"
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

		// REGISTRASI PENDAFTARAN
		protected.POST("/registrasi/getdata", registrasiController.GetData)
		protected.POST("/registrasi/postdata", registrasiController.PostData)
		protected.POST("/registrasi/addantrian", registrasiController.AddAntrian)
		protected.POST("/registrasi/batalantrian", registrasiController.BatalAntrian)
		protected.POST("/registrasi/checkin", registrasiController.Checkin)

		// TASKID INTERNAL
		protected.POST("/jkn/taskid/post", taskidController.PostTaskid)

		// SERVICE ANTROL BPJS KESEHATAN
		protected.GET("/antrol/ref-poli", antrolController.RefPoli)
		protected.GET("/antrol/ref-dokter", antrolController.RefDokter)
		protected.POST("/antrol/jadwal-dokter", antrolController.JadwalDokter)
		protected.GET("/antrol/ref-fingerprint", antrolController.RefFingerpoli)
		protected.POST("/antrol/ref-pasien-fingerprint", antrolController.RefPasienFingerpoli)
		protected.POST("/antrol/tambah-antrean", antrolController.TambahAntrean)
		protected.POST("/antrol/tambah-antrean-farmasi", antrolController.TambahAntreanFarmasi)
		protected.POST("/antrol/update-waktu-antrean", antrolController.UpdateWaktuAntrean)
		protected.POST("/antrol/batal-antrean", antrolController.BatalAntrean)

		protected.POST("/antrol/list-taskid", antrolController.ListTaskid)
		
		protected.POST("/antrol/antrean-per-tanggal", antrolController.AntreanPerTanggal)
		protected.POST("/antrol/antrean-per-kodebooking", antrolController.AntreanPerKodebooking)
		protected.GET("/antrol/belum-dilayani", antrolController.AntreanBelumDilayani)
		protected.POST("/antrol/belum-dilayani-detail", antrolController.AntreanBelumDilayaniDetail)
	}
}
