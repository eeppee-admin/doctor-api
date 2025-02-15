package route

import (
	"signin_and_signup/controller"

	"github.com/gin-gonic/gin"
)

// 定义患者相关的路由
func SetupPatientRoutes(r *gin.Engine) {
	patientGroup := r.Group("/api/patients")
	{
		patientGroup.GET(":id", controller.GetPatient)
		patientGroup.GET(":id/records", controller.GetMedicalRecords)
	}
}
